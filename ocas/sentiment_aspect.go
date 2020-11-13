// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v28/common"
)

// SentimentAspect aspect object
type SentimentAspect struct {

	// The number of Unicode code points preceding this entity in the submitted text
	Offset *int `mandatory:"false" json:"offset"`

	// length of text
	Length *int `mandatory:"false" json:"length"`

	// aspect text
	Text *string `mandatory:"false" json:"text"`

	// The highest-score sentiment for the aspect text.
	Sentiment *string `mandatory:"false" json:"sentiment"`

	// scores for each sentiment
	// Example: "{\"positive\": 1.0, \"negative\": 0.0}"
	Scores map[string]float64 `mandatory:"false" json:"scores"`
}

func (m SentimentAspect) String() string {
	return common.PointerString(m)
}
