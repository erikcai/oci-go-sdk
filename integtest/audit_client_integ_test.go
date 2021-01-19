// Copyright (c) 2016, 2017, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package integtest

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/oracle/oci-go-sdk/v33/audit"
	"github.com/oracle/oci-go-sdk/v33/common"
)

// ListEvents test
func TestAuditClient_ListEvents(t *testing.T) {
	c, clerr := audit.NewAuditClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	// list events for last 5 hour
	req := audit.ListEventsRequest{
		CompartmentId: common.String(getCompartmentID()),
		StartTime:     &common.SDKTime{time.Now().Add(time.Hour * -5)},
		EndTime:       &common.SDKTime{time.Now()},
	}

	resp, err := c.ListEvents(context.Background(), req)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
}
