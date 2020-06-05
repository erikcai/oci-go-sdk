// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ResyncResourceSummary Summary of the resync resource
type ResyncResourceSummary struct {

	// OCID of Tenant
	TenantId *string `mandatory:"true" json:"tenantId"`

	// List of detector recipes
	DetectorRecipes []ResyncDetectorRecipeSummary `mandatory:"false" json:"detectorRecipes"`

	// List of responder recipes
	ResponderRecipes []ResyncResponderRecipeSummary `mandatory:"false" json:"responderRecipes"`

	// List of managed List
	ManagedLists []ResyncManagedListSummary `mandatory:"false" json:"managedLists"`
}

func (m ResyncResourceSummary) String() string {
	return common.PointerString(m)
}
