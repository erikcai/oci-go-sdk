// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RestartInstanceDetails The representation of RestartInstanceDetails
type RestartInstanceDetails struct {

	// The InnoDB shutdown mode for MySQLaaS Instance, following the option
	// "innodb_fast_shutdown (https://dev.mysql.com/doc/refman/en/innodb-parameters.html#sysvar_innodb_fast_shutdown)".
	ShutdownType InnoDbShutdownModeEnum `mandatory:"true" json:"shutdownType"`
}

func (m RestartInstanceDetails) String() string {
	return common.PointerString(m)
}
