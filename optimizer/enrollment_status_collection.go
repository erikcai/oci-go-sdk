// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Optimizer API
//
// The API for the OCI Optimizer
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// EnrollmentStatusCollection Results of a enrollment status search. Returns the EnrollmentStatusSummary items. Can be used to return additional data.
type EnrollmentStatusCollection struct {

	// The array of EnrollmentStatusSummary objects.
	Items []EnrollmentStatusSummary `mandatory:"true" json:"items"`
}

func (m EnrollmentStatusCollection) String() string {
	return common.PointerString(m)
}
