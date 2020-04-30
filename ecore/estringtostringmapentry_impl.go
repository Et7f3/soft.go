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

// eStringToStringMapEntryImpl is the implementation of the model object 'EStringToStringMapEntry'
type eStringToStringMapEntryImpl struct {
	*EObjectImpl
	key   string
	value string
}

// newEStringToStringMapEntryImpl is the constructor of a eStringToStringMapEntryImpl
func newEStringToStringMapEntryImpl() *eStringToStringMapEntryImpl {
	eStringToStringMapEntry := new(eStringToStringMapEntryImpl)
	eStringToStringMapEntry.EObjectImpl = NewEObjectImpl()
	eStringToStringMapEntry.SetInterfaces(eStringToStringMapEntry)
	eStringToStringMapEntry.key = ""
	eStringToStringMapEntry.value = ""

	return eStringToStringMapEntry
}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) EStaticClass() EClass {
	return GetPackage().GetEStringToStringMapEntry()
}

// GetKey get the value of key
func (eStringToStringMapEntry *eStringToStringMapEntryImpl) GetKey() string {
	return eStringToStringMapEntry.key
}

// SetKey set the value of key
func (eStringToStringMapEntry *eStringToStringMapEntryImpl) SetKey(newKey string) {
	oldKey := eStringToStringMapEntry.key
	eStringToStringMapEntry.key = newKey
	if eStringToStringMapEntry.ENotificationRequired() {
		eStringToStringMapEntry.ENotify(NewNotificationByFeatureID(eStringToStringMapEntry.AsEObject(), SET, ESTRING_TO_STRING_MAP_ENTRY__KEY, oldKey, newKey, NO_INDEX))
	}
}

// GetValue get the value of value
func (eStringToStringMapEntry *eStringToStringMapEntryImpl) GetValue() string {
	return eStringToStringMapEntry.value
}

// SetValue set the value of value
func (eStringToStringMapEntry *eStringToStringMapEntryImpl) SetValue(newValue string) {
	oldValue := eStringToStringMapEntry.value
	eStringToStringMapEntry.value = newValue
	if eStringToStringMapEntry.ENotificationRequired() {
		eStringToStringMapEntry.ENotify(NewNotificationByFeatureID(eStringToStringMapEntry.AsEObject(), SET, ESTRING_TO_STRING_MAP_ENTRY__VALUE, oldValue, newValue, NO_INDEX))
	}
}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) EGetFromID(featureID int, resolve, coreType bool) interface{} {
	switch featureID {
	case ESTRING_TO_STRING_MAP_ENTRY__KEY:
		return eStringToStringMapEntry.GetKey()
	case ESTRING_TO_STRING_MAP_ENTRY__VALUE:
		return eStringToStringMapEntry.GetValue()
	default:
		return eStringToStringMapEntry.EObjectImpl.EGetFromID(featureID, resolve, coreType)
	}
}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) ESetFromID(featureID int, newValue interface{}) {
	switch featureID {
	case ESTRING_TO_STRING_MAP_ENTRY__KEY:
		k := newValue.(string)
		eStringToStringMapEntry.SetKey(k)
	case ESTRING_TO_STRING_MAP_ENTRY__VALUE:
		v := newValue.(string)
		eStringToStringMapEntry.SetValue(v)
	default:
		eStringToStringMapEntry.EObjectImpl.ESetFromID(featureID, newValue)
	}
}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case ESTRING_TO_STRING_MAP_ENTRY__KEY:
		eStringToStringMapEntry.SetKey("")
	case ESTRING_TO_STRING_MAP_ENTRY__VALUE:
		eStringToStringMapEntry.SetValue("")
	default:
		eStringToStringMapEntry.EObjectImpl.EUnsetFromID(featureID)
	}
}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case ESTRING_TO_STRING_MAP_ENTRY__KEY:
		return eStringToStringMapEntry.key != ""
	case ESTRING_TO_STRING_MAP_ENTRY__VALUE:
		return eStringToStringMapEntry.value != ""
	default:
		return eStringToStringMapEntry.EObjectImpl.EIsSetFromID(featureID)
	}
}
