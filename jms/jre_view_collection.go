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

// JreViewCollection Results of a Java Runtime search. Contains JreView items
type JreViewCollection struct {

	// List of Java Runtimes.
	Items []JreView `mandatory:"true" json:"items"`
}

func (m JreViewCollection) String() string {
	return common.PointerString(m)
}
