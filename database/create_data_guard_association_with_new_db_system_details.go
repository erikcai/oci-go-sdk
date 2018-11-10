// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDataGuardAssociationWithNewDbSystemDetails The configuration details for creating a Data Guard association to an existing database. A new DbSystem will be launched for standby database.
type CreateDataGuardAssociationWithNewDbSystemDetails struct {

	// A strong password for the `SYS`, `SYSTEM`, and `PDB Admin` users to apply during standby creation.
	// The password must contain no fewer than nine characters and include:
	// * At least two uppercase characters.
	// * At least two lowercase characters.
	// * At least two numeric characters.
	// * At least two special characters. Valid special characters include "_", "#", and "-" only.
	// **The password MUST be the same as the primary admin password.**
	DatabaseAdminPassword *string `mandatory:"true" json:"databaseAdminPassword"`

	// The user-friendly name for the DB System to create the standby database on. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The name of the Availability Domain that the standby database DB System will be located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID of the subnet the DB System is associated with.
	// **Subnet Restrictions:**
	// - For 1- and 2-node RAC DB Systems, do not use a subnet that overlaps with 192.168.16.16/28
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and backup subnet.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The host name for the DB Node.
	Hostname *string `mandatory:"false" json:"hostname"`

	// The protection mode to set up between the primary and standby databases. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	// **IMPORTANT** - The only protection mode currently supported by the Database Service is MAXIMUM_PERFORMANCE.
	ProtectionMode CreateDataGuardAssociationDetailsProtectionModeEnum `mandatory:"true" json:"protectionMode"`

	// The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:
	// * MAXIMUM_AVAILABILITY - SYNC or FASTSYNC
	// * MAXIMUM_PERFORMANCE - ASYNC
	// * MAXIMUM_PROTECTION - SYNC
	// For more information, see
	// Redo Transport Services (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400)
	// in the Oracle Data Guard documentation.
	// **IMPORTANT** - The only transport type currently supported by the Database Service is ASYNC.
	TransportType CreateDataGuardAssociationDetailsTransportTypeEnum `mandatory:"true" json:"transportType"`
}

//GetDatabaseAdminPassword returns DatabaseAdminPassword
func (m CreateDataGuardAssociationWithNewDbSystemDetails) GetDatabaseAdminPassword() *string {
	return m.DatabaseAdminPassword
}

//GetProtectionMode returns ProtectionMode
func (m CreateDataGuardAssociationWithNewDbSystemDetails) GetProtectionMode() CreateDataGuardAssociationDetailsProtectionModeEnum {
	return m.ProtectionMode
}

//GetTransportType returns TransportType
func (m CreateDataGuardAssociationWithNewDbSystemDetails) GetTransportType() CreateDataGuardAssociationDetailsTransportTypeEnum {
	return m.TransportType
}

func (m CreateDataGuardAssociationWithNewDbSystemDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDataGuardAssociationWithNewDbSystemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDataGuardAssociationWithNewDbSystemDetails CreateDataGuardAssociationWithNewDbSystemDetails
	s := struct {
		DiscriminatorParam string `json:"creationType"`
		MarshalTypeCreateDataGuardAssociationWithNewDbSystemDetails
	}{
		"NewDbSystem",
		(MarshalTypeCreateDataGuardAssociationWithNewDbSystemDetails)(m),
	}

	return json.Marshal(&s)
}
