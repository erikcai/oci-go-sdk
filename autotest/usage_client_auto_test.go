package autotest

import (
    "github.com/oracle/oci-go-sdk/usage"
    "github.com/oracle/oci-go-sdk/common"

    "context"
    "encoding/json"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

// IssueRoutingInfo email="" jiraProject="" opsJiraProject=""
func TestUsageClientGetSubscriptionInfo(t *testing.T) {
    enabled, err := testClient.isApiEnabled("usage", "GetSubscriptionInfo")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetSubscriptionInfo is not enabled by the testing service")
    }
    c, err := usage.NewUsageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("usage", "GetSubscriptionInfo")
    assert.NoError(t, err)

    type GetSubscriptionInfoRequestInfo struct {
        ContainerId string
        Request usage.GetSubscriptionInfoRequest
    }

    var requests []GetSubscriptionInfoRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetSubscriptionInfo(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="" jiraProject="" opsJiraProject=""
func TestUsageClientListUsageRecords(t *testing.T) {
    enabled, err := testClient.isApiEnabled("usage", "ListUsageRecords")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListUsageRecords is not enabled by the testing service")
    }
    c, err := usage.NewUsageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("usage", "ListUsageRecords")
    assert.NoError(t, err)

    type ListUsageRecordsRequestInfo struct {
        ContainerId string
        Request usage.ListUsageRecordsRequest
    }

    var requests []ListUsageRecordsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*usage.ListUsageRecordsRequest)
                return c.ListUsageRecords(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]usage.ListUsageRecordsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(usage.ListUsageRecordsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}
