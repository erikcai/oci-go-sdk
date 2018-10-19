// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OraCache Public API
//
// Oracle Caching Service Public API
//

package cache

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ReplicatedCache Details of a redis replicated cache
type ReplicatedCache struct {

	// The OCID of the redis replicated cache
	Id *string `mandatory:"true" json:"id"`

	// The compartment OCID of the redis replicated cache
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The VCN OCID of the redis replicated cache
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The name of the redis replicated cache
	Name *string `mandatory:"true" json:"name"`

	// The number of replicas in the redis replicated cache
	ReplicaCount *int `mandatory:"true" json:"replicaCount"`

	// The shape of the redis replicated cache
	Shape *string `mandatory:"true" json:"shape"`

	// The endpoints of the replicas in the redis replicated cache
	RedisEndPoints []EndPoint `mandatory:"true" json:"redisEndPoints"`

	// Cache creation timestamp. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The lifecycleState of the redis replicated cache
	LifecycleState ReplicatedCacheLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A brief description of the redis replicated cache
	Description *string `mandatory:"false" json:"description"`
}

func (m ReplicatedCache) String() string {
	return common.PointerString(m)
}

// ReplicatedCacheLifecycleStateEnum Enum with underlying type: string
type ReplicatedCacheLifecycleStateEnum string

// Set of constants representing the allowable values for ReplicatedCacheLifecycleStateEnum
const (
	ReplicatedCacheLifecycleStateCreating ReplicatedCacheLifecycleStateEnum = "CREATING"
	ReplicatedCacheLifecycleStateActive   ReplicatedCacheLifecycleStateEnum = "ACTIVE"
	ReplicatedCacheLifecycleStateUpdating ReplicatedCacheLifecycleStateEnum = "UPDATING"
	ReplicatedCacheLifecycleStateDeleting ReplicatedCacheLifecycleStateEnum = "DELETING"
	ReplicatedCacheLifecycleStateDeleted  ReplicatedCacheLifecycleStateEnum = "DELETED"
	ReplicatedCacheLifecycleStateFailed   ReplicatedCacheLifecycleStateEnum = "FAILED"
)

var mappingReplicatedCacheLifecycleState = map[string]ReplicatedCacheLifecycleStateEnum{
	"CREATING": ReplicatedCacheLifecycleStateCreating,
	"ACTIVE":   ReplicatedCacheLifecycleStateActive,
	"UPDATING": ReplicatedCacheLifecycleStateUpdating,
	"DELETING": ReplicatedCacheLifecycleStateDeleting,
	"DELETED":  ReplicatedCacheLifecycleStateDeleted,
	"FAILED":   ReplicatedCacheLifecycleStateFailed,
}

// GetReplicatedCacheLifecycleStateEnumValues Enumerates the set of values for ReplicatedCacheLifecycleStateEnum
func GetReplicatedCacheLifecycleStateEnumValues() []ReplicatedCacheLifecycleStateEnum {
	values := make([]ReplicatedCacheLifecycleStateEnum, 0)
	for _, v := range mappingReplicatedCacheLifecycleState {
		values = append(values, v)
	}
	return values
}
