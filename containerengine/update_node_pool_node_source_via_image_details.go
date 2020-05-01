// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateNodePoolNodeSourceViaImageDetails Details of the image running on the node.
type UpdateNodePoolNodeSourceViaImageDetails struct {

	// The OCID of the image.
	ImageId *string `mandatory:"false" json:"imageId"`

	// The size of the boot volume in GBs. Minimum value is 50 GB. See here (https://docs.cloud.oracle.com/en-us/iaas/Content/Block/Concepts/bootvolumes.htm) for max custom boot volume sizing and OS-specific requirements.
	BootVolumeSizeInGBs *int64 `mandatory:"false" json:"bootVolumeSizeInGBs"`
}

func (m UpdateNodePoolNodeSourceViaImageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateNodePoolNodeSourceViaImageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateNodePoolNodeSourceViaImageDetails UpdateNodePoolNodeSourceViaImageDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeUpdateNodePoolNodeSourceViaImageDetails
	}{
		"IMAGE",
		(MarshalTypeUpdateNodePoolNodeSourceViaImageDetails)(m),
	}

	return json.Marshal(&s)
}
