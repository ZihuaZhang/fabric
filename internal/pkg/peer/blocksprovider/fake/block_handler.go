// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/ZihuaZhang/fabric-protos-go/common"
	"github.com/ZihuaZhang/fabric/internal/pkg/peer/blocksprovider"
)

type BlockHandler struct {
	HandleBlockStub        func(string, *common.Block) error
	handleBlockMutex       sync.RWMutex
	handleBlockArgsForCall []struct {
		arg1 string
		arg2 *common.Block
	}
	handleBlockReturns struct {
		result1 error
	}
	handleBlockReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BlockHandler) HandleBlock(arg1 string, arg2 *common.Block) error {
	fake.handleBlockMutex.Lock()
	ret, specificReturn := fake.handleBlockReturnsOnCall[len(fake.handleBlockArgsForCall)]
	fake.handleBlockArgsForCall = append(fake.handleBlockArgsForCall, struct {
		arg1 string
		arg2 *common.Block
	}{arg1, arg2})
	fake.recordInvocation("HandleBlock", []interface{}{arg1, arg2})
	fake.handleBlockMutex.Unlock()
	if fake.HandleBlockStub != nil {
		return fake.HandleBlockStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.handleBlockReturns
	return fakeReturns.result1
}

func (fake *BlockHandler) HandleBlockCallCount() int {
	fake.handleBlockMutex.RLock()
	defer fake.handleBlockMutex.RUnlock()
	return len(fake.handleBlockArgsForCall)
}

func (fake *BlockHandler) HandleBlockCalls(stub func(string, *common.Block) error) {
	fake.handleBlockMutex.Lock()
	defer fake.handleBlockMutex.Unlock()
	fake.HandleBlockStub = stub
}

func (fake *BlockHandler) HandleBlockArgsForCall(i int) (string, *common.Block) {
	fake.handleBlockMutex.RLock()
	defer fake.handleBlockMutex.RUnlock()
	argsForCall := fake.handleBlockArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *BlockHandler) HandleBlockReturns(result1 error) {
	fake.handleBlockMutex.Lock()
	defer fake.handleBlockMutex.Unlock()
	fake.HandleBlockStub = nil
	fake.handleBlockReturns = struct {
		result1 error
	}{result1}
}

func (fake *BlockHandler) HandleBlockReturnsOnCall(i int, result1 error) {
	fake.handleBlockMutex.Lock()
	defer fake.handleBlockMutex.Unlock()
	fake.HandleBlockStub = nil
	if fake.handleBlockReturnsOnCall == nil {
		fake.handleBlockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.handleBlockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *BlockHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleBlockMutex.RLock()
	defer fake.handleBlockMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *BlockHandler) recordInvocation(key string, args []interface{}) {
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

var _ blocksprovider.BlockHandler = new(BlockHandler)
