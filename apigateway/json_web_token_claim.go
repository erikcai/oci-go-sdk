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

// JsonWebTokenClaim an individual claim represented by the token
type JsonWebTokenClaim struct {

	// the type of the claim
	Key *string `mandatory:"false" json:"key"`

	// The list of acceptable values for a given claim.
	// If this value is null or an empty array, then validation consists of making sure this claim exists.
	Values []string `mandatory:"false" json:"values"`

	// indicates whether this claim is required as part of the payload for a JWT
	IsRequired *bool `mandatory:"false" json:"isRequired"`
}

func (m JsonWebTokenClaim) String() string {
	return common.PointerString(m)
}
