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

// ImportConnectionDetails Import Connections from the connection metadata and oracle wallet file
type ImportConnectionDetails struct {

	// The information used to import the connections
	ConnectionsPayload []byte `mandatory:"true" json:"connectionsPayload"`

	// The connection details array that will be used to import the connections. If not , default names will be generated
	ConnectionsDetailArray []CreateConnectionDetails `mandatory:"false" json:"connectionsDetailArray"`
}

func (m ImportConnectionDetails) String() string {
	return common.PointerString(m)
}
