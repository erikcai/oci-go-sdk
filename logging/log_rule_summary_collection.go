// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// PublicLoggingControlplane API
//
// PublicLoggingControlplane API specification
//

package logging

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LogRuleSummaryCollection Log rule summary collection
type LogRuleSummaryCollection struct {

	// collection of LogRuleSummary objects
	Items []LogRuleSummary `mandatory:"true" json:"items"`
}

func (m LogRuleSummaryCollection) String() string {
	return common.PointerString(m)
}
