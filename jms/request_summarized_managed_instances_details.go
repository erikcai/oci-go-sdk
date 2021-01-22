// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets/Query API
//
// Java Management Service Fleets/Query API
//

package jms

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// RequestSummarizedManagedInstancesDetails Parameters for summarizing observed managed instances for a specified time period.
type RequestSummarizedManagedInstancesDetails struct {

	// The start of the time window in which resources are searched in RFC3339 format. Defaults to current time minus seven days.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// The end of the time window in which resources are searched in RFC3339 format. Defaults to current time.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The installation path of the related installation
	InstallationPath *string `mandatory:"false" json:"installationPath"`

	// The vendor of the related JRE
	JreVendor *string `mandatory:"false" json:"jreVendor"`

	// The distribution of the related JRE
	JreDistribution *string `mandatory:"false" json:"jreDistribution"`

	// The version of the related JRE
	JreVersion *string `mandatory:"false" json:"jreVersion"`

	// The id of the related application
	ApplicationId *string `mandatory:"false" json:"applicationId"`

	// The id of the managed instance
	ManagedInstanceId *string `mandatory:"false" json:"managedInstanceId"`

	// type of the managed instance
	ManagedInstanceType ManagedInstanceTypeEnum `mandatory:"false" json:"managedInstanceType,omitempty"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder SortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// The field to sort by the managed instance views. Only one sort order may be provided. Default order for firstSeen and lastSeen is descending. If no value is specified timeLastSeen is default.
	SortBy ManagedInstanceSortByEnum `mandatory:"false" json:"sortBy,omitempty"`
}

func (m RequestSummarizedManagedInstancesDetails) String() string {
	return common.PointerString(m)
}
