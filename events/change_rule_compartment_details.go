// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Events API
//
// API for the Events Service. Use this API to manage rules and actions that create automation
// in your tenancy. For more information, see Overview of Events (https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package events

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeRuleCompartmentDetails Specifies which compartment to move a rule to when changing a rule's compartment.
type ChangeRuleCompartmentDetails struct {

	// The OCID  (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment
	// into which the resource should be moved.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeRuleCompartmentDetails) String() string {
	return common.PointerString(m)
}
