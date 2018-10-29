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

// CreateReviewDetails The details for creating a new review
type CreateReviewDetails struct {

	// The rating of the application
	Rating *int64 `mandatory:"true" json:"rating"`

	// The subject of the review
	ReviewSubject *string `mandatory:"true" json:"reviewSubject"`

	// The body of the review
	ReviewText *string `mandatory:"true" json:"reviewText"`

	// The unique identifier of the reviewer
	ReviewerId *string `mandatory:"true" json:"reviewerId"`

	// The unique identifier of the application
	ApplicationId *string `mandatory:"true" json:"applicationId"`
}

func (m CreateReviewDetails) String() string {
	return common.PointerString(m)
}
