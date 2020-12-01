// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Core Services API

package example

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/v30/common"
	"github.com/oracle/oci-go-sdk/v30/core"
	"github.com/oracle/oci-go-sdk/v30/example/helpers"
)

const (
	vnicId      = "" // REQUIRED: VNIC ID on which the IPv6 attachment will be created
	displayName = "OCI-GOSDK-CreateIpv6-Example"
)

func ExampleCreateIpv6() {
	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)
	ctx := context.Background()

	// create the request
	request := core.CreateIpv6Request{}
	request.VnicId = common.String(vnicId)
	request.DisplayName = common.String(displayName)
	request.RequestMetadata = helpers.GetRequestMetadataWithDefaultRetryPolicy()

	_, err = client.CreateIpv6(ctx, request)
	helpers.FatalIfError(err)

	fmt.Println("IPv6 VNIC attachment created")

	// Output:
	// IPv6 VNIC attachment created
}
