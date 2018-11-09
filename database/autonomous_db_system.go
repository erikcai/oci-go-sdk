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

// AutonomousDbSystem The representation of AutonomousDbSystem
type AutonomousDbSystem struct {

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
	LifecycleState AutonomousDbSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	MaintenanceWindow *AutonomousDbSystemMaintenanceWindow `mandatory:"true" json:"maintenanceWindow"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The Oracle license model that applies to all the databases on the DB System. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel AutonomousDbSystemLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

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

func (m AutonomousDbSystem) String() string {
	return common.PointerString(m)
}

// AutonomousDbSystemLifecycleStateEnum Enum with underlying type: string
type AutonomousDbSystemLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDbSystemLifecycleStateEnum
const (
	AutonomousDbSystemLifecycleStateProvisioning AutonomousDbSystemLifecycleStateEnum = "PROVISIONING"
	AutonomousDbSystemLifecycleStateAvailable    AutonomousDbSystemLifecycleStateEnum = "AVAILABLE"
	AutonomousDbSystemLifecycleStateUpdating     AutonomousDbSystemLifecycleStateEnum = "UPDATING"
	AutonomousDbSystemLifecycleStateTerminating  AutonomousDbSystemLifecycleStateEnum = "TERMINATING"
	AutonomousDbSystemLifecycleStateTerminated   AutonomousDbSystemLifecycleStateEnum = "TERMINATED"
	AutonomousDbSystemLifecycleStateFailed       AutonomousDbSystemLifecycleStateEnum = "FAILED"
)

var mappingAutonomousDbSystemLifecycleState = map[string]AutonomousDbSystemLifecycleStateEnum{
	"PROVISIONING": AutonomousDbSystemLifecycleStateProvisioning,
	"AVAILABLE":    AutonomousDbSystemLifecycleStateAvailable,
	"UPDATING":     AutonomousDbSystemLifecycleStateUpdating,
	"TERMINATING":  AutonomousDbSystemLifecycleStateTerminating,
	"TERMINATED":   AutonomousDbSystemLifecycleStateTerminated,
	"FAILED":       AutonomousDbSystemLifecycleStateFailed,
}

// GetAutonomousDbSystemLifecycleStateEnumValues Enumerates the set of values for AutonomousDbSystemLifecycleStateEnum
func GetAutonomousDbSystemLifecycleStateEnumValues() []AutonomousDbSystemLifecycleStateEnum {
	values := make([]AutonomousDbSystemLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDbSystemLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousDbSystemLicenseModelEnum Enum with underlying type: string
type AutonomousDbSystemLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousDbSystemLicenseModelEnum
const (
	AutonomousDbSystemLicenseModelLicenseIncluded     AutonomousDbSystemLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousDbSystemLicenseModelBringYourOwnLicense AutonomousDbSystemLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousDbSystemLicenseModel = map[string]AutonomousDbSystemLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousDbSystemLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousDbSystemLicenseModelBringYourOwnLicense,
}

// GetAutonomousDbSystemLicenseModelEnumValues Enumerates the set of values for AutonomousDbSystemLicenseModelEnum
func GetAutonomousDbSystemLicenseModelEnumValues() []AutonomousDbSystemLicenseModelEnum {
	values := make([]AutonomousDbSystemLicenseModelEnum, 0)
	for _, v := range mappingAutonomousDbSystemLicenseModel {
		values = append(values, v)
	}
	return values
}
