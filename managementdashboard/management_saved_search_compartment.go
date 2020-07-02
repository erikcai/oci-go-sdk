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

// ManagementSavedSearchCompartment Properties for a saved search compartment.
type ManagementSavedSearchCompartment struct {

	// The ocid of the compartment that owns the saved search.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ManagementSavedSearchCompartment) String() string {
	return common.PointerString(m)
}
