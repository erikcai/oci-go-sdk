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






