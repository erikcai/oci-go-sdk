package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/dataflow"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createDataFlowClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := dataflow.NewDataFlowClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="sss_dev_ww_grp@oracle.com" jiraProject="SSS" opsJiraProject="SSS"
func TestDataFlowClientChangeApplicationCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataflow", "ChangeApplicationCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeApplicationCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataflow", "DataFlow", "ChangeApplicationCompartment", createDataFlowClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataflow.DataFlowClient)

	body, err := testClient.getRequests("dataflow", "ChangeApplicationCompartment")
	assert.NoError(t, err)

	type ChangeApplicationCompartmentRequestInfo struct {
		ContainerId string
		Request     dataflow.ChangeApplicationCompartmentRequest
	}

	var requests []ChangeApplicationCompartmentRequestInfo
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
			response, err := c.ChangeApplicationCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sss_dev_ww_grp@oracle.com" jiraProject="SSS" opsJiraProject="SSS"
func TestDataFlowClientChangeRunCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataflow", "ChangeRunCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeRunCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataflow", "DataFlow", "ChangeRunCompartment", createDataFlowClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataflow.DataFlowClient)

	body, err := testClient.getRequests("dataflow", "ChangeRunCompartment")
	assert.NoError(t, err)

	type ChangeRunCompartmentRequestInfo struct {
		ContainerId string
		Request     dataflow.ChangeRunCompartmentRequest
	}

	var requests []ChangeRunCompartmentRequestInfo
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
			response, err := c.ChangeRunCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sss_dev_ww_grp@oracle.com" jiraProject="SSS" opsJiraProject="SSS"
