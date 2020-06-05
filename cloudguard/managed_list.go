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

// ManagedList A cloud guard list containing one or more items of a list type
type ManagedList struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// ManagedList display name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// type of the list
	ListType ManagedListTypeEnum `mandatory:"true" json:"listType"`

	// ManagedList description
	Description *string `mandatory:"false" json:"description"`

	// List of ManagedListItem
	ListItems []string `mandatory:"false" json:"listItems"`

	// provider of the feed
	FeedProvider FeedProviderTypeEnum `mandatory:"false" json:"feedProvider,omitempty"`

	// If this list is editable or not
	IsEditable *bool `mandatory:"false" json:"isEditable"`

	// The time the the resource was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the resource.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ManagedList) String() string {
	return common.PointerString(m)
}
