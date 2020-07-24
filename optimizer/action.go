// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Optimizer API
//
// The API for the OCI Optimizer
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Action Recommended action.
type Action struct {

	// The resource action status.
	Type ActionTypeEnum `mandatory:"true" json:"type"`

	// Description of the recommended action
	Description *string `mandatory:"true" json:"description"`

	// URL of the KB Article on how to perform the action.
	Url *string `mandatory:"true" json:"url"`
}

func (m Action) String() string {
	return common.PointerString(m)
}
