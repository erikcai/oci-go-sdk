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

// ReplicatedCacheSummary Summary of a redis replicated cache
type ReplicatedCacheSummary struct {

	// The OCID of the redis replicated cache
	Id *string `mandatory:"true" json:"id"`

	// The name of the redis replicated cache
	Name *string `mandatory:"true" json:"name"`

	// The number of replicas in the redis replicated cache
	ReplicaCount *int `mandatory:"true" json:"replicaCount"`

	// The lifecycleState of the redis replicated cache
	LifecycleState ReplicatedCacheSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The URI to access the detailed information of the redis replicated cache
	ResourceUri *string `mandatory:"true" json:"resourceUri"`

	// The shape of the redis replicated cache
	Shape *string `mandatory:"true" json:"shape"`

	// A brief description of the redis replicated cache
	Description *string `mandatory:"false" json:"description"`
}

func (m ReplicatedCacheSummary) String() string {
	return common.PointerString(m)
}

// ReplicatedCacheSummaryLifecycleStateEnum Enum with underlying type: string
type ReplicatedCacheSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ReplicatedCacheSummaryLifecycleStateEnum
const (
	ReplicatedCacheSummaryLifecycleStateCreating ReplicatedCacheSummaryLifecycleStateEnum = "CREATING"
	ReplicatedCacheSummaryLifecycleStateActive   ReplicatedCacheSummaryLifecycleStateEnum = "ACTIVE"
	ReplicatedCacheSummaryLifecycleStateUpdating ReplicatedCacheSummaryLifecycleStateEnum = "UPDATING"
	ReplicatedCacheSummaryLifecycleStateDeleting ReplicatedCacheSummaryLifecycleStateEnum = "DELETING"
	ReplicatedCacheSummaryLifecycleStateDeleted  ReplicatedCacheSummaryLifecycleStateEnum = "DELETED"
	ReplicatedCacheSummaryLifecycleStateFailed   ReplicatedCacheSummaryLifecycleStateEnum = "FAILED"
)

var mappingReplicatedCacheSummaryLifecycleState = map[string]ReplicatedCacheSummaryLifecycleStateEnum{
	"CREATING": ReplicatedCacheSummaryLifecycleStateCreating,
	"ACTIVE":   ReplicatedCacheSummaryLifecycleStateActive,
	"UPDATING": ReplicatedCacheSummaryLifecycleStateUpdating,
	"DELETING": ReplicatedCacheSummaryLifecycleStateDeleting,
	"DELETED":  ReplicatedCacheSummaryLifecycleStateDeleted,
	"FAILED":   ReplicatedCacheSummaryLifecycleStateFailed,
}

// GetReplicatedCacheSummaryLifecycleStateEnumValues Enumerates the set of values for ReplicatedCacheSummaryLifecycleStateEnum
func GetReplicatedCacheSummaryLifecycleStateEnumValues() []ReplicatedCacheSummaryLifecycleStateEnum {
	values := make([]ReplicatedCacheSummaryLifecycleStateEnum, 0)
	for _, v := range mappingReplicatedCacheSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
