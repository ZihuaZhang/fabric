// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/ZihuaZhang/fabric-protos-go/orderer"
	"google.golang.org/grpc"
)

type ABClient struct {
	BroadcastStub        func(context.Context, ...grpc.CallOption) (orderer.AtomicBroadcast_BroadcastClient, error)
	broadcastMutex       sync.RWMutex
	broadcastArgsForCall []struct {
		arg1 context.Context
		arg2 []grpc.CallOption
	}
	broadcastReturns struct {
		result1 orderer.AtomicBroadcast_BroadcastClient
		result2 error
	}
	broadcastReturnsOnCall map[int]struct {
		result1 orderer.AtomicBroadcast_BroadcastClient
		result2 error
	}
	DeliverStub        func(context.Context, ...grpc.CallOption) (orderer.AtomicBroadcast_DeliverClient, error)
	deliverMutex       sync.RWMutex
	deliverArgsForCall []struct {
		arg1 context.Context
		arg2 []grpc.CallOption
	}
	deliverReturns struct {
		result1 orderer.AtomicBroadcast_DeliverClient
		result2 error
	}
	deliverReturnsOnCall map[int]struct {
		result1 orderer.AtomicBroadcast_DeliverClient
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ABClient) Broadcast(arg1 context.Context, arg2 ...grpc.CallOption) (orderer.AtomicBroadcast_BroadcastClient, error) {
	fake.broadcastMutex.Lock()
	ret, specificReturn := fake.broadcastReturnsOnCall[len(fake.broadcastArgsForCall)]
	fake.broadcastArgsForCall = append(fake.broadcastArgsForCall, struct {
		arg1 context.Context
		arg2 []grpc.CallOption
	}{arg1, arg2})
	stub := fake.BroadcastStub
	fakeReturns := fake.broadcastReturns
	fake.recordInvocation("Broadcast", []interface{}{arg1, arg2})
	fake.broadcastMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ABClient) BroadcastCallCount() int {
	fake.broadcastMutex.RLock()
	defer fake.broadcastMutex.RUnlock()
	return len(fake.broadcastArgsForCall)
}

func (fake *ABClient) BroadcastCalls(stub func(context.Context, ...grpc.CallOption) (orderer.AtomicBroadcast_BroadcastClient, error)) {
	fake.broadcastMutex.Lock()
	defer fake.broadcastMutex.Unlock()
	fake.BroadcastStub = stub
}

func (fake *ABClient) BroadcastArgsForCall(i int) (context.Context, []grpc.CallOption) {
	fake.broadcastMutex.RLock()
	defer fake.broadcastMutex.RUnlock()
	argsForCall := fake.broadcastArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ABClient) BroadcastReturns(result1 orderer.AtomicBroadcast_BroadcastClient, result2 error) {
	fake.broadcastMutex.Lock()
	defer fake.broadcastMutex.Unlock()
	fake.BroadcastStub = nil
	fake.broadcastReturns = struct {
		result1 orderer.AtomicBroadcast_BroadcastClient
		result2 error
	}{result1, result2}
}

func (fake *ABClient) BroadcastReturnsOnCall(i int, result1 orderer.AtomicBroadcast_BroadcastClient, result2 error) {
	fake.broadcastMutex.Lock()
	defer fake.broadcastMutex.Unlock()
	fake.BroadcastStub = nil
	if fake.broadcastReturnsOnCall == nil {
		fake.broadcastReturnsOnCall = make(map[int]struct {
			result1 orderer.AtomicBroadcast_BroadcastClient
			result2 error
		})
	}
	fake.broadcastReturnsOnCall[i] = struct {
		result1 orderer.AtomicBroadcast_BroadcastClient
		result2 error
	}{result1, result2}
}

func (fake *ABClient) Deliver(arg1 context.Context, arg2 ...grpc.CallOption) (orderer.AtomicBroadcast_DeliverClient, error) {
	fake.deliverMutex.Lock()
	ret, specificReturn := fake.deliverReturnsOnCall[len(fake.deliverArgsForCall)]
	fake.deliverArgsForCall = append(fake.deliverArgsForCall, struct {
		arg1 context.Context
		arg2 []grpc.CallOption
	}{arg1, arg2})
	stub := fake.DeliverStub
	fakeReturns := fake.deliverReturns
	fake.recordInvocation("Deliver", []interface{}{arg1, arg2})
	fake.deliverMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ABClient) DeliverCallCount() int {
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	return len(fake.deliverArgsForCall)
}

func (fake *ABClient) DeliverCalls(stub func(context.Context, ...grpc.CallOption) (orderer.AtomicBroadcast_DeliverClient, error)) {
	fake.deliverMutex.Lock()
	defer fake.deliverMutex.Unlock()
	fake.DeliverStub = stub
}

func (fake *ABClient) DeliverArgsForCall(i int) (context.Context, []grpc.CallOption) {
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	argsForCall := fake.deliverArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ABClient) DeliverReturns(result1 orderer.AtomicBroadcast_DeliverClient, result2 error) {
	fake.deliverMutex.Lock()
	defer fake.deliverMutex.Unlock()
	fake.DeliverStub = nil
	fake.deliverReturns = struct {
		result1 orderer.AtomicBroadcast_DeliverClient
		result2 error
	}{result1, result2}
}

func (fake *ABClient) DeliverReturnsOnCall(i int, result1 orderer.AtomicBroadcast_DeliverClient, result2 error) {
	fake.deliverMutex.Lock()
	defer fake.deliverMutex.Unlock()
	fake.DeliverStub = nil
	if fake.deliverReturnsOnCall == nil {
		fake.deliverReturnsOnCall = make(map[int]struct {
			result1 orderer.AtomicBroadcast_DeliverClient
			result2 error
		})
	}
	fake.deliverReturnsOnCall[i] = struct {
		result1 orderer.AtomicBroadcast_DeliverClient
		result2 error
	}{result1, result2}
}

func (fake *ABClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.broadcastMutex.RLock()
	defer fake.broadcastMutex.RUnlock()
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ABClient) recordInvocation(key string, args []interface{}) {
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
