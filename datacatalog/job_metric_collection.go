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

// JobMetricCollection Results of a Job Metrics Listing. Job Metrics are datum about a job execution in key value pairs.
type JobMetricCollection struct {

	// Collection of Job Metrics
	Items []JobMetricSummary `mandatory:"true" json:"items"`
}

func (m JobMetricCollection) String() string {
	return common.PointerString(m)
}
