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

// JobExecutionCollection Results of a Job Executions Listing. Job Executions are execution instances of a scheduled job.
type JobExecutionCollection struct {

	// Collection of Job Executions
	Items []JobExecutionSummary `mandatory:"true" json:"items"`
}

func (m JobExecutionCollection) String() string {
	return common.PointerString(m)
}
