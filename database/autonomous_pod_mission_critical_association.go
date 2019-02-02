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

// AutonomousPodMissionCriticalAssociation The properties that define the Mission Critical Association between two different Autonomous Pods.
type AutonomousPodMissionCriticalAssociation struct {

	// The OCID of the Mission Critical Autonomous Pod id.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Pod that has a relationship with the peer Autonomous Pod.
	AutonomousPodId *string `mandatory:"true" json:"autonomousPodId"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the peer Autonomous Pod.
	PeerAutonomousPodId *string `mandatory:"true" json:"peerAutonomousPodId"`

	// The role of the reporting Pod in this Autonomous Mission Critical Association.
	Role AutonomousPodMissionCriticalAssociationRoleEnum `mandatory:"true" json:"role"`

	// The role of the peer Pod in this Autonomous Mission Critical Association.
	PeerRole AutonomousPodMissionCriticalAssociationPeerRoleEnum `mandatory:"true" json:"peerRole"`

	// The OCID of the peer Autonomous Pod Mission Critical Association.
	PeerAutonomousPodMissionCriticalAssociationId *string `mandatory:"true" json:"peerAutonomousPodMissionCriticalAssociationId"`

	// The current state of the Autonomous Pod Mission Critical Association.
	LifecycleState AutonomousPodMissionCriticalAssociationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Additional information about the current lifecycleState, if available.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m AutonomousPodMissionCriticalAssociation) String() string {
	return common.PointerString(m)
}

// AutonomousPodMissionCriticalAssociationRoleEnum Enum with underlying type: string
type AutonomousPodMissionCriticalAssociationRoleEnum string

// Set of constants representing the allowable values for AutonomousPodMissionCriticalAssociationRoleEnum
const (
	AutonomousPodMissionCriticalAssociationRolePrimary AutonomousPodMissionCriticalAssociationRoleEnum = "PRIMARY"
	AutonomousPodMissionCriticalAssociationRoleStandby AutonomousPodMissionCriticalAssociationRoleEnum = "STANDBY"
)

var mappingAutonomousPodMissionCriticalAssociationRole = map[string]AutonomousPodMissionCriticalAssociationRoleEnum{
	"PRIMARY": AutonomousPodMissionCriticalAssociationRolePrimary,
	"STANDBY": AutonomousPodMissionCriticalAssociationRoleStandby,
}

// GetAutonomousPodMissionCriticalAssociationRoleEnumValues Enumerates the set of values for AutonomousPodMissionCriticalAssociationRoleEnum
func GetAutonomousPodMissionCriticalAssociationRoleEnumValues() []AutonomousPodMissionCriticalAssociationRoleEnum {
	values := make([]AutonomousPodMissionCriticalAssociationRoleEnum, 0)
	for _, v := range mappingAutonomousPodMissionCriticalAssociationRole {
		values = append(values, v)
	}
	return values
}

// AutonomousPodMissionCriticalAssociationPeerRoleEnum Enum with underlying type: string
type AutonomousPodMissionCriticalAssociationPeerRoleEnum string

// Set of constants representing the allowable values for AutonomousPodMissionCriticalAssociationPeerRoleEnum
const (
	AutonomousPodMissionCriticalAssociationPeerRolePrimary AutonomousPodMissionCriticalAssociationPeerRoleEnum = "PRIMARY"
	AutonomousPodMissionCriticalAssociationPeerRoleStandby AutonomousPodMissionCriticalAssociationPeerRoleEnum = "STANDBY"
)

var mappingAutonomousPodMissionCriticalAssociationPeerRole = map[string]AutonomousPodMissionCriticalAssociationPeerRoleEnum{
	"PRIMARY": AutonomousPodMissionCriticalAssociationPeerRolePrimary,
	"STANDBY": AutonomousPodMissionCriticalAssociationPeerRoleStandby,
}

// GetAutonomousPodMissionCriticalAssociationPeerRoleEnumValues Enumerates the set of values for AutonomousPodMissionCriticalAssociationPeerRoleEnum
func GetAutonomousPodMissionCriticalAssociationPeerRoleEnumValues() []AutonomousPodMissionCriticalAssociationPeerRoleEnum {
	values := make([]AutonomousPodMissionCriticalAssociationPeerRoleEnum, 0)
	for _, v := range mappingAutonomousPodMissionCriticalAssociationPeerRole {
		values = append(values, v)
	}
	return values
}

// AutonomousPodMissionCriticalAssociationLifecycleStateEnum Enum with underlying type: string
type AutonomousPodMissionCriticalAssociationLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousPodMissionCriticalAssociationLifecycleStateEnum
const (
	AutonomousPodMissionCriticalAssociationLifecycleStateProvisioning         AutonomousPodMissionCriticalAssociationLifecycleStateEnum = "PROVISIONING"
	AutonomousPodMissionCriticalAssociationLifecycleStateAvailable            AutonomousPodMissionCriticalAssociationLifecycleStateEnum = "AVAILABLE"
	AutonomousPodMissionCriticalAssociationLifecycleStateFailoverInProgress   AutonomousPodMissionCriticalAssociationLifecycleStateEnum = "FAILOVER_IN_PROGRESS"
	AutonomousPodMissionCriticalAssociationLifecycleStateSwitchoverInProgress AutonomousPodMissionCriticalAssociationLifecycleStateEnum = "SWITCHOVER_IN_PROGRESS"
	AutonomousPodMissionCriticalAssociationLifecycleStateTerminating          AutonomousPodMissionCriticalAssociationLifecycleStateEnum = "TERMINATING"
	AutonomousPodMissionCriticalAssociationLifecycleStateTerminated           AutonomousPodMissionCriticalAssociationLifecycleStateEnum = "TERMINATED"
	AutonomousPodMissionCriticalAssociationLifecycleStateFailed               AutonomousPodMissionCriticalAssociationLifecycleStateEnum = "FAILED"
)

var mappingAutonomousPodMissionCriticalAssociationLifecycleState = map[string]AutonomousPodMissionCriticalAssociationLifecycleStateEnum{
	"PROVISIONING":           AutonomousPodMissionCriticalAssociationLifecycleStateProvisioning,
	"AVAILABLE":              AutonomousPodMissionCriticalAssociationLifecycleStateAvailable,
	"FAILOVER_IN_PROGRESS":   AutonomousPodMissionCriticalAssociationLifecycleStateFailoverInProgress,
	"SWITCHOVER_IN_PROGRESS": AutonomousPodMissionCriticalAssociationLifecycleStateSwitchoverInProgress,
	"TERMINATING":            AutonomousPodMissionCriticalAssociationLifecycleStateTerminating,
	"TERMINATED":             AutonomousPodMissionCriticalAssociationLifecycleStateTerminated,
	"FAILED":                 AutonomousPodMissionCriticalAssociationLifecycleStateFailed,
}

// GetAutonomousPodMissionCriticalAssociationLifecycleStateEnumValues Enumerates the set of values for AutonomousPodMissionCriticalAssociationLifecycleStateEnum
func GetAutonomousPodMissionCriticalAssociationLifecycleStateEnumValues() []AutonomousPodMissionCriticalAssociationLifecycleStateEnum {
	values := make([]AutonomousPodMissionCriticalAssociationLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousPodMissionCriticalAssociationLifecycleState {
		values = append(values, v)
	}
	return values
}
