package ecore

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReflectiveEObjectImpl_EContainer(t *testing.T) {
	mockClass := new(MockEClass)
	mockContainer := new(MockEObjectInternal)
	mockResource := new(MockEResource)
	mockResourceSet := new(MockEResourceSet)

	o := NewReflectiveEObjectImpl()
	o.setEClass(mockClass)
	o.ESetInternalContainer(mockContainer, 0)
	o.ESetInternalResource(mockResource)

	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	o.EAdapters().Add(mockAdapter)
	mock.AssertExpectationsForObjects(t, mockAdapter)

	// non proxy
	mockContainer.On("EIsProxy").Return(false).Once()
	assert.Equal(t, mockContainer, o.EContainer())
	mock.AssertExpectationsForObjects(t, mockClass, mockContainer, mockAdapter)

	// proxy
	mockURI, _ := url.Parse("test://file.t")
	mockResolved := new(MockEObjectInternal)
	mockResolved.On("EProxyURI").Return(nil).Once()
	mockResource.On("GetResourceSet").Return(mockResourceSet).Once()
	mockResourceSet.On("GetEObject", mockURI, true).Return(mockResolved).Once()
	mockContainer.On("EIsProxy").Return(true).Once()
	mockContainer.On("EProxyURI").Return(mockURI).Twice()
	mockClass.On("GetEStructuralFeature", 0).Return(nil).Once()
	mockAdapter.On("NotifyChanged", mock.MatchedBy(func(notification ENotification) bool {
		return notification.GetEventType() == RESOLVE
	})).Once()
	assert.Equal(t, mockResolved, o.EContainer())
	mock.AssertExpectationsForObjects(t, mockClass, mockContainer, mockResolved, mockResource, mockResourceSet, mockAdapter)
}

func TestReflectiveEObjectImpl_ESetResource(t *testing.T) {
	mockClass := new(MockEClass)
	mockContainer := new(MockEObjectInternal)
	mockResource := new(MockEResource)
	mockNotifications := new(MockENotificationChain)
	mockReference := new(MockEReference)
	o := NewReflectiveEObjectImpl()
	o.setEClass(mockClass)
	o.ESetInternalContainer(mockContainer, 0)

	// set resource with container feature which is a reference proxy
	mockClass.On("GetEStructuralFeature", 0).Return(mockReference).Once()
	mockReference.On("IsResolveProxies").Return(true).Once()
	mockContainer.On("EInternalResource").Return(mockResource).Once()
	mockResource.On("Detached", o).Once()
	assert.Equal(t, mockNotifications, o.ESetResource(mockResource, mockNotifications))
	mock.AssertExpectationsForObjects(t, mockResource, mockContainer, mockReference, mockClass, mockNotifications)

	// reset resource with container feature as a reference proxy
	mockClass.On("GetEStructuralFeature", 0).Return(mockReference).Once()
	mockReference.On("IsResolveProxies").Return(true).Once()
	mockContainer.On("EInternalResource").Return(mockResource).Once()
	mockResource.On("Attached", o).Once()
	assert.Equal(t, mockNotifications, o.ESetResource(nil, mockNotifications))
	mock.AssertExpectationsForObjects(t, mockResource, mockContainer, mockReference, mockClass, mockNotifications)

	// set new resource with a container as a non proxy reference
	mockClass.On("GetEStructuralFeature", 0).Return(mockReference).Twice()
	mockReference.On("IsResolveProxies").Return(false).Once()
	mockReference.On("GetEOpposite").Return(nil).Once()
	mockContainer.On("EInternalResource").Return(mockResource).Once()
	mockResource.On("Detached", o).Once()
	assert.Equal(t, mockNotifications, o.ESetResource(mockResource, mockNotifications))
	mock.AssertExpectationsForObjects(t, mockResource, mockContainer, mockReference, mockClass, mockNotifications)
}

