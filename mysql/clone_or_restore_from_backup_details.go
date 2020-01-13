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

// CloneOrRestoreFromBackupDetails Parameters detailing the backup from which to restore the system.
type CloneOrRestoreFromBackupDetails struct {

	// The OCID of the backup to use for the restoration procedure. This
	// backup must be taken from (TBD or abstractly compatible with?) the
	// dbSystem it is being restored to.
	BackupId *string `mandatory:"true" json:"backupId"`

	// The specific position to recover to. This value maybe one of the following formats:
	//   - datetime (RFC 3339 (https://tools.ietf.org/rfc/rfc3339): describing the point-in-time of recovery.
	//   - integer: innodb log sequence number.
	//   - uuid (GTID): specific transaction ID.
	LogPosition *string `mandatory:"false" json:"logPosition"`
}

func (m CloneOrRestoreFromBackupDetails) String() string {
	return common.PointerString(m)
}
