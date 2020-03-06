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


    
 // CreateIncident Details of Incident
type CreateIncident struct {
    
 // Tenancy Ocid
    CompartmentId *string `mandatory:"true" json:"compartmentId"`
    
    Ticket *CreateTicketDetails `mandatory:"true" json:"ticket"`
    
 // States type of incident. eg: LIMIT, TECH
    ProblemType ProblemTypeEnum `mandatory:"true" json:"problemType"`
    
 // Customer Support Identifier of the support account
    Csi *string `mandatory:"false" json:"csi"`
}

func (m CreateIncident) String() string {
    return common.PointerString(m)
}







