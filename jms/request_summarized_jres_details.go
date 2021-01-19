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

// RequestSummarizedJresDetails Parameters for summarizing observed JREs for a specified time period.
type RequestSummarizedJresDetails struct {

	// The start of the time window in which resources are searched in RFC3339 format. Defaults to current time minus seven days.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// The end of the time window in which resources are searched in RFC3339 format. Defaults to current time.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The vendor of the JRE
	JreVendor *string `mandatory:"false" json:"jreVendor"`

	// The distribution of the JRE
	JreDistribution *string `mandatory:"false" json:"jreDistribution"`

	// The version of the JRE
	JreVersion *string `mandatory:"false" json:"jreVersion"`

	// The id of the related application
	ApplicationId *string `mandatory:"false" json:"applicationId"`

	// The id of the related managed instance
	ManagedInstanceId *string `mandatory:"false" json:"managedInstanceId"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder SortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// The field to sort by jre views. Only one sort order may be provided. Default order for firstSeen, lastSeen and version is descending. Default order for distribution and vendor is ascending. If no value is specified timeLastSeen is default.
	SortBy JreSortByEnum `mandatory:"false" json:"sortBy,omitempty"`
}

func (m RequestSummarizedJresDetails) String() string {
	return common.PointerString(m)
}
