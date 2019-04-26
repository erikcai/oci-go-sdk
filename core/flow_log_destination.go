// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// FlowLogDestination Where to store the flow logs.
type FlowLogDestination interface {
}

type flowlogdestination struct {
	JsonData        []byte
	DestinationType string `json:"destinationType"`
}

// UnmarshalJSON unmarshals json
func (m *flowlogdestination) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerflowlogdestination flowlogdestination
	s := struct {
		Model Unmarshalerflowlogdestination
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DestinationType = s.Model.DestinationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *flowlogdestination) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DestinationType {
	case "OBJECT_STORAGE":
		mm := FlowLogObjectStorageDestination{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m flowlogdestination) String() string {
	return common.PointerString(m)
}
