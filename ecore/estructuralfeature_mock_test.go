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

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func discardMockEStructuralFeature() {
	_ = assert.Equal
	_ = testing.Coverage
}

// TestMockEStructuralFeatureIsChangeable tests method IsChangeable
func TestMockEStructuralFeatureIsChangeable(t *testing.T) {
	o := &MockEStructuralFeature{}
	r := bool(true)
	o.On("IsChangeable").Once().Return(r)
	o.On("IsChangeable").Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsChangeable())
	assert.Equal(t, r, o.IsChangeable())
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureSetChangeable tests method SetChangeable
func TestMockEStructuralFeatureSetChangeable(t *testing.T) {
	o := &MockEStructuralFeature{}
	v := bool(true)
	o.On("SetChangeable", v).Once()
	o.SetChangeable(v)
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureGetDefaultValue tests method GetDefaultValue
func TestMockEStructuralFeatureGetDefaultValue(t *testing.T) {
	o := &MockEStructuralFeature{}
	r := interface{}(nil)
	o.On("GetDefaultValue").Once().Return(r)
	o.On("GetDefaultValue").Once().Return(func() interface{} {
		return r
	})
	assert.Equal(t, r, o.GetDefaultValue())
	assert.Equal(t, r, o.GetDefaultValue())
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureSetDefaultValue tests method SetDefaultValue
func TestMockEStructuralFeatureSetDefaultValue(t *testing.T) {
	o := &MockEStructuralFeature{}
	v := interface{}(nil)
	o.On("SetDefaultValue", v).Once()
	o.SetDefaultValue(v)
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureGetDefaultValueLiteral tests method GetDefaultValueLiteral
func TestMockEStructuralFeatureGetDefaultValueLiteral(t *testing.T) {
	o := &MockEStructuralFeature{}
	r := string("Test String")
	o.On("GetDefaultValueLiteral").Once().Return(r)
	o.On("GetDefaultValueLiteral").Once().Return(func() string {
		return r
	})
	assert.Equal(t, r, o.GetDefaultValueLiteral())
	assert.Equal(t, r, o.GetDefaultValueLiteral())
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureSetDefaultValueLiteral tests method SetDefaultValueLiteral
func TestMockEStructuralFeatureSetDefaultValueLiteral(t *testing.T) {
	o := &MockEStructuralFeature{}
	v := string("Test String")
	o.On("SetDefaultValueLiteral", v).Once()
	o.SetDefaultValueLiteral(v)
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureIsDerived tests method IsDerived
func TestMockEStructuralFeatureIsDerived(t *testing.T) {
	o := &MockEStructuralFeature{}
	r := bool(true)
	o.On("IsDerived").Once().Return(r)
	o.On("IsDerived").Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsDerived())
	assert.Equal(t, r, o.IsDerived())
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureSetDerived tests method SetDerived
func TestMockEStructuralFeatureSetDerived(t *testing.T) {
	o := &MockEStructuralFeature{}
	v := bool(true)
	o.On("SetDerived", v).Once()
	o.SetDerived(v)
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureGetEContainingClass tests method GetEContainingClass
func TestMockEStructuralFeatureGetEContainingClass(t *testing.T) {
	o := &MockEStructuralFeature{}
	r := new(MockEClass)
	o.On("GetEContainingClass").Once().Return(r)
	o.On("GetEContainingClass").Once().Return(func() EClass {
		return r
	})
	assert.Equal(t, r, o.GetEContainingClass())
	assert.Equal(t, r, o.GetEContainingClass())
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureGetFeatureID tests method GetFeatureID
func TestMockEStructuralFeatureGetFeatureID(t *testing.T) {
	o := &MockEStructuralFeature{}
	r := int(45)
	o.On("GetFeatureID").Once().Return(r)
	o.On("GetFeatureID").Once().Return(func() int {
		return r
	})
	assert.Equal(t, r, o.GetFeatureID())
	assert.Equal(t, r, o.GetFeatureID())
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureSetFeatureID tests method SetFeatureID
func TestMockEStructuralFeatureSetFeatureID(t *testing.T) {
	o := &MockEStructuralFeature{}
	v := int(45)
	o.On("SetFeatureID", v).Once()
	o.SetFeatureID(v)
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureIsTransient tests method IsTransient
func TestMockEStructuralFeatureIsTransient(t *testing.T) {
	o := &MockEStructuralFeature{}
	r := bool(true)
	o.On("IsTransient").Once().Return(r)
	o.On("IsTransient").Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsTransient())
	assert.Equal(t, r, o.IsTransient())
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureSetTransient tests method SetTransient
func TestMockEStructuralFeatureSetTransient(t *testing.T) {
	o := &MockEStructuralFeature{}
	v := bool(true)
	o.On("SetTransient", v).Once()
	o.SetTransient(v)
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureIsUnsettable tests method IsUnsettable
func TestMockEStructuralFeatureIsUnsettable(t *testing.T) {
	o := &MockEStructuralFeature{}
	r := bool(true)
	o.On("IsUnsettable").Once().Return(r)
	o.On("IsUnsettable").Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsUnsettable())
	assert.Equal(t, r, o.IsUnsettable())
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureSetUnsettable tests method SetUnsettable
func TestMockEStructuralFeatureSetUnsettable(t *testing.T) {
	o := &MockEStructuralFeature{}
	v := bool(true)
	o.On("SetUnsettable", v).Once()
	o.SetUnsettable(v)
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureIsVolatile tests method IsVolatile
func TestMockEStructuralFeatureIsVolatile(t *testing.T) {
	o := &MockEStructuralFeature{}
	r := bool(true)
	o.On("IsVolatile").Once().Return(r)
	o.On("IsVolatile").Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsVolatile())
	assert.Equal(t, r, o.IsVolatile())
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureSetVolatile tests method SetVolatile
func TestMockEStructuralFeatureSetVolatile(t *testing.T) {
	o := &MockEStructuralFeature{}
	v := bool(true)
	o.On("SetVolatile", v).Once()
	o.SetVolatile(v)
	o.AssertExpectations(t)
}

// TestMockEStructuralFeatureGetContainerClass tests method GetContainerClass
func TestMockEStructuralFeatureGetContainerClass(t *testing.T) {
	o := &MockEStructuralFeature{}
	r := reflect.Type(reflect.TypeOf(""))
	o.On("GetContainerClass").Return(r).Once()
	o.On("GetContainerClass").Return(func() reflect.Type {
		return r
	}).Once()
	assert.Equal(t, r, o.GetContainerClass())
	assert.Equal(t, r, o.GetContainerClass())
	o.AssertExpectations(t)
}
