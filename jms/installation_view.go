// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets/Query API
//
// Java Management Service Fleets/Query API
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v33/common"
)

// InstallationView A view resource on an installation in the specified time period. An installation is a collection of deployed instances of a sepcific JRE that share the same install path.
type InstallationView struct {

	// The vendor of the Java Runtime that is deployed with the installation
	JreVendor *string `mandatory:"true" json:"jreVendor"`

	// The distribution of the Java Runtime that is deployed with the installation
	JreDistribution *string `mandatory:"true" json:"jreDistribution"`

	// The version of the Java Runtime that is deployed with the installation
	JreVersion *string `mandatory:"true" json:"jreVersion"`

	// The path to the installation
	Path *string `mandatory:"true" json:"path"`

	// The Operating System for the the installation
	Os *string `mandatory:"true" json:"os"`

	// The architecture of the operating system for the installation
	Architecture *string `mandatory:"true" json:"architecture"`

	// Lower bound of the user specified time interval filter.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Upper bound of the user specified time interval filter.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The time the resource was first seen within the scope of filtering, but potentially outside the time interval provided by the filters.
	TimeFirstSeen *common.SDKTime `mandatory:"false" json:"timeFirstSeen"`

	// The time the resource was last seen within the scope of filtering, but potentially outside the time interval provided by the filters.
	TimeLastSeen *common.SDKTime `mandatory:"false" json:"timeLastSeen"`
}

func (m InstallationView) String() string {
	return common.PointerString(m)
}
