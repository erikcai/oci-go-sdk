package autotest

import (
	"github.com/oracle/oci-go-sdk/v27/bds"
	"github.com/oracle/oci-go-sdk/v27/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createBdsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := bds.NewBdsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientAddAutoScalingConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "AddAutoScalingConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AddAutoScalingConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "AddAutoScalingConfiguration", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "AddAutoScalingConfiguration")
	assert.NoError(t, err)

	type AddAutoScalingConfigurationRequestInfo struct {
		ContainerId string
		Request     bds.AddAutoScalingConfigurationRequest
	}

	var requests []AddAutoScalingConfigurationRequestInfo
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
			response, err := c.AddAutoScalingConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientAddBlockStorage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "AddBlockStorage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AddBlockStorage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "AddBlockStorage", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "AddBlockStorage")
	assert.NoError(t, err)

	type AddBlockStorageRequestInfo struct {
		ContainerId string
		Request     bds.AddBlockStorageRequest
	}

	var requests []AddBlockStorageRequestInfo
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
			response, err := c.AddBlockStorage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientAddCloudSql(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "AddCloudSql")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AddCloudSql is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "AddCloudSql", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "AddCloudSql")
	assert.NoError(t, err)

	type AddCloudSqlRequestInfo struct {
		ContainerId string
		Request     bds.AddCloudSqlRequest
	}

	var requests []AddCloudSqlRequestInfo
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
			response, err := c.AddCloudSql(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientAddWorkerNodes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "AddWorkerNodes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AddWorkerNodes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "AddWorkerNodes", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "AddWorkerNodes")
	assert.NoError(t, err)

	type AddWorkerNodesRequestInfo struct {
		ContainerId string
		Request     bds.AddWorkerNodesRequest
	}

	var requests []AddWorkerNodesRequestInfo
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
			response, err := c.AddWorkerNodes(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientChangeBdsInstanceCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "ChangeBdsInstanceCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeBdsInstanceCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "ChangeBdsInstanceCompartment", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "ChangeBdsInstanceCompartment")
	assert.NoError(t, err)

	type ChangeBdsInstanceCompartmentRequestInfo struct {
		ContainerId string
		Request     bds.ChangeBdsInstanceCompartmentRequest
	}

	var requests []ChangeBdsInstanceCompartmentRequestInfo
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
			response, err := c.ChangeBdsInstanceCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientChangeShape(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "ChangeShape")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeShape is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "ChangeShape", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "ChangeShape")
	assert.NoError(t, err)

	type ChangeShapeRequestInfo struct {
		ContainerId string
		Request     bds.ChangeShapeRequest
	}

	var requests []ChangeShapeRequestInfo
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
			response, err := c.ChangeShape(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientCreateBdsInstance(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "CreateBdsInstance")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBdsInstance is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "CreateBdsInstance", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "CreateBdsInstance")
	assert.NoError(t, err)

	type CreateBdsInstanceRequestInfo struct {
		ContainerId string
		Request     bds.CreateBdsInstanceRequest
	}

	var requests []CreateBdsInstanceRequestInfo
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
			response, err := c.CreateBdsInstance(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientDeleteBdsInstance(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "DeleteBdsInstance")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBdsInstance is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "DeleteBdsInstance", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "DeleteBdsInstance")
	assert.NoError(t, err)

	type DeleteBdsInstanceRequestInfo struct {
		ContainerId string
		Request     bds.DeleteBdsInstanceRequest
	}

	var requests []DeleteBdsInstanceRequestInfo
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
			response, err := c.DeleteBdsInstance(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientGetAutoScalingConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "GetAutoScalingConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutoScalingConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "GetAutoScalingConfiguration", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "GetAutoScalingConfiguration")
	assert.NoError(t, err)

	type GetAutoScalingConfigurationRequestInfo struct {
		ContainerId string
		Request     bds.GetAutoScalingConfigurationRequest
	}

	var requests []GetAutoScalingConfigurationRequestInfo
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
			response, err := c.GetAutoScalingConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientGetBdsInstance(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "GetBdsInstance")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBdsInstance is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "GetBdsInstance", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "GetBdsInstance")
	assert.NoError(t, err)

	type GetBdsInstanceRequestInfo struct {
		ContainerId string
		Request     bds.GetBdsInstanceRequest
	}

	var requests []GetBdsInstanceRequestInfo
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
			response, err := c.GetBdsInstance(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "GetWorkRequest", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     bds.GetWorkRequestRequest
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

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientListAutoScalingConfigurations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "ListAutoScalingConfigurations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutoScalingConfigurations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "ListAutoScalingConfigurations", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "ListAutoScalingConfigurations")
	assert.NoError(t, err)

	type ListAutoScalingConfigurationsRequestInfo struct {
		ContainerId string
		Request     bds.ListAutoScalingConfigurationsRequest
	}

	var requests []ListAutoScalingConfigurationsRequestInfo
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
				r := req.(*bds.ListAutoScalingConfigurationsRequest)
				return c.ListAutoScalingConfigurations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]bds.ListAutoScalingConfigurationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(bds.ListAutoScalingConfigurationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientListBdsInstances(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "ListBdsInstances")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBdsInstances is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "ListBdsInstances", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "ListBdsInstances")
	assert.NoError(t, err)

	type ListBdsInstancesRequestInfo struct {
		ContainerId string
		Request     bds.ListBdsInstancesRequest
	}

	var requests []ListBdsInstancesRequestInfo
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
				r := req.(*bds.ListBdsInstancesRequest)
				return c.ListBdsInstances(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]bds.ListBdsInstancesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(bds.ListBdsInstancesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "ListWorkRequestErrors", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     bds.ListWorkRequestErrorsRequest
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
				r := req.(*bds.ListWorkRequestErrorsRequest)
				return c.ListWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]bds.ListWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(bds.ListWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "ListWorkRequestLogs", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     bds.ListWorkRequestLogsRequest
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
				r := req.(*bds.ListWorkRequestLogsRequest)
				return c.ListWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]bds.ListWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(bds.ListWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "ListWorkRequests", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     bds.ListWorkRequestsRequest
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
				r := req.(*bds.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]bds.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(bds.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientRemoveAutoScalingConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "RemoveAutoScalingConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RemoveAutoScalingConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "RemoveAutoScalingConfiguration", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "RemoveAutoScalingConfiguration")
	assert.NoError(t, err)

	type RemoveAutoScalingConfigurationRequestInfo struct {
		ContainerId string
		Request     bds.RemoveAutoScalingConfigurationRequest
	}

	var requests []RemoveAutoScalingConfigurationRequestInfo
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
			response, err := c.RemoveAutoScalingConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientRemoveCloudSql(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "RemoveCloudSql")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RemoveCloudSql is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "RemoveCloudSql", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "RemoveCloudSql")
	assert.NoError(t, err)

	type RemoveCloudSqlRequestInfo struct {
		ContainerId string
		Request     bds.RemoveCloudSqlRequest
	}

	var requests []RemoveCloudSqlRequestInfo
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
			response, err := c.RemoveCloudSql(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientRestartNode(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "RestartNode")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestartNode is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "RestartNode", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "RestartNode")
	assert.NoError(t, err)

	type RestartNodeRequestInfo struct {
		ContainerId string
		Request     bds.RestartNodeRequest
	}

	var requests []RestartNodeRequestInfo
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
			response, err := c.RestartNode(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientUpdateAutoScalingConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "UpdateAutoScalingConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutoScalingConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "UpdateAutoScalingConfiguration", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "UpdateAutoScalingConfiguration")
	assert.NoError(t, err)

	type UpdateAutoScalingConfigurationRequestInfo struct {
		ContainerId string
		Request     bds.UpdateAutoScalingConfigurationRequest
	}

	var requests []UpdateAutoScalingConfigurationRequestInfo
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
			response, err := c.UpdateAutoScalingConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bdcs_dev_ww_grp@oracle.com" jiraProject="BDCS" opsJiraProject="OBDS"
func TestBdsClientUpdateBdsInstance(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bds", "UpdateBdsInstance")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBdsInstance is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bds", "Bds", "UpdateBdsInstance", createBdsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bds.BdsClient)

	body, err := testClient.getRequests("bds", "UpdateBdsInstance")
	assert.NoError(t, err)

	type UpdateBdsInstanceRequestInfo struct {
		ContainerId string
		Request     bds.UpdateBdsInstanceRequest
	}

	var requests []UpdateBdsInstanceRequestInfo
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
			response, err := c.UpdateBdsInstance(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
