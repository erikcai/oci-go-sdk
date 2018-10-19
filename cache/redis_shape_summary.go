// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OraCache Public API
//
// Oracle Caching Service Public API
//

package cache

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RedisShapeSummary Summary object describing a redis shape
type RedisShapeSummary struct {

	// Redis shape
	Shape *string `mandatory:"true" json:"shape"`
}

func (m RedisShapeSummary) String() string {
	return common.PointerString(m)
}
