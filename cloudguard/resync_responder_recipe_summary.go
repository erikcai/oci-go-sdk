// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ResyncResponderRecipeSummary Summary of the ResponderRecipe.
type ResyncResponderRecipeSummary struct {

	// Identifier for ResponderRecipe.
	Id *string `mandatory:"true" json:"id"`

	// ResponderRecipe Display Name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Owner of ResponderRecipe
	Owner OwnerTypeEnum `mandatory:"true" json:"owner"`

	// List of old responder rules associated with the recipe
	OldResponderRecipeResponderRules []ResponderRecipeResponderRule `mandatory:"true" json:"oldResponderRecipeResponderRules"`

	// List of new responder rules associated with the recipe
	NewResponderRecipeResponderRules []ResponderRecipeResponderRule `mandatory:"true" json:"newResponderRecipeResponderRules"`

	// The id of the source responder recipe.
	SourceResponderRecipeId *string `mandatory:"false" json:"sourceResponderRecipeId"`
}

func (m ResyncResponderRecipeSummary) String() string {
	return common.PointerString(m)
}
