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

// VersionSummary A summary of the supported MySQL Versions families, and a list oftheir supported minor versions.
// To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized, talk to an administrator.
// If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type VersionSummary struct {

	// The list of supported MySQL Versions.
	Versions []Version `mandatory:"true" json:"versions"`

	// A descriptive summary of a group of versions.
	VersionFamily *string `mandatory:"false" json:"versionFamily"`
}

func (m VersionSummary) String() string {
	return common.PointerString(m)
}
