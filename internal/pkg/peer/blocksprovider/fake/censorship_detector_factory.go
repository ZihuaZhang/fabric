// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/ZihuaZhang/fabric/internal/pkg/peer/blocksprovider"
	"github.com/ZihuaZhang/fabric/internal/pkg/peer/orderers"
)

type CensorshipDetectorFactory struct {
	CreateStub        func(string, blocksprovider.UpdatableBlockVerifier, blocksprovider.DeliverClientRequester, blocksprovider.BlockProgressReporter, []*orderers.Endpoint, int, blocksprovider.TimeoutConfig) blocksprovider.CensorshipDetector
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
		arg2 blocksprovider.UpdatableBlockVerifier
		arg3 blocksprovider.DeliverClientRequester
		arg4 blocksprovider.BlockProgressReporter
		arg5 []*orderers.Endpoint
		arg6 int
		arg7 blocksprovider.TimeoutConfig
	}
	createReturns struct {
		result1 blocksprovider.CensorshipDetector
	}
	createReturnsOnCall map[int]struct {
		result1 blocksprovider.CensorshipDetector
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CensorshipDetectorFactory) Create(arg1 string, arg2 blocksprovider.UpdatableBlockVerifier, arg3 blocksprovider.DeliverClientRequester, arg4 blocksprovider.BlockProgressReporter, arg5 []*orderers.Endpoint, arg6 int, arg7 blocksprovider.TimeoutConfig) blocksprovider.CensorshipDetector {
	var arg5Copy []*orderers.Endpoint
	if arg5 != nil {
		arg5Copy = make([]*orderers.Endpoint, len(arg5))
		copy(arg5Copy, arg5)
	}
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
		arg2 blocksprovider.UpdatableBlockVerifier
		arg3 blocksprovider.DeliverClientRequester
		arg4 blocksprovider.BlockProgressReporter
		arg5 []*orderers.Endpoint
		arg6 int
		arg7 blocksprovider.TimeoutConfig
	}{arg1, arg2, arg3, arg4, arg5Copy, arg6, arg7})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3, arg4, arg5Copy, arg6, arg7})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *CensorshipDetectorFactory) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *CensorshipDetectorFactory) CreateCalls(stub func(string, blocksprovider.UpdatableBlockVerifier, blocksprovider.DeliverClientRequester, blocksprovider.BlockProgressReporter, []*orderers.Endpoint, int, blocksprovider.TimeoutConfig) blocksprovider.CensorshipDetector) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *CensorshipDetectorFactory) CreateArgsForCall(i int) (string, blocksprovider.UpdatableBlockVerifier, blocksprovider.DeliverClientRequester, blocksprovider.BlockProgressReporter, []*orderers.Endpoint, int, blocksprovider.TimeoutConfig) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *CensorshipDetectorFactory) CreateReturns(result1 blocksprovider.CensorshipDetector) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 blocksprovider.CensorshipDetector
	}{result1}
}

func (fake *CensorshipDetectorFactory) CreateReturnsOnCall(i int, result1 blocksprovider.CensorshipDetector) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 blocksprovider.CensorshipDetector
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 blocksprovider.CensorshipDetector
	}{result1}
}

func (fake *CensorshipDetectorFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CensorshipDetectorFactory) recordInvocation(key string, args []interface{}) {
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

var _ blocksprovider.CensorshipDetectorFactory = new(CensorshipDetectorFactory)
