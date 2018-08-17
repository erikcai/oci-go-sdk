// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// ONS Gateway API
//
// A description of the ONS Gateway API
//

package ons

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PublishResult The response to a publishMessage call.
type PublishResult struct {

	// The UUID of the message.
	MessageId *string `mandatory:"false" json:"messageId"`

	// The time that the service received the message.
	TimeStamp *common.SDKTime `mandatory:"false" json:"timeStamp"`
}

func (m PublishResult) String() string {
	return common.PointerString(m)
}
