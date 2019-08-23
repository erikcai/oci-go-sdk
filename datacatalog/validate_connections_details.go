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

// ValidateConnectionsDetails Validate Connections from the connection metadata or oracle wallet file
type ValidateConnectionsDetails struct {

	// The connection details array that will be used to validate the connections.
	ConnectionsDetailArray []CreateConnectionDetails `mandatory:"false" json:"connectionsDetailArray"`

	// The information used to validate the connections
	ConnectionsPayload []byte `mandatory:"false" json:"connectionsPayload"`
}

func (m ValidateConnectionsDetails) String() string {
	return common.PointerString(m)
}
