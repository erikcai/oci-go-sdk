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

        
 // CreateTechSupportItemDetails Details of TechSupport Item
type CreateTechSupportItemDetails struct {
        
        Category *CreateCategoryDetails `mandatory:"false" json:"category"`
        
        SubCategory *CreateSubCategoryDetails `mandatory:"false" json:"subCategory"`
        
        IssueType *CreateIssueTypeDetails `mandatory:"false" json:"issueType"`
        
 // Name of the item
        Name *string `mandatory:"false" json:"name"`
}
        //GetCategory returns Category
        func (m CreateTechSupportItemDetails) GetCategory() *CreateCategoryDetails {
        return m.Category
        }
        //GetSubCategory returns SubCategory
        func (m CreateTechSupportItemDetails) GetSubCategory() *CreateSubCategoryDetails {
        return m.SubCategory
        }
        //GetIssueType returns IssueType
        func (m CreateTechSupportItemDetails) GetIssueType() *CreateIssueTypeDetails {
        return m.IssueType
        }
        //GetName returns Name
        func (m CreateTechSupportItemDetails) GetName() *string {
        return m.Name
        }

func (m CreateTechSupportItemDetails) String() string {
    return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateTechSupportItemDetails) MarshalJSON() (buff []byte, e error) {
    type MarshalTypeCreateTechSupportItemDetails CreateTechSupportItemDetails
    s := struct {
        DiscriminatorParam string `json:"type"`
        MarshalTypeCreateTechSupportItemDetails
    }{
        "tech",
        (MarshalTypeCreateTechSupportItemDetails)(m),
    }

    return json.Marshal(&s)
}






