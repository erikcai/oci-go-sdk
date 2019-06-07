// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// SourceVcnIpAddressCondition Source VCN/Overlay IP address based match condition. Should be used always in conjunction with SourceVcnIdCondition
type SourceVcnIpAddressCondition struct {

	// IPv4 address range to which the original client IP address (in customer VCN) of incoming packet would be matched against
	// Only classless inter-domain routing (CIDR) format(x.x.x.x/y or x:x::x/y) is accepted
	// Specify 0.0.0.0/0 or ::/0 to match all incoming traffic in the customer VCN
	AttributeValue *string `mandatory:"true" json:"attributeValue"`
}

func (m SourceVcnIpAddressCondition) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m SourceVcnIpAddressCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSourceVcnIpAddressCondition SourceVcnIpAddressCondition
	s := struct {
		DiscriminatorParam string `json:"attributeName"`
		MarshalTypeSourceVcnIpAddressCondition
	}{
		"SOURCE_VCN_IP_ADDRESS",
		(MarshalTypeSourceVcnIpAddressCondition)(m),
	}

	return json.Marshal(&s)
}
