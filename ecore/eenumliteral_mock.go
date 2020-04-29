// *****************************************************************************
//
// This file is part of a MASA library or program.
// Refer to the included end-user license agreement for restrictions.
//
// Copyright (c) 2020 MASA Group
//
// *****************************************************************************

// *****************************************************************************
//
// Warning: This file was generated by soft.generator.go Generator
//
// *****************************************************************************

package ecore

type MockEEnumLiteral struct {
	MockENamedElement
}

// GetEEnum get the value of eEnum
func (eEnumLiteral *MockEEnumLiteral) GetEEnum() EEnum {
	ret := eEnumLiteral.Called()

	var r EEnum
	if rf, ok := ret.Get(0).(func() EEnum); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EEnum)
		}
	}

	return r
}

// GetInstance get the value of instance
func (eEnumLiteral *MockEEnumLiteral) GetInstance() interface{} {
	ret := eEnumLiteral.Called()

	var r interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(interface{})
		}
	}

	return r
}

// SetInstance provides mock implementation for setting the value of instance
func (eEnumLiteral *MockEEnumLiteral) SetInstance(newInstance interface{}) {
	eEnumLiteral.Called(newInstance)
}

// GetLiteral get the value of literal
func (eEnumLiteral *MockEEnumLiteral) GetLiteral() string {
	ret := eEnumLiteral.Called()

	var r string
	if rf, ok := ret.Get(0).(func() string); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(string)
		}
	}

	return r
}

// SetLiteral provides mock implementation for setting the value of literal
func (eEnumLiteral *MockEEnumLiteral) SetLiteral(newLiteral string) {
	eEnumLiteral.Called(newLiteral)
}

// GetValue get the value of value
func (eEnumLiteral *MockEEnumLiteral) GetValue() int {
	ret := eEnumLiteral.Called()

	var r int
	if rf, ok := ret.Get(0).(func() int); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(int)
		}
	}

	return r
}

// SetValue provides mock implementation for setting the value of value
func (eEnumLiteral *MockEEnumLiteral) SetValue(newValue int) {
	eEnumLiteral.Called(newValue)
}
