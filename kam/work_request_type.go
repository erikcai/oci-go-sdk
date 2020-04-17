// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// KAM API
//
// description: |
//   Kubernetes Add-on Manager API for installing, uninstalling and upgrading
//   OKE add-ons or Marketplace applications on an OKE cluster
//

package kam

// WorkRequestTypeEnum Enum with underlying type: string
type WorkRequestTypeEnum string

// Set of constants representing the allowable values for WorkRequestTypeEnum
const (
	WorkRequestTypeInstall   WorkRequestTypeEnum = "INSTALL"
	WorkRequestTypeUninstall WorkRequestTypeEnum = "UNINSTALL"
	WorkRequestTypeUpgrade   WorkRequestTypeEnum = "UPGRADE"
)

var mappingWorkRequestType = map[string]WorkRequestTypeEnum{
	"INSTALL":   WorkRequestTypeInstall,
	"UNINSTALL": WorkRequestTypeUninstall,
	"UPGRADE":   WorkRequestTypeUpgrade,
}

// GetWorkRequestTypeEnumValues Enumerates the set of values for WorkRequestTypeEnum
func GetWorkRequestTypeEnumValues() []WorkRequestTypeEnum {
	values := make([]WorkRequestTypeEnum, 0)
	for _, v := range mappingWorkRequestType {
		values = append(values, v)
	}
	return values
}
