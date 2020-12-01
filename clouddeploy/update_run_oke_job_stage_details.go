// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v30/common"
)

// UpdateRunOkeJobStageDetails Specifies Run job on OKE stage.
type UpdateRunOkeJobStageDetails struct {

	// Stage Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A timeout in seconds, this stage may take to return.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	StagePredecessorCollection *StagePredecessorCollection `mandatory:"false" json:"stagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// OKE Cluster environment ocid.
	OkeClusterEnvironmentId *string `mandatory:"false" json:"okeClusterEnvironmentId"`

	// KubernetesManifest artifact ocid.
	KubernetesManifestArtifactId *string `mandatory:"false" json:"kubernetesManifestArtifactId"`

	// Default Namespace to be used for Kubernetes Job when not specified in the manifest
	Namespace *string `mandatory:"false" json:"namespace"`

	// Specifies synchronous or asynchronous mode of execution for this stage.
	ExecutionMode RunOkeJobStageExecutionModeEnum `mandatory:"false" json:"executionMode,omitempty"`
}

//GetDisplayName returns DisplayName
func (m UpdateRunOkeJobStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m UpdateRunOkeJobStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m UpdateRunOkeJobStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m UpdateRunOkeJobStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateRunOkeJobStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateRunOkeJobStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateRunOkeJobStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateRunOkeJobStageDetails UpdateRunOkeJobStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeUpdateRunOkeJobStageDetails
	}{
		"RUN_OKE_JOB",
		(MarshalTypeUpdateRunOkeJobStageDetails)(m),
	}

	return json.Marshal(&s)
}
