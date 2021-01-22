// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Synthetic
//
// API for APM Synthetic service. Use this API to query Synthethetic scripts and monitors.
//

package apmsynthetics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v33/common"
)

// ScriptedBrowserMonitorConfiguration Configuration for SCRIPTED_BROWSER monitor type.
type ScriptedBrowserMonitorConfiguration struct {

	// If isFailureRetried enabled, then if call is failed then it will be retried.
	IsFailureRetried *bool `mandatory:"false" json:"isFailureRetried"`

	// If certificate validation enabled, then call will fail for certificate errors.
	IsCertificateValidationEnabled *bool `mandatory:"false" json:"isCertificateValidationEnabled"`
}

//GetIsFailureRetried returns IsFailureRetried
func (m ScriptedBrowserMonitorConfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

func (m ScriptedBrowserMonitorConfiguration) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ScriptedBrowserMonitorConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeScriptedBrowserMonitorConfiguration ScriptedBrowserMonitorConfiguration
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeScriptedBrowserMonitorConfiguration
	}{
		"SCRIPTED_BROWSER_CONFIG",
		(MarshalTypeScriptedBrowserMonitorConfiguration)(m),
	}

	return json.Marshal(&s)
}
