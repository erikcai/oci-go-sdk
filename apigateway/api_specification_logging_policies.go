// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
//

package apigateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ApiSpecificationLoggingPolicies Policies controlling the pushing of logs to OCI Public Logging.
type ApiSpecificationLoggingPolicies struct {

	// Access log policy for API requests.
	AccessLog *AccessLogPolicy `mandatory:"false" json:"accessLog"`

	// Execution log policy for API requests.
	ExecutionLog *ExecutionLogPolicy `mandatory:"false" json:"executionLog"`
}

func (m ApiSpecificationLoggingPolicies) String() string {
	return common.PointerString(m)
}
