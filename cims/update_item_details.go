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





