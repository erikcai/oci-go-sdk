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

        
 // LimitItem Details of Limit Item
type LimitItem struct {
        
 // Unique ID that identifies an Item
        ItemKey *string `mandatory:"true" json:"itemKey"`
        
 // Name of item
        Name *string `mandatory:"false" json:"name"`
        
        Category *Category `mandatory:"false" json:"category"`
        
        SubCategory *SubCategory `mandatory:"false" json:"subCategory"`
        
        IssueType *IssueType `mandatory:"false" json:"issueType"`
        
 // Current available limit of the resource
        CurrentLimit *int `mandatory:"false" json:"currentLimit"`
        
 // Current used limit of the resource
        CurrentUsage *int `mandatory:"false" json:"currentUsage"`
        
 // Requested limit for the resource
        RequestedLimit *int `mandatory:"false" json:"requestedLimit"`
       
 // Status of the Limit
       LimitStatus LimitItemLimitStatusEnum `mandatory:"false" json:"limitStatus,omitempty"`
}
        //GetItemKey returns ItemKey
        func (m LimitItem) GetItemKey() *string {
        return m.ItemKey
        }
        //GetName returns Name
        func (m LimitItem) GetName() *string {
        return m.Name
        }
        //GetCategory returns Category
        func (m LimitItem) GetCategory() *Category {
        return m.Category
        }
        //GetSubCategory returns SubCategory
        func (m LimitItem) GetSubCategory() *SubCategory {
        return m.SubCategory
        }
        //GetIssueType returns IssueType
        func (m LimitItem) GetIssueType() *IssueType {
        return m.IssueType
        }

func (m LimitItem) String() string {
    return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m LimitItem) MarshalJSON() (buff []byte, e error) {
    type MarshalTypeLimitItem LimitItem
    s := struct {
        DiscriminatorParam string `json:"type"`
        MarshalTypeLimitItem
    }{
        "limit",
        (MarshalTypeLimitItem)(m),
    }

    return json.Marshal(&s)
}



// LimitItemLimitStatusEnum Enum with underlying type: string
type LimitItemLimitStatusEnum string

// Set of constants representing the allowable values for LimitItemLimitStatusEnum
const (
    LimitItemLimitStatusApproved LimitItemLimitStatusEnum = "APPROVED"
    LimitItemLimitStatusPartiallyApproved LimitItemLimitStatusEnum = "PARTIALLY_APPROVED"
    LimitItemLimitStatusNotApproved LimitItemLimitStatusEnum = "NOT_APPROVED"
)

var mappingLimitItemLimitStatus = map[string]LimitItemLimitStatusEnum { 
    "APPROVED": LimitItemLimitStatusApproved,
    "PARTIALLY_APPROVED": LimitItemLimitStatusPartiallyApproved,
    "NOT_APPROVED": LimitItemLimitStatusNotApproved,
}

// GetLimitItemLimitStatusEnumValues Enumerates the set of values for LimitItemLimitStatusEnum
func GetLimitItemLimitStatusEnumValues() []LimitItemLimitStatusEnum {
   values := make([]LimitItemLimitStatusEnum, 0)
   for _, v := range mappingLimitItemLimitStatus {
       values = append(values, v)
   }
   return values
}


