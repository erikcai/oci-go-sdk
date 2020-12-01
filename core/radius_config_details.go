// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm),
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm), and
// Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as virtual cloud networks (VCNs),
// compute instances, block storage volumes, and container images.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// RadiusConfigDetails The detail of RADIUS's authentication configuration.
type RadiusConfigDetails struct {

	// Allowed values:
	//   * `PAP`: An authentication option used by PPP to validate user.
	//   * `CHAP`: An authentication option used by PPP to validate user but performed only at the initial link establishment.
	//   * `MS-CHAP-v2`: An authentication option used by PPP to periodically validate the user.
	AuthenticationType RadiusConfigDetailsAuthenticationTypeEnum `mandatory:"false" json:"authenticationType,omitempty"`

	// A list of usable RADIUS server.
	Servers []RadiusServerDetails `mandatory:"false" json:"servers"`

	// Whether to enable RADIUS accounting. When this enabled, accouting port becomes required filed.
	IsRadiusAccountingEnabled *bool `mandatory:"false" json:"isRadiusAccountingEnabled"`

	// Enable case-sensitivity or not in RADIUS authentication.
	IsCaseSensitive *bool `mandatory:"false" json:"isCaseSensitive"`
}

func (m RadiusConfigDetails) String() string {
	return common.PointerString(m)
}

// RadiusConfigDetailsAuthenticationTypeEnum Enum with underlying type: string
type RadiusConfigDetailsAuthenticationTypeEnum string

// Set of constants representing the allowable values for RadiusConfigDetailsAuthenticationTypeEnum
const (
	RadiusConfigDetailsAuthenticationTypePap      RadiusConfigDetailsAuthenticationTypeEnum = "PAP"
	RadiusConfigDetailsAuthenticationTypeChap     RadiusConfigDetailsAuthenticationTypeEnum = "CHAP"
	RadiusConfigDetailsAuthenticationTypeMsChapV2 RadiusConfigDetailsAuthenticationTypeEnum = "MS_CHAP_V2"
)

var mappingRadiusConfigDetailsAuthenticationType = map[string]RadiusConfigDetailsAuthenticationTypeEnum{
	"PAP":        RadiusConfigDetailsAuthenticationTypePap,
	"CHAP":       RadiusConfigDetailsAuthenticationTypeChap,
	"MS_CHAP_V2": RadiusConfigDetailsAuthenticationTypeMsChapV2,
}

// GetRadiusConfigDetailsAuthenticationTypeEnumValues Enumerates the set of values for RadiusConfigDetailsAuthenticationTypeEnum
func GetRadiusConfigDetailsAuthenticationTypeEnumValues() []RadiusConfigDetailsAuthenticationTypeEnum {
	values := make([]RadiusConfigDetailsAuthenticationTypeEnum, 0)
	for _, v := range mappingRadiusConfigDetailsAuthenticationType {
		values = append(values, v)
	}
	return values
}
