// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
//

package apigateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateGatewayDetails The information about new Gateway.
type CreateGatewayDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which
	// resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Gateway endpoint type
	EndpointType GatewayEndpointTypeEnum `mandatory:"true" json:"endpointType"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet in which
	// related resources are created.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	// with no predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateGatewayDetails) String() string {
	return common.PointerString(m)
}
