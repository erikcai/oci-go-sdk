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

// UpdateIncidentDetails The representation of UpdateIncidentDetails
type UpdateIncidentDetails struct {
	Id *string `mandatory:"false" json:"id"`

	Reporter *Reporter `mandatory:"false" json:"reporter"`

	ContactList *ContactList `mandatory:"false" json:"contactList"`

	Tenant *Tenant `mandatory:"false" json:"tenant"`

	Ticket *Ticket `mandatory:"false" json:"ticket"`

	ClientSource UpdateIncidentDetailsClientSourceEnum `mandatory:"false" json:"clientSource,omitempty"`

	IncidentType *IncidentType `mandatory:"false" json:"incidentType"`
}

func (m UpdateIncidentDetails) String() string {
	return common.PointerString(m)
}

// UpdateIncidentDetailsClientSourceEnum Enum with underlying type: string
type UpdateIncidentDetailsClientSourceEnum string

// Set of constants representing the allowable values for UpdateIncidentDetailsClientSourceEnum
const (
	UpdateIncidentDetailsClientSourceOciConsole UpdateIncidentDetailsClientSourceEnum = "OCI_CONSOLE"
	UpdateIncidentDetailsClientSourceOracleMos  UpdateIncidentDetailsClientSourceEnum = "ORACLE_MOS"
)

var mappingUpdateIncidentDetailsClientSource = map[string]UpdateIncidentDetailsClientSourceEnum{
	"OCI_CONSOLE": UpdateIncidentDetailsClientSourceOciConsole,
	"ORACLE_MOS":  UpdateIncidentDetailsClientSourceOracleMos,
}

// GetUpdateIncidentDetailsClientSourceEnumValues Enumerates the set of values for UpdateIncidentDetailsClientSourceEnum
func GetUpdateIncidentDetailsClientSourceEnumValues() []UpdateIncidentDetailsClientSourceEnum {
	values := make([]UpdateIncidentDetailsClientSourceEnum, 0)
	for _, v := range mappingUpdateIncidentDetailsClientSource {
		values = append(values, v)
	}
	return values
}
