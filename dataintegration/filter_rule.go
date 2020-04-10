// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// FilterRule auto generated description
type FilterRule struct {

	// rule
	Rule *string `mandatory:"false" json:"rule"`

	// ruleType
	RuleType *string `mandatory:"false" json:"ruleType"`
}

func (m FilterRule) String() string {
	return common.PointerString(m)
}
