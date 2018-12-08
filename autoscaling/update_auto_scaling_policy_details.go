// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Auto Scaling API
//
// Auto Scaling API spec
//

package autoscaling

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateAutoScalingPolicyDetails The representation of UpdateAutoScalingPolicyDetails
type UpdateAutoScalingPolicyDetails interface {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string
}

type updateautoscalingpolicydetails struct {
	JsonData    []byte
	DisplayName *string `mandatory:"false" json:"displayName"`
	PolicyType  string  `json:"policyType"`
}

// UnmarshalJSON unmarshals json
func (m *updateautoscalingpolicydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateautoscalingpolicydetails updateautoscalingpolicydetails
	s := struct {
		Model Unmarshalerupdateautoscalingpolicydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.PolicyType = s.Model.PolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateautoscalingpolicydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PolicyType {
	case "threshold":
		mm := UpdateThresholdPolicyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m updateautoscalingpolicydetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m updateautoscalingpolicydetails) String() string {
	return common.PointerString(m)
}
