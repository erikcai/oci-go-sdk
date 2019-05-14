// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Quotas APIs
//
// APIs for managing Compartment Resource Quotas.
//

package limits

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateQuotaDetails Request object for create quota operation.
type CreateQuotaDetails struct {

	// The OCID of the compartment containing the resource this quota applies to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The description you assign to the quota.
	Description *string `mandatory:"true" json:"description"`

	// The name you assign to the quota during creation. The name must be unique across all quotas
	// in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// An array of quota statements written in the declarative language.
	Statements []string `mandatory:"true" json:"statements"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateQuotaDetails) String() string {
	return common.PointerString(m)
}
