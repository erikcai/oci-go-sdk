// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Cpe An object you create when setting up an IPSec VPN between your on-premise network
// and VCN. The `Cpe` is a virtual representation of your Customer-Premises Equipment,
// which is the actual router on-premise at your site at your end of the IPSec VPN connection.
// For more information,
// see Overview of the Networking Service (https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/overview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
type Cpe struct {

	// The OCID of the compartment containing the CPE.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The CPE's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The public IP address of the on-premise router.
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The date and time the CPE was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m Cpe) String() string {
	return common.PointerString(m)
}
