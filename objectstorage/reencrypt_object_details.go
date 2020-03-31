// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object Storage and Archive Storage APIs for managing buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ReencryptObjectDetails This reference describes the information used to re-encrypt the encryption keys associated with an object.
// You can only specify either a kmsKeyId or an sseCustomerKey in the request payload (not both).
// If the request payload is empty, the object is encrypted using the encryption mechanism specified for the
// bucket. The encryption mechanism can either be a master encryption key managed by Oracle or the Key Management service.
// - The sseCustomerKey field specifies the customer-provided encryption key that will be used to re-encrypt the
//   object and its chunks.
// - The sourceSSECustomerKey field specifies information about the customer-provided encryption key that is currently
//   associated with the object source. Specify a value for the sourceSSECustomerKey only if the object
//   is encrypted with a customer-provided encryption key.
type ReencryptObjectDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the master encryption key used to call the Key Management
	// service to re-encrypt the data encryption key associated with the object and its chunks. If the kmsKeyId value is
	// empty, whether null or an empty string, the API will perform re-encryption by using the kmsKeyId associated with the
	// bucket.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	SseCustomerKey *SseCustomerKeyDetails `mandatory:"false" json:"sseCustomerKey"`

	SourceSseCustomerKey *SseCustomerKeyDetails `mandatory:"false" json:"sourceSseCustomerKey"`
}

func (m ReencryptObjectDetails) String() string {
	return common.PointerString(m)
}