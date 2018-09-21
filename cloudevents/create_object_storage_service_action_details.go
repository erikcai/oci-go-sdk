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

// CreateObjectStorageServiceActionDetails deliver to an Oracle Object Storage bucket
type CreateObjectStorageServiceActionDetails struct {

	// whether or not this aciton is currently enabled
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The top-level namespace for buckets
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// The name of the bucket
	BucketName *string `mandatory:"true" json:"bucketName"`
}

//GetIsEnabled returns IsEnabled
func (m CreateObjectStorageServiceActionDetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

func (m CreateObjectStorageServiceActionDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateObjectStorageServiceActionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateObjectStorageServiceActionDetails CreateObjectStorageServiceActionDetails
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeCreateObjectStorageServiceActionDetails
	}{
		"OBJECTSTORAGE",
		(MarshalTypeCreateObjectStorageServiceActionDetails)(m),
	}

	return json.Marshal(&s)
}