func TestReflectiveEObjectImpl_EGetFromID(t *testing.T) {
	mockClass := new(MockEClass)
	o := NewReflectiveEObjectImpl()
	o.setEClass(mockClass)

	mockClass.On("GetEStructuralFeature", 0).Return(nil).Once()
	assert.Panics(t, func() { o.EGetFromID(0, false) })
	mock.AssertExpectationsForObjects(t, mockClass)
}

func TestReflectiveEObjectImpl_ESetFromID(t *testing.T) {
	mockClass := new(MockEClass)
	o := NewReflectiveEObjectImpl()
	o.setEClass(mockClass)

	mockClass.On("GetEStructuralFeature", 0).Return(nil).Once()
	assert.Panics(t, func() { o.ESetFromID(0, false) })
	mock.AssertExpectationsForObjects(t, mockClass)
}

func TestReflectiveEObjectImpl_EUnSetFromID(t *testing.T) {
	mockClass := new(MockEClass)
	o := NewReflectiveEObjectImpl()
	o.setEClass(mockClass)

	mockClass.On("GetEStructuralFeature", 0).Return(nil).Once()
	assert.Panics(t, func() { o.EUnsetFromID(0) })
	mock.AssertExpectationsForObjects(t, mockClass)
}

func TestReflectiveEObjectImpl_GetAttribute(t *testing.T) {
	mockAttribute := new(MockEAttribute)
	mockAttribute.On("IsMany").Return(false).Once()

	mockClass := new(MockEClass)
	mockClass.On("GetFeatureCount").Return(2).Once()
	mockClass.On("GetEStructuralFeature", 0).Return(mockAttribute).Once()

	o := NewReflectiveEObjectImpl()
	o.setEClass(mockClass)

	// get unitialized value
	assert.Nil(t, o.EGetFromID(0, true))

	mock.AssertExpectationsForObjects(t, mockClass, mockAttribute)
}

func TestReflectiveEObjectImpl_GetAttribute_Many(t *testing.T) {
	mockAttribute := new(MockEAttribute)
	mockClass := new(MockEClass)
	o := NewReflectiveEObjectImpl()
	o.setEClass(mockClass)

	// get unitialized value
	mockClass.On("GetFeatureCount").Return(2).Once()
	mockClass.On("GetEStructuralFeature", 0).Return(mockAttribute).Once()
	mockAttribute.On("IsMany").Return(true).Once()
	mockAttribute.On("IsUnique").Return(true).Once()
	val := o.EGetFromID(0, false)
	assert.NotNil(t, val)
	l, _ := val.(EList)
	assert.NotNil(t, l)

	mock.AssertExpectationsForObjects(t, mockClass, mockAttribute)
}

func TestReflectiveEObjectImpl_GetReference_Many(t *testing.T) {
	mockReference := new(MockEReference)
	mockClass := new(MockEClass)
	o := NewReflectiveEObjectImpl()
	o.setEClass(mockClass)

	// get unitialized value
	mockClass.On("GetFeatureCount").Return(2).Once()
	mockClass.On("GetEStructuralFeature", 0).Return(mockReference).Once()
	mockReference.On("IsMany").Return(true).Once()
	mockReference.On("GetEOpposite").Return(nil).Twice()
	mockReference.On("IsContainment").Return(false).Once()
	mockReference.On("GetFeatureID").Return(0).Once()
	mockReference.On("EIsProxy").Return(false).Once()
	mockReference.On("IsUnsettable").Return(false).Once()
	val := o.EGetFromID(0, false)
	assert.NotNil(t, val)

	// check its is an object list
	l, _ := val.(EObjectList)
	assert.NotNil(t, l)

	mock.AssertExpectationsForObjects(t, mockClass, mockReference)
}

