package autotest

import (
	"github.com/oracle/oci-go-sdk/v30/common"
	"github.com/oracle/oci-go-sdk/v30/operatoraccesscontrol"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createOperatoraccesscontrolAccessRequestsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := operatoraccesscontrol.NewAccessRequestsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolAccessRequestsClientApproveAccessRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "ApproveAccessRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ApproveAccessRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "AccessRequests", "ApproveAccessRequest", createOperatoraccesscontrolAccessRequestsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.AccessRequestsClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "ApproveAccessRequest")
	assert.NoError(t, err)

	type ApproveAccessRequestRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.ApproveAccessRequestRequest
	}

	var requests []ApproveAccessRequestRequestInfo
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
			response, err := c.ApproveAccessRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolAccessRequestsClientGetAccessRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "GetAccessRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAccessRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "AccessRequests", "GetAccessRequest", createOperatoraccesscontrolAccessRequestsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.AccessRequestsClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "GetAccessRequest")
	assert.NoError(t, err)

	type GetAccessRequestRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.GetAccessRequestRequest
	}

	var requests []GetAccessRequestRequestInfo
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
			response, err := c.GetAccessRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolAccessRequestsClientListAccessRequestHistories(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "ListAccessRequestHistories")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAccessRequestHistories is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "AccessRequests", "ListAccessRequestHistories", createOperatoraccesscontrolAccessRequestsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.AccessRequestsClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "ListAccessRequestHistories")
	assert.NoError(t, err)

	type ListAccessRequestHistoriesRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.ListAccessRequestHistoriesRequest
	}

	var requests []ListAccessRequestHistoriesRequestInfo
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
				r := req.(*operatoraccesscontrol.ListAccessRequestHistoriesRequest)
				return c.ListAccessRequestHistories(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operatoraccesscontrol.ListAccessRequestHistoriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operatoraccesscontrol.ListAccessRequestHistoriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolAccessRequestsClientListAccessRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "ListAccessRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAccessRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "AccessRequests", "ListAccessRequests", createOperatoraccesscontrolAccessRequestsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.AccessRequestsClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "ListAccessRequests")
	assert.NoError(t, err)

	type ListAccessRequestsRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.ListAccessRequestsRequest
	}

	var requests []ListAccessRequestsRequestInfo
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
				r := req.(*operatoraccesscontrol.ListAccessRequestsRequest)
				return c.ListAccessRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operatoraccesscontrol.ListAccessRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operatoraccesscontrol.ListAccessRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolAccessRequestsClientRejectAccessRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "RejectAccessRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RejectAccessRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "AccessRequests", "RejectAccessRequest", createOperatoraccesscontrolAccessRequestsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.AccessRequestsClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "RejectAccessRequest")
	assert.NoError(t, err)

	type RejectAccessRequestRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.RejectAccessRequestRequest
	}

	var requests []RejectAccessRequestRequestInfo
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
			response, err := c.RejectAccessRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolAccessRequestsClientRevokeAccessRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "RevokeAccessRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RevokeAccessRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "AccessRequests", "RevokeAccessRequest", createOperatoraccesscontrolAccessRequestsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.AccessRequestsClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "RevokeAccessRequest")
	assert.NoError(t, err)

	type RevokeAccessRequestRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.RevokeAccessRequestRequest
	}

	var requests []RevokeAccessRequestRequestInfo
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
			response, err := c.RevokeAccessRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
