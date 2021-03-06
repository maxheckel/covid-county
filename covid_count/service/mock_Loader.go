// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

// Autogenerated from `make mockery` do not edit manually

package service

import (
	csv "encoding/csv"

	domain "github.com/maxheckel/covid_county/covid_count/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockLoader is an autogenerated mock type for the Loader type
type MockLoader struct {
	mock.Mock
}

// Load provides a mock function with given fields:
func (_m *MockLoader) Load() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// csvToRecords provides a mock function with given fields: csvReader, lineCount
func (_m *MockLoader) csvToRecords(csvReader *csv.Reader, lineCount int) []domain.Record {
	ret := _m.Called(csvReader, lineCount)

	var r0 []domain.Record
	if rf, ok := ret.Get(0).(func(*csv.Reader, int) []domain.Record); ok {
		r0 = rf(csvReader, lineCount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Record)
		}
	}

	return r0
}

// downloadFile provides a mock function with given fields: filepath, url
func (_m *MockLoader) downloadFile(filepath string, url string) error {
	ret := _m.Called(filepath, url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(filepath, url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// getLatestFile provides a mock function with given fields:
func (_m *MockLoader) getLatestFile() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// getLineCount provides a mock function with given fields: path
func (_m *MockLoader) getLineCount(path string) int {
	ret := _m.Called(path)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
