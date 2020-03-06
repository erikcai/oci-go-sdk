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

        
 // Item Details of Item
type Item interface {
    
 // Unique ID that identifies an Item
    GetItemKey() *string
    
 // Name of item
    GetName() *string
    
    GetCategory() *Category
    
    GetSubCategory() *SubCategory
    
    GetIssueType() *IssueType
}

type item struct {
    JsonData []byte
    ItemKey *string `mandatory:"true" json:"itemKey"`
    Name *string `mandatory:"false" json:"name"`
    Category *Category `mandatory:"false" json:"category"`
    SubCategory *SubCategory `mandatory:"false" json:"subCategory"`
    IssueType *IssueType `mandatory:"false" json:"issueType"`
    Type string `json:"type"`
}


// UnmarshalJSON unmarshals json
func (m *item) UnmarshalJSON(data []byte) error {
    m.JsonData = data
    type Unmarshaleritem item
    s := struct {
        Model Unmarshaleritem
    }{}
    err := json.Unmarshal(data, &s.Model)
    if err != nil {
        return err
    }
        m.ItemKey = s.Model.ItemKey
        m.Name = s.Model.Name
        m.Category = s.Model.Category
        m.SubCategory = s.Model.SubCategory
        m.IssueType = s.Model.IssueType
    m.Type = s.Model.Type

    return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *item) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

    if data == nil || string(data) == "null" {
            return nil, nil
    }

    var err error
    switch(m.Type) {
       case "limit":
        mm := LimitItem{}
        err = json.Unmarshal(data, &mm)
        return mm, err
       case "tech":
        mm := TechSupportItem{}
        err = json.Unmarshal(data, &mm)
        return mm, err
       case "activity":
        mm := ActivityItem{}
        err = json.Unmarshal(data, &mm)
        return mm, err
    default:
        return *m, nil
    }
}

    //GetItemKey returns ItemKey
    func (m item) GetItemKey()  *string {
        return m.ItemKey
    }
    //GetName returns Name
    func (m item) GetName()  *string {
    return m.Name
    }
    //GetCategory returns Category
    func (m item) GetCategory()  *Category {
    return m.Category
    }
    //GetSubCategory returns SubCategory
    func (m item) GetSubCategory()  *SubCategory {
    return m.SubCategory
    }
    //GetIssueType returns IssueType
    func (m item) GetIssueType()  *IssueType {
    return m.IssueType
    }

func (m item) String() string {
    return common.PointerString(m)
}





