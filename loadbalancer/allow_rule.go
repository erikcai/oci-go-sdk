// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// AllowRule Configuration of an access control rule that permits access to the application
// resources exposed by a listener based on user specified match conditions.
// This rule applies to only HTTP listeners.
// Atleast one match condition should be specified in the rule.
// **NOTE:** When user does not specify any access control rules, the implicit default rule is to allow all traffic.
// **NOTE:** When user adds any access control rules, all the traffic not matching with any rules will be implicitly denied.
// **NOTE:** User can specify this rule only with the following combinations of RuleCondition types:
//             - ["SOURCE_IP_ADDRESS"]
//             - ["SOURCE_VCN_ID"]
//             - ["SOURCE_VCN_ID", "SOURCE_VCN_IP_ADDRESS"]
type AllowRule struct {
	Conditions []RuleCondition `mandatory:"true" json:"conditions"`

	// Brief description of the access control rule.
	Description *string `mandatory:"false" json:"description"`
}

func (m AllowRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AllowRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAllowRule AllowRule
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeAllowRule
	}{
		"ALLOW",
		(MarshalTypeAllowRule)(m),
	}

	return json.Marshal(&s)
}
