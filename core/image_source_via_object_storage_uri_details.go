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

// ImageSourceViaObjectStorageUriDetails The representation of ImageSourceViaObjectStorageUriDetails
type ImageSourceViaObjectStorageUriDetails struct {
	OperatingSystem *string `mandatory:"true" json:"operatingSystem"`

	OperatingSystemVersion *string `mandatory:"true" json:"operatingSystemVersion"`

	// Object store URI to export the image to
	SourceUri *string `mandatory:"true" json:"sourceUri"`
}

//GetOperatingSystem returns OperatingSystem
func (m ImageSourceViaObjectStorageUriDetails) GetOperatingSystem() *string {
	return m.OperatingSystem
}

//GetOperatingSystemVersion returns OperatingSystemVersion
func (m ImageSourceViaObjectStorageUriDetails) GetOperatingSystemVersion() *string {
	return m.OperatingSystemVersion
}

func (m ImageSourceViaObjectStorageUriDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ImageSourceViaObjectStorageUriDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImageSourceViaObjectStorageUriDetails ImageSourceViaObjectStorageUriDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeImageSourceViaObjectStorageUriDetails
	}{
		"objectStorageUri",
		(MarshalTypeImageSourceViaObjectStorageUriDetails)(m),
	}

	return json.Marshal(&s)
}
