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

// AllowRule An object that represents the action of configuring an access control rule. Access control rules permit access
// to application resources based on user-specified match conditions. This rule applies only to HTTP listeners.
// **NOTES:**
// *  If you do not specify any access control rules, the default rule is to allow all traffic.
// *  If you add access control rules, the load balancer denies any traffic that does not match the rules.
// *  Maximum of two match conditions can be specified in a rule.
// *  You can specify this rule only with the following `RuleCondition` combinations:
//     *  `SOURCE_IP_ADDRESS`
//     *  `SOURCE_VCN_ID`
//     *  `SOURCE_VCN_ID", "SOURCE_VCN_IP_ADDRESS`
type AllowRule struct {
	Conditions []RuleCondition `mandatory:"true" json:"conditions"`

	// A brief description of the access control rule. Avoid entering confidential information.
	// example: `192.168.0.0/16 and 2001:db8::/32 are trusted clients. Whitelist them.`
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
