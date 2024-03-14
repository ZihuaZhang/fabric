// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/ZihuaZhang/fabric-protos-go/common"
	"github.com/ZihuaZhang/fabric-protos-go/orderer"
	"github.com/ZihuaZhang/fabric/common/ledger/blockledger"
)

type BlockReader struct {
	HeightStub        func() uint64
	heightMutex       sync.RWMutex
	heightArgsForCall []struct {
	}
	heightReturns struct {
		result1 uint64
	}
	heightReturnsOnCall map[int]struct {
		result1 uint64
	}
	IteratorStub        func(*orderer.SeekPosition) (blockledger.Iterator, uint64)
	iteratorMutex       sync.RWMutex
	iteratorArgsForCall []struct {
		arg1 *orderer.SeekPosition
	}
	iteratorReturns struct {
		result1 blockledger.Iterator
		result2 uint64
	}
	iteratorReturnsOnCall map[int]struct {
		result1 blockledger.Iterator
		result2 uint64
	}
	RetrieveBlockByNumberStub        func(uint64) (*common.Block, error)
	retrieveBlockByNumberMutex       sync.RWMutex
	retrieveBlockByNumberArgsForCall []struct {
		arg1 uint64
	}
	retrieveBlockByNumberReturns struct {
		result1 *common.Block
		result2 error
	}
	retrieveBlockByNumberReturnsOnCall map[int]struct {
		result1 *common.Block
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BlockReader) Height() uint64 {
	fake.heightMutex.Lock()
	ret, specificReturn := fake.heightReturnsOnCall[len(fake.heightArgsForCall)]
	fake.heightArgsForCall = append(fake.heightArgsForCall, struct {
	}{})
	stub := fake.HeightStub
	fakeReturns := fake.heightReturns
	fake.recordInvocation("Height", []interface{}{})
	fake.heightMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *BlockReader) HeightCallCount() int {
	fake.heightMutex.RLock()
	defer fake.heightMutex.RUnlock()
	return len(fake.heightArgsForCall)
}

func (fake *BlockReader) HeightCalls(stub func() uint64) {
	fake.heightMutex.Lock()
	defer fake.heightMutex.Unlock()
	fake.HeightStub = stub
}

func (fake *BlockReader) HeightReturns(result1 uint64) {
	fake.heightMutex.Lock()
	defer fake.heightMutex.Unlock()
	fake.HeightStub = nil
	fake.heightReturns = struct {
		result1 uint64
	}{result1}
}

func (fake *BlockReader) HeightReturnsOnCall(i int, result1 uint64) {
	fake.heightMutex.Lock()
	defer fake.heightMutex.Unlock()
	fake.HeightStub = nil
	if fake.heightReturnsOnCall == nil {
		fake.heightReturnsOnCall = make(map[int]struct {
			result1 uint64
		})
	}
	fake.heightReturnsOnCall[i] = struct {
		result1 uint64
	}{result1}
}

func (fake *BlockReader) Iterator(arg1 *orderer.SeekPosition) (blockledger.Iterator, uint64) {
	fake.iteratorMutex.Lock()
	ret, specificReturn := fake.iteratorReturnsOnCall[len(fake.iteratorArgsForCall)]
	fake.iteratorArgsForCall = append(fake.iteratorArgsForCall, struct {
		arg1 *orderer.SeekPosition
	}{arg1})
	stub := fake.IteratorStub
	fakeReturns := fake.iteratorReturns
	fake.recordInvocation("Iterator", []interface{}{arg1})
	fake.iteratorMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *BlockReader) IteratorCallCount() int {
	fake.iteratorMutex.RLock()
	defer fake.iteratorMutex.RUnlock()
	return len(fake.iteratorArgsForCall)
}

func (fake *BlockReader) IteratorCalls(stub func(*orderer.SeekPosition) (blockledger.Iterator, uint64)) {
	fake.iteratorMutex.Lock()
	defer fake.iteratorMutex.Unlock()
	fake.IteratorStub = stub
}

func (fake *BlockReader) IteratorArgsForCall(i int) *orderer.SeekPosition {
	fake.iteratorMutex.RLock()
	defer fake.iteratorMutex.RUnlock()
	argsForCall := fake.iteratorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *BlockReader) IteratorReturns(result1 blockledger.Iterator, result2 uint64) {
	fake.iteratorMutex.Lock()
	defer fake.iteratorMutex.Unlock()
	fake.IteratorStub = nil
	fake.iteratorReturns = struct {
		result1 blockledger.Iterator
		result2 uint64
	}{result1, result2}
}

func (fake *BlockReader) IteratorReturnsOnCall(i int, result1 blockledger.Iterator, result2 uint64) {
	fake.iteratorMutex.Lock()
	defer fake.iteratorMutex.Unlock()
	fake.IteratorStub = nil
	if fake.iteratorReturnsOnCall == nil {
		fake.iteratorReturnsOnCall = make(map[int]struct {
			result1 blockledger.Iterator
			result2 uint64
		})
	}
	fake.iteratorReturnsOnCall[i] = struct {
		result1 blockledger.Iterator
		result2 uint64
	}{result1, result2}
}

func (fake *BlockReader) RetrieveBlockByNumber(arg1 uint64) (*common.Block, error) {
	fake.retrieveBlockByNumberMutex.Lock()
	ret, specificReturn := fake.retrieveBlockByNumberReturnsOnCall[len(fake.retrieveBlockByNumberArgsForCall)]
	fake.retrieveBlockByNumberArgsForCall = append(fake.retrieveBlockByNumberArgsForCall, struct {
		arg1 uint64
	}{arg1})
	stub := fake.RetrieveBlockByNumberStub
	fakeReturns := fake.retrieveBlockByNumberReturns
	fake.recordInvocation("RetrieveBlockByNumber", []interface{}{arg1})
	fake.retrieveBlockByNumberMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *BlockReader) RetrieveBlockByNumberCallCount() int {
	fake.retrieveBlockByNumberMutex.RLock()
	defer fake.retrieveBlockByNumberMutex.RUnlock()
	return len(fake.retrieveBlockByNumberArgsForCall)
}

func (fake *BlockReader) RetrieveBlockByNumberCalls(stub func(uint64) (*common.Block, error)) {
	fake.retrieveBlockByNumberMutex.Lock()
	defer fake.retrieveBlockByNumberMutex.Unlock()
	fake.RetrieveBlockByNumberStub = stub
}

func (fake *BlockReader) RetrieveBlockByNumberArgsForCall(i int) uint64 {
	fake.retrieveBlockByNumberMutex.RLock()
	defer fake.retrieveBlockByNumberMutex.RUnlock()
	argsForCall := fake.retrieveBlockByNumberArgsForCall[i]
	return argsForCall.arg1
}

func (fake *BlockReader) RetrieveBlockByNumberReturns(result1 *common.Block, result2 error) {
	fake.retrieveBlockByNumberMutex.Lock()
	defer fake.retrieveBlockByNumberMutex.Unlock()
	fake.RetrieveBlockByNumberStub = nil
	fake.retrieveBlockByNumberReturns = struct {
		result1 *common.Block
		result2 error
	}{result1, result2}
}

func (fake *BlockReader) RetrieveBlockByNumberReturnsOnCall(i int, result1 *common.Block, result2 error) {
	fake.retrieveBlockByNumberMutex.Lock()
	defer fake.retrieveBlockByNumberMutex.Unlock()
	fake.RetrieveBlockByNumberStub = nil
	if fake.retrieveBlockByNumberReturnsOnCall == nil {
		fake.retrieveBlockByNumberReturnsOnCall = make(map[int]struct {
			result1 *common.Block
			result2 error
		})
	}
	fake.retrieveBlockByNumberReturnsOnCall[i] = struct {
		result1 *common.Block
		result2 error
	}{result1, result2}
}

func (fake *BlockReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.heightMutex.RLock()
	defer fake.heightMutex.RUnlock()
	fake.iteratorMutex.RLock()
	defer fake.iteratorMutex.RUnlock()
	fake.retrieveBlockByNumberMutex.RLock()
	defer fake.retrieveBlockByNumberMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *BlockReader) recordInvocation(key string, args []interface{}) {
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
