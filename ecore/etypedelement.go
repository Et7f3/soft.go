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

// ETypedElement is the representation of the model object 'ETypedElement'
type ETypedElement interface {
	ENamedElement

	IsOrdered() bool
	SetOrdered(bool)

	IsUnique() bool
	SetUnique(bool)

	GetLowerBound() int
	SetLowerBound(int)

	GetUpperBound() int
	SetUpperBound(int)

	IsMany() bool

	IsRequired() bool

	GetEType() EClassifier
	SetEType(EClassifier)
	UnsetEType()

	// Start of user code ETypedElement
	// End of user code
}
