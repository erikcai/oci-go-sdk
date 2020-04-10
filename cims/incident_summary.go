// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Support Management API
// 
 // Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
    "github.com/oracle/oci-go-sdk/common"
)


    
 // IncidentSummary Details of Incident
type IncidentSummary struct {
    
 // Unique ID that identifies an Incident
    Key *string `mandatory:"true" json:"key"`
    
 // States type of incident. eg: LIMIT, TECH
    ProblemType ProblemTypeEnum `mandatory:"true" json:"problemType"`
    
 // Tenancy Ocid
    CompartmentId *string `mandatory:"false" json:"compartmentId"`
    
    ContactList *ContactList `mandatory:"false" json:"contactList"`
    
    TenancyInformation *TenancyInformation `mandatory:"false" json:"tenancyInformation"`
    
    Ticket *Ticket `mandatory:"false" json:"ticket"`
    
    IncidentType *IncidentResourceType `mandatory:"false" json:"incidentType"`
}

func (m IncidentSummary) String() string {
    return common.PointerString(m)
}







