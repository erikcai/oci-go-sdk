// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compliance Documents Service API
//
// API for downloading Oracle Cloud Infrastructure compliance documents from the Oracle Cloud Infrastructure Console.
//

package compdocsapi

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// ComplianceDocumentSummary A summary representation of the compliance document.
type ComplianceDocumentSummary struct {

	// A unique identifier for the document that is assigned when you create
	// the document as an Oracle Cloud Infrastructure resource and is immutable.
	Id *string `mandatory:"true" json:"id"`

	// A friendly name or title for the compliance document. You cannot update this value later.
	// Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	// The current lifecycle state of the compliance document.
	LifecycleState ComplianceDocumentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The file name of the compliance document.
	DocumentFileName *string `mandatory:"true" json:"documentFileName"`

	// The version number of the compliance document.
	Version *int `mandatory:"true" json:"version"`

	// The type of compliance document.
	Type ComplianceDocumentTypeEnum `mandatory:"true" json:"type"`

	// The environment, also known as platform or business pillar, to which the compliance document belongs.
	Platform ComplianceDocumentPlatformEnum `mandatory:"true" json:"platform"`

	// The date and time the compliance document was last updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The date and time the compliance document was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m ComplianceDocumentSummary) String() string {
	return common.PointerString(m)
}
