// *****************************************************************************
//
// This file is part of a MASA library or program.
// Refer to the included end-user license agreement for restrictions.
//
// Copyright (c) 2019 MASA Group
//
// *****************************************************************************

// *****************************************************************************
//
// Warning: This file was generated by soft.generator.go Generator
//
// *****************************************************************************

package ecore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func discardEEnumLiteral() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage

	_ = time.Now()
}

func TestEEnumLiteralValueGet(t *testing.T) {
	obj := newEEnumLiteralImpl()
	obj.SetValue(45)
	assert.Equal(t, 45, obj.GetValue())
}

func TestEEnumLiteralValueSet(t *testing.T) {
	mockAdapter := &MockEAdapter{}
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	obj := newEEnumLiteralImpl()
	obj.EAdapters().Add(mockAdapter)
	obj.SetValue(45)
	mockAdapter.AssertExpectations(t)
}

func TestEEnumLiteralValueEGet(t *testing.T) {
	obj := newEEnumLiteralImpl()
	{
		assert.Equal(t, obj.GetEEnum(), obj.EGetFromID(EENUM_LITERAL__EENUM, false, false))
	}
	{
		assert.Equal(t, obj.GetInstance(), obj.EGetFromID(EENUM_LITERAL__INSTANCE, false, false))
	}
	{
		assert.Equal(t, obj.GetLiteral(), obj.EGetFromID(EENUM_LITERAL__LITERAL, false, false))
	}
	{
		assert.Equal(t, obj.GetValue(), obj.EGetFromID(EENUM_LITERAL__VALUE, false, false))
	}
}

func TestEEnumLiteralValueEIsSet(t *testing.T) {
	obj := newEEnumLiteralImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEEnumLiteralValueEUnset(t *testing.T) {
	obj := newEEnumLiteralImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEEnumLiteralValueESet(t *testing.T) {
	obj := newEEnumLiteralImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEEnumLiteralValueEInverseAdd(t *testing.T) {
	{
	}
}

func TestEEnumLiteralValueEInverseRemove(t *testing.T) {
	{
	}
}

func TestEEnumLiteralInstanceGet(t *testing.T) {
	obj := newEEnumLiteralImpl()
	obj.SetInstance(nil)
	assert.Equal(t, nil, obj.GetInstance())
}

func TestEEnumLiteralInstanceSet(t *testing.T) {
	mockAdapter := &MockEAdapter{}
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	obj := newEEnumLiteralImpl()
	obj.EAdapters().Add(mockAdapter)
	obj.SetInstance(nil)
	mockAdapter.AssertExpectations(t)
}

func TestEEnumLiteralInstanceEGet(t *testing.T) {
	obj := newEEnumLiteralImpl()
	{
		assert.Equal(t, obj.GetEEnum(), obj.EGetFromID(EENUM_LITERAL__EENUM, false, false))
	}
	{
		assert.Equal(t, obj.GetInstance(), obj.EGetFromID(EENUM_LITERAL__INSTANCE, false, false))
	}
	{
		assert.Equal(t, obj.GetLiteral(), obj.EGetFromID(EENUM_LITERAL__LITERAL, false, false))
	}
	{
		assert.Equal(t, obj.GetValue(), obj.EGetFromID(EENUM_LITERAL__VALUE, false, false))
	}
}

func TestEEnumLiteralInstanceEIsSet(t *testing.T) {
	obj := newEEnumLiteralImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEEnumLiteralInstanceEUnset(t *testing.T) {
	obj := newEEnumLiteralImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEEnumLiteralInstanceESet(t *testing.T) {
	obj := newEEnumLiteralImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEEnumLiteralInstanceEInverseAdd(t *testing.T) {
	{
	}
}

func TestEEnumLiteralInstanceEInverseRemove(t *testing.T) {
	{
	}
}

func TestEEnumLiteralLiteralGet(t *testing.T) {
	obj := newEEnumLiteralImpl()
	obj.SetLiteral("Test String")
	assert.Equal(t, "Test String", obj.GetLiteral())
}

func TestEEnumLiteralLiteralSet(t *testing.T) {
	mockAdapter := &MockEAdapter{}
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	obj := newEEnumLiteralImpl()
	obj.EAdapters().Add(mockAdapter)
	obj.SetLiteral("Test String")
	mockAdapter.AssertExpectations(t)
}

func TestEEnumLiteralLiteralEGet(t *testing.T) {
	obj := newEEnumLiteralImpl()
	{
		assert.Equal(t, obj.GetEEnum(), obj.EGetFromID(EENUM_LITERAL__EENUM, false, false))
	}
	{
		assert.Equal(t, obj.GetInstance(), obj.EGetFromID(EENUM_LITERAL__INSTANCE, false, false))
	}
	{
		assert.Equal(t, obj.GetLiteral(), obj.EGetFromID(EENUM_LITERAL__LITERAL, false, false))
	}
	{
		assert.Equal(t, obj.GetValue(), obj.EGetFromID(EENUM_LITERAL__VALUE, false, false))
	}
}

func TestEEnumLiteralLiteralEIsSet(t *testing.T) {
	obj := newEEnumLiteralImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEEnumLiteralLiteralEUnset(t *testing.T) {
	obj := newEEnumLiteralImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEEnumLiteralLiteralESet(t *testing.T) {
	obj := newEEnumLiteralImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEEnumLiteralLiteralEInverseAdd(t *testing.T) {
	{
	}
}

func TestEEnumLiteralLiteralEInverseRemove(t *testing.T) {
	{
	}
}
