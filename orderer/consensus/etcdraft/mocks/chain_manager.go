// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/ZihuaZhang/fabric/orderer/common/types"
	"github.com/ZihuaZhang/fabric/orderer/consensus"
	"github.com/ZihuaZhang/fabric/orderer/consensus/etcdraft"
)

type ChainManager struct {
	CreateChainStub        func(string)
	createChainMutex       sync.RWMutex
	createChainArgsForCall []struct {
		arg1 string
	}
	GetConsensusChainStub        func(string) consensus.Chain
	getConsensusChainMutex       sync.RWMutex
	getConsensusChainArgsForCall []struct {
		arg1 string
	}
	getConsensusChainReturns struct {
		result1 consensus.Chain
	}
	getConsensusChainReturnsOnCall map[int]struct {
		result1 consensus.Chain
	}
	ReportConsensusRelationAndStatusMetricsStub        func(string, types.ConsensusRelation, types.Status)
	reportConsensusRelationAndStatusMetricsMutex       sync.RWMutex
	reportConsensusRelationAndStatusMetricsArgsForCall []struct {
		arg1 string
		arg2 types.ConsensusRelation
		arg3 types.Status
	}
	SwitchChainToFollowerStub        func(string)
	switchChainToFollowerMutex       sync.RWMutex
	switchChainToFollowerArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ChainManager) CreateChain(arg1 string) {
	fake.createChainMutex.Lock()
	fake.createChainArgsForCall = append(fake.createChainArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CreateChainStub
	fake.recordInvocation("CreateChain", []interface{}{arg1})
	fake.createChainMutex.Unlock()
	if stub != nil {
		fake.CreateChainStub(arg1)
	}
}

func (fake *ChainManager) CreateChainCallCount() int {
	fake.createChainMutex.RLock()
	defer fake.createChainMutex.RUnlock()
	return len(fake.createChainArgsForCall)
}

func (fake *ChainManager) CreateChainCalls(stub func(string)) {
	fake.createChainMutex.Lock()
	defer fake.createChainMutex.Unlock()
	fake.CreateChainStub = stub
}

func (fake *ChainManager) CreateChainArgsForCall(i int) string {
	fake.createChainMutex.RLock()
	defer fake.createChainMutex.RUnlock()
	argsForCall := fake.createChainArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChainManager) GetConsensusChain(arg1 string) consensus.Chain {
	fake.getConsensusChainMutex.Lock()
	ret, specificReturn := fake.getConsensusChainReturnsOnCall[len(fake.getConsensusChainArgsForCall)]
	fake.getConsensusChainArgsForCall = append(fake.getConsensusChainArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetConsensusChainStub
	fakeReturns := fake.getConsensusChainReturns
	fake.recordInvocation("GetConsensusChain", []interface{}{arg1})
	fake.getConsensusChainMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ChainManager) GetConsensusChainCallCount() int {
	fake.getConsensusChainMutex.RLock()
	defer fake.getConsensusChainMutex.RUnlock()
	return len(fake.getConsensusChainArgsForCall)
}

func (fake *ChainManager) GetConsensusChainCalls(stub func(string) consensus.Chain) {
	fake.getConsensusChainMutex.Lock()
	defer fake.getConsensusChainMutex.Unlock()
	fake.GetConsensusChainStub = stub
}

func (fake *ChainManager) GetConsensusChainArgsForCall(i int) string {
	fake.getConsensusChainMutex.RLock()
	defer fake.getConsensusChainMutex.RUnlock()
	argsForCall := fake.getConsensusChainArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChainManager) GetConsensusChainReturns(result1 consensus.Chain) {
	fake.getConsensusChainMutex.Lock()
	defer fake.getConsensusChainMutex.Unlock()
	fake.GetConsensusChainStub = nil
	fake.getConsensusChainReturns = struct {
		result1 consensus.Chain
	}{result1}
}

func (fake *ChainManager) GetConsensusChainReturnsOnCall(i int, result1 consensus.Chain) {
	fake.getConsensusChainMutex.Lock()
	defer fake.getConsensusChainMutex.Unlock()
	fake.GetConsensusChainStub = nil
	if fake.getConsensusChainReturnsOnCall == nil {
		fake.getConsensusChainReturnsOnCall = make(map[int]struct {
			result1 consensus.Chain
		})
	}
	fake.getConsensusChainReturnsOnCall[i] = struct {
		result1 consensus.Chain
	}{result1}
}

func (fake *ChainManager) ReportConsensusRelationAndStatusMetrics(arg1 string, arg2 types.ConsensusRelation, arg3 types.Status) {
	fake.reportConsensusRelationAndStatusMetricsMutex.Lock()
	fake.reportConsensusRelationAndStatusMetricsArgsForCall = append(fake.reportConsensusRelationAndStatusMetricsArgsForCall, struct {
		arg1 string
		arg2 types.ConsensusRelation
		arg3 types.Status
	}{arg1, arg2, arg3})
	stub := fake.ReportConsensusRelationAndStatusMetricsStub
	fake.recordInvocation("ReportConsensusRelationAndStatusMetrics", []interface{}{arg1, arg2, arg3})
	fake.reportConsensusRelationAndStatusMetricsMutex.Unlock()
	if stub != nil {
		fake.ReportConsensusRelationAndStatusMetricsStub(arg1, arg2, arg3)
	}
}

func (fake *ChainManager) ReportConsensusRelationAndStatusMetricsCallCount() int {
	fake.reportConsensusRelationAndStatusMetricsMutex.RLock()
	defer fake.reportConsensusRelationAndStatusMetricsMutex.RUnlock()
	return len(fake.reportConsensusRelationAndStatusMetricsArgsForCall)
}

func (fake *ChainManager) ReportConsensusRelationAndStatusMetricsCalls(stub func(string, types.ConsensusRelation, types.Status)) {
	fake.reportConsensusRelationAndStatusMetricsMutex.Lock()
	defer fake.reportConsensusRelationAndStatusMetricsMutex.Unlock()
	fake.ReportConsensusRelationAndStatusMetricsStub = stub
}

func (fake *ChainManager) ReportConsensusRelationAndStatusMetricsArgsForCall(i int) (string, types.ConsensusRelation, types.Status) {
	fake.reportConsensusRelationAndStatusMetricsMutex.RLock()
	defer fake.reportConsensusRelationAndStatusMetricsMutex.RUnlock()
	argsForCall := fake.reportConsensusRelationAndStatusMetricsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ChainManager) SwitchChainToFollower(arg1 string) {
	fake.switchChainToFollowerMutex.Lock()
	fake.switchChainToFollowerArgsForCall = append(fake.switchChainToFollowerArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SwitchChainToFollowerStub
	fake.recordInvocation("SwitchChainToFollower", []interface{}{arg1})
	fake.switchChainToFollowerMutex.Unlock()
	if stub != nil {
		fake.SwitchChainToFollowerStub(arg1)
	}
}

func (fake *ChainManager) SwitchChainToFollowerCallCount() int {
	fake.switchChainToFollowerMutex.RLock()
	defer fake.switchChainToFollowerMutex.RUnlock()
	return len(fake.switchChainToFollowerArgsForCall)
}

func (fake *ChainManager) SwitchChainToFollowerCalls(stub func(string)) {
	fake.switchChainToFollowerMutex.Lock()
	defer fake.switchChainToFollowerMutex.Unlock()
	fake.SwitchChainToFollowerStub = stub
}

func (fake *ChainManager) SwitchChainToFollowerArgsForCall(i int) string {
	fake.switchChainToFollowerMutex.RLock()
	defer fake.switchChainToFollowerMutex.RUnlock()
	argsForCall := fake.switchChainToFollowerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChainManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createChainMutex.RLock()
	defer fake.createChainMutex.RUnlock()
	fake.getConsensusChainMutex.RLock()
	defer fake.getConsensusChainMutex.RUnlock()
	fake.reportConsensusRelationAndStatusMetricsMutex.RLock()
	defer fake.reportConsensusRelationAndStatusMetricsMutex.RUnlock()
	fake.switchChainToFollowerMutex.RLock()
	defer fake.switchChainToFollowerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ChainManager) recordInvocation(key string, args []interface{}) {
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

var _ etcdraft.ChainManager = new(ChainManager)
