package autotest

import (
	"github.com/oracle/oci-go-sdk/v30/common"
	"github.com/oracle/oci-go-sdk/v30/mysql"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createDbSystemClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := mysql.NewDbSystemClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbSystemClientAddAnalyticsCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "AddAnalyticsCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AddAnalyticsCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "AddAnalyticsCluster", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "AddAnalyticsCluster")
	assert.NoError(t, err)

	type AddAnalyticsClusterRequestInfo struct {
		ContainerId string
		Request     mysql.AddAnalyticsClusterRequest
	}

	var requests []AddAnalyticsClusterRequestInfo
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
			response, err := c.AddAnalyticsCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbSystemClientCreateDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "CreateDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "CreateDbSystem", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "CreateDbSystem")
	assert.NoError(t, err)

	type CreateDbSystemRequestInfo struct {
		ContainerId string
		Request     mysql.CreateDbSystemRequest
	}

	var requests []CreateDbSystemRequestInfo
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
			response, err := c.CreateDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbSystemClientDeleteAnalyticsCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "DeleteAnalyticsCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAnalyticsCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "DeleteAnalyticsCluster", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

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
func TestDbSystemClientDeleteDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "DeleteDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "DeleteDbSystem", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "DeleteDbSystem")
	assert.NoError(t, err)

	type DeleteDbSystemRequestInfo struct {
		ContainerId string
		Request     mysql.DeleteDbSystemRequest
	}

	var requests []DeleteDbSystemRequestInfo
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
			response, err := c.DeleteDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbSystemClientGenerateAnalyticsClusterMemoryEstimate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "GenerateAnalyticsClusterMemoryEstimate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GenerateAnalyticsClusterMemoryEstimate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "GenerateAnalyticsClusterMemoryEstimate", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "GenerateAnalyticsClusterMemoryEstimate")
	assert.NoError(t, err)

	type GenerateAnalyticsClusterMemoryEstimateRequestInfo struct {
		ContainerId string
		Request     mysql.GenerateAnalyticsClusterMemoryEstimateRequest
	}

	var requests []GenerateAnalyticsClusterMemoryEstimateRequestInfo
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
			response, err := c.GenerateAnalyticsClusterMemoryEstimate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbSystemClientGetAnalyticsCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "GetAnalyticsCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAnalyticsCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "GetAnalyticsCluster", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

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
func TestDbSystemClientGetAnalyticsClusterMemoryEstimate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "GetAnalyticsClusterMemoryEstimate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAnalyticsClusterMemoryEstimate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "GetAnalyticsClusterMemoryEstimate", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "GetAnalyticsClusterMemoryEstimate")
	assert.NoError(t, err)

	type GetAnalyticsClusterMemoryEstimateRequestInfo struct {
		ContainerId string
		Request     mysql.GetAnalyticsClusterMemoryEstimateRequest
	}

	var requests []GetAnalyticsClusterMemoryEstimateRequestInfo
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
			response, err := c.GetAnalyticsClusterMemoryEstimate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbSystemClientGetDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "GetDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "GetDbSystem", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "GetDbSystem")
	assert.NoError(t, err)

	type GetDbSystemRequestInfo struct {
		ContainerId string
		Request     mysql.GetDbSystemRequest
	}

	var requests []GetDbSystemRequestInfo
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
			response, err := c.GetDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbSystemClientListDbSystems(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "ListDbSystems")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbSystems is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "ListDbSystems", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "ListDbSystems")
	assert.NoError(t, err)

	type ListDbSystemsRequestInfo struct {
		ContainerId string
		Request     mysql.ListDbSystemsRequest
	}

	var requests []ListDbSystemsRequestInfo
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
				r := req.(*mysql.ListDbSystemsRequest)
				return c.ListDbSystems(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]mysql.ListDbSystemsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(mysql.ListDbSystemsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbSystemClientRestartAnalyticsCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "RestartAnalyticsCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestartAnalyticsCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "RestartAnalyticsCluster", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "RestartAnalyticsCluster")
	assert.NoError(t, err)

	type RestartAnalyticsClusterRequestInfo struct {
		ContainerId string
		Request     mysql.RestartAnalyticsClusterRequest
	}

	var requests []RestartAnalyticsClusterRequestInfo
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
			response, err := c.RestartAnalyticsCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbSystemClientRestartDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "RestartDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestartDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "RestartDbSystem", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "RestartDbSystem")
	assert.NoError(t, err)

	type RestartDbSystemRequestInfo struct {
		ContainerId string
		Request     mysql.RestartDbSystemRequest
	}

	var requests []RestartDbSystemRequestInfo
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
			response, err := c.RestartDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbSystemClientStartAnalyticsCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "StartAnalyticsCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StartAnalyticsCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "StartAnalyticsCluster", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

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
func TestDbSystemClientStartDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "StartDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StartDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "StartDbSystem", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "StartDbSystem")
	assert.NoError(t, err)

	type StartDbSystemRequestInfo struct {
		ContainerId string
		Request     mysql.StartDbSystemRequest
	}

	var requests []StartDbSystemRequestInfo
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
			response, err := c.StartDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbSystemClientStopAnalyticsCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "StopAnalyticsCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StopAnalyticsCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "StopAnalyticsCluster", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

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
func TestDbSystemClientStopDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "StopDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StopDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "StopDbSystem", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "StopDbSystem")
	assert.NoError(t, err)

	type StopDbSystemRequestInfo struct {
		ContainerId string
		Request     mysql.StopDbSystemRequest
	}

	var requests []StopDbSystemRequestInfo
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
			response, err := c.StopDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbSystemClientUpdateAnalyticsCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "UpdateAnalyticsCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAnalyticsCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "UpdateAnalyticsCluster", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

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

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbSystemClientUpdateDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "UpdateDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "UpdateDbSystem", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "UpdateDbSystem")
	assert.NoError(t, err)

	type UpdateDbSystemRequestInfo struct {
		ContainerId string
		Request     mysql.UpdateDbSystemRequest
	}

	var requests []UpdateDbSystemRequestInfo
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
			response, err := c.UpdateDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
