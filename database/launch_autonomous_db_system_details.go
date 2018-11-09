// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LaunchAutonomousDbSystemDetails Describes the input parameters to launch a new autonomous DB System.
type LaunchAutonomousDbSystemDetails struct {

	// The Oracle Cloud ID (OCID) of the compartment the Autonomous DB System  belongs in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Availability Domain where the Autonomous DB System is located.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the subnet the Autonomous DB System is associated with.
	// **Subnet Restrictions:**
	// - For Exadata DB Systems, do not use a subnet that overlaps with 192.168.128.0/20
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and backup subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The shape of the Autonomous DB System. The shape determines resources allocated to the DB System - CPU cores, memory and storage. To get a list of shapes, use the ListDbSystemShapes operation.
	Shape *string `mandatory:"true" json:"shape"`

	MaintenanceWindowDetails *AutonomousDbSystemMaintenanceWindow `mandatory:"true" json:"maintenanceWindowDetails"`

	// The user-friendly name for the Autonomous DB System. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The host name for the Autonomous DB System. The host name must begin with an alphabetic character and
	// can contain a maximum of 30 alphanumeric characters, including hyphens (-).
	// The maximum length of the combined hostname and domain is 63 characters.
	// **Note:** The hostname must be unique within the subnet. If it is not unique,
	// the DB System will fail to provision.
	Hostname *string `mandatory:"false" json:"hostname"`

	// A domain name used for the Autonomous DB System. If the Oracle-provided Internet and VCN
	// Resolver is enabled for the specified subnet, the domain name for the subnet is used
	// (don't provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted.
	Domain *string `mandatory:"false" json:"domain"`

	// The Oracle license model that applies to all the databases on the DB System. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel LaunchAutonomousDbSystemDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m LaunchAutonomousDbSystemDetails) String() string {
	return common.PointerString(m)
}

// LaunchAutonomousDbSystemDetailsLicenseModelEnum Enum with underlying type: string
type LaunchAutonomousDbSystemDetailsLicenseModelEnum string

// Set of constants representing the allowable values for LaunchAutonomousDbSystemDetailsLicenseModelEnum
const (
	LaunchAutonomousDbSystemDetailsLicenseModelLicenseIncluded     LaunchAutonomousDbSystemDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	LaunchAutonomousDbSystemDetailsLicenseModelBringYourOwnLicense LaunchAutonomousDbSystemDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingLaunchAutonomousDbSystemDetailsLicenseModel = map[string]LaunchAutonomousDbSystemDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       LaunchAutonomousDbSystemDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": LaunchAutonomousDbSystemDetailsLicenseModelBringYourOwnLicense,
}

// GetLaunchAutonomousDbSystemDetailsLicenseModelEnumValues Enumerates the set of values for LaunchAutonomousDbSystemDetailsLicenseModelEnum
func GetLaunchAutonomousDbSystemDetailsLicenseModelEnumValues() []LaunchAutonomousDbSystemDetailsLicenseModelEnum {
	values := make([]LaunchAutonomousDbSystemDetailsLicenseModelEnum, 0)
	for _, v := range mappingLaunchAutonomousDbSystemDetailsLicenseModel {
		values = append(values, v)
	}
	return values
}
