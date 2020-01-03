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

// JobDefinitionCollection Results of a Job Definition Listing. Job Definitions are resources that describe the scope and type of jobs (eg: harvest , profiling , sampling) that are defined by users in the system.
type JobDefinitionCollection struct {

	// Collection of JobDefinitions
	Items []JobDefinitionSummary `mandatory:"true" json:"items"`
}

func (m JobDefinitionCollection) String() string {
	return common.PointerString(m)
}
