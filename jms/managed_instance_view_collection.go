// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets/Query API
//
// Java Management Service Fleets/Query API
//

package jms

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// ManagedInstanceViewCollection Results of a managed instance search. Contains ManagedInstanceView items.
type ManagedInstanceViewCollection struct {

	// List of managed instances.
	Items []ManagedInstanceView `mandatory:"true" json:"items"`
}

func (m ManagedInstanceViewCollection) String() string {
	return common.PointerString(m)
}
