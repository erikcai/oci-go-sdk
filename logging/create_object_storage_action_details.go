// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PublicLoggingControlplane API
//
// PublicLoggingControlplane API specification
//

package logging

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateObjectStorageActionDetails Create an action that delivers to Object Storage.
type CreateObjectStorageActionDetails struct {

	// Whether or not this resource is currently enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// name of bucket
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The OCID of the resource.
	Id *string `mandatory:"false" json:"id"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// object name prefix
	ObjectNamePrefix *string `mandatory:"false" json:"objectNamePrefix"`
}

//GetId returns Id
func (m CreateObjectStorageActionDetails) GetId() *string {
	return m.Id
}

//GetIsEnabled returns IsEnabled
func (m CreateObjectStorageActionDetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

//GetDescription returns Description
func (m CreateObjectStorageActionDetails) GetDescription() *string {
	return m.Description
}

func (m CreateObjectStorageActionDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateObjectStorageActionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateObjectStorageActionDetails CreateObjectStorageActionDetails
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeCreateObjectStorageActionDetails
	}{
		"OBJECT_STORAGE",
		(MarshalTypeCreateObjectStorageActionDetails)(m),
	}

	return json.Marshal(&s)
}
