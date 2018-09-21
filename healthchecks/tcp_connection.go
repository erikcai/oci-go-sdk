// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Oracle Health Checks Service API
//
// Health Checks Service API.  This API allows clients to configure and run probes (tests)
// that will be executed on one or more global vantage points to monitor OCI assets.  The API
// supports running on-demand probes as well as retrieving historical results.
//

package healthchecks

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TcpConnection TCP connection results.  All durations are in milliseconds.
type TcpConnection struct {

	// The connection IP address.
	Address *string `mandatory:"false" json:"address"`

	// The port.
	Port *int `mandatory:"false" json:"port"`

	// Total connect duration, calculated by connectEnd - connectStart.
	ConnectDuration *float64 `mandatory:"false" json:"connectDuration"`

	// The duration to secure the connection.  This value will be zero for
	// insecure connections.  Calculated by connectEnd - secureConnectionStart.
	SecureConnectDuration *float64 `mandatory:"false" json:"secureConnectDuration"`
}

func (m TcpConnection) String() string {
	return common.PointerString(m)
}
