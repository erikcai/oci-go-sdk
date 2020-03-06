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


    
 // TenancyInformation Details of Customer Tenant
type TenancyInformation struct {
    
 // Tenant customer support identifier
    CustomerSupportKey *string `mandatory:"true" json:"customerSupportKey"`
    
 // Tenant OCID
    TenancyId *string `mandatory:"true" json:"tenancyId"`
}

func (m TenancyInformation) String() string {
    return common.PointerString(m)
}







