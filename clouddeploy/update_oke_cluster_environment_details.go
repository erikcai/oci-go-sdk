// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v33/common"
)

// UpdateOkeClusterEnvironmentDetails Specifies the OKE Cluster environment.
type UpdateOkeClusterEnvironmentDetails struct {

	// Environment Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the OKE cluster.
	ClusterId *string `mandatory:"true" json:"clusterId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetDisplayName returns DisplayName
func (m UpdateOkeClusterEnvironmentDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m UpdateOkeClusterEnvironmentDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateOkeClusterEnvironmentDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateOkeClusterEnvironmentDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateOkeClusterEnvironmentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOkeClusterEnvironmentDetails UpdateOkeClusterEnvironmentDetails
	s := struct {
		DiscriminatorParam string `json:"environmentType"`
		MarshalTypeUpdateOkeClusterEnvironmentDetails
	}{
		"OKE_CLUSTER",
		(MarshalTypeUpdateOkeClusterEnvironmentDetails)(m),
	}

	return json.Marshal(&s)
}
