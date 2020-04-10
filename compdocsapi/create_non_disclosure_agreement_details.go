// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Compliance Documents Service API
//
// API for downloading Oracle Cloud Infrastructure compliance documents from the Oracle Cloud Infrastructure Console.
//

package compdocsapi

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateNonDisclosureAgreementDetails Details to use to create a new non-disclosure agreement for a particular compliance document.
type CreateNonDisclosureAgreementDetails struct {

	// The ID of the compliance document associated with the non-disclosure agreement.
	DocumentId *string `mandatory:"true" json:"documentId"`

	// The OCID of the compartment that contains the non-disclosure agreement.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m CreateNonDisclosureAgreementDetails) String() string {
	return common.PointerString(m)
}
