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

// StaticPublicKey a type of public key that is specified as part of the deployment.
type StaticPublicKey interface {

	// the ID of the key
	GetKid() *string
}

type staticpublickey struct {
	JsonData []byte
	Kid      *string `mandatory:"true" json:"kid"`
	Format   string  `json:"format"`
}

// UnmarshalJSON unmarshals json
func (m *staticpublickey) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstaticpublickey staticpublickey
	s := struct {
		Model Unmarshalerstaticpublickey
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kid = s.Model.Kid
	m.Format = s.Model.Format

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *staticpublickey) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Format {
	case "JSON_WEB_KEY":
		mm := JsonWebKey{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PEM":
		mm := PemEncodedPublicKey{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "X509_CERTIFICATE":
		mm := X509Certificate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKid returns Kid
func (m staticpublickey) GetKid() *string {
	return m.Kid
}

func (m staticpublickey) String() string {
	return common.PointerString(m)
}

// StaticPublicKeyFormatEnum Enum with underlying type: string
type StaticPublicKeyFormatEnum string

// Set of constants representing the allowable values for StaticPublicKeyFormatEnum
const (
	StaticPublicKeyFormatJsonWebKey      StaticPublicKeyFormatEnum = "JSON_WEB_KEY"
	StaticPublicKeyFormatPem             StaticPublicKeyFormatEnum = "PEM"
	StaticPublicKeyFormatX509Certificate StaticPublicKeyFormatEnum = "X509_CERTIFICATE"
)

var mappingStaticPublicKeyFormat = map[string]StaticPublicKeyFormatEnum{
	"JSON_WEB_KEY":     StaticPublicKeyFormatJsonWebKey,
	"PEM":              StaticPublicKeyFormatPem,
	"X509_CERTIFICATE": StaticPublicKeyFormatX509Certificate,
}

// GetStaticPublicKeyFormatEnumValues Enumerates the set of values for StaticPublicKeyFormatEnum
func GetStaticPublicKeyFormatEnumValues() []StaticPublicKeyFormatEnum {
	values := make([]StaticPublicKeyFormatEnum, 0)
	for _, v := range mappingStaticPublicKeyFormat {
		values = append(values, v)
	}
	return values
}
