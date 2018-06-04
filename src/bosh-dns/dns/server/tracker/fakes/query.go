// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"bosh-dns/dns/server/criteria"
	"bosh-dns/dns/server/record"
	"sync"
)

type Query struct {
	FilterStub        func(criteria.Criteria, []record.Record) []record.Record
	filterMutex       sync.RWMutex
	filterArgsForCall []struct {
		arg1 criteria.Criteria
		arg2 []record.Record
	}
	filterReturns struct {
		result1 []record.Record
	}
	filterReturnsOnCall map[int]struct {
		result1 []record.Record
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Query) Filter(arg1 criteria.Criteria, arg2 []record.Record) []record.Record {
	var arg2Copy []record.Record
	if arg2 != nil {
		arg2Copy = make([]record.Record, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.filterMutex.Lock()
	ret, specificReturn := fake.filterReturnsOnCall[len(fake.filterArgsForCall)]
	fake.filterArgsForCall = append(fake.filterArgsForCall, struct {
		arg1 criteria.Criteria
		arg2 []record.Record
	}{arg1, arg2Copy})
	fake.recordInvocation("Filter", []interface{}{arg1, arg2Copy})
	fake.filterMutex.Unlock()
	if fake.FilterStub != nil {
		return fake.FilterStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.filterReturns.result1
}

func (fake *Query) FilterCallCount() int {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	return len(fake.filterArgsForCall)
}

func (fake *Query) FilterArgsForCall(i int) (criteria.Criteria, []record.Record) {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	return fake.filterArgsForCall[i].arg1, fake.filterArgsForCall[i].arg2
}

func (fake *Query) FilterReturns(result1 []record.Record) {
	fake.FilterStub = nil
	fake.filterReturns = struct {
		result1 []record.Record
	}{result1}
}

func (fake *Query) FilterReturnsOnCall(i int, result1 []record.Record) {
	fake.FilterStub = nil
	if fake.filterReturnsOnCall == nil {
		fake.filterReturnsOnCall = make(map[int]struct {
			result1 []record.Record
		})
	}
	fake.filterReturnsOnCall[i] = struct {
		result1 []record.Record
	}{result1}
}

func (fake *Query) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	return fake.invocations
}

func (fake *Query) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
