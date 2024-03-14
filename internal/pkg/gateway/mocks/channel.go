// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/ZihuaZhang/fabric/common/channelconfig"
)

type Channel struct {
	BlockDataHashingStructureWidthStub        func() uint32
	blockDataHashingStructureWidthMutex       sync.RWMutex
	blockDataHashingStructureWidthArgsForCall []struct {
	}
	blockDataHashingStructureWidthReturns struct {
		result1 uint32
	}
	blockDataHashingStructureWidthReturnsOnCall map[int]struct {
		result1 uint32
	}
	CapabilitiesStub        func() channelconfig.ChannelCapabilities
	capabilitiesMutex       sync.RWMutex
	capabilitiesArgsForCall []struct {
	}
	capabilitiesReturns struct {
		result1 channelconfig.ChannelCapabilities
	}
	capabilitiesReturnsOnCall map[int]struct {
		result1 channelconfig.ChannelCapabilities
	}
	HashingAlgorithmStub        func() func(input []byte) []byte
	hashingAlgorithmMutex       sync.RWMutex
	hashingAlgorithmArgsForCall []struct {
	}
	hashingAlgorithmReturns struct {
		result1 func(input []byte) []byte
	}
	hashingAlgorithmReturnsOnCall map[int]struct {
		result1 func(input []byte) []byte
	}
	OrdererAddressesStub        func() []string
	ordererAddressesMutex       sync.RWMutex
	ordererAddressesArgsForCall []struct {
	}
	ordererAddressesReturns struct {
		result1 []string
	}
	ordererAddressesReturnsOnCall map[int]struct {
		result1 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Channel) BlockDataHashingStructureWidth() uint32 {
	fake.blockDataHashingStructureWidthMutex.Lock()
	ret, specificReturn := fake.blockDataHashingStructureWidthReturnsOnCall[len(fake.blockDataHashingStructureWidthArgsForCall)]
	fake.blockDataHashingStructureWidthArgsForCall = append(fake.blockDataHashingStructureWidthArgsForCall, struct {
	}{})
	stub := fake.BlockDataHashingStructureWidthStub
	fakeReturns := fake.blockDataHashingStructureWidthReturns
	fake.recordInvocation("BlockDataHashingStructureWidth", []interface{}{})
	fake.blockDataHashingStructureWidthMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Channel) BlockDataHashingStructureWidthCallCount() int {
	fake.blockDataHashingStructureWidthMutex.RLock()
	defer fake.blockDataHashingStructureWidthMutex.RUnlock()
	return len(fake.blockDataHashingStructureWidthArgsForCall)
}

func (fake *Channel) BlockDataHashingStructureWidthCalls(stub func() uint32) {
	fake.blockDataHashingStructureWidthMutex.Lock()
	defer fake.blockDataHashingStructureWidthMutex.Unlock()
	fake.BlockDataHashingStructureWidthStub = stub
}

func (fake *Channel) BlockDataHashingStructureWidthReturns(result1 uint32) {
	fake.blockDataHashingStructureWidthMutex.Lock()
	defer fake.blockDataHashingStructureWidthMutex.Unlock()
	fake.BlockDataHashingStructureWidthStub = nil
	fake.blockDataHashingStructureWidthReturns = struct {
		result1 uint32
	}{result1}
}

func (fake *Channel) BlockDataHashingStructureWidthReturnsOnCall(i int, result1 uint32) {
	fake.blockDataHashingStructureWidthMutex.Lock()
	defer fake.blockDataHashingStructureWidthMutex.Unlock()
	fake.BlockDataHashingStructureWidthStub = nil
	if fake.blockDataHashingStructureWidthReturnsOnCall == nil {
		fake.blockDataHashingStructureWidthReturnsOnCall = make(map[int]struct {
			result1 uint32
		})
	}
	fake.blockDataHashingStructureWidthReturnsOnCall[i] = struct {
		result1 uint32
	}{result1}
}

func (fake *Channel) Capabilities() channelconfig.ChannelCapabilities {
	fake.capabilitiesMutex.Lock()
	ret, specificReturn := fake.capabilitiesReturnsOnCall[len(fake.capabilitiesArgsForCall)]
	fake.capabilitiesArgsForCall = append(fake.capabilitiesArgsForCall, struct {
	}{})
	stub := fake.CapabilitiesStub
	fakeReturns := fake.capabilitiesReturns
	fake.recordInvocation("Capabilities", []interface{}{})
	fake.capabilitiesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Channel) CapabilitiesCallCount() int {
	fake.capabilitiesMutex.RLock()
	defer fake.capabilitiesMutex.RUnlock()
	return len(fake.capabilitiesArgsForCall)
}

func (fake *Channel) CapabilitiesCalls(stub func() channelconfig.ChannelCapabilities) {
	fake.capabilitiesMutex.Lock()
	defer fake.capabilitiesMutex.Unlock()
	fake.CapabilitiesStub = stub
}

func (fake *Channel) CapabilitiesReturns(result1 channelconfig.ChannelCapabilities) {
	fake.capabilitiesMutex.Lock()
	defer fake.capabilitiesMutex.Unlock()
	fake.CapabilitiesStub = nil
	fake.capabilitiesReturns = struct {
		result1 channelconfig.ChannelCapabilities
	}{result1}
}

func (fake *Channel) CapabilitiesReturnsOnCall(i int, result1 channelconfig.ChannelCapabilities) {
	fake.capabilitiesMutex.Lock()
	defer fake.capabilitiesMutex.Unlock()
	fake.CapabilitiesStub = nil
	if fake.capabilitiesReturnsOnCall == nil {
		fake.capabilitiesReturnsOnCall = make(map[int]struct {
			result1 channelconfig.ChannelCapabilities
		})
	}
	fake.capabilitiesReturnsOnCall[i] = struct {
		result1 channelconfig.ChannelCapabilities
	}{result1}
}

func (fake *Channel) HashingAlgorithm() func(input []byte) []byte {
	fake.hashingAlgorithmMutex.Lock()
	ret, specificReturn := fake.hashingAlgorithmReturnsOnCall[len(fake.hashingAlgorithmArgsForCall)]
	fake.hashingAlgorithmArgsForCall = append(fake.hashingAlgorithmArgsForCall, struct {
	}{})
	stub := fake.HashingAlgorithmStub
	fakeReturns := fake.hashingAlgorithmReturns
	fake.recordInvocation("HashingAlgorithm", []interface{}{})
	fake.hashingAlgorithmMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Channel) HashingAlgorithmCallCount() int {
	fake.hashingAlgorithmMutex.RLock()
	defer fake.hashingAlgorithmMutex.RUnlock()
	return len(fake.hashingAlgorithmArgsForCall)
}

func (fake *Channel) HashingAlgorithmCalls(stub func() func(input []byte) []byte) {
	fake.hashingAlgorithmMutex.Lock()
	defer fake.hashingAlgorithmMutex.Unlock()
	fake.HashingAlgorithmStub = stub
}

func (fake *Channel) HashingAlgorithmReturns(result1 func(input []byte) []byte) {
	fake.hashingAlgorithmMutex.Lock()
	defer fake.hashingAlgorithmMutex.Unlock()
	fake.HashingAlgorithmStub = nil
	fake.hashingAlgorithmReturns = struct {
		result1 func(input []byte) []byte
	}{result1}
}

func (fake *Channel) HashingAlgorithmReturnsOnCall(i int, result1 func(input []byte) []byte) {
	fake.hashingAlgorithmMutex.Lock()
	defer fake.hashingAlgorithmMutex.Unlock()
	fake.HashingAlgorithmStub = nil
	if fake.hashingAlgorithmReturnsOnCall == nil {
		fake.hashingAlgorithmReturnsOnCall = make(map[int]struct {
			result1 func(input []byte) []byte
		})
	}
	fake.hashingAlgorithmReturnsOnCall[i] = struct {
		result1 func(input []byte) []byte
	}{result1}
}

func (fake *Channel) OrdererAddresses() []string {
	fake.ordererAddressesMutex.Lock()
	ret, specificReturn := fake.ordererAddressesReturnsOnCall[len(fake.ordererAddressesArgsForCall)]
	fake.ordererAddressesArgsForCall = append(fake.ordererAddressesArgsForCall, struct {
	}{})
	stub := fake.OrdererAddressesStub
	fakeReturns := fake.ordererAddressesReturns
	fake.recordInvocation("OrdererAddresses", []interface{}{})
	fake.ordererAddressesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Channel) OrdererAddressesCallCount() int {
	fake.ordererAddressesMutex.RLock()
	defer fake.ordererAddressesMutex.RUnlock()
	return len(fake.ordererAddressesArgsForCall)
}

func (fake *Channel) OrdererAddressesCalls(stub func() []string) {
	fake.ordererAddressesMutex.Lock()
	defer fake.ordererAddressesMutex.Unlock()
	fake.OrdererAddressesStub = stub
}

func (fake *Channel) OrdererAddressesReturns(result1 []string) {
	fake.ordererAddressesMutex.Lock()
	defer fake.ordererAddressesMutex.Unlock()
	fake.OrdererAddressesStub = nil
	fake.ordererAddressesReturns = struct {
		result1 []string
	}{result1}
}

func (fake *Channel) OrdererAddressesReturnsOnCall(i int, result1 []string) {
	fake.ordererAddressesMutex.Lock()
	defer fake.ordererAddressesMutex.Unlock()
	fake.OrdererAddressesStub = nil
	if fake.ordererAddressesReturnsOnCall == nil {
		fake.ordererAddressesReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.ordererAddressesReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *Channel) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.blockDataHashingStructureWidthMutex.RLock()
	defer fake.blockDataHashingStructureWidthMutex.RUnlock()
	fake.capabilitiesMutex.RLock()
	defer fake.capabilitiesMutex.RUnlock()
	fake.hashingAlgorithmMutex.RLock()
	defer fake.hashingAlgorithmMutex.RUnlock()
	fake.ordererAddressesMutex.RLock()
	defer fake.ordererAddressesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Channel) recordInvocation(key string, args []interface{}) {
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
