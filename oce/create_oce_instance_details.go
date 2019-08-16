// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OceInstance API
//
// A description of the OceInstance API
//

package oce

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateOceInstanceDetails The information about new OceInstance.
type CreateOceInstanceDetails struct {

	// OceInstance Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OceInstance Identifier
	OceInstanceType *string `mandatory:"true" json:"oceInstanceType"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Region
	Region *string `mandatory:"true" json:"region"`

	// Storage Compartment Identifier
	StorageCompartmentId *string `mandatory:"true" json:"storageCompartmentId"`

	// Service Name
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// Tenancy Identifier
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// IDCS Proof of Stripe Token
	IdcsAt *string `mandatory:"true" json:"idcsAt"`

	// Namespace
	Namespace *string `mandatory:"true" json:"namespace"`

	// Admin Email for Notification
	AdminEmail *string `mandatory:"true" json:"adminEmail"`

	// OceInstance description
	Description *string `mandatory:"false" json:"description"`

	// Account Name
	AccountName *string `mandatory:"false" json:"accountName"`

	// Target POD Name
	PodName *string `mandatory:"false" json:"podName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateOceInstanceDetails) String() string {
	return common.PointerString(m)
}
