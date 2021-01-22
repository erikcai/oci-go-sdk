// Copyright (c) 2016, 2017, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Clusters Service API
//
// Test for Container Engine for Kubernetes API

package integtest

import (
	"context"
	"testing"

	"github.com/erikcai/oci-go-sdk/v33/common"
	"github.com/erikcai/oci-go-sdk/v33/containerengine"
	"github.com/stretchr/testify/assert"
)

// ListClusters test -- simple test to make sure it can reach to container engine service API
func TestContainerEngineClient_ListClusters(t *testing.T) {
	c, clerr := containerengine.NewContainerEngineClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	// list events for last 5 hour
	req := containerengine.ListClustersRequest{
		CompartmentId: common.String(getCompartmentID()),
	}

	resp, err := c.ListClusters(context.Background(), req)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
}
