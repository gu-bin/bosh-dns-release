// Code generated by counterfeiter. DO NOT EDIT.
package configfakes

import (
	"bosh-dns/dns/config"
	"sync"
)

type FakeStringShuffler struct {
	ShuffleStub        func(src []string) []string
	shuffleMutex       sync.RWMutex
	shuffleArgsForCall []struct {
		src []string
	}
	shuffleReturns struct {
		result1 []string
	}
	shuffleReturnsOnCall map[int]struct {
		result1 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStringShuffler) Shuffle(src []string) []string {
	var srcCopy []string
	if src != nil {
		srcCopy = make([]string, len(src))
		copy(srcCopy, src)
	}
	fake.shuffleMutex.Lock()
	ret, specificReturn := fake.shuffleReturnsOnCall[len(fake.shuffleArgsForCall)]
	fake.shuffleArgsForCall = append(fake.shuffleArgsForCall, struct {
		src []string
	}{srcCopy})
	fake.recordInvocation("Shuffle", []interface{}{srcCopy})
	fake.shuffleMutex.Unlock()
	if fake.ShuffleStub != nil {
		return fake.ShuffleStub(src)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.shuffleReturns.result1
}

func (fake *FakeStringShuffler) ShuffleCallCount() int {
	fake.shuffleMutex.RLock()
	defer fake.shuffleMutex.RUnlock()
	return len(fake.shuffleArgsForCall)
}

func (fake *FakeStringShuffler) ShuffleArgsForCall(i int) []string {
	fake.shuffleMutex.RLock()
	defer fake.shuffleMutex.RUnlock()
	return fake.shuffleArgsForCall[i].src
}

func (fake *FakeStringShuffler) ShuffleReturns(result1 []string) {
	fake.ShuffleStub = nil
	fake.shuffleReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeStringShuffler) ShuffleReturnsOnCall(i int, result1 []string) {
	fake.ShuffleStub = nil
	if fake.shuffleReturnsOnCall == nil {
		fake.shuffleReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.shuffleReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeStringShuffler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.shuffleMutex.RLock()
	defer fake.shuffleMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStringShuffler) recordInvocation(key string, args []interface{}) {
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

var _ config.StringShuffler = new(FakeStringShuffler)
