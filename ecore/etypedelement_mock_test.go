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
	"testing"
)

func discardMockETypedElement() {
	_ = assert.Equal
	_ = testing.Coverage
}

// TestMockETypedElementGetEType tests method GetEType
func TestMockETypedElementGetEType(t *testing.T) {
	o := &MockETypedElement{}
	r := new(MockEClassifier)
	o.On("GetEType").Once().Return(r)
	o.On("GetEType").Once().Return(func() EClassifier {
		return r
	})
	assert.Equal(t, r, o.GetEType())
	assert.Equal(t, r, o.GetEType())
	o.AssertExpectations(t)
}

// TestMockETypedElementSetEType tests method SetEType
func TestMockETypedElementSetEType(t *testing.T) {
	o := &MockETypedElement{}
	v := new(MockEClassifier)
	o.On("SetEType", v).Once()
	o.SetEType(v)
	o.AssertExpectations(t)
}

// TestMockETypedElementUnsetEType tests method UnsetEType
func TestMockETypedElementUnsetEType(t *testing.T) {
	o := &MockETypedElement{}
	o.On("UnsetEType").Once()
	o.UnsetEType()
	o.AssertExpectations(t)
}

// TestMockETypedElementGetLowerBound tests method GetLowerBound
func TestMockETypedElementGetLowerBound(t *testing.T) {
	o := &MockETypedElement{}
	r := int(45)
	o.On("GetLowerBound").Once().Return(r)
	o.On("GetLowerBound").Once().Return(func() int {
		return r
	})
	assert.Equal(t, r, o.GetLowerBound())
	assert.Equal(t, r, o.GetLowerBound())
	o.AssertExpectations(t)
}

// TestMockETypedElementSetLowerBound tests method SetLowerBound
func TestMockETypedElementSetLowerBound(t *testing.T) {
	o := &MockETypedElement{}
	v := int(45)
	o.On("SetLowerBound", v).Once()
	o.SetLowerBound(v)
	o.AssertExpectations(t)
}

// TestMockETypedElementIsMany tests method IsMany
func TestMockETypedElementIsMany(t *testing.T) {
	o := &MockETypedElement{}
	r := bool(true)
	o.On("IsMany").Once().Return(r)
	o.On("IsMany").Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsMany())
	assert.Equal(t, r, o.IsMany())
	o.AssertExpectations(t)
}

// TestMockETypedElementIsOrdered tests method IsOrdered
func TestMockETypedElementIsOrdered(t *testing.T) {
	o := &MockETypedElement{}
	r := bool(true)
	o.On("IsOrdered").Once().Return(r)
	o.On("IsOrdered").Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsOrdered())
	assert.Equal(t, r, o.IsOrdered())
	o.AssertExpectations(t)
}

// TestMockETypedElementSetOrdered tests method SetOrdered
func TestMockETypedElementSetOrdered(t *testing.T) {
	o := &MockETypedElement{}
	v := bool(true)
	o.On("SetOrdered", v).Once()
	o.SetOrdered(v)
	o.AssertExpectations(t)
}

// TestMockETypedElementIsRequired tests method IsRequired
func TestMockETypedElementIsRequired(t *testing.T) {
	o := &MockETypedElement{}
	r := bool(true)
	o.On("IsRequired").Once().Return(r)
	o.On("IsRequired").Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsRequired())
	assert.Equal(t, r, o.IsRequired())
	o.AssertExpectations(t)
}

// TestMockETypedElementIsUnique tests method IsUnique
func TestMockETypedElementIsUnique(t *testing.T) {
	o := &MockETypedElement{}
	r := bool(true)
	o.On("IsUnique").Once().Return(r)
	o.On("IsUnique").Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsUnique())
	assert.Equal(t, r, o.IsUnique())
	o.AssertExpectations(t)
}

// TestMockETypedElementSetUnique tests method SetUnique
func TestMockETypedElementSetUnique(t *testing.T) {
	o := &MockETypedElement{}
	v := bool(true)
	o.On("SetUnique", v).Once()
	o.SetUnique(v)
	o.AssertExpectations(t)
}

// TestMockETypedElementGetUpperBound tests method GetUpperBound
func TestMockETypedElementGetUpperBound(t *testing.T) {
	o := &MockETypedElement{}
	r := int(45)
	o.On("GetUpperBound").Once().Return(r)
	o.On("GetUpperBound").Once().Return(func() int {
		return r
	})
	assert.Equal(t, r, o.GetUpperBound())
	assert.Equal(t, r, o.GetUpperBound())
	o.AssertExpectations(t)
}

// TestMockETypedElementSetUpperBound tests method SetUpperBound
func TestMockETypedElementSetUpperBound(t *testing.T) {
	o := &MockETypedElement{}
	v := int(45)
	o.On("SetUpperBound", v).Once()
	o.SetUpperBound(v)
	o.AssertExpectations(t)
}
