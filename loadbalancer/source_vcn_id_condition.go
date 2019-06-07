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

// SourceVcnIdCondition Source VCN OCID based match condition
type SourceVcnIdCondition struct {

	// OCID of the customer VCN to which the service gateway embedded VCN ID of incoming packet would be matched against
	// This condition can be used in conjunction with `SourceVcnIpAddressCondition`.
	// **NOTE:** If this condition is defined on a rule without `SourceVcnIpAddressCondition`, then this condition matches all incoming traffic in the specified customer VCN
	AttributeValue *string `mandatory:"true" json:"attributeValue"`
}

func (m SourceVcnIdCondition) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m SourceVcnIdCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSourceVcnIdCondition SourceVcnIdCondition
	s := struct {
		DiscriminatorParam string `json:"attributeName"`
		MarshalTypeSourceVcnIdCondition
	}{
		"SOURCE_VCN_ID",
		(MarshalTypeSourceVcnIdCondition)(m),
	}

	return json.Marshal(&s)
}
