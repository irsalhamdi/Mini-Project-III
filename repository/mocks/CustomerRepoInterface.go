// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entity "customer-relationship-management/entity"

	mock "github.com/stretchr/testify/mock"
)

// CustomerRepoInterface is an autogenerated mock type for the CustomerRepoInterface type
type CustomerRepoInterface struct {
	mock.Mock
}

// CreateCustomer provides a mock function with given fields: customer
func (_m *CustomerRepoInterface) CreateCustomer(customer *entity.Customer) (entity.Customer, error) {
	ret := _m.Called(customer)

	var r0 entity.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Customer) (entity.Customer, error)); ok {
		return rf(customer)
	}
	if rf, ok := ret.Get(0).(func(*entity.Customer) entity.Customer); ok {
		r0 = rf(customer)
	} else {
		r0 = ret.Get(0).(entity.Customer)
	}

	if rf, ok := ret.Get(1).(func(*entity.Customer) error); ok {
		r1 = rf(customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCustomerById provides a mock function with given fields: id
func (_m *CustomerRepoInterface) DeleteCustomerById(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllCustomer provides a mock function with given fields: page, username
func (_m *CustomerRepoInterface) GetAllCustomer(page uint, username string) (uint, uint, int, uint, []entity.Customer, error) {
	ret := _m.Called(page, username)

	var r0 uint
	var r1 uint
	var r2 int
	var r3 uint
	var r4 []entity.Customer
	var r5 error
	if rf, ok := ret.Get(0).(func(uint, string) (uint, uint, int, uint, []entity.Customer, error)); ok {
		return rf(page, username)
	}
	if rf, ok := ret.Get(0).(func(uint, string) uint); ok {
		r0 = rf(page, username)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(uint, string) uint); ok {
		r1 = rf(page, username)
	} else {
		r1 = ret.Get(1).(uint)
	}

	if rf, ok := ret.Get(2).(func(uint, string) int); ok {
		r2 = rf(page, username)
	} else {
		r2 = ret.Get(2).(int)
	}

	if rf, ok := ret.Get(3).(func(uint, string) uint); ok {
		r3 = rf(page, username)
	} else {
		r3 = ret.Get(3).(uint)
	}

	if rf, ok := ret.Get(4).(func(uint, string) []entity.Customer); ok {
		r4 = rf(page, username)
	} else {
		if ret.Get(4) != nil {
			r4 = ret.Get(4).([]entity.Customer)
		}
	}

	if rf, ok := ret.Get(5).(func(uint, string) error); ok {
		r5 = rf(page, username)
	} else {
		r5 = ret.Error(5)
	}

	return r0, r1, r2, r3, r4, r5
}

// GetCustomerById provides a mock function with given fields: id
func (_m *CustomerRepoInterface) GetCustomerById(id uint) (entity.Customer, error) {
	ret := _m.Called(id)

	var r0 entity.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (entity.Customer, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) entity.Customer); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Customer)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCustomerById provides a mock function with given fields: id, customer
func (_m *CustomerRepoInterface) UpdateCustomerById(id uint, customer *entity.Customer) (entity.Customer, error) {
	ret := _m.Called(id, customer)

	var r0 entity.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, *entity.Customer) (entity.Customer, error)); ok {
		return rf(id, customer)
	}
	if rf, ok := ret.Get(0).(func(uint, *entity.Customer) entity.Customer); ok {
		r0 = rf(id, customer)
	} else {
		r0 = ret.Get(0).(entity.Customer)
	}

	if rf, ok := ret.Get(1).(func(uint, *entity.Customer) error); ok {
		r1 = rf(id, customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCustomerRepoInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewCustomerRepoInterface creates a new instance of CustomerRepoInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCustomerRepoInterface(t mockConstructorTestingTNewCustomerRepoInterface) *CustomerRepoInterface {
	mock := &CustomerRepoInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
