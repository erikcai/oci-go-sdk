// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// CreateCloudAutonomousVmClusterDetails Details for the create cloud Autonomous VM cluster operation.
type CreateCloudAutonomousVmClusterDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the availability domain that the Cloud Autonomous VM cluster is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the cloud Autonomous VM Cluster is associated with.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The user-friendly name for the cloud Autonomous VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
	CloudExadataInfrastructureId *string `mandatory:"true" json:"cloudExadataInfrastructureId"`

	// User defined description of the Cloud Autonomous Vm cluster.
	Description *string `mandatory:"false" json:"description"`

	// The number of nodes in Autonomous VM Cluster.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm), this attribute must be null because the attribute is already set at the
	// Autonomous Exadata Infrastructure level. When using shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`.
	LicenseModel CreateCloudAutonomousVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateCloudAutonomousVmClusterDetails) String() string {
	return common.PointerString(m)
}

// CreateCloudAutonomousVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type CreateCloudAutonomousVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for CreateCloudAutonomousVmClusterDetailsLicenseModelEnum
const (
	CreateCloudAutonomousVmClusterDetailsLicenseModelLicenseIncluded     CreateCloudAutonomousVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	CreateCloudAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense CreateCloudAutonomousVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateCloudAutonomousVmClusterDetailsLicenseModel = map[string]CreateCloudAutonomousVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateCloudAutonomousVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateCloudAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetCreateCloudAutonomousVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for CreateCloudAutonomousVmClusterDetailsLicenseModelEnum
func GetCreateCloudAutonomousVmClusterDetailsLicenseModelEnumValues() []CreateCloudAutonomousVmClusterDetailsLicenseModelEnum {
	values := make([]CreateCloudAutonomousVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingCreateCloudAutonomousVmClusterDetailsLicenseModel {
		values = append(values, v)
	}
	return values
}
