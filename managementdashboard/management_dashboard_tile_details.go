// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ManagementDashboard API
//
// Management Dashboard micro-service provides APIs to support dashboard and saved search metadata preservation, as follows 1) Save and retrieve metadata to support UI activities (Create an empty dashboard, open a saved search, etc.). 2) Check authority for each CRUD operation. 3) Validate input: Are all properties present?  Any empty values?  Are all saved searches OOB when dashboard is OOB?  Etc. 4) Import and export dashboards.
//

package managementdashboard

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ManagementDashboardTileDetails Properties of tile for a saved search.  Tile is a UI rendition of a saved search.
type ManagementDashboardTileDetails struct {

	// Display name for saved search.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Title of saved search.
	Title *string `mandatory:"true" json:"title"`

	// Description for saved search.
	Description *string `mandatory:"true" json:"description"`

	// Id of saved search.
	SavedSearchId *string `mandatory:"true" json:"savedSearchId"`

	// Description
	Type *string `mandatory:"true" json:"type"`

	// Row, Y position
	Row *int `mandatory:"true" json:"row"`

	// Column, X position
	Column *int `mandatory:"true" json:"column"`

	// Height position
	Height *int `mandatory:"true" json:"height"`

	// Width position
	Width *int `mandatory:"true" json:"width"`

	// Drill down configuration
	DrillConfig *interface{} `mandatory:"true" json:"drillConfig"`

	// Tile parameters that override certain saved search properties.
	TileParameters *interface{} `mandatory:"true" json:"tileParameters"`
}

func (m ManagementDashboardTileDetails) String() string {
	return common.PointerString(m)
}
