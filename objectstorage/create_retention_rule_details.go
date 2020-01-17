// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object Storage and Archive Storage APIs for managing buckets, objects, and related resources.
//

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateRetentionRuleDetails The details to create a retention rule.
type CreateRetentionRuleDetails struct {

	// A user-specified name for the retention rule. Names can be helpful in identifying retention rules.
	DisplayName *string `mandatory:"false" json:"displayName"`

	Duration *Duration `mandatory:"false" json:"duration"`

	// The date and time as per RFC 3339 (https://tools.ietf.org/html/rfc3339) after which this rule is locked
	// and can only be deleted by deleting the bucket. Once a rule is locked, only increases in the duration are
	// allowed and no other properties can be changed. This property cannot be updated for rules that are in a
	// locked state. Specifying it when a duration is not specified is considered an error.
	TimeRuleLocked *common.SDKTime `mandatory:"false" json:"timeRuleLocked"`
}

func (m CreateRetentionRuleDetails) String() string {
	return common.PointerString(m)
}
