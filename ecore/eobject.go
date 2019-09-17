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

// EObject is the representation of the model object 'EObject'
type EObject interface {
	ENotifier

	EClass() EClass
	EIsProxy() bool
	EResource() EResource
	EContainer() EObject
	EContainingFeature() EStructuralFeature
	EContainmentFeature() EReference
	EContents() EList
	EAllContents() EIterator
	ECrossReferences() EList
	EGet(EStructuralFeature) interface{}
	EGetResolve(EStructuralFeature, bool) interface{}
	ESet(EStructuralFeature, interface{})
	EIsSet(EStructuralFeature) bool
	EUnset(EStructuralFeature)
	EInvoke(EOperation, EList) interface{}

	// Start of user code EObject
	// End of user code
}
