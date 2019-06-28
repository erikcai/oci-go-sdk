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

// PaginatedResponseIncidentType The representation of PaginatedResponseIncidentType
type PaginatedResponseIncidentType struct {
	Results []IncidentType `mandatory:"false" json:"results"`

	ResultsType *ModelType `mandatory:"false" json:"resultsType"`

	Next *string `mandatory:"false" json:"next"`

	Prev *string `mandatory:"false" json:"prev"`
}

func (m PaginatedResponseIncidentType) String() string {
	return common.PointerString(m)
}
