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


    
 // CreateSubCategoryDetails Details of Sub Category of the incident
type CreateSubCategoryDetails struct {
    
 // Unique ID that identifies a Sub Category
    SubCategoryKey *string `mandatory:"false" json:"subCategoryKey"`
}

func (m CreateSubCategoryDetails) String() string {
    return common.PointerString(m)
}







