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

// X509Certificate a public key represented by an X509 certificate
type X509Certificate struct {

	// the ID of the key
	Kid *string `mandatory:"true" json:"kid"`

	// x509 certificate
	Certificate *string `mandatory:"true" json:"certificate"`
}

//GetKid returns Kid
func (m X509Certificate) GetKid() *string {
	return m.Kid
}

func (m X509Certificate) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m X509Certificate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeX509Certificate X509Certificate
	s := struct {
		DiscriminatorParam string `json:"format"`
		MarshalTypeX509Certificate
	}{
		"X509_CERTIFICATE",
		(MarshalTypeX509Certificate)(m),
	}

	return json.Marshal(&s)
}
