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

// AutonomousDbSystemSummary Infrastructure that enables the customer to run multiple Autonomous Dedicated Databases.
// For details about autonomous DB Systems, see:
// - Autonomous DB Systems (https://docs.us-phoenix-1.oraclecloud.com/Content/Database/Concepts/exaoverview.htm)
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
// For information about access control and compartments, see
// Overview of the Identity Service (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
// For information about Availability Domains, see
// Regions and Availability Domains (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/regions.htm).
// To get a list of Availability Domains, use the `ListAvailabilityDomains` operation
// in the Identity Service API.
type AutonomousDbSystemSummary struct {

	// The OCID of the Autonomous DB System.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Autonomous DB System.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the Availability Domain that the Autonomous DB System is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the subnet the Autonomous DB System is associated with.
	// **Subnet Restrictions:**
	// - For Exadata Autonomous DB Systems, do not use a subnet that overlaps with 192.168.128.0/20
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and backup subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The shape of the Autonomous DB System. The shape determines resources to allocate to the Autonomous DB system - CPU cores, memory and storage.
	Shape *string `mandatory:"true" json:"shape"`

	// The host name for the Autonomous DB Node.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The domain name for the Autonomous DB System.
	Domain *string `mandatory:"true" json:"domain"`

	// The current state of the Autonomous DB System.
	LifecycleState AutonomousDbSystemSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	MaintenanceWindow *AutonomousDbSystemMaintenanceWindow `mandatory:"true" json:"maintenanceWindow"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The Oracle license model that applies to all the databases on the DB System. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel AutonomousDbSystemSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The date and time the Autonomous DB System was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	LastMaintenanceRun *MaintenanceRun `mandatory:"false" json:"lastMaintenanceRun"`

	NextMaintenanceRun *MaintenanceRun `mandatory:"false" json:"nextMaintenanceRun"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AutonomousDbSystemSummary) String() string {
	return common.PointerString(m)
}

// AutonomousDbSystemSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousDbSystemSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDbSystemSummaryLifecycleStateEnum
const (
	AutonomousDbSystemSummaryLifecycleStateProvisioning AutonomousDbSystemSummaryLifecycleStateEnum = "PROVISIONING"
	AutonomousDbSystemSummaryLifecycleStateAvailable    AutonomousDbSystemSummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousDbSystemSummaryLifecycleStateUpdating     AutonomousDbSystemSummaryLifecycleStateEnum = "UPDATING"
	AutonomousDbSystemSummaryLifecycleStateTerminating  AutonomousDbSystemSummaryLifecycleStateEnum = "TERMINATING"
	AutonomousDbSystemSummaryLifecycleStateTerminated   AutonomousDbSystemSummaryLifecycleStateEnum = "TERMINATED"
	AutonomousDbSystemSummaryLifecycleStateFailed       AutonomousDbSystemSummaryLifecycleStateEnum = "FAILED"
)

var mappingAutonomousDbSystemSummaryLifecycleState = map[string]AutonomousDbSystemSummaryLifecycleStateEnum{
	"PROVISIONING": AutonomousDbSystemSummaryLifecycleStateProvisioning,
	"AVAILABLE":    AutonomousDbSystemSummaryLifecycleStateAvailable,
	"UPDATING":     AutonomousDbSystemSummaryLifecycleStateUpdating,
	"TERMINATING":  AutonomousDbSystemSummaryLifecycleStateTerminating,
	"TERMINATED":   AutonomousDbSystemSummaryLifecycleStateTerminated,
	"FAILED":       AutonomousDbSystemSummaryLifecycleStateFailed,
}

// GetAutonomousDbSystemSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousDbSystemSummaryLifecycleStateEnum
func GetAutonomousDbSystemSummaryLifecycleStateEnumValues() []AutonomousDbSystemSummaryLifecycleStateEnum {
	values := make([]AutonomousDbSystemSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDbSystemSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousDbSystemSummaryLicenseModelEnum Enum with underlying type: string
type AutonomousDbSystemSummaryLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousDbSystemSummaryLicenseModelEnum
const (
	AutonomousDbSystemSummaryLicenseModelLicenseIncluded     AutonomousDbSystemSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousDbSystemSummaryLicenseModelBringYourOwnLicense AutonomousDbSystemSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousDbSystemSummaryLicenseModel = map[string]AutonomousDbSystemSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousDbSystemSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousDbSystemSummaryLicenseModelBringYourOwnLicense,
}

// GetAutonomousDbSystemSummaryLicenseModelEnumValues Enumerates the set of values for AutonomousDbSystemSummaryLicenseModelEnum
func GetAutonomousDbSystemSummaryLicenseModelEnumValues() []AutonomousDbSystemSummaryLicenseModelEnum {
	values := make([]AutonomousDbSystemSummaryLicenseModelEnum, 0)
	for _, v := range mappingAutonomousDbSystemSummaryLicenseModel {
		values = append(values, v)
	}
	return values
}
