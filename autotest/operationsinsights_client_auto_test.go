package autotest

import (
	"github.com/oracle/oci-go-sdk/v30/common"
	"github.com/oracle/oci-go-sdk/v30/operationsinsights"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createOperationsinsightsOperationsInsightsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := operationsinsights.NewOperationsInsightsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="sqlLoader" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientIngestSqlBucket(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "IngestSqlBucket")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("IngestSqlBucket is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "IngestSqlBucket", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "IngestSqlBucket")
	assert.NoError(t, err)

	type IngestSqlBucketRequestInfo struct {
		ContainerId string
		Request     operationsinsights.IngestSqlBucketRequest
	}

	var requests []IngestSqlBucketRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.IngestSqlBucket(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="sqlLoader" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientIngestSqlPlanLines(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "IngestSqlPlanLines")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("IngestSqlPlanLines is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "IngestSqlPlanLines", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "IngestSqlPlanLines")
	assert.NoError(t, err)

	type IngestSqlPlanLinesRequestInfo struct {
		ContainerId string
		Request     operationsinsights.IngestSqlPlanLinesRequest
	}

	var requests []IngestSqlPlanLinesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.IngestSqlPlanLines(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="sqlLoader" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientIngestSqlText(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "IngestSqlText")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("IngestSqlText is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "IngestSqlText", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "IngestSqlText")
	assert.NoError(t, err)

	type IngestSqlTextRequestInfo struct {
		ContainerId string
		Request     operationsinsights.IngestSqlTextRequest
	}

	var requests []IngestSqlTextRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.IngestSqlText(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="resourcePlanning" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientListDatabaseInsights(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "ListDatabaseInsights")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDatabaseInsights is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "ListDatabaseInsights", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "ListDatabaseInsights")
	assert.NoError(t, err)

	type ListDatabaseInsightsRequestInfo struct {
		ContainerId string
		Request     operationsinsights.ListDatabaseInsightsRequest
	}

	var requests []ListDatabaseInsightsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.ListDatabaseInsightsRequest)
				return c.ListDatabaseInsights(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.ListDatabaseInsightsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.ListDatabaseInsightsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="sqlWarehouse" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientListSqlPlans(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "ListSqlPlans")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSqlPlans is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "ListSqlPlans", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "ListSqlPlans")
	assert.NoError(t, err)

	type ListSqlPlansRequestInfo struct {
		ContainerId string
		Request     operationsinsights.ListSqlPlansRequest
	}

	var requests []ListSqlPlansRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.ListSqlPlansRequest)
				return c.ListSqlPlans(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.ListSqlPlansResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.ListSqlPlansResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="sqlWarehouse" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientListSqlSearches(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "ListSqlSearches")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSqlSearches is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "ListSqlSearches", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "ListSqlSearches")
	assert.NoError(t, err)

	type ListSqlSearchesRequestInfo struct {
		ContainerId string
		Request     operationsinsights.ListSqlSearchesRequest
	}

	var requests []ListSqlSearchesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.ListSqlSearchesRequest)
				return c.ListSqlSearches(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.ListSqlSearchesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.ListSqlSearchesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="sqlWarehouse" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientListSqlTexts(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "ListSqlTexts")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSqlTexts is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "ListSqlTexts", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "ListSqlTexts")
	assert.NoError(t, err)

	type ListSqlTextsRequestInfo struct {
		ContainerId string
		Request     operationsinsights.ListSqlTextsRequest
	}

	var requests []ListSqlTextsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.ListSqlTextsRequest)
				return c.ListSqlTexts(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.ListSqlTextsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.ListSqlTextsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="resourcePlanning" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientSummarizeDatabaseInsightResourceCapacityTrend(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "SummarizeDatabaseInsightResourceCapacityTrend")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeDatabaseInsightResourceCapacityTrend is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "SummarizeDatabaseInsightResourceCapacityTrend", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "SummarizeDatabaseInsightResourceCapacityTrend")
	assert.NoError(t, err)

	type SummarizeDatabaseInsightResourceCapacityTrendRequestInfo struct {
		ContainerId string
		Request     operationsinsights.SummarizeDatabaseInsightResourceCapacityTrendRequest
	}

	var requests []SummarizeDatabaseInsightResourceCapacityTrendRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.SummarizeDatabaseInsightResourceCapacityTrendRequest)
				return c.SummarizeDatabaseInsightResourceCapacityTrend(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.SummarizeDatabaseInsightResourceCapacityTrendResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.SummarizeDatabaseInsightResourceCapacityTrendResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="resourcePlanning" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientSummarizeDatabaseInsightResourceForecastTrend(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "SummarizeDatabaseInsightResourceForecastTrend")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeDatabaseInsightResourceForecastTrend is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "SummarizeDatabaseInsightResourceForecastTrend", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "SummarizeDatabaseInsightResourceForecastTrend")
	assert.NoError(t, err)

	type SummarizeDatabaseInsightResourceForecastTrendRequestInfo struct {
		ContainerId string
		Request     operationsinsights.SummarizeDatabaseInsightResourceForecastTrendRequest
	}

	var requests []SummarizeDatabaseInsightResourceForecastTrendRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.SummarizeDatabaseInsightResourceForecastTrendRequest)
				return c.SummarizeDatabaseInsightResourceForecastTrend(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.SummarizeDatabaseInsightResourceForecastTrendResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.SummarizeDatabaseInsightResourceForecastTrendResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="resourcePlanning" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientSummarizeDatabaseInsightResourceStatistics(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "SummarizeDatabaseInsightResourceStatistics")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeDatabaseInsightResourceStatistics is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "SummarizeDatabaseInsightResourceStatistics", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "SummarizeDatabaseInsightResourceStatistics")
	assert.NoError(t, err)

	type SummarizeDatabaseInsightResourceStatisticsRequestInfo struct {
		ContainerId string
		Request     operationsinsights.SummarizeDatabaseInsightResourceStatisticsRequest
	}

	var requests []SummarizeDatabaseInsightResourceStatisticsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.SummarizeDatabaseInsightResourceStatisticsRequest)
				return c.SummarizeDatabaseInsightResourceStatistics(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.SummarizeDatabaseInsightResourceStatisticsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.SummarizeDatabaseInsightResourceStatisticsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="resourcePlanning" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientSummarizeDatabaseInsightResourceUsage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "SummarizeDatabaseInsightResourceUsage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeDatabaseInsightResourceUsage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "SummarizeDatabaseInsightResourceUsage", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "SummarizeDatabaseInsightResourceUsage")
	assert.NoError(t, err)

	type SummarizeDatabaseInsightResourceUsageRequestInfo struct {
		ContainerId string
		Request     operationsinsights.SummarizeDatabaseInsightResourceUsageRequest
	}

	var requests []SummarizeDatabaseInsightResourceUsageRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.SummarizeDatabaseInsightResourceUsageRequest)
				return c.SummarizeDatabaseInsightResourceUsage(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.SummarizeDatabaseInsightResourceUsageResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.SummarizeDatabaseInsightResourceUsageResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="resourcePlanning" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientSummarizeDatabaseInsightResourceUsageTrend(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "SummarizeDatabaseInsightResourceUsageTrend")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeDatabaseInsightResourceUsageTrend is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "SummarizeDatabaseInsightResourceUsageTrend", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "SummarizeDatabaseInsightResourceUsageTrend")
	assert.NoError(t, err)

	type SummarizeDatabaseInsightResourceUsageTrendRequestInfo struct {
		ContainerId string
		Request     operationsinsights.SummarizeDatabaseInsightResourceUsageTrendRequest
	}

	var requests []SummarizeDatabaseInsightResourceUsageTrendRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.SummarizeDatabaseInsightResourceUsageTrendRequest)
				return c.SummarizeDatabaseInsightResourceUsageTrend(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.SummarizeDatabaseInsightResourceUsageTrendResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.SummarizeDatabaseInsightResourceUsageTrendResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="resourcePlanning" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientSummarizeDatabaseInsightResourceUtilizationInsight(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "SummarizeDatabaseInsightResourceUtilizationInsight")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeDatabaseInsightResourceUtilizationInsight is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "SummarizeDatabaseInsightResourceUtilizationInsight", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "SummarizeDatabaseInsightResourceUtilizationInsight")
	assert.NoError(t, err)

	type SummarizeDatabaseInsightResourceUtilizationInsightRequestInfo struct {
		ContainerId string
		Request     operationsinsights.SummarizeDatabaseInsightResourceUtilizationInsightRequest
	}

	var requests []SummarizeDatabaseInsightResourceUtilizationInsightRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.SummarizeDatabaseInsightResourceUtilizationInsightRequest)
				return c.SummarizeDatabaseInsightResourceUtilizationInsight(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.SummarizeDatabaseInsightResourceUtilizationInsightResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.SummarizeDatabaseInsightResourceUtilizationInsightResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="sqlWarehouse" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientSummarizeSqlInsights(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "SummarizeSqlInsights")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeSqlInsights is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "SummarizeSqlInsights", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "SummarizeSqlInsights")
	assert.NoError(t, err)

	type SummarizeSqlInsightsRequestInfo struct {
		ContainerId string
		Request     operationsinsights.SummarizeSqlInsightsRequest
	}

	var requests []SummarizeSqlInsightsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.SummarizeSqlInsightsRequest)
				return c.SummarizeSqlInsights(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.SummarizeSqlInsightsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.SummarizeSqlInsightsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="sqlWarehouse" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientSummarizeSqlPlanInsights(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "SummarizeSqlPlanInsights")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeSqlPlanInsights is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "SummarizeSqlPlanInsights", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "SummarizeSqlPlanInsights")
	assert.NoError(t, err)

	type SummarizeSqlPlanInsightsRequestInfo struct {
		ContainerId string
		Request     operationsinsights.SummarizeSqlPlanInsightsRequest
	}

	var requests []SummarizeSqlPlanInsightsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.SummarizeSqlPlanInsightsRequest)
				return c.SummarizeSqlPlanInsights(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.SummarizeSqlPlanInsightsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.SummarizeSqlPlanInsightsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="sqlWarehouse" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientSummarizeSqlResponseTimeDistributions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "SummarizeSqlResponseTimeDistributions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeSqlResponseTimeDistributions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "SummarizeSqlResponseTimeDistributions", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "SummarizeSqlResponseTimeDistributions")
	assert.NoError(t, err)

	type SummarizeSqlResponseTimeDistributionsRequestInfo struct {
		ContainerId string
		Request     operationsinsights.SummarizeSqlResponseTimeDistributionsRequest
	}

	var requests []SummarizeSqlResponseTimeDistributionsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.SummarizeSqlResponseTimeDistributionsRequest)
				return c.SummarizeSqlResponseTimeDistributions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.SummarizeSqlResponseTimeDistributionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.SummarizeSqlResponseTimeDistributionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="sqlWarehouse" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientSummarizeSqlStatistics(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "SummarizeSqlStatistics")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeSqlStatistics is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "SummarizeSqlStatistics", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "SummarizeSqlStatistics")
	assert.NoError(t, err)

	type SummarizeSqlStatisticsRequestInfo struct {
		ContainerId string
		Request     operationsinsights.SummarizeSqlStatisticsRequest
	}

	var requests []SummarizeSqlStatisticsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.SummarizeSqlStatisticsRequest)
				return c.SummarizeSqlStatistics(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.SummarizeSqlStatisticsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.SummarizeSqlStatisticsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="sqlWarehouse" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientSummarizeSqlStatisticsTimeSeries(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "SummarizeSqlStatisticsTimeSeries")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeSqlStatisticsTimeSeries is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "SummarizeSqlStatisticsTimeSeries", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "SummarizeSqlStatisticsTimeSeries")
	assert.NoError(t, err)

	type SummarizeSqlStatisticsTimeSeriesRequestInfo struct {
		ContainerId string
		Request     operationsinsights.SummarizeSqlStatisticsTimeSeriesRequest
	}

	var requests []SummarizeSqlStatisticsTimeSeriesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.SummarizeSqlStatisticsTimeSeriesRequest)
				return c.SummarizeSqlStatisticsTimeSeries(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.SummarizeSqlStatisticsTimeSeriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.SummarizeSqlStatisticsTimeSeriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="sqlWarehouse" email="dbx_dev_ww_grp@oracle.com" jiraProject="DBX" opsJiraProject="DBXSD"
func TestOperationsinsightsOperationsInsightsClientSummarizeSqlStatisticsTimeSeriesByPlan(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operationsinsights", "SummarizeSqlStatisticsTimeSeriesByPlan")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeSqlStatisticsTimeSeriesByPlan is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operationsinsights", "OperationsInsights", "SummarizeSqlStatisticsTimeSeriesByPlan", createOperationsinsightsOperationsInsightsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operationsinsights.OperationsInsightsClient)

	body, err := testClient.getRequests("operationsinsights", "SummarizeSqlStatisticsTimeSeriesByPlan")
	assert.NoError(t, err)

	type SummarizeSqlStatisticsTimeSeriesByPlanRequestInfo struct {
		ContainerId string
		Request     operationsinsights.SummarizeSqlStatisticsTimeSeriesByPlanRequest
	}

	var requests []SummarizeSqlStatisticsTimeSeriesByPlanRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*operationsinsights.SummarizeSqlStatisticsTimeSeriesByPlanRequest)
				return c.SummarizeSqlStatisticsTimeSeriesByPlan(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operationsinsights.SummarizeSqlStatisticsTimeSeriesByPlanResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operationsinsights.SummarizeSqlStatisticsTimeSeriesByPlanResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
