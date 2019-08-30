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

// NotebookSessionShapeSummary A compute shape used to launch a notebook session compute instance.
type NotebookSessionShapeSummary struct {

	// The name of the shape.
	Name *string `mandatory:"true" json:"name"`

	// The number of cores associated with this notebook session shape
	CoreCount *int `mandatory:"true" json:"coreCount"`

	// The amount of memory in GBs associated with this notebook session shape
	MemoryInGBs *int `mandatory:"true" json:"memoryInGBs"`
}

func (m NotebookSessionShapeSummary) String() string {
	return common.PointerString(m)
}
