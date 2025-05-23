// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	net "net"

	mock "github.com/stretchr/testify/mock"

	netlink "github.com/vishvananda/netlink"
)

// NetlinkManager is an autogenerated mock type for the NetlinkManager type
type NetlinkManager struct {
	mock.Mock
}

// LinkByName provides a mock function with given fields: _a0
func (_m *NetlinkManager) LinkByName(_a0 string) (netlink.Link, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for LinkByName")
	}

	var r0 netlink.Link
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (netlink.Link, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) netlink.Link); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(netlink.Link)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkDelAltName provides a mock function with given fields: _a0, _a1
func (_m *NetlinkManager) LinkDelAltName(_a0 netlink.Link, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for LinkDelAltName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetDown provides a mock function with given fields: _a0
func (_m *NetlinkManager) LinkSetDown(_a0 netlink.Link) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for LinkSetDown")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetName provides a mock function with given fields: _a0, _a1
func (_m *NetlinkManager) LinkSetName(_a0 netlink.Link, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for LinkSetName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetNsFd provides a mock function with given fields: _a0, _a1
func (_m *NetlinkManager) LinkSetNsFd(_a0 netlink.Link, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for LinkSetNsFd")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetUp provides a mock function with given fields: _a0
func (_m *NetlinkManager) LinkSetUp(_a0 netlink.Link) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for LinkSetUp")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetVfNodeGUID provides a mock function with given fields: _a0, _a1, _a2
func (_m *NetlinkManager) LinkSetVfNodeGUID(_a0 netlink.Link, _a1 int, _a2 net.HardwareAddr) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for LinkSetVfNodeGUID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int, net.HardwareAddr) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetVfPortGUID provides a mock function with given fields: _a0, _a1, _a2
func (_m *NetlinkManager) LinkSetVfPortGUID(_a0 netlink.Link, _a1 int, _a2 net.HardwareAddr) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for LinkSetVfPortGUID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int, net.HardwareAddr) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetVfState provides a mock function with given fields: _a0, _a1, _a2
func (_m *NetlinkManager) LinkSetVfState(_a0 netlink.Link, _a1 int, _a2 uint32) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for LinkSetVfState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int, uint32) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewNetlinkManager creates a new instance of NetlinkManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNetlinkManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *NetlinkManager {
	mock := &NetlinkManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
