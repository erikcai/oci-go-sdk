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

// NotebookSessionConfigurationDetails Details for the notebook session configuration.
type NotebookSessionConfigurationDetails struct {

	// The shape used to launch the notebook session compute instance.  The list of available shapes in a given compartment can be retrieved from the ListNotebookSessionShapes endpoint.
	Shape *string `mandatory:"true" json:"shape"`

	// A notebook session instance is provided with a VNIC for network access.  This specifies the OCID of the subnet to create a VNIC in.  The subnet should be in a VCN with a NAT gateway for egress to the internet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A notebook session instance is provided with a block storage volume.  This specifies the size of the volume in GBs.
	BlockStorageSizeInGBs *int `mandatory:"false" json:"blockStorageSizeInGBs"`
}

func (m NotebookSessionConfigurationDetails) String() string {
	return common.PointerString(m)
}