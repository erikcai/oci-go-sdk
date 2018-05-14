// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Clusters API Specification
//
// Container Engine for Kubernetes API
//

package containerengine

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateClusterDetails The properties that define a request to update a cluster.
type UpdateClusterDetails struct {

	// The new name for the cluster. Avoid entering confidential information.
	Name *string `mandatory:"false" json:"name"`

	// The version of Kubernetes to which the cluster masters should be upgraded.
	KubernetesVersion *string `mandatory:"false" json:"kubernetesVersion"`
}

func (m UpdateClusterDetails) String() string {
	return common.PointerString(m)
}
