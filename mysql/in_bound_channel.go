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

// InBoundChannel The representation of InBoundChannel
type InBoundChannel struct {

	// The OCID of the InBoundChannel used to receive data streams.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the DbSystem the InBoundChannel is associated to.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// Name that uniquely identifies the channel.
	Name *string `mandatory:"true" json:"name"`

	// Address that a MySQLaaS instance or DbSystem shall connect to receive data changes.
	Address *string `mandatory:"true" json:"address"`

	// The user's name that MySQLaaS instance or DbSystemm shall use to connect to an entity to
	// receive data changes.
	Username *string `mandatory:"true" json:"username"`

	// The Oracle Cloud ID (OCID) of the compartment the DbSystem belongs in.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// User-provided data about the in-bound channel.
	Description *string `mandatory:"false" json:"description"`

	// Port that a MySQLaaS instance or DbSystem shall connect to receive data changes. Default 3306.
	Port *int `mandatory:"false" json:"port"`
}

func (m InBoundChannel) String() string {
	return common.PointerString(m)
}
