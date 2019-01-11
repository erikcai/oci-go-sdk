package autotest

import (
    "github.com/oracle/oci-go-sdk/cache"
    "github.com/oracle/oci-go-sdk/common"

    "context"
    "encoding/json"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func createOracacheClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {
    
    client, err := cache.NewOracacheClientWithConfigurationProvider(p)
    if testConfig.Endpoint != "" {
        client.Host = testConfig.Endpoint
    }else {
        client.SetRegion(testConfig.Region)
    }
   return client, err
}



// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestOracacheClientCancelWorkRequest(t *testing.T) {
    enabled, err := testClient.isApiEnabled("cache", "CancelWorkRequest")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CancelWorkRequest is not enabled by the testing service")
    }
    c, err := cache.NewOracacheClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("cache", "CancelWorkRequest")
    assert.NoError(t, err)

    type CancelWorkRequestRequestInfo struct {
        ContainerId string
        Request cache.CancelWorkRequestRequest
    }

    var requests []CancelWorkRequestRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CancelWorkRequest(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestOracacheClientCreateReplicatedCache(t *testing.T) {
    enabled, err := testClient.isApiEnabled("cache", "CreateReplicatedCache")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateReplicatedCache is not enabled by the testing service")
    }
    c, err := cache.NewOracacheClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("cache", "CreateReplicatedCache")
    assert.NoError(t, err)

    type CreateReplicatedCacheRequestInfo struct {
        ContainerId string
        Request cache.CreateReplicatedCacheRequest
    }

    var requests []CreateReplicatedCacheRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateReplicatedCache(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestOracacheClientDeleteReplicatedCache(t *testing.T) {
    enabled, err := testClient.isApiEnabled("cache", "DeleteReplicatedCache")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteReplicatedCache is not enabled by the testing service")
    }
    c, err := cache.NewOracacheClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("cache", "DeleteReplicatedCache")
    assert.NoError(t, err)

    type DeleteReplicatedCacheRequestInfo struct {
        ContainerId string
        Request cache.DeleteReplicatedCacheRequest
    }

    var requests []DeleteReplicatedCacheRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteReplicatedCache(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestOracacheClientGetReplicatedCache(t *testing.T) {
    enabled, err := testClient.isApiEnabled("cache", "GetReplicatedCache")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetReplicatedCache is not enabled by the testing service")
    }
    c, err := cache.NewOracacheClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("cache", "GetReplicatedCache")
    assert.NoError(t, err)

    type GetReplicatedCacheRequestInfo struct {
        ContainerId string
        Request cache.GetReplicatedCacheRequest
    }

    var requests []GetReplicatedCacheRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetReplicatedCache(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestOracacheClientGetWorkRequest(t *testing.T) {
    enabled, err := testClient.isApiEnabled("cache", "GetWorkRequest")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetWorkRequest is not enabled by the testing service")
    }
    c, err := cache.NewOracacheClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("cache", "GetWorkRequest")
    assert.NoError(t, err)

    type GetWorkRequestRequestInfo struct {
        ContainerId string
        Request cache.GetWorkRequestRequest
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
func TestOracacheClientListCachingConsumption(t *testing.T) {
    enabled, err := testClient.isApiEnabled("cache", "ListCachingConsumption")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListCachingConsumption is not enabled by the testing service")
    }
    c, err := cache.NewOracacheClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("cache", "ListCachingConsumption")
    assert.NoError(t, err)

    type ListCachingConsumptionRequestInfo struct {
        ContainerId string
        Request cache.ListCachingConsumptionRequest
    }

    var requests []ListCachingConsumptionRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.ListCachingConsumption(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestOracacheClientListRedisShapes(t *testing.T) {
    enabled, err := testClient.isApiEnabled("cache", "ListRedisShapes")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListRedisShapes is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("cache", "Oracache", "ListRedisShapes", createOracacheClientWithProvider)
    assert.NoError(t, err)
    c := cc.(cache.OracacheClient)

    body, err := testClient.getRequests("cache", "ListRedisShapes")
    assert.NoError(t, err)

    type ListRedisShapesRequestInfo struct {
        ContainerId string
        Request cache.ListRedisShapesRequest
    }

    var requests []ListRedisShapesRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*cache.ListRedisShapesRequest)
                return c.ListRedisShapes(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]cache.ListRedisShapesResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(cache.ListRedisShapesResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestOracacheClientListReplicatedCaches(t *testing.T) {
    enabled, err := testClient.isApiEnabled("cache", "ListReplicatedCaches")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListReplicatedCaches is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("cache", "Oracache", "ListReplicatedCaches", createOracacheClientWithProvider)
    assert.NoError(t, err)
    c := cc.(cache.OracacheClient)

    body, err := testClient.getRequests("cache", "ListReplicatedCaches")
    assert.NoError(t, err)

    type ListReplicatedCachesRequestInfo struct {
        ContainerId string
        Request cache.ListReplicatedCachesRequest
    }

    var requests []ListReplicatedCachesRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*cache.ListReplicatedCachesRequest)
                return c.ListReplicatedCaches(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]cache.ListReplicatedCachesResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(cache.ListReplicatedCachesResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestOracacheClientListVersions(t *testing.T) {
    enabled, err := testClient.isApiEnabled("cache", "ListVersions")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListVersions is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("cache", "Oracache", "ListVersions", createOracacheClientWithProvider)
    assert.NoError(t, err)
    c := cc.(cache.OracacheClient)

    body, err := testClient.getRequests("cache", "ListVersions")
    assert.NoError(t, err)

    type ListVersionsRequestInfo struct {
        ContainerId string
        Request cache.ListVersionsRequest
    }

    var requests []ListVersionsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*cache.ListVersionsRequest)
                return c.ListVersions(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]cache.ListVersionsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(cache.ListVersionsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestOracacheClientListWorkRequestErrors(t *testing.T) {
    enabled, err := testClient.isApiEnabled("cache", "ListWorkRequestErrors")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequestErrors is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("cache", "Oracache", "ListWorkRequestErrors", createOracacheClientWithProvider)
    assert.NoError(t, err)
    c := cc.(cache.OracacheClient)

    body, err := testClient.getRequests("cache", "ListWorkRequestErrors")
    assert.NoError(t, err)

    type ListWorkRequestErrorsRequestInfo struct {
        ContainerId string
        Request cache.ListWorkRequestErrorsRequest
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
                r := req.(*cache.ListWorkRequestErrorsRequest)
                return c.ListWorkRequestErrors(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]cache.ListWorkRequestErrorsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(cache.ListWorkRequestErrorsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestOracacheClientListWorkRequestLogs(t *testing.T) {
    enabled, err := testClient.isApiEnabled("cache", "ListWorkRequestLogs")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequestLogs is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("cache", "Oracache", "ListWorkRequestLogs", createOracacheClientWithProvider)
    assert.NoError(t, err)
    c := cc.(cache.OracacheClient)

    body, err := testClient.getRequests("cache", "ListWorkRequestLogs")
    assert.NoError(t, err)

    type ListWorkRequestLogsRequestInfo struct {
        ContainerId string
        Request cache.ListWorkRequestLogsRequest
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
                r := req.(*cache.ListWorkRequestLogsRequest)
                return c.ListWorkRequestLogs(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]cache.ListWorkRequestLogsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(cache.ListWorkRequestLogsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestOracacheClientListWorkRequests(t *testing.T) {
    enabled, err := testClient.isApiEnabled("cache", "ListWorkRequests")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequests is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("cache", "Oracache", "ListWorkRequests", createOracacheClientWithProvider)
    assert.NoError(t, err)
    c := cc.(cache.OracacheClient)

    body, err := testClient.getRequests("cache", "ListWorkRequests")
    assert.NoError(t, err)

    type ListWorkRequestsRequestInfo struct {
        ContainerId string
        Request cache.ListWorkRequestsRequest
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
                r := req.(*cache.ListWorkRequestsRequest)
                return c.ListWorkRequests(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]cache.ListWorkRequestsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(cache.ListWorkRequestsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}
