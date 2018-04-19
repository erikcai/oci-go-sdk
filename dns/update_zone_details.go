// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DNS Service API
//
// API for managing DNS zones, records, and policies.
//

package dns

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateZoneDetails The body for updating a zone.
type UpdateZoneDetails struct {

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// External master servers for the zone.
	ExternalMasters []ExternalMaster `mandatory:"false" json:"externalMasters"`
}

func (m UpdateZoneDetails) String() string {
	return common.PointerString(m)
}
