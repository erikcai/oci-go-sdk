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

// Incident The representation of Incident
type Incident struct {
	Id *string `mandatory:"false" json:"id"`

	Reporter *Reporter `mandatory:"false" json:"reporter"`

	ContactList *ContactList `mandatory:"false" json:"contactList"`

	Tenant *Tenant `mandatory:"false" json:"tenant"`

	Ticket *Ticket `mandatory:"false" json:"ticket"`

	ClientSource IncidentClientSourceEnum `mandatory:"false" json:"clientSource,omitempty"`

	IncidentType *IncidentType `mandatory:"false" json:"incidentType"`
}

func (m Incident) String() string {
	return common.PointerString(m)
}

// IncidentClientSourceEnum Enum with underlying type: string
type IncidentClientSourceEnum string

// Set of constants representing the allowable values for IncidentClientSourceEnum
const (
	IncidentClientSourceOciConsole IncidentClientSourceEnum = "OCI_CONSOLE"
	IncidentClientSourceOracleMos  IncidentClientSourceEnum = "ORACLE_MOS"
)

var mappingIncidentClientSource = map[string]IncidentClientSourceEnum{
	"OCI_CONSOLE": IncidentClientSourceOciConsole,
	"ORACLE_MOS":  IncidentClientSourceOracleMos,
}

// GetIncidentClientSourceEnumValues Enumerates the set of values for IncidentClientSourceEnum
func GetIncidentClientSourceEnumValues() []IncidentClientSourceEnum {
	values := make([]IncidentClientSourceEnum, 0)
	for _, v := range mappingIncidentClientSource {
		values = append(values, v)
	}
	return values
}
