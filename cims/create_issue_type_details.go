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


    
 // CreateIssueTypeDetails Details Issue Type of the incident
type CreateIssueTypeDetails struct {
    
 // Unique ID that identifies an Issue Type
    IssueTypeKey *string `mandatory:"false" json:"issueTypeKey"`
}

func (m CreateIssueTypeDetails) String() string {
    return common.PointerString(m)
}







