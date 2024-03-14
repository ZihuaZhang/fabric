// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import api "github.com/ZihuaZhang/fabric/gossip/api"
import common "github.com/ZihuaZhang/fabric/gossip/common"
import discovery "github.com/ZihuaZhang/fabric/gossip/discovery"
import mock "github.com/stretchr/testify/mock"

// GossipSupport is an autogenerated mock type for the GossipSupport type
type GossipSupport struct {
	mock.Mock
}

// ChannelExists provides a mock function with given fields: channel
func (_m *GossipSupport) ChannelExists(channel string) bool {
	ret := _m.Called(channel)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(channel)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IdentityInfo provides a mock function with given fields:
func (_m *GossipSupport) IdentityInfo() api.PeerIdentitySet {
	ret := _m.Called()

	var r0 api.PeerIdentitySet
	if rf, ok := ret.Get(0).(func() api.PeerIdentitySet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.PeerIdentitySet)
		}
	}

	return r0
}

// Peers provides a mock function with given fields:
func (_m *GossipSupport) Peers() discovery.Members {
	ret := _m.Called()

	var r0 discovery.Members
	if rf, ok := ret.Get(0).(func() discovery.Members); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(discovery.Members)
		}
	}

	return r0
}

// PeersOfChannel provides a mock function with given fields: _a0
func (_m *GossipSupport) PeersOfChannel(_a0 common.ChannelID) discovery.Members {
	ret := _m.Called(_a0)

	var r0 discovery.Members
	if rf, ok := ret.Get(0).(func(common.ChannelID) discovery.Members); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(discovery.Members)
		}
	}

	return r0
}
