// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// RebootToRecoveryImagePowerActionDetails The representation of RebootToRecoveryImagePowerActionDetails
type RebootToRecoveryImagePowerActionDetails struct {

	// Custom metadata key/value pairs that you provide. This field is reserved for a forthcoming feature.
	RecoveryMetadata map[string]string `mandatory:"false" json:"recoveryMetadata"`
}

func (m RebootToRecoveryImagePowerActionDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RebootToRecoveryImagePowerActionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRebootToRecoveryImagePowerActionDetails RebootToRecoveryImagePowerActionDetails
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeRebootToRecoveryImagePowerActionDetails
	}{
		"reboot_to_recovery_image",
		(MarshalTypeRebootToRecoveryImagePowerActionDetails)(m),
	}

	return json.Marshal(&s)
}
