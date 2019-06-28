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

// CreateIncidentDetails The representation of CreateIncidentDetails
type CreateIncidentDetails struct {
	Id *string `mandatory:"false" json:"id"`

	Reporter *Reporter `mandatory:"false" json:"reporter"`

	ContactList *ContactList `mandatory:"false" json:"contactList"`

	Tenant *Tenant `mandatory:"false" json:"tenant"`

	Ticket *Ticket `mandatory:"false" json:"ticket"`

	ClientSource CreateIncidentDetailsClientSourceEnum `mandatory:"false" json:"clientSource,omitempty"`

	IncidentType *IncidentType `mandatory:"false" json:"incidentType"`
}

func (m CreateIncidentDetails) String() string {
	return common.PointerString(m)
}

// CreateIncidentDetailsClientSourceEnum Enum with underlying type: string
type CreateIncidentDetailsClientSourceEnum string

// Set of constants representing the allowable values for CreateIncidentDetailsClientSourceEnum
const (
	CreateIncidentDetailsClientSourceOciConsole CreateIncidentDetailsClientSourceEnum = "OCI_CONSOLE"
	CreateIncidentDetailsClientSourceOracleMos  CreateIncidentDetailsClientSourceEnum = "ORACLE_MOS"
)

var mappingCreateIncidentDetailsClientSource = map[string]CreateIncidentDetailsClientSourceEnum{
	"OCI_CONSOLE": CreateIncidentDetailsClientSourceOciConsole,
	"ORACLE_MOS":  CreateIncidentDetailsClientSourceOracleMos,
}

// GetCreateIncidentDetailsClientSourceEnumValues Enumerates the set of values for CreateIncidentDetailsClientSourceEnum
func GetCreateIncidentDetailsClientSourceEnumValues() []CreateIncidentDetailsClientSourceEnum {
	values := make([]CreateIncidentDetailsClientSourceEnum, 0)
	for _, v := range mappingCreateIncidentDetailsClientSource {
		values = append(values, v)
	}
	return values
}
