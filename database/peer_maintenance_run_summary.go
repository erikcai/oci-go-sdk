// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PeerMaintenanceRunSummary Details of the Data Guard association's peer container database maintenance run.
type PeerMaintenanceRunSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the peer.
	Id *string `mandatory:"false" json:"id"`

	// The date and time the peer maintenance run is scheduled to begin.
	TimeScheduled *common.SDKTime `mandatory:"false" json:"timeScheduled"`

	// The ID of the target resource being updated by the peer maintenance run.
	TargetResourceId *string `mandatory:"false" json:"targetResourceId"`
}

func (m PeerMaintenanceRunSummary) String() string {
	return common.PointerString(m)
}
