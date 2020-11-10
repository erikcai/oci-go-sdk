package autotest

import (
	"github.com/oracle/oci-go-sdk/v28/common"
	"github.com/oracle/oci-go-sdk/v28/computeinstanceagent"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createComputeinstanceagentComputeInstanceAgentClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := computeinstanceagent.NewComputeInstanceAgentClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="computeInstanceAgent" email="oci_osi_grp@oracle.com" jiraProject="OSI" opsJiraProject="IMAGE"
func TestComputeinstanceagentComputeInstanceAgentClientCancelInstanceAgentCommand(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("computeinstanceagent", "CancelInstanceAgentCommand")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CancelInstanceAgentCommand is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("computeinstanceagent", "ComputeInstanceAgent", "CancelInstanceAgentCommand", createComputeinstanceagentComputeInstanceAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(computeinstanceagent.ComputeInstanceAgentClient)

	body, err := testClient.getRequests("computeinstanceagent", "CancelInstanceAgentCommand")
	assert.NoError(t, err)

	type CancelInstanceAgentCommandRequestInfo struct {
		ContainerId string
		Request     computeinstanceagent.CancelInstanceAgentCommandRequest
	}

	var requests []CancelInstanceAgentCommandRequestInfo
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
			response, err := c.CancelInstanceAgentCommand(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeInstanceAgent" email="oci_osi_grp@oracle.com" jiraProject="OSI" opsJiraProject="IMAGE"
func TestComputeinstanceagentComputeInstanceAgentClientCreateInstanceAgentCommand(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("computeinstanceagent", "CreateInstanceAgentCommand")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateInstanceAgentCommand is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("computeinstanceagent", "ComputeInstanceAgent", "CreateInstanceAgentCommand", createComputeinstanceagentComputeInstanceAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(computeinstanceagent.ComputeInstanceAgentClient)

	body, err := testClient.getRequests("computeinstanceagent", "CreateInstanceAgentCommand")
	assert.NoError(t, err)

	type CreateInstanceAgentCommandRequestInfo struct {
		ContainerId string
		Request     computeinstanceagent.CreateInstanceAgentCommandRequest
	}

	var requests []CreateInstanceAgentCommandRequestInfo
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
			response, err := c.CreateInstanceAgentCommand(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeInstanceAgent" email="oci_osi_grp@oracle.com" jiraProject="OSI" opsJiraProject="IMAGE"
func TestComputeinstanceagentComputeInstanceAgentClientGetInstanceAgentCommand(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("computeinstanceagent", "GetInstanceAgentCommand")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInstanceAgentCommand is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("computeinstanceagent", "ComputeInstanceAgent", "GetInstanceAgentCommand", createComputeinstanceagentComputeInstanceAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(computeinstanceagent.ComputeInstanceAgentClient)

	body, err := testClient.getRequests("computeinstanceagent", "GetInstanceAgentCommand")
	assert.NoError(t, err)

	type GetInstanceAgentCommandRequestInfo struct {
		ContainerId string
		Request     computeinstanceagent.GetInstanceAgentCommandRequest
	}

	var requests []GetInstanceAgentCommandRequestInfo
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
			response, err := c.GetInstanceAgentCommand(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeInstanceAgent" email="oci_osi_grp@oracle.com" jiraProject="OSI" opsJiraProject="IMAGE"
func TestComputeinstanceagentComputeInstanceAgentClientGetInstanceAgentCommandExecution(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("computeinstanceagent", "GetInstanceAgentCommandExecution")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInstanceAgentCommandExecution is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("computeinstanceagent", "ComputeInstanceAgent", "GetInstanceAgentCommandExecution", createComputeinstanceagentComputeInstanceAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(computeinstanceagent.ComputeInstanceAgentClient)

	body, err := testClient.getRequests("computeinstanceagent", "GetInstanceAgentCommandExecution")
	assert.NoError(t, err)

	type GetInstanceAgentCommandExecutionRequestInfo struct {
		ContainerId string
		Request     computeinstanceagent.GetInstanceAgentCommandExecutionRequest
	}

	var requests []GetInstanceAgentCommandExecutionRequestInfo
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
			response, err := c.GetInstanceAgentCommandExecution(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeInstanceAgent" email="oci_osi_grp@oracle.com" jiraProject="OSI" opsJiraProject="IMAGE"
func TestComputeinstanceagentComputeInstanceAgentClientListInstanceAgentCommandExecutions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("computeinstanceagent", "ListInstanceAgentCommandExecutions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInstanceAgentCommandExecutions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("computeinstanceagent", "ComputeInstanceAgent", "ListInstanceAgentCommandExecutions", createComputeinstanceagentComputeInstanceAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(computeinstanceagent.ComputeInstanceAgentClient)

	body, err := testClient.getRequests("computeinstanceagent", "ListInstanceAgentCommandExecutions")
	assert.NoError(t, err)

	type ListInstanceAgentCommandExecutionsRequestInfo struct {
		ContainerId string
		Request     computeinstanceagent.ListInstanceAgentCommandExecutionsRequest
	}

	var requests []ListInstanceAgentCommandExecutionsRequestInfo
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
				r := req.(*computeinstanceagent.ListInstanceAgentCommandExecutionsRequest)
				return c.ListInstanceAgentCommandExecutions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]computeinstanceagent.ListInstanceAgentCommandExecutionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(computeinstanceagent.ListInstanceAgentCommandExecutionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeInstanceAgent" email="oci_osi_grp@oracle.com" jiraProject="OSI" opsJiraProject="IMAGE"
func TestComputeinstanceagentComputeInstanceAgentClientListInstanceAgentCommands(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("computeinstanceagent", "ListInstanceAgentCommands")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInstanceAgentCommands is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("computeinstanceagent", "ComputeInstanceAgent", "ListInstanceAgentCommands", createComputeinstanceagentComputeInstanceAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(computeinstanceagent.ComputeInstanceAgentClient)

	body, err := testClient.getRequests("computeinstanceagent", "ListInstanceAgentCommands")
	assert.NoError(t, err)

	type ListInstanceAgentCommandsRequestInfo struct {
		ContainerId string
		Request     computeinstanceagent.ListInstanceAgentCommandsRequest
	}

	var requests []ListInstanceAgentCommandsRequestInfo
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
				r := req.(*computeinstanceagent.ListInstanceAgentCommandsRequest)
				return c.ListInstanceAgentCommands(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]computeinstanceagent.ListInstanceAgentCommandsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(computeinstanceagent.ListInstanceAgentCommandsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
