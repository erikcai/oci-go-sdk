// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Listing The model for a listing in Oracle Cloud Infrastructure Marketplace.
type Listing struct {

	// The unique identifier for the listing in the Oracle Cloud Infrastructure Marketplace.
	Id *string `mandatory:"false" json:"id"`

	// The name of the listing.
	Name *string `mandatory:"false" json:"name"`

	// The version of the listing.
	Version *string `mandatory:"false" json:"version"`

	// The tagline of the listing.
	Tagline *string `mandatory:"false" json:"tagline"`

	// Keywords associated with the listing.
	Keywords *string `mandatory:"false" json:"keywords"`

	// A short description of the listing.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// Usage Information about Listing
	UsageInformation *string `mandatory:"false" json:"usageInformation"`

	// A long description of the listing.
	LongDescription *string `mandatory:"false" json:"longDescription"`

	// A description of the publisher's licensing model for the listing.
	LicenseModelDescription *string `mandatory:"false" json:"licenseModelDescription"`

	// System requirements for the listing.
	SystemRequirements *string `mandatory:"false" json:"systemRequirements"`

	// The release date of the listing.
	TimeReleased *common.SDKTime `mandatory:"false" json:"timeReleased"`

	// The release notes for the listing.
	ReleaseNotes *string `mandatory:"false" json:"releaseNotes"`

	// Categories that this listing belongs to.
	Categories []string `mandatory:"false" json:"categories"`

	// The publisher of the listing.
	Publisher *Publisher `mandatory:"false" json:"publisher"`

	// The languages supported by the application.
	Languages []Item `mandatory:"false" json:"languages"`

	// Screenshots of the listing.
	Screenshots []Screenshot `mandatory:"false" json:"screenshots"`

	// Videos of the listing.
	Videos []NamedLink `mandatory:"false" json:"videos"`

	// Contact information to use to get support for the listing.
	SupportContacts []SupportContact `mandatory:"false" json:"supportContacts"`

	// Links to support resources for the listing.
	SupportLinks []NamedLink `mandatory:"false" json:"supportLinks"`

	// Links to additional documentation resources for the listing.
	DocumentationLinks []DocumentationLink `mandatory:"false" json:"documentationLinks"`

	// The URL of the listing icon.
	Icon *UploadData `mandatory:"false" json:"icon"`

	// The URL of the banner.
	Banner *UploadData `mandatory:"false" json:"banner"`

	// The regions where the listing is available.
	Regions []Region `mandatory:"false" json:"regions"`

	// The package type of the listing.
	PackageType PackageTypeEnumEnum `mandatory:"false" json:"packageType,omitempty"`

	// The default package version.
	DefaultPackageVersion *string `mandatory:"false" json:"defaultPackageVersion"`

	// Reference links.
	Links []Link `mandatory:"false" json:"links"`
}

func (m Listing) String() string {
	return common.PointerString(m)
}

// ListingPackageTypeEnum is an alias to type: PackageTypeEnumEnum
// Consider using PackageTypeEnumEnum instead
// Deprecated
type ListingPackageTypeEnum = PackageTypeEnumEnum

// Set of constants representing the allowable values for PackageTypeEnumEnum
// Deprecated
const (
	ListingPackageTypeImage PackageTypeEnumEnum = "IMAGE"
)

// GetListingPackageTypeEnumValues Enumerates the set of values for PackageTypeEnumEnum
// Consider using GetPackageTypeEnumEnumValue
// Deprecated
var GetListingPackageTypeEnumValues = GetPackageTypeEnumEnumValues
