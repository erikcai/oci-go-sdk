// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Key Management Service API
// 
 // APIs for managing and performing operations with keys and vaults.
//

package kms

import (
    "github.com/oracle/oci-go-sdk/common"
)


    
 // KeyShape TODO
type KeyShape struct {
    
 // The algorithm used by a Key's KeyVersions.  TODO description
    Algorithm KeyShapeAlgorithmEnum `mandatory:"true" json:"algorithm"`
    
 // The length of the Key.  TODO description
    Length *int `mandatory:"true" json:"length"`
}

func (m KeyShape) String() string {
    return common.PointerString(m)
}


// KeyShapeAlgorithmEnum Enum with underlying type: string
type KeyShapeAlgorithmEnum string

// Set of constants representing the allowable values for KeyShapeAlgorithm
const (
    KeyShapeAlgorithmAes KeyShapeAlgorithmEnum = "AES"
)

var mappingKeyShapeAlgorithm = map[string]KeyShapeAlgorithmEnum { 
    "AES": KeyShapeAlgorithmAes,
}

// GetKeyShapeAlgorithmEnumValues Enumerates the set of values for KeyShapeAlgorithm
func GetKeyShapeAlgorithmEnumValues() []KeyShapeAlgorithmEnum {
   values := make([]KeyShapeAlgorithmEnum, 0)
   for _, v := range mappingKeyShapeAlgorithm {
       values = append(values, v)
   }
   return values
}



