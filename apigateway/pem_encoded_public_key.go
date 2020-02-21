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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// PemEncodedPublicKey a public key represented by its PEM encoding
type PemEncodedPublicKey struct {

	// the ID of the key
	Kid *string `mandatory:"true" json:"kid"`

	// pem encoded public key
	Key *string `mandatory:"true" json:"key"`
}

//GetKid returns Kid
func (m PemEncodedPublicKey) GetKid() *string {
	return m.Kid
}

func (m PemEncodedPublicKey) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PemEncodedPublicKey) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePemEncodedPublicKey PemEncodedPublicKey
	s := struct {
		DiscriminatorParam string `json:"format"`
		MarshalTypePemEncodedPublicKey
	}{
		"PEM",
		(MarshalTypePemEncodedPublicKey)(m),
	}

	return json.Marshal(&s)
}
