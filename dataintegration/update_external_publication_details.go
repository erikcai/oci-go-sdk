// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateExternalPublicationDetails Properties used to update a published a OCI DataFlow object.
type UpdateExternalPublicationDetails struct {

	// The OCID of the compartment where the application is created in the OCI Dataflow Service.
	ApplicationCompartmentId *string `mandatory:"true" json:"applicationCompartmentId"`

	// The name of the application
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The unique OCID of the identifier that is returned after creating the OCI Dataflow application.
	ApplicationId *string `mandatory:"false" json:"applicationId"`

	// The details of the dataflow or the application
	Description *string `mandatory:"false" json:"description"`

	ResourceConfiguration *ResourceConfiguration `mandatory:"false" json:"resourceConfiguration"`

	ConfigurationDetails *ConfigurationDetails `mandatory:"false" json:"configurationDetails"`
}

func (m UpdateExternalPublicationDetails) String() string {
	return common.PointerString(m)
}
