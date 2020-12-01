// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// CreateFunctionDetails Properties to create a new function.
type CreateFunctionDetails struct {

	// The display name of the function. The display name must be unique within the application containing the function. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the application this function belongs to.
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// The qualified name of the Docker image to use in the function, including the image tag.
	// The image should be in the OCI Registry that is in the same region as the function itself.
	// Example: `phx.ocir.io/ten/functions/function:0.0.1`
	Image *string `mandatory:"true" json:"image"`

	// Maximum usable memory for the function (MiB).
	MemoryInMBs *int64 `mandatory:"true" json:"memoryInMBs"`

	// The image digest for the version of the image that will be pulled when invoking this function.
	// If no value is specified, the digest currently associated with the image in the OCI Registry will be used.
	// Example: `sha256:ca0eeb6fb05351dfc8759c20733c91def84cb8007aa89a5bf606bc8b315b9fc7`
	ImageDigest *string `mandatory:"false" json:"imageDigest"`

	// Function configuration. These values are passed on to the function as environment variables, this overrides application configuration values.
	// Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.
	// Example: `{"MY_FUNCTION_CONFIG": "ConfVal"}`
	// The maximum size for all configuration keys and values is limited to 4KB. This is measured as the sum of octets necessary to represent each key and value in UTF-8.
	Config map[string]string `mandatory:"false" json:"config"`

	// Timeout for executions of the function. Value in seconds.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateFunctionDetails) String() string {
	return common.PointerString(m)
}
