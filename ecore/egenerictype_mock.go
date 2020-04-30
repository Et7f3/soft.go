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

type MockEGenericType struct {
	MockEObject
}

// GetEClassifier get the value of eClassifier
func (eGenericType *MockEGenericType) GetEClassifier() EClassifier {
	ret := eGenericType.Called()

	var r EClassifier
	if rf, ok := ret.Get(0).(func() EClassifier); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EClassifier)
		}
	}

	return r
}

// SetEClassifier provides mock implementation for setting the value of eClassifier
func (eGenericType *MockEGenericType) SetEClassifier(newEClassifier EClassifier) {
	eGenericType.Called(newEClassifier)
}

// GetELowerBound get the value of eLowerBound
func (eGenericType *MockEGenericType) GetELowerBound() EGenericType {
	ret := eGenericType.Called()

	var r EGenericType
	if rf, ok := ret.Get(0).(func() EGenericType); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EGenericType)
		}
	}

	return r
}

// SetELowerBound provides mock implementation for setting the value of eLowerBound
func (eGenericType *MockEGenericType) SetELowerBound(newELowerBound EGenericType) {
	eGenericType.Called(newELowerBound)
}

// GetERawType get the value of eRawType
func (eGenericType *MockEGenericType) GetERawType() EClassifier {
	ret := eGenericType.Called()

	var r EClassifier
	if rf, ok := ret.Get(0).(func() EClassifier); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EClassifier)
		}
	}

	return r
}

// GetETypeArguments get the value of eTypeArguments
func (eGenericType *MockEGenericType) GetETypeArguments() EList {
	ret := eGenericType.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

// GetETypeParameter get the value of eTypeParameter
func (eGenericType *MockEGenericType) GetETypeParameter() ETypeParameter {
	ret := eGenericType.Called()

	var r ETypeParameter
	if rf, ok := ret.Get(0).(func() ETypeParameter); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(ETypeParameter)
		}
	}

	return r
}

// SetETypeParameter provides mock implementation for setting the value of eTypeParameter
func (eGenericType *MockEGenericType) SetETypeParameter(newETypeParameter ETypeParameter) {
	eGenericType.Called(newETypeParameter)
}

// GetEUpperBound get the value of eUpperBound
func (eGenericType *MockEGenericType) GetEUpperBound() EGenericType {
	ret := eGenericType.Called()

	var r EGenericType
	if rf, ok := ret.Get(0).(func() EGenericType); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EGenericType)
		}
	}

	return r
}

// SetEUpperBound provides mock implementation for setting the value of eUpperBound
func (eGenericType *MockEGenericType) SetEUpperBound(newEUpperBound EGenericType) {
	eGenericType.Called(newEUpperBound)
}

// IsInstance provides mock implementation
func (eGenericType *MockEGenericType) IsInstance(object interface{}) bool {
	ret := eGenericType.Called(object)

	var r bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(bool)
		}
	}

	return r
}
