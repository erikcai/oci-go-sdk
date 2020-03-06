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


    
 // ErrorInfo The representation of ErrorInfo
type ErrorInfo struct {
    
 // Describes which exception has occurred
    ErrorCode ErrorCodeEnum `mandatory:"false" json:"errorCode,omitempty"`
    
 // Details about the exceptions
    Message *string `mandatory:"false" json:"message"`
}

func (m ErrorInfo) String() string {
    return common.PointerString(m)
}







