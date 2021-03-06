// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm),
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm), and
// Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as virtual cloud networks (VCNs),
// compute instances, block storage volumes, and container images.
//

package core

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// MigrateDrgDetails The request to update a DrgAttachment and migrate the Drg to the destination Drg Type.
type MigrateDrgDetails struct {

	// Type of the Drg to be migrated to.
	DestinationDrgType MigrateDrgDetailsDestinationDrgTypeEnum `mandatory:"true" json:"destinationDrgType"`

	// The DRG attachment's Oracle ID (OCID).
	DrgAttachmentId *string `mandatory:"false" json:"drgAttachmentId"`

	// NextHop target's MPLS label.
	MplsLabel *string `mandatory:"false" json:"mplsLabel"`

	// The string in the form ASN:mplsLabel.
	RouteTarget *string `mandatory:"false" json:"routeTarget"`
}

func (m MigrateDrgDetails) String() string {
	return common.PointerString(m)
}

// MigrateDrgDetailsDestinationDrgTypeEnum Enum with underlying type: string
type MigrateDrgDetailsDestinationDrgTypeEnum string

// Set of constants representing the allowable values for MigrateDrgDetailsDestinationDrgTypeEnum
const (
	MigrateDrgDetailsDestinationDrgTypeClassical  MigrateDrgDetailsDestinationDrgTypeEnum = "DRG_CLASSICAL"
	MigrateDrgDetailsDestinationDrgTypeTransitHub MigrateDrgDetailsDestinationDrgTypeEnum = "DRG_TRANSIT_HUB"
)

var mappingMigrateDrgDetailsDestinationDrgType = map[string]MigrateDrgDetailsDestinationDrgTypeEnum{
	"DRG_CLASSICAL":   MigrateDrgDetailsDestinationDrgTypeClassical,
	"DRG_TRANSIT_HUB": MigrateDrgDetailsDestinationDrgTypeTransitHub,
}

// GetMigrateDrgDetailsDestinationDrgTypeEnumValues Enumerates the set of values for MigrateDrgDetailsDestinationDrgTypeEnum
func GetMigrateDrgDetailsDestinationDrgTypeEnumValues() []MigrateDrgDetailsDestinationDrgTypeEnum {
	values := make([]MigrateDrgDetailsDestinationDrgTypeEnum, 0)
	for _, v := range mappingMigrateDrgDetailsDestinationDrgType {
		values = append(values, v)
	}
	return values
}
