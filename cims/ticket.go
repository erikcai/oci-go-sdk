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

// Ticket Ticket class
type Ticket struct {
	Title *string `mandatory:"true" json:"title"`

	Description *string `mandatory:"true" json:"description"`

	Id *string `mandatory:"false" json:"id"`

	Severity TicketSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	ResourceList []Resource `mandatory:"false" json:"resourceList"`

	TimeCreated *int `mandatory:"false" json:"timeCreated"`

	TimeUpdated *int `mandatory:"false" json:"timeUpdated"`

	Status *Status `mandatory:"false" json:"status"`
}

func (m Ticket) String() string {
	return common.PointerString(m)
}

// TicketSeverityEnum Enum with underlying type: string
type TicketSeverityEnum string

// Set of constants representing the allowable values for TicketSeverityEnum
const (
	TicketSeverityHighest TicketSeverityEnum = "HIGHEST"
	TicketSeverityHigh    TicketSeverityEnum = "HIGH"
	TicketSeverityMedium  TicketSeverityEnum = "MEDIUM"
	TicketSeverityLow     TicketSeverityEnum = "LOW"
	TicketSeverityLowest  TicketSeverityEnum = "LOWEST"
)

var mappingTicketSeverity = map[string]TicketSeverityEnum{
	"HIGHEST": TicketSeverityHighest,
	"HIGH":    TicketSeverityHigh,
	"MEDIUM":  TicketSeverityMedium,
	"LOW":     TicketSeverityLow,
	"LOWEST":  TicketSeverityLowest,
}

// GetTicketSeverityEnumValues Enumerates the set of values for TicketSeverityEnum
func GetTicketSeverityEnumValues() []TicketSeverityEnum {
	values := make([]TicketSeverityEnum, 0)
	for _, v := range mappingTicketSeverity {
		values = append(values, v)
	}
	return values
}
