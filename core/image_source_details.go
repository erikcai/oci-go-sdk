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

// ImageSourceDetails The representation of ImageSourceDetails
type ImageSourceDetails interface {
	GetOperatingSystem() *string

	GetOperatingSystemVersion() *string
}

type imagesourcedetails struct {
	JsonData               []byte
	OperatingSystem        *string `mandatory:"true" json:"operatingSystem"`
	OperatingSystemVersion *string `mandatory:"true" json:"operatingSystemVersion"`
	SourceType             string  `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *imagesourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerimagesourcedetails imagesourcedetails
	s := struct {
		Model Unmarshalerimagesourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.OperatingSystem = s.Model.OperatingSystem
	m.OperatingSystemVersion = s.Model.OperatingSystemVersion
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *imagesourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.SourceType {
	case "objectStorageTuple":
		mm := ImageSourceViaObjectStorageTupleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "objectStorageUri":
		mm := ImageSourceViaObjectStorageUriDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

//GetOperatingSystem returns OperatingSystem
func (m imagesourcedetails) GetOperatingSystem() *string {
	return m.OperatingSystem
}

//GetOperatingSystemVersion returns OperatingSystemVersion
func (m imagesourcedetails) GetOperatingSystemVersion() *string {
	return m.OperatingSystemVersion
}

func (m imagesourcedetails) String() string {
	return common.PointerString(m)
}
