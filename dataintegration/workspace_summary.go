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

// WorkspaceSummary Summary of a Workspace.
type WorkspaceSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"false" json:"id"`

	// Workspace description
	Description *string `mandatory:"false" json:"description"`

	// Data Integration Workspace display name, workspaces can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The time the Data Integration Workspace was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Data Integration Workspace was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the workspace.
	LifecycleState WorkspaceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`
}

func (m WorkspaceSummary) String() string {
	return common.PointerString(m)
}
