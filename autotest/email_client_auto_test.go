package autotest

import (
    "github.com/oracle/oci-go-sdk/email"
    "github.com/oracle/oci-go-sdk/common"

    "context"
    "encoding/json"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

// IssueRoutingInfo email="" jiraProject="" opsJiraProject=""
func TestEmailClientCreateSender(t *testing.T) {
    enabled, err := testClient.isApiEnabled("email", "CreateSender")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateSender is not enabled by the testing service")
    }
    c, err := email.NewEmailClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("email", "CreateSender")
    assert.NoError(t, err)

    type CreateSenderRequestInfo struct {
        ContainerId string
        Request email.CreateSenderRequest
    }

    var requests []CreateSenderRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateSender(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="" jiraProject="" opsJiraProject=""
func TestEmailClientCreateSuppression(t *testing.T) {
    enabled, err := testClient.isApiEnabled("email", "CreateSuppression")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateSuppression is not enabled by the testing service")
    }
    c, err := email.NewEmailClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("email", "CreateSuppression")
    assert.NoError(t, err)

    type CreateSuppressionRequestInfo struct {
        ContainerId string
        Request email.CreateSuppressionRequest
    }

    var requests []CreateSuppressionRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateSuppression(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="" jiraProject="" opsJiraProject=""
func TestEmailClientDeleteSender(t *testing.T) {
    enabled, err := testClient.isApiEnabled("email", "DeleteSender")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteSender is not enabled by the testing service")
    }
    c, err := email.NewEmailClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("email", "DeleteSender")
    assert.NoError(t, err)

    type DeleteSenderRequestInfo struct {
        ContainerId string
        Request email.DeleteSenderRequest
    }

    var requests []DeleteSenderRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteSender(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="" jiraProject="" opsJiraProject=""
func TestEmailClientDeleteSuppression(t *testing.T) {
    enabled, err := testClient.isApiEnabled("email", "DeleteSuppression")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteSuppression is not enabled by the testing service")
    }
    c, err := email.NewEmailClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("email", "DeleteSuppression")
    assert.NoError(t, err)

    type DeleteSuppressionRequestInfo struct {
        ContainerId string
        Request email.DeleteSuppressionRequest
    }

    var requests []DeleteSuppressionRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteSuppression(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="" jiraProject="" opsJiraProject=""
func TestEmailClientGetSender(t *testing.T) {
    enabled, err := testClient.isApiEnabled("email", "GetSender")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetSender is not enabled by the testing service")
    }
    c, err := email.NewEmailClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("email", "GetSender")
    assert.NoError(t, err)

    type GetSenderRequestInfo struct {
        ContainerId string
        Request email.GetSenderRequest
    }

    var requests []GetSenderRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetSender(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="" jiraProject="" opsJiraProject=""
func TestEmailClientGetSuppression(t *testing.T) {
    enabled, err := testClient.isApiEnabled("email", "GetSuppression")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetSuppression is not enabled by the testing service")
    }
    c, err := email.NewEmailClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("email", "GetSuppression")
    assert.NoError(t, err)

    type GetSuppressionRequestInfo struct {
        ContainerId string
        Request email.GetSuppressionRequest
    }

    var requests []GetSuppressionRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetSuppression(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="" jiraProject="" opsJiraProject=""
func TestEmailClientListSenders(t *testing.T) {
    enabled, err := testClient.isApiEnabled("email", "ListSenders")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListSenders is not enabled by the testing service")
    }
    c, err := email.NewEmailClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("email", "ListSenders")
    assert.NoError(t, err)

    type ListSendersRequestInfo struct {
        ContainerId string
        Request email.ListSendersRequest
    }

    var requests []ListSendersRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*email.ListSendersRequest)
                return c.ListSenders(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]email.ListSendersResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(email.ListSendersResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="" jiraProject="" opsJiraProject=""
func TestEmailClientListSuppressions(t *testing.T) {
    enabled, err := testClient.isApiEnabled("email", "ListSuppressions")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListSuppressions is not enabled by the testing service")
    }
    c, err := email.NewEmailClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("email", "ListSuppressions")
    assert.NoError(t, err)

    type ListSuppressionsRequestInfo struct {
        ContainerId string
        Request email.ListSuppressionsRequest
    }

    var requests []ListSuppressionsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*email.ListSuppressionsRequest)
                return c.ListSuppressions(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]email.ListSuppressionsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(email.ListSuppressionsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}
