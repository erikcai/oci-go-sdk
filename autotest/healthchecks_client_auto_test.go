package autotest

import (
    "github.com/oracle/oci-go-sdk/healthchecks"
    "github.com/oracle/oci-go-sdk/common"

    "context"
    "encoding/json"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientCreateHttpMonitor(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "CreateHttpMonitor")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateHttpMonitor is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "CreateHttpMonitor")
    assert.NoError(t, err)

    type CreateHttpMonitorRequestInfo struct {
        ContainerId string
        Request healthchecks.CreateHttpMonitorRequest
    }

    var requests []CreateHttpMonitorRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateHttpMonitor(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientCreateOnDemandHttpProbe(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "CreateOnDemandHttpProbe")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateOnDemandHttpProbe is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "CreateOnDemandHttpProbe")
    assert.NoError(t, err)

    type CreateOnDemandHttpProbeRequestInfo struct {
        ContainerId string
        Request healthchecks.CreateOnDemandHttpProbeRequest
    }

    var requests []CreateOnDemandHttpProbeRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateOnDemandHttpProbe(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientCreateOnDemandPingProbe(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "CreateOnDemandPingProbe")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateOnDemandPingProbe is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "CreateOnDemandPingProbe")
    assert.NoError(t, err)

    type CreateOnDemandPingProbeRequestInfo struct {
        ContainerId string
        Request healthchecks.CreateOnDemandPingProbeRequest
    }

    var requests []CreateOnDemandPingProbeRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateOnDemandPingProbe(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientCreatePingMonitor(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "CreatePingMonitor")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreatePingMonitor is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "CreatePingMonitor")
    assert.NoError(t, err)

    type CreatePingMonitorRequestInfo struct {
        ContainerId string
        Request healthchecks.CreatePingMonitorRequest
    }

    var requests []CreatePingMonitorRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreatePingMonitor(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientDeleteHttpMonitor(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "DeleteHttpMonitor")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteHttpMonitor is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "DeleteHttpMonitor")
    assert.NoError(t, err)

    type DeleteHttpMonitorRequestInfo struct {
        ContainerId string
        Request healthchecks.DeleteHttpMonitorRequest
    }

    var requests []DeleteHttpMonitorRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteHttpMonitor(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientDeletePingMonitor(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "DeletePingMonitor")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeletePingMonitor is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "DeletePingMonitor")
    assert.NoError(t, err)

    type DeletePingMonitorRequestInfo struct {
        ContainerId string
        Request healthchecks.DeletePingMonitorRequest
    }

    var requests []DeletePingMonitorRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeletePingMonitor(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientGetHttpMonitor(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "GetHttpMonitor")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetHttpMonitor is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "GetHttpMonitor")
    assert.NoError(t, err)

    type GetHttpMonitorRequestInfo struct {
        ContainerId string
        Request healthchecks.GetHttpMonitorRequest
    }

    var requests []GetHttpMonitorRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetHttpMonitor(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientGetPingMonitor(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "GetPingMonitor")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetPingMonitor is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "GetPingMonitor")
    assert.NoError(t, err)

    type GetPingMonitorRequestInfo struct {
        ContainerId string
        Request healthchecks.GetPingMonitorRequest
    }

    var requests []GetPingMonitorRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetPingMonitor(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientListHealthChecksVantagePoints(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "ListHealthChecksVantagePoints")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListHealthChecksVantagePoints is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "ListHealthChecksVantagePoints")
    assert.NoError(t, err)

    type ListHealthChecksVantagePointsRequestInfo struct {
        ContainerId string
        Request healthchecks.ListHealthChecksVantagePointsRequest
    }

    var requests []ListHealthChecksVantagePointsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*healthchecks.ListHealthChecksVantagePointsRequest)
                return c.ListHealthChecksVantagePoints(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]healthchecks.ListHealthChecksVantagePointsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(healthchecks.ListHealthChecksVantagePointsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientListHttpMonitors(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "ListHttpMonitors")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListHttpMonitors is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "ListHttpMonitors")
    assert.NoError(t, err)

    type ListHttpMonitorsRequestInfo struct {
        ContainerId string
        Request healthchecks.ListHttpMonitorsRequest
    }

    var requests []ListHttpMonitorsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*healthchecks.ListHttpMonitorsRequest)
                return c.ListHttpMonitors(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]healthchecks.ListHttpMonitorsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(healthchecks.ListHttpMonitorsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientListHttpProbeResults(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "ListHttpProbeResults")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListHttpProbeResults is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "ListHttpProbeResults")
    assert.NoError(t, err)

    type ListHttpProbeResultsRequestInfo struct {
        ContainerId string
        Request healthchecks.ListHttpProbeResultsRequest
    }

    var requests []ListHttpProbeResultsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*healthchecks.ListHttpProbeResultsRequest)
                return c.ListHttpProbeResults(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]healthchecks.ListHttpProbeResultsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(healthchecks.ListHttpProbeResultsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientListPingMonitors(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "ListPingMonitors")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListPingMonitors is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "ListPingMonitors")
    assert.NoError(t, err)

    type ListPingMonitorsRequestInfo struct {
        ContainerId string
        Request healthchecks.ListPingMonitorsRequest
    }

    var requests []ListPingMonitorsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*healthchecks.ListPingMonitorsRequest)
                return c.ListPingMonitors(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]healthchecks.ListPingMonitorsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(healthchecks.ListPingMonitorsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientListPingProbeResults(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "ListPingProbeResults")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListPingProbeResults is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "ListPingProbeResults")
    assert.NoError(t, err)

    type ListPingProbeResultsRequestInfo struct {
        ContainerId string
        Request healthchecks.ListPingProbeResultsRequest
    }

    var requests []ListPingProbeResultsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*healthchecks.ListPingProbeResultsRequest)
                return c.ListPingProbeResults(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]healthchecks.ListPingProbeResultsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(healthchecks.ListPingProbeResultsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientUpdateHttpMonitor(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "UpdateHttpMonitor")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("UpdateHttpMonitor is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "UpdateHttpMonitor")
    assert.NoError(t, err)

    type UpdateHttpMonitorRequestInfo struct {
        ContainerId string
        Request healthchecks.UpdateHttpMonitorRequest
    }

    var requests []UpdateHttpMonitorRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.UpdateHttpMonitor(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="groan-chomskies_us_grp@oracle.com" jiraProject="OHC" opsJiraProject="HC"
func TestHealthChecksClientUpdatePingMonitor(t *testing.T) {
    enabled, err := testClient.isApiEnabled("healthchecks", "UpdatePingMonitor")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("UpdatePingMonitor is not enabled by the testing service")
    }
    c, err := healthchecks.NewHealthChecksClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("healthchecks", "UpdatePingMonitor")
    assert.NoError(t, err)

    type UpdatePingMonitorRequestInfo struct {
        ContainerId string
        Request healthchecks.UpdatePingMonitorRequest
    }

    var requests []UpdatePingMonitorRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.UpdatePingMonitor(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}
