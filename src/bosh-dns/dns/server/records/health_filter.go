package records

import (
	"sync"
	"time"

	"code.cloudfoundry.org/workpool"

	"bosh-dns/dns/server/criteria"
	"bosh-dns/dns/server/healthiness"
	"bosh-dns/dns/server/record"
)

type healthFilter struct {
	nextFilter     reducer
	health         chan<- record.Host
	w              healthWatcher
	shouldTrack    bool
	domain         string
	filterWorkPool *workpool.WorkPool
}

type reducer interface {
	Filter(criteria.Criteria, []record.Record) []record.Record
}

type healthTracker interface {
	MonitorRecordHealth(ip, fqdn string)
}

type healthWatcher interface {
	HealthState(ip string) healthiness.HealthState
	Track(ip string)
	RunCheck(ip string)
}

func NewHealthFilter(nextFilter reducer, health chan<- record.Host, w healthWatcher, shouldTrack bool) healthFilter {
	wp, _ := workpool.NewWorkPool(1000)
	return healthFilter{
		nextFilter:     nextFilter,
		health:         health,
		w:              w,
		shouldTrack:    shouldTrack,
		filterWorkPool: wp,
	}
}

func (q *healthFilter) Filter(crit criteria.Criteria, recs []record.Record) []record.Record {
	records := q.nextFilter.Filter(crit, recs)
	var wg sync.WaitGroup

	q.addRecordsToWorkPool(crit, records, &wg)

	q.waitForWaitGroupOrTimeout(&wg)

	healthyRecords, unhealthyRecords, maybeHealthyRecords := q.sortRecords(records)

	healthStrategy := "0"
	if len(crit["s"]) > 0 {
		healthStrategy = crit["s"][0]
	}

	switch healthStrategy {
	case "1": // unhealthy ones
		return unhealthyRecords
	case "3": // healthy
		return healthyRecords
	case "4": // all
		return append(maybeHealthyRecords, unhealthyRecords...)
	default: // smart strategy
		if len(maybeHealthyRecords) == 0 {
			return unhealthyRecords
		}

		return maybeHealthyRecords
	}
}

func (q *healthFilter) addRecordsToWorkPool(criteria criteria.Criteria, records []record.Record, wg *sync.WaitGroup) {
	for _, r := range records {
		wg.Add(1)
		r := r
		q.filterWorkPool.Submit(func() {
			if q.shouldTrack {
				defer wg.Done()
				if fqdn, ok := criteria["fqdn"]; ok {
					q.health <- record.Host{IP: r.IP, FQDN: fqdn[0]}

					if len(criteria["y"]) > 0 {
						q.synchronousHealthCheck(criteria["y"][0], r.IP)
					}
				}
			}
		})
	}
}

func (q *healthFilter) waitForWaitGroupOrTimeout(wg *sync.WaitGroup) {
	timeout := time.After(1 * time.Second)
	success := make(chan bool)

	go func() {
		wg.Wait()
		success <- true
	}()

	for {
		select {
		case <-timeout:
			return
		case <-success:
			return
		}
	}
}

func (q *healthFilter) sortRecords(records []record.Record) (healthyRecords, unhealthyRecords, maybeHealthyRecords []record.Record) {
	var unknownRecords, uncheckedRecords []record.Record

	for _, r := range records {
		switch q.w.HealthState(r.IP) {
		case healthiness.StateHealthy:
			healthyRecords = append(healthyRecords, r)
		case healthiness.StateUnhealthy:
			unhealthyRecords = append(unhealthyRecords, r)
		case healthiness.StateUnknown:
			unknownRecords = append(unknownRecords, r)
		case healthiness.StateUnchecked:
			uncheckedRecords = append(uncheckedRecords, r)
		}
	}

	maybeHealthyRecords = append(healthyRecords, unknownRecords...)
	maybeHealthyRecords = append(maybeHealthyRecords, uncheckedRecords...)

	return healthyRecords, unhealthyRecords, maybeHealthyRecords
}

func (q *healthFilter) synchronousHealthCheck(strategy, ip string) {
	switch strategy {
	case "0":
		return
	case "1":
		if q.w.HealthState(ip) == healthiness.StateUnchecked {
			q.w.RunCheck(ip)
		}
	case "2":
		// q.runCheck(ip) to be implemented in a future story
		return
	}
}
