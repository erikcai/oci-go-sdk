// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
//
// Cloud Incident Management Service
//

package cims

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// LimitItem The representation of LimitItem
type LimitItem struct {
	Id *string `mandatory:"false" json:"id"`

	Name *string `mandatory:"false" json:"name"`

	CurrentLimit *int `mandatory:"false" json:"currentLimit"`

	CurrentUsage *int `mandatory:"false" json:"currentUsage"`

	RequestedLimit *int `mandatory:"false" json:"requestedLimit"`
}

//GetId returns Id
func (m LimitItem) GetId() *string {
	return m.Id
}

//GetName returns Name
func (m LimitItem) GetName() *string {
	return m.Name
}

func (m LimitItem) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m LimitItem) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLimitItem LimitItem
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeLimitItem
	}{
		"limit",
		(MarshalTypeLimitItem)(m),
	}

	return json.Marshal(&s)
}
