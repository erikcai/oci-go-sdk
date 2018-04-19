// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ImageSourceViaObjectStorageTupleDetails The representation of ImageSourceViaObjectStorageTupleDetails
type ImageSourceViaObjectStorageTupleDetails struct {
	OperatingSystem *string `mandatory:"true" json:"operatingSystem"`

	OperatingSystemVersion *string `mandatory:"true" json:"operatingSystemVersion"`

	// Object store bucket to export the image to
	BucketName *string `mandatory:"false" json:"bucketName"`

	// Object store namespace to export the image to
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// Object store object name for the exported image
	ObjectName *string `mandatory:"false" json:"objectName"`
}

//GetOperatingSystem returns OperatingSystem
func (m ImageSourceViaObjectStorageTupleDetails) GetOperatingSystem() *string {
	return m.OperatingSystem
}

//GetOperatingSystemVersion returns OperatingSystemVersion
func (m ImageSourceViaObjectStorageTupleDetails) GetOperatingSystemVersion() *string {
	return m.OperatingSystemVersion
}

func (m ImageSourceViaObjectStorageTupleDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ImageSourceViaObjectStorageTupleDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImageSourceViaObjectStorageTupleDetails ImageSourceViaObjectStorageTupleDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeImageSourceViaObjectStorageTupleDetails
	}{
		"objectStorageTuple",
		(MarshalTypeImageSourceViaObjectStorageTupleDetails)(m),
	}

	return json.Marshal(&s)
}
