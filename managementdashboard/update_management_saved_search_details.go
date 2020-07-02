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

// UpdateManagementSavedSearchDetails Properties of a saved search.
type UpdateManagementSavedSearchDetails struct {

	// Display name for saved search.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Id for application (LA, APM, etc.) that owners this saved search.  Each owner has a unique Id.
	ProviderId *string `mandatory:"true" json:"providerId"`

	// Version.
	ProviderVersion *string `mandatory:"true" json:"providerVersion"`

	// Name for application (LA, APM, etc.) that owners this saved search.
	ProviderName *string `mandatory:"true" json:"providerName"`

	// String boolean ("true" or "false").
	IsWidget *bool `mandatory:"true" json:"isWidget"`

	// String boolean ("true" or "false").
	IsDashboardIneligible *bool `mandatory:"true" json:"isDashboardIneligible"`

	// The ocid of the compartment that owns the saved search.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// String boolean ("true" or "false") to indicate Out Of the Box saved search.
	IsOOBSavedSearch *bool `mandatory:"true" json:"isOOBSavedSearch"`

	// Description.
	Description *string `mandatory:"true" json:"description"`

	// Template.
	WidgetTemplate *string `mandatory:"true" json:"widgetTemplate"`

	// View Model
	WidgetVM *string `mandatory:"true" json:"widgetVM"`

	// Dashboard JSON
	WidgetOption *interface{} `mandatory:"true" json:"widgetOption"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateManagementSavedSearchDetails) String() string {
	return common.PointerString(m)
}
