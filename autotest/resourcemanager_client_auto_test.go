package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/resourcemanager"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientCancelJob(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "CancelJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CancelJob is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "CancelJob")
	assert.NoError(t, err)

	type CancelJobRequestInfo struct {
		ContainerId string
		Request     resourcemanager.CancelJobRequest
	}

	var requests []CancelJobRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CancelJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientCreateJob(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "CreateJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateJob is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "CreateJob")
	assert.NoError(t, err)

	type CreateJobRequestInfo struct {
		ContainerId string
		Request     resourcemanager.CreateJobRequest
	}

	var requests []CreateJobRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientCreateStack(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "CreateStack")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateStack is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "CreateStack")
	assert.NoError(t, err)

	type CreateStackRequestInfo struct {
		ContainerId string
		Request     resourcemanager.CreateStackRequest
	}

	var requests []CreateStackRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateStack(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientDeleteStack(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "DeleteStack")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteStack is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "DeleteStack")
	assert.NoError(t, err)

	type DeleteStackRequestInfo struct {
		ContainerId string
		Request     resourcemanager.DeleteStackRequest
	}

	var requests []DeleteStackRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteStack(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientGetJob(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJob is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "GetJob")
	assert.NoError(t, err)

	type GetJobRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobRequest
	}

	var requests []GetJobRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientGetJobLogs(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJobLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJobLogs is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "GetJobLogs")
	assert.NoError(t, err)

	type GetJobLogsRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobLogsRequest
	}

	var requests []GetJobLogsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*resourcemanager.GetJobLogsRequest)
				return c.GetJobLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.GetJobLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.GetJobLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientGetJobLogsContent(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJobLogsContent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJobLogsContent is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "GetJobLogsContent")
	assert.NoError(t, err)

	type GetJobLogsContentRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobLogsContentRequest
	}

	var requests []GetJobLogsContentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetJobLogsContent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientGetJobTfConfig(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJobTfConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJobTfConfig is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "GetJobTfConfig")
	assert.NoError(t, err)

	type GetJobTfConfigRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobTfConfigRequest
	}

	var requests []GetJobTfConfigRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetJobTfConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientGetJobTfState(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJobTfState")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJobTfState is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "GetJobTfState")
	assert.NoError(t, err)

	type GetJobTfStateRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobTfStateRequest
	}

	var requests []GetJobTfStateRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetJobTfState(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientGetStack(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "GetStack")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStack is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "GetStack")
	assert.NoError(t, err)

	type GetStackRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetStackRequest
	}

	var requests []GetStackRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetStack(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientGetStackTfConfig(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "GetStackTfConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStackTfConfig is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "GetStackTfConfig")
	assert.NoError(t, err)

	type GetStackTfConfigRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetStackTfConfigRequest
	}

	var requests []GetStackTfConfigRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetStackTfConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientListJobs(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "ListJobs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListJobs is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "ListJobs")
	assert.NoError(t, err)

	type ListJobsRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListJobsRequest
	}

	var requests []ListJobsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*resourcemanager.ListJobsRequest)
				return c.ListJobs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.ListJobsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.ListJobsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientListStacks(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "ListStacks")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListStacks is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "ListStacks")
	assert.NoError(t, err)

	type ListStacksRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListStacksRequest
	}

	var requests []ListStacksRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*resourcemanager.ListStacksRequest)
				return c.ListStacks(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.ListStacksResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.ListStacksResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientUpdateJob(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "UpdateJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateJob is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "UpdateJob")
	assert.NoError(t, err)

	type UpdateJobRequestInfo struct {
		ContainerId string
		Request     resourcemanager.UpdateJobRequest
	}

	var requests []UpdateJobRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestResourceManagerClientUpdateStack(t *testing.T) {
	enabled, err := testClient.isApiEnabled("resourcemanager", "UpdateStack")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateStack is not enabled by the testing service")
	}
	c, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("resourcemanager", "UpdateStack")
	assert.NoError(t, err)

	type UpdateStackRequestInfo struct {
		ContainerId string
		Request     resourcemanager.UpdateStackRequest
	}

	var requests []UpdateStackRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateStack(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
