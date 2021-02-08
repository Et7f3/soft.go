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

func discardMockEObject() {
	_ = assert.Equal
	_ = testing.Coverage
}

// TestMockEObjectEAllContents tests method EAllContents
func TestMockEObjectEAllContents(t *testing.T) {
	o := &MockEObject{}
	r := EIterator(nil)
	o.On("EAllContents").Return(r).Once()
	o.On("EAllContents").Return(func() EIterator {
		return r
	}).Once()
	assert.Equal(t, r, o.EAllContents())
	assert.Equal(t, r, o.EAllContents())
	o.AssertExpectations(t)
}

// TestMockEObjectEClass tests method EClass
func TestMockEObjectEClass(t *testing.T) {
	o := &MockEObject{}
	r := new(MockEClass)
	o.On("EClass").Return(r).Once()
	o.On("EClass").Return(func() EClass {
		return r
	}).Once()
	assert.Equal(t, r, o.EClass())
	assert.Equal(t, r, o.EClass())
	o.AssertExpectations(t)
}

// TestMockEObjectEContainer tests method EContainer
func TestMockEObjectEContainer(t *testing.T) {
	o := &MockEObject{}
	r := new(MockEObjectInternal)
	o.On("EContainer").Return(r).Once()
	o.On("EContainer").Return(func() EObject {
		return r
	}).Once()
	assert.Equal(t, r, o.EContainer())
	assert.Equal(t, r, o.EContainer())
	o.AssertExpectations(t)
}

// TestMockEObjectEContainingFeature tests method EContainingFeature
func TestMockEObjectEContainingFeature(t *testing.T) {
	o := &MockEObject{}
	r := new(MockEStructuralFeature)
	o.On("EContainingFeature").Return(r).Once()
	o.On("EContainingFeature").Return(func() EStructuralFeature {
		return r
	}).Once()
	assert.Equal(t, r, o.EContainingFeature())
	assert.Equal(t, r, o.EContainingFeature())
	o.AssertExpectations(t)
}

// TestMockEObjectEContainmentFeature tests method EContainmentFeature
func TestMockEObjectEContainmentFeature(t *testing.T) {
	o := &MockEObject{}
	r := new(MockEReference)
	o.On("EContainmentFeature").Return(r).Once()
	o.On("EContainmentFeature").Return(func() EReference {
		return r
	}).Once()
	assert.Equal(t, r, o.EContainmentFeature())
	assert.Equal(t, r, o.EContainmentFeature())
	o.AssertExpectations(t)
}

// TestMockEObjectEContents tests method EContents
func TestMockEObjectEContents(t *testing.T) {
	o := &MockEObject{}
	r := EList(NewEmptyBasicEList())
	o.On("EContents").Return(r).Once()
	o.On("EContents").Return(func() EList {
		return r
	}).Once()
	assert.Equal(t, r, o.EContents())
	assert.Equal(t, r, o.EContents())
	o.AssertExpectations(t)
}

// TestMockEObjectECrossReferences tests method ECrossReferences
func TestMockEObjectECrossReferences(t *testing.T) {
	o := &MockEObject{}
	r := EList(NewEmptyBasicEList())
	o.On("ECrossReferences").Return(r).Once()
	o.On("ECrossReferences").Return(func() EList {
		return r
	}).Once()
	assert.Equal(t, r, o.ECrossReferences())
	assert.Equal(t, r, o.ECrossReferences())
	o.AssertExpectations(t)
}

// TestMockEObjectEGet tests method EGet
func TestMockEObjectEGet(t *testing.T) {
	o := &MockEObject{}
	feature := new(MockEStructuralFeature)
	r := interface{}(nil)
	o.On("EGet", feature).Return(r).Once()
	o.On("EGet", feature).Return(func() interface{} {
		return r
	}).Once()
	assert.Equal(t, r, o.EGet(feature))
	assert.Equal(t, r, o.EGet(feature))
	o.AssertExpectations(t)
}

// TestMockEObjectEGetResolve tests method EGetResolve
func TestMockEObjectEGetResolve(t *testing.T) {
	o := &MockEObject{}
	feature := new(MockEStructuralFeature)
	resolve := bool(true)
	r := interface{}(nil)
	o.On("EGetResolve", feature, resolve).Return(r).Once()
	o.On("EGetResolve", feature, resolve).Return(func() interface{} {
		return r
	}).Once()
	assert.Equal(t, r, o.EGetResolve(feature, resolve))
	assert.Equal(t, r, o.EGetResolve(feature, resolve))
	o.AssertExpectations(t)
}

// TestMockEObjectEInvoke tests method EInvoke
func TestMockEObjectEInvoke(t *testing.T) {
	o := &MockEObject{}
	operation := new(MockEOperation)
	arguments := EList(NewEmptyBasicEList())
	r := interface{}(nil)
	o.On("EInvoke", operation, arguments).Return(r).Once()
	o.On("EInvoke", operation, arguments).Return(func() interface{} {
		return r
	}).Once()
	assert.Equal(t, r, o.EInvoke(operation, arguments))
	assert.Equal(t, r, o.EInvoke(operation, arguments))
	o.AssertExpectations(t)
}

// TestMockEObjectEIsProxy tests method EIsProxy
func TestMockEObjectEIsProxy(t *testing.T) {
	o := &MockEObject{}
	r := bool(true)
	o.On("EIsProxy").Return(r).Once()
	o.On("EIsProxy").Return(func() bool {
		return r
	}).Once()
	assert.Equal(t, r, o.EIsProxy())
	assert.Equal(t, r, o.EIsProxy())
	o.AssertExpectations(t)
}

// TestMockEObjectEIsSet tests method EIsSet
func TestMockEObjectEIsSet(t *testing.T) {
	o := &MockEObject{}
	feature := new(MockEStructuralFeature)
	r := bool(true)
	o.On("EIsSet", feature).Return(r).Once()
	o.On("EIsSet", feature).Return(func() bool {
		return r
	}).Once()
	assert.Equal(t, r, o.EIsSet(feature))
	assert.Equal(t, r, o.EIsSet(feature))
	o.AssertExpectations(t)
}

// TestMockEObjectEResource tests method EResource
func TestMockEObjectEResource(t *testing.T) {
	o := &MockEObject{}
	r := EResource(nil)
	o.On("EResource").Return(r).Once()
	o.On("EResource").Return(func() EResource {
		return r
	}).Once()
	assert.Equal(t, r, o.EResource())
	assert.Equal(t, r, o.EResource())
	o.AssertExpectations(t)
}

// TestMockEObjectESet tests method ESet
func TestMockEObjectESet(t *testing.T) {
	o := &MockEObject{}
	feature := new(MockEStructuralFeature)
	newValue := interface{}(nil)
	o.On("ESet", feature, newValue).Once()
	o.ESet(feature, newValue)
	o.AssertExpectations(t)
}

// TestMockEObjectEUnset tests method EUnset
func TestMockEObjectEUnset(t *testing.T) {
	o := &MockEObject{}
	feature := new(MockEStructuralFeature)
	o.On("EUnset", feature).Once()
	o.EUnset(feature)
	o.AssertExpectations(t)
}
