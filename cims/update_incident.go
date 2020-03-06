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


    
 // UpdateIncident Details of Resource Item to be updated
type UpdateIncident struct {
    
    Ticket *UpdateTicketDetails `mandatory:"true" json:"ticket"`
}

func (m UpdateIncident) String() string {
    return common.PointerString(m)
}







