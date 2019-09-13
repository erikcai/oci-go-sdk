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

// AuditEvent All the attributes of an audit event. For more information, see Viewing Audit Log Events (https://docs.cloud.oracle.com/iaas/Content/Audit/Tasks/viewinglogevents.htm).
type AuditEvent struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the tenant.
	TenantId *string `mandatory:"false" json:"tenantId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The name of the compartment. This value is the friendly name associated with compartmentId.
	// This value can change, but the service logs the value that appeared at the time of the audit
	// event.
	// Example: `CompartmentA`
	CompartmentName *string `mandatory:"false" json:"compartmentName"`

	// The GUID of the event.
	// Example: `examplea9-f488-4842-96cb-a10f2893b369`
	EventId *string `mandatory:"false" json:"eventId"`

	// The name of the event.
	// Example: `LaunchInstance`
	EventName *string `mandatory:"false" json:"eventName"`

	// The source of the event.
	// Example: `TestService`
	EventSource *string `mandatory:"false" json:"eventSource"`

	// The type of event.
	// Example: `ServiceApi`
	EventType *string `mandatory:"false" json:"eventType"`

	// The time the event occurred, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-06-14T22:24:37.671Z`
	EventTime *common.SDKTime `mandatory:"false" json:"eventTime"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the user whose action triggered
	// the event.
	PrincipalId *string `mandatory:"false" json:"principalId"`

	// The credential ID of the user. This value is extracted from the HTTP 'Authorization' request
	// header. It consists of the tenantId, userId, and user fingerprint, all delimited by a slash (/).
	CredentialId *string `mandatory:"false" json:"credentialId"`

	// The HTTP method of the request.
	// Example: `GET`
	RequestAction *string `mandatory:"false" json:"requestAction"`

	// The opc-request-id of the request.
	RequestId *string `mandatory:"false" json:"requestId"`

	// The user agent of the client that made the request.
	RequestAgent *string `mandatory:"false" json:"requestAgent"`

	// The HTTP header fields and values in the request.
	// Example:
	//   -----
	//     {
	//       <The HTTP header fields and values from the request.>
	//     }
	//   -----
	RequestHeaders map[string][]string `mandatory:"false" json:"requestHeaders"`

	// The IP address of the source of the request.
	// Example: `172.24.96.35`
	RequestOrigin *string `mandatory:"false" json:"requestOrigin"`

	// The query parameter fields and values for the request.
	// Example:
	//   -----
	//     {
	//       "compartmentId": [
	//         "example.compartment.region1..aaaaaaaaexample5gsmzi3hnumfjzstsgkzha74uc6jzpcxlxhx33hewsq"
	//         ],
	//       "availabilityDomain": [
	//         "Example-AD-2"
	//         ],
	//       "sortOrder": [
	//         "DESC"
	//         ],
	//       "limit": [
	//         "25"
	//         ],
	//       "sortBy": [
	//         "timeCreated"
	//         ]
	//     }
	//   -----
	RequestParameters map[string][]string `mandatory:"false" json:"requestParameters"`

	// The resource targeted by the request.
	// Example: `/hello-world`
	RequestResource *string `mandatory:"false" json:"requestResource"`

	// The headers of the response.
	// Example:
	//   -----
	//     {
	//       "Access-Control-Expose-Headers": [
	//         "opc-previous-page,opc-next-page,opc-client-info,ETag,opc-request-id,Location"
	//         ],
	//       "Access-Control-Allow-Origin": [
	//         "https://console.us-phoenix-1.oraclecloud.com"
	//         ],
	//       "Access-Control-Allow-Credentials": [
	//         "true"
	//         ],
	//       "Connection": [
	//         "close"
	//         ],
	//       "Content-Length": [
	//         "3"
	//         ],
	//       "opc-request-id": [
	//         "example-4092-8233-EXAMPLEB863DC6CEC1BF1CE734676108C6345FF/51FE3CACE106DD8F825508D04E91E261"
	//         ],
	//       "Date": [
	//         "Thu, 14 Jun 2018 22:24:37 GMT"
	//         ],
	//       "Content-Type": [
	//         "application/json"
	//         ]
	//     }
	//   -----
	ResponseHeaders map[string][]string `mandatory:"false" json:"responseHeaders"`

	// The status code of the response.
	// Example: `200`
	ResponseStatus *string `mandatory:"false" json:"responseStatus"`

	// The time of the response to the audited request, expressed in
	// RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-06-14T22:24:37.713Z`
	ResponseTime *common.SDKTime `mandatory:"false" json:"responseTime"`

	// Metadata of interest from the response payload. For example, the
	// OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of a resource.
	ResponsePayload map[string]interface{} `mandatory:"false" json:"responsePayload"`

	// The name of the user or service. This value is the friendly name associated with principalId.
	// Example: `JohnSmith`
	UserName *string `mandatory:"false" json:"userName"`
}

func (m AuditEvent) String() string {
	return common.PointerString(m)
}
