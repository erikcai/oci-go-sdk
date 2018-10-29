// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateZoneDetails The body for updating a zone.
// *Warning:* Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateZoneDetails struct {

	// External master servers for the zone. `externalMasters` becomes a
	// required parameter when the `zoneType` value is `SECONDARY`.
	ExternalMasters []ExternalMaster `mandatory:"false" json:"externalMasters"`
}

func (m UpdateZoneDetails) String() string {
	return common.PointerString(m)
}
