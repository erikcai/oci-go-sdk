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

// RequestSummarizedApplicationsDetails Parameters for summarizing observed applications for a specified time period.
type RequestSummarizedApplicationsDetails struct {

	// The start of the time window in which resources are searched in RFC3339 format. Defaults to current time minus seven days.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// The end of the time window in which resources are searched in RFC3339 format. Defaults to current time.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The display name of the application
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The installation path of the related installation
	InstallationPath *string `mandatory:"false" json:"installationPath"`

	// The vendor of the related JRE
	JreVendor *string `mandatory:"false" json:"jreVendor"`

	// The distribution of the related JRE
	JreDistribution *string `mandatory:"false" json:"jreDistribution"`

	// The version of the related JRE
	JreVersion *string `mandatory:"false" json:"jreVersion"`

	// The id of the application
	ApplicationId *string `mandatory:"false" json:"applicationId"`

	// The way the application was started
	ApplicationType *string `mandatory:"false" json:"applicationType"`

	// The id of the related managed instance
	ManagedInstanceId *string `mandatory:"false" json:"managedInstanceId"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder SortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// The field to sort by application views. Only one sort order may be provided. Default order for firstSeen and lastSeen is descending. Default order for displayName is ascending. If no value is specified timeLastSeen is default.
	SortBy ApplicationSortByEnum `mandatory:"false" json:"sortBy,omitempty"`
}

func (m RequestSummarizedApplicationsDetails) String() string {
	return common.PointerString(m)
}
