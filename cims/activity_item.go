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

        
 // ActivityItem Details of Activity Item
type ActivityItem struct {
        
 // Unique ID that identifies an Item
        ItemKey *string `mandatory:"true" json:"itemKey"`
        
 // Name of item
        Name *string `mandatory:"false" json:"name"`
        
        Category *Category `mandatory:"false" json:"category"`
        
        SubCategory *SubCategory `mandatory:"false" json:"subCategory"`
        
        IssueType *IssueType `mandatory:"false" json:"issueType"`
        
 // Comments to update as part of Activity
        Comments *string `mandatory:"false" json:"comments"`
        
 // Epoch time when activity was created
        TimeCreated *int `mandatory:"false" json:"timeCreated"`
        
 // Epoch time when activity was updated
        TimeUpdated *int `mandatory:"false" json:"timeUpdated"`
       
 // Type of activity. eg: NOTES, UPDATE
       ActivityType ActivityItemActivityTypeEnum `mandatory:"false" json:"activityType,omitempty"`
       
 // Person who updates the activity
       ActivityAuthor ActivityItemActivityAuthorEnum `mandatory:"false" json:"activityAuthor,omitempty"`
}
        //GetItemKey returns ItemKey
        func (m ActivityItem) GetItemKey() *string {
        return m.ItemKey
        }
        //GetName returns Name
        func (m ActivityItem) GetName() *string {
        return m.Name
        }
        //GetCategory returns Category
        func (m ActivityItem) GetCategory() *Category {
        return m.Category
        }
        //GetSubCategory returns SubCategory
        func (m ActivityItem) GetSubCategory() *SubCategory {
        return m.SubCategory
        }
        //GetIssueType returns IssueType
        func (m ActivityItem) GetIssueType() *IssueType {
        return m.IssueType
        }

func (m ActivityItem) String() string {
    return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ActivityItem) MarshalJSON() (buff []byte, e error) {
    type MarshalTypeActivityItem ActivityItem
    s := struct {
        DiscriminatorParam string `json:"type"`
        MarshalTypeActivityItem
    }{
        "activity",
        (MarshalTypeActivityItem)(m),
    }

    return json.Marshal(&s)
}



// ActivityItemActivityTypeEnum Enum with underlying type: string
type ActivityItemActivityTypeEnum string

// Set of constants representing the allowable values for ActivityItemActivityTypeEnum
const (
    ActivityItemActivityTypeNotes ActivityItemActivityTypeEnum = "NOTES"
    ActivityItemActivityTypeProblemDescription ActivityItemActivityTypeEnum = "PROBLEM_DESCRIPTION"
    ActivityItemActivityTypeUpdate ActivityItemActivityTypeEnum = "UPDATE"
    ActivityItemActivityTypeClose ActivityItemActivityTypeEnum = "CLOSE"
)

var mappingActivityItemActivityType = map[string]ActivityItemActivityTypeEnum { 
    "NOTES": ActivityItemActivityTypeNotes,
    "PROBLEM_DESCRIPTION": ActivityItemActivityTypeProblemDescription,
    "UPDATE": ActivityItemActivityTypeUpdate,
    "CLOSE": ActivityItemActivityTypeClose,
}

// GetActivityItemActivityTypeEnumValues Enumerates the set of values for ActivityItemActivityTypeEnum
func GetActivityItemActivityTypeEnumValues() []ActivityItemActivityTypeEnum {
   values := make([]ActivityItemActivityTypeEnum, 0)
   for _, v := range mappingActivityItemActivityType {
       values = append(values, v)
   }
   return values
}
// ActivityItemActivityAuthorEnum Enum with underlying type: string
type ActivityItemActivityAuthorEnum string

// Set of constants representing the allowable values for ActivityItemActivityAuthorEnum
const (
    ActivityItemActivityAuthorCustomer ActivityItemActivityAuthorEnum = "CUSTOMER"
    ActivityItemActivityAuthorOracle ActivityItemActivityAuthorEnum = "ORACLE"
)

var mappingActivityItemActivityAuthor = map[string]ActivityItemActivityAuthorEnum { 
    "CUSTOMER": ActivityItemActivityAuthorCustomer,
    "ORACLE": ActivityItemActivityAuthorOracle,
}

// GetActivityItemActivityAuthorEnumValues Enumerates the set of values for ActivityItemActivityAuthorEnum
func GetActivityItemActivityAuthorEnumValues() []ActivityItemActivityAuthorEnum {
   values := make([]ActivityItemActivityAuthorEnum, 0)
   for _, v := range mappingActivityItemActivityAuthor {
       values = append(values, v)
   }
   return values
}


