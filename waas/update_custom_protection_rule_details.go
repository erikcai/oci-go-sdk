// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateCustomProtectionRuleDetails Updates the configuration details of a Custom Protection rule. The update is possible only if the rule hasn't been selected as active in any WAAS policy (remains in `Off` state).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateCustomProtectionRuleDetails struct {

	// A user-friendly name for the Custom Protection rule.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description for the Custom Protection rule.
	Description *string `mandatory:"false" json:"description"`

	// The template text of the Custom Protection rule. The syntax is based on ModSecurity Rule Language. Additionaly it needs to include two variables / placeholders which will be replaced during publishing.
	// - **{{mode}}** - rule action, defined by user in UI, like `OFF`, `DETECT` or `BLOCK`.
	// - **{{id_1}}** - unique rule ID which identifies a `SecRule`, generated by the system. Multiple IDs can be used by increasing the number of the variable for every `SecRule` defined in the template.
	// *Example usage:*
	//   ```
	//   SecRule REQUEST_COOKIES "regex matching SQL injection - part 1/2" \
	//           "phase:2,                                                 \
	//           msg:'Detects chained SQL injection attempts 1/2.',        \
	//           id: {{id_1}},                                             \
	//           ctl:ruleEngine={{mode}},                                  \
	//           deny"
	//   SecRule REQUEST_COOKIES "regex matching SQL injection - part 2/2" \
	//           "phase:2,                                                 \
	//           msg:'Detects chained SQL injection attempts 2/2.',        \
	//           id: {{id_2}},                                             \
	//           ctl:ruleEngine={{mode}},                                  \
	//           deny"
	//   ```
	//   The example contains two `SecRules` each having distinct regex expression to match
	//   `Cookie` header value during second input analysis phase.
	//   The disruptive `deny` action takes effect only when `{{mode}}` is set to `BLOCK`.
	//   The message is logged either when `{{mode}}` is set to `DETECT` or `BLOCK`.
	//
	// For more information about ModSecurity's open source WAF rules, see Mod Security's documentation (https://www.modsecurity.org/CRS/Documentation/making.html).
	Template *string `mandatory:"false" json:"template"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateCustomProtectionRuleDetails) String() string {
	return common.PointerString(m)
}
