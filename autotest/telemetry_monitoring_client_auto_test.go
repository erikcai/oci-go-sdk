package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/telemetry"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientCreateAlarm(t *testing.T) {
	enabled, err := testClient.isApiEnabled("telemetry", "CreateAlarm")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAlarm is not enabled by the testing service")
	}
	c, err := telemetry.NewMonitoringClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("telemetry", "CreateAlarm")
	assert.NoError(t, err)

	type CreateAlarmRequestInfo struct {
		ContainerId string
		Request     telemetry.CreateAlarmRequest
	}

	var requests []CreateAlarmRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateAlarm(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientDeleteAlarm(t *testing.T) {
	enabled, err := testClient.isApiEnabled("telemetry", "DeleteAlarm")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAlarm is not enabled by the testing service")
	}
	c, err := telemetry.NewMonitoringClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("telemetry", "DeleteAlarm")
	assert.NoError(t, err)

	type DeleteAlarmRequestInfo struct {
		ContainerId string
		Request     telemetry.DeleteAlarmRequest
	}

	var requests []DeleteAlarmRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteAlarm(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientGetAlarm(t *testing.T) {
	enabled, err := testClient.isApiEnabled("telemetry", "GetAlarm")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAlarm is not enabled by the testing service")
	}
	c, err := telemetry.NewMonitoringClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("telemetry", "GetAlarm")
	assert.NoError(t, err)

	type GetAlarmRequestInfo struct {
		ContainerId string
		Request     telemetry.GetAlarmRequest
	}

	var requests []GetAlarmRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAlarm(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientGetAlarmHistory(t *testing.T) {
	enabled, err := testClient.isApiEnabled("telemetry", "GetAlarmHistory")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAlarmHistory is not enabled by the testing service")
	}
	c, err := telemetry.NewMonitoringClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("telemetry", "GetAlarmHistory")
	assert.NoError(t, err)

	type GetAlarmHistoryRequestInfo struct {
		ContainerId string
		Request     telemetry.GetAlarmHistoryRequest
	}

	var requests []GetAlarmHistoryRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
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

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientListAlarms(t *testing.T) {
	enabled, err := testClient.isApiEnabled("telemetry", "ListAlarms")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAlarms is not enabled by the testing service")
	}
	c, err := telemetry.NewMonitoringClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("telemetry", "ListAlarms")
	assert.NoError(t, err)

	type ListAlarmsRequestInfo struct {
		ContainerId string
		Request     telemetry.ListAlarmsRequest
	}

	var requests []ListAlarmsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
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

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientListAlarmsStatus(t *testing.T) {
	enabled, err := testClient.isApiEnabled("telemetry", "ListAlarmsStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAlarmsStatus is not enabled by the testing service")
	}
	c, err := telemetry.NewMonitoringClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("telemetry", "ListAlarmsStatus")
	assert.NoError(t, err)

	type ListAlarmsStatusRequestInfo struct {
		ContainerId string
		Request     telemetry.ListAlarmsStatusRequest
	}

	var requests []ListAlarmsStatusRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*telemetry.ListAlarmsStatusRequest)
				return c.ListAlarmsStatus(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]telemetry.ListAlarmsStatusResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(telemetry.ListAlarmsStatusResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientListMetrics(t *testing.T) {
	enabled, err := testClient.isApiEnabled("telemetry", "ListMetrics")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListMetrics is not enabled by the testing service")
	}
	c, err := telemetry.NewMonitoringClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("telemetry", "ListMetrics")
	assert.NoError(t, err)

	type ListMetricsRequestInfo struct {
		ContainerId string
		Request     telemetry.ListMetricsRequest
	}

	var requests []ListMetricsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
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

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientPostMetricData(t *testing.T) {
	enabled, err := testClient.isApiEnabled("telemetry", "PostMetricData")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PostMetricData is not enabled by the testing service")
	}
	c, err := telemetry.NewMonitoringClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("telemetry", "PostMetricData")
	assert.NoError(t, err)

	type PostMetricDataRequestInfo struct {
		ContainerId string
		Request     telemetry.PostMetricDataRequest
	}

	var requests []PostMetricDataRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.PostMetricData(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientRemoveAlarmSuppression(t *testing.T) {
	enabled, err := testClient.isApiEnabled("telemetry", "RemoveAlarmSuppression")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RemoveAlarmSuppression is not enabled by the testing service")
	}
	c, err := telemetry.NewMonitoringClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("telemetry", "RemoveAlarmSuppression")
	assert.NoError(t, err)

	type RemoveAlarmSuppressionRequestInfo struct {
		ContainerId string
		Request     telemetry.RemoveAlarmSuppressionRequest
	}

	var requests []RemoveAlarmSuppressionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.RemoveAlarmSuppression(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientSummarizeMetricsData(t *testing.T) {
	enabled, err := testClient.isApiEnabled("telemetry", "SummarizeMetricsData")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeMetricsData is not enabled by the testing service")
	}
	c, err := telemetry.NewMonitoringClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("telemetry", "SummarizeMetricsData")
	assert.NoError(t, err)

	type SummarizeMetricsDataRequestInfo struct {
		ContainerId string
		Request     telemetry.SummarizeMetricsDataRequest
	}

	var requests []SummarizeMetricsDataRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.SummarizeMetricsData(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientUpdateAlarm(t *testing.T) {
	enabled, err := testClient.isApiEnabled("telemetry", "UpdateAlarm")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAlarm is not enabled by the testing service")
	}
	c, err := telemetry.NewMonitoringClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("telemetry", "UpdateAlarm")
	assert.NoError(t, err)

	type UpdateAlarmRequestInfo struct {
		ContainerId string
		Request     telemetry.UpdateAlarmRequest
	}

	var requests []UpdateAlarmRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateAlarm(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
