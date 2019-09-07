// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// PublicLoggingControlplane API
//
// PublicLoggingControlplane API specification
//

package publiclogging

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// OciService OCI service logging configuration.
type OciService struct {

	// Service generating log.
	Service *string `mandatory:"true" json:"service"`

	// The unique identifier of the resource emitting the log.
	Resource *string `mandatory:"true" json:"resource"`

	// Log object category.
	Category *string `mandatory:"true" json:"category"`

	// Log category parameters are stored here.
	Parameters map[string]string `mandatory:"false" json:"parameters"`
}

func (m OciService) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OciService) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOciService OciService
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeOciService
	}{
		"OCISERVICE",
		(MarshalTypeOciService)(m),
	}

	return json.Marshal(&s)
}
