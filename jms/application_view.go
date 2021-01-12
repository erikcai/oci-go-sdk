// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets/Query API
//
// Java Management Service Fleets/Query API
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v32/common"
)

// ApplicationView A view resource on the application in the specified time period. An application is a Java application that can be executed by JRE installations. Application is independent to the Java Runtime or the installation.
type ApplicationView struct {

	// An internal identifier for the application that is unique to the fleet
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// The name of the application
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of the application, which denotes how the application was started
	ApplicationType *string `mandatory:"true" json:"applicationType"`

	// Lower bound of the user specified time interval filter.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Upper bound of the user specified time interval filter.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The time the resource was first seen within the scope of filtering, but potentially outside the time interval provided by the filters.
	TimeFirstSeen *common.SDKTime `mandatory:"false" json:"timeFirstSeen"`

	// The time the resource was last seen within the scope of filtering, but potentially outside the time interval provided by the filters.
	TimeLastSeen *common.SDKTime `mandatory:"false" json:"timeLastSeen"`
}

func (m ApplicationView) String() string {
	return common.PointerString(m)
}
