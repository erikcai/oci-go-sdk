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

// ApplicationSummary The summary of application details
type ApplicationSummary struct {

	// The unique identifier of the application
	Id *string `mandatory:"false" json:"id"`

	// The name of the application
	Name *string `mandatory:"false" json:"name"`

	// The tagline of the application
	TagLine *string `mandatory:"false" json:"tagLine"`

	// The tags of the application
	Tags *string `mandatory:"false" json:"tags"`

	// The rating of the application
	Rating *int `mandatory:"false" json:"rating"`

	// The count of reviews of the application
	ReviewCount *int64 `mandatory:"false" json:"reviewCount"`

	// The URL of the icon
	Icon *UploadData `mandatory:"false" json:"icon"`

	// The pricing of the application
	Pricing *Pricing `mandatory:"false" json:"pricing"`

	// Whether the application is installable
	IsInstallable *bool `mandatory:"false" json:"isInstallable"`

	// Whether the application is ready for cloud deployment
	IsCloudReady *bool `mandatory:"false" json:"isCloudReady"`
}

func (m ApplicationSummary) String() string {
	return common.PointerString(m)
}
