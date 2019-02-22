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

// ResetPeriod The reset period for the budget. We will start with MONTHLY and look into QUARTERLY and maybe ANNUAL post-MVP.
type ResetPeriod struct {
}

func (m ResetPeriod) String() string {
	return common.PointerString(m)
}
