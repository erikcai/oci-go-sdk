// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connectors API
//
// A description of the Connectors API
//

package sch

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// LoggingSourceDetails The logging source.
type LoggingSourceDetails struct {

	// The resources affected by this work request.
	LogSources []LogSource `mandatory:"true" json:"logSources"`
}

func (m LoggingSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m LoggingSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLoggingSourceDetails LoggingSourceDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeLoggingSourceDetails
	}{
		"logging",
		(MarshalTypeLoggingSourceDetails)(m),
	}

	return json.Marshal(&s)
}
