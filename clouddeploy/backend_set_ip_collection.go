// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v32/common"
)

// BackendSetIpCollection Collection of Backend environment IP address.
type BackendSetIpCollection struct {

	// The IP address of the backend server.
	// A server could be a Compute instance or a Load Balancer.
	Items []string `mandatory:"false" json:"items"`
}

func (m BackendSetIpCollection) String() string {
	return common.PointerString(m)
}
