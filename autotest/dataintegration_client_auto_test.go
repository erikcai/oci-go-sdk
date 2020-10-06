package autotest

import (
	"github.com/oracle/oci-go-sdk/v26/common"
	"github.com/oracle/oci-go-sdk/v26/dataintegration"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createDataIntegrationClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := dataintegration.NewDataIntegrationClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientChangeCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ChangeCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ChangeCompartment", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ChangeCompartment")
	assert.NoError(t, err)

	type ChangeCompartmentRequestInfo struct {
		ContainerId string
		Request     dataintegration.ChangeCompartmentRequest
	}

	var requests []ChangeCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateApplication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateApplication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateApplication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateApplication", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateApplication")
	assert.NoError(t, err)

	type CreateApplicationRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateApplicationRequest
	}

	var requests []CreateApplicationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateApplication(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateConnection(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateConnection", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateConnection")
	assert.NoError(t, err)

	type CreateConnectionRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateConnectionRequest
	}

	var requests []CreateConnectionRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateConnectionRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateConnectionDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "modelType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"MYSQL_CONNECTION":                 &dataintegration.CreateConnectionFromMySql{},
				"GENERIC_JDBC_CONNECTION":          &dataintegration.CreateConnectionFromJdbc{},
				"ORACLE_ATP_CONNECTION":            &dataintegration.CreateConnectionFromAtp{},
				"ORACLE_ADWC_CONNECTION":           &dataintegration.CreateConnectionFromAdwc{},
				"ORACLEDB_CONNECTION":              &dataintegration.CreateConnectionFromOracle{},
				"ORACLE_OBJECT_STORAGE_CONNECTION": &dataintegration.CreateConnectionFromObjectStorage{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateConnectionValidation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateConnectionValidation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateConnectionValidation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateConnectionValidation", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateConnectionValidation")
	assert.NoError(t, err)

	type CreateConnectionValidationRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateConnectionValidationRequest
	}

	var requests []CreateConnectionValidationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateConnectionValidation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateDataAsset(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateDataAsset")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDataAsset is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateDataAsset", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateDataAsset")
	assert.NoError(t, err)

	type CreateDataAssetRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateDataAssetRequest
	}

	var requests []CreateDataAssetRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateDataAssetRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateDataAssetDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "modelType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"GENERIC_JDBC_DATA_ASSET":          &dataintegration.CreateDataAssetFromJdbc{},
				"MYSQL_DATA_ASSET":                 &dataintegration.CreateDataAssetFromMySql{},
				"ORACLE_DATA_ASSET":                &dataintegration.CreateDataAssetFromOracle{},
				"ORACLE_ADWC_DATA_ASSET":           &dataintegration.CreateDataAssetFromAdwc{},
				"ORACLE_ATP_DATA_ASSET":            &dataintegration.CreateDataAssetFromAtp{},
				"ORACLE_OBJECT_STORAGE_DATA_ASSET": &dataintegration.CreateDataAssetFromObjectStorage{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateDataAsset(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateDataFlow(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateDataFlow")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDataFlow is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateDataFlow", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateDataFlow")
	assert.NoError(t, err)

	type CreateDataFlowRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateDataFlowRequest
	}

	var requests []CreateDataFlowRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateDataFlow(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateDataFlowValidation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateDataFlowValidation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDataFlowValidation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateDataFlowValidation", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateDataFlowValidation")
	assert.NoError(t, err)

	type CreateDataFlowValidationRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateDataFlowValidationRequest
	}

	var requests []CreateDataFlowValidationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateDataFlowValidation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateEntityShape(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateEntityShape")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateEntityShape is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateEntityShape", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateEntityShape")
	assert.NoError(t, err)

	type CreateEntityShapeRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateEntityShapeRequest
	}

	var requests []CreateEntityShapeRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateEntityShapeRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateEntityShapeDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "modelType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"FILE_ENTITY": &dataintegration.CreateEntityShapeFromFile{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateEntityShape(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateExternalPublication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateExternalPublication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateExternalPublication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateExternalPublication", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateExternalPublication")
	assert.NoError(t, err)

	type CreateExternalPublicationRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateExternalPublicationRequest
	}

	var requests []CreateExternalPublicationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateExternalPublication(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateExternalPublicationValidation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateExternalPublicationValidation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateExternalPublicationValidation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateExternalPublicationValidation", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateExternalPublicationValidation")
	assert.NoError(t, err)

	type CreateExternalPublicationValidationRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateExternalPublicationValidationRequest
	}

	var requests []CreateExternalPublicationValidationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateExternalPublicationValidation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateFolder(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateFolder")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateFolder is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateFolder", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateFolder")
	assert.NoError(t, err)

	type CreateFolderRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateFolderRequest
	}

	var requests []CreateFolderRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateFolder(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreatePatch(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreatePatch")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreatePatch is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreatePatch", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreatePatch")
	assert.NoError(t, err)

	type CreatePatchRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreatePatchRequest
	}

	var requests []CreatePatchRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreatePatch(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateProject(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateProject")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateProject is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateProject", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateProject")
	assert.NoError(t, err)

	type CreateProjectRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateProjectRequest
	}

	var requests []CreateProjectRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateProject(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateTask(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateTask")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTask is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateTask", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateTask")
	assert.NoError(t, err)

	type CreateTaskRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateTaskRequest
	}

	var requests []CreateTaskRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateTaskRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateTaskDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "modelType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"INTEGRATION_TASK": &dataintegration.CreateTaskFromIntegrationTask{},
				"DATA_LOADER_TASK": &dataintegration.CreateTaskFromDataLoaderTask{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateTask(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateTaskRun(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateTaskRun")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTaskRun is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateTaskRun", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateTaskRun")
	assert.NoError(t, err)

	type CreateTaskRunRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateTaskRunRequest
	}

	var requests []CreateTaskRunRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateTaskRun(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateTaskValidation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateTaskValidation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTaskValidation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateTaskValidation", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateTaskValidation")
	assert.NoError(t, err)

	type CreateTaskValidationRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateTaskValidationRequest
	}

	var requests []CreateTaskValidationRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateTaskValidationRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateTaskValidationDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "modelType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"DATA_LOADER_TASK": &dataintegration.CreateTaskValidationFromDataLoaderTask{},
				"INTEGRATION_TASK": &dataintegration.CreateTaskValidationFromIntegrationTask{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateTaskValidation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientCreateWorkspace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "CreateWorkspace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateWorkspace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "CreateWorkspace", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "CreateWorkspace")
	assert.NoError(t, err)

	type CreateWorkspaceRequestInfo struct {
		ContainerId string
		Request     dataintegration.CreateWorkspaceRequest
	}

	var requests []CreateWorkspaceRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateWorkspace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeleteApplication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeleteApplication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteApplication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeleteApplication", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeleteApplication")
	assert.NoError(t, err)

	type DeleteApplicationRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeleteApplicationRequest
	}

	var requests []DeleteApplicationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteApplication(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeleteConnection(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeleteConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeleteConnection", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeleteConnection")
	assert.NoError(t, err)

	type DeleteConnectionRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeleteConnectionRequest
	}

	var requests []DeleteConnectionRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeleteConnectionValidation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeleteConnectionValidation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteConnectionValidation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeleteConnectionValidation", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeleteConnectionValidation")
	assert.NoError(t, err)

	type DeleteConnectionValidationRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeleteConnectionValidationRequest
	}

	var requests []DeleteConnectionValidationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteConnectionValidation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeleteDataAsset(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeleteDataAsset")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDataAsset is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeleteDataAsset", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeleteDataAsset")
	assert.NoError(t, err)

	type DeleteDataAssetRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeleteDataAssetRequest
	}

	var requests []DeleteDataAssetRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteDataAsset(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeleteDataFlow(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeleteDataFlow")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDataFlow is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeleteDataFlow", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeleteDataFlow")
	assert.NoError(t, err)

	type DeleteDataFlowRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeleteDataFlowRequest
	}

	var requests []DeleteDataFlowRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteDataFlow(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeleteDataFlowValidation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeleteDataFlowValidation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDataFlowValidation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeleteDataFlowValidation", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeleteDataFlowValidation")
	assert.NoError(t, err)

	type DeleteDataFlowValidationRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeleteDataFlowValidationRequest
	}

	var requests []DeleteDataFlowValidationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteDataFlowValidation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeleteExternalPublication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeleteExternalPublication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteExternalPublication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeleteExternalPublication", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeleteExternalPublication")
	assert.NoError(t, err)

	type DeleteExternalPublicationRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeleteExternalPublicationRequest
	}

	var requests []DeleteExternalPublicationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteExternalPublication(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeleteExternalPublicationValidation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeleteExternalPublicationValidation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteExternalPublicationValidation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeleteExternalPublicationValidation", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeleteExternalPublicationValidation")
	assert.NoError(t, err)

	type DeleteExternalPublicationValidationRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeleteExternalPublicationValidationRequest
	}

	var requests []DeleteExternalPublicationValidationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteExternalPublicationValidation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeleteFolder(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeleteFolder")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteFolder is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeleteFolder", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeleteFolder")
	assert.NoError(t, err)

	type DeleteFolderRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeleteFolderRequest
	}

	var requests []DeleteFolderRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteFolder(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeletePatch(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeletePatch")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeletePatch is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeletePatch", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeletePatch")
	assert.NoError(t, err)

	type DeletePatchRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeletePatchRequest
	}

	var requests []DeletePatchRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeletePatch(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeleteProject(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeleteProject")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteProject is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeleteProject", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeleteProject")
	assert.NoError(t, err)

	type DeleteProjectRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeleteProjectRequest
	}

	var requests []DeleteProjectRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteProject(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeleteTask(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeleteTask")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteTask is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeleteTask", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeleteTask")
	assert.NoError(t, err)

	type DeleteTaskRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeleteTaskRequest
	}

	var requests []DeleteTaskRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteTask(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeleteTaskRun(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeleteTaskRun")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteTaskRun is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeleteTaskRun", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeleteTaskRun")
	assert.NoError(t, err)

	type DeleteTaskRunRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeleteTaskRunRequest
	}

	var requests []DeleteTaskRunRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteTaskRun(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeleteTaskValidation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeleteTaskValidation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteTaskValidation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeleteTaskValidation", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeleteTaskValidation")
	assert.NoError(t, err)

	type DeleteTaskValidationRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeleteTaskValidationRequest
	}

	var requests []DeleteTaskValidationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteTaskValidation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientDeleteWorkspace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "DeleteWorkspace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteWorkspace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "DeleteWorkspace", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "DeleteWorkspace")
	assert.NoError(t, err)

	type DeleteWorkspaceRequestInfo struct {
		ContainerId string
		Request     dataintegration.DeleteWorkspaceRequest
	}

	var requests []DeleteWorkspaceRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteWorkspace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetApplication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetApplication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetApplication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetApplication", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetApplication")
	assert.NoError(t, err)

	type GetApplicationRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetApplicationRequest
	}

	var requests []GetApplicationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetApplication(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetConnection(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetConnection", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetConnection")
	assert.NoError(t, err)

	type GetConnectionRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetConnectionRequest
	}

	var requests []GetConnectionRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetConnectionValidation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetConnectionValidation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConnectionValidation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetConnectionValidation", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetConnectionValidation")
	assert.NoError(t, err)

	type GetConnectionValidationRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetConnectionValidationRequest
	}

	var requests []GetConnectionValidationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetConnectionValidation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetCountStatistic(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetCountStatistic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCountStatistic is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetCountStatistic", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetCountStatistic")
	assert.NoError(t, err)

	type GetCountStatisticRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetCountStatisticRequest
	}

	var requests []GetCountStatisticRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetCountStatistic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetDataAsset(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetDataAsset")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDataAsset is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetDataAsset", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetDataAsset")
	assert.NoError(t, err)

	type GetDataAssetRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetDataAssetRequest
	}

	var requests []GetDataAssetRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDataAsset(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetDataEntity(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetDataEntity")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDataEntity is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetDataEntity", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetDataEntity")
	assert.NoError(t, err)

	type GetDataEntityRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetDataEntityRequest
	}

	var requests []GetDataEntityRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDataEntity(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetDataFlow(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetDataFlow")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDataFlow is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetDataFlow", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetDataFlow")
	assert.NoError(t, err)

	type GetDataFlowRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetDataFlowRequest
	}

	var requests []GetDataFlowRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDataFlow(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetDataFlowValidation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetDataFlowValidation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDataFlowValidation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetDataFlowValidation", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetDataFlowValidation")
	assert.NoError(t, err)

	type GetDataFlowValidationRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetDataFlowValidationRequest
	}

	var requests []GetDataFlowValidationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDataFlowValidation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetDependentObject(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetDependentObject")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDependentObject is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetDependentObject", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetDependentObject")
	assert.NoError(t, err)

	type GetDependentObjectRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetDependentObjectRequest
	}

	var requests []GetDependentObjectRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDependentObject(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetExternalPublication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetExternalPublication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetExternalPublication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetExternalPublication", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetExternalPublication")
	assert.NoError(t, err)

	type GetExternalPublicationRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetExternalPublicationRequest
	}

	var requests []GetExternalPublicationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetExternalPublication(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetExternalPublicationValidation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetExternalPublicationValidation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetExternalPublicationValidation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetExternalPublicationValidation", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetExternalPublicationValidation")
	assert.NoError(t, err)

	type GetExternalPublicationValidationRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetExternalPublicationValidationRequest
	}

	var requests []GetExternalPublicationValidationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetExternalPublicationValidation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetFolder(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetFolder")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetFolder is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetFolder", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetFolder")
	assert.NoError(t, err)

	type GetFolderRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetFolderRequest
	}

	var requests []GetFolderRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetFolder(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetPatch(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetPatch")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPatch is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetPatch", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetPatch")
	assert.NoError(t, err)

	type GetPatchRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetPatchRequest
	}

	var requests []GetPatchRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetPatch(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetProject(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetProject")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetProject is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetProject", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetProject")
	assert.NoError(t, err)

	type GetProjectRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetProjectRequest
	}

	var requests []GetProjectRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetProject(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetPublishedObject(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetPublishedObject")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPublishedObject is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetPublishedObject", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetPublishedObject")
	assert.NoError(t, err)

	type GetPublishedObjectRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetPublishedObjectRequest
	}

	var requests []GetPublishedObjectRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetPublishedObject(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetReference(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetReference")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetReference is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetReference", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetReference")
	assert.NoError(t, err)

	type GetReferenceRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetReferenceRequest
	}

	var requests []GetReferenceRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetReference(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetSchema(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetSchema")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSchema is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetSchema", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetSchema")
	assert.NoError(t, err)

	type GetSchemaRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetSchemaRequest
	}

	var requests []GetSchemaRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetSchema(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetTask(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetTask")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTask is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetTask", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetTask")
	assert.NoError(t, err)

	type GetTaskRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetTaskRequest
	}

	var requests []GetTaskRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetTask(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetTaskRun(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetTaskRun")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTaskRun is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetTaskRun", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetTaskRun")
	assert.NoError(t, err)

	type GetTaskRunRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetTaskRunRequest
	}

	var requests []GetTaskRunRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetTaskRun(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetTaskValidation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetTaskValidation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTaskValidation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetTaskValidation", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetTaskValidation")
	assert.NoError(t, err)

	type GetTaskValidationRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetTaskValidationRequest
	}

	var requests []GetTaskValidationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetTaskValidation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetWorkRequest", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetWorkRequestRequest
	}

	var requests []GetWorkRequestRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientGetWorkspace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "GetWorkspace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkspace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "GetWorkspace", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "GetWorkspace")
	assert.NoError(t, err)

	type GetWorkspaceRequestInfo struct {
		ContainerId string
		Request     dataintegration.GetWorkspaceRequest
	}

	var requests []GetWorkspaceRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetWorkspace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListApplications(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListApplications")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListApplications is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListApplications", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListApplications")
	assert.NoError(t, err)

	type ListApplicationsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListApplicationsRequest
	}

	var requests []ListApplicationsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListApplicationsRequest)
				return c.ListApplications(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListApplicationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListApplicationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListConnectionValidations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListConnectionValidations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListConnectionValidations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListConnectionValidations", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListConnectionValidations")
	assert.NoError(t, err)

	type ListConnectionValidationsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListConnectionValidationsRequest
	}

	var requests []ListConnectionValidationsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListConnectionValidationsRequest)
				return c.ListConnectionValidations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListConnectionValidationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListConnectionValidationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListConnections(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListConnections")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListConnections is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListConnections", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListConnections")
	assert.NoError(t, err)

	type ListConnectionsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListConnectionsRequest
	}

	var requests []ListConnectionsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListConnectionsRequest)
				return c.ListConnections(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListConnectionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListConnectionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListDataAssets(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListDataAssets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDataAssets is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListDataAssets", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListDataAssets")
	assert.NoError(t, err)

	type ListDataAssetsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListDataAssetsRequest
	}

	var requests []ListDataAssetsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListDataAssetsRequest)
				return c.ListDataAssets(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListDataAssetsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListDataAssetsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListDataEntities(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListDataEntities")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDataEntities is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListDataEntities", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListDataEntities")
	assert.NoError(t, err)

	type ListDataEntitiesRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListDataEntitiesRequest
	}

	var requests []ListDataEntitiesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListDataEntitiesRequest)
				return c.ListDataEntities(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListDataEntitiesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListDataEntitiesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListDataFlowValidations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListDataFlowValidations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDataFlowValidations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListDataFlowValidations", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListDataFlowValidations")
	assert.NoError(t, err)

	type ListDataFlowValidationsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListDataFlowValidationsRequest
	}

	var requests []ListDataFlowValidationsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListDataFlowValidationsRequest)
				return c.ListDataFlowValidations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListDataFlowValidationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListDataFlowValidationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListDataFlows(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListDataFlows")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDataFlows is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListDataFlows", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListDataFlows")
	assert.NoError(t, err)

	type ListDataFlowsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListDataFlowsRequest
	}

	var requests []ListDataFlowsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListDataFlowsRequest)
				return c.ListDataFlows(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListDataFlowsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListDataFlowsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListDependentObjects(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListDependentObjects")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDependentObjects is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListDependentObjects", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListDependentObjects")
	assert.NoError(t, err)

	type ListDependentObjectsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListDependentObjectsRequest
	}

	var requests []ListDependentObjectsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListDependentObjectsRequest)
				return c.ListDependentObjects(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListDependentObjectsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListDependentObjectsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListExternalPublicationValidations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListExternalPublicationValidations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListExternalPublicationValidations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListExternalPublicationValidations", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListExternalPublicationValidations")
	assert.NoError(t, err)

	type ListExternalPublicationValidationsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListExternalPublicationValidationsRequest
	}

	var requests []ListExternalPublicationValidationsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListExternalPublicationValidationsRequest)
				return c.ListExternalPublicationValidations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListExternalPublicationValidationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListExternalPublicationValidationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListExternalPublications(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListExternalPublications")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListExternalPublications is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListExternalPublications", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListExternalPublications")
	assert.NoError(t, err)

	type ListExternalPublicationsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListExternalPublicationsRequest
	}

	var requests []ListExternalPublicationsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListExternalPublicationsRequest)
				return c.ListExternalPublications(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListExternalPublicationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListExternalPublicationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListFolders(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListFolders")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListFolders is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListFolders", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListFolders")
	assert.NoError(t, err)

	type ListFoldersRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListFoldersRequest
	}

	var requests []ListFoldersRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListFoldersRequest)
				return c.ListFolders(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListFoldersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListFoldersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListPatchChanges(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListPatchChanges")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPatchChanges is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListPatchChanges", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListPatchChanges")
	assert.NoError(t, err)

	type ListPatchChangesRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListPatchChangesRequest
	}

	var requests []ListPatchChangesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListPatchChangesRequest)
				return c.ListPatchChanges(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListPatchChangesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListPatchChangesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListPatches(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListPatches")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPatches is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListPatches", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListPatches")
	assert.NoError(t, err)

	type ListPatchesRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListPatchesRequest
	}

	var requests []ListPatchesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListPatchesRequest)
				return c.ListPatches(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListPatchesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListPatchesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListProjects(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListProjects")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListProjects is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListProjects", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListProjects")
	assert.NoError(t, err)

	type ListProjectsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListProjectsRequest
	}

	var requests []ListProjectsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListProjectsRequest)
				return c.ListProjects(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListProjectsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListProjectsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListPublishedObjects(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListPublishedObjects")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPublishedObjects is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListPublishedObjects", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListPublishedObjects")
	assert.NoError(t, err)

	type ListPublishedObjectsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListPublishedObjectsRequest
	}

	var requests []ListPublishedObjectsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListPublishedObjectsRequest)
				return c.ListPublishedObjects(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListPublishedObjectsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListPublishedObjectsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListReferences(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListReferences")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListReferences is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListReferences", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListReferences")
	assert.NoError(t, err)

	type ListReferencesRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListReferencesRequest
	}

	var requests []ListReferencesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListReferencesRequest)
				return c.ListReferences(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListReferencesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListReferencesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListSchemas(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListSchemas")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSchemas is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListSchemas", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListSchemas")
	assert.NoError(t, err)

	type ListSchemasRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListSchemasRequest
	}

	var requests []ListSchemasRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListSchemasRequest)
				return c.ListSchemas(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListSchemasResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListSchemasResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListTaskRunLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListTaskRunLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTaskRunLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListTaskRunLogs", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListTaskRunLogs")
	assert.NoError(t, err)

	type ListTaskRunLogsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListTaskRunLogsRequest
	}

	var requests []ListTaskRunLogsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListTaskRunLogsRequest)
				return c.ListTaskRunLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListTaskRunLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListTaskRunLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListTaskRuns(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListTaskRuns")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTaskRuns is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListTaskRuns", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListTaskRuns")
	assert.NoError(t, err)

	type ListTaskRunsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListTaskRunsRequest
	}

	var requests []ListTaskRunsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListTaskRunsRequest)
				return c.ListTaskRuns(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListTaskRunsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListTaskRunsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListTaskValidations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListTaskValidations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTaskValidations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListTaskValidations", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListTaskValidations")
	assert.NoError(t, err)

	type ListTaskValidationsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListTaskValidationsRequest
	}

	var requests []ListTaskValidationsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListTaskValidationsRequest)
				return c.ListTaskValidations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListTaskValidationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListTaskValidationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListTasks(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListTasks")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTasks is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListTasks", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListTasks")
	assert.NoError(t, err)

	type ListTasksRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListTasksRequest
	}

	var requests []ListTasksRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListTasksRequest)
				return c.ListTasks(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListTasksResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListTasksResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListWorkRequestErrors", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListWorkRequestErrorsRequest
	}

	var requests []ListWorkRequestErrorsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListWorkRequestErrorsRequest)
				return c.ListWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListWorkRequestLogs", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListWorkRequestLogsRequest
	}

	var requests []ListWorkRequestLogsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListWorkRequestLogsRequest)
				return c.ListWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListWorkRequests", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListWorkRequestsRequest
	}

	var requests []ListWorkRequestsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientListWorkspaces(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "ListWorkspaces")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkspaces is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "ListWorkspaces", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "ListWorkspaces")
	assert.NoError(t, err)

	type ListWorkspacesRequestInfo struct {
		ContainerId string
		Request     dataintegration.ListWorkspacesRequest
	}

	var requests []ListWorkspacesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dataintegration.ListWorkspacesRequest)
				return c.ListWorkspaces(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataintegration.ListWorkspacesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataintegration.ListWorkspacesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientStartWorkspace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "StartWorkspace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StartWorkspace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "StartWorkspace", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "StartWorkspace")
	assert.NoError(t, err)

	type StartWorkspaceRequestInfo struct {
		ContainerId string
		Request     dataintegration.StartWorkspaceRequest
	}

	var requests []StartWorkspaceRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.StartWorkspace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientStopWorkspace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "StopWorkspace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StopWorkspace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "StopWorkspace", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "StopWorkspace")
	assert.NoError(t, err)

	type StopWorkspaceRequestInfo struct {
		ContainerId string
		Request     dataintegration.StopWorkspaceRequest
	}

	var requests []StopWorkspaceRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.StopWorkspace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientUpdateApplication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "UpdateApplication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateApplication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "UpdateApplication", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "UpdateApplication")
	assert.NoError(t, err)

	type UpdateApplicationRequestInfo struct {
		ContainerId string
		Request     dataintegration.UpdateApplicationRequest
	}

	var requests []UpdateApplicationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateApplication(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientUpdateConnection(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "UpdateConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "UpdateConnection", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "UpdateConnection")
	assert.NoError(t, err)

	type UpdateConnectionRequestInfo struct {
		ContainerId string
		Request     dataintegration.UpdateConnectionRequest
	}

	var requests []UpdateConnectionRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]UpdateConnectionRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["UpdateConnectionDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "modelType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"GENERIC_JDBC_CONNECTION":          &dataintegration.UpdateConnectionFromJdbc{},
				"ORACLE_OBJECT_STORAGE_CONNECTION": &dataintegration.UpdateConnectionFromObjectStorage{},
				"ORACLE_ATP_CONNECTION":            &dataintegration.UpdateConnectionFromAtp{},
				"ORACLEDB_CONNECTION":              &dataintegration.UpdateConnectionFromOracle{},
				"ORACLE_ADWC_CONNECTION":           &dataintegration.UpdateConnectionFromAdwc{},
				"MYSQL_CONNECTION":                 &dataintegration.UpdateConnectionFromMySql{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientUpdateDataAsset(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "UpdateDataAsset")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDataAsset is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "UpdateDataAsset", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "UpdateDataAsset")
	assert.NoError(t, err)

	type UpdateDataAssetRequestInfo struct {
		ContainerId string
		Request     dataintegration.UpdateDataAssetRequest
	}

	var requests []UpdateDataAssetRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]UpdateDataAssetRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["UpdateDataAssetDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "modelType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"ORACLE_ATP_DATA_ASSET":            &dataintegration.UpdateDataAssetFromAtp{},
				"ORACLE_ADWC_DATA_ASSET":           &dataintegration.UpdateDataAssetFromAdwc{},
				"GENERIC_JDBC_DATA_ASSET":          &dataintegration.UpdateDataAssetFromJdbc{},
				"ORACLE_OBJECT_STORAGE_DATA_ASSET": &dataintegration.UpdateDataAssetFromObjectStorage{},
				"MYSQL_DATA_ASSET":                 &dataintegration.UpdateDataAssetFromMySql{},
				"ORACLE_DATA_ASSET":                &dataintegration.UpdateDataAssetFromOracle{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateDataAsset(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientUpdateDataFlow(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "UpdateDataFlow")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDataFlow is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "UpdateDataFlow", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "UpdateDataFlow")
	assert.NoError(t, err)

	type UpdateDataFlowRequestInfo struct {
		ContainerId string
		Request     dataintegration.UpdateDataFlowRequest
	}

	var requests []UpdateDataFlowRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateDataFlow(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientUpdateExternalPublication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "UpdateExternalPublication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateExternalPublication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "UpdateExternalPublication", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "UpdateExternalPublication")
	assert.NoError(t, err)

	type UpdateExternalPublicationRequestInfo struct {
		ContainerId string
		Request     dataintegration.UpdateExternalPublicationRequest
	}

	var requests []UpdateExternalPublicationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateExternalPublication(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientUpdateFolder(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "UpdateFolder")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateFolder is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "UpdateFolder", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "UpdateFolder")
	assert.NoError(t, err)

	type UpdateFolderRequestInfo struct {
		ContainerId string
		Request     dataintegration.UpdateFolderRequest
	}

	var requests []UpdateFolderRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateFolder(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientUpdateProject(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "UpdateProject")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateProject is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "UpdateProject", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "UpdateProject")
	assert.NoError(t, err)

	type UpdateProjectRequestInfo struct {
		ContainerId string
		Request     dataintegration.UpdateProjectRequest
	}

	var requests []UpdateProjectRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateProject(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientUpdateReference(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "UpdateReference")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateReference is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "UpdateReference", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "UpdateReference")
	assert.NoError(t, err)

	type UpdateReferenceRequestInfo struct {
		ContainerId string
		Request     dataintegration.UpdateReferenceRequest
	}

	var requests []UpdateReferenceRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateReference(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientUpdateTask(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "UpdateTask")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTask is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "UpdateTask", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "UpdateTask")
	assert.NoError(t, err)

	type UpdateTaskRequestInfo struct {
		ContainerId string
		Request     dataintegration.UpdateTaskRequest
	}

	var requests []UpdateTaskRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]UpdateTaskRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["UpdateTaskDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "modelType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"DATA_LOADER_TASK": &dataintegration.UpdateTaskFromDataLoaderTask{},
				"INTEGRATION_TASK": &dataintegration.UpdateTaskFromIntegrationTask{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateTask(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientUpdateTaskRun(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "UpdateTaskRun")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTaskRun is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "UpdateTaskRun", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "UpdateTaskRun")
	assert.NoError(t, err)

	type UpdateTaskRunRequestInfo struct {
		ContainerId string
		Request     dataintegration.UpdateTaskRunRequest
	}

	var requests []UpdateTaskRunRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateTaskRun(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="di_dis_ww_grp@oracle.com" jiraProject="DI" opsJiraProject="DIS"
func TestDataIntegrationClientUpdateWorkspace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataintegration", "UpdateWorkspace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateWorkspace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataintegration", "DataIntegration", "UpdateWorkspace", createDataIntegrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataintegration.DataIntegrationClient)

	body, err := testClient.getRequests("dataintegration", "UpdateWorkspace")
	assert.NoError(t, err)

	type UpdateWorkspaceRequestInfo struct {
		ContainerId string
		Request     dataintegration.UpdateWorkspaceRequest
	}

	var requests []UpdateWorkspaceRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateWorkspace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
