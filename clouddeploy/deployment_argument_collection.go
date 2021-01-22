// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// DeploymentArgumentCollection Specifies list of arguments passed along with the deployment.
type DeploymentArgumentCollection struct {

	// List of arguments provided at the time of deployment.
	Items []DeploymentArgument `mandatory:"true" json:"items"`
}

func (m DeploymentArgumentCollection) String() string {
	return common.PointerString(m)
}