func TestDataFlowClientCreateApplication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataflow", "CreateApplication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateApplication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataflow", "DataFlow", "CreateApplication", createDataFlowClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataflow.DataFlowClient)

	body, err := testClient.getRequests("dataflow", "CreateApplication")
	assert.NoError(t, err)

	type CreateApplicationRequestInfo struct {
		ContainerId string
		Request     dataflow.CreateApplicationRequest
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

// IssueRoutingInfo tag="default" email="sss_dev_ww_grp@oracle.com" jiraProject="SSS" opsJiraProject="SSS"
func TestDataFlowClientCreateRun(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataflow", "CreateRun")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateRun is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataflow", "DataFlow", "CreateRun", createDataFlowClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataflow.DataFlowClient)

	body, err := testClient.getRequests("dataflow", "CreateRun")
	assert.NoError(t, err)

	type CreateRunRequestInfo struct {
		ContainerId string
		Request     dataflow.CreateRunRequest
	}

	var requests []CreateRunRequestInfo
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
			response, err := c.CreateRun(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sss_dev_ww_grp@oracle.com" jiraProject="SSS" opsJiraProject="SSS"
func TestDataFlowClientDeleteApplication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataflow", "DeleteApplication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteApplication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataflow", "DataFlow", "DeleteApplication", createDataFlowClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataflow.DataFlowClient)

	body, err := testClient.getRequests("dataflow", "DeleteApplication")
	assert.NoError(t, err)

	type DeleteApplicationRequestInfo struct {
		ContainerId string
		Request     dataflow.DeleteApplicationRequest
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

// IssueRoutingInfo tag="default" email="sss_dev_ww_grp@oracle.com" jiraProject="SSS" opsJiraProject="SSS"
func TestDataFlowClientDeleteRun(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataflow", "DeleteRun")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteRun is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataflow", "DataFlow", "DeleteRun", createDataFlowClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataflow.DataFlowClient)

	body, err := testClient.getRequests("dataflow", "DeleteRun")
	assert.NoError(t, err)

	type DeleteRunRequestInfo struct {
		ContainerId string
		Request     dataflow.DeleteRunRequest
	}

	var requests []DeleteRunRequestInfo
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
			response, err := c.DeleteRun(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sss_dev_ww_grp@oracle.com" jiraProject="SSS" opsJiraProject="SSS"
func TestDataFlowClientGetApplication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataflow", "GetApplication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetApplication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataflow", "DataFlow", "GetApplication", createDataFlowClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataflow.DataFlowClient)

	body, err := testClient.getRequests("dataflow", "GetApplication")
	assert.NoError(t, err)

	type GetApplicationRequestInfo struct {
		ContainerId string
		Request     dataflow.GetApplicationRequest
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

// IssueRoutingInfo tag="default" email="sss_dev_ww_grp@oracle.com" jiraProject="SSS" opsJiraProject="SSS"
func TestDataFlowClientGetRun(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataflow", "GetRun")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRun is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataflow", "DataFlow", "GetRun", createDataFlowClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataflow.DataFlowClient)

	body, err := testClient.getRequests("dataflow", "GetRun")
	assert.NoError(t, err)

	type GetRunRequestInfo struct {
		ContainerId string
		Request     dataflow.GetRunRequest
	}

	var requests []GetRunRequestInfo
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
			response, err := c.GetRun(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sss_dev_ww_grp@oracle.com" jiraProject="SSS" opsJiraProject="SSS"
func TestDataFlowClientGetRunLog(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataflow", "GetRunLog")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRunLog is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataflow", "DataFlow", "GetRunLog", createDataFlowClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataflow.DataFlowClient)

	body, err := testClient.getRequests("dataflow", "GetRunLog")
	assert.NoError(t, err)

	type GetRunLogRequestInfo struct {
		ContainerId string
		Request     dataflow.GetRunLogRequest
	}

	var requests []GetRunLogRequestInfo
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
			response, err := c.GetRunLog(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sss_dev_ww_grp@oracle.com" jiraProject="SSS" opsJiraProject="SSS"
func TestDataFlowClientListApplications(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataflow", "ListApplications")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListApplications is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataflow", "DataFlow", "ListApplications", createDataFlowClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataflow.DataFlowClient)

	body, err := testClient.getRequests("dataflow", "ListApplications")
	assert.NoError(t, err)

	type ListApplicationsRequestInfo struct {
		ContainerId string
		Request     dataflow.ListApplicationsRequest
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
				r := req.(*dataflow.ListApplicationsRequest)
				return c.ListApplications(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataflow.ListApplicationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataflow.ListApplicationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sss_dev_ww_grp@oracle.com" jiraProject="SSS" opsJiraProject="SSS"
func TestDataFlowClientListRunLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataflow", "ListRunLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRunLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataflow", "DataFlow", "ListRunLogs", createDataFlowClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataflow.DataFlowClient)

	body, err := testClient.getRequests("dataflow", "ListRunLogs")
	assert.NoError(t, err)

	type ListRunLogsRequestInfo struct {
		ContainerId string
		Request     dataflow.ListRunLogsRequest
	}

	var requests []ListRunLogsRequestInfo
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
				r := req.(*dataflow.ListRunLogsRequest)
				return c.ListRunLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataflow.ListRunLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataflow.ListRunLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sss_dev_ww_grp@oracle.com" jiraProject="SSS" opsJiraProject="SSS"
func TestDataFlowClientListRuns(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataflow", "ListRuns")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRuns is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataflow", "DataFlow", "ListRuns", createDataFlowClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataflow.DataFlowClient)

	body, err := testClient.getRequests("dataflow", "ListRuns")
	assert.NoError(t, err)

	type ListRunsRequestInfo struct {
		ContainerId string
		Request     dataflow.ListRunsRequest
	}

	var requests []ListRunsRequestInfo
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
				r := req.(*dataflow.ListRunsRequest)
				return c.ListRuns(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dataflow.ListRunsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dataflow.ListRunsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sss_dev_ww_grp@oracle.com" jiraProject="SSS" opsJiraProject="SSS"
func TestDataFlowClientUpdateApplication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataflow", "UpdateApplication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateApplication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataflow", "DataFlow", "UpdateApplication", createDataFlowClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataflow.DataFlowClient)

	body, err := testClient.getRequests("dataflow", "UpdateApplication")
	assert.NoError(t, err)

	type UpdateApplicationRequestInfo struct {
		ContainerId string
		Request     dataflow.UpdateApplicationRequest
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

// IssueRoutingInfo tag="default" email="sss_dev_ww_grp@oracle.com" jiraProject="SSS" opsJiraProject="SSS"
func TestDataFlowClientUpdateRun(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dataflow", "UpdateRun")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRun is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dataflow", "DataFlow", "UpdateRun", createDataFlowClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dataflow.DataFlowClient)

	body, err := testClient.getRequests("dataflow", "UpdateRun")
	assert.NoError(t, err)

	type UpdateRunRequestInfo struct {
		ContainerId string
		Request     dataflow.UpdateRunRequest
	}

	var requests []UpdateRunRequestInfo
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
			response, err := c.UpdateRun(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
