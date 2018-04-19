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

// CreateVolumeAttachmentDetails The representation of CreateVolumeAttachmentDetails
type CreateVolumeAttachmentDetails interface {

	// The OCID of the volume.
	GetVolumeId() *string

	// A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.
	GetDisplayName() *string
}

type createvolumeattachmentdetails struct {
	JsonData    []byte
	VolumeId    *string `mandatory:"true" json:"volumeId"`
	DisplayName *string `mandatory:"false" json:"displayName"`
	Type        string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createvolumeattachmentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatevolumeattachmentdetails createvolumeattachmentdetails
	s := struct {
		Model Unmarshalercreatevolumeattachmentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.VolumeId = s.Model.VolumeId
	m.DisplayName = s.Model.DisplayName
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createvolumeattachmentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.Type {
	case "iscsi":
		mm := CreateIscsiVolumeAttachmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

//GetVolumeId returns VolumeId
func (m createvolumeattachmentdetails) GetVolumeId() *string {
	return m.VolumeId
}

//GetDisplayName returns DisplayName
func (m createvolumeattachmentdetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m createvolumeattachmentdetails) String() string {
	return common.PointerString(m)
}
