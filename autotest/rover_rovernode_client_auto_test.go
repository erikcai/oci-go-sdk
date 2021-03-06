package autotest

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"github.com/erikcai/oci-go-sdk/v33/rover"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createRoverRoverNodeClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := rover.NewRoverNodeClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverNodeClientChangeRoverNodeCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "ChangeRoverNodeCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeRoverNodeCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverNode", "ChangeRoverNodeCompartment", createRoverRoverNodeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverNodeClient)

	body, err := testClient.getRequests("rover", "ChangeRoverNodeCompartment")
	assert.NoError(t, err)

	type ChangeRoverNodeCompartmentRequestInfo struct {
		ContainerId string
		Request     rover.ChangeRoverNodeCompartmentRequest
	}

	var requests []ChangeRoverNodeCompartmentRequestInfo
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
			response, err := c.ChangeRoverNodeCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverNodeClientCreateRoverNode(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "CreateRoverNode")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateRoverNode is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverNode", "CreateRoverNode", createRoverRoverNodeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverNodeClient)

	body, err := testClient.getRequests("rover", "CreateRoverNode")
	assert.NoError(t, err)

	type CreateRoverNodeRequestInfo struct {
		ContainerId string
		Request     rover.CreateRoverNodeRequest
	}

	var requests []CreateRoverNodeRequestInfo
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
			response, err := c.CreateRoverNode(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverNodeClientDeleteRoverNode(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "DeleteRoverNode")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteRoverNode is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverNode", "DeleteRoverNode", createRoverRoverNodeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverNodeClient)

	body, err := testClient.getRequests("rover", "DeleteRoverNode")
	assert.NoError(t, err)

	type DeleteRoverNodeRequestInfo struct {
		ContainerId string
		Request     rover.DeleteRoverNodeRequest
	}

	var requests []DeleteRoverNodeRequestInfo
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
			response, err := c.DeleteRoverNode(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverNodeClientGetRoverNode(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "GetRoverNode")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRoverNode is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverNode", "GetRoverNode", createRoverRoverNodeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverNodeClient)

	body, err := testClient.getRequests("rover", "GetRoverNode")
	assert.NoError(t, err)

	type GetRoverNodeRequestInfo struct {
		ContainerId string
		Request     rover.GetRoverNodeRequest
	}

	var requests []GetRoverNodeRequestInfo
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
			response, err := c.GetRoverNode(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverNodeClientGetRoverNodeCertificate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "GetRoverNodeCertificate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRoverNodeCertificate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverNode", "GetRoverNodeCertificate", createRoverRoverNodeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverNodeClient)

	body, err := testClient.getRequests("rover", "GetRoverNodeCertificate")
	assert.NoError(t, err)

	type GetRoverNodeCertificateRequestInfo struct {
		ContainerId string
		Request     rover.GetRoverNodeCertificateRequest
	}

	var requests []GetRoverNodeCertificateRequestInfo
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
			response, err := c.GetRoverNodeCertificate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverNodeClientGetRoverNodeEncryptionKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "GetRoverNodeEncryptionKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRoverNodeEncryptionKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverNode", "GetRoverNodeEncryptionKey", createRoverRoverNodeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverNodeClient)

	body, err := testClient.getRequests("rover", "GetRoverNodeEncryptionKey")
	assert.NoError(t, err)

	type GetRoverNodeEncryptionKeyRequestInfo struct {
		ContainerId string
		Request     rover.GetRoverNodeEncryptionKeyRequest
	}

	var requests []GetRoverNodeEncryptionKeyRequestInfo
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
			response, err := c.GetRoverNodeEncryptionKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverNodeClientGetRoverNodeGetRpt(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "GetRoverNodeGetRpt")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRoverNodeGetRpt is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverNode", "GetRoverNodeGetRpt", createRoverRoverNodeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverNodeClient)

	body, err := testClient.getRequests("rover", "GetRoverNodeGetRpt")
	assert.NoError(t, err)

	type GetRoverNodeGetRptRequestInfo struct {
		ContainerId string
		Request     rover.GetRoverNodeGetRptRequest
	}

	var requests []GetRoverNodeGetRptRequestInfo
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
			response, err := c.GetRoverNodeGetRpt(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverNodeClientListRoverNodes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "ListRoverNodes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRoverNodes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverNode", "ListRoverNodes", createRoverRoverNodeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverNodeClient)

	body, err := testClient.getRequests("rover", "ListRoverNodes")
	assert.NoError(t, err)

	type ListRoverNodesRequestInfo struct {
		ContainerId string
		Request     rover.ListRoverNodesRequest
	}

	var requests []ListRoverNodesRequestInfo
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
				r := req.(*rover.ListRoverNodesRequest)
				return c.ListRoverNodes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]rover.ListRoverNodesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(rover.ListRoverNodesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverNodeClientRoverNodeActionSetKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "RoverNodeActionSetKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RoverNodeActionSetKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverNode", "RoverNodeActionSetKey", createRoverRoverNodeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverNodeClient)

	body, err := testClient.getRequests("rover", "RoverNodeActionSetKey")
	assert.NoError(t, err)

	type RoverNodeActionSetKeyRequestInfo struct {
		ContainerId string
		Request     rover.RoverNodeActionSetKeyRequest
	}

	var requests []RoverNodeActionSetKeyRequestInfo
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
			response, err := c.RoverNodeActionSetKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverNodeClientUpdateRoverNode(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "UpdateRoverNode")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRoverNode is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverNode", "UpdateRoverNode", createRoverRoverNodeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverNodeClient)

	body, err := testClient.getRequests("rover", "UpdateRoverNode")
	assert.NoError(t, err)

	type UpdateRoverNodeRequestInfo struct {
		ContainerId string
		Request     rover.UpdateRoverNodeRequest
	}

	var requests []UpdateRoverNodeRequestInfo
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
			response, err := c.UpdateRoverNode(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
