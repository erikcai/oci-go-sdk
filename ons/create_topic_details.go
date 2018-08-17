// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// ONS Gateway API
//
// A description of the ONS Gateway API
//

package ons

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateTopicDetails The details of the topic being created.
type CreateTopicDetails struct {

	// The name of the topic being created.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment where the topic is being created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The description of the topic being created.
	Description *string `mandatory:"false" json:"description"`
}

func (m CreateTopicDetails) String() string {
	return common.PointerString(m)
}
