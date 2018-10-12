// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Event Control Service API
//
// API for managing event rules and actions.
// For more information, see Overview of Events (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package cloudevents

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RuleSummary The summary details of event rules. For more information, see
// Managing Rules for Events (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Events/Task/managingrulesactions.htm)
type RuleSummary struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) of this rule.
	Id *string `mandatory:"true" json:"id"`

	// A string that describes the rule. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	// Example: `"This rule sends a notification upon completion of DbaaS backup."`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Specifies the event that will trigger the actions associated with this rule.
	// Example: `"eventType": "DbaaSBackupCompleted"`
	Condition *string `mandatory:"true" json:"condition"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this rule belongs.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Whether or not this rule is currently enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The time this rule was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2018-09-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m RuleSummary) String() string {
	return common.PointerString(m)
}
