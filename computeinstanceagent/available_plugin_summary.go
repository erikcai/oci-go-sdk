// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Agent API
//
// API for the Oracle Cloud Agent software running on compute instances. Oracle Cloud Agent
// is a lightweight process that monitors and manages compute instances.
//

package computeinstanceagent

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// AvailablePluginSummary Describes where the plugin is supported
type AvailablePluginSummary struct {

	// The plugin name
	Name *string `mandatory:"true" json:"name"`

	// Is the plugin supported or not
	IsSupported *bool `mandatory:"true" json:"isSupported"`

	// Is the plugin enabled or disabled by default
	IsEnabledByDefault *bool `mandatory:"true" json:"isEnabledByDefault"`

	// A brief description of the plugin functionality
	Summary *string `mandatory:"false" json:"summary"`
}

func (m AvailablePluginSummary) String() string {
	return common.PointerString(m)
}
