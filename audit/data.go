// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Audit API
//
// API for the Audit Service. Use this API for compliance monitoring in your tenancy.
// For more information, see Overview of Audit (https://docs.cloud.oracle.com/iaas/Content/Audit/Concepts/auditoverview.htm).
// **Tip**: This API is good for queries, but not bulk-export operations.
//

package audit

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Data The representation of Data
type Data struct {

	// The GUID of the event grouping id.
	// Example: `examplea9-f488-4842-96cb-a10f2893b369`
	EventGroupingId *string `mandatory:"false" json:"eventGroupingId"`

	// The name of the event.
	// Example: `LaunchInstance`
	EventName *string `mandatory:"false" json:"eventName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The name of the compartment. This value is the friendly name associated with compartmentId.
	// This value can change, but the service logs the value that appeared at the time of the audit
	// event.
	// Example: `CompartmentA`
	CompartmentName *string `mandatory:"false" json:"compartmentName"`

	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. Exists for cross-compatibility only.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	Identity *Identity `mandatory:"false" json:"identity"`

	Request *Request `mandatory:"false" json:"request"`

	Response *Response `mandatory:"false" json:"response"`

	StateChange *StateChange `mandatory:"false" json:"stateChange"`

	AdditionalDetails map[string]interface{} `mandatory:"false" json:"additionalDetails"`
}

func (m Data) String() string {
	return common.PointerString(m)
}
