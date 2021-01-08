// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets/Query API
//
// Java Management Service Fleets/Query API
//

package jms

// InstallationSortByEnum Enum with underlying type: string
type InstallationSortByEnum string

// Set of constants representing the allowable values for InstallationSortByEnum
const (
	InstallationSortByTimeFirstSeen   InstallationSortByEnum = "timeFirstSeen"
	InstallationSortByTimeLastSeen    InstallationSortByEnum = "timeLastSeen"
	InstallationSortByJreVersion      InstallationSortByEnum = "jreVersion"
	InstallationSortByJreDistribution InstallationSortByEnum = "jreDistribution"
	InstallationSortByJreVendor       InstallationSortByEnum = "jreVendor"
	InstallationSortByPath            InstallationSortByEnum = "path"
)

var mappingInstallationSortBy = map[string]InstallationSortByEnum{
	"timeFirstSeen":   InstallationSortByTimeFirstSeen,
	"timeLastSeen":    InstallationSortByTimeLastSeen,
	"jreVersion":      InstallationSortByJreVersion,
	"jreDistribution": InstallationSortByJreDistribution,
	"jreVendor":       InstallationSortByJreVendor,
	"path":            InstallationSortByPath,
}

// GetInstallationSortByEnumValues Enumerates the set of values for InstallationSortByEnum
func GetInstallationSortByEnumValues() []InstallationSortByEnum {
	values := make([]InstallationSortByEnum, 0)
	for _, v := range mappingInstallationSortBy {
		values = append(values, v)
	}
	return values
}
