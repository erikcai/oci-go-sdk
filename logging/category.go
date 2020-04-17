// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// PublicLoggingControlplane API
//
// PublicLoggingControlplane API specification
//

package logging

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Category categories for resources.
type Category struct {

	// Category name
	Name *string `mandatory:"false" json:"name"`

	// Parameters category supports.
	Parameters []Parameter `mandatory:"false" json:"parameters"`
}

func (m Category) String() string {
	return common.PointerString(m)
}
