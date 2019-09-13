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

// Request The representation of Request
type Request struct {

	// The opc-request-id of the request.
	Id *string `mandatory:"false" json:"id"`

	Path *string `mandatory:"false" json:"path"`

	// The HTTP method of the request.
	// Example: `GET`
	Action *string `mandatory:"false" json:"action"`

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
	Parameters map[string][]string `mandatory:"false" json:"parameters"`

	// The HTTP header fields and values in the request.
	// Example:
	//   -----
	//     {
	//       <The HTTP header fields and values from the request.>
	//     }
	//   -----
	Headers map[string][]string `mandatory:"false" json:"headers"`
}

func (m Request) String() string {
	return common.PointerString(m)
}
