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
	"github.com/oracle/oci-go-sdk/v25/common"
)

// DatabaseRegistration Description of DatabaseRegistration.
type DatabaseRegistration struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the databaseRegistration being referenced.
	Id *string `mandatory:"true" json:"id"`

	// Object Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	Fqdn *string `mandatory:"true" json:"fqdn"`

	// The private IP address (in the customer's VCN) of the customer's endpoint (typically a database).
	IpAddress *string `mandatory:"true" json:"ipAddress"`

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

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database being referenced.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The pivate IP address (in the GGS Service VCN) which will permit access to the Customer VCN.
	RcePrivateIp *string `mandatory:"false" json:"rcePrivateIp"`

	// last accepted request token which is generated by splat store.
	LastAcceptedRequestToken *string `mandatory:"false" json:"lastAcceptedRequestToken"`

	// last commpleted request token which is generated by splat store.
	LastCompletedRequestToken *string `mandatory:"false" json:"lastCompletedRequestToken"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The Username Oracle GoldenGate will utilize to connect the associated RDBMS.  This Username must already exist and be available for use by the database.  It must conform to the specific security requirements implemented by the database including length, case sensitivity, etc.
	Username *string `mandatory:"false" json:"username"`

	// The Password Oracle GoldenGate will utilize to connect the associated RDBMS.  It must conform to the specific security requirements implemented by the database including length, case sensitivity, etc.
	Password *string `mandatory:"false" json:"password"`

	// The wallet utilized by Oracle GoldenGate to make connections to a database.  It must contain the files cwallet.sso and tnsnames.ora.  This attribute is expected to be base64 encoded.
	Wallet []byte `mandatory:"false" json:"wallet"`
}

func (m DatabaseRegistration) String() string {
	return common.PointerString(m)
}
