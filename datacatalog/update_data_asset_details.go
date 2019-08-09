// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DataCatalog API
//
// A description of the DataCatalog API
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateDataAssetDetails Properties used in Data Asset update operations.
type UpdateDataAssetDetails struct {

	// The display name of a user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the Data Asset.
	Description *string `mandatory:"false" json:"description"`
}

func (m UpdateDataAssetDetails) String() string {
	return common.PointerString(m)
}
