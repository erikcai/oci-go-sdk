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

// AppCatalogPackageVersion The model of an App Catalog Package Version
type AppCatalogPackageVersion struct {

	// Unique identifier for OPC package version
	PackageVersionId *string `mandatory:"false" json:"packageVersionId"`

	// Unique identifier for PIC resource version
	AppCatalogPackageVersion *string `mandatory:"false" json:"appCatalogPackageVersion"`

	// Array of Region codes
	Regions []string `mandatory:"false" json:"regions"`
}

func (m AppCatalogPackageVersion) String() string {
	return common.PointerString(m)
}
