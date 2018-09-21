// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// EventsControlService API
//
// This service exposes APIs to create, update and delete Rules. Rules are used to tap into the Events stream.
//

package cloudevents

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ObjectStorageServiceAction deliver to an Oracle Object Storage bucket
type ObjectStorageServiceAction struct {

	// OCID of the action
	Id *string `mandatory:"true" json:"id"`

	// Service-generated message about the current state of this rule
	LifecycleMessage *string `mandatory:"true" json:"lifecycleMessage"`

	// The top-level namespace of buckets
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// The name of the bucket
	BucketName *string `mandatory:"false" json:"bucketName"`

	// specifies current state of the action
	LifecycleState ActionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m ObjectStorageServiceAction) GetId() *string {
	return m.Id
}

//GetLifecycleMessage returns LifecycleMessage
func (m ObjectStorageServiceAction) GetLifecycleMessage() *string {
	return m.LifecycleMessage
}

//GetLifecycleState returns LifecycleState
func (m ObjectStorageServiceAction) GetLifecycleState() ActionLifecycleStateEnum {
	return m.LifecycleState
}

func (m ObjectStorageServiceAction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ObjectStorageServiceAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageServiceAction ObjectStorageServiceAction
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeObjectStorageServiceAction
	}{
		"OBJECTSTORAGE",
		(MarshalTypeObjectStorageServiceAction)(m),
	}

	return json.Marshal(&s)
}
