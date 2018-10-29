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

// PackageVersion The model of the package version
type PackageVersion struct {

	// The unique identifier of the package version
	Id *string `mandatory:"false" json:"id"`

	// The version of the package
	Version *string `mandatory:"false" json:"version"`

	// Status of the package version
	Status *Status `mandatory:"false" json:"status"`

	// Primary resource of the package
	PrimaryResource *Resource `mandatory:"false" json:"primaryResource"`

	// Publishing date of the package version
	PublishedDate *common.SDKTime `mandatory:"false" json:"publishedDate"`

	// Whether this version is default
	IsDefault *bool `mandatory:"false" json:"isDefault"`
}

func (m PackageVersion) String() string {
	return common.PointerString(m)
}
