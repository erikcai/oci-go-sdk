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

// DbSystemEndpoint The representation of DbSystemEndpoint
type DbSystemEndpoint struct {

	// The Type of Endpoint for the DbSystem.
	EndpointType DbSystemEndpointTypeEnum `mandatory:"true" json:"endpointType"`

	// The IP address MySQLaaS instance is configured to listen on.
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The port for MySQLaaS instance to listen on.
	Port *int `mandatory:"true" json:"port"`

	// The network port on which X Plugin listens for TCP/IP connections. This is the X Plugin equivalent of port.
	MysqlxPort *int `mandatory:"true" json:"mysqlxPort"`
}

func (m DbSystemEndpoint) String() string {
	return common.PointerString(m)
}
