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


    
 // IssueType Details Issue Type of the incident
type IssueType struct {
    
 // Unique ID that identifies an Issue Type
    IssueTypeKey *string `mandatory:"false" json:"issueTypeKey"`
    
 // Label of issue type. eg: Instance Performance
    Label *string `mandatory:"false" json:"label"`
}

func (m IssueType) String() string {
    return common.PointerString(m)
}







