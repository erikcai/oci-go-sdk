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


    
 // KeySummary The representation of KeySummary
type KeySummary struct {
    
 // The OCID of the Compartment containing this resource.
    CompartmentId *string `mandatory:"true" json:"compartmentId"`
    
 // A user-friendly name. Does not have to be unique, and it's changeable.
 // Avoid entering confidential information.
    DisplayName *string `mandatory:"true" json:"displayName"`
    
 // The OCID of the resource.
    Id *string `mandatory:"true" json:"id"`
    
 // The Key's current state.
    LifecycleState KeySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
    
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

func (m KeySummary) String() string {
    return common.PointerString(m)
}


// KeySummaryLifecycleStateEnum Enum with underlying type: string
type KeySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for KeySummaryLifecycleState
const (
    KeySummaryLifecycleStateCreating KeySummaryLifecycleStateEnum = "CREATING"
    KeySummaryLifecycleStateEnabling KeySummaryLifecycleStateEnum = "ENABLING"
    KeySummaryLifecycleStateEnabled KeySummaryLifecycleStateEnum = "ENABLED"
    KeySummaryLifecycleStateDisabling KeySummaryLifecycleStateEnum = "DISABLING"
    KeySummaryLifecycleStateDisabled KeySummaryLifecycleStateEnum = "DISABLED"
    KeySummaryLifecycleStateDeleting KeySummaryLifecycleStateEnum = "DELETING"
    KeySummaryLifecycleStateDeleted KeySummaryLifecycleStateEnum = "DELETED"
)

var mappingKeySummaryLifecycleState = map[string]KeySummaryLifecycleStateEnum { 
    "CREATING": KeySummaryLifecycleStateCreating,
    "ENABLING": KeySummaryLifecycleStateEnabling,
    "ENABLED": KeySummaryLifecycleStateEnabled,
    "DISABLING": KeySummaryLifecycleStateDisabling,
    "DISABLED": KeySummaryLifecycleStateDisabled,
    "DELETING": KeySummaryLifecycleStateDeleting,
    "DELETED": KeySummaryLifecycleStateDeleted,
}

// GetKeySummaryLifecycleStateEnumValues Enumerates the set of values for KeySummaryLifecycleState
func GetKeySummaryLifecycleStateEnumValues() []KeySummaryLifecycleStateEnum {
   values := make([]KeySummaryLifecycleStateEnum, 0)
   for _, v := range mappingKeySummaryLifecycleState {
       values = append(values, v)
   }
   return values
}



