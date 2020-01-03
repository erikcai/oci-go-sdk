// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DataCatalog API
//
// A description of the DataCatalog API
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeCatalogCompartmentDetails The representation of ChangeCatalogCompartmentDetails
type ChangeCatalogCompartmentDetails struct {

	// The identifier of the Compartment into which the resource should be moved.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeCatalogCompartmentDetails) String() string {
	return common.PointerString(m)
}
