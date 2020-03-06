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


    
 // CreateResourceDetails Details of Ticket Item
type CreateResourceDetails struct {
    
    Item CreateItemDetails `mandatory:"false" json:"item"`
    
 // List of OCI regions
    Region RegionEnum `mandatory:"false" json:"region,omitempty"`
    
 // List of OCI ADs
    AvailabilityDomain AvailabilityDomainEnum `mandatory:"false" json:"availabilityDomain,omitempty"`
}

func (m CreateResourceDetails) String() string {
    return common.PointerString(m)
}

    // UnmarshalJSON unmarshals from json
    func (m *CreateResourceDetails) UnmarshalJSON(data []byte) (e error){
        model := struct{
                Item createitemdetails `json:"item"`
                Region RegionEnum `json:"region"`
                AvailabilityDomain AvailabilityDomainEnum `json:"availabilityDomain"`
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
            m.Item = nn.(CreateItemDetails)
        } else {
            m.Item = nil
        }
            
        m.Region = model.Region
            
        m.AvailabilityDomain = model.AvailabilityDomain
    return
    }






