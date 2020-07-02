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

// UpdateManagementDashboardDetails Properties for a dashboard.
type UpdateManagementDashboardDetails struct {

	// Provider Id.
	ProviderId *string `mandatory:"true" json:"providerId"`

	// Provider name.
	ProviderName *string `mandatory:"true" json:"providerName"`

	// Provider version.
	ProviderVersion *string `mandatory:"true" json:"providerVersion"`

	// Dashboard tiles array
	Tiles []ManagementDashboardTileDetails `mandatory:"true" json:"tiles"`

	// Display name for dashboard.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Dashboard's description.
	Description *string `mandatory:"true" json:"description"`

	// Enable filtering.
	IsFilteringEnabled *bool `mandatory:"true" json:"isFilteringEnabled"`

	// The ocid of the tenant that owns the dashboard.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The ocid of the compartment that owns the dashboard.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// String boolean ("true" or "false").  OOB dashboards are only provided by Oracle.  They cannot be modified by non-Oracle.
	IsOOBDashboard *bool `mandatory:"true" json:"isOOBDashboard"`

	// String boolean ("true" or "false").  When false, description is not shown.
	IsDescriptionEnabled *bool `mandatory:"true" json:"isDescriptionEnabled"`

	// String boolean ("true" or "false").  When false, time range is disabled.
	IsTimeRangeEnabled *bool `mandatory:"true" json:"isTimeRangeEnabled"`

	// String boolean ("true" or "false").  When false, dashboard is not automatically refreshed in intervals.
	IsRefreshEnabled *bool `mandatory:"true" json:"isRefreshEnabled"`

	// String boolean ("true" or "false").  When false, dashboard is not shown in dashboard home.
	IsShowInHome *bool `mandatory:"true" json:"isShowInHome"`

	// Screenshot.
	ScreenShot *string `mandatory:"true" json:"screenShot"`

	// NORMAL meaning single dashboard, or SET meaning dashboard set
	Type *string `mandatory:"true" json:"type"`

	// String boolean ("true" or "false").
	IsFavorite *bool `mandatory:"true" json:"isFavorite"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateManagementDashboardDetails) String() string {
	return common.PointerString(m)
}
