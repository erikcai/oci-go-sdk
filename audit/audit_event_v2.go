// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Audit API
//
// API for the Audit Service. Use this API for compliance monitoring in your tenancy.
// For more information, see Overview of Audit (https://docs.cloud.oracle.com/iaas/Content/Audit/Concepts/auditoverview.htm).
// **Tip**: This API is good for queries, but not bulk-export operations.
//

package audit

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AuditEventV2 All the attributes of an audit event. For more information, see Viewing Audit Log Events (https://docs.cloud.oracle.com/iaas/Content/Audit/Tasks/viewinglogevents.htm).
type AuditEventV2 struct {
	EventType *string `mandatory:"false" json:"eventType"`

	CloudEventsVersion *string `mandatory:"false" json:"cloudEventsVersion"`

	EventTypeVersion *string `mandatory:"false" json:"eventTypeVersion"`

	// The source of the event.
	// Example: `TestService`
	Source *string `mandatory:"false" json:"source"`

	// The GUID of the event.
	// Example: `examplea9-f488-4842-96cb-a10f2893b369`
	EventId *string `mandatory:"false" json:"eventId"`

	// The time the event occurred, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-06-14T22:24:37.671Z`
	EventTime *common.SDKTime `mandatory:"false" json:"eventTime"`

	ContentType *string `mandatory:"false" json:"contentType"`

	Data *Data `mandatory:"false" json:"data"`
}

func (m AuditEventV2) String() string {
	return common.PointerString(m)
}
