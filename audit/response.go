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

// Response The representation of Response
type Response struct {

	// The status code of the response.
	// Example: `200`
	Status *string `mandatory:"false" json:"status"`

	// The time of the response to the audited request, expressed in
	// RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-06-14T22:24:37.713Z`
	ResponseTime *common.SDKTime `mandatory:"false" json:"responseTime"`

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
	Headers map[string][]string `mandatory:"false" json:"headers"`

	// Metadata of interest from the response payload.
	Payload map[string]interface{} `mandatory:"false" json:"payload"`

	Message *string `mandatory:"false" json:"message"`
}

func (m Response) String() string {
	return common.PointerString(m)
}
