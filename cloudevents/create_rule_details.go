// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// EventsControlService API
//
// This service exposes APIs to create, update and delete Rules. Rules are used to tap into the Events stream.
//

package cloudevents

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateRuleDetails Object used to create a rule.
type CreateRuleDetails struct {

	// A string that describes the rule
	DisplayName *string `mandatory:"true" json:"displayName"`

	// whether or not this rule is currently enabled
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// specifies what the rule matches
	Condition *string `mandatory:"true" json:"condition"`

	// OCID of the compartment to create the rule in
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A list of action object.
	Actions *ActionDetailsList `mandatory:"true" json:"actions"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. Exists for cross-compatibility only.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}'
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateRuleDetails) String() string {
	return common.PointerString(m)
}
