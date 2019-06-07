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

// SourceIpAddressCondition Source IP address based match condition
type SourceIpAddressCondition struct {

	// IPv4 or IPv6 address range to which the source IP address of incoming packet would be matched against
	// Only classless inter-domain routing (CIDR) format(x.x.x.x/y or x:x::x/y) is accepted
	// Specify 0.0.0.0/0 or ::/0 to match all incoming traffic
	AttributeValue *string `mandatory:"true" json:"attributeValue"`
}

func (m SourceIpAddressCondition) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m SourceIpAddressCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSourceIpAddressCondition SourceIpAddressCondition
	s := struct {
		DiscriminatorParam string `json:"attributeName"`
		MarshalTypeSourceIpAddressCondition
	}{
		"SOURCE_IP_ADDRESS",
		(MarshalTypeSourceIpAddressCondition)(m),
	}

	return json.Marshal(&s)
}
