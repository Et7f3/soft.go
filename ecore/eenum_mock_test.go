// Code generated by soft.generator.go. DO NOT EDIT.

// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func discardMockEEnum() {
	_ = assert.Equal
	_ = testing.Coverage
}

// TestMockEEnumGetELiterals tests method GetELiterals
func TestMockEEnumGetELiterals(t *testing.T) {
	o := &MockEEnum{}
	l := &MockEList{}
	// return a value
	o.On("GetELiterals").Once().Return(l)
	o.On("GetELiterals").Once().Return(func() EList {
		return l
	})
	assert.Equal(t, l, o.GetELiterals())
	assert.Equal(t, l, o.GetELiterals())
	o.AssertExpectations(t)
}

// TestMockEEnumGetEEnumLiteralByLiteral tests method GetEEnumLiteralByLiteral
func TestMockEEnumGetEEnumLiteralByLiteral(t *testing.T) {
	o := &MockEEnum{}
	literal := string("Test String")
	r := new(MockEEnumLiteral)
	o.On("GetEEnumLiteralByLiteral", literal).Return(r).Once()
	o.On("GetEEnumLiteralByLiteral", literal).Return(func() EEnumLiteral {
		return r
	}).Once()
	assert.Equal(t, r, o.GetEEnumLiteralByLiteral(literal))
	assert.Equal(t, r, o.GetEEnumLiteralByLiteral(literal))
	o.AssertExpectations(t)
}

// TestMockEEnumGetEEnumLiteralByName tests method GetEEnumLiteralByName
func TestMockEEnumGetEEnumLiteralByName(t *testing.T) {
	o := &MockEEnum{}
	name := string("Test String")
	r := new(MockEEnumLiteral)
	o.On("GetEEnumLiteralByName", name).Return(r).Once()
	o.On("GetEEnumLiteralByName", name).Return(func() EEnumLiteral {
		return r
	}).Once()
	assert.Equal(t, r, o.GetEEnumLiteralByName(name))
	assert.Equal(t, r, o.GetEEnumLiteralByName(name))
	o.AssertExpectations(t)
}

// TestMockEEnumGetEEnumLiteralByValue tests method GetEEnumLiteralByValue
func TestMockEEnumGetEEnumLiteralByValue(t *testing.T) {
	o := &MockEEnum{}
	value := int(45)
	r := new(MockEEnumLiteral)
	o.On("GetEEnumLiteralByValue", value).Return(r).Once()
	o.On("GetEEnumLiteralByValue", value).Return(func() EEnumLiteral {
		return r
	}).Once()
	assert.Equal(t, r, o.GetEEnumLiteralByValue(value))
	assert.Equal(t, r, o.GetEEnumLiteralByValue(value))
	o.AssertExpectations(t)
}
