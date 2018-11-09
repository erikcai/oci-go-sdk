// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousDatabaseMissionCriticalAssociation The properties that define the Mission Critical assocation between two different Autonomous Databases.
type AutonomousDatabaseMissionCriticalAssociation struct {

	// The OCID of the Autonomous Database Mission Critical Id.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database that has a relationship with the peer Autonomous Dedicated Database.
	AutonomousDatabaseId *string `mandatory:"true" json:"autonomousDatabaseId"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the peer Autonomous Database.
	PeerAutonomousDatabaseId *string `mandatory:"true" json:"peerAutonomousDatabaseId"`

	// The OCID of the peer Autonomous Database Mission Critical association id.
	PeerAutonomousDatabaseMissionCriticalAssociationId *string `mandatory:"true" json:"peerAutonomousDatabaseMissionCriticalAssociationId"`

	// The current state of the Autonomous Database Mission Critical association.
	LifecycleState AutonomousDatabaseMissionCriticalAssociationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The role of the reporting Autonomous Database in this Autonomous Mission Critical association.
	Role AutonomousDatabaseMissionCriticalAssociationRoleEnum `mandatory:"false" json:"role,omitempty"`

	// The role of the peer Autonomous Database in this Autonomous Mission Critical association.
	PeerRole AutonomousDatabaseMissionCriticalAssociationPeerRoleEnum `mandatory:"false" json:"peerRole,omitempty"`

	// Additional information about the current lifecycleState, if available.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m AutonomousDatabaseMissionCriticalAssociation) String() string {
	return common.PointerString(m)
}

// AutonomousDatabaseMissionCriticalAssociationRoleEnum Enum with underlying type: string
type AutonomousDatabaseMissionCriticalAssociationRoleEnum string

// Set of constants representing the allowable values for AutonomousDatabaseMissionCriticalAssociationRoleEnum
const (
	AutonomousDatabaseMissionCriticalAssociationRolePrimary AutonomousDatabaseMissionCriticalAssociationRoleEnum = "PRIMARY"
	AutonomousDatabaseMissionCriticalAssociationRoleStandby AutonomousDatabaseMissionCriticalAssociationRoleEnum = "STANDBY"
)

var mappingAutonomousDatabaseMissionCriticalAssociationRole = map[string]AutonomousDatabaseMissionCriticalAssociationRoleEnum{
	"PRIMARY": AutonomousDatabaseMissionCriticalAssociationRolePrimary,
	"STANDBY": AutonomousDatabaseMissionCriticalAssociationRoleStandby,
}

// GetAutonomousDatabaseMissionCriticalAssociationRoleEnumValues Enumerates the set of values for AutonomousDatabaseMissionCriticalAssociationRoleEnum
func GetAutonomousDatabaseMissionCriticalAssociationRoleEnumValues() []AutonomousDatabaseMissionCriticalAssociationRoleEnum {
	values := make([]AutonomousDatabaseMissionCriticalAssociationRoleEnum, 0)
	for _, v := range mappingAutonomousDatabaseMissionCriticalAssociationRole {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseMissionCriticalAssociationPeerRoleEnum Enum with underlying type: string
type AutonomousDatabaseMissionCriticalAssociationPeerRoleEnum string

// Set of constants representing the allowable values for AutonomousDatabaseMissionCriticalAssociationPeerRoleEnum
const (
	AutonomousDatabaseMissionCriticalAssociationPeerRolePrimary AutonomousDatabaseMissionCriticalAssociationPeerRoleEnum = "PRIMARY"
	AutonomousDatabaseMissionCriticalAssociationPeerRoleStandby AutonomousDatabaseMissionCriticalAssociationPeerRoleEnum = "STANDBY"
)

var mappingAutonomousDatabaseMissionCriticalAssociationPeerRole = map[string]AutonomousDatabaseMissionCriticalAssociationPeerRoleEnum{
	"PRIMARY": AutonomousDatabaseMissionCriticalAssociationPeerRolePrimary,
	"STANDBY": AutonomousDatabaseMissionCriticalAssociationPeerRoleStandby,
}

// GetAutonomousDatabaseMissionCriticalAssociationPeerRoleEnumValues Enumerates the set of values for AutonomousDatabaseMissionCriticalAssociationPeerRoleEnum
func GetAutonomousDatabaseMissionCriticalAssociationPeerRoleEnumValues() []AutonomousDatabaseMissionCriticalAssociationPeerRoleEnum {
	values := make([]AutonomousDatabaseMissionCriticalAssociationPeerRoleEnum, 0)
	for _, v := range mappingAutonomousDatabaseMissionCriticalAssociationPeerRole {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseMissionCriticalAssociationLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseMissionCriticalAssociationLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseMissionCriticalAssociationLifecycleStateEnum
const (
	AutonomousDatabaseMissionCriticalAssociationLifecycleStateProvisioning AutonomousDatabaseMissionCriticalAssociationLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseMissionCriticalAssociationLifecycleStateAvailable    AutonomousDatabaseMissionCriticalAssociationLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseMissionCriticalAssociationLifecycleStateUpdating     AutonomousDatabaseMissionCriticalAssociationLifecycleStateEnum = "UPDATING"
	AutonomousDatabaseMissionCriticalAssociationLifecycleStateTerminating  AutonomousDatabaseMissionCriticalAssociationLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseMissionCriticalAssociationLifecycleStateTerminated   AutonomousDatabaseMissionCriticalAssociationLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseMissionCriticalAssociationLifecycleStateFailed       AutonomousDatabaseMissionCriticalAssociationLifecycleStateEnum = "FAILED"
)

var mappingAutonomousDatabaseMissionCriticalAssociationLifecycleState = map[string]AutonomousDatabaseMissionCriticalAssociationLifecycleStateEnum{
	"PROVISIONING": AutonomousDatabaseMissionCriticalAssociationLifecycleStateProvisioning,
	"AVAILABLE":    AutonomousDatabaseMissionCriticalAssociationLifecycleStateAvailable,
	"UPDATING":     AutonomousDatabaseMissionCriticalAssociationLifecycleStateUpdating,
	"TERMINATING":  AutonomousDatabaseMissionCriticalAssociationLifecycleStateTerminating,
	"TERMINATED":   AutonomousDatabaseMissionCriticalAssociationLifecycleStateTerminated,
	"FAILED":       AutonomousDatabaseMissionCriticalAssociationLifecycleStateFailed,
}

// GetAutonomousDatabaseMissionCriticalAssociationLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseMissionCriticalAssociationLifecycleStateEnum
func GetAutonomousDatabaseMissionCriticalAssociationLifecycleStateEnumValues() []AutonomousDatabaseMissionCriticalAssociationLifecycleStateEnum {
	values := make([]AutonomousDatabaseMissionCriticalAssociationLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseMissionCriticalAssociationLifecycleState {
		values = append(values, v)
	}
	return values
}
