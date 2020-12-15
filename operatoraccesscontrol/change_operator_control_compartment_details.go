// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// ChangeOperatorControlCompartmentDetails Details of the new compartment to contain the operator control.
type ChangeOperatorControlCompartmentDetails struct {

	// The OCID of the new compartment to contain the operator contol.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m ChangeOperatorControlCompartmentDetails) String() string {
	return common.PointerString(m)
}
