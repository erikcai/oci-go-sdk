// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Overview of the Compute Service (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
// **Note:** Autoscaling is not available in Government Cloud tenancies. For more information, see
// Information for Oracle Cloud Infrastructure Government Cloud Customers (https://docs.cloud.oracle.com/Content/General/Concepts/govinfo.htm).
//

package autoscaling

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutoScalingPolicySummary Summary information for an autoscaling policy.
type AutoScalingPolicySummary struct {

	// The ID of the autoscaling policy that is assigned after creation.
	Id *string `mandatory:"true" json:"id"`

	// The type of autoscaling policy.
	PolicyType *string `mandatory:"true" json:"policyType"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Boolean field indicated whether this policy is enabled or not.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

func (m AutoScalingPolicySummary) String() string {
	return common.PointerString(m)
}
