// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousPodBackupConfig Backup options for the Autonomous Pod.
type AutonomousPodBackupConfig struct {

	// Number of days between the current and the earliest point of recoverability covered by automatic backups.
	// This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window.
	// When the value is updated, it is applied to all existing automatic backups.
	RecoveryWindowInDays *int `mandatory:"false" json:"recoveryWindowInDays"`
}

func (m AutonomousPodBackupConfig) String() string {
	return common.PointerString(m)
}
