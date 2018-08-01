// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AppCatalogListingResourceVersionAgreements Agreements for a listing resource version.
type AppCatalogListingResourceVersionAgreements struct {

	// Oracle TOU link
	OracleTermsOfUseLink *string `mandatory:"true" json:"oracleTermsOfUseLink"`

	// The OCID of the listing associated with these agreements.
	ListingId *string `mandatory:"true" json:"listingId"`

	// The OCID of the listing resource version associated with these agreements.
	ListingResourceVersion *string `mandatory:"true" json:"listingResourceVersion"`

	// Date and time the agreements were retrieved, in RFC3339 format.
	// Example: `2018-03-20T12:32:53.532Z`
	RetrieveDatetime *common.SDKTime `mandatory:"true" json:"retrieveDatetime"`

	// A signature for this agreement retrieval operation.
	Signature *string `mandatory:"true" json:"signature"`

	// EULA link
	EulaLink *string `mandatory:"false" json:"eulaLink"`
}

func (m AppCatalogListingResourceVersionAgreements) String() string {
	return common.PointerString(m)
}
