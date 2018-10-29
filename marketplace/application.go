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

// Application The model of an application of Oracle Marketplace
type Application struct {

	// Unique identifier of of the application
	Id *string `mandatory:"false" json:"id"`

	// Name of the application
	Name *string `mandatory:"false" json:"name"`

	// Tagline of the application
	TagLine *string `mandatory:"false" json:"tagLine"`

	// Tags of the application
	Tags *string `mandatory:"false" json:"tags"`

	// Device type of the application
	DeviceType *DeviceType `mandatory:"false" json:"deviceType"`

	// Short description of the application
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// Long description of the application
	LongDescription *string `mandatory:"false" json:"longDescription"`

	// Publisher of the application
	Publisher *Publisher `mandatory:"false" json:"publisher"`

	// Related documents of the application
	RelatedDocuments []ApplicationDoc `mandatory:"false" json:"relatedDocuments"`

	// List of languages that the application supports
	Languages []Item `mandatory:"false" json:"languages"`

	// Review highlights of the application
	ReviewHighlights *ReviewHighlights `mandatory:"false" json:"reviewHighlights"`

	// Screenshot of the application
	ScreenShots []ScreenShot `mandatory:"false" json:"screenShots"`

	// Application Videos of the application
	ApplicationVideos []ApplicationVideo `mandatory:"false" json:"applicationVideos"`

	// Products of the application
	Products []Products `mandatory:"false" json:"products"`

	// The demo URL
	DemoURL *string `mandatory:"false" json:"demoURL"`

	// Marketplace Self paced training URL
	SelfPacedTrainingURL *string `mandatory:"false" json:"selfPacedTrainingURL"`

	// The URL of the icon
	Icon *UploadData `mandatory:"false" json:"icon"`

	// The URL of the banner
	Banner *UploadData `mandatory:"false" json:"banner"`

	// The pricing of the banner
	Pricing *Pricing `mandatory:"false" json:"pricing"`

	// The version details of the application
	VersionDetails *VersionDetails `mandatory:"false" json:"versionDetails"`

	// The version history of the application
	VersionHistory []VersionDetails `mandatory:"false" json:"versionHistory"`

	// The support information of the application
	Support *SupportInfo `mandatory:"false" json:"support"`

	// Defines whether the application is installable
	IsInstallable *bool `mandatory:"false" json:"isInstallable"`

	// Defines whether the application is ready for deployment in cloud
	IsCloudReady *bool `mandatory:"false" json:"isCloudReady"`

	// System requirements of the application
	SystemRequirements *string `mandatory:"false" json:"systemRequirements"`

	// Defines whether the application is available in all regions
	IsAvailableInAllRegions *bool `mandatory:"false" json:"isAvailableInAllRegions"`

	// The list of regions where the application is available
	Regions []Region `mandatory:"false" json:"regions"`

	// The package type of the application
	PackageType *string `mandatory:"false" json:"packageType"`

	// The default package version
	DefaultPackageVersion *string `mandatory:"false" json:"defaultPackageVersion"`

	// The list of available package versions
	PackageVersions []PackageVersion `mandatory:"false" json:"packageVersions"`

	// Reference link
	Links []Link `mandatory:"false" json:"links"`
}

func (m Application) String() string {
	return common.PointerString(m)
}
