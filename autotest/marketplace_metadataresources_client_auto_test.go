package autotest

import (
    "github.com/oracle/oci-go-sdk/marketplace"
    "github.com/oracle/oci-go-sdk/common"

    "context"
    "encoding/json"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func createMetadataResourcesClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {
    
    client, err := marketplace.NewMetadataResourcesClientWithConfigurationProvider(p)
    if testConfig.Endpoint != "" {
        client.Host = testConfig.Endpoint
    }else {
        client.SetRegion(testConfig.Region)
    }
   return client, err
}



// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestMetadataResourcesClientListCategories(t *testing.T) {
    enabled, err := testClient.isApiEnabled("marketplace", "ListCategories")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListCategories is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("marketplace", "MetadataResources", "ListCategories", createMetadataResourcesClientWithProvider)
    assert.NoError(t, err)
    c := cc.(marketplace.MetadataResourcesClient)

    body, err := testClient.getRequests("marketplace", "ListCategories")
    assert.NoError(t, err)

    type ListCategoriesRequestInfo struct {
        ContainerId string
        Request marketplace.ListCategoriesRequest
    }

    var requests []ListCategoriesRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*marketplace.ListCategoriesRequest)
                return c.ListCategories(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]marketplace.ListCategoriesResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(marketplace.ListCategoriesResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestMetadataResourcesClientListProductFilters(t *testing.T) {
    enabled, err := testClient.isApiEnabled("marketplace", "ListProductFilters")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListProductFilters is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("marketplace", "MetadataResources", "ListProductFilters", createMetadataResourcesClientWithProvider)
    assert.NoError(t, err)
    c := cc.(marketplace.MetadataResourcesClient)

    body, err := testClient.getRequests("marketplace", "ListProductFilters")
    assert.NoError(t, err)

    type ListProductFiltersRequestInfo struct {
        ContainerId string
        Request marketplace.ListProductFiltersRequest
    }

    var requests []ListProductFiltersRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*marketplace.ListProductFiltersRequest)
                return c.ListProductFilters(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]marketplace.ListProductFiltersResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(marketplace.ListProductFiltersResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}
