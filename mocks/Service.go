// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: user
func (_m *Service) Create(user interface{}) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *Service) Delete(id primitive.ObjectID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id, user
func (_m *Service) Get(id primitive.ObjectID, user interface{}) error {
	ret := _m.Called(id, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, interface{}) error); ok {
		r0 = rf(id, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: users
func (_m *Service) GetAll(users interface{}) error {
	ret := _m.Called(users)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(users)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: id, username, password
func (_m *Service) Update(id primitive.ObjectID, username string, password string) error {
	ret := _m.Called(id, username, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, string, string) error); ok {
		r0 = rf(id, username, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}