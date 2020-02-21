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

// JsonWebKey a key represented by a JSON structure
type JsonWebKey struct {

	// the ID of the key
	Kid *string `mandatory:"true" json:"kid"`

	// the algorithm intended for use with this key
	Alg *string `mandatory:"false" json:"alg"`

	// Base64urlUint-encoded modulus for RSA public key.  Required when type is RSA.
	N *string `mandatory:"false" json:"n"`

	// Base64urlUint-encoded exponent for RSA public key. Required when type is RSA.
	E *string `mandatory:"false" json:"e"`

	// URI for an x.509 certificate or certificate chain
	X5u *string `mandatory:"false" json:"x5u"`

	// list of base64 encoded certificates in the certificate chain
	X5c []string `mandatory:"false" json:"x5c"`

	// thumbprint/fingerprint of the x.509 certificate
	X5t *string `mandatory:"false" json:"x5t"`

	// The type of the key
	Kty JsonWebKeyKtyEnum `mandatory:"true" json:"kty"`

	// intended use of the public key, whether for encrypting or verifying signatures
	Use JsonWebKeyUseEnum `mandatory:"false" json:"use,omitempty"`

	// the operations for which this key is intended to be used
	KeyOps JsonWebKeyKeyOpsEnum `mandatory:"false" json:"keyOps,omitempty"`
}

//GetKid returns Kid
func (m JsonWebKey) GetKid() *string {
	return m.Kid
}

func (m JsonWebKey) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m JsonWebKey) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeJsonWebKey JsonWebKey
	s := struct {
		DiscriminatorParam string `json:"format"`
		MarshalTypeJsonWebKey
	}{
		"JSON_WEB_KEY",
		(MarshalTypeJsonWebKey)(m),
	}

	return json.Marshal(&s)
}

// JsonWebKeyKtyEnum Enum with underlying type: string
type JsonWebKeyKtyEnum string

// Set of constants representing the allowable values for JsonWebKeyKtyEnum
const (
	JsonWebKeyKtyRsa JsonWebKeyKtyEnum = "RSA"
	JsonWebKeyKtyEc  JsonWebKeyKtyEnum = "EC"
	JsonWebKeyKtyOct JsonWebKeyKtyEnum = "oct"
)

var mappingJsonWebKeyKty = map[string]JsonWebKeyKtyEnum{
	"RSA": JsonWebKeyKtyRsa,
	"EC":  JsonWebKeyKtyEc,
	"oct": JsonWebKeyKtyOct,
}

// GetJsonWebKeyKtyEnumValues Enumerates the set of values for JsonWebKeyKtyEnum
func GetJsonWebKeyKtyEnumValues() []JsonWebKeyKtyEnum {
	values := make([]JsonWebKeyKtyEnum, 0)
	for _, v := range mappingJsonWebKeyKty {
		values = append(values, v)
	}
	return values
}

// JsonWebKeyUseEnum Enum with underlying type: string
type JsonWebKeyUseEnum string

// Set of constants representing the allowable values for JsonWebKeyUseEnum
const (
	JsonWebKeyUseSig JsonWebKeyUseEnum = "sig"
	JsonWebKeyUseEnc JsonWebKeyUseEnum = "enc"
)

var mappingJsonWebKeyUse = map[string]JsonWebKeyUseEnum{
	"sig": JsonWebKeyUseSig,
	"enc": JsonWebKeyUseEnc,
}

// GetJsonWebKeyUseEnumValues Enumerates the set of values for JsonWebKeyUseEnum
func GetJsonWebKeyUseEnumValues() []JsonWebKeyUseEnum {
	values := make([]JsonWebKeyUseEnum, 0)
	for _, v := range mappingJsonWebKeyUse {
		values = append(values, v)
	}
	return values
}

// JsonWebKeyKeyOpsEnum Enum with underlying type: string
type JsonWebKeyKeyOpsEnum string

// Set of constants representing the allowable values for JsonWebKeyKeyOpsEnum
const (
	JsonWebKeyKeyOpsVerify JsonWebKeyKeyOpsEnum = "verify"
)

var mappingJsonWebKeyKeyOps = map[string]JsonWebKeyKeyOpsEnum{
	"verify": JsonWebKeyKeyOpsVerify,
}

// GetJsonWebKeyKeyOpsEnumValues Enumerates the set of values for JsonWebKeyKeyOpsEnum
func GetJsonWebKeyKeyOpsEnumValues() []JsonWebKeyKeyOpsEnum {
	values := make([]JsonWebKeyKeyOpsEnum, 0)
	for _, v := range mappingJsonWebKeyKeyOps {
		values = append(values, v)
	}
	return values
}
