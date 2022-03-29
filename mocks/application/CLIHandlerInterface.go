// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	cli "github.com/urfave/cli/v2"
)

// CLIHandlerInterface is an autogenerated mock type for the CLIHandlerInterface type
type CLIHandlerInterface struct {
	mock.Mock
}

// CreateCLIHandler provides a mock function with given fields: c
func (_m *CLIHandlerInterface) CreateCLIHandler(c *cli.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*cli.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCLIHandler provides a mock function with given fields: c
func (_m *CLIHandlerInterface) DeleteCLIHandler(c *cli.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*cli.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCLIHandler provides a mock function with given fields: c
func (_m *CLIHandlerInterface) GetCLIHandler(c *cli.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*cli.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListCLIHandler provides a mock function with given fields: c
func (_m *CLIHandlerInterface) ListCLIHandler(c *cli.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*cli.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterCommands provides a mock function with given fields: app
func (_m *CLIHandlerInterface) RegisterCommands(app *cli.App) {
	_m.Called(app)
}

// UpdateCLIHandler provides a mock function with given fields: c
func (_m *CLIHandlerInterface) UpdateCLIHandler(c *cli.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*cli.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
