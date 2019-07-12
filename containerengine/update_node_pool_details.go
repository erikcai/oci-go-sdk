// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateNodePoolDetails The properties that define a request to update a node pool.
type UpdateNodePoolDetails struct {

	// The new name for the cluster. Avoid entering confidential information.
	Name *string `mandatory:"false" json:"name"`

	// The version of Kubernetes to which the nodes in the node pool should be upgraded.
	KubernetesVersion *string `mandatory:"false" json:"kubernetesVersion"`

	// A list of key/value pairs to add to nodes after they join the Kubernetes cluster.
	InitialNodeLabels []KeyValue `mandatory:"false" json:"initialNodeLabels"`

	// The number of nodes to have in each subnet specified in subnetIds property. This property is deprecated,
	// use nodeConfigurationDetails instead.
	QuantityPerSubnet *int `mandatory:"false" json:"quantityPerSubnet"`

	// The OCIDs of the subnets in which to place nodes for this node pool. This property is deprecated,
	// use nodeConfigurationDetails instead. Only one of the subnetIds or nodeConfigurationDetails
	// properties can be specified.
	SubnetIds []string `mandatory:"false" json:"subnetIds"`

	// The nodes configuration of the node pool. Only one of the subnetIds or nodeConfigurationDetails
	// properties should be specified.
	NodeConfigurationDetails *UpdateNodePoolNodeConfigurationDetails `mandatory:"false" json:"nodeConfigurationDetails"`
}

func (m UpdateNodePoolDetails) String() string {
	return common.PointerString(m)
}
