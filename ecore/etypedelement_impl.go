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

// eTypedElementImpl is the implementation of the model object 'ETypedElement'
type eTypedElementImpl struct {
	*eNamedElementImpl
	eType      EClassifier
	isMany     bool
	isOrdered  bool
	isRequired bool
	isUnique   bool
	lowerBound int
	upperBound int
}

// newETypedElementImpl is the constructor of a eTypedElementImpl
func newETypedElementImpl() *eTypedElementImpl {
	eTypedElement := new(eTypedElementImpl)
	eTypedElement.eNamedElementImpl = newENamedElementImpl()
	eTypedElement.SetInterfaces(eTypedElement)
	eTypedElement.isOrdered = true
	eTypedElement.isUnique = true
	eTypedElement.lowerBound = 0
	eTypedElement.upperBound = 1

	return eTypedElement
}

func (eTypedElement *eTypedElementImpl) EStaticClass() EClass {
	return GetPackage().GetETypedElement()
}

// GetEType get the value of eType
func (eTypedElement *eTypedElementImpl) GetEType() EClassifier {
	if eTypedElement.eType != nil && eTypedElement.eType.EIsProxy() {
		oldEType := eTypedElement.eType
		newEType := eTypedElement.EResolveProxy(oldEType).(EClassifier)
		eTypedElement.eType = newEType
		if newEType != oldEType {
			if eTypedElement.ENotificationRequired() {
				eTypedElement.ENotify(NewNotificationByFeatureID(eTypedElement, RESOLVE, ETYPED_ELEMENT__ETYPE, oldEType, newEType, NO_INDEX))
			}
		}
	}
	return eTypedElement.eType
}

func (eTypedElement *eTypedElementImpl) basicGetEType() EClassifier {
	return eTypedElement.eType
}

// SetEType set the value of eType
func (eTypedElement *eTypedElementImpl) SetEType(newEType EClassifier) {
	oldEType := eTypedElement.eType
	eTypedElement.eType = newEType
	if eTypedElement.ENotificationRequired() {
		eTypedElement.ENotify(NewNotificationByFeatureID(eTypedElement.GetEObject(), SET, ETYPED_ELEMENT__ETYPE, oldEType, newEType, NO_INDEX))
	}
}

// UnsetEType unset the value of eType
func (eTypedElement *eTypedElementImpl) UnsetEType() {
	if eTypedElement.ENotificationRequired() {
		eTypedElement.ENotify(NewNotificationByFeatureID(eTypedElement.GetEObject(), UNSET, ETYPED_ELEMENT__ETYPE, nil, nil, NO_INDEX))
	}
}

// IsMany get the value of isMany
func (eTypedElement *eTypedElementImpl) IsMany() bool {
	panic("IsMany not implemented")
}

// IsOrdered get the value of isOrdered
func (eTypedElement *eTypedElementImpl) IsOrdered() bool {
	return eTypedElement.isOrdered
}

// SetOrdered set the value of isOrdered
func (eTypedElement *eTypedElementImpl) SetOrdered(newIsOrdered bool) {
	oldIsOrdered := eTypedElement.isOrdered
	eTypedElement.isOrdered = newIsOrdered
	if eTypedElement.ENotificationRequired() {
		eTypedElement.ENotify(NewNotificationByFeatureID(eTypedElement.GetEObject(), SET, ETYPED_ELEMENT__ORDERED, oldIsOrdered, newIsOrdered, NO_INDEX))
	}
}

// IsRequired get the value of isRequired
func (eTypedElement *eTypedElementImpl) IsRequired() bool {
	panic("IsRequired not implemented")
}

// IsUnique get the value of isUnique
func (eTypedElement *eTypedElementImpl) IsUnique() bool {
	return eTypedElement.isUnique
}

// SetUnique set the value of isUnique
func (eTypedElement *eTypedElementImpl) SetUnique(newIsUnique bool) {
	oldIsUnique := eTypedElement.isUnique
	eTypedElement.isUnique = newIsUnique
	if eTypedElement.ENotificationRequired() {
		eTypedElement.ENotify(NewNotificationByFeatureID(eTypedElement.GetEObject(), SET, ETYPED_ELEMENT__UNIQUE, oldIsUnique, newIsUnique, NO_INDEX))
	}
}

// GetLowerBound get the value of lowerBound
func (eTypedElement *eTypedElementImpl) GetLowerBound() int {
	return eTypedElement.lowerBound
}

