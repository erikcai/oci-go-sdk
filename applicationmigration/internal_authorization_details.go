// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration API
//
// Application Migration simplifies the migration of applications from Oracle Cloud Infrastructure Classic to Oracle Cloud Infrastructure.
// You can use Application Migration API to migrate applications, such as Oracle Java Cloud Service, SOA Cloud Service, and Integration Classic
// instances, to Oracle Cloud Infrastructure. For more information, see
// Overview of Application Migration (https://docs.cloud.oracle.com/iaas/application-migration/appmigrationoverview.htm).
//

package applicationmigration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v29/common"
)

// InternalAuthorizationDetails Credentials to access Oracle Cloud Infrastructure - Classic, which is the source environment from which you want to migrate the application.
type InternalAuthorizationDetails struct {

	// User with Compute Operations role in Oracle Cloud Infrastructure - Classic.
	Username *string `mandatory:"true" json:"username"`

	// Password for this user.
	Password *string `mandatory:"true" json:"password"`
}

func (m InternalAuthorizationDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InternalAuthorizationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInternalAuthorizationDetails InternalAuthorizationDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeInternalAuthorizationDetails
	}{
		"INTERNAL_COMPUTE",
		(MarshalTypeInternalAuthorizationDetails)(m),
	}

	return json.Marshal(&s)
}
