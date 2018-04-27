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

// RemoveHttpResponseHeaderRule An object that represents the action of removing a header from a response.
// If same header presented more than once then rule will be applie to all occurances of the header.
// This rule will be applied only to HTTP or HTTP2 listeners.
type RemoveHttpResponseHeaderRule struct {

	// A header name that conforms to RFC 7230.
	// Example: `example_header_name`
	Header *string `mandatory:"true" json:"header"`
}

func (m RemoveHttpResponseHeaderRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RemoveHttpResponseHeaderRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRemoveHttpResponseHeaderRule RemoveHttpResponseHeaderRule
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeRemoveHttpResponseHeaderRule
	}{
		"REMOVE_HTTP_RESPONSE_HEADER",
		(MarshalTypeRemoveHttpResponseHeaderRule)(m),
	}

	return json.Marshal(&s)
}
