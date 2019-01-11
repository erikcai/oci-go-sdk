// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WafLog A list of Web Application Firewall log entries. Each entry is a JSON object whose fields vary based on log type. Logs record what rules and countermeasures are triggered by requests and are used as a basis to move request handling into block mode.
type WafLog struct {
}

func (m WafLog) String() string {
	return common.PointerString(m)
}
