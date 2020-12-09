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

import "strconv"

type ecoreFactoryInternal interface {
	createEBigDecimalFromString(eDataType EDataType, literalValue string) interface{}
	createEBigIntegerFromString(eDataType EDataType, literalValue string) interface{}
	createEBooleanFromString(eDataType EDataType, literalValue string) interface{}
	createEBooleanObjectFromString(eDataType EDataType, literalValue string) interface{}
	createEByteFromString(eDataType EDataType, literalValue string) interface{}
	createEByteArrayFromString(eDataType EDataType, literalValue string) interface{}
	createEByteObjectFromString(eDataType EDataType, literalValue string) interface{}
	createECharFromString(eDataType EDataType, literalValue string) interface{}
	createECharacterObjectFromString(eDataType EDataType, literalValue string) interface{}
	createEDateFromString(eDataType EDataType, literalValue string) interface{}
	createEDoubleFromString(eDataType EDataType, literalValue string) interface{}
	createEDoubleObjectFromString(eDataType EDataType, literalValue string) interface{}
	createEFloatFromString(eDataType EDataType, literalValue string) interface{}
	createEFloatObjectFromString(eDataType EDataType, literalValue string) interface{}
	createEIntFromString(eDataType EDataType, literalValue string) interface{}
	createEIntegerObjectFromString(eDataType EDataType, literalValue string) interface{}
	createEJavaClassFromString(eDataType EDataType, literalValue string) interface{}
	createEJavaObjectFromString(eDataType EDataType, literalValue string) interface{}
	createELongFromString(eDataType EDataType, literalValue string) interface{}
	createELongObjectFromString(eDataType EDataType, literalValue string) interface{}
	createEShortFromString(eDataType EDataType, literalValue string) interface{}
	createEShortObjectFromString(eDataType EDataType, literalValue string) interface{}
	createEStringFromString(eDataType EDataType, literalValue string) interface{}
	convertEBigDecimalToString(eDataType EDataType, literalValue interface{}) string
	convertEBigIntegerToString(eDataType EDataType, literalValue interface{}) string
	convertEBooleanToString(eDataType EDataType, literalValue interface{}) string
	convertEBooleanObjectToString(eDataType EDataType, literalValue interface{}) string
	convertEByteToString(eDataType EDataType, literalValue interface{}) string
	convertEByteArrayToString(eDataType EDataType, literalValue interface{}) string
	convertEByteObjectToString(eDataType EDataType, literalValue interface{}) string
	convertECharToString(eDataType EDataType, literalValue interface{}) string
	convertECharacterObjectToString(eDataType EDataType, literalValue interface{}) string
	convertEDateToString(eDataType EDataType, literalValue interface{}) string
	convertEDoubleToString(eDataType EDataType, literalValue interface{}) string
	convertEDoubleObjectToString(eDataType EDataType, literalValue interface{}) string
	convertEFloatToString(eDataType EDataType, literalValue interface{}) string
	convertEFloatObjectToString(eDataType EDataType, literalValue interface{}) string
	convertEIntToString(eDataType EDataType, literalValue interface{}) string
	convertEIntegerObjectToString(eDataType EDataType, literalValue interface{}) string
	convertEJavaClassToString(eDataType EDataType, literalValue interface{}) string
	convertEJavaObjectToString(eDataType EDataType, literalValue interface{}) string
	convertELongToString(eDataType EDataType, literalValue interface{}) string
	convertELongObjectToString(eDataType EDataType, literalValue interface{}) string
	convertEShortToString(eDataType EDataType, literalValue interface{}) string
	convertEShortObjectToString(eDataType EDataType, literalValue interface{}) string
	convertEStringToString(eDataType EDataType, literalValue interface{}) string
}

type ecoreFactoryImpl struct {
	*EFactoryExt
}

func newEcoreFactoryImpl() *ecoreFactoryImpl {
	factory := new(ecoreFactoryImpl)
	factory.EFactoryExt = NewEFactoryExt()
	factory.SetInterfaces(factory)
	return factory
}

