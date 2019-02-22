// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// BudgetsControlPlane API
//
// A description of the BudgetsControlPlane API
//

package budget

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AlertType ACTUAL means the alert will trigger based on actual usage.
// FORECAST means the alert will trigger based on predicted usage.
type AlertType struct {
}

func (m AlertType) String() string {
	return common.PointerString(m)
}
