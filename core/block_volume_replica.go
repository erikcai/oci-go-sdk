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

// BlockVolumeReplica An asynchronous replica of a block volume that can then be used to create
// a new block volume or recover a block volume. For more information, see Overview
// of Block Volume Replicas (https://docs.cloud.oracle.com/Content/Block/Concepts/blockvolumereplicas.htm)
// To use any of the API operations, you must be authorized in an IAM policy.
// If you're not authorized, talk to an administrator. If you're an administrator
// who needs to write policies to give users access, see Getting Started with
// Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you
// supply string values using the API.
type BlockVolumeReplica struct {

	// The availability domain of the block volume replica.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the compartment that contains the block volume replica.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The block volume replica's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The current state of a block volume replica.
	LifecycleState BlockVolumeReplicaLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The size of the source block volume, in GBs.
	SizeInGBs *int64 `mandatory:"true" json:"sizeInGBs"`

	// The date and time the block volume replica was created. Format defined
	// by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the block volume replica was last synced from the source block volume.
	// Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeLastSynced *common.SDKTime `mandatory:"true" json:"timeLastSynced"`

	// The OCID of the source block volume.
	BlockVolumeId *string `mandatory:"true" json:"blockVolumeId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m BlockVolumeReplica) String() string {
	return common.PointerString(m)
}

// BlockVolumeReplicaLifecycleStateEnum Enum with underlying type: string
type BlockVolumeReplicaLifecycleStateEnum string

// Set of constants representing the allowable values for BlockVolumeReplicaLifecycleStateEnum
const (
	BlockVolumeReplicaLifecycleStateProvisioning BlockVolumeReplicaLifecycleStateEnum = "PROVISIONING"
	BlockVolumeReplicaLifecycleStateAvailable    BlockVolumeReplicaLifecycleStateEnum = "AVAILABLE"
	BlockVolumeReplicaLifecycleStateActivating   BlockVolumeReplicaLifecycleStateEnum = "ACTIVATING"
	BlockVolumeReplicaLifecycleStateTerminating  BlockVolumeReplicaLifecycleStateEnum = "TERMINATING"
	BlockVolumeReplicaLifecycleStateTerminated   BlockVolumeReplicaLifecycleStateEnum = "TERMINATED"
	BlockVolumeReplicaLifecycleStateFaulty       BlockVolumeReplicaLifecycleStateEnum = "FAULTY"
)

var mappingBlockVolumeReplicaLifecycleState = map[string]BlockVolumeReplicaLifecycleStateEnum{
	"PROVISIONING": BlockVolumeReplicaLifecycleStateProvisioning,
	"AVAILABLE":    BlockVolumeReplicaLifecycleStateAvailable,
	"ACTIVATING":   BlockVolumeReplicaLifecycleStateActivating,
	"TERMINATING":  BlockVolumeReplicaLifecycleStateTerminating,
	"TERMINATED":   BlockVolumeReplicaLifecycleStateTerminated,
	"FAULTY":       BlockVolumeReplicaLifecycleStateFaulty,
}

// GetBlockVolumeReplicaLifecycleStateEnumValues Enumerates the set of values for BlockVolumeReplicaLifecycleStateEnum
func GetBlockVolumeReplicaLifecycleStateEnumValues() []BlockVolumeReplicaLifecycleStateEnum {
	values := make([]BlockVolumeReplicaLifecycleStateEnum, 0)
	for _, v := range mappingBlockVolumeReplicaLifecycleState {
		values = append(values, v)
	}
	return values
}
