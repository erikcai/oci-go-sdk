// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science APIs to organize your data science work, access data and computing resources, and build, train, deploy, and manage models on Oracle Cloud.
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// ModelDeploymentResourcePrincipalToken This is a JsonWebToken (JWT) containing the mapping information between the instance principal and model deployment resource.
type ModelDeploymentResourcePrincipalToken struct {

	// The base64 encoded intermediate Resource Principal Token.
	ResourcePrincipalToken *string `mandatory:"true" json:"resourcePrincipalToken"`

	// The base64 encoded Session Token of the Service Principal that provided the intermediate Resource Principal Token.
	ServicePrincipalSessionToken *string `mandatory:"true" json:"servicePrincipalSessionToken"`
}

func (m ModelDeploymentResourcePrincipalToken) String() string {
	return common.PointerString(m)
}
