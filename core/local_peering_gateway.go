// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LocalPeeringGateway Details regarding a local peering gateway, which is an entity that allows two VCNs to communicate
// without traversing the Internet.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
type LocalPeeringGateway struct {

	// The OCID of the compartment containing the local peering gateway.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The local peering gateway's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// Indicates whether the peer local peering gateway is contained within another tenancy.
	IsCrossTenancyPeering *bool `mandatory:"true" json:"isCrossTenancyPeering"`

	// The local peering gateway's current lifecycle state.
	LifecycleState LocalPeeringGatewayLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Indicates whether the local peering gateway is peered with another local peering gateway.
	PeeringStatus LocalPeeringGatewayPeeringStatusEnum `mandatory:"true" json:"peeringStatus"`

	// The date and time the local peering gateway was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the VCN the local peering gateway belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Indicates the range of IPs available on the peer. `null` if not peered.
	PeerAdvertisedCidr *string `mandatory:"false" json:"peerAdvertisedCidr"`

	// Additional information regarding the peering status if applicable.
	PeeringStatusDetails *string `mandatory:"false" json:"peeringStatusDetails"`
}

func (m LocalPeeringGateway) String() string {
	return common.PointerString(m)
}

// LocalPeeringGatewayLifecycleStateEnum Enum with underlying type: string
type LocalPeeringGatewayLifecycleStateEnum string

// Set of constants representing the allowable values for LocalPeeringGatewayLifecycleState
const (
	LocalPeeringGatewayLifecycleStateProvisioning LocalPeeringGatewayLifecycleStateEnum = "PROVISIONING"
	LocalPeeringGatewayLifecycleStateAvailable    LocalPeeringGatewayLifecycleStateEnum = "AVAILABLE"
	LocalPeeringGatewayLifecycleStateTerminating  LocalPeeringGatewayLifecycleStateEnum = "TERMINATING"
	LocalPeeringGatewayLifecycleStateTerminated   LocalPeeringGatewayLifecycleStateEnum = "TERMINATED"
)

var mappingLocalPeeringGatewayLifecycleState = map[string]LocalPeeringGatewayLifecycleStateEnum{
	"PROVISIONING": LocalPeeringGatewayLifecycleStateProvisioning,
	"AVAILABLE":    LocalPeeringGatewayLifecycleStateAvailable,
	"TERMINATING":  LocalPeeringGatewayLifecycleStateTerminating,
	"TERMINATED":   LocalPeeringGatewayLifecycleStateTerminated,
}

// GetLocalPeeringGatewayLifecycleStateEnumValues Enumerates the set of values for LocalPeeringGatewayLifecycleState
func GetLocalPeeringGatewayLifecycleStateEnumValues() []LocalPeeringGatewayLifecycleStateEnum {
	values := make([]LocalPeeringGatewayLifecycleStateEnum, 0)
	for _, v := range mappingLocalPeeringGatewayLifecycleState {
		values = append(values, v)
	}
	return values
}

// LocalPeeringGatewayPeeringStatusEnum Enum with underlying type: string
type LocalPeeringGatewayPeeringStatusEnum string

// Set of constants representing the allowable values for LocalPeeringGatewayPeeringStatus
const (
	LocalPeeringGatewayPeeringStatusInvalid LocalPeeringGatewayPeeringStatusEnum = "INVALID"
	LocalPeeringGatewayPeeringStatusNew     LocalPeeringGatewayPeeringStatusEnum = "NEW"
	LocalPeeringGatewayPeeringStatusPeered  LocalPeeringGatewayPeeringStatusEnum = "PEERED"
	LocalPeeringGatewayPeeringStatusPending LocalPeeringGatewayPeeringStatusEnum = "PENDING"
	LocalPeeringGatewayPeeringStatusRevoked LocalPeeringGatewayPeeringStatusEnum = "REVOKED"
)

var mappingLocalPeeringGatewayPeeringStatus = map[string]LocalPeeringGatewayPeeringStatusEnum{
	"INVALID": LocalPeeringGatewayPeeringStatusInvalid,
	"NEW":     LocalPeeringGatewayPeeringStatusNew,
	"PEERED":  LocalPeeringGatewayPeeringStatusPeered,
	"PENDING": LocalPeeringGatewayPeeringStatusPending,
	"REVOKED": LocalPeeringGatewayPeeringStatusRevoked,
}

// GetLocalPeeringGatewayPeeringStatusEnumValues Enumerates the set of values for LocalPeeringGatewayPeeringStatus
func GetLocalPeeringGatewayPeeringStatusEnumValues() []LocalPeeringGatewayPeeringStatusEnum {
	values := make([]LocalPeeringGatewayPeeringStatusEnum, 0)
	for _, v := range mappingLocalPeeringGatewayPeeringStatus {
		values = append(values, v)
	}
	return values
}
