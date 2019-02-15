package autotest

import (
    "github.com/oracle/oci-go-sdk/datasecurity"
    "github.com/oracle/oci-go-sdk/common"

    "context"
    "encoding/json"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func createDataSecurityClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {
    
    client, err := datasecurity.NewDataSecurityClientWithConfigurationProvider(p)
    if testConfig.Endpoint != "" {
        client.Host = testConfig.Endpoint
    }else {
        client.SetRegion(testConfig.Region)
    }
   return client, err
}



// IssueRoutingInfo tag="default" email="datasecurity_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSecurityClientChangeDataSecurityInstanceCompartment(t *testing.T) {
    enabled, err := testClient.isApiEnabled("datasecurity", "ChangeDataSecurityInstanceCompartment")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ChangeDataSecurityInstanceCompartment is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("datasecurity", "DataSecurity", "ChangeDataSecurityInstanceCompartment", createDataSecurityClientWithProvider)
    assert.NoError(t, err)
    c := cc.(datasecurity.DataSecurityClient)

    body, err := testClient.getRequests("datasecurity", "ChangeDataSecurityInstanceCompartment")
    assert.NoError(t, err)

    type ChangeDataSecurityInstanceCompartmentRequestInfo struct {
        ContainerId string
        Request datasecurity.ChangeDataSecurityInstanceCompartmentRequest
    }

    var requests []ChangeDataSecurityInstanceCompartmentRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.ChangeDataSecurityInstanceCompartment(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="datasecurity_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSecurityClientCreateDataSecurityInstance(t *testing.T) {
    enabled, err := testClient.isApiEnabled("datasecurity", "CreateDataSecurityInstance")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateDataSecurityInstance is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("datasecurity", "DataSecurity", "CreateDataSecurityInstance", createDataSecurityClientWithProvider)
    assert.NoError(t, err)
    c := cc.(datasecurity.DataSecurityClient)

    body, err := testClient.getRequests("datasecurity", "CreateDataSecurityInstance")
    assert.NoError(t, err)

    type CreateDataSecurityInstanceRequestInfo struct {
        ContainerId string
        Request datasecurity.CreateDataSecurityInstanceRequest
    }

    var requests []CreateDataSecurityInstanceRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateDataSecurityInstance(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="datasecurity_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSecurityClientDeleteDataSecurityInstance(t *testing.T) {
    enabled, err := testClient.isApiEnabled("datasecurity", "DeleteDataSecurityInstance")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteDataSecurityInstance is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("datasecurity", "DataSecurity", "DeleteDataSecurityInstance", createDataSecurityClientWithProvider)
    assert.NoError(t, err)
    c := cc.(datasecurity.DataSecurityClient)

    body, err := testClient.getRequests("datasecurity", "DeleteDataSecurityInstance")
    assert.NoError(t, err)

    type DeleteDataSecurityInstanceRequestInfo struct {
        ContainerId string
        Request datasecurity.DeleteDataSecurityInstanceRequest
    }

    var requests []DeleteDataSecurityInstanceRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteDataSecurityInstance(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="datasecurity_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSecurityClientGetDataSecurityInstance(t *testing.T) {
    enabled, err := testClient.isApiEnabled("datasecurity", "GetDataSecurityInstance")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetDataSecurityInstance is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("datasecurity", "DataSecurity", "GetDataSecurityInstance", createDataSecurityClientWithProvider)
    assert.NoError(t, err)
    c := cc.(datasecurity.DataSecurityClient)

    body, err := testClient.getRequests("datasecurity", "GetDataSecurityInstance")
    assert.NoError(t, err)

    type GetDataSecurityInstanceRequestInfo struct {
        ContainerId string
        Request datasecurity.GetDataSecurityInstanceRequest
    }

    var requests []GetDataSecurityInstanceRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetDataSecurityInstance(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="datasecurity_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSecurityClientGetWorkRequest(t *testing.T) {
    enabled, err := testClient.isApiEnabled("datasecurity", "GetWorkRequest")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetWorkRequest is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("datasecurity", "DataSecurity", "GetWorkRequest", createDataSecurityClientWithProvider)
    assert.NoError(t, err)
    c := cc.(datasecurity.DataSecurityClient)

    body, err := testClient.getRequests("datasecurity", "GetWorkRequest")
    assert.NoError(t, err)

    type GetWorkRequestRequestInfo struct {
        ContainerId string
        Request datasecurity.GetWorkRequestRequest
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

// IssueRoutingInfo tag="default" email="datasecurity_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSecurityClientListDataSecurityInstances(t *testing.T) {
    enabled, err := testClient.isApiEnabled("datasecurity", "ListDataSecurityInstances")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListDataSecurityInstances is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("datasecurity", "DataSecurity", "ListDataSecurityInstances", createDataSecurityClientWithProvider)
    assert.NoError(t, err)
    c := cc.(datasecurity.DataSecurityClient)

    body, err := testClient.getRequests("datasecurity", "ListDataSecurityInstances")
    assert.NoError(t, err)

    type ListDataSecurityInstancesRequestInfo struct {
        ContainerId string
        Request datasecurity.ListDataSecurityInstancesRequest
    }

    var requests []ListDataSecurityInstancesRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*datasecurity.ListDataSecurityInstancesRequest)
                return c.ListDataSecurityInstances(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]datasecurity.ListDataSecurityInstancesResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(datasecurity.ListDataSecurityInstancesResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="datasecurity_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSecurityClientListWorkRequestErrors(t *testing.T) {
    enabled, err := testClient.isApiEnabled("datasecurity", "ListWorkRequestErrors")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequestErrors is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("datasecurity", "DataSecurity", "ListWorkRequestErrors", createDataSecurityClientWithProvider)
    assert.NoError(t, err)
    c := cc.(datasecurity.DataSecurityClient)

    body, err := testClient.getRequests("datasecurity", "ListWorkRequestErrors")
    assert.NoError(t, err)

    type ListWorkRequestErrorsRequestInfo struct {
        ContainerId string
        Request datasecurity.ListWorkRequestErrorsRequest
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
                r := req.(*datasecurity.ListWorkRequestErrorsRequest)
                return c.ListWorkRequestErrors(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]datasecurity.ListWorkRequestErrorsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(datasecurity.ListWorkRequestErrorsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="datasecurity_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSecurityClientListWorkRequestLogs(t *testing.T) {
    enabled, err := testClient.isApiEnabled("datasecurity", "ListWorkRequestLogs")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequestLogs is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("datasecurity", "DataSecurity", "ListWorkRequestLogs", createDataSecurityClientWithProvider)
    assert.NoError(t, err)
    c := cc.(datasecurity.DataSecurityClient)

    body, err := testClient.getRequests("datasecurity", "ListWorkRequestLogs")
    assert.NoError(t, err)

    type ListWorkRequestLogsRequestInfo struct {
        ContainerId string
        Request datasecurity.ListWorkRequestLogsRequest
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
                r := req.(*datasecurity.ListWorkRequestLogsRequest)
                return c.ListWorkRequestLogs(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]datasecurity.ListWorkRequestLogsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(datasecurity.ListWorkRequestLogsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="datasecurity_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSecurityClientListWorkRequests(t *testing.T) {
    enabled, err := testClient.isApiEnabled("datasecurity", "ListWorkRequests")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequests is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("datasecurity", "DataSecurity", "ListWorkRequests", createDataSecurityClientWithProvider)
    assert.NoError(t, err)
    c := cc.(datasecurity.DataSecurityClient)

    body, err := testClient.getRequests("datasecurity", "ListWorkRequests")
    assert.NoError(t, err)

    type ListWorkRequestsRequestInfo struct {
        ContainerId string
        Request datasecurity.ListWorkRequestsRequest
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
                r := req.(*datasecurity.ListWorkRequestsRequest)
                return c.ListWorkRequests(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]datasecurity.ListWorkRequestsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(datasecurity.ListWorkRequestsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="datasecurity_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSecurityClientUpdateDataSecurityInstance(t *testing.T) {
    enabled, err := testClient.isApiEnabled("datasecurity", "UpdateDataSecurityInstance")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("UpdateDataSecurityInstance is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("datasecurity", "DataSecurity", "UpdateDataSecurityInstance", createDataSecurityClientWithProvider)
    assert.NoError(t, err)
    c := cc.(datasecurity.DataSecurityClient)

    body, err := testClient.getRequests("datasecurity", "UpdateDataSecurityInstance")
    assert.NoError(t, err)

    type UpdateDataSecurityInstanceRequestInfo struct {
        ContainerId string
        Request datasecurity.UpdateDataSecurityInstanceRequest
    }

    var requests []UpdateDataSecurityInstanceRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.UpdateDataSecurityInstance(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}
