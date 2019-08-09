// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DataCatalog API
//
// A description of the DataCatalog API
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateJobDefinitionDetails Update information for a Job Definition resource.
type UpdateJobDefinitionDetails struct {

	// The display name of a user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Specifies if the Job Definition is incremental or full.
	IsIncremental *bool `mandatory:"false" json:"isIncremental"`

	// The key of the Data Asset for which the job is defined.
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`

	// Detailed description of the Job Definition.
	Description *string `mandatory:"false" json:"description"`

	// The key of the connection resource to be used for harvest, sampling, profiling  jobs.
	ConnectionKey *string `mandatory:"false" json:"connectionKey"`

	// Scope for the Job Definition. This property specifies the filters, object types and other criteria
	// which define the extract scope.
	JobDefinitionScope []ExtractJobDefinitionScope `mandatory:"false" json:"jobDefinitionScope"`
}

func (m UpdateJobDefinitionDetails) String() string {
	return common.PointerString(m)
}
