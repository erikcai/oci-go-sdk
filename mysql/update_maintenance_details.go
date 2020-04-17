// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateMaintenanceDetails The Maintenance Policy for the DB System.
type UpdateMaintenanceDetails struct {

	// The start of the 2 hour maintenance window.
	// This string is of the format: "{day-of-week} {time-of-day}".
	// "{day-of-week}" is a case-insensitive string like "mon", "tue", &c.
	// "{time-of-day}" is the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero.
	WindowStartTime *string `mandatory:"false" json:"windowStartTime"`
}

func (m UpdateMaintenanceDetails) String() string {
	return common.PointerString(m)
}
