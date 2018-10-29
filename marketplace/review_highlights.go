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

// ReviewHighlights The model of review highlights
type ReviewHighlights struct {

	// The average rating of the application
	AverageRating *float64 `mandatory:"false" json:"averageRating"`

	// The maximum rating the platform allows
	MaxRating *int64 `mandatory:"false" json:"maxRating"`

	// The count of reviews for the application
	ReviewCount *int64 `mandatory:"false" json:"reviewCount"`

	// The sum of one star reviews
	OneStarCount *int64 `mandatory:"false" json:"oneStarCount"`

	// The sum of two star reviews
	TwoStarCount *int64 `mandatory:"false" json:"twoStarCount"`

	// The sum of three star reviews
	ThreeStarCount *int64 `mandatory:"false" json:"threeStarCount"`

	// The sum of four star reviews
	FourStarCount *int64 `mandatory:"false" json:"fourStarCount"`

	// The sum of five star reviews
	FiveStarCount *int64 `mandatory:"false" json:"fiveStarCount"`
}

func (m ReviewHighlights) String() string {
	return common.PointerString(m)
}
