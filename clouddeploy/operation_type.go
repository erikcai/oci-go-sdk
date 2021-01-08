// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateApplication OperationTypeEnum = "CREATE_APPLICATION"
	OperationTypeUpdateApplication OperationTypeEnum = "UPDATE_APPLICATION"
	OperationTypeDeleteApplication OperationTypeEnum = "DELETE_APPLICATION"
	OperationTypeCreatePipeline    OperationTypeEnum = "CREATE_PIPELINE"
	OperationTypeUpdatePipeline    OperationTypeEnum = "UPDATE_PIPELINE"
	OperationTypeDeletePipeline    OperationTypeEnum = "DELETE_PIPELINE"
	OperationTypeCreateStage       OperationTypeEnum = "CREATE_STAGE"
	OperationTypeUpdateStage       OperationTypeEnum = "UPDATE_STAGE"
	OperationTypeDeleteStage       OperationTypeEnum = "DELETE_STAGE"
	OperationTypeCreateArtifact    OperationTypeEnum = "CREATE_ARTIFACT"
	OperationTypeUpdateArtifact    OperationTypeEnum = "UPDATE_ARTIFACT"
	OperationTypeDeleteArtifact    OperationTypeEnum = "DELETE_ARTIFACT"
	OperationTypeCreateEnvironment OperationTypeEnum = "CREATE_ENVIRONMENT"
	OperationTypeUpdateEnvironment OperationTypeEnum = "UPDATE_ENVIRONMENT"
	OperationTypeDeleteEnvironment OperationTypeEnum = "DELETE_ENVIRONMENT"
	OperationTypeCreateDeployment  OperationTypeEnum = "CREATE_DEPLOYMENT"
	OperationTypeUpdateDeployment  OperationTypeEnum = "UPDATE_DEPLOYMENT"
	OperationTypeDeleteDeployment  OperationTypeEnum = "DELETE_DEPLOYMENT"
)

var mappingOperationType = map[string]OperationTypeEnum{
	"CREATE_APPLICATION": OperationTypeCreateApplication,
	"UPDATE_APPLICATION": OperationTypeUpdateApplication,
	"DELETE_APPLICATION": OperationTypeDeleteApplication,
	"CREATE_PIPELINE":    OperationTypeCreatePipeline,
	"UPDATE_PIPELINE":    OperationTypeUpdatePipeline,
	"DELETE_PIPELINE":    OperationTypeDeletePipeline,
	"CREATE_STAGE":       OperationTypeCreateStage,
	"UPDATE_STAGE":       OperationTypeUpdateStage,
	"DELETE_STAGE":       OperationTypeDeleteStage,
	"CREATE_ARTIFACT":    OperationTypeCreateArtifact,
	"UPDATE_ARTIFACT":    OperationTypeUpdateArtifact,
	"DELETE_ARTIFACT":    OperationTypeDeleteArtifact,
	"CREATE_ENVIRONMENT": OperationTypeCreateEnvironment,
	"UPDATE_ENVIRONMENT": OperationTypeUpdateEnvironment,
	"DELETE_ENVIRONMENT": OperationTypeDeleteEnvironment,
	"CREATE_DEPLOYMENT":  OperationTypeCreateDeployment,
	"UPDATE_DEPLOYMENT":  OperationTypeUpdateDeployment,
	"DELETE_DEPLOYMENT":  OperationTypeDeleteDeployment,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationType {
		values = append(values, v)
	}
	return values
}
