// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	pflag "github.com/spf13/pflag"
	mock "github.com/stretchr/testify/mock"
)

// FlagSetUtils is an autogenerated mock type for the FlagSetUtils type
type FlagSetUtils struct {
	mock.Mock
}

// GetLogFileName provides a mock function with given fields: _a0
func (_m *FlagSetUtils) GetLogFileName(_a0 *pflag.FlagSet) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}