// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Synthetic
//
// API for APM Synthetic service. Use this API to query Synthethetic scripts and monitors.
//

package apmsynthetics

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// MonitorScriptParameterInfo Script Parameter Details of the monitor.
// isOverwritten specfies that if the parameters used script resource are overwritten by monitor.
// If User overrides the parameter value in monitor, then, parameter is considered overwritten. Over written values
// will be used to run the monitor.
type MonitorScriptParameterInfo struct {
	MonitorScriptParameter *MonitorScriptParameter `mandatory:"true" json:"monitorScriptParameter"`

	// Describes if  the parameter value is secret and should be kept confidential.
	// isSecret was specified in either CreateScript or UpdateScript API
	IsSecret *bool `mandatory:"true" json:"isSecret"`

	// If parameter value is default or overwritten.
	IsOverwritten *bool `mandatory:"true" json:"isOverwritten"`
}

func (m MonitorScriptParameterInfo) String() string {
	return common.PointerString(m)
}
