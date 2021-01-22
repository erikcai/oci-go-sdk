// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Synthetic
//
// API for APM Synthetic service. Use this API to query Synthethetic scripts and monitors.
//

package apmsynthetics

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// MonitorStatusCountMap monitor count per state.
// example: `{ "total" : 5, "enabled" : 3 , "disabled" : 2, "invalid" : 0 }`
type MonitorStatusCountMap struct {

	// Total number of monitors using the script
	Total *int `mandatory:"true" json:"total"`

	// Number of enabled monitors using the script
	Enabled *int `mandatory:"true" json:"enabled"`

	// Number of disabled monitors using the script
	Disabled *int `mandatory:"true" json:"disabled"`

	// Number of invalid monitors using the script
	Invalid *int `mandatory:"true" json:"invalid"`
}

func (m MonitorStatusCountMap) String() string {
	return common.PointerString(m)
}
