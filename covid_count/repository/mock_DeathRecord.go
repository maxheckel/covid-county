// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

// Autogenerated from `make mockery` do not edit manually

package repository

import (
	domain "github.com/maxheckel/covid_county/covid_count/domain"
	mock "github.com/stretchr/testify/mock"

	sync "sync"
)

// MockDeathRecord is an autogenerated mock type for the DeathRecord type
type MockDeathRecord struct {
	mock.Mock
}

// ClearPreviousMonthlyCountyDeaths provides a mock function with given fields:
func (_m *MockDeathRecord) ClearPreviousMonthlyCountyDeaths() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMultiple provides a mock function with given fields: records
func (_m *MockDeathRecord) CreateMultiple(records []domain.MonthlyCountyDeaths) error {
	ret := _m.Called(records)

	var r0 error
	if rf, ok := ret.Get(0).(func([]domain.MonthlyCountyDeaths) error); ok {
		r0 = rf(records)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetForCounty provides a mock function with given fields: county
func (_m *MockDeathRecord) GetForCounty(county string) ([]*domain.MonthlyCountyDeaths, error) {
	ret := _m.Called(county)

	var r0 []*domain.MonthlyCountyDeaths
	if rf, ok := ret.Get(0).(func(string) []*domain.MonthlyCountyDeaths); ok {
		r0 = rf(county)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.MonthlyCountyDeaths)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(county)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// insertAsync provides a mock function with given fields: records, wg
func (_m *MockDeathRecord) insertAsync(records []domain.MonthlyCountyDeaths, wg *sync.WaitGroup) {
	_m.Called(records, wg)
}
