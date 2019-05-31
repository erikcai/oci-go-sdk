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

// GatewaySummary Summary of the Gateway.
type GatewaySummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which
	// resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Gateway Endpoint Type
	EndpointType GatewayEndpointTypeEnum `mandatory:"true" json:"endpointType"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet in which
	// related resources are created.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The time this resource was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the gateway.
	LifecycleState GatewayLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// An message describing the current state in more detail.
	// For example, can be used to provide actionable information for a
	// resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The hostname for the APIs deployed on the Gateway
	Hostname *string `mandatory:"false" json:"hostname"`
}

func (m GatewaySummary) String() string {
	return common.PointerString(m)
}
