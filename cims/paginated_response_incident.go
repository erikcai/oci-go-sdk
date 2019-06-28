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

// PaginatedResponseIncident The representation of PaginatedResponseIncident
type PaginatedResponseIncident struct {
	Results []Incident `mandatory:"false" json:"results"`

	ResultsType *ModelType `mandatory:"false" json:"resultsType"`

	Next *string `mandatory:"false" json:"next"`

	Prev *string `mandatory:"false" json:"prev"`
}

func (m PaginatedResponseIncident) String() string {
	return common.PointerString(m)
}
