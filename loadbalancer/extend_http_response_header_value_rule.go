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

// ExtendHttpResponseHeaderValueRule An object that represents the action of modifying a response header value.
// This rule will be applied only to HTTP or HTTP2 listeners.
// It will concatenate prefix and suffix to header value.
// If same header presented more than once then rule will be applie to only one value
// which will be chosen randomly.
// For example:
// my_suffix_header: ValueOne, ValueTwo, ValueThree
// OR
// my_suffix_header: ValueOne
// my_suffix_header: ValueTwo
// my_suffix_header: ValueThree
// will result in LoadBalancer modifying any one of the values and leaving the rest intact.
type ExtendHttpResponseHeaderValueRule struct {

	// A header name that conforms to RFC 7230.
	// Example: `example_header_name`
	Header *string `mandatory:"true" json:"header"`

	// A string to prepend to the header value. The resulting header value must still conform to RFC 7230.
	// Example: `example_prefix_value`
	Prefix *string `mandatory:"false" json:"prefix"`

	// A string to append to the header value. The resulting header value must still conform to RFC 7230.
	// Example: `example_suffix_value`
	Suffix *string `mandatory:"false" json:"suffix"`
}

func (m ExtendHttpResponseHeaderValueRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ExtendHttpResponseHeaderValueRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExtendHttpResponseHeaderValueRule ExtendHttpResponseHeaderValueRule
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeExtendHttpResponseHeaderValueRule
	}{
		"EXTEND_HTTP_RESPONSE_HEADER_VALUE",
		(MarshalTypeExtendHttpResponseHeaderValueRule)(m),
	}

	return json.Marshal(&s)
}
