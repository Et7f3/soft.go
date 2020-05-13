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

// eTypeParameterImpl is the implementation of the model object 'ETypeParameter'
type eTypeParameterImpl struct {
	*eNamedElementImpl
	eBounds EList
}

// newETypeParameterImpl is the constructor of a eTypeParameterImpl
func newETypeParameterImpl() *eTypeParameterImpl {
	eTypeParameter := new(eTypeParameterImpl)
	eTypeParameter.eNamedElementImpl = newENamedElementImpl()
	eTypeParameter.SetInterfaces(eTypeParameter)

	return eTypeParameter
}

type eTypeParameterImplInitializers interface {
	initEBounds() EList
}

func (eTypeParameter *eTypeParameterImpl) getInitializers() eTypeParameterImplInitializers {
	return eTypeParameter.AsEObject().(eTypeParameterImplInitializers)
}

func (eTypeParameter *eTypeParameterImpl) EStaticClass() EClass {
	return GetPackage().GetETypeParameter()
}

// GetEBounds get the value of eBounds
func (eTypeParameter *eTypeParameterImpl) GetEBounds() EList {
	if eTypeParameter.eBounds == nil {
		eTypeParameter.eBounds = eTypeParameter.getInitializers().initEBounds()
	}
	return eTypeParameter.eBounds
}

func (eTypeParameter *eTypeParameterImpl) initEBounds() EList {
	return NewBasicEObjectList(eTypeParameter.AsEObjectInternal(), ETYPE_PARAMETER__EBOUNDS, -1, true, true, false, false, false)
}

func (eTypeParameter *eTypeParameterImpl) EGetFromID(featureID int, resolve bool) interface{} {
	switch featureID {
	case ETYPE_PARAMETER__EBOUNDS:
		return eTypeParameter.GetEBounds()
	default:
		return eTypeParameter.eNamedElementImpl.EGetFromID(featureID, resolve)
	}
}

func (eTypeParameter *eTypeParameterImpl) ESetFromID(featureID int, newValue interface{}) {
	switch featureID {
	case ETYPE_PARAMETER__EBOUNDS:
		e := newValue.(EList)
		eTypeParameter.GetEBounds().Clear()
		eTypeParameter.GetEBounds().AddAll(e)
	default:
		eTypeParameter.eNamedElementImpl.ESetFromID(featureID, newValue)
	}
}

func (eTypeParameter *eTypeParameterImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case ETYPE_PARAMETER__EBOUNDS:
		eTypeParameter.GetEBounds().Clear()
	default:
		eTypeParameter.eNamedElementImpl.EUnsetFromID(featureID)
	}
}

func (eTypeParameter *eTypeParameterImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case ETYPE_PARAMETER__EBOUNDS:
		return eTypeParameter.eBounds != nil && eTypeParameter.eBounds.Size() != 0
	default:
		return eTypeParameter.eNamedElementImpl.EIsSetFromID(featureID)
	}
}

func (eTypeParameter *eTypeParameterImpl) EBasicInverseRemove(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case ETYPE_PARAMETER__EBOUNDS:
		list := eTypeParameter.GetEBounds().(ENotifyingList)
		return list.RemoveWithNotification(otherEnd, notifications)
	default:
		return eTypeParameter.eNamedElementImpl.EBasicInverseRemove(otherEnd, featureID, notifications)
	}
}
