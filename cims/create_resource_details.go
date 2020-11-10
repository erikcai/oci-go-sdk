// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
// 
 // Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
    "github.com/oracle/oci-go-sdk/v28/common"
        "encoding/json"
)


    
 // CreateResourceDetails Details about the resource that the support ticket relates to.
 // **Caution:** Avoid using any confidential information when you supply string values using the API.
type CreateResourceDetails struct {
    
    Item CreateItemDetails `mandatory:"false" json:"item"`
    
 // The list of available Oracle Cloud Infrastructure regions.
    Region RegionEnum `mandatory:"false" json:"region,omitempty"`
    
 // The list of available Oracle Cloud Infrastructure availability domains.
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






