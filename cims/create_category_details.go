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


    
 // CreateCategoryDetails Details of Category of the incident
type CreateCategoryDetails struct {
    
 // Unique ID that identifies a Category
    CategoryKey *string `mandatory:"false" json:"categoryKey"`
}

func (m CreateCategoryDetails) String() string {
    return common.PointerString(m)
}







