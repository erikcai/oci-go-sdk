// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OraCache Public API
//
// API for the Data Caching Service. Use this service to manage Redis replicated caches.
//

package cache

import (
	"github.com/oracle/oci-go-sdk/v28/common"
)

// RedisShapeSummary The amount of memory allocated to the Redis replicated cache.
type RedisShapeSummary struct {

	// Redis shape
	Shape *string `mandatory:"true" json:"shape"`
}

func (m RedisShapeSummary) String() string {
	return common.PointerString(m)
}
