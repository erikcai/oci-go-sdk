// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CertificateCollection Collection of certificate summaries.
type CertificateCollection struct {

	// Certificate summaries.
	Items []CertificateSummary `mandatory:"true" json:"items"`
}

func (m CertificateCollection) String() string {
	return common.PointerString(m)
}
