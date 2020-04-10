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


    
 // IncidentType Details of incident type
type IncidentType struct {
    
 // Unique ID that identifies an Incident Type
    Id *string `mandatory:"false" json:"id"`
    
 // Name of Incident type
    Name *string `mandatory:"false" json:"name"`
    
 // Label associated with Incident Type
    Label *string `mandatory:"false" json:"label"`
    
 // Details of Incident Type
    Description *string `mandatory:"false" json:"description"`
    
 // List of classifiers
    ClassifierList []Classifier `mandatory:"false" json:"classifierList"`
}

func (m IncidentType) String() string {
    return common.PointerString(m)
}







