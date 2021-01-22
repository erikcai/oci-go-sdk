// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Synthetic
//
// API for APM Synthetic service. Use this API to query Synthethetic scripts and monitors.
//

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v33/common"
)

// VantagePointInfo Vantage Point Details.
type VantagePointInfo struct {

	// Vantage point name.
	Name *string `mandatory:"true" json:"name"`

	// A user-friendly name. Must be unique, and it can be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`
}

func (m VantagePointInfo) String() string {
	return common.PointerString(m)
}
