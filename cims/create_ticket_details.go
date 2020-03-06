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


    
 // CreateTicketDetails Details of Ticket created
type CreateTicketDetails struct {
    
 // Severity of the ticket. eg: HIGH, MEDIUM
    Severity CreateTicketDetailsSeverityEnum `mandatory:"true" json:"severity"`
    
 // Title of ticket
    Title *string `mandatory:"true" json:"title"`
    
 // Details of ticket
    Description *string `mandatory:"true" json:"description"`
    
 // List of resources
    ResourceList []CreateResourceDetails `mandatory:"false" json:"resourceList"`
}

func (m CreateTicketDetails) String() string {
    return common.PointerString(m)
}




// CreateTicketDetailsSeverityEnum Enum with underlying type: string
type CreateTicketDetailsSeverityEnum string

// Set of constants representing the allowable values for CreateTicketDetailsSeverityEnum
const (
    CreateTicketDetailsSeverityHighest CreateTicketDetailsSeverityEnum = "HIGHEST"
    CreateTicketDetailsSeverityHigh CreateTicketDetailsSeverityEnum = "HIGH"
    CreateTicketDetailsSeverityMedium CreateTicketDetailsSeverityEnum = "MEDIUM"
)

var mappingCreateTicketDetailsSeverity = map[string]CreateTicketDetailsSeverityEnum { 
    "HIGHEST": CreateTicketDetailsSeverityHighest,
    "HIGH": CreateTicketDetailsSeverityHigh,
    "MEDIUM": CreateTicketDetailsSeverityMedium,
}

// GetCreateTicketDetailsSeverityEnumValues Enumerates the set of values for CreateTicketDetailsSeverityEnum
func GetCreateTicketDetailsSeverityEnumValues() []CreateTicketDetailsSeverityEnum {
   values := make([]CreateTicketDetailsSeverityEnum, 0)
   for _, v := range mappingCreateTicketDetailsSeverity {
       values = append(values, v)
   }
   return values
}



