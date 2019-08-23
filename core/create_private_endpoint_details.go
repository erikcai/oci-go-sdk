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

// CreatePrivateEndpointDetails Details to create a Private Endpoint.
type CreatePrivateEndpointDetails struct {

	// The OCID  (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment to contain the Private Endpoint
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Endpoint Service OCID. These is the Service which will be enabled on the Private Endpoint.
	EndpointServiceId *string `mandatory:"true" json:"endpointServiceId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// A user-friendly name. Does not have to be unique.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of this particular Private Endpoint.
	Description *string `mandatory:"false" json:"description"`

	// Service's 3 label FQDN representing the Endpoint Service.
	EndpointFqdn *string `mandatory:"false" json:"endpointFqdn"`

	// The network security group OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) that this Private Endpoint is associated with.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// An IP in the subnet that shoudl be assigned to this private endpoint.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`
}

func (m CreatePrivateEndpointDetails) String() string {
	return common.PointerString(m)
}
