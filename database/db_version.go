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

// DbVersion The Oracle Database version details.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type DbVersion struct {

	// A valid Oracle Database version.
	Version *string `mandatory:"true" json:"version"`

	// True if the Oracle Database software is a preview version.
	IsPreviewDbVersion *bool `mandatory:"true" json:"isPreviewDbVersion"`

	// List of the DB system StorageManagement summary field.
	StorageManagement []StorageManagementFieldSummary `mandatory:"true" json:"storageManagement"`

	// List of database service resources.
	DatabaseServiceResources []DatabaseServiceResource `mandatory:"true" json:"databaseServiceResources"`
}

func (m DbVersion) String() string {
	return common.PointerString(m)
}
