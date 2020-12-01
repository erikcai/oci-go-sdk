// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm),
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm), and
// Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as virtual cloud networks (VCNs),
// compute instances, block storage volumes, and container images.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// UpdateEndpointServiceDetails Information that can be updated for an endpoint service.
type UpdateEndpointServiceDetails struct {

	// A description of the endpoint service.
	Description *string `mandatory:"false" json:"description"`

	// A friendly name for the endpoint service. Must be unique within the VCN. Avoid entering
	// confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Some services want to restrict access to the resources represented by an endpoint service so
	// that only a single private endpoint in the customer VCN has access.
	// For example, the endpoint service might represent a particular service resource (such as a
	// particular database). The service might want to allow access to that particular resource
	// from only a single private endpoint.
	// Defaults to `false`.
	// Example: `true`
	AreMultiplePrivateEndpointsPerVcnAllowed *bool `mandatory:"false" json:"areMultiplePrivateEndpointsPerVcnAllowed"`

	// List of service endpoints (in the service VCN) that handle requests to the endpoint service.
	ServiceIps []EndpointServiceIpDetails `mandatory:"false" json:"serviceIps"`

	// The ports on the endpoint service's endpoints that are open for private endpoint traffic for this
	// endpoint service. If you provide no ports, all open ports on the service endpoints are accessible.
	Ports []int `mandatory:"false" json:"ports"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m UpdateEndpointServiceDetails) String() string {
	return common.PointerString(m)
}
