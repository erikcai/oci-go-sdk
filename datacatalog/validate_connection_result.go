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

// ValidateConnectionResult Details regarding the validation of a connection resource.
type ValidateConnectionResult struct {

	// The message from the connection validation
	Message *string `mandatory:"true" json:"message"`

	// The name of the connection validated
	ConnectionName *string `mandatory:"false" json:"connectionName"`

	// The status code returned from the connection validation
	StatusCode *string `mandatory:"false" json:"statusCode"`
}

func (m ValidateConnectionResult) String() string {
	return common.PointerString(m)
}
