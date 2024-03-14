// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/ZihuaZhang/fabric-protos-go/common"
	"github.com/ZihuaZhang/fabric/orderer/common/follower"
)

type BlockPullerFactory struct {
	BlockPullerStub        func(*common.Block, chan struct{}) (follower.ChannelPuller, error)
	blockPullerMutex       sync.RWMutex
	blockPullerArgsForCall []struct {
		arg1 *common.Block
		arg2 chan struct{}
	}
	blockPullerReturns struct {
		result1 follower.ChannelPuller
		result2 error
	}
	blockPullerReturnsOnCall map[int]struct {
		result1 follower.ChannelPuller
		result2 error
	}
	UpdateVerifierFromConfigBlockStub        func(*common.Block) error
	updateVerifierFromConfigBlockMutex       sync.RWMutex
	updateVerifierFromConfigBlockArgsForCall []struct {
		arg1 *common.Block
	}
	updateVerifierFromConfigBlockReturns struct {
		result1 error
	}
	updateVerifierFromConfigBlockReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BlockPullerFactory) BlockPuller(arg1 *common.Block, arg2 chan struct{}) (follower.ChannelPuller, error) {
	fake.blockPullerMutex.Lock()
	ret, specificReturn := fake.blockPullerReturnsOnCall[len(fake.blockPullerArgsForCall)]
	fake.blockPullerArgsForCall = append(fake.blockPullerArgsForCall, struct {
		arg1 *common.Block
		arg2 chan struct{}
	}{arg1, arg2})
	stub := fake.BlockPullerStub
	fakeReturns := fake.blockPullerReturns
	fake.recordInvocation("BlockPuller", []interface{}{arg1, arg2})
	fake.blockPullerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *BlockPullerFactory) BlockPullerCallCount() int {
	fake.blockPullerMutex.RLock()
	defer fake.blockPullerMutex.RUnlock()
	return len(fake.blockPullerArgsForCall)
}

func (fake *BlockPullerFactory) BlockPullerCalls(stub func(*common.Block, chan struct{}) (follower.ChannelPuller, error)) {
	fake.blockPullerMutex.Lock()
	defer fake.blockPullerMutex.Unlock()
	fake.BlockPullerStub = stub
}

func (fake *BlockPullerFactory) BlockPullerArgsForCall(i int) (*common.Block, chan struct{}) {
	fake.blockPullerMutex.RLock()
	defer fake.blockPullerMutex.RUnlock()
	argsForCall := fake.blockPullerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *BlockPullerFactory) BlockPullerReturns(result1 follower.ChannelPuller, result2 error) {
	fake.blockPullerMutex.Lock()
	defer fake.blockPullerMutex.Unlock()
	fake.BlockPullerStub = nil
	fake.blockPullerReturns = struct {
		result1 follower.ChannelPuller
		result2 error
	}{result1, result2}
}

func (fake *BlockPullerFactory) BlockPullerReturnsOnCall(i int, result1 follower.ChannelPuller, result2 error) {
	fake.blockPullerMutex.Lock()
	defer fake.blockPullerMutex.Unlock()
	fake.BlockPullerStub = nil
	if fake.blockPullerReturnsOnCall == nil {
		fake.blockPullerReturnsOnCall = make(map[int]struct {
			result1 follower.ChannelPuller
			result2 error
		})
	}
	fake.blockPullerReturnsOnCall[i] = struct {
		result1 follower.ChannelPuller
		result2 error
	}{result1, result2}
}

func (fake *BlockPullerFactory) UpdateVerifierFromConfigBlock(arg1 *common.Block) error {
	fake.updateVerifierFromConfigBlockMutex.Lock()
	ret, specificReturn := fake.updateVerifierFromConfigBlockReturnsOnCall[len(fake.updateVerifierFromConfigBlockArgsForCall)]
	fake.updateVerifierFromConfigBlockArgsForCall = append(fake.updateVerifierFromConfigBlockArgsForCall, struct {
		arg1 *common.Block
	}{arg1})
	stub := fake.UpdateVerifierFromConfigBlockStub
	fakeReturns := fake.updateVerifierFromConfigBlockReturns
	fake.recordInvocation("UpdateVerifierFromConfigBlock", []interface{}{arg1})
	fake.updateVerifierFromConfigBlockMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *BlockPullerFactory) UpdateVerifierFromConfigBlockCallCount() int {
	fake.updateVerifierFromConfigBlockMutex.RLock()
	defer fake.updateVerifierFromConfigBlockMutex.RUnlock()
	return len(fake.updateVerifierFromConfigBlockArgsForCall)
}

func (fake *BlockPullerFactory) UpdateVerifierFromConfigBlockCalls(stub func(*common.Block) error) {
	fake.updateVerifierFromConfigBlockMutex.Lock()
	defer fake.updateVerifierFromConfigBlockMutex.Unlock()
	fake.UpdateVerifierFromConfigBlockStub = stub
}

func (fake *BlockPullerFactory) UpdateVerifierFromConfigBlockArgsForCall(i int) *common.Block {
	fake.updateVerifierFromConfigBlockMutex.RLock()
	defer fake.updateVerifierFromConfigBlockMutex.RUnlock()
	argsForCall := fake.updateVerifierFromConfigBlockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *BlockPullerFactory) UpdateVerifierFromConfigBlockReturns(result1 error) {
	fake.updateVerifierFromConfigBlockMutex.Lock()
	defer fake.updateVerifierFromConfigBlockMutex.Unlock()
	fake.UpdateVerifierFromConfigBlockStub = nil
	fake.updateVerifierFromConfigBlockReturns = struct {
		result1 error
	}{result1}
}

func (fake *BlockPullerFactory) UpdateVerifierFromConfigBlockReturnsOnCall(i int, result1 error) {
	fake.updateVerifierFromConfigBlockMutex.Lock()
	defer fake.updateVerifierFromConfigBlockMutex.Unlock()
	fake.UpdateVerifierFromConfigBlockStub = nil
	if fake.updateVerifierFromConfigBlockReturnsOnCall == nil {
		fake.updateVerifierFromConfigBlockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateVerifierFromConfigBlockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *BlockPullerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.blockPullerMutex.RLock()
	defer fake.blockPullerMutex.RUnlock()
	fake.updateVerifierFromConfigBlockMutex.RLock()
	defer fake.updateVerifierFromConfigBlockMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *BlockPullerFactory) recordInvocation(key string, args []interface{}) {
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

var _ follower.BlockPullerFactory = new(BlockPullerFactory)
