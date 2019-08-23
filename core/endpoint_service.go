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

// EndpointService Information of a particular Endpoint Service that is created for the provider service.
type EndpointService struct {

	// The Endpoint Service's Oracle ID (OCID) (/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// The OCID  (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)  of the VCN to contain the Endpoint service.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID  (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)  of the compartment to contain the Endpoint service.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// List of service IPs that will expose the service.
	ServiceIps []EndpointServiceIpDetails `mandatory:"true" json:"serviceIps"`

	// The ports on this service's IPs that are open for communication via Private Endpoints associated to this Endpoint Service.
	// If no ports are provided, all open ports are accessible on your service IPs
	Ports []int `mandatory:"true" json:"ports"`

	// Service's 3 label FQDN representing the Endpoint Service. This is the fqdn which will be updated in consumer VCN's dns whenever a Private Endpoint is created.
	EndpointFqdn *string `mandatory:"true" json:"endpointFqdn"`

	// The date and time the Service was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The endpoint service's current state.
	LifecycleState EndpointServiceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Description of this particular Endpoint Service, provided by the service owner.
	Description *string `mandatory:"false" json:"description"`

	// Name of the Endpoint Service. This has to be unique within a VCN.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Allow multiple Private Endpoints to be created for this Endpoint Service in the same customer VCN. Defaults to false.
	AreMultiplePrivateEndpointsPerVcnAllowed *bool `mandatory:"false" json:"areMultiplePrivateEndpointsPerVcnAllowed"`

	// Indicates if the incoming traffic should include VCN metadata of the source.
	IsVcnMetadataEnabled *bool `mandatory:"false" json:"isVcnMetadataEnabled"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m EndpointService) String() string {
	return common.PointerString(m)
}

// EndpointServiceLifecycleStateEnum Enum with underlying type: string
type EndpointServiceLifecycleStateEnum string

// Set of constants representing the allowable values for EndpointServiceLifecycleStateEnum
const (
	EndpointServiceLifecycleStateProvisioning EndpointServiceLifecycleStateEnum = "PROVISIONING"
	EndpointServiceLifecycleStateAvailable    EndpointServiceLifecycleStateEnum = "AVAILABLE"
	EndpointServiceLifecycleStateTerminating  EndpointServiceLifecycleStateEnum = "TERMINATING"
	EndpointServiceLifecycleStateTerminated   EndpointServiceLifecycleStateEnum = "TERMINATED"
)

var mappingEndpointServiceLifecycleState = map[string]EndpointServiceLifecycleStateEnum{
	"PROVISIONING": EndpointServiceLifecycleStateProvisioning,
	"AVAILABLE":    EndpointServiceLifecycleStateAvailable,
	"TERMINATING":  EndpointServiceLifecycleStateTerminating,
	"TERMINATED":   EndpointServiceLifecycleStateTerminated,
}

// GetEndpointServiceLifecycleStateEnumValues Enumerates the set of values for EndpointServiceLifecycleStateEnum
func GetEndpointServiceLifecycleStateEnumValues() []EndpointServiceLifecycleStateEnum {
	values := make([]EndpointServiceLifecycleStateEnum, 0)
	for _, v := range mappingEndpointServiceLifecycleState {
		values = append(values, v)
	}
	return values
}
