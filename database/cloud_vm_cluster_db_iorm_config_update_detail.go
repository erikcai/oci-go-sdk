// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CloudVmClusterDbIormConfigUpdateDetail IORM Config setting request for this database
type CloudVmClusterDbIormConfigUpdateDetail struct {

	// Database Name. For updating default DbPlan, pass in dbName as `default`
	DbName *string `mandatory:"false" json:"dbName"`

	// Relative priority of a database
	Share *int64 `mandatory:"false" json:"share"`
}

func (m CloudVmClusterDbIormConfigUpdateDetail) String() string {
	return common.PointerString(m)
}
