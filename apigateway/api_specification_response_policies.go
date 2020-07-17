// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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

// ApiSpecificationResponsePolicies Global behavior applied to all responses sent by the API.
type ApiSpecificationResponsePolicies struct {

	// Perform these transformations on the HTTP response headers in the order listed.
	HeaderTransformations []HeaderTransformation `mandatory:"false" json:"headerTransformations"`
}

func (m ApiSpecificationResponsePolicies) String() string {
	return common.PointerString(m)
}
