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

// Classifier Incident Classifier
type Classifier struct {
	Id *string `mandatory:"false" json:"id"`

	Name *string `mandatory:"false" json:"name"`

	Label *string `mandatory:"false" json:"label"`

	Description *string `mandatory:"false" json:"description"`

	Scope *Scope `mandatory:"false" json:"scope"`

	Unit *Unit `mandatory:"false" json:"unit"`
}

func (m Classifier) String() string {
	return common.PointerString(m)
}
