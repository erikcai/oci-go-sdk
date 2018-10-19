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

// RedisVersionSummary Summary object describing a redis version
type RedisVersionSummary struct {

	// Redis version
	Version *string `mandatory:"true" json:"version"`
}

func (m RedisVersionSummary) String() string {
	return common.PointerString(m)
}
