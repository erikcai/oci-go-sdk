package autotest

import (
    "github.com/oracle/oci-go-sdk/resourcesearch"
    "github.com/oracle/oci-go-sdk/common"

    "context"
    "encoding/json"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestResourceSearchClientGetResourceType(t *testing.T) {
    enabled, err := testClient.isApiEnabled("resourcesearch", "GetResourceType")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetResourceType is not enabled by the testing service")
    }
    c, err := resourcesearch.NewResourceSearchClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("resourcesearch", "GetResourceType")
    assert.NoError(t, err)

    type GetResourceTypeRequestInfo struct {
        ContainerId string
        Request resourcesearch.GetResourceTypeRequest
    }

    var requests []GetResourceTypeRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetResourceType(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestResourceSearchClientListResourceTypes(t *testing.T) {
    enabled, err := testClient.isApiEnabled("resourcesearch", "ListResourceTypes")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListResourceTypes is not enabled by the testing service")
    }
    c, err := resourcesearch.NewResourceSearchClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("resourcesearch", "ListResourceTypes")
    assert.NoError(t, err)

    type ListResourceTypesRequestInfo struct {
        ContainerId string
        Request resourcesearch.ListResourceTypesRequest
    }

    var requests []ListResourceTypesRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*resourcesearch.ListResourceTypesRequest)
                return c.ListResourceTypes(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]resourcesearch.ListResourceTypesResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(resourcesearch.ListResourceTypesResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestResourceSearchClientSearchResources(t *testing.T) {
    enabled, err := testClient.isApiEnabled("resourcesearch", "SearchResources")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("SearchResources is not enabled by the testing service")
    }
    c, err := resourcesearch.NewResourceSearchClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("resourcesearch", "SearchResources")
    assert.NoError(t, err)

    type SearchResourcesRequestInfo struct {
        ContainerId string
        Request resourcesearch.SearchResourcesRequest
    }

    var requests []SearchResourcesRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*resourcesearch.SearchResourcesRequest)
                return c.SearchResources(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]resourcesearch.SearchResourcesResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(resourcesearch.SearchResourcesResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}
