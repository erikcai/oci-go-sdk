// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts.
//

package budget

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Budget A budget.
type Budget struct {

	// The OCID of the budget
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the budget.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The amount of the budget expressed in the currency of the customer's rate card.
	Amount *float32 `mandatory:"true" json:"amount"`

	// The reset period for the budget.
	ResetPeriod BudgetResetPeriodEnum `mandatory:"true" json:"resetPeriod"`

	// The current state of the budget.
	LifecycleState BudgetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Total number of alert rules in the budget
	AlertRuleCount *int `mandatory:"true" json:"alertRuleCount"`

	// Time that budget was created
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time that budget was updated
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// This is DEPRECATED. For backwards compatability, the property will be populated when
	// targetType is "COMPARTMENT" AND targets contains EXACT ONE target compartment ocid.
	// For all other scenarios, this property will be left empty.
	TargetCompartmentId *string `mandatory:"false" json:"targetCompartmentId"`

	// The description of the budget.
	Description *string `mandatory:"false" json:"description"`

	// The type of target on which the budget is applied.
	TargetType BudgetTargetTypeEnum `mandatory:"false" json:"targetType,omitempty"`

	// The list of targets on which the budget is applied.
	//   If targetType is "COMPARTMENT", targets contains list of compartment OCIDs.
	//   If targetType is "TAG", targets contains list of tag identifiers in the form of "{tagNamespace}.{tagKey}.{tagValue}".
	Targets []string `mandatory:"false" json:"targets"`

	// Version of the budget. Starts from 1 and increments by 1.
	Version *int `mandatory:"false" json:"version"`

	// The actual spend in currency for the current budget cycle
	ActualSpend *float32 `mandatory:"false" json:"actualSpend"`

	// The forecasted spend in currency by the end of the current budget cycle
	ForecastedSpend *float32 `mandatory:"false" json:"forecastedSpend"`

	// The time that the budget spend was last computed
	TimeSpendComputed *common.SDKTime `mandatory:"false" json:"timeSpendComputed"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Budget) String() string {
	return common.PointerString(m)
}

// BudgetResetPeriodEnum Enum with underlying type: string
type BudgetResetPeriodEnum string

// Set of constants representing the allowable values for BudgetResetPeriodEnum
const (
	BudgetResetPeriodMonthly BudgetResetPeriodEnum = "MONTHLY"
)

var mappingBudgetResetPeriod = map[string]BudgetResetPeriodEnum{
	"MONTHLY": BudgetResetPeriodMonthly,
}

// GetBudgetResetPeriodEnumValues Enumerates the set of values for BudgetResetPeriodEnum
func GetBudgetResetPeriodEnumValues() []BudgetResetPeriodEnum {
	values := make([]BudgetResetPeriodEnum, 0)
	for _, v := range mappingBudgetResetPeriod {
		values = append(values, v)
	}
	return values
}

// BudgetTargetTypeEnum Enum with underlying type: string
type BudgetTargetTypeEnum string

// Set of constants representing the allowable values for BudgetTargetTypeEnum
const (
	BudgetTargetTypeCompartment BudgetTargetTypeEnum = "COMPARTMENT"
	BudgetTargetTypeTag         BudgetTargetTypeEnum = "TAG"
)

var mappingBudgetTargetType = map[string]BudgetTargetTypeEnum{
	"COMPARTMENT": BudgetTargetTypeCompartment,
	"TAG":         BudgetTargetTypeTag,
}

// GetBudgetTargetTypeEnumValues Enumerates the set of values for BudgetTargetTypeEnum
func GetBudgetTargetTypeEnumValues() []BudgetTargetTypeEnum {
	values := make([]BudgetTargetTypeEnum, 0)
	for _, v := range mappingBudgetTargetType {
		values = append(values, v)
	}
	return values
}

// BudgetLifecycleStateEnum Enum with underlying type: string
type BudgetLifecycleStateEnum string

// Set of constants representing the allowable values for BudgetLifecycleStateEnum
const (
	BudgetLifecycleStateActive   BudgetLifecycleStateEnum = "ACTIVE"
	BudgetLifecycleStateInactive BudgetLifecycleStateEnum = "INACTIVE"
)

var mappingBudgetLifecycleState = map[string]BudgetLifecycleStateEnum{
	"ACTIVE":   BudgetLifecycleStateActive,
	"INACTIVE": BudgetLifecycleStateInactive,
}

// GetBudgetLifecycleStateEnumValues Enumerates the set of values for BudgetLifecycleStateEnum
func GetBudgetLifecycleStateEnumValues() []BudgetLifecycleStateEnum {
	values := make([]BudgetLifecycleStateEnum, 0)
	for _, v := range mappingBudgetLifecycleState {
		values = append(values, v)
	}
	return values
}
