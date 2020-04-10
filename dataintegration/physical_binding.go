// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// PhysicalBinding The physical binding object.
type PhysicalBinding struct {

	// supportedOperations
	SupportedOperations *int `mandatory:"false" json:"supportedOperations"`

	// The physical entity reference.
	PhysicalEntity *interface{} `mandatory:"false" json:"physicalEntity"`
}

func (m PhysicalBinding) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PhysicalBinding) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePhysicalBinding PhysicalBinding
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypePhysicalBinding
	}{
		"PHYSICAL_BINDING",
		(MarshalTypePhysicalBinding)(m),
	}

	return json.Marshal(&s)
}
