// Code generated by counterfeiter. DO NOT EDIT.
package handlersfakes

import (
	"bosh-dns/dns/server/handlers"
	"sync"

	"github.com/miekg/dns"
)

type FakeServerMux struct {
	HandleStub        func(pattern string, handler dns.Handler)
	handleMutex       sync.RWMutex
	handleArgsForCall []struct {
		pattern string
		handler dns.Handler
	}
	HandleRemoveStub        func(pattern string)
	handleRemoveMutex       sync.RWMutex
	handleRemoveArgsForCall []struct {
		pattern string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServerMux) Handle(pattern string, handler dns.Handler) {
	fake.handleMutex.Lock()
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct {
		pattern string
		handler dns.Handler
	}{pattern, handler})
	fake.recordInvocation("Handle", []interface{}{pattern, handler})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		fake.HandleStub(pattern, handler)
	}
}

func (fake *FakeServerMux) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeServerMux) HandleArgsForCall(i int) (string, dns.Handler) {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return fake.handleArgsForCall[i].pattern, fake.handleArgsForCall[i].handler
}

func (fake *FakeServerMux) HandleRemove(pattern string) {
	fake.handleRemoveMutex.Lock()
	fake.handleRemoveArgsForCall = append(fake.handleRemoveArgsForCall, struct {
		pattern string
	}{pattern})
	fake.recordInvocation("HandleRemove", []interface{}{pattern})
	fake.handleRemoveMutex.Unlock()
	if fake.HandleRemoveStub != nil {
		fake.HandleRemoveStub(pattern)
	}
}

func (fake *FakeServerMux) HandleRemoveCallCount() int {
	fake.handleRemoveMutex.RLock()
	defer fake.handleRemoveMutex.RUnlock()
	return len(fake.handleRemoveArgsForCall)
}

func (fake *FakeServerMux) HandleRemoveArgsForCall(i int) string {
	fake.handleRemoveMutex.RLock()
	defer fake.handleRemoveMutex.RUnlock()
	return fake.handleRemoveArgsForCall[i].pattern
}

func (fake *FakeServerMux) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	fake.handleRemoveMutex.RLock()
	defer fake.handleRemoveMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeServerMux) recordInvocation(key string, args []interface{}) {
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

var _ handlers.ServerMux = new(FakeServerMux)
