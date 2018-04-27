// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// AddHttpResponseHeaderRule An object that represents the action of adding a header to a response.
// This rule will be applied only to HTTP or HTTP2 listeners.
type AddHttpResponseHeaderRule struct {

	// A header name that conforms to RFC 7230.
	// Example: `example_header_name`
	Header *string `mandatory:"true" json:"header"`

	// A header value that conforms to RFC 7230.
	// Example: `example_value`
	Value *string `mandatory:"true" json:"value"`
}

func (m AddHttpResponseHeaderRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AddHttpResponseHeaderRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAddHttpResponseHeaderRule AddHttpResponseHeaderRule
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeAddHttpResponseHeaderRule
	}{
		"ADD_HTTP_RESPONSE_HEADER",
		(MarshalTypeAddHttpResponseHeaderRule)(m),
	}

	return json.Marshal(&s)
}
