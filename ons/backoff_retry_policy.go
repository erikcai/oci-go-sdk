// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// ONS Gateway API
//
// A description of the ONS Gateway API
//

package ons

import (
	"github.com/oracle/oci-go-sdk/common"
)

// BackoffRetryPolicy Backoff retry policy
type BackoffRetryPolicy struct {

	// The maximum retry duration in milliseconds
	MaxRetryDuration *int `mandatory:"false" json:"maxRetryDuration"`

	// Deliver policy in a json string
	PolicyType BackoffRetryPolicyPolicyTypeEnum `mandatory:"false" json:"policyType,omitempty"`
}

func (m BackoffRetryPolicy) String() string {
	return common.PointerString(m)
}

// BackoffRetryPolicyPolicyTypeEnum Enum with underlying type: string
type BackoffRetryPolicyPolicyTypeEnum string

// Set of constants representing the allowable values for BackoffRetryPolicyPolicyTypeEnum
const (
	BackoffRetryPolicyPolicyTypeExponential BackoffRetryPolicyPolicyTypeEnum = "EXPONENTIAL"
)

var mappingBackoffRetryPolicyPolicyType = map[string]BackoffRetryPolicyPolicyTypeEnum{
	"EXPONENTIAL": BackoffRetryPolicyPolicyTypeExponential,
}

// GetBackoffRetryPolicyPolicyTypeEnumValues Enumerates the set of values for BackoffRetryPolicyPolicyTypeEnum
func GetBackoffRetryPolicyPolicyTypeEnumValues() []BackoffRetryPolicyPolicyTypeEnum {
	values := make([]BackoffRetryPolicyPolicyTypeEnum, 0)
	for _, v := range mappingBackoffRetryPolicyPolicyType {
		values = append(values, v)
	}
	return values
}
