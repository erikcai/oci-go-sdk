// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AnalyticsClusterMemoryEstimates Returns the OCID of the MySQLaaS Instance with the most recent Analytics Cluster memory estimate
// that can be used to provision the Analytics Cluster size. For each MySQL user table the estimated memory
// footprint when the table is loaded to the Analytics Cluster memory is returned.
type AnalyticsClusterMemoryEstimates struct {

	// The OCID of the MySQLaaS Instance that was used to estimate the memory footprints.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// Current status of the Work Request generating the estimate.
	Status WorkRequestStatusTypeEnum `mandatory:"true" json:"status"`

	// The date and time that the Work Request for the estimate was issued, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc333).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time that the estimate was generated, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc333).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Collection of schemas with estimated memory footprints for MySQL user tables of each schema
	// when loaded to Analytics Cluster memory.
	TableSchemas []TableSchema `mandatory:"true" json:"tableSchemas"`
}

func (m AnalyticsClusterMemoryEstimates) String() string {
	return common.PointerString(m)
}
