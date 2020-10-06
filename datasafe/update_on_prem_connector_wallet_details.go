// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v26/common"
)

// UpdateOnPremConnectorWalletDetails The details used to update an on-premises connector's wallet.
type UpdateOnPremConnectorWalletDetails struct {

	// Indicates whether to update or not. If false, the wallet will not be updated. Default is false.
	IsUpdate *bool `mandatory:"false" json:"isUpdate"`
}

func (m UpdateOnPremConnectorWalletDetails) String() string {
	return common.PointerString(m)
}
