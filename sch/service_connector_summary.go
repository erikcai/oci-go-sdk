// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connectors API
//
// A description of the Connectors API
//

package sch

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ServiceConnectorSummary Summary of the ServiceConnector.
type ServiceConnectorSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// The ServiceConnector name,.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the ServiceConnector was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the ServiceConnector was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the ServiceConnector.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the resource.
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m ServiceConnectorSummary) String() string {
	return common.PointerString(m)
}
