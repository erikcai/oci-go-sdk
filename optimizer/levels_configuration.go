// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Optimizer API
//
// The API for the OCI Optimizer
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LevelsConfiguration An array of configuration levels for each recommendation.
type LevelsConfiguration struct {

	// The array of configuration level.
	Items []LevelConfiguration `mandatory:"false" json:"items"`
}

func (m LevelsConfiguration) String() string {
	return common.PointerString(m)
}
