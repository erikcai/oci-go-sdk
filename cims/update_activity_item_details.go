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

        
 // UpdateActivityItemDetails Details of Activity Item
type UpdateActivityItemDetails struct {
        
 // Comments to update as part of Activity
        Comments *string `mandatory:"false" json:"comments"`
       
 // Type of activity. eg: NOTES, UPDATE
       ActivityType UpdateActivityItemDetailsActivityTypeEnum `mandatory:"false" json:"activityType,omitempty"`
}

func (m UpdateActivityItemDetails) String() string {
    return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateActivityItemDetails) MarshalJSON() (buff []byte, e error) {
    type MarshalTypeUpdateActivityItemDetails UpdateActivityItemDetails
    s := struct {
        DiscriminatorParam string `json:"type"`
        MarshalTypeUpdateActivityItemDetails
    }{
        "activity",
        (MarshalTypeUpdateActivityItemDetails)(m),
    }

    return json.Marshal(&s)
}



// UpdateActivityItemDetailsActivityTypeEnum Enum with underlying type: string
type UpdateActivityItemDetailsActivityTypeEnum string

// Set of constants representing the allowable values for UpdateActivityItemDetailsActivityTypeEnum
const (
    UpdateActivityItemDetailsActivityTypeNotes UpdateActivityItemDetailsActivityTypeEnum = "NOTES"
    UpdateActivityItemDetailsActivityTypeProblemDescription UpdateActivityItemDetailsActivityTypeEnum = "PROBLEM_DESCRIPTION"
    UpdateActivityItemDetailsActivityTypeUpdate UpdateActivityItemDetailsActivityTypeEnum = "UPDATE"
    UpdateActivityItemDetailsActivityTypeClose UpdateActivityItemDetailsActivityTypeEnum = "CLOSE"
)

var mappingUpdateActivityItemDetailsActivityType = map[string]UpdateActivityItemDetailsActivityTypeEnum { 
    "NOTES": UpdateActivityItemDetailsActivityTypeNotes,
    "PROBLEM_DESCRIPTION": UpdateActivityItemDetailsActivityTypeProblemDescription,
    "UPDATE": UpdateActivityItemDetailsActivityTypeUpdate,
    "CLOSE": UpdateActivityItemDetailsActivityTypeClose,
}

// GetUpdateActivityItemDetailsActivityTypeEnumValues Enumerates the set of values for UpdateActivityItemDetailsActivityTypeEnum
func GetUpdateActivityItemDetailsActivityTypeEnumValues() []UpdateActivityItemDetailsActivityTypeEnum {
   values := make([]UpdateActivityItemDetailsActivityTypeEnum, 0)
   for _, v := range mappingUpdateActivityItemDetailsActivityType {
       values = append(values, v)
   }
   return values
}



