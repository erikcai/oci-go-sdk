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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v33/common"
)

// RemotePeeringConnectionDrgAttachmentNetworkDetails Specifies the DRG Attachment
type RemotePeeringConnectionDrgAttachmentNetworkDetails struct {

	// The OCID of the network attached to the DRG.
	Id *string `mandatory:"true" json:"id"`

	// The OCI region name
	PeerRegionName *string `mandatory:"false" json:"peerRegionName"`

	// The attachment route target.
	PeerAttachmentRouteTarget *string `mandatory:"false" json:"peerAttachmentRouteTarget"`

	// Routes which may be imported from the attachment (subject to import policy) appear in the route reflectors
	// tagged with the attachment's import route target.
	ImportRouteTarget *string `mandatory:"false" json:"importRouteTarget"`

	// Routes which are exported to the attachment are exported by El Paso as L3VPN routes to the route reflectors
	// with the route target set to the value of the attachment's export route target.
	ExportRouteTarget *string `mandatory:"false" json:"exportRouteTarget"`

	// The MPLS label of the DRG Attachment
	MplsLabel *int `mandatory:"false" json:"mplsLabel"`

	// The BGP ASN to use for the IPSec Connection's route target
	RegionalOciAsn *string `mandatory:"false" json:"regionalOciAsn"`
}

//GetId returns Id
func (m RemotePeeringConnectionDrgAttachmentNetworkDetails) GetId() *string {
	return m.Id
}

func (m RemotePeeringConnectionDrgAttachmentNetworkDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RemotePeeringConnectionDrgAttachmentNetworkDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRemotePeeringConnectionDrgAttachmentNetworkDetails RemotePeeringConnectionDrgAttachmentNetworkDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRemotePeeringConnectionDrgAttachmentNetworkDetails
	}{
		"REMOTE_PEERING_CONNECTION",
		(MarshalTypeRemotePeeringConnectionDrgAttachmentNetworkDetails)(m),
	}

	return json.Marshal(&s)
}
