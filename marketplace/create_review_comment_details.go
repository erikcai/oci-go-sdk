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

// CreateReviewCommentDetails The details for creating a new review comment
type CreateReviewCommentDetails struct {

	// The body of the review comment
	CommentText *string `mandatory:"true" json:"commentText"`
}

func (m CreateReviewCommentDetails) String() string {
	return common.PointerString(m)
}
