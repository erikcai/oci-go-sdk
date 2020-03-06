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

        
 // CreateItemDetails Details of Item
type CreateItemDetails interface {
    
    GetCategory() *CreateCategoryDetails
    
    GetSubCategory() *CreateSubCategoryDetails
    
    GetIssueType() *CreateIssueTypeDetails
    
 // Name of the item
    GetName() *string
}

type createitemdetails struct {
    JsonData []byte
    Category *CreateCategoryDetails `mandatory:"false" json:"category"`
    SubCategory *CreateSubCategoryDetails `mandatory:"false" json:"subCategory"`
    IssueType *CreateIssueTypeDetails `mandatory:"false" json:"issueType"`
    Name *string `mandatory:"false" json:"name"`
    Type string `json:"type"`
}


// UnmarshalJSON unmarshals json
func (m *createitemdetails) UnmarshalJSON(data []byte) error {
    m.JsonData = data
    type Unmarshalercreateitemdetails createitemdetails
    s := struct {
        Model Unmarshalercreateitemdetails
    }{}
    err := json.Unmarshal(data, &s.Model)
    if err != nil {
        return err
    }
        m.Category = s.Model.Category
        m.SubCategory = s.Model.SubCategory
        m.IssueType = s.Model.IssueType
        m.Name = s.Model.Name
    m.Type = s.Model.Type

    return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createitemdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

    if data == nil || string(data) == "null" {
            return nil, nil
    }

    var err error
    switch(m.Type) {
       case "tech":
        mm := CreateTechSupportItemDetails{}
        err = json.Unmarshal(data, &mm)
        return mm, err
       case "limit":
        mm := CreateLimitItemDetails{}
        err = json.Unmarshal(data, &mm)
        return mm, err
    default:
        return *m, nil
    }
}

    //GetCategory returns Category
    func (m createitemdetails) GetCategory()  *CreateCategoryDetails {
    return m.Category
    }
    //GetSubCategory returns SubCategory
    func (m createitemdetails) GetSubCategory()  *CreateSubCategoryDetails {
    return m.SubCategory
    }
    //GetIssueType returns IssueType
    func (m createitemdetails) GetIssueType()  *CreateIssueTypeDetails {
    return m.IssueType
    }
    //GetName returns Name
    func (m createitemdetails) GetName()  *string {
    return m.Name
    }

func (m createitemdetails) String() string {
    return common.PointerString(m)
}





