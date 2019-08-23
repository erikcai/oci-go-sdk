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

// ParseConnectionDetails Parse Connections from the connection metadata and oracle wallet file
type ParseConnectionDetails struct {
	ConnectionDetail *Connection `mandatory:"false" json:"connectionDetail"`

	// The information used to parse the connections from the wallet file payload
	ConnectionsPayload []byte `mandatory:"false" json:"connectionsPayload"`
}

func (m ParseConnectionDetails) String() string {
	return common.PointerString(m)
}
