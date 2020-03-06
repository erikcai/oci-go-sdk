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


    
 // UpdateTicketDetails Details of Ticket updated
type UpdateTicketDetails struct {
    
 // List of resources
    Resource *interface{} `mandatory:"true" json:"resource"`
}

func (m UpdateTicketDetails) String() string {
    return common.PointerString(m)
}







