// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v26/common"
)

// CreateVnicWorkerDetails The data to create vnicWorker.
type CreateVnicWorkerDetails struct {

	// The OCID of the compartment that contains the VNIC worker.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of associated service VNIC.
	ServiceVnicId *string `mandatory:"true" json:"serviceVnicId"`

	// A user-friendly name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// List of vnicWorker IP OCIDs.
	WorkerIps []string `mandatory:"false" json:"workerIps"`

	// The instance where vnicWorker needs to be attached.
	WorkerInstanceId *string `mandatory:"false" json:"workerInstanceId"`

	// Which physical network interface card (NIC) the VNIC worker uses. Defaults to 0.
	// Certain bare metal instance shapes have two active physical NICs (0 and 1). If
	// you add a VNIC worker to one of these instances, you can specify which NIC
	// the VNIC worker will use. Note that it is required for NIC to have at least a single
	// VNIC attached before attaching a VNIC worker.
	WorkerNicIndex *int `mandatory:"false" json:"workerNicIndex"`

	// Specifies whether the vnicworker is enabled for forwarding traffic at creation
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

func (m CreateVnicWorkerDetails) String() string {
	return common.PointerString(m)
}
