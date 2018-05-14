// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object and Archive Storage APIs for managing buckets and objects.
//

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ObjectLifecycleRule To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
type ObjectLifecycleRule struct {

	// The name of the lifecycle rule to be applied.
	Name *string `mandatory:"true" json:"name"`

	// The action of the object lifecycle policy rule. Rules using the action 'ARCHIVE' move objects into the
	// Archival Storage tier (https://docs.us-phoenix-1.oraclecloud.com/Content/Archive/Concepts/archivestorageoverview.htm). Rules using the action
	// 'DELETE' permanently delete objects from buckets.
	Action ObjectLifecycleRuleActionEnum `mandatory:"true" json:"action"`

	// Specifies the age (in days) of objects to apply the rule to. An object's daysSinceLastModified value is based on
	// its Last-Modified time. Days are defined as starting and ending at midnight UTC.
	TimeSinceLastModificationInDays *int `mandatory:"true" json:"timeSinceLastModificationInDays"`

	// A boolean that determines whether this rule is currently enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// A filter limiting object names that the rule will apply to.
	ObjectNameFilter *ObjectNameFilter `mandatory:"false" json:"objectNameFilter"`
}

func (m ObjectLifecycleRule) String() string {
	return common.PointerString(m)
}

// ObjectLifecycleRuleActionEnum Enum with underlying type: string
type ObjectLifecycleRuleActionEnum string

// Set of constants representing the allowable values for ObjectLifecycleRuleAction
const (
	ObjectLifecycleRuleActionArchive ObjectLifecycleRuleActionEnum = "ARCHIVE"
	ObjectLifecycleRuleActionDelete  ObjectLifecycleRuleActionEnum = "DELETE"
)

var mappingObjectLifecycleRuleAction = map[string]ObjectLifecycleRuleActionEnum{
	"ARCHIVE": ObjectLifecycleRuleActionArchive,
	"DELETE":  ObjectLifecycleRuleActionDelete,
}

// GetObjectLifecycleRuleActionEnumValues Enumerates the set of values for ObjectLifecycleRuleAction
func GetObjectLifecycleRuleActionEnumValues() []ObjectLifecycleRuleActionEnum {
	values := make([]ObjectLifecycleRuleActionEnum, 0)
	for _, v := range mappingObjectLifecycleRuleAction {
		values = append(values, v)
	}
	return values
}
