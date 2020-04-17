// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

// OsFamiliesEnum Enum with underlying type: string
type OsFamiliesEnum string

// Set of constants representing the allowable values for OsFamiliesEnum
const (
	OsFamiliesLinux   OsFamiliesEnum = "LINUX"
	OsFamiliesWindows OsFamiliesEnum = "WINDOWS"
	OsFamiliesAll     OsFamiliesEnum = "ALL"
)

var mappingOsFamilies = map[string]OsFamiliesEnum{
	"LINUX":   OsFamiliesLinux,
	"WINDOWS": OsFamiliesWindows,
	"ALL":     OsFamiliesAll,
}

// GetOsFamiliesEnumValues Enumerates the set of values for OsFamiliesEnum
func GetOsFamiliesEnumValues() []OsFamiliesEnum {
	values := make([]OsFamiliesEnum, 0)
	for _, v := range mappingOsFamilies {
		values = append(values, v)
	}
	return values
}
