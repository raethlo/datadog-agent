// Code generated by mockery v1.0.0. DO NOT EDIT.

package checks

import (
	check "github.com/DataDog/datadog-agent/pkg/collector/check"
	compliance "github.com/DataDog/datadog-agent/pkg/compliance"

	mock "github.com/stretchr/testify/mock"
)

// MockBuilder is an autogenerated mock type for the Builder type
type MockBuilder struct {
	mock.Mock
}

// ChecksFromRule provides a mock function with given fields: meta, rule
func (_m *MockBuilder) ChecksFromRule(meta *compliance.SuiteMeta, rule *compliance.Rule) ([]check.Check, error) {
	ret := _m.Called(meta, rule)

	var r0 []check.Check
	if rf, ok := ret.Get(0).(func(*compliance.SuiteMeta, *compliance.Rule) []check.Check); ok {
		r0 = rf(meta, rule)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]check.Check)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*compliance.SuiteMeta, *compliance.Rule) error); ok {
		r1 = rf(meta, rule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
