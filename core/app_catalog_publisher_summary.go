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

// AppCatalogPublisherSummary Publisher details.
type AppCatalogPublisherSummary struct {

	// Name of the publisher.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the publisher.
	Description *string `mandatory:"false" json:"description"`

	// Publisher's website URL.
	WebsiteUrl *string `mandatory:"false" json:"websiteUrl"`

	// Publisher's logo URL.
	LogoUrl *string `mandatory:"false" json:"logoUrl"`

	// Publisher's type.
	PublisherType AppCatalogPublisherSummaryPublisherTypeEnum `mandatory:"false" json:"publisherType,omitempty"`

	// Publisher's contact email.
	Email *string `mandatory:"false" json:"email"`

	// Publisher's contact phone number.
	// Example: "+1 555-555-1234"
	Phone *string `mandatory:"false" json:"phone"`

	// Publisher's address.
	// Example: 500 Oracle Parkway, Redwood Shores, CA, USA 94065
	Address *string `mandatory:"false" json:"address"`
}

func (m AppCatalogPublisherSummary) String() string {
	return common.PointerString(m)
}

// AppCatalogPublisherSummaryPublisherTypeEnum Enum with underlying type: string
type AppCatalogPublisherSummaryPublisherTypeEnum string

// Set of constants representing the allowable values for AppCatalogPublisherSummaryPublisherType
const (
	AppCatalogPublisherSummaryPublisherTypeOci      AppCatalogPublisherSummaryPublisherTypeEnum = "OCI"
	AppCatalogPublisherSummaryPublisherTypeOracle   AppCatalogPublisherSummaryPublisherTypeEnum = "ORACLE"
	AppCatalogPublisherSummaryPublisherTypeTrusted  AppCatalogPublisherSummaryPublisherTypeEnum = "TRUSTED"
	AppCatalogPublisherSummaryPublisherTypeStandard AppCatalogPublisherSummaryPublisherTypeEnum = "STANDARD"
)

var mappingAppCatalogPublisherSummaryPublisherType = map[string]AppCatalogPublisherSummaryPublisherTypeEnum{
	"OCI":      AppCatalogPublisherSummaryPublisherTypeOci,
	"ORACLE":   AppCatalogPublisherSummaryPublisherTypeOracle,
	"TRUSTED":  AppCatalogPublisherSummaryPublisherTypeTrusted,
	"STANDARD": AppCatalogPublisherSummaryPublisherTypeStandard,
}

// GetAppCatalogPublisherSummaryPublisherTypeEnumValues Enumerates the set of values for AppCatalogPublisherSummaryPublisherType
func GetAppCatalogPublisherSummaryPublisherTypeEnumValues() []AppCatalogPublisherSummaryPublisherTypeEnum {
	values := make([]AppCatalogPublisherSummaryPublisherTypeEnum, 0)
	for _, v := range mappingAppCatalogPublisherSummaryPublisherType {
		values = append(values, v)
	}
	return values
}
