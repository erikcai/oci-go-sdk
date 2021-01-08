// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets/Query API
//
// Java Management Service Fleets/Query API
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// ManagedInstanceView A view resource on the managed instance in the specified time period. Entities that emit usage events to the JMS are represented as managed instances. A Managed Instance has a unique identity which is used by the JMS in order to distinguish it from all other managed instances. A Managed Instance has a distinguishing tagging key-value pair(s) that associates it with its containing fleet. Currently, JMS supports only one kind of managed instance, a MACS Agent.
type ManagedInstanceView struct {

	// The ocid of the related managed instance
	ManagedInstanceId *string `mandatory:"true" json:"managedInstanceId"`

	// The type of the source
	ManagedInstanceType ManagedInstanceTypeEnum `mandatory:"true" json:"managedInstanceType"`

	// Lower bound of the user specified time interval filter.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Upper bound of the user specified time interval filter.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The time the resource was first seen within the scope of filtering, but potentially outside the time interval provided by the filters.
	TimeFirstSeen *common.SDKTime `mandatory:"false" json:"timeFirstSeen"`

	// The time the resource was last seen within the scope of filtering, but potentially outside the time interval provided by the filters.
	TimeLastSeen *common.SDKTime `mandatory:"false" json:"timeLastSeen"`
}

func (m ManagedInstanceView) String() string {
	return common.PointerString(m)
}
