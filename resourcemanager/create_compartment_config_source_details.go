// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service. Use this API to install, configure, and manage resources via the "infrastructure-as-code" model. For more information, see Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateCompartmentConfigSourceDetails Property details for uploading the specified compartment as the configuration source.
type CreateCompartmentConfigSourceDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to use for creating the stack. The new stack will include definitions for supported resource types in this compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// File path to the directory from which Terraform runs.
	// If not specified, the root directory is used.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`
}

//GetWorkingDirectory returns WorkingDirectory
func (m CreateCompartmentConfigSourceDetails) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m CreateCompartmentConfigSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateCompartmentConfigSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateCompartmentConfigSourceDetails CreateCompartmentConfigSourceDetails
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeCreateCompartmentConfigSourceDetails
	}{
		"COMPARTMENT_CONFIG_SOURCE",
		(MarshalTypeCreateCompartmentConfigSourceDetails)(m),
	}

	return json.Marshal(&s)
}
