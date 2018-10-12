package autotest

import (
    "github.com/oracle/oci-go-sdk/telemetry"
    "github.com/oracle/oci-go-sdk/common"

    "context"
    "encoding/json"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestTelemetryClientCreateAlarm(t *testing.T) {
    enabled, err := testClient.isApiEnabled("telemetry", "CreateAlarm")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateAlarm is not enabled by the testing service")
    }
    c, err := telemetry.NewTelemetryClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("telemetry", "CreateAlarm")
    assert.NoError(t, err)

    type CreateAlarmRequestInfo struct {
        ContainerId string
        Request telemetry.CreateAlarmRequest
    }

    var requests []CreateAlarmRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateAlarm(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestTelemetryClientDeleteAlarm(t *testing.T) {
    enabled, err := testClient.isApiEnabled("telemetry", "DeleteAlarm")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteAlarm is not enabled by the testing service")
    }
    c, err := telemetry.NewTelemetryClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("telemetry", "DeleteAlarm")
    assert.NoError(t, err)

    type DeleteAlarmRequestInfo struct {
        ContainerId string
        Request telemetry.DeleteAlarmRequest
    }

    var requests []DeleteAlarmRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteAlarm(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestTelemetryClientGetAlarm(t *testing.T) {
    enabled, err := testClient.isApiEnabled("telemetry", "GetAlarm")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetAlarm is not enabled by the testing service")
    }
    c, err := telemetry.NewTelemetryClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("telemetry", "GetAlarm")
    assert.NoError(t, err)

    type GetAlarmRequestInfo struct {
        ContainerId string
        Request telemetry.GetAlarmRequest
    }

    var requests []GetAlarmRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetAlarm(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestTelemetryClientGetAlarmHistory(t *testing.T) {
    enabled, err := testClient.isApiEnabled("telemetry", "GetAlarmHistory")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetAlarmHistory is not enabled by the testing service")
    }
    c, err := telemetry.NewTelemetryClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("telemetry", "GetAlarmHistory")
    assert.NoError(t, err)

    type GetAlarmHistoryRequestInfo struct {
        ContainerId string
        Request telemetry.GetAlarmHistoryRequest
    }

    var requests []GetAlarmHistoryRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*telemetry.GetAlarmHistoryRequest)
                return c.GetAlarmHistory(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]telemetry.GetAlarmHistoryResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(telemetry.GetAlarmHistoryResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestTelemetryClientListAlarms(t *testing.T) {
    enabled, err := testClient.isApiEnabled("telemetry", "ListAlarms")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListAlarms is not enabled by the testing service")
    }
    c, err := telemetry.NewTelemetryClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("telemetry", "ListAlarms")
    assert.NoError(t, err)

    type ListAlarmsRequestInfo struct {
        ContainerId string
        Request telemetry.ListAlarmsRequest
    }

    var requests []ListAlarmsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*telemetry.ListAlarmsRequest)
                return c.ListAlarms(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]telemetry.ListAlarmsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(telemetry.ListAlarmsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestTelemetryClientListMetrics(t *testing.T) {
    enabled, err := testClient.isApiEnabled("telemetry", "ListMetrics")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListMetrics is not enabled by the testing service")
    }
    c, err := telemetry.NewTelemetryClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("telemetry", "ListMetrics")
    assert.NoError(t, err)

    type ListMetricsRequestInfo struct {
        ContainerId string
        Request telemetry.ListMetricsRequest
    }

    var requests []ListMetricsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*telemetry.ListMetricsRequest)
                return c.ListMetrics(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]telemetry.ListMetricsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(telemetry.ListMetricsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestTelemetryClientRemoveAlarmSuppression(t *testing.T) {
    enabled, err := testClient.isApiEnabled("telemetry", "RemoveAlarmSuppression")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("RemoveAlarmSuppression is not enabled by the testing service")
    }
    c, err := telemetry.NewTelemetryClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("telemetry", "RemoveAlarmSuppression")
    assert.NoError(t, err)

    type RemoveAlarmSuppressionRequestInfo struct {
        ContainerId string
        Request telemetry.RemoveAlarmSuppressionRequest
    }

    var requests []RemoveAlarmSuppressionRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.RemoveAlarmSuppression(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestTelemetryClientSummarizeMetricsData(t *testing.T) {
    enabled, err := testClient.isApiEnabled("telemetry", "SummarizeMetricsData")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("SummarizeMetricsData is not enabled by the testing service")
    }
    c, err := telemetry.NewTelemetryClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("telemetry", "SummarizeMetricsData")
    assert.NoError(t, err)

    type SummarizeMetricsDataRequestInfo struct {
        ContainerId string
        Request telemetry.SummarizeMetricsDataRequest
    }

    var requests []SummarizeMetricsDataRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.SummarizeMetricsData(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestTelemetryClientUpdateAlarm(t *testing.T) {
    enabled, err := testClient.isApiEnabled("telemetry", "UpdateAlarm")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("UpdateAlarm is not enabled by the testing service")
    }
    c, err := telemetry.NewTelemetryClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("telemetry", "UpdateAlarm")
    assert.NoError(t, err)

    type UpdateAlarmRequestInfo struct {
        ContainerId string
        Request telemetry.UpdateAlarmRequest
    }

    var requests []UpdateAlarmRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.UpdateAlarm(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}
