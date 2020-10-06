// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GgsControlPlane API
//
// The GoldenGate Control Plane API specifying the operations and metadata for creating and maintaining the control infrastructure
// related to operating an OCI native Oracle managed GoldenGate service.
//

package goldengate

import (
	"github.com/oracle/oci-go-sdk/v26/common"
)

// UserNamePasswordPair The definition of a UserName/Password Pair.
type UserNamePasswordPair struct {

	// The UserName.
	UserName *string `mandatory:"true" json:"userName"`

	// The password.
	Password *string `mandatory:"true" json:"password"`
}

func (m UserNamePasswordPair) String() string {
	return common.PointerString(m)
}
