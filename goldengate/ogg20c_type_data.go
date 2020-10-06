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
	"github.com/oracle/oci-go-sdk/v26/common"
)

// Ogg20cTypeData The properties and values for the details of the type "UUID_1D7CAE6C_4CB0_3ABD_97E2_CDC1F558B638"
// an example of this is:
// {
//   deploymentName: "NewYork",
//   oggAdmin: {
//     userName: "oggadmin",
//     password: "NotASecret"
//   }
// }
type Ogg20cTypeData struct {

	// The definition of an OGG deploymentName.
	DeploymentName *string `mandatory:"true" json:"deploymentName"`

	OggAdmin *UserNamePasswordPair `mandatory:"true" json:"oggAdmin"`

	SslCertificate *CertificateKeyPair `mandatory:"false" json:"sslCertificate"`
}

func (m Ogg20cTypeData) String() string {
	return common.PointerString(m)
}
