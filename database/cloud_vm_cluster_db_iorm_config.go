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

// CloudVmClusterDbIormConfig IORM Config setting response for this database
type CloudVmClusterDbIormConfig struct {

	// Database Name. For default DbPlan, the dbName will always be `default`
	DbName *string `mandatory:"false" json:"dbName"`

	// Relative priority of a database
	Share *int64 `mandatory:"false" json:"share"`

	// Flash Cache limit, internally configured based on shares
	FlashCacheLimit *string `mandatory:"false" json:"flashCacheLimit"`
}

func (m CloudVmClusterDbIormConfig) String() string {
	return common.PointerString(m)
}
