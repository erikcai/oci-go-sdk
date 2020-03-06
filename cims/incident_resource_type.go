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


    
 // IncidentResourceType Details of incident type
type IncidentResourceType struct {
    
 // Label associated with Incident Type
    Label *string `mandatory:"true" json:"label"`
    
 // Unique ID that identifies an Incident Type
    ResourceTypeKey *string `mandatory:"false" json:"resourceTypeKey"`
    
 // Name of Incident type
    Name *string `mandatory:"false" json:"name"`
    
 // Details of Incident Type
    Description *string `mandatory:"false" json:"description"`
    
 // Service Category List
    ServiceCategoryList []ServiceCategory `mandatory:"false" json:"serviceCategoryList"`
}

func (m IncidentResourceType) String() string {
    return common.PointerString(m)
}







