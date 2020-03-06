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


    
 // Status Details of Ticket Status
type Status struct {
    
 // Unique code
    Code *string `mandatory:"true" json:"code"`
    
 // Status message
    Message *string `mandatory:"true" json:"message"`
}

func (m Status) String() string {
    return common.PointerString(m)
}







