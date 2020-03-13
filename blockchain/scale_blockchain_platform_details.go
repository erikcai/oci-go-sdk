// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Plane API
//

package blockchain

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ScaleBlockchainPlatformDetails Scale operation details for a blockchain platform. The scale operation payload has multiple options
// - Add one or more Ordering Service Node (addOsns)
// - Add one or more Peers (addPeers)
// - Add more replicas of CA, Console and Rest Proxy (addReplicas)
// - Add more storage to the platform (addStorage)
// - Modify the CPU allocation for Peer Nodes (modifyPeers)
// - Remove one or more replicas of CA, Console and Rest Proxy (removeReplicas)
// - Remove one or more Ordering Service Node (removeOsns)
// - Remove one or more Peers (removePeers).
// The scale operation payload must have at least one of the above options.
type ScaleBlockchainPlatformDetails struct {

	// new OSNs to add
	AddOsns []ScaleOutOsnDetails `mandatory:"false" json:"addOsns"`

	AddReplicas *ScaleReplicaInfoDetails `mandatory:"false" json:"addReplicas"`

	// new Peers to add
	AddPeers []ScaleOutPeerDetails `mandatory:"false" json:"addPeers"`

	AddStorage *ScaleStorageDetails `mandatory:"false" json:"addStorage"`

	// modify ocpu allocation to existing Peers
	ModifyPeers []ModifyPeerDetails `mandatory:"false" json:"modifyPeers"`

	RemoveReplicas *ScaleReplicaInfoDetails `mandatory:"false" json:"removeReplicas"`

	// OSN id list to remove
	RemoveOsns []string `mandatory:"false" json:"removeOsns"`

	// Peer id list to remove
	RemovePeers []string `mandatory:"false" json:"removePeers"`
}

func (m ScaleBlockchainPlatformDetails) String() string {
	return common.PointerString(m)
}
