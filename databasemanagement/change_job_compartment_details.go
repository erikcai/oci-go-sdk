// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management Service APIs.
//
// This file contains the customer facing APIs for Database Management service.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// ChangeJobCompartmentDetails Change the compartment of the Job.
type ChangeJobCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the
	// compartment into which the resource should be moved.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeJobCompartmentDetails) String() string {
	return common.PointerString(m)
}
