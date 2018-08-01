// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// AppCatalogListingResourceVersion Listing Resource Version
type AppCatalogListingResourceVersion struct {

	// The OCID of the listing this resource version belongs to.
	ListingId *string `mandatory:"true" json:"listingId"`

	// Date and time the listing resource version was published, in RFC3339 format.
	// Example: `2018-03-20T12:32:53.532Z`
	PublishedDate *common.SDKTime `mandatory:"true" json:"publishedDate"`

	// OCID of the listing resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Resource Version.
	ResourceVersion *string `mandatory:"true" json:"resourceVersion"`

	// List of accessible ports for instances launched with this listing resource version.
	AccessiblePorts []int `mandatory:"true" json:"accessiblePorts"`

	// List of regions that this listing resource version is available. in
	AvailableRegions []string `mandatory:"false" json:"availableRegions"`

	// Array of shapes compatible with this resource.
	CompatibleShapes []string `mandatory:"false" json:"compatibleShapes"`

	// Allowed actions for the listing resource.
	AllowedActions []AppCatalogListingResourceVersionAllowedActionsEnum `mandatory:"true" json:"allowedActions"`
}

//GetListingId returns ListingId
func (m AppCatalogListingResourceVersion) GetListingId() *string {
	return m.ListingId
}

//GetPublishedDate returns PublishedDate
func (m AppCatalogListingResourceVersion) GetPublishedDate() *common.SDKTime {
	return m.PublishedDate
}

//GetResourceId returns ResourceId
func (m AppCatalogListingResourceVersion) GetResourceId() *string {
	return m.ResourceId
}

//GetResourceVersion returns ResourceVersion
func (m AppCatalogListingResourceVersion) GetResourceVersion() *string {
	return m.ResourceVersion
}

func (m AppCatalogListingResourceVersion) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AppCatalogListingResourceVersion) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAppCatalogListingResourceVersion AppCatalogListingResourceVersion
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAppCatalogListingResourceVersion
	}{
		"listingResourceVersion",
		(MarshalTypeAppCatalogListingResourceVersion)(m),
	}

	return json.Marshal(&s)
}

// AppCatalogListingResourceVersionAllowedActionsEnum Enum with underlying type: string
type AppCatalogListingResourceVersionAllowedActionsEnum string

// Set of constants representing the allowable values for AppCatalogListingResourceVersionAllowedActions
const (
	AppCatalogListingResourceVersionAllowedActionsSnapshot              AppCatalogListingResourceVersionAllowedActionsEnum = "SNAPSHOT"
	AppCatalogListingResourceVersionAllowedActionsBootVolumeDetach      AppCatalogListingResourceVersionAllowedActionsEnum = "BOOT_VOLUME_DETACH"
	AppCatalogListingResourceVersionAllowedActionsPreserveBootVolume    AppCatalogListingResourceVersionAllowedActionsEnum = "PRESERVE_BOOT_VOLUME"
	AppCatalogListingResourceVersionAllowedActionsSerialConsoleAccess   AppCatalogListingResourceVersionAllowedActionsEnum = "SERIAL_CONSOLE_ACCESS"
	AppCatalogListingResourceVersionAllowedActionsBootRecovery          AppCatalogListingResourceVersionAllowedActionsEnum = "BOOT_RECOVERY"
	AppCatalogListingResourceVersionAllowedActionsBackupBootVolume      AppCatalogListingResourceVersionAllowedActionsEnum = "BACKUP_BOOT_VOLUME"
	AppCatalogListingResourceVersionAllowedActionsCaptureConsoleHistory AppCatalogListingResourceVersionAllowedActionsEnum = "CAPTURE_CONSOLE_HISTORY"
)

var mappingAppCatalogListingResourceVersionAllowedActions = map[string]AppCatalogListingResourceVersionAllowedActionsEnum{
	"SNAPSHOT":                AppCatalogListingResourceVersionAllowedActionsSnapshot,
	"BOOT_VOLUME_DETACH":      AppCatalogListingResourceVersionAllowedActionsBootVolumeDetach,
	"PRESERVE_BOOT_VOLUME":    AppCatalogListingResourceVersionAllowedActionsPreserveBootVolume,
	"SERIAL_CONSOLE_ACCESS":   AppCatalogListingResourceVersionAllowedActionsSerialConsoleAccess,
	"BOOT_RECOVERY":           AppCatalogListingResourceVersionAllowedActionsBootRecovery,
	"BACKUP_BOOT_VOLUME":      AppCatalogListingResourceVersionAllowedActionsBackupBootVolume,
	"CAPTURE_CONSOLE_HISTORY": AppCatalogListingResourceVersionAllowedActionsCaptureConsoleHistory,
}

// GetAppCatalogListingResourceVersionAllowedActionsEnumValues Enumerates the set of values for AppCatalogListingResourceVersionAllowedActions
func GetAppCatalogListingResourceVersionAllowedActionsEnumValues() []AppCatalogListingResourceVersionAllowedActionsEnum {
	values := make([]AppCatalogListingResourceVersionAllowedActionsEnum, 0)
	for _, v := range mappingAppCatalogListingResourceVersionAllowedActions {
		values = append(values, v)
	}
	return values
}
