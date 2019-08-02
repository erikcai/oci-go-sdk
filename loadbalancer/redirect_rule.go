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

// RedirectRule An object that represents the action of returning a specified response code along with the redirect URI
// configured for the listener for a specific path.
// The default response code is `302 Moved Temporarily`.
// This rule applies only to HTTP listeners. Only one `RedirectRule` object can be present in
// a given listener for a given originalPath.
// **NOTE:** User can specify this rule only with the RuleCondition of type 'PATH'.
type RedirectRule struct {
	Conditions []RuleCondition `mandatory:"true" json:"conditions"`

	// The HTTP status code to return when the redirection is performed.
	// The associated status line returned with the code is mapped from the standard HTTP specification.
	// Valid response codes for redirection are 301, 302, 303, 307 and 308.
	// The default value is `302 (Moved Temporarily)`.
	// Example: 301
	ResponseCode *int `mandatory:"false" json:"responseCode"`

	RedirectUri *RedirectUri `mandatory:"false" json:"redirectUri"`
}

func (m RedirectRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RedirectRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRedirectRule RedirectRule
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeRedirectRule
	}{
		"REDIRECT",
		(MarshalTypeRedirectRule)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *RedirectRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ResponseCode *int            `json:"responseCode"`
		RedirectUri  *RedirectUri    `json:"redirectUri"`
		Conditions   []rulecondition `json:"conditions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.ResponseCode = model.ResponseCode
	m.RedirectUri = model.RedirectUri
	m.Conditions = make([]RuleCondition, len(model.Conditions))
	for i, n := range model.Conditions {
		nn, err := n.UnmarshalPolymorphicJSON(n.JsonData)
		if err != nil {
			return err
		}
		if nn != nil {
			m.Conditions[i] = nn.(RuleCondition)
		} else {
			m.Conditions[i] = nil
		}
	}
	return
}
