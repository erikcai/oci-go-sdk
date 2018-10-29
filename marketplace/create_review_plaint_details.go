// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace REST API endpoint
//
// Marketplace REST API for loom plugin
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateReviewPlaintDetails The details of the review plait
type CreateReviewPlaintDetails struct {

	// The justification/reason of the plaint
	AccusationText *string `mandatory:"true" json:"accusationText"`
}

func (m CreateReviewPlaintDetails) String() string {
	return common.PointerString(m)
}
