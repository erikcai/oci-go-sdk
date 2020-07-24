package autotest

import (
	"github.com/oracle/oci-go-sdk/bastions"
	"github.com/oracle/oci-go-sdk/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createBastionsSessionsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := bastions.NewSessionsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionsSessionsClientChangeSessionCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastions", "ChangeSessionCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeSessionCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastions", "Sessions", "ChangeSessionCompartment", createBastionsSessionsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastions.SessionsClient)

	body, err := testClient.getRequests("bastions", "ChangeSessionCompartment")
	assert.NoError(t, err)

	type ChangeSessionCompartmentRequestInfo struct {
		ContainerId string
		Request     bastions.ChangeSessionCompartmentRequest
	}

	var requests []ChangeSessionCompartmentRequestInfo
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
			response, err := c.ChangeSessionCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionsSessionsClientCreateSession(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastions", "CreateSession")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSession is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastions", "Sessions", "CreateSession", createBastionsSessionsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastions.SessionsClient)

	body, err := testClient.getRequests("bastions", "CreateSession")
	assert.NoError(t, err)

	type CreateSessionRequestInfo struct {
		ContainerId string
		Request     bastions.CreateSessionRequest
	}

	var requests []CreateSessionRequestInfo
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
			response, err := c.CreateSession(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionsSessionsClientDeleteSession(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastions", "DeleteSession")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSession is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastions", "Sessions", "DeleteSession", createBastionsSessionsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastions.SessionsClient)

	body, err := testClient.getRequests("bastions", "DeleteSession")
	assert.NoError(t, err)

	type DeleteSessionRequestInfo struct {
		ContainerId string
		Request     bastions.DeleteSessionRequest
	}

	var requests []DeleteSessionRequestInfo
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
			response, err := c.DeleteSession(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionsSessionsClientGetSession(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastions", "GetSession")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSession is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastions", "Sessions", "GetSession", createBastionsSessionsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastions.SessionsClient)

	body, err := testClient.getRequests("bastions", "GetSession")
	assert.NoError(t, err)

	type GetSessionRequestInfo struct {
		ContainerId string
		Request     bastions.GetSessionRequest
	}

	var requests []GetSessionRequestInfo
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
			response, err := c.GetSession(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionsSessionsClientListSessions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastions", "ListSessions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSessions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastions", "Sessions", "ListSessions", createBastionsSessionsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastions.SessionsClient)

	body, err := testClient.getRequests("bastions", "ListSessions")
	assert.NoError(t, err)

	type ListSessionsRequestInfo struct {
		ContainerId string
		Request     bastions.ListSessionsRequest
	}

	var requests []ListSessionsRequestInfo
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
				r := req.(*bastions.ListSessionsRequest)
				return c.ListSessions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]bastions.ListSessionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(bastions.ListSessionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionsSessionsClientUpdateSession(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastions", "UpdateSession")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSession is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastions", "Sessions", "UpdateSession", createBastionsSessionsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastions.SessionsClient)

	body, err := testClient.getRequests("bastions", "UpdateSession")
	assert.NoError(t, err)

	type UpdateSessionRequestInfo struct {
		ContainerId string
		Request     bastions.UpdateSessionRequest
	}

	var requests []UpdateSessionRequestInfo
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
			response, err := c.UpdateSession(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
