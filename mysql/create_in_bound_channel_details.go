// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateInBoundChannelDetails The representation of CreateInBoundChannelDetails
type CreateInBoundChannelDetails struct {

	// The Oracle Cloud ID (OCID) of the compartment the DbSystem belongs in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name that uniquely identifies the channel.
	Name *string `mandatory:"true" json:"name"`

	// Address a MySQLaaS instance or DbSystem connects to receive data changes.
	Address *string `mandatory:"true" json:"address"`

	// The username the MySQLaaS instance or DbSystem uses to connect to an entity to
	// receive data changes.
	Username *string `mandatory:"true" json:"username"`

	// The user's password MySQLaaS instance or DbSystem uses to connect to an entity
	// to receive data changes.
	Password *string `mandatory:"true" json:"password"`

	// User-provided data about the in-bound channel.
	Description *string `mandatory:"false" json:"description"`

	// Port a MySQLaaS instance or DbSystem connects to receive data changes. Default 3306.
	Port *int `mandatory:"false" json:"port"`
}

func (m CreateInBoundChannelDetails) String() string {
	return common.PointerString(m)
}