// SetLowerBound set the value of lowerBound
func (eTypedElement *eTypedElementImpl) SetLowerBound(newLowerBound int) {
	oldLowerBound := eTypedElement.lowerBound
	eTypedElement.lowerBound = newLowerBound
	if eTypedElement.ENotificationRequired() {
		eTypedElement.ENotify(NewNotificationByFeatureID(eTypedElement.GetEObject(), SET, ETYPED_ELEMENT__LOWER_BOUND, oldLowerBound, newLowerBound, NO_INDEX))
	}
}

// GetUpperBound get the value of upperBound
func (eTypedElement *eTypedElementImpl) GetUpperBound() int {
	return eTypedElement.upperBound
}

// SetUpperBound set the value of upperBound
func (eTypedElement *eTypedElementImpl) SetUpperBound(newUpperBound int) {
	oldUpperBound := eTypedElement.upperBound
	eTypedElement.upperBound = newUpperBound
	if eTypedElement.ENotificationRequired() {
		eTypedElement.ENotify(NewNotificationByFeatureID(eTypedElement.GetEObject(), SET, ETYPED_ELEMENT__UPPER_BOUND, oldUpperBound, newUpperBound, NO_INDEX))
	}
}

func (eTypedElement *eTypedElementImpl) EGetFromID(featureID int, resolve, coreType bool) interface{} {
	switch featureID {
	case ETYPED_ELEMENT__ETYPE:
		return eTypedElement.GetEType()
	case ETYPED_ELEMENT__LOWER_BOUND:
		return eTypedElement.GetLowerBound()
	case ETYPED_ELEMENT__MANY:
		return eTypedElement.IsMany()
	case ETYPED_ELEMENT__ORDERED:
		return eTypedElement.IsOrdered()
	case ETYPED_ELEMENT__REQUIRED:
		return eTypedElement.IsRequired()
	case ETYPED_ELEMENT__UNIQUE:
		return eTypedElement.IsUnique()
	case ETYPED_ELEMENT__UPPER_BOUND:
		return eTypedElement.GetUpperBound()
	default:
		return eTypedElement.eNamedElementImpl.EGetFromID(featureID, resolve, coreType)
	}
}

func (eTypedElement *eTypedElementImpl) ESetFromID(featureID int, newValue interface{}) {
	switch featureID {
	case ETYPED_ELEMENT__ETYPE:
		e := newValue.(EClassifier)
		eTypedElement.SetEType(e)
	case ETYPED_ELEMENT__LOWER_BOUND:
		l := newValue.(int)
		eTypedElement.SetLowerBound(l)
	case ETYPED_ELEMENT__ORDERED:
		o := newValue.(bool)
		eTypedElement.SetOrdered(o)
	case ETYPED_ELEMENT__UNIQUE:
		u := newValue.(bool)
		eTypedElement.SetUnique(u)
	case ETYPED_ELEMENT__UPPER_BOUND:
		u := newValue.(int)
		eTypedElement.SetUpperBound(u)
	default:
		eTypedElement.eNamedElementImpl.ESetFromID(featureID, newValue)
	}
}

func (eTypedElement *eTypedElementImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case ETYPED_ELEMENT__ETYPE:
		eTypedElement.UnsetEType()
	case ETYPED_ELEMENT__LOWER_BOUND:
		eTypedElement.SetLowerBound(0)
	case ETYPED_ELEMENT__ORDERED:
		eTypedElement.SetOrdered(true)
	case ETYPED_ELEMENT__UNIQUE:
		eTypedElement.SetUnique(true)
	case ETYPED_ELEMENT__UPPER_BOUND:
		eTypedElement.SetUpperBound(1)
	default:
		eTypedElement.eNamedElementImpl.EUnsetFromID(featureID)
	}
}

func (eTypedElement *eTypedElementImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case ETYPED_ELEMENT__ETYPE:
		return eTypedElement.eType != nil
	case ETYPED_ELEMENT__LOWER_BOUND:
		return eTypedElement.lowerBound != 0
	case ETYPED_ELEMENT__MANY:
		return eTypedElement.IsMany() != false
	case ETYPED_ELEMENT__ORDERED:
		return eTypedElement.isOrdered != true
	case ETYPED_ELEMENT__REQUIRED:
		return eTypedElement.IsRequired() != false
	case ETYPED_ELEMENT__UNIQUE:
		return eTypedElement.isUnique != true
	case ETYPED_ELEMENT__UPPER_BOUND:
		return eTypedElement.upperBound != 1
	default:
		return eTypedElement.eNamedElementImpl.EIsSetFromID(featureID)
	}
}
