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

// ListenerRuleSummary Information about a Rule.
type ListenerRuleSummary struct {

	// Rule object that was applied to a listener.
	Rule Rule `mandatory:"false" json:"rule"`

	// Name of the ruleset to which rule belongs to
	RuleSetName *string `mandatory:"false" json:"ruleSetName"`
}

func (m ListenerRuleSummary) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ListenerRuleSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Rule        rule    `json:"rule"`
		RuleSetName *string `json:"ruleSetName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	nn, e := model.Rule.UnmarshalPolymorphicJSON(model.Rule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Rule = nn.(Rule)
	} else {
		m.Rule = nil
	}
	m.RuleSetName = model.RuleSetName
	return
}
