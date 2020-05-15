package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBasicEObjectListAccessors(t *testing.T) {
	{
		list := NewBasicEObjectList(nil, 1, -1, false, true, false, false, false)
		assert.Equal(t, nil, list.GetNotifier())
		assert.Equal(t, nil, list.GetFeature())
		assert.Equal(t, 1, list.GetFeatureID())
	}
	{
		mockOwner := &MockEObjectInternal{}
		list := NewBasicEObjectList(mockOwner, 1, -1, false, true, false, false, false)
		assert.Equal(t, mockOwner, list.GetNotifier())
		assert.Equal(t, 1, list.GetFeatureID())
		mockClass := &MockEClass{}
		mockFeature := &MockEStructuralFeature{}
		mockClass.On("GetEStructuralFeature", 1).Return(mockFeature)
		mockOwner.On("EClass").Return(mockClass)
		assert.Equal(t, mockFeature, list.GetFeature())
		mock.AssertExpectationsForObjects(t, mockOwner, mockClass, mockFeature, mockClass)
	}
}

func TestBasicEObjectListInverseNoOpposite(t *testing.T) {
	mockOwner := &MockEObjectInternal{}
	mockOwner.On("EDeliver").Return(false)

	mockObject := &MockEObjectInternal{}
	list := NewBasicEObjectList(mockOwner, 1, -1, false, true, false, false, false)
	mockObject.On("EInverseAdd", mockOwner, -2, nil).Return(nil)

	assert.True(t, list.Add(mockObject))

	mockObject.On("EInverseRemove", mockOwner, -2, nil).Return(nil)
	assert.True(t, list.Remove(mockObject))

	mock.AssertExpectationsForObjects(t, mockObject, mockOwner)
}

func TestBasicEObjectListInverseOpposite(t *testing.T) {
	mockOwner := &MockEObjectInternal{}
	mockOwner.On("EDeliver").Return(false)

	mockObject := &MockEObjectInternal{}
	list := NewBasicEObjectList(mockOwner, 1, 2, false, true, true, false, false)

	mockObject.On("EInverseAdd", mockOwner, 2, nil).Return(nil)
	assert.True(t, list.Add(mockObject))

	mockObject.On("EInverseRemove", mockOwner, 2, nil).Return(nil)
	assert.True(t, list.Remove(mockObject))

	mock.AssertExpectationsForObjects(t, mockObject, mockOwner)
}

func TestBasicEObjectListContains(t *testing.T) {
	mockOwner := &MockEObjectInternal{}
	mockOwner.On("EDeliver").Return(false)

	// no proxy
	{
		list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, false, false)
		mockObject := &MockEObjectInternal{}
		list.Add(mockObject)
		assert.True(t, list.Contains(mockObject))
		assert.True(t, list.Contains(mockObject))

		mock.AssertExpectationsForObjects(t, mockObject, mockOwner)
	}

	// with proxy
	{
		list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
		mockObject := &MockEObjectInternal{}
		list.Add(mockObject)
		assert.True(t, list.Contains(mockObject))

		mockResolved := &MockEObjectInternal{}
		mockOwner.On("EResolveProxy", mockObject).Return(mockResolved)
		mockObject.On("EIsProxy").Return(true)
		assert.True(t, list.Contains(mockResolved))

		mock.AssertExpectationsForObjects(t, mockObject, mockOwner)
	}
}

func TestBasicEObjectListGet(t *testing.T) {
	mockOwner := &MockEObjectInternal{}
	mockOwner.On("EDeliver").Return(false)

	// no proxy
	{
		list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, false, false)
		mockObject := &MockEObjectInternal{}
		list.Add(mockObject)
		assert.Equal(t, mockObject, list.Get(0))

		mock.AssertExpectationsForObjects(t, mockObject, mockOwner)
	}

	// with proxy
	{
		list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
		mockObject := &MockEObjectInternal{}
		list.Add(mockObject)

		mockResolved := &MockEObjectInternal{}
		mockOwner.On("EResolveProxy", mockObject).Return(mockResolved)
		mockObject.On("EIsProxy").Return(true)
		assert.Equal(t, mockResolved, list.Get(0))

		mock.AssertExpectationsForObjects(t, mockObject, mockOwner)
	}
}

func TestBasicEObjectListUnResolved(t *testing.T) {
	mockOwner := &MockEObjectInternal{}
	// no proxy
	{
		list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, false, false)
		assert.Equal(t, list, list.GetUnResolvedList())
	}
	// with proxy
	{
		list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
		unresolved := list.GetUnResolvedList()
		assert.NotEqual(t, list, unresolved)
		unresolvedAsEObjectList, _ := unresolved.(EObjectList)
		assert.NotNil(t, unresolvedAsEObjectList)
		assert.Equal(t, unresolved, unresolvedAsEObjectList.GetUnResolvedList())
		unresolvedAsENotifyingList, _ := unresolved.(ENotifyingList)
		assert.NotNil(t, unresolvedAsENotifyingList)
	}
}

func TestBasicEObjectListUnResolvedGet(t *testing.T) {
	mockOwner := &MockEObjectInternal{}
	mockOwner.On("EDeliver").Return(false)

	// add an object to unresolved
	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()
	mockObject := &MockEObjectInternal{}
	unresolved.Add(mockObject)
	// check that in unresolved it is the same
	assert.Equal(t, mockObject, unresolved.Get(0))

	// check that in original list , there is a resolution
	mockResolved := &MockEObjectInternal{}
	mockOwner.On("EResolveProxy", mockObject).Once().Return(mockResolved)
	mockObject.On("EIsProxy").Return(true)
	assert.Equal(t, mockResolved, list.Get(0))

	// check that now it is the resolved one in the unresolved list
	assert.Equal(t, mockResolved, unresolved.Get(0))
	assert.Panics(t, func() { unresolved.Get(1) })

	mock.AssertExpectationsForObjects(t, mockOwner, mockObject, mockResolved)
}

func TestBasicEObjectListUnResolvedContains(t *testing.T) {
	mockOwner := &MockEObjectInternal{}
	mockOwner.On("EDeliver").Return(false)

	// add an object to unresolved
	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()
	mockObject := &MockEObjectInternal{}
	unresolved.Add(mockObject)

	assert.True(t, unresolved.Contains(mockObject))

	// check that in original list , there is a resolution
	mockResolved := &MockEObjectInternal{}
	mockOwner.On("EResolveProxy", mockObject).Once().Return(mockResolved)
	mockObject.On("EIsProxy").Return(true)
	assert.True(t, !unresolved.Contains(mockResolved))
	assert.True(t, list.Contains(mockResolved))

	mock.AssertExpectationsForObjects(t, mockOwner, mockObject, mockResolved)
}

func TestBasicEObjectListUnResolvedSet(t *testing.T) {
	mockOwner := &MockEObjectInternal{}
	mockOwner.On("EDeliver").Return(false)

	// add an object to unresolved
	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()
	mockObject := &MockEObjectInternal{}
	unresolved.Add(mockObject)

	// set first index as another object & check that it has been replaced
	mockObject1 := &MockEObjectInternal{}
	unresolved.Set(0, mockObject1)
	assert.Equal(t, mockObject1, unresolved.Get(0))

	// check that invalid range is supported
	assert.Panics(t, func() { unresolved.Set(1, mockObject) })

	mock.AssertExpectationsForObjects(t, mockOwner, mockObject)
}