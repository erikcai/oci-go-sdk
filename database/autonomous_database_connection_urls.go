// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousDatabaseConnectionUrls The URLs for accessing Oracle Application Express (APEX) and SQL Developer Web with a browser from a Compute instance within your VCN or that has a direct connection to your VCN. Note that these URLs are provided by the console only for dedicated deployments (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm).
// Example: `{"sqlDevWebUrl": "https://<hostname>/ords...", "apexUrl", "https://<hostname>/ords..."}`
type AutonomousDatabaseConnectionUrls struct {

	// Oracle SQL Developer Web URL.
	SqlDevWebUrl *string `mandatory:"false" json:"sqlDevWebUrl"`

	// Oracle Application Express (APEX) URL.
	ApexUrl *string `mandatory:"false" json:"apexUrl"`

	// Oracle Machine Learning User Management URL.
	MachineLearningUserManagementUrl *string `mandatory:"false" json:"machineLearningUserManagementUrl"`
}

func (m AutonomousDatabaseConnectionUrls) String() string {
	return common.PointerString(m)
}
