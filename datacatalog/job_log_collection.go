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

// JobLogCollection Results of a Job Logs Listing. Job Log is an audit log record inserted during the lifecycle of a job execution instance.
type JobLogCollection struct {

	// Collection of Job Logs
	Items []JobLogSummary `mandatory:"true" json:"items"`
}

func (m JobLogCollection) String() string {
	return common.PointerString(m)
}
