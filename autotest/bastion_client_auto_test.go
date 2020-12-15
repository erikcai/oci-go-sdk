package autotest

import (
	"github.com/oracle/oci-go-sdk/v31/bastion"
	"github.com/oracle/oci-go-sdk/v31/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createBastionBastionClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := bastion.NewBastionClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionBastionClientChangeBastionCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "ChangeBastionCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeBastionCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "ChangeBastionCompartment", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "ChangeBastionCompartment")
	assert.NoError(t, err)

	type ChangeBastionCompartmentRequestInfo struct {
		ContainerId string
		Request     bastion.ChangeBastionCompartmentRequest
	}

	var requests []ChangeBastionCompartmentRequestInfo
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
			response, err := c.ChangeBastionCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionBastionClientCreateBastion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "CreateBastion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBastion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "CreateBastion", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "CreateBastion")
	assert.NoError(t, err)

	type CreateBastionRequestInfo struct {
		ContainerId string
		Request     bastion.CreateBastionRequest
	}

	var requests []CreateBastionRequestInfo
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
			response, err := c.CreateBastion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionBastionClientCreateSession(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "CreateSession")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSession is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "CreateSession", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "CreateSession")
	assert.NoError(t, err)

	type CreateSessionRequestInfo struct {
		ContainerId string
		Request     bastion.CreateSessionRequest
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
func TestBastionBastionClientDeleteBastion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "DeleteBastion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBastion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "DeleteBastion", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "DeleteBastion")
	assert.NoError(t, err)

	type DeleteBastionRequestInfo struct {
		ContainerId string
		Request     bastion.DeleteBastionRequest
	}

	var requests []DeleteBastionRequestInfo
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
			response, err := c.DeleteBastion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionBastionClientDeleteSession(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "DeleteSession")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSession is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "DeleteSession", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "DeleteSession")
	assert.NoError(t, err)

	type DeleteSessionRequestInfo struct {
		ContainerId string
		Request     bastion.DeleteSessionRequest
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
func TestBastionBastionClientGetBastion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "GetBastion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBastion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "GetBastion", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "GetBastion")
	assert.NoError(t, err)

	type GetBastionRequestInfo struct {
		ContainerId string
		Request     bastion.GetBastionRequest
	}

	var requests []GetBastionRequestInfo
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
			response, err := c.GetBastion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionBastionClientGetSession(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "GetSession")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSession is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "GetSession", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "GetSession")
	assert.NoError(t, err)

	type GetSessionRequestInfo struct {
		ContainerId string
		Request     bastion.GetSessionRequest
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
func TestBastionBastionClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "GetWorkRequest", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     bastion.GetWorkRequestRequest
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

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionBastionClientListBastions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "ListBastions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBastions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "ListBastions", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "ListBastions")
	assert.NoError(t, err)

	type ListBastionsRequestInfo struct {
		ContainerId string
		Request     bastion.ListBastionsRequest
	}

	var requests []ListBastionsRequestInfo
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
				r := req.(*bastion.ListBastionsRequest)
				return c.ListBastions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]bastion.ListBastionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(bastion.ListBastionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionBastionClientListSessions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "ListSessions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSessions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "ListSessions", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "ListSessions")
	assert.NoError(t, err)

	type ListSessionsRequestInfo struct {
		ContainerId string
		Request     bastion.ListSessionsRequest
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
				r := req.(*bastion.ListSessionsRequest)
				return c.ListSessions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]bastion.ListSessionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(bastion.ListSessionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionBastionClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "ListWorkRequestErrors", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     bastion.ListWorkRequestErrorsRequest
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
				r := req.(*bastion.ListWorkRequestErrorsRequest)
				return c.ListWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]bastion.ListWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(bastion.ListWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionBastionClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "ListWorkRequestLogs", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     bastion.ListWorkRequestLogsRequest
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
				r := req.(*bastion.ListWorkRequestLogsRequest)
				return c.ListWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]bastion.ListWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(bastion.ListWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionBastionClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "ListWorkRequests", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     bastion.ListWorkRequestsRequest
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
				r := req.(*bastion.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]bastion.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(bastion.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionBastionClientUpdateBastion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "UpdateBastion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBastion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "UpdateBastion", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "UpdateBastion")
	assert.NoError(t, err)

	type UpdateBastionRequestInfo struct {
		ContainerId string
		Request     bastion.UpdateBastionRequest
	}

	var requests []UpdateBastionRequestInfo
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
			response, err := c.UpdateBastion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionBastionClientUpdateSession(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastion", "UpdateSession")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSession is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastion", "Bastion", "UpdateSession", createBastionBastionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastion.BastionClient)

	body, err := testClient.getRequests("bastion", "UpdateSession")
	assert.NoError(t, err)

	type UpdateSessionRequestInfo struct {
		ContainerId string
		Request     bastion.UpdateSessionRequest
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
