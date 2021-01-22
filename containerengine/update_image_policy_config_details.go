// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// UpdateImagePolicyConfigDetails The properties that define a image verification policy.
type UpdateImagePolicyConfigDetails struct {

	// Whether the image verification policy is enabled. Defaults to false. If set to true, the images will be verified against the policy at runtime.
	IsPolicyEnabled *bool `mandatory:"false" json:"isPolicyEnabled"`

	// The OCIDs of the KMS keys that will be used to verify whether the images are signed by an approved source.
	KmsKeyIds []string `mandatory:"false" json:"kmsKeyIds"`
}

func (m UpdateImagePolicyConfigDetails) String() string {
	return common.PointerString(m)
}
