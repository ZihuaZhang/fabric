// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/ZihuaZhang/fabric-protos-go/common"
)

type ConfigGetter struct {
	GetCurrConfigStub        func(string) *common.Config
	getCurrConfigMutex       sync.RWMutex
	getCurrConfigArgsForCall []struct {
		arg1 string
	}
	getCurrConfigReturns struct {
		result1 *common.Config
	}
	getCurrConfigReturnsOnCall map[int]struct {
		result1 *common.Config
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ConfigGetter) GetCurrConfig(arg1 string) *common.Config {
	fake.getCurrConfigMutex.Lock()
	ret, specificReturn := fake.getCurrConfigReturnsOnCall[len(fake.getCurrConfigArgsForCall)]
	fake.getCurrConfigArgsForCall = append(fake.getCurrConfigArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetCurrConfig", []interface{}{arg1})
	fake.getCurrConfigMutex.Unlock()
	if fake.GetCurrConfigStub != nil {
		return fake.GetCurrConfigStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getCurrConfigReturns
	return fakeReturns.result1
}

func (fake *ConfigGetter) GetCurrConfigCallCount() int {
	fake.getCurrConfigMutex.RLock()
	defer fake.getCurrConfigMutex.RUnlock()
	return len(fake.getCurrConfigArgsForCall)
}

func (fake *ConfigGetter) GetCurrConfigCalls(stub func(string) *common.Config) {
	fake.getCurrConfigMutex.Lock()
	defer fake.getCurrConfigMutex.Unlock()
	fake.GetCurrConfigStub = stub
}

func (fake *ConfigGetter) GetCurrConfigArgsForCall(i int) string {
	fake.getCurrConfigMutex.RLock()
	defer fake.getCurrConfigMutex.RUnlock()
	argsForCall := fake.getCurrConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigGetter) GetCurrConfigReturns(result1 *common.Config) {
	fake.getCurrConfigMutex.Lock()
	defer fake.getCurrConfigMutex.Unlock()
	fake.GetCurrConfigStub = nil
	fake.getCurrConfigReturns = struct {
		result1 *common.Config
	}{result1}
}

func (fake *ConfigGetter) GetCurrConfigReturnsOnCall(i int, result1 *common.Config) {
	fake.getCurrConfigMutex.Lock()
	defer fake.getCurrConfigMutex.Unlock()
	fake.GetCurrConfigStub = nil
	if fake.getCurrConfigReturnsOnCall == nil {
		fake.getCurrConfigReturnsOnCall = make(map[int]struct {
			result1 *common.Config
		})
	}
	fake.getCurrConfigReturnsOnCall[i] = struct {
		result1 *common.Config
	}{result1}
}

func (fake *ConfigGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getCurrConfigMutex.RLock()
	defer fake.getCurrConfigMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ConfigGetter) recordInvocation(key string, args []interface{}) {
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
