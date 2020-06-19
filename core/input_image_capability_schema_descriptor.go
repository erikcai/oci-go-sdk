// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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

// InputImageCapabilitySchemaDescriptor Image Capability Schema Descriptor is a type of capability for an image.
type InputImageCapabilitySchemaDescriptor interface {
}

type inputimagecapabilityschemadescriptor struct {
	JsonData       []byte
	DescriptorType string `json:"descriptorType"`
}

// UnmarshalJSON unmarshals json
func (m *inputimagecapabilityschemadescriptor) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinputimagecapabilityschemadescriptor inputimagecapabilityschemadescriptor
	s := struct {
		Model Unmarshalerinputimagecapabilityschemadescriptor
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DescriptorType = s.Model.DescriptorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *inputimagecapabilityschemadescriptor) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DescriptorType {
	case "enuminteger":
		mm := EnumIntegerInputImageCapabilityDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "boolean":
		mm := BooleanInputImageCapabilitySchemaDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "enumstring":
		mm := EnumStringInputImageCapabilitySchemaDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m inputimagecapabilityschemadescriptor) String() string {
	return common.PointerString(m)
}
