// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// KAM API
//
// description: |
//   Kubernetes Add-on Manager API for installing, uninstalling and upgrading
//   OKE add-ons or Marketplace applications on an OKE cluster
//

package kam

import (
	"github.com/oracle/oci-go-sdk/v26/common"
)

// CreateKamReleaseDetails Details of the installation request. The KAM chart OCID can be looked up
// using the ListKamCharts API
type CreateKamReleaseDetails struct {

	// The OCID of the OKE Add-on or Marketplace app to deploy
	KamChartId *string `mandatory:"true" json:"kamChartId"`

	// List of overrides for default configuration
	Overrides []KamOverride `mandatory:"false" json:"overrides"`

	// Namespace to install to (optional)
	Namespace *string `mandatory:"false" json:"namespace"`

	// Release name to install (optional)
	ReleaseName *string `mandatory:"false" json:"releaseName"`
}

func (m CreateKamReleaseDetails) String() string {
	return common.PointerString(m)
}