func (ecoreFactoryImpl *ecoreFactoryImpl) Create(eClass EClass) EObject {
	classID := eClass.GetClassifierID()
	switch classID {
	case EANNOTATION:
		return ecoreFactoryImpl.CreateEAnnotation()
	case EATTRIBUTE:
		return ecoreFactoryImpl.CreateEAttribute()
	case ECLASS:
		return ecoreFactoryImpl.CreateEClass()
	case EDATA_TYPE:
		return ecoreFactoryImpl.CreateEDataType()
	case EENUM:
		return ecoreFactoryImpl.CreateEEnum()
	case EENUM_LITERAL:
		return ecoreFactoryImpl.CreateEEnumLiteral()
	case EFACTORY:
		return ecoreFactoryImpl.CreateEFactory()
	case EGENERIC_TYPE:
		return ecoreFactoryImpl.CreateEGenericType()
	case EOBJECT:
		return ecoreFactoryImpl.CreateEObject()
	case EOPERATION:
		return ecoreFactoryImpl.CreateEOperation()
	case EPACKAGE:
		return ecoreFactoryImpl.CreateEPackage()
	case EPARAMETER:
		return ecoreFactoryImpl.CreateEParameter()
	case EREFERENCE:
		return ecoreFactoryImpl.CreateEReference()
	case ESTRING_TO_STRING_MAP_ENTRY:
		return ecoreFactoryImpl.CreateEStringToStringMapEntry()
	case ETYPE_PARAMETER:
		return ecoreFactoryImpl.CreateETypeParameter()
	default:
		panic("Create: " + strconv.Itoa(classID) + " not found")
	}
}

