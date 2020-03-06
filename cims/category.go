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


    
 // Category Details of Category of the incident
type Category struct {
    
 // Unique ID that identifies a Category
    CategoryKey *string `mandatory:"false" json:"categoryKey"`
    
 // Name of category. eg: Compute, Identity
    Name *string `mandatory:"false" json:"name"`
}

func (m Category) String() string {
    return common.PointerString(m)
}







