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

// CreateReplicatedCacheDetails The details to create the replicated cache with.
type CreateReplicatedCacheDetails struct {

	// The OCID of the customer's compartment that the redis replicated cache is being created in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the redis replicated cache.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the customer's VCN that the redis replicated cache is being created in.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The number of replicas.
	ReplicaCount *int `mandatory:"true" json:"replicaCount"`

	// The shape of the redis replicated cache.
	Shape *string `mandatory:"true" json:"shape"`

	// The description of the redis replicated cache.
	Description *string `mandatory:"false" json:"description"`

	// The ad/subnet within which each redis node (primary/replica) should be created in.
	RedisNodeDetailsList []RedisNodeDetails `mandatory:"false" json:"redisNodeDetailsList"`
}

func (m CreateReplicatedCacheDetails) String() string {
	return common.PointerString(m)
}