func TestReflectiveEObjectImpl_SetAttribute(t *testing.T) {
	mockAttribute := new(MockEAttribute)

	mockClass := new(MockEClass)
	mockClass.On("GetFeatureCount").Return(2).Once()
	mockClass.On("GetEStructuralFeature", 0).Return(mockAttribute).Twice()

	mockObject := new(MockEObject)

	o := NewReflectiveEObjectImpl()
	o.setEClass(mockClass)

	// set
	o.ESetFromID(0, mockObject)

	// check that value is well set
	assert.Equal(t, mockObject, o.EGetFromID(0, true))

	mock.AssertExpectationsForObjects(t, mockClass, mockAttribute)
}

func TestReflectiveEObjectImpl_UnsetAttribute(t *testing.T) {
	mockAttribute := new(MockEAttribute)
	mockAttribute.On("IsMany").Return(false).Once()

	mockClass := new(MockEClass)
	mockClass.On("GetFeatureCount").Return(2).Once()
	mockClass.On("GetEStructuralFeature", 0).Return(mockAttribute).Times(3)

	mockObject := new(MockEObject)

	o := NewReflectiveEObjectImpl()
	o.setEClass(mockClass)

	// set - unset
	o.ESetFromID(0, mockObject)
	o.EUnsetFromID(0)

	// check that value is unset
	assert.Nil(t, o.EGetFromID(0, true))

	mock.AssertExpectationsForObjects(t, mockClass, mockAttribute)
}

func TestReflectiveEObjectImpl_GetContainer(t *testing.T) {
	mockOpposite := new(MockEReference)
	mockReference := new(MockEReference)
	mockClass := new(MockEClass)

	o := NewReflectiveEObjectImpl()
	o.setEClass(mockClass)

	// get non initialized container
	mockOpposite.On("IsContainment").Return(true).Once()
	mockReference.On("GetEOpposite").Return(mockOpposite).Once()
	mockClass.On("GetEStructuralFeature", 0).Return(mockReference).Once()
	mockClass.On("GetFeatureID", mockReference).Return(0).Once()
	assert.Nil(t, o.EGetFromID(0, true))
	mock.AssertExpectationsForObjects(t, mockClass, mockReference, mockOpposite)
}

func TestReflectiveEObjectImpl_SetContainer(t *testing.T) {
	mockObject := new(MockEObjectInternal)
	mockObjectClass := new(MockEClass)
	mockOpposite := new(MockEReference)
	mockReference := new(MockEReference)
	mockClass := new(MockEClass)

	o := NewReflectiveEObjectImpl()
	o.setEClass(mockClass)

	// set reference as mockObject
	mockObject.On("EInternalResource").Return(nil).Once()
	mockObject.On("EInverseAdd", o, 0, nil).Return(nil).Once()
	mockObject.On("EClass").Return(mockObjectClass).Once()
	mockOpposite.On("IsContainment").Return(true).Once()
	mockReference.On("GetEOpposite").Return(mockOpposite).Twice()
	mockClass.On("GetEStructuralFeature", 0).Return(mockReference).Once()
	mockClass.On("GetFeatureID", mockReference).Return(0).Once()
	mockObjectClass.On("GetFeatureID", mockOpposite).Return(0).Once()
	o.ESetFromID(0, mockObject)
	mock.AssertExpectationsForObjects(t, mockClass, mockReference, mockOpposite, mockObject, mockObjectClass)

	// get unresolved
	mockOpposite.On("IsContainment").Return(true).Once()
	mockReference.On("GetEOpposite").Return(mockOpposite).Once()
	mockClass.On("GetEStructuralFeature", 0).Return(mockReference).Once()
	mockClass.On("GetFeatureID", mockReference).Return(0).Once()
	assert.Equal(t, mockObject, o.EGetFromID(0, false))
	mock.AssertExpectationsForObjects(t, mockClass, mockReference, mockOpposite, mockObject)

	// get resolved
	mockObject.On("EIsProxy").Return(false).Once()
	mockOpposite.On("IsContainment").Return(true).Once()
	mockReference.On("GetEOpposite").Return(mockOpposite).Once()
	mockClass.On("GetEStructuralFeature", 0).Return(mockReference).Once()
	mockClass.On("GetFeatureID", mockReference).Return(0).Once()
	assert.Equal(t, mockObject, o.EGetFromID(0, true))
	mock.AssertExpectationsForObjects(t, mockClass, mockReference, mockOpposite, mockObject)

	// set reference as nil
	mockObject.On("EInverseRemove", o.GetInterfaces(), 0, nil).Return(nil).Once()
	mockObject.On("EInternalResource").Return(nil).Once()
	mockOpposite.On("IsContainment").Return(true).Once()
	mockOpposite.On("GetFeatureID").Return(0).Once()
	mockReference.On("GetEOpposite").Return(mockOpposite).Twice()
	mockClass.On("GetEStructuralFeature", 0).Return(mockReference).Twice()
	mockClass.On("GetFeatureID", mockReference).Return(0).Once()
	o.ESetFromID(0, nil)
	mock.AssertExpectationsForObjects(t, mockClass, mockReference, mockOpposite, mockObject)

	// get unresolved
	mockOpposite.On("IsContainment").Return(true).Once()
	mockReference.On("GetEOpposite").Return(mockOpposite).Once()
	mockClass.On("GetEStructuralFeature", 0).Return(mockReference).Once()
	mockClass.On("GetFeatureID", mockReference).Return(0).Once()
	assert.Nil(t, o.EGetFromID(0, false))
	mock.AssertExpectationsForObjects(t, mockClass, mockReference, mockOpposite, mockObject)
}

