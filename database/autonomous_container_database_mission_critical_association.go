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

// AutonomousContainerDatabaseMissionCriticalAssociation The properties that define the Mission Critical Association between two different Autonomous Container Databases.
type AutonomousContainerDatabaseMissionCriticalAssociation struct {

	// The OCID of the Mission Critical Autonomous Container Database id.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database that has a relationship with the peer Autonomous Container Database.
	AutonomousContainerDatabaseId *string `mandatory:"true" json:"autonomousContainerDatabaseId"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the peer Autonomous Container Database.
	PeerAutonomousContainerDatabaseId *string `mandatory:"true" json:"peerAutonomousContainerDatabaseId"`

	// The role of the reporting Container Database in this Autonomous Mission Critical Association.
	Role AutonomousContainerDatabaseMissionCriticalAssociationRoleEnum `mandatory:"true" json:"role"`

	// The role of the peer Container Database in this Autonomous Mission Critical Association.
	PeerRole AutonomousContainerDatabaseMissionCriticalAssociationPeerRoleEnum `mandatory:"true" json:"peerRole"`

	// The OCID of the peer Autonomous Container Database Mission Critical Association.
	PeerMissionCriticalAssociationId *string `mandatory:"true" json:"peerMissionCriticalAssociationId"`

	// The current state of the Autonomous Container Database Mission Critical Association.
	LifecycleState AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Additional information about the current lifecycleState, if available.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m AutonomousContainerDatabaseMissionCriticalAssociation) String() string {
	return common.PointerString(m)
}

// AutonomousContainerDatabaseMissionCriticalAssociationRoleEnum Enum with underlying type: string
type AutonomousContainerDatabaseMissionCriticalAssociationRoleEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseMissionCriticalAssociationRoleEnum
const (
	AutonomousContainerDatabaseMissionCriticalAssociationRolePrimary AutonomousContainerDatabaseMissionCriticalAssociationRoleEnum = "PRIMARY"
	AutonomousContainerDatabaseMissionCriticalAssociationRoleStandby AutonomousContainerDatabaseMissionCriticalAssociationRoleEnum = "STANDBY"
)

var mappingAutonomousContainerDatabaseMissionCriticalAssociationRole = map[string]AutonomousContainerDatabaseMissionCriticalAssociationRoleEnum{
	"PRIMARY": AutonomousContainerDatabaseMissionCriticalAssociationRolePrimary,
	"STANDBY": AutonomousContainerDatabaseMissionCriticalAssociationRoleStandby,
}

// GetAutonomousContainerDatabaseMissionCriticalAssociationRoleEnumValues Enumerates the set of values for AutonomousContainerDatabaseMissionCriticalAssociationRoleEnum
func GetAutonomousContainerDatabaseMissionCriticalAssociationRoleEnumValues() []AutonomousContainerDatabaseMissionCriticalAssociationRoleEnum {
	values := make([]AutonomousContainerDatabaseMissionCriticalAssociationRoleEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseMissionCriticalAssociationRole {
		values = append(values, v)
	}
	return values
}

// AutonomousContainerDatabaseMissionCriticalAssociationPeerRoleEnum Enum with underlying type: string
type AutonomousContainerDatabaseMissionCriticalAssociationPeerRoleEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseMissionCriticalAssociationPeerRoleEnum
const (
	AutonomousContainerDatabaseMissionCriticalAssociationPeerRolePrimary AutonomousContainerDatabaseMissionCriticalAssociationPeerRoleEnum = "PRIMARY"
	AutonomousContainerDatabaseMissionCriticalAssociationPeerRoleStandby AutonomousContainerDatabaseMissionCriticalAssociationPeerRoleEnum = "STANDBY"
)

var mappingAutonomousContainerDatabaseMissionCriticalAssociationPeerRole = map[string]AutonomousContainerDatabaseMissionCriticalAssociationPeerRoleEnum{
	"PRIMARY": AutonomousContainerDatabaseMissionCriticalAssociationPeerRolePrimary,
	"STANDBY": AutonomousContainerDatabaseMissionCriticalAssociationPeerRoleStandby,
}

// GetAutonomousContainerDatabaseMissionCriticalAssociationPeerRoleEnumValues Enumerates the set of values for AutonomousContainerDatabaseMissionCriticalAssociationPeerRoleEnum
func GetAutonomousContainerDatabaseMissionCriticalAssociationPeerRoleEnumValues() []AutonomousContainerDatabaseMissionCriticalAssociationPeerRoleEnum {
	values := make([]AutonomousContainerDatabaseMissionCriticalAssociationPeerRoleEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseMissionCriticalAssociationPeerRole {
		values = append(values, v)
	}
	return values
}

// AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum Enum with underlying type: string
type AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum
const (
	AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateProvisioning         AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum = "PROVISIONING"
	AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateAvailable            AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum = "AVAILABLE"
	AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateFailoverInProgress   AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum = "FAILOVER_IN_PROGRESS"
	AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateSwitchoverInProgress AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum = "SWITCHOVER_IN_PROGRESS"
	AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateTerminating          AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum = "TERMINATING"
	AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateTerminated           AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum = "TERMINATED"
	AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateFailed               AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum = "FAILED"
)

var mappingAutonomousContainerDatabaseMissionCriticalAssociationLifecycleState = map[string]AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum{
	"PROVISIONING":           AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateProvisioning,
	"AVAILABLE":              AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateAvailable,
	"FAILOVER_IN_PROGRESS":   AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateFailoverInProgress,
	"SWITCHOVER_IN_PROGRESS": AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateSwitchoverInProgress,
	"TERMINATING":            AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateTerminating,
	"TERMINATED":             AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateTerminated,
	"FAILED":                 AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateFailed,
}

// GetAutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnumValues Enumerates the set of values for AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum
func GetAutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnumValues() []AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum {
	values := make([]AutonomousContainerDatabaseMissionCriticalAssociationLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseMissionCriticalAssociationLifecycleState {
		values = append(values, v)
	}
	return values
}
