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

// JreView A view resource on the Java Runtime Environment in the specified time period. A Java Runtime Environment (JRE) is a certain distribution of a Java Runtime identified by its vendor and version
type JreView struct {

	// The vendor of the Java Runtime
	Vendor *string `mandatory:"true" json:"vendor"`

	// The distribution of a Java Runtime is the name of the lineage of products it belongs to, for example Java(TM) SE Runtime Environment
	Distribution *string `mandatory:"true" json:"distribution"`

	// The version of the Java Runtime
	Version *string `mandatory:"true" json:"version"`

	// Lower bound of the user specified time interval filter.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Upper bound of the user specified time interval filter.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The time the resource was first seen within the scope of filtering, but potentially outside the time interval provided by the filters.
	TimeFirstSeen *common.SDKTime `mandatory:"false" json:"timeFirstSeen"`

	// The time the resource was last seen within the scope of filtering, but potentially outside the time interval provided by the filters.
	TimeLastSeen *common.SDKTime `mandatory:"false" json:"timeLastSeen"`
}

func (m JreView) String() string {
	return common.PointerString(m)
}
