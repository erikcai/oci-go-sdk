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

// CreateJobDefinitionDetails Representation of a Job Definition Resource. Job Definitions define the harvest scope and includes the list of
// objects to be harvested along with a schedule. The list of objects is usually specified through a combination of
// object type, regular expressions or specific names of objects and a sample size for the data harvested.
type CreateJobDefinitionDetails struct {

	// The display name of a user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of the Job Definition.
	JobType JobTypeEnum `mandatory:"true" json:"jobType"`

	// The key of the Data Asset for which the job is defined.
	DataAssetKey *string `mandatory:"true" json:"dataAssetKey"`

	// The key of the connection resource to be used for the job.
	ConnectionKey *string `mandatory:"true" json:"connectionKey"`

	// Detailed description of the Job Definition.
	Description *string `mandatory:"false" json:"description"`

	// Specifies if the Job Definition is incremental or full.
	IsIncremental *bool `mandatory:"false" json:"isIncremental"`

	// Scope for the job definition.
	JobDefinitionScope []ExtractJobDefinitionScope `mandatory:"false" json:"jobDefinitionScope"`
}

func (m CreateJobDefinitionDetails) String() string {
	return common.PointerString(m)
}
