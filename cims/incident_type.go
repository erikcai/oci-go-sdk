// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
//
// Cloud Incident Management Service
//

package cims

import (
	"github.com/oracle/oci-go-sdk/common"
)

// IncidentType Incident type
type IncidentType struct {
	Id *string `mandatory:"false" json:"id"`

	Name *string `mandatory:"false" json:"name"`

	Label *string `mandatory:"false" json:"label"`

	Description *string `mandatory:"false" json:"description"`

	ClassifierList []Classifier `mandatory:"false" json:"classifierList"`
}

func (m IncidentType) String() string {
	return common.PointerString(m)
}
