// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// UsageApi API
//
// A description of the UsageApi API.
//

package usage

import (
	"github.com/oracle/oci-go-sdk/v29/common"
)

// CostRecord Object describing the cost and the usage for a specific resource type within the defined time interval.
type CostRecord struct {

	// The type of cost. Note that this property is not available for filtered queries (by compartmentId or cost tracking tag).
	ComputeType *string `mandatory:"false" json:"computeType"`

	// The amount of usage of the target resource (measured in {displayUnitName} units).
	ComputedQuantity *float32 `mandatory:"false" json:"computedQuantity"`

	// The cost of the target resource (measured in {currency} currency).
	ComputedAmount *float32 `mandatory:"false" json:"computedAmount"`

	// The unit price. Note that this property is not available for filtered queries (by compartmentId or cost tracking tag).
	UnitPrice *float32 `mandatory:"false" json:"unitPrice"`
}

func (m CostRecord) String() string {
	return common.PointerString(m)
}
