// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"github.com/oracle/oci-go-sdk/v32/common"
)

// AccessRequestSummary Summary of access request.
type AccessRequestSummary struct {

	// The OCID of the access request.
	Id *string `mandatory:"true" json:"id"`

	// Comment associated with the access request.
	AccessReasonSummary *string `mandatory:"true" json:"accessReasonSummary"`

	// The OCID of the target resource associated with the access request. The operator raises an access request to get approval to
	// access the target resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// This is a system-generated identifier.
	RequestId *string `mandatory:"false" json:"requestId"`

	// The OCID of the compartment that contains the access request.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The name of the target resource.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The current state of the AccessRequest.
	LifecycleState AccessRequestLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Time when the access request was created by the operator user in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeOfCreation *common.SDKTime `mandatory:"false" json:"timeOfCreation"`

	// Time when the access request was last modified in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeOfModification *common.SDKTime `mandatory:"false" json:"timeOfModification"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AccessRequestSummary) String() string {
	return common.PointerString(m)
}
