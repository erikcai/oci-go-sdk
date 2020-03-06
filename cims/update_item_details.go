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

        
 // UpdateItemDetails Details of Item
type UpdateItemDetails interface {
}

type updateitemdetails struct {
    JsonData []byte
    Type string `json:"type"`
}


// UnmarshalJSON unmarshals json
func (m *updateitemdetails) UnmarshalJSON(data []byte) error {
    m.JsonData = data
    type Unmarshalerupdateitemdetails updateitemdetails
    s := struct {
        Model Unmarshalerupdateitemdetails
    }{}
    err := json.Unmarshal(data, &s.Model)
    if err != nil {
        return err
    }
    m.Type = s.Model.Type

    return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateitemdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

    if data == nil || string(data) == "null" {
            return nil, nil
    }

    var err error
    switch(m.Type) {
       case "activity":
        mm := UpdateActivityItemDetails{}
        err = json.Unmarshal(data, &mm)
        return mm, err
    default:
        return *m, nil
    }
}


func (m updateitemdetails) String() string {
    return common.PointerString(m)
}





