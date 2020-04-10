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


    
 // UpdateTicketDetails Details of Ticket updated
type UpdateTicketDetails struct {
    
 // List of resources
    Resource *interface{} `mandatory:"true" json:"resource"`
}

func (m UpdateTicketDetails) String() string {
    return common.PointerString(m)
}







