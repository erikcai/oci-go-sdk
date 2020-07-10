// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connectors API
//
// A description of the Connectors API
//

package sch

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// NotificationsTargetDetails The notifications target.
type NotificationsTargetDetails struct {

	// The topic identifier.
	TopicId *string `mandatory:"true" json:"topicId"`
}

func (m NotificationsTargetDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m NotificationsTargetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNotificationsTargetDetails NotificationsTargetDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeNotificationsTargetDetails
	}{
		"notifications",
		(MarshalTypeNotificationsTargetDetails)(m),
	}

	return json.Marshal(&s)
}
