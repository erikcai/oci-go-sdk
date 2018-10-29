// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace REST API endpoint
//
// Marketplace REST API for loom plugin
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ActionProperty The model of user action property for install configurations
type ActionProperty struct {

	// Name of the action property
	Name *string `mandatory:"false" json:"name"`

	// Display Name of the action property
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Resource tag of the action property
	ResourceTag *string `mandatory:"false" json:"resourceTag"`

	// Value of the action property
	Value *string `mandatory:"false" json:"value"`

	// Description of the action property
	Description *string `mandatory:"false" json:"description"`

	// Data type of the action property
	DataType *string `mandatory:"false" json:"dataType"`

	// Whether the action property is mandatory
	IsMandatory *bool `mandatory:"false" json:"isMandatory"`

	// Hint message of the Action Property
	HintMessage *string `mandatory:"false" json:"hintMessage"`
}

func (m ActionProperty) String() string {
	return common.PointerString(m)
}
