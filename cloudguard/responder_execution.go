// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ResponderExecution Responder Execution Object.
type ResponderExecution struct {

	// The Unique identifier of the responder execution
	Id *string `mandatory:"true" json:"id"`

	// Responder Rule id for the responder execution
	ResponderRuleId *string `mandatory:"true" json:"responderRuleId"`

	// Rule Type for the responder execution
	ResponderRuleType *string `mandatory:"true" json:"responderRuleType"`

	// Rule name for the responder execution
	ResponderRuleName *string `mandatory:"true" json:"responderRuleName"`

	// Problem id associated with the responder execution
	ProblemId *string `mandatory:"true" json:"problemId"`

	// region where the problem is found
	Region *string `mandatory:"true" json:"region"`

	// targetId of the problem for the responder execution
	TargetId *string `mandatory:"true" json:"targetId"`

	// tenant id of the problem for the responder execution
	TenantId *string `mandatory:"true" json:"tenantId"`

	// compartment id of the problem for the responder execution
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// resource type of the problem for the responder execution
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// resource name of the problem for the responder execution
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The time at which responder execution was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// possible type of responder execution states
	ResponderExecutionStatus ResponderExecutionStatesEnum `mandatory:"true" json:"responderExecutionStatus"`

	// possible type of responder execution modes
	ResponderExecutionMode ResponderExecutionModesEnum `mandatory:"true" json:"responderExecutionMode"`

	// The time at which responder execution was updated. An RFC3339 formatted datetime string
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// Message about the responder execution.
	Message *string `mandatory:"false" json:"message"`

	ResponderRuleExecutionDetails *ResponderRuleExecutionDetails `mandatory:"false" json:"responderRuleExecutionDetails"`
}

func (m ResponderExecution) String() string {
	return common.PointerString(m)
}
