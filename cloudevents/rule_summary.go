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

// RuleSummary Summary representation of a Rule.
type RuleSummary struct {

	// OCID of the rule
	Id *string `mandatory:"true" json:"id"`

	// display name of the rule
	DisplayName *string `mandatory:"true" json:"displayName"`

	// specifies what the rule matches
	Condition *string `mandatory:"true" json:"condition"`

	// OCID of the compartment the rule belongs to
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// whether or not this rule is currently enabled
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// time rule was created
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m RuleSummary) String() string {
	return common.PointerString(m)
}
