// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OraCache Public API
//
// API for the Data Caching Service. Use this service to manage Redis replicated caches.
//

package cache

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ReplicatedCache Details of the Redis replicated cache. Redis replicated caches are comprised of Oracle-managed Redis nodes that each contain a replica of the cached data. The cache is accessible from a tenant's compartment using a published endpoint.
type ReplicatedCache struct {

	// The OCID of the Redis replicated cache.
	Id *string `mandatory:"true" json:"id"`

	// The compartment OCID from which the Redis replicated cache is accessible.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The VCN OCID that contains the network resources and subnets to which the Redis nodes are attached.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The name of the Redis replicated cache
	Name *string `mandatory:"true" json:"name"`

	// The number of replica nodes that make up the Redis replicated cache.
	ReplicaCount *int `mandatory:"true" json:"replicaCount"`

	// The amount of memory allocated to the Redis replicated cache.
	Shape *string `mandatory:"true" json:"shape"`

	// The endpoints of the replicas that make up the Redis replicated cache.
	RedisNodes []EndPoint `mandatory:"true" json:"redisNodes"`

	// Cache creation timestamp. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The `lifecycleState` of the Redis replicated cache.
	LifecycleState ReplicatedCacheLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A brief description of the Redis replicated cache
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
