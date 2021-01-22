// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// UsageApi API
//
// A description of the UsageApi API.
//

package usage

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// SubscriptionInfo The response object for the GetSubscriptionInfo API call. It provides information about the specified billing cycle.
type SubscriptionInfo struct {

	// The OCID of the target tenancy.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// The billing cycle start date (UTC), in the format specified by RFC3339 (https://tools.ietf.org/html/rfc3339).
	BillStart *common.SDKTime `mandatory:"false" json:"billStart"`

	// The billing cycle end date (UTC), in the format specified by RFC3339 (https://tools.ietf.org/html/rfc3339).
	BillEnd *common.SDKTime `mandatory:"false" json:"billEnd"`

	// Total balance for the billing cycle.
	BillTotalBalance *float32 `mandatory:"false" json:"billTotalBalance"`

	// Total available credit for the billing cycle.
	BillTotalPurchased *float32 `mandatory:"false" json:"billTotalPurchased"`

	// Date and time (UTC) the balance was last modified.
	BillModifiedTime *common.SDKTime `mandatory:"false" json:"billModifiedTime"`

	// The currency of the subscription, in the format specified by ISO-4217 (https://www.iso.org/iso-4217-currency-codes.html).
	BillCurrency *string `mandatory:"false" json:"billCurrency"`
}

func (m SubscriptionInfo) String() string {
	return common.PointerString(m)
}
