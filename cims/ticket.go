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


    
 // Ticket Details of Ticket created
type Ticket struct {
    
 // Severity of the ticket. eg: HIGH, MEDIUM
    Severity TicketSeverityEnum `mandatory:"true" json:"severity"`
    
 // Title of ticket
    Title *string `mandatory:"true" json:"title"`
    
 // Details of ticket
    Description *string `mandatory:"true" json:"description"`
    
 // Unique ID that identifies a Ticket
    TicketNumber *string `mandatory:"false" json:"ticketNumber"`
    
 // List of resources
    ResourceList []Resource `mandatory:"false" json:"resourceList"`
    
 // Epoch time of ticket creation
    TimeCreated *int `mandatory:"false" json:"timeCreated"`
    
 // Epoch time of ticket updated
    TimeUpdated *int `mandatory:"false" json:"timeUpdated"`
    
 // Ticket's SEV1 Work hours
    IsWork247 *bool `mandatory:"false" json:"isWork247"`
    
 // Describes the lifecycles of a ticket
    LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
    
 // Describes the lifecycle details of a ticket
    LifecycleDetails LifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`
}

func (m Ticket) String() string {
    return common.PointerString(m)
}




// TicketSeverityEnum Enum with underlying type: string
type TicketSeverityEnum string

// Set of constants representing the allowable values for TicketSeverityEnum
const (
    TicketSeverityHighest TicketSeverityEnum = "HIGHEST"
    TicketSeverityHigh TicketSeverityEnum = "HIGH"
    TicketSeverityMedium TicketSeverityEnum = "MEDIUM"
)

var mappingTicketSeverity = map[string]TicketSeverityEnum { 
    "HIGHEST": TicketSeverityHighest,
    "HIGH": TicketSeverityHigh,
    "MEDIUM": TicketSeverityMedium,
}

// GetTicketSeverityEnumValues Enumerates the set of values for TicketSeverityEnum
func GetTicketSeverityEnumValues() []TicketSeverityEnum {
   values := make([]TicketSeverityEnum, 0)
   for _, v := range mappingTicketSeverity {
       values = append(values, v)
   }
   return values
}



