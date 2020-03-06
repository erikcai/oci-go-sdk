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

        
 // TechSupportItem Details of TechSupport Item
type TechSupportItem struct {
        
 // Unique ID that identifies an Item
        ItemKey *string `mandatory:"true" json:"itemKey"`
        
 // Name of item
        Name *string `mandatory:"false" json:"name"`
        
        Category *Category `mandatory:"false" json:"category"`
        
        SubCategory *SubCategory `mandatory:"false" json:"subCategory"`
        
        IssueType *IssueType `mandatory:"false" json:"issueType"`
}
        //GetItemKey returns ItemKey
        func (m TechSupportItem) GetItemKey() *string {
        return m.ItemKey
        }
        //GetName returns Name
        func (m TechSupportItem) GetName() *string {
        return m.Name
        }
        //GetCategory returns Category
        func (m TechSupportItem) GetCategory() *Category {
        return m.Category
        }
        //GetSubCategory returns SubCategory
        func (m TechSupportItem) GetSubCategory() *SubCategory {
        return m.SubCategory
        }
        //GetIssueType returns IssueType
        func (m TechSupportItem) GetIssueType() *IssueType {
        return m.IssueType
        }

func (m TechSupportItem) String() string {
    return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TechSupportItem) MarshalJSON() (buff []byte, e error) {
    type MarshalTypeTechSupportItem TechSupportItem
    s := struct {
        DiscriminatorParam string `json:"type"`
        MarshalTypeTechSupportItem
    }{
        "tech",
        (MarshalTypeTechSupportItem)(m),
    }

    return json.Marshal(&s)
}






