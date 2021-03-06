// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

// Autogenerated from `make mockery` do not edit manually

package repository

import (
	domain "github.com/maxheckel/covid_county/covid_count/domain"
	mock "github.com/stretchr/testify/mock"

	sync "sync"

	time "time"
)

// MockRecord is an autogenerated mock type for the Record type
type MockRecord struct {
	mock.Mock
}

// ClearPreviousRecords provides a mock function with given fields:
func (_m *MockRecord) ClearPreviousRecords() error {
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
func (_m *MockRecord) CreateMultiple(records []domain.Record) error {
	ret := _m.Called(records)

	var r0 error
	if rf, ok := ret.Get(0).(func([]domain.Record) error); ok {
		r0 = rf(records)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MaxDate provides a mock function with given fields: col
func (_m *MockRecord) MaxDate(col string) (*time.Time, error) {
	ret := _m.Called(col)

	var r0 *time.Time
	if rf, ok := ret.Get(0).(func(string) *time.Time); ok {
		r0 = rf(col)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Time)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(col)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// insertAsync provides a mock function with given fields: records, wg
func (_m *MockRecord) insertAsync(records []domain.Record, wg *sync.WaitGroup) {
	_m.Called(records, wg)
}
