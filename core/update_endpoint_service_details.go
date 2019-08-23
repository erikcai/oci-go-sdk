// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateEndpointServiceDetails Information that can be updated for a particular Endpoint Service.
type UpdateEndpointServiceDetails struct {

	// Description of this particular Endpoint Service, provided by the service owner.
	Description *string `mandatory:"false" json:"description"`

	// Name of the Endpoint Service.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Allow multiple Private Endpoints to be created for this Endpoint Service in the same customer VCN. Defaults to false.
	AreMultiplePrivateEndpointsPerVcnAllowed *bool `mandatory:"false" json:"areMultiplePrivateEndpointsPerVcnAllowed"`

	// Indicates if the incoming traffic should include VCN metadata of the source.
	IsVcnMetadataEnabled *bool `mandatory:"false" json:"isVcnMetadataEnabled"`

	// List of unique service IPs that will service the requests. If you are directly providing service IPs, then you will also have to indicate if these are public IPs.
	ServiceIps []EndpointServiceIpDetails `mandatory:"false" json:"serviceIps"`

	// Ports that are open on the provided service IPs for the Endpoint Service.
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
