// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// InstancePowerActionDetails A base object for all types of Instance Power Action requests.
type InstancePowerActionDetails interface {
}

type instancepoweractiondetails struct {
	JsonData   []byte
	ActionType string `json:"actionType"`
}

// UnmarshalJSON unmarshals json
func (m *instancepoweractiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinstancepoweractiondetails instancepoweractiondetails
	s := struct {
		Model Unmarshalerinstancepoweractiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ActionType = s.Model.ActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *instancepoweractiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.ActionType {
	case "reboot_to_recovery_image":
		mm := RebootToRecoveryImagePowerActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

func (m instancepoweractiondetails) String() string {
	return common.PointerString(m)
}
