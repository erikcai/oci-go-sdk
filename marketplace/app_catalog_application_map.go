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

// AppCatalogApplicationMap This is a mapping for OPC identifiers to be mapped to Partner Image Catalog identifiers.
type AppCatalogApplicationMap struct {

	// App Catalog application id
	AppCatalogApplicationId *string `mandatory:"false" json:"appCatalogApplicationId"`

	// Marketplace Application Id
	MarketplaceApplicationId *string `mandatory:"false" json:"marketplaceApplicationId"`

	// List of Package versions
	PackageVersions []AppCatalogPackageVersion `mandatory:"false" json:"packageVersions"`

	// Reference link
	Links []Link `mandatory:"false" json:"links"`
}

func (m AppCatalogApplicationMap) String() string {
	return common.PointerString(m)
}
