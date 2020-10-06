// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// KAM API
//
// description: |
//   Kubernetes Add-on Manager API for installing, uninstalling and upgrading
//   OKE add-ons or Marketplace applications on an OKE cluster
//

package kam

import (
	"github.com/oracle/oci-go-sdk/v26/common"
)

// KamOverride KAM charts come with default configuration values. This may be used to
// override the defaults for the KAM chart
type KamOverride struct {

	// KAM chart property to override
	Name *string `mandatory:"false" json:"name"`

	// KAM chart property value to override
	Value *string `mandatory:"false" json:"value"`
}

func (m KamOverride) String() string {
	return common.PointerString(m)
}
