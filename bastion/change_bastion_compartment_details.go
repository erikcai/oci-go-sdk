// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciBastions API
//
// A description of the OciBastions API
//

package bastion

import (
	"github.com/oracle/oci-go-sdk/v26/common"
)

// ChangeBastionCompartmentDetails update compartment details.
type ChangeBastionCompartmentDetails struct {

	// the target compartment ocid which the bastion should be moved to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeBastionCompartmentDetails) String() string {
	return common.PointerString(m)
}
