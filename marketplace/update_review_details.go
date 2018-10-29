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

// UpdateReviewDetails Details for updating review
type UpdateReviewDetails struct {

	// The unique identifier of the review
	Id *string `mandatory:"false" json:"id"`

	// Rating for the application
	Rating *int `mandatory:"false" json:"rating"`

	// Unique Identifier of the application
	ApplicationId *string `mandatory:"false" json:"applicationId"`

	// Status of the update request
	StatusMessage *string `mandatory:"false" json:"statusMessage"`
}

func (m UpdateReviewDetails) String() string {
	return common.PointerString(m)
}
