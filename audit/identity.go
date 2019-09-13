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

// Identity The representation of Identity
type Identity struct {

	// The name of the user or service. This value is the friendly name associated with principalId.
	// Example: `JohnSmith`
	PrincipalName *string `mandatory:"false" json:"principalName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the principal.
	PrincipalId *string `mandatory:"false" json:"principalId"`

	AuthType *string `mandatory:"false" json:"authType"`

	// The name of the user or service. This value is the friendly name associated with callerId.
	// Example: `JohnSmith`
	CallerName *string `mandatory:"false" json:"callerName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the caller. The caller that made a request on behalf of the prinicpal.
	CallerId *string `mandatory:"false" json:"callerId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the tenant.
	TenantId *string `mandatory:"false" json:"tenantId"`

	// The IP address of the source of the request.
	// Example: `172.24.96.35`
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The credential ID of the user. This value is extracted from the HTTP 'Authorization' request
	// header. It consists of the tenantId, userId, and user fingerprint, all delimited by a slash (/).
	Credentials *string `mandatory:"false" json:"credentials"`

	// The user agent of the client that made the request.
	UserAgent *string `mandatory:"false" json:"userAgent"`

	// The identifier of the console session
	ConsoleSessionId *string `mandatory:"false" json:"consoleSessionId"`
}

func (m Identity) String() string {
	return common.PointerString(m)
}
