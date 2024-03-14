// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/ZihuaZhang/fabric/core/ledger/kvledger/txmgmt/statedb"
)

type NamespaceProvider struct {
	PossibleNamespacesStub        func(statedb.VersionedDB) ([]string, error)
	possibleNamespacesMutex       sync.RWMutex
	possibleNamespacesArgsForCall []struct {
		arg1 statedb.VersionedDB
	}
	possibleNamespacesReturns struct {
		result1 []string
		result2 error
	}
	possibleNamespacesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NamespaceProvider) PossibleNamespaces(arg1 statedb.VersionedDB) ([]string, error) {
	fake.possibleNamespacesMutex.Lock()
	ret, specificReturn := fake.possibleNamespacesReturnsOnCall[len(fake.possibleNamespacesArgsForCall)]
	fake.possibleNamespacesArgsForCall = append(fake.possibleNamespacesArgsForCall, struct {
		arg1 statedb.VersionedDB
	}{arg1})
	fake.recordInvocation("PossibleNamespaces", []interface{}{arg1})
	fake.possibleNamespacesMutex.Unlock()
	if fake.PossibleNamespacesStub != nil {
		return fake.PossibleNamespacesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.possibleNamespacesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *NamespaceProvider) PossibleNamespacesCallCount() int {
	fake.possibleNamespacesMutex.RLock()
	defer fake.possibleNamespacesMutex.RUnlock()
	return len(fake.possibleNamespacesArgsForCall)
}

func (fake *NamespaceProvider) PossibleNamespacesCalls(stub func(statedb.VersionedDB) ([]string, error)) {
	fake.possibleNamespacesMutex.Lock()
	defer fake.possibleNamespacesMutex.Unlock()
	fake.PossibleNamespacesStub = stub
}

func (fake *NamespaceProvider) PossibleNamespacesArgsForCall(i int) statedb.VersionedDB {
	fake.possibleNamespacesMutex.RLock()
	defer fake.possibleNamespacesMutex.RUnlock()
	argsForCall := fake.possibleNamespacesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *NamespaceProvider) PossibleNamespacesReturns(result1 []string, result2 error) {
	fake.possibleNamespacesMutex.Lock()
	defer fake.possibleNamespacesMutex.Unlock()
	fake.PossibleNamespacesStub = nil
	fake.possibleNamespacesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *NamespaceProvider) PossibleNamespacesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.possibleNamespacesMutex.Lock()
	defer fake.possibleNamespacesMutex.Unlock()
	fake.PossibleNamespacesStub = nil
	if fake.possibleNamespacesReturnsOnCall == nil {
		fake.possibleNamespacesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.possibleNamespacesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *NamespaceProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.possibleNamespacesMutex.RLock()
	defer fake.possibleNamespacesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *NamespaceProvider) recordInvocation(key string, args []interface{}) {
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

var _ statedb.NamespaceProvider = new(NamespaceProvider)
