// Copyright (c) 2016 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// File Storage Service API Tests
//
// The API for the File Storage Service.
//

package integtest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oracle/oci-go-sdk/v29/common"

	"github.com/oracle/oci-go-sdk/v29/filestorage"
)

// sanity test for file storage service as the current plan is rely on Java test for other APIs
func TestFileStorageClient_ListFileSystems(t *testing.T) {
	c, clerr := filestorage.NewFileStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	req := filestorage.ListFileSystemsRequest{
		CompartmentId:      common.String(getCompartmentID()),
		AvailabilityDomain: common.String(validAD()),
	}

	resp, err := c.ListFileSystems(context.Background(), req)
	failIfError(t, err)

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.OpcRequestId)
}
