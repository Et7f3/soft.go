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

type EcoreFactory interface {
	CreateEAnnotation() EAnnotation
	CreateEAnnotationFromContainer(eContainer EModelElement) EAnnotation
	CreateEAttribute() EAttribute
	CreateEAttributeFromContainer(eContainer EClass) EAttribute
	CreateEAttributeFromContainerAndClassID(eContainer EClass, classID int) EAttribute
	CreateEClass() EClass
	CreateEClassFromContainer(eContainer EPackage) EClass
	CreateEClassFromContainerAndClassID(eContainer EPackage, classID int) EClass
	CreateEDataType() EDataType
	CreateEDataTypeFromContainer(eContainer EPackage) EDataType
	CreateEDataTypeFromContainerAndClassID(eContainer EPackage, classID int) EDataType
	CreateEEnum() EEnum
	CreateEEnumFromContainer(eContainer EPackage) EEnum
	CreateEEnumFromContainerAndClassID(eContainer EPackage, classID int) EEnum
	CreateEEnumLiteral() EEnumLiteral
	CreateEEnumLiteralFromContainer(eContainer EEnum) EEnumLiteral
	CreateEFactory() EFactory
	CreateEFactoryFromContainer(eContainer EPackage) EFactory
	CreateEGenericType() EGenericType
	CreateEObject() EObject
	CreateEOperation() EOperation
	CreateEOperationFromContainer(eContainer EClass) EOperation
	CreateEOperationFromContainerAndClassID(eContainer EClass, classID int) EOperation
	CreateEPackage() EPackage
	CreateEPackageFromContainer(eContainer EPackage) EPackage
	CreateEParameter() EParameter
	CreateEParameterFromContainer(eContainer EOperation) EParameter
	CreateEReference() EReference
	CreateEReferenceFromContainer(eContainer EClass) EReference
	CreateEReferenceFromContainerAndClassID(eContainer EClass, classID int) EReference
	CreateEStringToStringMapEntry() EStringToStringMapEntry
	CreateETypeParameter() ETypeParameter
}

var factoryInstance EcoreFactory

// GetFactory returns the factory for the model ecore
func GetFactory() EcoreFactory {
	if factoryInstance == nil {
		factoryInstance = newEcoreFactoryExt()
	}
	return factoryInstance
}
