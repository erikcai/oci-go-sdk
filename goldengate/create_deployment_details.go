// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GgsControlPlane API
//
// The GoldenGate Control Plane API specifying the operations and metadata for creating and maintaining the control infrastructure
// related to operating an OCI native Oracle managed GoldenGate service.
//

package goldengate

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDeploymentDetails The information about new Deployment.
type CreateDeploymentDetails struct {

	// Object Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The unique identifier for the "type" of this resource.  The value of this "typeKey" will determine the Json Schema structure of the "typeData" property. As an example, a GGS Deployment Resource is a generic resource representing any 'kind' of 'deployment', such as a particular version of Oracle GoldenGate, or maybe a particular version of Oracle Streaming Analytics.  It is expected that each 'kind' will have some information that it needs to do its task. This information is really of no-use to the GGS ControlPlane nor the GGS DataPlane, but it does need to be passed to the deploymed container (similar in concept to "CloudInit" for a compute instance). There is only a single deployment kind being delivered at this point.  It has the value "UUID_1D7CAE6C_4CB0_3ABD_97E2_CDC1F558B638" and represents the released version of Oracle GoldenGate.
	TypeKey CreateDeploymentDetailsTypeKeyEnum `mandatory:"true" json:"typeKey"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup being referenced.
	BackupId *string `mandatory:"false" json:"backupId"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	Fqdn *string `mandatory:"false" json:"fqdn"`

	// Array of Network Security Group OCIDs.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// True if this object is publicly available.
	IsPublic *bool `mandatory:"false" json:"isPublic"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	PublicFqdn *string `mandatory:"false" json:"publicFqdn"`

	TypeData *Ogg20cTypeData `mandatory:"false" json:"typeData"`
}

func (m CreateDeploymentDetails) String() string {
	return common.PointerString(m)
}

// CreateDeploymentDetailsTypeKeyEnum Enum with underlying type: string
type CreateDeploymentDetailsTypeKeyEnum string

// Set of constants representing the allowable values for CreateDeploymentDetailsTypeKeyEnum
const (
	CreateDeploymentDetailsTypeKeyUuid1d7cae6c4cb03abd97e2Cdc1f558b638 CreateDeploymentDetailsTypeKeyEnum = "UUID_1D7CAE6C_4CB0_3ABD_97E2_CDC1F558B638"
)

var mappingCreateDeploymentDetailsTypeKey = map[string]CreateDeploymentDetailsTypeKeyEnum{
	"UUID_1D7CAE6C_4CB0_3ABD_97E2_CDC1F558B638": CreateDeploymentDetailsTypeKeyUuid1d7cae6c4cb03abd97e2Cdc1f558b638,
}

// GetCreateDeploymentDetailsTypeKeyEnumValues Enumerates the set of values for CreateDeploymentDetailsTypeKeyEnum
func GetCreateDeploymentDetailsTypeKeyEnumValues() []CreateDeploymentDetailsTypeKeyEnum {
	values := make([]CreateDeploymentDetailsTypeKeyEnum, 0)
	for _, v := range mappingCreateDeploymentDetailsTypeKey {
		values = append(values, v)
	}
	return values
}
