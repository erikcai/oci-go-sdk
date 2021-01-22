package autotest

import (
	"github.com/erikcai/oci-go-sdk/v33/clouddeploy"
	"github.com/erikcai/oci-go-sdk/v33/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createClouddeployPipelineClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := clouddeploy.NewPipelineClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployPipelineClientCreatePipeline(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "CreatePipeline")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreatePipeline is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Pipeline", "CreatePipeline", createClouddeployPipelineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.PipelineClient)

	body, err := testClient.getRequests("clouddeploy", "CreatePipeline")
	assert.NoError(t, err)

	type CreatePipelineRequestInfo struct {
		ContainerId string
		Request     clouddeploy.CreatePipelineRequest
	}

	var requests []CreatePipelineRequestInfo
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
			response, err := c.CreatePipeline(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployPipelineClientDeletePipeline(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "DeletePipeline")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeletePipeline is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Pipeline", "DeletePipeline", createClouddeployPipelineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.PipelineClient)

	body, err := testClient.getRequests("clouddeploy", "DeletePipeline")
	assert.NoError(t, err)

	type DeletePipelineRequestInfo struct {
		ContainerId string
		Request     clouddeploy.DeletePipelineRequest
	}

	var requests []DeletePipelineRequestInfo
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
			response, err := c.DeletePipeline(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployPipelineClientGetPipeline(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "GetPipeline")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPipeline is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Pipeline", "GetPipeline", createClouddeployPipelineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.PipelineClient)

	body, err := testClient.getRequests("clouddeploy", "GetPipeline")
	assert.NoError(t, err)

	type GetPipelineRequestInfo struct {
		ContainerId string
		Request     clouddeploy.GetPipelineRequest
	}

	var requests []GetPipelineRequestInfo
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
			response, err := c.GetPipeline(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployPipelineClientListPipelines(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "ListPipelines")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPipelines is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Pipeline", "ListPipelines", createClouddeployPipelineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.PipelineClient)

	body, err := testClient.getRequests("clouddeploy", "ListPipelines")
	assert.NoError(t, err)

	type ListPipelinesRequestInfo struct {
		ContainerId string
		Request     clouddeploy.ListPipelinesRequest
	}

	var requests []ListPipelinesRequestInfo
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
				r := req.(*clouddeploy.ListPipelinesRequest)
				return c.ListPipelines(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]clouddeploy.ListPipelinesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(clouddeploy.ListPipelinesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployPipelineClientUpdatePipeline(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "UpdatePipeline")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdatePipeline is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Pipeline", "UpdatePipeline", createClouddeployPipelineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.PipelineClient)

	body, err := testClient.getRequests("clouddeploy", "UpdatePipeline")
	assert.NoError(t, err)

	type UpdatePipelineRequestInfo struct {
		ContainerId string
		Request     clouddeploy.UpdatePipelineRequest
	}

	var requests []UpdatePipelineRequestInfo
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
			response, err := c.UpdatePipeline(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
