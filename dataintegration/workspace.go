// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Workspace A workspace is an organizational construct to keep multiple data integration solutions and their resources (data assets, data flows, tasks, and so on) separate from each other, helping you to stay organized. For example, you could have separate workspaces for development, testing, and production.
type Workspace struct {

	// VCN OCID where the subnet resides.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Subnet OCID for the customer connected network where databases for example reside.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Data Integration Workspace display name, workspaces can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// IP of the custom DNS server
	DnsServerIp *string `mandatory:"false" json:"dnsServerIp"`

	// DNS Zone of the custom DNS to use to resolve names.
	DnsServerZone *string `mandatory:"false" json:"dnsServerZone"`

	// Enable/disable private network connection.
	IsPrivateNetworkEnabled *bool `mandatory:"false" json:"isPrivateNetworkEnabled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Workspace description
	Description *string `mandatory:"false" json:"description"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The time the Data Integration Workspace was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Data Integration Workspace was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Lifecycle states for workspaces in Data Integration Service
	// CREATING - The resource is being created and may not be usable until the entire metadata is defined
	// UPDATING - The resource is being updated and may not be usable until all changes are commited
	// DELETING - The resource is being deleted and might require deep cleanup of children.
	// ACTIVE   - The resource is valid and available for access
	// INACTIVE - The resource might be incomplete in its definition or might have been made unavailable for
	//          administrative reasons
	// DELETED  - The resource has been deleted and isn't available
	// FAILED   - The resource is in a failed state due to validation or other errors
	LifecycleState WorkspaceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`
}

func (m Workspace) String() string {
	return common.PointerString(m)
}

// WorkspaceLifecycleStateEnum Enum with underlying type: string
type WorkspaceLifecycleStateEnum string

// Set of constants representing the allowable values for WorkspaceLifecycleStateEnum
const (
	WorkspaceLifecycleStateCreating WorkspaceLifecycleStateEnum = "CREATING"
	WorkspaceLifecycleStateActive   WorkspaceLifecycleStateEnum = "ACTIVE"
	WorkspaceLifecycleStateInactive WorkspaceLifecycleStateEnum = "INACTIVE"
	WorkspaceLifecycleStateUpdating WorkspaceLifecycleStateEnum = "UPDATING"
	WorkspaceLifecycleStateDeleting WorkspaceLifecycleStateEnum = "DELETING"
	WorkspaceLifecycleStateDeleted  WorkspaceLifecycleStateEnum = "DELETED"
	WorkspaceLifecycleStateFailed   WorkspaceLifecycleStateEnum = "FAILED"
)

var mappingWorkspaceLifecycleState = map[string]WorkspaceLifecycleStateEnum{
	"CREATING": WorkspaceLifecycleStateCreating,
	"ACTIVE":   WorkspaceLifecycleStateActive,
	"INACTIVE": WorkspaceLifecycleStateInactive,
	"UPDATING": WorkspaceLifecycleStateUpdating,
	"DELETING": WorkspaceLifecycleStateDeleting,
	"DELETED":  WorkspaceLifecycleStateDeleted,
	"FAILED":   WorkspaceLifecycleStateFailed,
}

// GetWorkspaceLifecycleStateEnumValues Enumerates the set of values for WorkspaceLifecycleStateEnum
func GetWorkspaceLifecycleStateEnumValues() []WorkspaceLifecycleStateEnum {
	values := make([]WorkspaceLifecycleStateEnum, 0)
	for _, v := range mappingWorkspaceLifecycleState {
		values = append(values, v)
	}
	return values
}
