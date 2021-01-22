// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package ocas

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// DetectLanguageSentimentsResult Result of sentiments detect call.
type DetectLanguageSentimentsResult struct {

	// List of aspects
	Aspects []SentimentAspect `mandatory:"true" json:"aspects"`
}

func (m DetectLanguageSentimentsResult) String() string {
	return common.PointerString(m)
}
