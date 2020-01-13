// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// VersionOptionsMetadata Metadata about the collection of options valid for a specific MySQL version.
type VersionOptionsMetadata struct {

	// The MySQL version for which the options are valid.
	ForVersion *string `mandatory:"true" json:"forVersion"`

	Options []OptionMetadata `mandatory:"true" json:"options"`
}

func (m VersionOptionsMetadata) String() string {
	return common.PointerString(m)
}
