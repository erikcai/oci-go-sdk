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


    
 // Incident Details of Incident
type Incident struct {
    
 // Unique ID that identifies an Incident
    Key *string `mandatory:"true" json:"key"`
    
 // Tenancy Ocid
    CompartmentId *string `mandatory:"false" json:"compartmentId"`
    
    ContactList *ContactList `mandatory:"false" json:"contactList"`
    
    TenancyInformation *TenancyInformation `mandatory:"false" json:"tenancyInformation"`
    
    Ticket *Ticket `mandatory:"false" json:"ticket"`
    
    IncidentType *IncidentType `mandatory:"false" json:"incidentType"`
    
 // States type of incident. eg: LIMIT, TECH
    ProblemType ProblemTypeEnum `mandatory:"false" json:"problemType,omitempty"`
}

func (m Incident) String() string {
    return common.PointerString(m)
}







