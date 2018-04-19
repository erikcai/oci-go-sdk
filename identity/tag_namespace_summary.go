// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TagNamespaceSummary A bag of tags that is attached to a compartment and has unique existence in tenancy.
type TagNamespaceSummary struct {

	// The OCID of the tagNamespace.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the compartment which the namespace is attached to.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The name of the tagNamespace. It must be unique across all tagNamespaces in the tenancy and cannot be changed.
	Name *string `mandatory:"false" json:"name"`

	// The description you assign to the tagNamespace.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Indicated whether or not the tagNamespace is retired
	IsRetired *bool `mandatory:"false" json:"isRetired"`

	// Date and time the tagNamespace was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m TagNamespaceSummary) String() string {
	return common.PointerString(m)
}
