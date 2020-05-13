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

// eEnumLiteralImpl is the implementation of the model object 'EEnumLiteral'
type eEnumLiteralImpl struct {
	*eNamedElementImpl
	eEnum    EEnum
	instance interface{}
	literal  string
	value    int
}

// newEEnumLiteralImpl is the constructor of a eEnumLiteralImpl
func newEEnumLiteralImpl() *eEnumLiteralImpl {
	eEnumLiteral := new(eEnumLiteralImpl)
	eEnumLiteral.eNamedElementImpl = newENamedElementImpl()
	eEnumLiteral.SetInterfaces(eEnumLiteral)
	eEnumLiteral.instance = nil
	eEnumLiteral.literal = ""
	eEnumLiteral.value = 0

	return eEnumLiteral
}

func (eEnumLiteral *eEnumLiteralImpl) asEEnumLiteral() EEnumLiteral {
	return eEnumLiteral.GetInterfaces().(EEnumLiteral)
}

func (eEnumLiteral *eEnumLiteralImpl) EStaticClass() EClass {
	return GetPackage().GetEEnumLiteral()
}

// GetEEnum get the value of eEnum
func (eEnumLiteral *eEnumLiteralImpl) GetEEnum() EEnum {
	if eEnumLiteral.EContainerFeatureID() == EENUM_LITERAL__EENUM {
		return eEnumLiteral.EContainer().(EEnum)
	}
	return nil
}

// GetInstance get the value of instance
func (eEnumLiteral *eEnumLiteralImpl) GetInstance() interface{} {
	return eEnumLiteral.instance
}

// SetInstance set the value of instance
func (eEnumLiteral *eEnumLiteralImpl) SetInstance(newInstance interface{}) {
	oldInstance := eEnumLiteral.instance
	eEnumLiteral.instance = newInstance
	if eEnumLiteral.ENotificationRequired() {
		eEnumLiteral.ENotify(NewNotificationByFeatureID(eEnumLiteral.AsEObject(), SET, EENUM_LITERAL__INSTANCE, oldInstance, newInstance, NO_INDEX))
	}
}

// GetLiteral get the value of literal
func (eEnumLiteral *eEnumLiteralImpl) GetLiteral() string {
	return eEnumLiteral.literal
}

// SetLiteral set the value of literal
func (eEnumLiteral *eEnumLiteralImpl) SetLiteral(newLiteral string) {
	oldLiteral := eEnumLiteral.literal
	eEnumLiteral.literal = newLiteral
	if eEnumLiteral.ENotificationRequired() {
		eEnumLiteral.ENotify(NewNotificationByFeatureID(eEnumLiteral.AsEObject(), SET, EENUM_LITERAL__LITERAL, oldLiteral, newLiteral, NO_INDEX))
	}
}

// GetValue get the value of value
func (eEnumLiteral *eEnumLiteralImpl) GetValue() int {
	return eEnumLiteral.value
}

// SetValue set the value of value
func (eEnumLiteral *eEnumLiteralImpl) SetValue(newValue int) {
	oldValue := eEnumLiteral.value
	eEnumLiteral.value = newValue
	if eEnumLiteral.ENotificationRequired() {
		eEnumLiteral.ENotify(NewNotificationByFeatureID(eEnumLiteral.AsEObject(), SET, EENUM_LITERAL__VALUE, oldValue, newValue, NO_INDEX))
	}
}

func (eEnumLiteral *eEnumLiteralImpl) EGetFromID(featureID int, resolve bool) interface{} {
	switch featureID {
	case EENUM_LITERAL__EENUM:
		return eEnumLiteral.asEEnumLiteral().GetEEnum()
	case EENUM_LITERAL__INSTANCE:
		return eEnumLiteral.asEEnumLiteral().GetInstance()
	case EENUM_LITERAL__LITERAL:
		return eEnumLiteral.asEEnumLiteral().GetLiteral()
	case EENUM_LITERAL__VALUE:
		return eEnumLiteral.asEEnumLiteral().GetValue()
	default:
		return eEnumLiteral.eNamedElementImpl.EGetFromID(featureID, resolve)
	}
}

func (eEnumLiteral *eEnumLiteralImpl) ESetFromID(featureID int, newValue interface{}) {
	switch featureID {
	case EENUM_LITERAL__INSTANCE:
		eEnumLiteral.asEEnumLiteral().SetInstance(newValue.(interface{}))
	case EENUM_LITERAL__LITERAL:
		eEnumLiteral.asEEnumLiteral().SetLiteral(newValue.(string))
	case EENUM_LITERAL__VALUE:
		eEnumLiteral.asEEnumLiteral().SetValue(newValue.(int))
	default:
		eEnumLiteral.eNamedElementImpl.ESetFromID(featureID, newValue)
	}
}

func (eEnumLiteral *eEnumLiteralImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case EENUM_LITERAL__INSTANCE:
		eEnumLiteral.asEEnumLiteral().SetInstance(nil)
	case EENUM_LITERAL__LITERAL:
		eEnumLiteral.asEEnumLiteral().SetLiteral("")
	case EENUM_LITERAL__VALUE:
		eEnumLiteral.asEEnumLiteral().SetValue(0)
	default:
		eEnumLiteral.eNamedElementImpl.EUnsetFromID(featureID)
	}
}

func (eEnumLiteral *eEnumLiteralImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case EENUM_LITERAL__EENUM:
		return eEnumLiteral.GetEEnum() != nil
	case EENUM_LITERAL__INSTANCE:
		return eEnumLiteral.GetInstance() != ""
	case EENUM_LITERAL__LITERAL:
		return eEnumLiteral.literal != ""
	case EENUM_LITERAL__VALUE:
		return eEnumLiteral.value != 0
	default:
		return eEnumLiteral.eNamedElementImpl.EIsSetFromID(featureID)
	}
}

func (eEnumLiteral *eEnumLiteralImpl) EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EENUM_LITERAL__EENUM:
		msgs := notifications
		if eEnumLiteral.EContainer() != nil {
			msgs = eEnumLiteral.EBasicRemoveFromContainer(msgs)
		}
		return eEnumLiteral.EBasicSetContainer(otherEnd, EENUM_LITERAL__EENUM, msgs)
	default:
		return eEnumLiteral.eNamedElementImpl.EBasicInverseAdd(otherEnd, featureID, notifications)
	}
}

func (eEnumLiteral *eEnumLiteralImpl) EBasicInverseRemove(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EENUM_LITERAL__EENUM:
		return eEnumLiteral.EBasicSetContainer(nil, EENUM_LITERAL__EENUM, notifications)
	default:
		return eEnumLiteral.eNamedElementImpl.EBasicInverseRemove(otherEnd, featureID, notifications)
	}
}
