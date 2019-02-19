// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	ldap "github.com/pivotalservices/cf-mgmt/ldap"
	user "github.com/pivotalservices/cf-mgmt/user"
)

type FakeLdapManager struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	GetUserByDNStub        func(string) (*ldap.User, error)
	getUserByDNMutex       sync.RWMutex
	getUserByDNArgsForCall []struct {
		arg1 string
	}
	getUserByDNReturns struct {
		result1 *ldap.User
		result2 error
	}
	getUserByDNReturnsOnCall map[int]struct {
		result1 *ldap.User
		result2 error
	}
	GetUserByIDStub        func(string) (*ldap.User, error)
	getUserByIDMutex       sync.RWMutex
	getUserByIDArgsForCall []struct {
		arg1 string
	}
	getUserByIDReturns struct {
		result1 *ldap.User
		result2 error
	}
	getUserByIDReturnsOnCall map[int]struct {
		result1 *ldap.User
		result2 error
	}
	GetUserDNsStub        func(string) ([]string, error)
	getUserDNsMutex       sync.RWMutex
	getUserDNsArgsForCall []struct {
		arg1 string
	}
	getUserDNsReturns struct {
		result1 []string
		result2 error
	}
	getUserDNsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLdapManager) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeLdapManager) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeLdapManager) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeLdapManager) GetUserByDN(arg1 string) (*ldap.User, error) {
	fake.getUserByDNMutex.Lock()
	ret, specificReturn := fake.getUserByDNReturnsOnCall[len(fake.getUserByDNArgsForCall)]
	fake.getUserByDNArgsForCall = append(fake.getUserByDNArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetUserByDN", []interface{}{arg1})
	fake.getUserByDNMutex.Unlock()
	if fake.GetUserByDNStub != nil {
		return fake.GetUserByDNStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getUserByDNReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLdapManager) GetUserByDNCallCount() int {
	fake.getUserByDNMutex.RLock()
	defer fake.getUserByDNMutex.RUnlock()
	return len(fake.getUserByDNArgsForCall)
}

func (fake *FakeLdapManager) GetUserByDNCalls(stub func(string) (*ldap.User, error)) {
	fake.getUserByDNMutex.Lock()
	defer fake.getUserByDNMutex.Unlock()
	fake.GetUserByDNStub = stub
}

func (fake *FakeLdapManager) GetUserByDNArgsForCall(i int) string {
	fake.getUserByDNMutex.RLock()
	defer fake.getUserByDNMutex.RUnlock()
	argsForCall := fake.getUserByDNArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLdapManager) GetUserByDNReturns(result1 *ldap.User, result2 error) {
	fake.getUserByDNMutex.Lock()
	defer fake.getUserByDNMutex.Unlock()
	fake.GetUserByDNStub = nil
	fake.getUserByDNReturns = struct {
		result1 *ldap.User
		result2 error
	}{result1, result2}
}

func (fake *FakeLdapManager) GetUserByDNReturnsOnCall(i int, result1 *ldap.User, result2 error) {
	fake.getUserByDNMutex.Lock()
	defer fake.getUserByDNMutex.Unlock()
	fake.GetUserByDNStub = nil
	if fake.getUserByDNReturnsOnCall == nil {
		fake.getUserByDNReturnsOnCall = make(map[int]struct {
			result1 *ldap.User
			result2 error
		})
	}
	fake.getUserByDNReturnsOnCall[i] = struct {
		result1 *ldap.User
		result2 error
	}{result1, result2}
}

func (fake *FakeLdapManager) GetUserByID(arg1 string) (*ldap.User, error) {
	fake.getUserByIDMutex.Lock()
	ret, specificReturn := fake.getUserByIDReturnsOnCall[len(fake.getUserByIDArgsForCall)]
	fake.getUserByIDArgsForCall = append(fake.getUserByIDArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetUserByID", []interface{}{arg1})
	fake.getUserByIDMutex.Unlock()
	if fake.GetUserByIDStub != nil {
		return fake.GetUserByIDStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getUserByIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLdapManager) GetUserByIDCallCount() int {
	fake.getUserByIDMutex.RLock()
	defer fake.getUserByIDMutex.RUnlock()
	return len(fake.getUserByIDArgsForCall)
}

func (fake *FakeLdapManager) GetUserByIDCalls(stub func(string) (*ldap.User, error)) {
	fake.getUserByIDMutex.Lock()
	defer fake.getUserByIDMutex.Unlock()
	fake.GetUserByIDStub = stub
}

func (fake *FakeLdapManager) GetUserByIDArgsForCall(i int) string {
	fake.getUserByIDMutex.RLock()
	defer fake.getUserByIDMutex.RUnlock()
	argsForCall := fake.getUserByIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLdapManager) GetUserByIDReturns(result1 *ldap.User, result2 error) {
	fake.getUserByIDMutex.Lock()
	defer fake.getUserByIDMutex.Unlock()
	fake.GetUserByIDStub = nil
	fake.getUserByIDReturns = struct {
		result1 *ldap.User
		result2 error
	}{result1, result2}
}

func (fake *FakeLdapManager) GetUserByIDReturnsOnCall(i int, result1 *ldap.User, result2 error) {
	fake.getUserByIDMutex.Lock()
	defer fake.getUserByIDMutex.Unlock()
	fake.GetUserByIDStub = nil
	if fake.getUserByIDReturnsOnCall == nil {
		fake.getUserByIDReturnsOnCall = make(map[int]struct {
			result1 *ldap.User
			result2 error
		})
	}
	fake.getUserByIDReturnsOnCall[i] = struct {
		result1 *ldap.User
		result2 error
	}{result1, result2}
}

func (fake *FakeLdapManager) GetUserDNs(arg1 string) ([]string, error) {
	fake.getUserDNsMutex.Lock()
	ret, specificReturn := fake.getUserDNsReturnsOnCall[len(fake.getUserDNsArgsForCall)]
	fake.getUserDNsArgsForCall = append(fake.getUserDNsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetUserDNs", []interface{}{arg1})
	fake.getUserDNsMutex.Unlock()
	if fake.GetUserDNsStub != nil {
		return fake.GetUserDNsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getUserDNsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLdapManager) GetUserDNsCallCount() int {
	fake.getUserDNsMutex.RLock()
	defer fake.getUserDNsMutex.RUnlock()
	return len(fake.getUserDNsArgsForCall)
}

func (fake *FakeLdapManager) GetUserDNsCalls(stub func(string) ([]string, error)) {
	fake.getUserDNsMutex.Lock()
	defer fake.getUserDNsMutex.Unlock()
	fake.GetUserDNsStub = stub
}

func (fake *FakeLdapManager) GetUserDNsArgsForCall(i int) string {
	fake.getUserDNsMutex.RLock()
	defer fake.getUserDNsMutex.RUnlock()
	argsForCall := fake.getUserDNsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLdapManager) GetUserDNsReturns(result1 []string, result2 error) {
	fake.getUserDNsMutex.Lock()
	defer fake.getUserDNsMutex.Unlock()
	fake.GetUserDNsStub = nil
	fake.getUserDNsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeLdapManager) GetUserDNsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.getUserDNsMutex.Lock()
	defer fake.getUserDNsMutex.Unlock()
	fake.GetUserDNsStub = nil
	if fake.getUserDNsReturnsOnCall == nil {
		fake.getUserDNsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getUserDNsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeLdapManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.getUserByDNMutex.RLock()
	defer fake.getUserByDNMutex.RUnlock()
	fake.getUserByIDMutex.RLock()
	defer fake.getUserByIDMutex.RUnlock()
	fake.getUserDNsMutex.RLock()
	defer fake.getUserDNsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLdapManager) recordInvocation(key string, args []interface{}) {
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

var _ user.LdapManager = new(FakeLdapManager)
