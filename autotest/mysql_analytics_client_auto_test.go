package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/mysql"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createMysqlAnalyticsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := mysql.NewAnalyticsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlAnalyticsClientCreateAnalyticsCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "CreateAnalyticsCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAnalyticsCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Analytics", "CreateAnalyticsCluster", createMysqlAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.AnalyticsClient)

	body, err := testClient.getRequests("mysql", "CreateAnalyticsCluster")
	assert.NoError(t, err)

	type CreateAnalyticsClusterRequestInfo struct {
		ContainerId string
		Request     mysql.CreateAnalyticsClusterRequest
	}

	var requests []CreateAnalyticsClusterRequestInfo
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
			response, err := c.CreateAnalyticsCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlAnalyticsClientDeleteAnalyticsCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "DeleteAnalyticsCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAnalyticsCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Analytics", "DeleteAnalyticsCluster", createMysqlAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.AnalyticsClient)

	body, err := testClient.getRequests("mysql", "DeleteAnalyticsCluster")
	assert.NoError(t, err)

	type DeleteAnalyticsClusterRequestInfo struct {
		ContainerId string
		Request     mysql.DeleteAnalyticsClusterRequest
	}

	var requests []DeleteAnalyticsClusterRequestInfo
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
			response, err := c.DeleteAnalyticsCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlAnalyticsClientGetAnalyticsCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "GetAnalyticsCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAnalyticsCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Analytics", "GetAnalyticsCluster", createMysqlAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.AnalyticsClient)

	body, err := testClient.getRequests("mysql", "GetAnalyticsCluster")
	assert.NoError(t, err)

	type GetAnalyticsClusterRequestInfo struct {
		ContainerId string
		Request     mysql.GetAnalyticsClusterRequest
	}

	var requests []GetAnalyticsClusterRequestInfo
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
			response, err := c.GetAnalyticsCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlAnalyticsClientListAnalyticsClusterShapes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "ListAnalyticsClusterShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAnalyticsClusterShapes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Analytics", "ListAnalyticsClusterShapes", createMysqlAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.AnalyticsClient)

	body, err := testClient.getRequests("mysql", "ListAnalyticsClusterShapes")
	assert.NoError(t, err)

	type ListAnalyticsClusterShapesRequestInfo struct {
		ContainerId string
		Request     mysql.ListAnalyticsClusterShapesRequest
	}

	var requests []ListAnalyticsClusterShapesRequestInfo
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
			response, err := c.ListAnalyticsClusterShapes(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlAnalyticsClientListAnalyticsClusters(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "ListAnalyticsClusters")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAnalyticsClusters is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Analytics", "ListAnalyticsClusters", createMysqlAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.AnalyticsClient)

	body, err := testClient.getRequests("mysql", "ListAnalyticsClusters")
	assert.NoError(t, err)

	type ListAnalyticsClustersRequestInfo struct {
		ContainerId string
		Request     mysql.ListAnalyticsClustersRequest
	}

	var requests []ListAnalyticsClustersRequestInfo
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
				r := req.(*mysql.ListAnalyticsClustersRequest)
				return c.ListAnalyticsClusters(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]mysql.ListAnalyticsClustersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(mysql.ListAnalyticsClustersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlAnalyticsClientStartAnalyticsCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "StartAnalyticsCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StartAnalyticsCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Analytics", "StartAnalyticsCluster", createMysqlAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.AnalyticsClient)

	body, err := testClient.getRequests("mysql", "StartAnalyticsCluster")
	assert.NoError(t, err)

	type StartAnalyticsClusterRequestInfo struct {
		ContainerId string
		Request     mysql.StartAnalyticsClusterRequest
	}

	var requests []StartAnalyticsClusterRequestInfo
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
			response, err := c.StartAnalyticsCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlAnalyticsClientStopAnalyticsCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "StopAnalyticsCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StopAnalyticsCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Analytics", "StopAnalyticsCluster", createMysqlAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.AnalyticsClient)

	body, err := testClient.getRequests("mysql", "StopAnalyticsCluster")
	assert.NoError(t, err)

	type StopAnalyticsClusterRequestInfo struct {
		ContainerId string
		Request     mysql.StopAnalyticsClusterRequest
	}

	var requests []StopAnalyticsClusterRequestInfo
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
			response, err := c.StopAnalyticsCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlAnalyticsClientUpdateAnalyticsCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "UpdateAnalyticsCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAnalyticsCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Analytics", "UpdateAnalyticsCluster", createMysqlAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.AnalyticsClient)

	body, err := testClient.getRequests("mysql", "UpdateAnalyticsCluster")
	assert.NoError(t, err)

	type UpdateAnalyticsClusterRequestInfo struct {
		ContainerId string
		Request     mysql.UpdateAnalyticsClusterRequest
	}

	var requests []UpdateAnalyticsClusterRequestInfo
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
			response, err := c.UpdateAnalyticsCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
