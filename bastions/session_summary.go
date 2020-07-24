// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciBastions API
//
// A description of the OciBastions API
//

package bastions

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SessionSummary summary of session resource.
type SessionSummary struct {

	// OCID of session
	Id *string `mandatory:"true" json:"id"`

	// bastion display name that the session is created on
	BastionDisplayName *string `mandatory:"true" json:"bastionDisplayName"`

	// The OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of the session.
	SessionType SessionTypeEnum `mandatory:"true" json:"sessionType"`

	// display name of the target compute instance
	TargetComputeInstanceDisplayName *string `mandatory:"true" json:"targetComputeInstanceDisplayName"`

	// The OCID of the target compute instance to connect to.
	TargetComputeInstanceId *string `mandatory:"true" json:"targetComputeInstanceId"`

	// name of the user to use on target compute instance
	TargetComputeInstanceUserName *string `mandatory:"true" json:"targetComputeInstanceUserName"`

	// The time the session was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the session.
	LifecycleState SessionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// TTL of the session.
	SessionTtlInSeconds *int `mandatory:"true" json:"sessionTtlInSeconds"`

	// The time the session was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m SessionSummary) String() string {
	return common.PointerString(m)
}