func TestReflectiveEObjectImpl_UnSetContainer(t *testing.T) {

	mockObject := new(MockEObjectInternal)
	mockObjectClass := new(MockEClass)
	mockOpposite := new(MockEReference)
	mockReference := new(MockEReference)
	mockClass := new(MockEClass)

	o := NewReflectiveEObjectImpl()
	o.setEClass(mockClass)

	// set reference as mockObject
	mockObject.On("EInternalResource").Return(nil).Once()
	mockObject.On("EInverseAdd", o, 0, nil).Return(nil).Once()
	mockObject.On("EClass").Return(mockObjectClass).Once()
	mockOpposite.On("IsContainment").Return(true).Once()
	mockReference.On("GetEOpposite").Return(mockOpposite).Twice()
	mockClass.On("GetEStructuralFeature", 0).Return(mockReference).Once()
	mockClass.On("GetFeatureID", mockReference).Return(0).Once()
	mockObjectClass.On("GetFeatureID", mockOpposite).Return(0).Once()
	o.ESetFromID(0, mockObject)
	mock.AssertExpectationsForObjects(t, mockClass, mockReference, mockOpposite, mockObject, mockObjectClass)

	// unset
	mockObject.On("EInverseRemove", o.GetInterfaces(), 0, nil).Return(nil).Once()
	mockObject.On("EInternalResource").Return(nil).Once()
	mockOpposite.On("IsContainment").Return(true).Once()
	mockOpposite.On("GetFeatureID").Return(0).Once()
	mockReference.On("GetEOpposite").Return(mockOpposite).Twice()
	mockClass.On("GetEStructuralFeature", 0).Return(mockReference).Twice()
	mockClass.On("GetFeatureID", mockReference).Return(0).Once()
	o.EUnsetFromID(0)

	// get unresolved
	mockOpposite.On("IsContainment").Return(true).Once()
	mockReference.On("GetEOpposite").Return(mockOpposite).Once()
	mockClass.On("GetEStructuralFeature", 0).Return(mockReference).Once()
	mockClass.On("GetFeatureID", mockReference).Return(0).Once()
	assert.Nil(t, o.EGetFromID(0, false))
	mock.AssertExpectationsForObjects(t, mockClass, mockReference, mockOpposite, mockObject)
}
