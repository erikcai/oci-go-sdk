// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v31/common"
)

// ClientVpnEndpointUser The user of a certain clientVpnEndpoint.
type ClientVpnEndpointUser struct {

	// The unique username of the user want to create.
	UserName *string `mandatory:"true" json:"userName"`

	// The current state of the ClientVPNendpointUser.
	LifecycleState ClientVpnEndpointUserLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Whether to log in the user by cert-authentication only or not.
	IsCertificateAuthenticationOnly *bool `mandatory:"false" json:"isCertificateAuthenticationOnly"`

	// The time ClientVpnEndpointUser was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m ClientVpnEndpointUser) String() string {
	return common.PointerString(m)
}

// ClientVpnEndpointUserLifecycleStateEnum Enum with underlying type: string
type ClientVpnEndpointUserLifecycleStateEnum string

// Set of constants representing the allowable values for ClientVpnEndpointUserLifecycleStateEnum
const (
	ClientVpnEndpointUserLifecycleStateCreating ClientVpnEndpointUserLifecycleStateEnum = "CREATING"
	ClientVpnEndpointUserLifecycleStateActive   ClientVpnEndpointUserLifecycleStateEnum = "ACTIVE"
	ClientVpnEndpointUserLifecycleStateInactive ClientVpnEndpointUserLifecycleStateEnum = "INACTIVE"
	ClientVpnEndpointUserLifecycleStateFailed   ClientVpnEndpointUserLifecycleStateEnum = "FAILED"
	ClientVpnEndpointUserLifecycleStateDeleted  ClientVpnEndpointUserLifecycleStateEnum = "DELETED"
	ClientVpnEndpointUserLifecycleStateDeleting ClientVpnEndpointUserLifecycleStateEnum = "DELETING"
	ClientVpnEndpointUserLifecycleStateUpdating ClientVpnEndpointUserLifecycleStateEnum = "UPDATING"
)

var mappingClientVpnEndpointUserLifecycleState = map[string]ClientVpnEndpointUserLifecycleStateEnum{
	"CREATING": ClientVpnEndpointUserLifecycleStateCreating,
	"ACTIVE":   ClientVpnEndpointUserLifecycleStateActive,
	"INACTIVE": ClientVpnEndpointUserLifecycleStateInactive,
	"FAILED":   ClientVpnEndpointUserLifecycleStateFailed,
	"DELETED":  ClientVpnEndpointUserLifecycleStateDeleted,
	"DELETING": ClientVpnEndpointUserLifecycleStateDeleting,
	"UPDATING": ClientVpnEndpointUserLifecycleStateUpdating,
}

// GetClientVpnEndpointUserLifecycleStateEnumValues Enumerates the set of values for ClientVpnEndpointUserLifecycleStateEnum
func GetClientVpnEndpointUserLifecycleStateEnumValues() []ClientVpnEndpointUserLifecycleStateEnum {
	values := make([]ClientVpnEndpointUserLifecycleStateEnum, 0)
	for _, v := range mappingClientVpnEndpointUserLifecycleState {
		values = append(values, v)
	}
	return values
}
