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







