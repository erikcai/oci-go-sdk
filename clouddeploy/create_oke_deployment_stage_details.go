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

// CreateOkeDeploymentStageDetails Specifies the OKEDeployment stage.
type CreateOkeDeploymentStageDetails struct {

	// Stage Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Pipeline Identifier
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	StagePredecessorCollection *StagePredecessorCollection `mandatory:"true" json:"stagePredecessorCollection"`

	// OkeCluster environment ocid for deployment.
	OkeClusterEnvironmentId *string `mandatory:"true" json:"okeClusterEnvironmentId"`

	// KubernetesManifest artifact ocid, this manifest should not include any Job resource.
	KubernetesManifestArtifactId *string `mandatory:"true" json:"kubernetesManifestArtifactId"`

	// A timeout in seconds, this stage may take to return.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Default Namespace to be used for Kubernetes Deployment when not specified in the manifest
	Namespace *string `mandatory:"false" json:"namespace"`
}

//GetDisplayName returns DisplayName
func (m CreateOkeDeploymentStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetPipelineId returns PipelineId
func (m CreateOkeDeploymentStageDetails) GetPipelineId() *string {
	return m.PipelineId
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m CreateOkeDeploymentStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m CreateOkeDeploymentStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m CreateOkeDeploymentStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateOkeDeploymentStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateOkeDeploymentStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateOkeDeploymentStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOkeDeploymentStageDetails CreateOkeDeploymentStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeCreateOkeDeploymentStageDetails
	}{
		"OKE_DEPLOYMENT",
		(MarshalTypeCreateOkeDeploymentStageDetails)(m),
	}

	return json.Marshal(&s)
}
