// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
// 
 // Cloud Incident Management Service
//

package cims

import (
    "github.com/oracle/oci-go-sdk/common"
        "encoding/json"
)

        
 // CreateLimitItemDetails Details of Limit Item
type CreateLimitItemDetails struct {
        
        Category *CreateCategoryDetails `mandatory:"false" json:"category"`
        
        SubCategory *CreateSubCategoryDetails `mandatory:"false" json:"subCategory"`
        
        IssueType *CreateIssueTypeDetails `mandatory:"false" json:"issueType"`
        
 // Name of the item
        Name *string `mandatory:"false" json:"name"`
        
 // Current available limit of the resource
        CurrentLimit *int `mandatory:"false" json:"currentLimit"`
        
 // Current used limit of the resource
        CurrentUsage *int `mandatory:"false" json:"currentUsage"`
        
 // Requested limit for the resource
        RequestedLimit *int `mandatory:"false" json:"requestedLimit"`
       
 // Status of the Limit
       LimitStatus CreateLimitItemDetailsLimitStatusEnum `mandatory:"false" json:"limitStatus,omitempty"`
}
        //GetCategory returns Category
        func (m CreateLimitItemDetails) GetCategory() *CreateCategoryDetails {
        return m.Category
        }
        //GetSubCategory returns SubCategory
        func (m CreateLimitItemDetails) GetSubCategory() *CreateSubCategoryDetails {
        return m.SubCategory
        }
        //GetIssueType returns IssueType
        func (m CreateLimitItemDetails) GetIssueType() *CreateIssueTypeDetails {
        return m.IssueType
        }
        //GetName returns Name
        func (m CreateLimitItemDetails) GetName() *string {
        return m.Name
        }

func (m CreateLimitItemDetails) String() string {
    return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateLimitItemDetails) MarshalJSON() (buff []byte, e error) {
    type MarshalTypeCreateLimitItemDetails CreateLimitItemDetails
    s := struct {
        DiscriminatorParam string `json:"type"`
        MarshalTypeCreateLimitItemDetails
    }{
        "limit",
        (MarshalTypeCreateLimitItemDetails)(m),
    }

    return json.Marshal(&s)
}



// CreateLimitItemDetailsLimitStatusEnum Enum with underlying type: string
type CreateLimitItemDetailsLimitStatusEnum string

// Set of constants representing the allowable values for CreateLimitItemDetailsLimitStatusEnum
const (
    CreateLimitItemDetailsLimitStatusApproved CreateLimitItemDetailsLimitStatusEnum = "APPROVED"
    CreateLimitItemDetailsLimitStatusPartiallyApproved CreateLimitItemDetailsLimitStatusEnum = "PARTIALLY_APPROVED"
    CreateLimitItemDetailsLimitStatusNotApproved CreateLimitItemDetailsLimitStatusEnum = "NOT_APPROVED"
)

var mappingCreateLimitItemDetailsLimitStatus = map[string]CreateLimitItemDetailsLimitStatusEnum { 
    "APPROVED": CreateLimitItemDetailsLimitStatusApproved,
    "PARTIALLY_APPROVED": CreateLimitItemDetailsLimitStatusPartiallyApproved,
    "NOT_APPROVED": CreateLimitItemDetailsLimitStatusNotApproved,
}

// GetCreateLimitItemDetailsLimitStatusEnumValues Enumerates the set of values for CreateLimitItemDetailsLimitStatusEnum
func GetCreateLimitItemDetailsLimitStatusEnumValues() []CreateLimitItemDetailsLimitStatusEnum {
   values := make([]CreateLimitItemDetailsLimitStatusEnum, 0)
   for _, v := range mappingCreateLimitItemDetailsLimitStatus {
       values = append(values, v)
   }
   return values
}



