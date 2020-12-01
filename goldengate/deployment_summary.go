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
	"github.com/oracle/oci-go-sdk/v30/common"
)

// DeploymentSummary Summary of the Deployment.
type DeploymentSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the deployment being referenced.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier for the "type" of this resource.  The value of this "typeKey" will determine the Json Schema structure of the "typeData" property. As an example, a GGS Deployment Resource is a generic resource representing any 'kind' of 'deployment', such as a particular version of Oracle GoldenGate, or maybe a particular version of Oracle Streaming Analytics.  It is expected that each 'kind' will have some information that it needs to do its task. This information is really of no-use to the GGS ControlPlane nor the GGS DataPlane, but it does need to be passed to the deploymed container (similar in concept to "CloudInit" for a compute instance). There is only a single deployment kind being delivered at this point.  It has the value "1D7CAE6C_4CB0_3ABD_97E2_CDC1F558B638" and represents the released version of Oracle GoldenGate.
	TypeKey *string `mandatory:"true" json:"typeKey"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Object Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// The time the Resource was created. Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Resource was updated.  Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Possible lifecycle states.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	TypeData *Ogg20cTypeData `mandatory:"false" json:"typeData"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	Fqdn *string `mandatory:"false" json:"fqdn"`

	// True if this object is publicly available.
	IsPublic *bool `mandatory:"false" json:"isPublic"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	PublicFqdn *string `mandatory:"false" json:"publicFqdn"`

	// The private IP address (in the customer's VCN) representing the access point for the associated endpoint service in the GGS Service VCN.
	PePrivateIp *string `mandatory:"false" json:"pePrivateIp"`

	// The URL of a resource.
	PrivateDeploymentUrl *string `mandatory:"false" json:"privateDeploymentUrl"`

	// The URL of a resource.
	PublicDeploymentUrl *string `mandatory:"false" json:"publicDeploymentUrl"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DeploymentSummary) String() string {
	return common.PointerString(m)
}
