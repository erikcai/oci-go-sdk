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

// ResyncManagedListSummary Summary of Managed List
type ResyncManagedListSummary struct {

	// Unique identifier for ManagedList.
	Id *string `mandatory:"true" json:"id"`

	// Managed List Identifier Name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Feed provider of Managed List
	FeedProvider FeedProviderTypeEnum `mandatory:"true" json:"feedProvider"`

	// Old list Items of Managed List
	OldListItems []string `mandatory:"true" json:"oldListItems"`

	// New list Items of Managed List
	NewListItems []string `mandatory:"true" json:"newListItems"`

	// The id of the source Managed List.
	SourceManagedListId *string `mandatory:"false" json:"sourceManagedListId"`
}

func (m ResyncManagedListSummary) String() string {
	return common.PointerString(m)
}
