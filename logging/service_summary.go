// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// PublicLoggingControlplane API
//
// PublicLoggingControlplane API specification
//

package logging

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ServiceSummary Summary of Services that are integrated with public logging
type ServiceSummary struct {

	// Tenant OCID.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// Service id as set in Service Principal.
	ServicePrincipalName *string `mandatory:"true" json:"servicePrincipalName"`

	// Service endpoint.
	Endpoint *string `mandatory:"true" json:"endpoint"`

	// User friendly service name.
	Name *string `mandatory:"true" json:"name"`

	// Type of Resource that a Service provides.
	ResourceTypes []ResourceType `mandatory:"true" json:"resourceTypes"`

	// Apollo project namespace if any.
	Namespace *string `mandatory:"false" json:"namespace"`

	// Service id.
	Id *string `mandatory:"false" json:"id"`
}

func (m ServiceSummary) String() string {
	return common.PointerString(m)
}
