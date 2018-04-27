// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Resource Query Service
//
// Query for resources across your cloud infrastructure
//

package resourcequery

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// FreeTextSearchDetails A request containing free text search parameters.
type FreeTextSearchDetails struct {

	// The text to search for.
	Text *string `mandatory:"true" json:"text"`

	// Defines the type of matching context returned in response. If HIGHLIGHTS then there will be highlighting fragments returned from the service (see ResourceSummary.searchContext and SearchContext). Default is NONE.
	MatchingContextType SearchDetailsMatchingContextTypeEnum `mandatory:"false" json:"matchingContextType,omitempty"`
}

//GetMatchingContextType returns MatchingContextType
func (m FreeTextSearchDetails) GetMatchingContextType() SearchDetailsMatchingContextTypeEnum {
	return m.MatchingContextType
}

func (m FreeTextSearchDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m FreeTextSearchDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFreeTextSearchDetails FreeTextSearchDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeFreeTextSearchDetails
	}{
		"FreeText",
		(MarshalTypeFreeTextSearchDetails)(m),
	}

	return json.Marshal(&s)
}
