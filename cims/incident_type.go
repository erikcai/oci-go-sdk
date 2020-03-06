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







