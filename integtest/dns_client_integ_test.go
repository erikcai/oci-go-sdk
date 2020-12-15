// Copyright (c) 2016 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
//
// Public DNS Service Tests
//
// API for managing DNS zones, records, and policies.
//

package integtest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oracle/oci-go-sdk/v31/common"

	"github.com/oracle/oci-go-sdk/v31/dns"
)

// sanity test for dns service as the current plan is rely on Java test for other APIs
func TestDNSClient_ListZones(t *testing.T) {
	c, clerr := dns.NewDnsClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	req := dns.ListZonesRequest{
		CompartmentId: common.String(getCompartmentID()),
	}

	resp, err := c.ListZones(context.Background(), req)
	failIfError(t, err)

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.OpcRequestId)
}
