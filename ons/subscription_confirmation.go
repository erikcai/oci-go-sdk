// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Notification API
//
// Use the Notification API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notification Overview (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SubscriptionConfirmation The subscription confirmation.
type SubscriptionConfirmation struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription.
	Id *string `mandatory:"true" json:"id"`

	// The URL for confirming the subscription. Contains an authentication token and identifies the subscription by ID and protocol.
	Url *string `mandatory:"true" json:"url"`
}

func (m SubscriptionConfirmation) String() string {
	return common.PointerString(m)
}
