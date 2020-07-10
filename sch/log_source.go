// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connectors API
//
// A description of the Connectors API
//

package sch

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LogSource The log source.
type LogSource struct {

	// The compartment OCID of the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The log group identifier.
	LogGroupId *string `mandatory:"false" json:"logGroupId"`

	// The log identifier.
	LogId *string `mandatory:"false" json:"logId"`
}

func (m LogSource) String() string {
	return common.PointerString(m)
}