func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEAnnotation() EAnnotation {
	return newEAnnotationImpl()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEAnnotationFromContainer(eContainer EModelElement) EAnnotation {
	element := newEAnnotationImpl()
	if eContainer != nil {
		eContainer.GetEAnnotations().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEAttribute() EAttribute {
	return newEAttributeExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEAttributeFromContainer(eContainer EClass) EAttribute {
	element := newEAttributeExt()
	if eContainer != nil {
		eContainer.GetEStructuralFeatures().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEAttributeFromContainerAndClassID(eContainer EClass, classID int) EAttribute {
	element := newEAttributeExt()
	element.SetFeatureID(classID)
	if eContainer != nil {
		eContainer.GetEStructuralFeatures().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEClass() EClass {
	return newEClassExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEClassFromContainer(eContainer EPackage) EClass {
	element := newEClassExt()
	if eContainer != nil {
		eContainer.GetEClassifiers().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEClassFromContainerAndClassID(eContainer EPackage, classID int) EClass {
	element := newEClassExt()
	element.SetClassifierID(classID)
	if eContainer != nil {
		eContainer.GetEClassifiers().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEDataType() EDataType {
	return newEDataTypeExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEDataTypeFromContainer(eContainer EPackage) EDataType {
	element := newEDataTypeExt()
	if eContainer != nil {
		eContainer.GetEClassifiers().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEDataTypeFromContainerAndClassID(eContainer EPackage, classID int) EDataType {
	element := newEDataTypeExt()
	element.SetClassifierID(classID)
	if eContainer != nil {
		eContainer.GetEClassifiers().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEEnum() EEnum {
	return newEEnumImpl()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEEnumFromContainer(eContainer EPackage) EEnum {
	element := newEEnumImpl()
	if eContainer != nil {
		eContainer.GetEClassifiers().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEEnumFromContainerAndClassID(eContainer EPackage, classID int) EEnum {
	element := newEEnumImpl()
	element.SetClassifierID(classID)
	if eContainer != nil {
		eContainer.GetEClassifiers().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEEnumLiteral() EEnumLiteral {
	return newEEnumLiteralImpl()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEEnumLiteralFromContainer(eContainer EEnum) EEnumLiteral {
	element := newEEnumLiteralImpl()
	if eContainer != nil {
		eContainer.GetELiterals().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEFactory() EFactory {
	return NewEFactoryExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEFactoryFromContainer(eContainer EPackage) EFactory {
	element := NewEFactoryExt()
	if eContainer != nil {
		eContainer.SetEFactoryInstance(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEGenericType() EGenericType {
	return newEGenericTypeImpl()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEObject() EObject {
	return NewEObjectImpl()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEOperation() EOperation {
	return newEOperationExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEOperationFromContainer(eContainer EClass) EOperation {
	element := newEOperationExt()
	if eContainer != nil {
		eContainer.GetEOperations().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEOperationFromContainerAndClassID(eContainer EClass, classID int) EOperation {
	element := newEOperationExt()
	element.SetOperationID(classID)
	if eContainer != nil {
		eContainer.GetEOperations().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEPackage() EPackage {
	return NewEPackageExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEPackageFromContainer(eContainer EPackage) EPackage {
	element := NewEPackageExt()
	if eContainer != nil {
		eContainer.GetESubPackages().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEParameter() EParameter {
	return newEParameterImpl()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEParameterFromContainer(eContainer EOperation) EParameter {
	element := newEParameterImpl()
	if eContainer != nil {
		eContainer.GetEParameters().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEReference() EReference {
	return newEReferenceExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEReferenceFromContainer(eContainer EClass) EReference {
	element := newEReferenceExt()
	if eContainer != nil {
		eContainer.GetEStructuralFeatures().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEReferenceFromContainerAndClassID(eContainer EClass, classID int) EReference {
	element := newEReferenceExt()
	element.SetFeatureID(classID)
	if eContainer != nil {
		eContainer.GetEStructuralFeatures().Add(element)
	}
	return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEStringToStringMapEntry() EStringToStringMapEntry {
	return newEStringToStringMapEntryImpl()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateETypeParameter() ETypeParameter {
	return newETypeParameterImpl()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateFromString(eDataType EDataType, literalValue string) interface{} {
	classID := eDataType.GetClassifierID()
	internal := ecoreFactoryImpl.GetInterfaces().(ecoreFactoryInternal)
	switch classID {
	case EBIG_DECIMAL:
		return internal.createEBigDecimalFromString(eDataType, literalValue)
	case EBIG_INTEGER:
		return internal.createEBigIntegerFromString(eDataType, literalValue)
	case EBOOLEAN:
		return internal.createEBooleanFromString(eDataType, literalValue)
	case EBOOLEAN_OBJECT:
		return internal.createEBooleanObjectFromString(eDataType, literalValue)
	case EBYTE:
		return internal.createEByteFromString(eDataType, literalValue)
	case EBYTE_ARRAY:
		return internal.createEByteArrayFromString(eDataType, literalValue)
	case EBYTE_OBJECT:
		return internal.createEByteObjectFromString(eDataType, literalValue)
	case ECHAR:
		return internal.createECharFromString(eDataType, literalValue)
	case ECHARACTER_OBJECT:
		return internal.createECharacterObjectFromString(eDataType, literalValue)
	case EDATE:
		return internal.createEDateFromString(eDataType, literalValue)
	case EDOUBLE:
		return internal.createEDoubleFromString(eDataType, literalValue)
	case EDOUBLE_OBJECT:
		return internal.createEDoubleObjectFromString(eDataType, literalValue)
	case EFLOAT:
		return internal.createEFloatFromString(eDataType, literalValue)
	case EFLOAT_OBJECT:
		return internal.createEFloatObjectFromString(eDataType, literalValue)
	case EINT:
		return internal.createEIntFromString(eDataType, literalValue)
	case EINTEGER_OBJECT:
		return internal.createEIntegerObjectFromString(eDataType, literalValue)
	case EJAVA_CLASS:
		return internal.createEJavaClassFromString(eDataType, literalValue)
	case EJAVA_OBJECT:
		return internal.createEJavaObjectFromString(eDataType, literalValue)
	case ELONG:
		return internal.createELongFromString(eDataType, literalValue)
	case ELONG_OBJECT:
		return internal.createELongObjectFromString(eDataType, literalValue)
	case ESHORT:
		return internal.createEShortFromString(eDataType, literalValue)
	case ESHORT_OBJECT:
		return internal.createEShortObjectFromString(eDataType, literalValue)
	case ESTRING:
		return internal.createEStringFromString(eDataType, literalValue)
	default:
		panic("The datatype '" + eDataType.GetName() + "' is not a valid classifier")
	}
}

func (ecoreFactoryImpl *ecoreFactoryImpl) ConvertToString(eDataType EDataType, instanceValue interface{}) string {
	classID := eDataType.GetClassifierID()
	internal := ecoreFactoryImpl.GetInterfaces().(ecoreFactoryInternal)
	switch classID {
	case EBIG_DECIMAL:
		return internal.convertEBigDecimalToString(eDataType, instanceValue)
	case EBIG_INTEGER:
		return internal.convertEBigIntegerToString(eDataType, instanceValue)
	case EBOOLEAN:
		return internal.convertEBooleanToString(eDataType, instanceValue)
	case EBOOLEAN_OBJECT:
		return internal.convertEBooleanObjectToString(eDataType, instanceValue)
	case EBYTE:
		return internal.convertEByteToString(eDataType, instanceValue)
	case EBYTE_ARRAY:
		return internal.convertEByteArrayToString(eDataType, instanceValue)
	case EBYTE_OBJECT:
		return internal.convertEByteObjectToString(eDataType, instanceValue)
	case ECHAR:
		return internal.convertECharToString(eDataType, instanceValue)
	case ECHARACTER_OBJECT:
		return internal.convertECharacterObjectToString(eDataType, instanceValue)
	case EDATE:
		return internal.convertEDateToString(eDataType, instanceValue)
	case EDOUBLE:
		return internal.convertEDoubleToString(eDataType, instanceValue)
	case EDOUBLE_OBJECT:
		return internal.convertEDoubleObjectToString(eDataType, instanceValue)
	case EFLOAT:
		return internal.convertEFloatToString(eDataType, instanceValue)
	case EFLOAT_OBJECT:
		return internal.convertEFloatObjectToString(eDataType, instanceValue)
	case EINT:
		return internal.convertEIntToString(eDataType, instanceValue)
	case EINTEGER_OBJECT:
		return internal.convertEIntegerObjectToString(eDataType, instanceValue)
	case EJAVA_CLASS:
		return internal.convertEJavaClassToString(eDataType, instanceValue)
	case EJAVA_OBJECT:
		return internal.convertEJavaObjectToString(eDataType, instanceValue)
	case ELONG:
		return internal.convertELongToString(eDataType, instanceValue)
	case ELONG_OBJECT:
		return internal.convertELongObjectToString(eDataType, instanceValue)
	case ESHORT:
		return internal.convertEShortToString(eDataType, instanceValue)
	case ESHORT_OBJECT:
		return internal.convertEShortObjectToString(eDataType, instanceValue)
	case ESTRING:
		return internal.convertEStringToString(eDataType, instanceValue)
	default:
		panic("The datatype '" + eDataType.GetName() + "' is not a valid classifier")
	}
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEBigDecimalFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEBigDecimalToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEBigIntegerFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEBigIntegerToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEBooleanFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEBooleanToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEBooleanObjectFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEBooleanObjectToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEByteFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEByteToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEByteArrayFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEByteArrayToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEByteObjectFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEByteObjectToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createECharFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertECharToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createECharacterObjectFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertECharacterObjectToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEDateFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEDateToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEDoubleFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEDoubleToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEDoubleObjectFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEDoubleObjectToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEFloatFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEFloatToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEFloatObjectFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEFloatObjectToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEIntFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEIntToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEIntegerObjectFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEIntegerObjectToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEJavaClassFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEJavaClassToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEJavaObjectFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEJavaObjectToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createELongFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertELongToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createELongObjectFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertELongObjectToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEShortFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEShortToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEShortObjectFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEShortObjectToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEStringFromString(eDataType EDataType, literalValue string) interface{} {
	panic("NotImplementedException")
}
func (ecoreFactoryImpl *ecoreFactoryImpl) convertEStringToString(eDataType EDataType, instanceValue interface{}) string {
	panic("NotImplementedException")
}
