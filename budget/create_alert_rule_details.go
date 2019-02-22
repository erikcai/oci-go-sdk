// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// BudgetsControlPlane API
//
// A description of the BudgetsControlPlane API
//

package budget

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateAlertRuleDetails The create alert rule details. This is a batch-create.
type CreateAlertRuleDetails struct {

	// ACTUAL means the alert will trigger based on actual usage.
	// FORECAST means the alert will trigger based on predicted usage.
	Type CreateAlertRuleDetailsTypeEnum `mandatory:"true" json:"type"`

	// The threshold for triggering the alert. If thresholdType is PERCENTAGE, the maximum value is 10000.
	Threshold *float32 `mandatory:"true" json:"threshold"`

	// The type of threshold.
	ThresholdType CreateAlertRuleDetailsThresholdTypeEnum `mandatory:"true" json:"thresholdType"`

	// The audience that will received the alert when it triggers.
	Recipients *string `mandatory:"true" json:"recipients"`

	// The name of the alert rule.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the alert rule.
	Description *string `mandatory:"false" json:"description"`

	// The message to be sent to the recipients when alert rule is triggered.
	Message *string `mandatory:"false" json:"message"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateAlertRuleDetails) String() string {
	return common.PointerString(m)
}

// CreateAlertRuleDetailsTypeEnum Enum with underlying type: string
type CreateAlertRuleDetailsTypeEnum string

// Set of constants representing the allowable values for CreateAlertRuleDetailsTypeEnum
const (
	CreateAlertRuleDetailsTypeActual   CreateAlertRuleDetailsTypeEnum = "ACTUAL"
	CreateAlertRuleDetailsTypeForecast CreateAlertRuleDetailsTypeEnum = "FORECAST"
)

var mappingCreateAlertRuleDetailsType = map[string]CreateAlertRuleDetailsTypeEnum{
	"ACTUAL":   CreateAlertRuleDetailsTypeActual,
	"FORECAST": CreateAlertRuleDetailsTypeForecast,
}

// GetCreateAlertRuleDetailsTypeEnumValues Enumerates the set of values for CreateAlertRuleDetailsTypeEnum
func GetCreateAlertRuleDetailsTypeEnumValues() []CreateAlertRuleDetailsTypeEnum {
	values := make([]CreateAlertRuleDetailsTypeEnum, 0)
	for _, v := range mappingCreateAlertRuleDetailsType {
		values = append(values, v)
	}
	return values
}

// CreateAlertRuleDetailsThresholdTypeEnum Enum with underlying type: string
type CreateAlertRuleDetailsThresholdTypeEnum string

// Set of constants representing the allowable values for CreateAlertRuleDetailsThresholdTypeEnum
const (
	CreateAlertRuleDetailsThresholdTypePercentage CreateAlertRuleDetailsThresholdTypeEnum = "PERCENTAGE"
	CreateAlertRuleDetailsThresholdTypeAbsolute   CreateAlertRuleDetailsThresholdTypeEnum = "ABSOLUTE"
)

var mappingCreateAlertRuleDetailsThresholdType = map[string]CreateAlertRuleDetailsThresholdTypeEnum{
	"PERCENTAGE": CreateAlertRuleDetailsThresholdTypePercentage,
	"ABSOLUTE":   CreateAlertRuleDetailsThresholdTypeAbsolute,
}

// GetCreateAlertRuleDetailsThresholdTypeEnumValues Enumerates the set of values for CreateAlertRuleDetailsThresholdTypeEnum
func GetCreateAlertRuleDetailsThresholdTypeEnumValues() []CreateAlertRuleDetailsThresholdTypeEnum {
	values := make([]CreateAlertRuleDetailsThresholdTypeEnum, 0)
	for _, v := range mappingCreateAlertRuleDetailsThresholdType {
		values = append(values, v)
	}
	return values
}
