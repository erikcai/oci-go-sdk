// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

// PackageTypeEnumEnum Enum with underlying type: string
type PackageTypeEnumEnum string

// Set of constants representing the allowable values for PackageTypeEnumEnum
const (
	PackageTypeEnumImage PackageTypeEnumEnum = "IMAGE"
)

var mappingPackageTypeEnum = map[string]PackageTypeEnumEnum{
	"IMAGE": PackageTypeEnumImage,
}

// GetPackageTypeEnumEnumValues Enumerates the set of values for PackageTypeEnumEnum
func GetPackageTypeEnumEnumValues() []PackageTypeEnumEnum {
	values := make([]PackageTypeEnumEnum, 0)
	for _, v := range mappingPackageTypeEnum {
		values = append(values, v)
	}
	return values
}
