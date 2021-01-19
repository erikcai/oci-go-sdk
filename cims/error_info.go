// Copyright (c) 2016 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// CIMS API
// 
 // Cloud Incident Management Service
//

package cims

import (
    "github.com/oracle/oci-go-sdk/v33/common"
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







