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

// ProblemSummary Summary of the Problem.
type ProblemSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Identifier of the rule
	DetectorRuleId *string `mandatory:"false" json:"detectorRuleId"`

	// The Risk Level
	RiskLevel RiskLevelEnum `mandatory:"false" json:"riskLevel,omitempty"`

	// Identifier of the Resource
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// DisplayName of the Identifier
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// Type of the Resource
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// user defined labels on the problem
	Labels []string `mandatory:"false" json:"labels"`

	// The time the Problem was last detected. An RFC3339 formatted datetime string
	TimeLastDetected *common.SDKTime `mandatory:"false" json:"timeLastDetected"`

	// The current state of the Problem.
	LifecycleState ProblemLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The lifecycleDetail will give more detail on the substate of the lifecycleState.
	LifecycleDetail ProblemLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetail,omitempty"`

	// possible type of detectors(activity/configuration/ERPCloud) associated with the Problem.
	DetectorId DetectorEnumEnum `mandatory:"false" json:"detectorId,omitempty"`

	// region where the problem is found
	Region *string `mandatory:"false" json:"region"`

	// targetId associated with the problem.
	TargetId *string `mandatory:"false" json:"targetId"`
}

func (m ProblemSummary) String() string {
	return common.PointerString(m)
}
