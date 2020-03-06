// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
// 
 // Cloud Incident Management Service
//

package cims

import (
    "github.com/oracle/oci-go-sdk/common"
)


    
 // ContactList List of contacts
type ContactList struct {
    
 // List of contacts
    ContactList []Contact `mandatory:"true" json:"contactList"`
}

func (m ContactList) String() string {
    return common.PointerString(m)
}







