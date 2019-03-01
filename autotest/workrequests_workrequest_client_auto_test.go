package autotest

import (
    "github.com/oracle/oci-go-sdk/workrequests"
    "github.com/oracle/oci-go-sdk/common"

    "context"
    "encoding/json"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func createWorkRequestClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {
    
    client, err := workrequests.NewWorkRequestClientWithConfigurationProvider(p)
    if testConfig.Endpoint != "" {
        client.Host = testConfig.Endpoint
    }else {
        client.SetRegion(testConfig.Region)
    }
   return client, err
}



// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestWorkRequestClientGetWorkRequest(t *testing.T) {
    enabled, err := testClient.isApiEnabled("workrequests", "GetWorkRequest")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetWorkRequest is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("workrequests", "WorkRequest", "GetWorkRequest", createWorkRequestClientWithProvider)
    assert.NoError(t, err)
    c := cc.(workrequests.WorkRequestClient)

    body, err := testClient.getRequests("workrequests", "GetWorkRequest")
    assert.NoError(t, err)

    type GetWorkRequestRequestInfo struct {
        ContainerId string
        Request workrequests.GetWorkRequestRequest
    }

    var requests []GetWorkRequestRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetWorkRequest(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestWorkRequestClientListWorkRequestErrors(t *testing.T) {
    enabled, err := testClient.isApiEnabled("workrequests", "ListWorkRequestErrors")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequestErrors is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("workrequests", "WorkRequest", "ListWorkRequestErrors", createWorkRequestClientWithProvider)
    assert.NoError(t, err)
    c := cc.(workrequests.WorkRequestClient)

    body, err := testClient.getRequests("workrequests", "ListWorkRequestErrors")
    assert.NoError(t, err)

    type ListWorkRequestErrorsRequestInfo struct {
        ContainerId string
        Request workrequests.ListWorkRequestErrorsRequest
    }

    var requests []ListWorkRequestErrorsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*workrequests.ListWorkRequestErrorsRequest)
                return c.ListWorkRequestErrors(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]workrequests.ListWorkRequestErrorsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(workrequests.ListWorkRequestErrorsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestWorkRequestClientListWorkRequestLogs(t *testing.T) {
    enabled, err := testClient.isApiEnabled("workrequests", "ListWorkRequestLogs")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequestLogs is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("workrequests", "WorkRequest", "ListWorkRequestLogs", createWorkRequestClientWithProvider)
    assert.NoError(t, err)
    c := cc.(workrequests.WorkRequestClient)

    body, err := testClient.getRequests("workrequests", "ListWorkRequestLogs")
    assert.NoError(t, err)

    type ListWorkRequestLogsRequestInfo struct {
        ContainerId string
        Request workrequests.ListWorkRequestLogsRequest
    }

    var requests []ListWorkRequestLogsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*workrequests.ListWorkRequestLogsRequest)
                return c.ListWorkRequestLogs(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]workrequests.ListWorkRequestLogsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(workrequests.ListWorkRequestLogsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestWorkRequestClientListWorkRequests(t *testing.T) {
    enabled, err := testClient.isApiEnabled("workrequests", "ListWorkRequests")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequests is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("workrequests", "WorkRequest", "ListWorkRequests", createWorkRequestClientWithProvider)
    assert.NoError(t, err)
    c := cc.(workrequests.WorkRequestClient)

    body, err := testClient.getRequests("workrequests", "ListWorkRequests")
    assert.NoError(t, err)

    type ListWorkRequestsRequestInfo struct {
        ContainerId string
        Request workrequests.ListWorkRequestsRequest
    }

    var requests []ListWorkRequestsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*workrequests.ListWorkRequestsRequest)
                return c.ListWorkRequests(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]workrequests.ListWorkRequestsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(workrequests.ListWorkRequestsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}
