// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entity "customer-relationship-management/entity"

	mock "github.com/stretchr/testify/mock"
)

// ActorRepoInterface is an autogenerated mock type for the ActorRepoInterface type
type ActorRepoInterface struct {
	mock.Mock
}

// ActivateActorById provides a mock function with given fields: id
func (_m *ActorRepoInterface) ActivateActorById(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateActor provides a mock function with given fields: actor
func (_m *ActorRepoInterface) CreateActor(actor *entity.Actor) (entity.Actor, error) {
	ret := _m.Called(actor)

	var r0 entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Actor) (entity.Actor, error)); ok {
		return rf(actor)
	}
	if rf, ok := ret.Get(0).(func(*entity.Actor) entity.Actor); ok {
		r0 = rf(actor)
	} else {
		r0 = ret.Get(0).(entity.Actor)
	}

	if rf, ok := ret.Get(1).(func(*entity.Actor) error); ok {
		r1 = rf(actor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeactivateActorById provides a mock function with given fields: id
func (_m *ActorRepoInterface) DeactivateActorById(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteActorById provides a mock function with given fields: id
func (_m *ActorRepoInterface) DeleteActorById(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetActorById provides a mock function with given fields: id
func (_m *ActorRepoInterface) GetActorById(id uint) (entity.Actor, error) {
	ret := _m.Called(id)

	var r0 entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (entity.Actor, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) entity.Actor); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Actor)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllActor provides a mock function with given fields: page, username
func (_m *ActorRepoInterface) GetAllActor(page uint, username string) (uint, uint, int, uint, []entity.Actor, error) {
	ret := _m.Called(page, username)

	var r0 uint
	var r1 uint
	var r2 int
	var r3 uint
	var r4 []entity.Actor
	var r5 error
	if rf, ok := ret.Get(0).(func(uint, string) (uint, uint, int, uint, []entity.Actor, error)); ok {
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

	if rf, ok := ret.Get(4).(func(uint, string) []entity.Actor); ok {
		r4 = rf(page, username)
	} else {
		if ret.Get(4) != nil {
			r4 = ret.Get(4).([]entity.Actor)
		}
	}

	if rf, ok := ret.Get(5).(func(uint, string) error); ok {
		r5 = rf(page, username)
	} else {
		r5 = ret.Error(5)
	}

	return r0, r1, r2, r3, r4, r5
}

// LoginActor provides a mock function with given fields: actor
func (_m *ActorRepoInterface) LoginActor(actor *entity.Actor) (entity.Actor, error) {
	ret := _m.Called(actor)

	var r0 entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Actor) (entity.Actor, error)); ok {
		return rf(actor)
	}
	if rf, ok := ret.Get(0).(func(*entity.Actor) entity.Actor); ok {
		r0 = rf(actor)
	} else {
		r0 = ret.Get(0).(entity.Actor)
	}

	if rf, ok := ret.Get(1).(func(*entity.Actor) error); ok {
		r1 = rf(actor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateActorById provides a mock function with given fields: id, actor
func (_m *ActorRepoInterface) UpdateActorById(id uint, actor *entity.Actor) (entity.Actor, error) {
	ret := _m.Called(id, actor)

	var r0 entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, *entity.Actor) (entity.Actor, error)); ok {
		return rf(id, actor)
	}
	if rf, ok := ret.Get(0).(func(uint, *entity.Actor) entity.Actor); ok {
		r0 = rf(id, actor)
	} else {
		r0 = ret.Get(0).(entity.Actor)
	}

	if rf, ok := ret.Get(1).(func(uint, *entity.Actor) error); ok {
		r1 = rf(id, actor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewActorRepoInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewActorRepoInterface creates a new instance of ActorRepoInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewActorRepoInterface(t mockConstructorTestingTNewActorRepoInterface) *ActorRepoInterface {
	mock := &ActorRepoInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
