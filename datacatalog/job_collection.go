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

// JobCollection Results of a Jobs Listing. Jobs are scheduled instances of a job Definition.
type JobCollection struct {

	// Collection of Jobs
	Items []JobSummary `mandatory:"true" json:"items"`
}

func (m JobCollection) String() string {
	return common.PointerString(m)
}
