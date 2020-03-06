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


    
 // ServiceCategory Incident Classifier details
type ServiceCategory struct {
    
 // Unique ID that identifies a classifier
    Key *string `mandatory:"false" json:"key"`
    
 // Name of classifier. eg: LIMIT Increase
    Name *string `mandatory:"false" json:"name"`
    
 // Label of classifier
    Label *string `mandatory:"false" json:"label"`
    
 // Description of classifier
    Description *string `mandatory:"false" json:"description"`
    
 // List of Issues
    IssueTypeList []IssueType `mandatory:"false" json:"issueTypeList"`
    
 // List of Scope
    Scope ScopeEnum `mandatory:"false" json:"scope,omitempty"`
    
 // List of Units
    Unit UnitEnum `mandatory:"false" json:"unit,omitempty"`
    
 // Limit's unique id
    LimitId *string `mandatory:"false" json:"limitId"`
}

func (m ServiceCategory) String() string {
    return common.PointerString(m)
}







