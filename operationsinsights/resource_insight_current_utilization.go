// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperationsInsights API
//
// The OperationsInsights API for OPSI service.
//

package operationsinsights

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ResourceInsightCurrentUtilization Current utilization(High/low) for cpu or storage
type ResourceInsightCurrentUtilization struct {

	// List of db ids with low usage
	Low []string `mandatory:"false" json:"low"`

	// List of db ids with high usage
	High []string `mandatory:"false" json:"high"`
}

func (m ResourceInsightCurrentUtilization) String() string {
	return common.PointerString(m)
}
