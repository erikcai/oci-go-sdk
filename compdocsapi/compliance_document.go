// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compliance Documents Service API
//
// API for downloading Oracle Cloud Infrastructure compliance documents from the Oracle Cloud Infrastructure Console.
//

package compdocsapi

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// ComplianceDocument A compliance document that exists in the tenancy.
type ComplianceDocument struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compliance document, which is assigned
	// when you create the document as an Oracle Cloud Infrastructure resource and is immutable.
	Id *string `mandatory:"true" json:"id"`

	// A friendly name or title for the compliance document. You cannot update this value later.
	// Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	// The date and time the compliance document was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current lifecycle state of the compliance document.
	LifecycleState ComplianceDocumentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The file name of the compliance document.
	DocumentFileName *string `mandatory:"true" json:"documentFileName"`

	// The version number of the compliance document.
	Version *int `mandatory:"true" json:"version"`

	// The type of compliance document. For definitions of supported types of compliance documents, see Types of Compliance Documents (https://docs.cloud.oracle.com/en-us/iaas/Content/ComplianceDocuments/Concepts/compliancedocsoverview.htm#DocTypes).
	Type ComplianceDocumentTypeEnum `mandatory:"true" json:"type"`

	// The information technology infrastructure platform, or set of services, to which the compliance document belongs. A platform
	// can also be described as an environment or a business pillar. For definitions of supported environments, see Types of Environments (https://docs.cloud.oracle.com/en-us/iaas/Content/ComplianceDocuments/Concepts/compliancedocsoverview.htm#EnvironmentTypes).
	Platform ComplianceDocumentPlatformEnum `mandatory:"true" json:"platform"`

	// The date and time the compliance document was last updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m ComplianceDocument) String() string {
	return common.PointerString(m)
}

// ComplianceDocumentLifecycleStateEnum Enum with underlying type: string
type ComplianceDocumentLifecycleStateEnum string

// Set of constants representing the allowable values for ComplianceDocumentLifecycleStateEnum
const (
	ComplianceDocumentLifecycleStateActive   ComplianceDocumentLifecycleStateEnum = "ACTIVE"
	ComplianceDocumentLifecycleStateInactive ComplianceDocumentLifecycleStateEnum = "INACTIVE"
)

var mappingComplianceDocumentLifecycleState = map[string]ComplianceDocumentLifecycleStateEnum{
	"ACTIVE":   ComplianceDocumentLifecycleStateActive,
	"INACTIVE": ComplianceDocumentLifecycleStateInactive,
}

// GetComplianceDocumentLifecycleStateEnumValues Enumerates the set of values for ComplianceDocumentLifecycleStateEnum
func GetComplianceDocumentLifecycleStateEnumValues() []ComplianceDocumentLifecycleStateEnum {
	values := make([]ComplianceDocumentLifecycleStateEnum, 0)
	for _, v := range mappingComplianceDocumentLifecycleState {
		values = append(values, v)
	}
	return values
}

// ComplianceDocumentTypeEnum Enum with underlying type: string
type ComplianceDocumentTypeEnum string

// Set of constants representing the allowable values for ComplianceDocumentTypeEnum
const (
	ComplianceDocumentTypeSod          ComplianceDocumentTypeEnum = "SOD"
	ComplianceDocumentTypeAttestation  ComplianceDocumentTypeEnum = "ATTESTATION"
	ComplianceDocumentTypeBridgeletter ComplianceDocumentTypeEnum = "BRIDGELETTER"
	ComplianceDocumentTypePentest      ComplianceDocumentTypeEnum = "PENTEST"
	ComplianceDocumentTypeAudit        ComplianceDocumentTypeEnum = "AUDIT"
	ComplianceDocumentTypeCertificate  ComplianceDocumentTypeEnum = "CERTIFICATE"
	ComplianceDocumentTypeSoc3         ComplianceDocumentTypeEnum = "SOC3"
	ComplianceDocumentTypeOther        ComplianceDocumentTypeEnum = "OTHER"
)

var mappingComplianceDocumentType = map[string]ComplianceDocumentTypeEnum{
	"SOD":          ComplianceDocumentTypeSod,
	"ATTESTATION":  ComplianceDocumentTypeAttestation,
	"BRIDGELETTER": ComplianceDocumentTypeBridgeletter,
	"PENTEST":      ComplianceDocumentTypePentest,
	"AUDIT":        ComplianceDocumentTypeAudit,
	"CERTIFICATE":  ComplianceDocumentTypeCertificate,
	"SOC3":         ComplianceDocumentTypeSoc3,
	"OTHER":        ComplianceDocumentTypeOther,
}

// GetComplianceDocumentTypeEnumValues Enumerates the set of values for ComplianceDocumentTypeEnum
func GetComplianceDocumentTypeEnumValues() []ComplianceDocumentTypeEnum {
	values := make([]ComplianceDocumentTypeEnum, 0)
	for _, v := range mappingComplianceDocumentType {
		values = append(values, v)
	}
	return values
}

// ComplianceDocumentPlatformEnum Enum with underlying type: string
type ComplianceDocumentPlatformEnum string

// Set of constants representing the allowable values for ComplianceDocumentPlatformEnum
const (
	ComplianceDocumentPlatformOciedgeservices ComplianceDocumentPlatformEnum = "OCIEDGESERVICES"
	ComplianceDocumentPlatformOci             ComplianceDocumentPlatformEnum = "OCI"
	ComplianceDocumentPlatformPaas            ComplianceDocumentPlatformEnum = "PAAS"
	ComplianceDocumentPlatformCloudconsole    ComplianceDocumentPlatformEnum = "CLOUDCONSOLE"
	ComplianceDocumentPlatformOmcs            ComplianceDocumentPlatformEnum = "OMCS"
	ComplianceDocumentPlatformOciCIaas        ComplianceDocumentPlatformEnum = "OCI_C_IAAS"
	ComplianceDocumentPlatformOther           ComplianceDocumentPlatformEnum = "OTHER"
)

var mappingComplianceDocumentPlatform = map[string]ComplianceDocumentPlatformEnum{
	"OCIEDGESERVICES": ComplianceDocumentPlatformOciedgeservices,
	"OCI":             ComplianceDocumentPlatformOci,
	"PAAS":            ComplianceDocumentPlatformPaas,
	"CLOUDCONSOLE":    ComplianceDocumentPlatformCloudconsole,
	"OMCS":            ComplianceDocumentPlatformOmcs,
	"OCI_C_IAAS":      ComplianceDocumentPlatformOciCIaas,
	"OTHER":           ComplianceDocumentPlatformOther,
}

// GetComplianceDocumentPlatformEnumValues Enumerates the set of values for ComplianceDocumentPlatformEnum
func GetComplianceDocumentPlatformEnumValues() []ComplianceDocumentPlatformEnum {
	values := make([]ComplianceDocumentPlatformEnum, 0)
	for _, v := range mappingComplianceDocumentPlatform {
		values = append(values, v)
	}
	return values
}
