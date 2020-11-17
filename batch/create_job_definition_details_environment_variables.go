// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Batch Service
//
// This is a Oracle Batch Service. You can find out more about at
// wiki (https://confluence.oraclecorp.com/confluence/display/C9QA/OCI+Batch+Service+-+Core+Functionality+Technical+Design+and+Explanation+for+Phase+I).
//

package batch

import (
	"github.com/oracle/oci-go-sdk/v29/common"
)

// CreateJobDefinitionDetailsEnvironmentVariables The representation of CreateJobDefinitionDetailsEnvironmentVariables
type CreateJobDefinitionDetailsEnvironmentVariables struct {
	Name *string `mandatory:"false" json:"name"`

	Value *string `mandatory:"false" json:"value"`
}

func (m CreateJobDefinitionDetailsEnvironmentVariables) String() string {
	return common.PointerString(m)
}
