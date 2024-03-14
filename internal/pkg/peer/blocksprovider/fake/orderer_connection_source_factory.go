// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/ZihuaZhang/fabric/common/flogging"
	"github.com/ZihuaZhang/fabric/internal/pkg/peer/blocksprovider"
	"github.com/ZihuaZhang/fabric/internal/pkg/peer/orderers"
)

type OrdererConnectionSourceFactory struct {
	CreateConnectionSourceStub        func(*flogging.FabricLogger) orderers.ConnectionSourcer
	createConnectionSourceMutex       sync.RWMutex
	createConnectionSourceArgsForCall []struct {
		arg1 *flogging.FabricLogger
	}
	createConnectionSourceReturns struct {
		result1 orderers.ConnectionSourcer
	}
	createConnectionSourceReturnsOnCall map[int]struct {
		result1 orderers.ConnectionSourcer
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OrdererConnectionSourceFactory) CreateConnectionSource(arg1 *flogging.FabricLogger) orderers.ConnectionSourcer {
	fake.createConnectionSourceMutex.Lock()
	ret, specificReturn := fake.createConnectionSourceReturnsOnCall[len(fake.createConnectionSourceArgsForCall)]
	fake.createConnectionSourceArgsForCall = append(fake.createConnectionSourceArgsForCall, struct {
		arg1 *flogging.FabricLogger
	}{arg1})
	stub := fake.CreateConnectionSourceStub
	fakeReturns := fake.createConnectionSourceReturns
	fake.recordInvocation("CreateConnectionSource", []interface{}{arg1})
	fake.createConnectionSourceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *OrdererConnectionSourceFactory) CreateConnectionSourceCallCount() int {
	fake.createConnectionSourceMutex.RLock()
	defer fake.createConnectionSourceMutex.RUnlock()
	return len(fake.createConnectionSourceArgsForCall)
}

func (fake *OrdererConnectionSourceFactory) CreateConnectionSourceCalls(stub func(*flogging.FabricLogger) orderers.ConnectionSourcer) {
	fake.createConnectionSourceMutex.Lock()
	defer fake.createConnectionSourceMutex.Unlock()
	fake.CreateConnectionSourceStub = stub
}

func (fake *OrdererConnectionSourceFactory) CreateConnectionSourceArgsForCall(i int) *flogging.FabricLogger {
	fake.createConnectionSourceMutex.RLock()
	defer fake.createConnectionSourceMutex.RUnlock()
	argsForCall := fake.createConnectionSourceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *OrdererConnectionSourceFactory) CreateConnectionSourceReturns(result1 orderers.ConnectionSourcer) {
	fake.createConnectionSourceMutex.Lock()
	defer fake.createConnectionSourceMutex.Unlock()
	fake.CreateConnectionSourceStub = nil
	fake.createConnectionSourceReturns = struct {
		result1 orderers.ConnectionSourcer
	}{result1}
}

func (fake *OrdererConnectionSourceFactory) CreateConnectionSourceReturnsOnCall(i int, result1 orderers.ConnectionSourcer) {
	fake.createConnectionSourceMutex.Lock()
	defer fake.createConnectionSourceMutex.Unlock()
	fake.CreateConnectionSourceStub = nil
	if fake.createConnectionSourceReturnsOnCall == nil {
		fake.createConnectionSourceReturnsOnCall = make(map[int]struct {
			result1 orderers.ConnectionSourcer
		})
	}
	fake.createConnectionSourceReturnsOnCall[i] = struct {
		result1 orderers.ConnectionSourcer
	}{result1}
}

func (fake *OrdererConnectionSourceFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createConnectionSourceMutex.RLock()
	defer fake.createConnectionSourceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OrdererConnectionSourceFactory) recordInvocation(key string, args []interface{}) {
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

var _ blocksprovider.OrdererConnectionSourceFactory = new(OrdererConnectionSourceFactory)