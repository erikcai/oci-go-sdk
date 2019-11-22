// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object Storage and Archive Storage APIs for managing buckets, objects, and related resources.
//

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateReplicationPolicyDetails The details to create a replication policy.
type CreateReplicationPolicyDetails struct {

	// The name of the policy.
	Name *string `mandatory:"true" json:"name"`

	// The destination region to replicate to, for example "us-ashburn-1".
	DestinationRegionName *string `mandatory:"true" json:"destinationRegionName"`

	// The bucket to replicate to in the destination region.
	DestinationBucketName *string `mandatory:"true" json:"destinationBucketName"`
}

func (m CreateReplicationPolicyDetails) String() string {
	return common.PointerString(m)
}
