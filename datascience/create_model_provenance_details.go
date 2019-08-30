// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Science API
//
// The Data Science service enables data science teams to organize their work, easily access data and computing resources, and build, train, deploy, and manage ML/AI models on the Oracle Cloud.
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateModelProvenanceDetails Model provenance gives data scientists information about the origin of their model. This information allows data scientists to reproduce the development environment in which the model was trained.
type CreateModelProvenanceDetails struct {

	// For model reproducibility purposes. URL of the git repository associated with model training.
	RepositoryUrl *string `mandatory:"false" json:"repositoryUrl"`

	// For model reproducibility purposes. Branch of the git repository associated with model training.
	GitBranch *string `mandatory:"false" json:"gitBranch"`

	// For model reproducibility purposes. Commit ID of the git repository associated with model training.
	GitCommit *string `mandatory:"false" json:"gitCommit"`

	// For model reproducibility purposes. Path to model artifacts.
	ScriptDir *string `mandatory:"false" json:"scriptDir"`

	// For model reproducibility purposes. Path to the python script or notebook in which the model was trained."
	TrainingScript *string `mandatory:"false" json:"trainingScript"`
}

func (m CreateModelProvenanceDetails) String() string {
	return common.PointerString(m)
}
