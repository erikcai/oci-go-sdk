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


    
 // Key The representation of Key
type Key struct {
    
 // The OCID of the Compartment containing this resource.
    CompartmentId *string `mandatory:"true" json:"compartmentId"`
    
 // The OCID of the KeyVersion that is used in Encrypt operations.
    CurrentKeyVersion *string `mandatory:"true" json:"currentKeyVersion"`
    
 // A user-friendly name. Does not have to be unique, and it's changeable.
 // Avoid entering confidential information.
    DisplayName *string `mandatory:"true" json:"displayName"`
    
 // The OCID of the resource.
    Id *string `mandatory:"true" json:"id"`
    
    KeyShape *KeyShape `mandatory:"true" json:"keyShape"`
    
 // The Key's current state.
    LifecycleState KeyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
    
 // The date and time this was created, in the format defined by RFC3339.
 // Example: `2016-08-25T21:10:29.600Z`
    TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
    
 // The OCID of the Vault containing this resource.
    VaultId *string `mandatory:"true" json:"vaultId"`
    
 // Usage of predefined tag keys. These predefined keys are scoped to namespaces.
 // Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
    DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
    
 // Simple key-value pair that is applied without any predefined name, type or scope.
 // Exists for cross-compatibility only.
 // Example: `{"bar-key": "value"}`
    FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m Key) String() string {
    return common.PointerString(m)
}


// KeyLifecycleStateEnum Enum with underlying type: string
type KeyLifecycleStateEnum string

// Set of constants representing the allowable values for KeyLifecycleState
const (
    KeyLifecycleStateCreating KeyLifecycleStateEnum = "CREATING"
    KeyLifecycleStateEnabling KeyLifecycleStateEnum = "ENABLING"
    KeyLifecycleStateEnabled KeyLifecycleStateEnum = "ENABLED"
    KeyLifecycleStateDisabling KeyLifecycleStateEnum = "DISABLING"
    KeyLifecycleStateDisabled KeyLifecycleStateEnum = "DISABLED"
    KeyLifecycleStateDeleting KeyLifecycleStateEnum = "DELETING"
    KeyLifecycleStateDeleted KeyLifecycleStateEnum = "DELETED"
)

var mappingKeyLifecycleState = map[string]KeyLifecycleStateEnum { 
    "CREATING": KeyLifecycleStateCreating,
    "ENABLING": KeyLifecycleStateEnabling,
    "ENABLED": KeyLifecycleStateEnabled,
    "DISABLING": KeyLifecycleStateDisabling,
    "DISABLED": KeyLifecycleStateDisabled,
    "DELETING": KeyLifecycleStateDeleting,
    "DELETED": KeyLifecycleStateDeleted,
}

// GetKeyLifecycleStateEnumValues Enumerates the set of values for KeyLifecycleState
func GetKeyLifecycleStateEnumValues() []KeyLifecycleStateEnum {
   values := make([]KeyLifecycleStateEnum, 0)
   for _, v := range mappingKeyLifecycleState {
       values = append(values, v)
   }
   return values
}



