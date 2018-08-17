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

// TopicAttributesDetails topic attributes details
type TopicAttributesDetails struct {

	// The description of the topic.
	Description *string `mandatory:"true" json:"description"`
}

func (m TopicAttributesDetails) String() string {
	return common.PointerString(m)
}
