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

// ReplicationPolicySummary The summary of a cross region replication policy.
type ReplicationPolicySummary struct {

	// The name of the policy.
	Name *string `mandatory:"true" json:"name"`

	// The destination region to replicate to, for example "us-ashburn-1".
	DestinationRegionName *string `mandatory:"true" json:"destinationRegionName"`

	// The bucket to replicate to in the destination region.
	DestinationBucketName *string `mandatory:"true" json:"destinationBucketName"`

	// The id of the replication policy.
	Id *string `mandatory:"false" json:"id"`

	// The date when the replication policy was created as per RFC 3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Changes made to the source bucket before this time has been replicated.
	TimeLastSync *common.SDKTime `mandatory:"false" json:"timeLastSync"`

	// The replication status of the policy. If the status is CLIENT_ERROR, once the user fixes the issue
	// described in the status message, the status will become ACTIVE.
	Status ReplicationPolicySummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// A human-readable description of the status.
	StatusMessage *string `mandatory:"false" json:"statusMessage"`
}

func (m ReplicationPolicySummary) String() string {
	return common.PointerString(m)
}

// ReplicationPolicySummaryStatusEnum Enum with underlying type: string
type ReplicationPolicySummaryStatusEnum string

// Set of constants representing the allowable values for ReplicationPolicySummaryStatusEnum
const (
	ReplicationPolicySummaryStatusActive      ReplicationPolicySummaryStatusEnum = "ACTIVE"
	ReplicationPolicySummaryStatusClientError ReplicationPolicySummaryStatusEnum = "CLIENT_ERROR"
)

var mappingReplicationPolicySummaryStatus = map[string]ReplicationPolicySummaryStatusEnum{
	"ACTIVE":       ReplicationPolicySummaryStatusActive,
	"CLIENT_ERROR": ReplicationPolicySummaryStatusClientError,
}

// GetReplicationPolicySummaryStatusEnumValues Enumerates the set of values for ReplicationPolicySummaryStatusEnum
func GetReplicationPolicySummaryStatusEnumValues() []ReplicationPolicySummaryStatusEnum {
	values := make([]ReplicationPolicySummaryStatusEnum, 0)
	for _, v := range mappingReplicationPolicySummaryStatus {
		values = append(values, v)
	}
	return values
}
