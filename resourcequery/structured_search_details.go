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

// StructuredSearchDetails A request containing structured query parameters.
type StructuredSearchDetails struct {

	// The structured query describing which resources to search for.
	Query *string `mandatory:"true" json:"query"`

	// Defines the type of matching context returned in response. If HIGHLIGHTS then there will be highlighting fragments returned from the service (see ResourceSummary.searchContext and SearchContext). Default is NONE.
	MatchingContextType SearchDetailsMatchingContextTypeEnum `mandatory:"false" json:"matchingContextType,omitempty"`
}

//GetMatchingContextType returns MatchingContextType
func (m StructuredSearchDetails) GetMatchingContextType() SearchDetailsMatchingContextTypeEnum {
	return m.MatchingContextType
}

func (m StructuredSearchDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m StructuredSearchDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStructuredSearchDetails StructuredSearchDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeStructuredSearchDetails
	}{
		"Structured",
		(MarshalTypeStructuredSearchDetails)(m),
	}

	return json.Marshal(&s)
}
