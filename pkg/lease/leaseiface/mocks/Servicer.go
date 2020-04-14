// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import lease "github.com/Optum/dce/pkg/lease"

import mock "github.com/stretchr/testify/mock"

// Servicer is an autogenerated mock type for the Servicer type
type Servicer struct {
	mock.Mock
}

// Create provides a mock function with given fields: data
func (_m *Servicer) Create(data *lease.Lease) (*lease.Lease, error) {
	ret := _m.Called(data)

	var r0 *lease.Lease
	if rf, ok := ret.Get(0).(func(*lease.Lease) *lease.Lease); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lease.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*lease.Lease) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ID
func (_m *Servicer) Delete(ID string) (*lease.Lease, error) {
	ret := _m.Called(ID)

	var r0 *lease.Lease
	if rf, ok := ret.Get(0).(func(string) *lease.Lease); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lease.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ID
func (_m *Servicer) Get(ID string) (*lease.Lease, error) {
	ret := _m.Called(ID)

	var r0 *lease.Lease
	if rf, ok := ret.Get(0).(func(string) *lease.Lease); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lease.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByAccountIDAndPrincipalID provides a mock function with given fields: accountID, principalID
func (_m *Servicer) GetByAccountIDAndPrincipalID(accountID string, principalID string) (*lease.Lease, error) {
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

// List provides a mock function with given fields: query
func (_m *Servicer) List(query *lease.Lease) (*lease.Leases, error) {
	ret := _m.Called(query)

	var r0 *lease.Leases
	if rf, ok := ret.Get(0).(func(*lease.Lease) *lease.Leases); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lease.Leases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*lease.Lease) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPages provides a mock function with given fields: query, fn
func (_m *Servicer) ListPages(query *lease.Lease, fn func(*lease.Leases) bool) error {
	ret := _m.Called(query, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(*lease.Lease, func(*lease.Leases) bool) error); ok {
		r0 = rf(query, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
