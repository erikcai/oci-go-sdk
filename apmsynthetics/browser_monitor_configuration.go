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

// BrowserMonitorConfiguration Configuration for BROWSER monitor type.
type BrowserMonitorConfiguration struct {

	// If isFailureRetried enabled, then if call is failed then it will be retried.
	IsFailureRetried *bool `mandatory:"false" json:"isFailureRetried"`

	// If validation enabled, then call will fail for certificate errors.
	IsCertificateValidationEnabled *bool `mandatory:"false" json:"isCertificateValidationEnabled"`

	// Verify all the search strings present in response.
	// If any search string is not present in the response, then it will be considered as a failure.
	VerifyTexts []VerifyText `mandatory:"false" json:"verifyTexts"`
}

//GetIsFailureRetried returns IsFailureRetried
func (m BrowserMonitorConfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

func (m BrowserMonitorConfiguration) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m BrowserMonitorConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBrowserMonitorConfiguration BrowserMonitorConfiguration
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeBrowserMonitorConfiguration
	}{
		"BROWSER_CONFIG",
		(MarshalTypeBrowserMonitorConfiguration)(m),
	}

	return json.Marshal(&s)
}
