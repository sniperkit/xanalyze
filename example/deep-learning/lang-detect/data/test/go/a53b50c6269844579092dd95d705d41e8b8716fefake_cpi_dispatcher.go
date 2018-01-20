// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/tscolari/bosh-c3pi/cpi"
)

type FakeCpiDispatcher struct {
	DispatchStub        func(methodName string, Arguments cpi.RequestArguments) (interface{}, error)
	dispatchMutex       sync.RWMutex
	dispatchArgsForCall []struct {
		methodName string
		Arguments  cpi.RequestArguments
	}
	dispatchReturns struct {
		result1 interface{}
		result2 error
	}
}

func (fake *FakeCpiDispatcher) Dispatch(methodName string, Arguments cpi.RequestArguments) (interface{}, error) {
	fake.dispatchMutex.Lock()
	fake.dispatchArgsForCall = append(fake.dispatchArgsForCall, struct {
		methodName string
		Arguments  cpi.RequestArguments
	}{methodName, Arguments})
	fake.dispatchMutex.Unlock()
	if fake.DispatchStub != nil {
		return fake.DispatchStub(methodName, Arguments)
	} else {
		return fake.dispatchReturns.result1, fake.dispatchReturns.result2
	}
}

func (fake *FakeCpiDispatcher) DispatchCallCount() int {
	fake.dispatchMutex.RLock()
	defer fake.dispatchMutex.RUnlock()
	return len(fake.dispatchArgsForCall)
}

func (fake *FakeCpiDispatcher) DispatchArgsForCall(i int) (string, cpi.RequestArguments) {
	fake.dispatchMutex.RLock()
	defer fake.dispatchMutex.RUnlock()
	return fake.dispatchArgsForCall[i].methodName, fake.dispatchArgsForCall[i].Arguments
}

func (fake *FakeCpiDispatcher) DispatchReturns(result1 interface{}, result2 error) {
	fake.DispatchStub = nil
	fake.dispatchReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

var _ cpi.CpiDispatcher = new(FakeCpiDispatcher)
