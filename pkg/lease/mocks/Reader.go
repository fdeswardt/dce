// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import lease "github.com/Optum/dce/pkg/lease"
import mock "github.com/stretchr/testify/mock"

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

// Get provides a mock function with given fields: leaseID
func (_m *Reader) Get(leaseID string) (*lease.Lease, error) {
	ret := _m.Called(leaseID)

	var r0 *lease.Lease
	if rf, ok := ret.Get(0).(func(string) *lease.Lease); ok {
		r0 = rf(leaseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lease.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(leaseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByAccountIDAndPrincipalID provides a mock function with given fields: accountID, principalID
func (_m *Reader) GetByAccountIDAndPrincipalID(accountID string, principalID string) (*lease.Lease, error) {
	ret := _m.Called(accountID, principalID)

	var r0 *lease.Lease
	if rf, ok := ret.Get(0).(func(string, string) *lease.Lease); ok {
		r0 = rf(accountID, principalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lease.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(accountID, principalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0
func (_m *Reader) List(_a0 *lease.Lease) (*lease.Leases, error) {
	ret := _m.Called(_a0)

	var r0 *lease.Leases
	if rf, ok := ret.Get(0).(func(*lease.Lease) *lease.Leases); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lease.Leases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*lease.Lease) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
