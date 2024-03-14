// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/ZihuaZhang/fabric/core/ledger"
)

type BlockListener struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	ReceiveBlockStub        func(*ledger.CommitNotification)
	receiveBlockMutex       sync.RWMutex
	receiveBlockArgsForCall []struct {
		arg1 *ledger.CommitNotification
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BlockListener) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		fake.CloseStub()
	}
}

func (fake *BlockListener) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *BlockListener) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *BlockListener) ReceiveBlock(arg1 *ledger.CommitNotification) {
	fake.receiveBlockMutex.Lock()
	fake.receiveBlockArgsForCall = append(fake.receiveBlockArgsForCall, struct {
		arg1 *ledger.CommitNotification
	}{arg1})
	stub := fake.ReceiveBlockStub
	fake.recordInvocation("ReceiveBlock", []interface{}{arg1})
	fake.receiveBlockMutex.Unlock()
	if stub != nil {
		fake.ReceiveBlockStub(arg1)
	}
}

func (fake *BlockListener) ReceiveBlockCallCount() int {
	fake.receiveBlockMutex.RLock()
	defer fake.receiveBlockMutex.RUnlock()
	return len(fake.receiveBlockArgsForCall)
}

func (fake *BlockListener) ReceiveBlockCalls(stub func(*ledger.CommitNotification)) {
	fake.receiveBlockMutex.Lock()
	defer fake.receiveBlockMutex.Unlock()
	fake.ReceiveBlockStub = stub
}

func (fake *BlockListener) ReceiveBlockArgsForCall(i int) *ledger.CommitNotification {
	fake.receiveBlockMutex.RLock()
	defer fake.receiveBlockMutex.RUnlock()
	argsForCall := fake.receiveBlockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *BlockListener) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.receiveBlockMutex.RLock()
	defer fake.receiveBlockMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *BlockListener) recordInvocation(key string, args []interface{}) {
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
