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

func createApplicationResourcesClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {
    
    client, err := marketplace.NewApplicationResourcesClientWithConfigurationProvider(p)
    if testConfig.Endpoint != "" {
        client.Host = testConfig.Endpoint
    }else {
        client.SetRegion(testConfig.Region)
    }
   return client, err
}



// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestApplicationResourcesClientGetApplication(t *testing.T) {
    enabled, err := testClient.isApiEnabled("marketplace", "GetApplication")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetApplication is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("marketplace", "ApplicationResources", "GetApplication", createApplicationResourcesClientWithProvider)
    assert.NoError(t, err)
    c := cc.(marketplace.ApplicationResourcesClient)

    body, err := testClient.getRequests("marketplace", "GetApplication")
    assert.NoError(t, err)

    type GetApplicationRequestInfo struct {
        ContainerId string
        Request marketplace.GetApplicationRequest
    }

    var requests []GetApplicationRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetApplication(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestApplicationResourcesClientListApplications(t *testing.T) {
    enabled, err := testClient.isApiEnabled("marketplace", "ListApplications")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListApplications is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("marketplace", "ApplicationResources", "ListApplications", createApplicationResourcesClientWithProvider)
    assert.NoError(t, err)
    c := cc.(marketplace.ApplicationResourcesClient)

    body, err := testClient.getRequests("marketplace", "ListApplications")
    assert.NoError(t, err)

    type ListApplicationsRequestInfo struct {
        ContainerId string
        Request marketplace.ListApplicationsRequest
    }

    var requests []ListApplicationsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*marketplace.ListApplicationsRequest)
                return c.ListApplications(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]marketplace.ListApplicationsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(marketplace.ListApplicationsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}
