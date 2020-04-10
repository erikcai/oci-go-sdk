// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Support Management API
// 
 // Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
    "github.com/oracle/oci-go-sdk/common"
        "encoding/json"
)


    
 // UpdateResourceDetails Update Resource details
type UpdateResourceDetails struct {
    
    Item UpdateItemDetails `mandatory:"false" json:"item"`
}

func (m UpdateResourceDetails) String() string {
    return common.PointerString(m)
}

    // UnmarshalJSON unmarshals from json
    func (m *UpdateResourceDetails) UnmarshalJSON(data []byte) (e error){
        model := struct{
                Item updateitemdetails `json:"item"`
        }{}

        e = json.Unmarshal(data, &model)
        if e != nil {
            return
        }
            var nn interface{}
        nn, e = model.Item.UnmarshalPolymorphicJSON(model.Item.JsonData)
        if e != nil {
            return
        }
        if nn != nil {
            m.Item = nn.(UpdateItemDetails)
        } else {
            m.Item = nil
        }
    return
    }






