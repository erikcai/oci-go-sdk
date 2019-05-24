// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ApplicationOciLogging Configure OCI based logging.
type ApplicationOciLogging struct {

	// Enables OCI based logging. This overrides any configured `syslogUrl`.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

func (m ApplicationOciLogging) String() string {
	return common.PointerString(m)
}
