// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	model "testify/model"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// PriceProvider is an autogenerated mock type for the PriceProvider type
type PriceProvider struct {
	mock.Mock
}

// Latest provides a mock function with given fields:
func (_m *PriceProvider) Latest() (*model.TimeAndPrice, error) {
	ret := _m.Called()

	var r0 *model.TimeAndPrice
	var r1 error
	if rf, ok := ret.Get(0).(func() (*model.TimeAndPrice, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *model.TimeAndPrice); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TimeAndPrice)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: date
func (_m *PriceProvider) List(date time.Time) ([]*model.TimeAndPrice, error) {
	ret := _m.Called(date)

	var r0 []*model.TimeAndPrice
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time) ([]*model.TimeAndPrice, error)); ok {
		return rf(date)
	}
	if rf, ok := ret.Get(0).(func(time.Time) []*model.TimeAndPrice); ok {
		r0 = rf(date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TimeAndPrice)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Time) error); ok {
		r1 = rf(date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPriceProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewPriceProvider creates a new instance of PriceProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPriceProvider(t mockConstructorTestingTNewPriceProvider) *PriceProvider {
	mock := &PriceProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
