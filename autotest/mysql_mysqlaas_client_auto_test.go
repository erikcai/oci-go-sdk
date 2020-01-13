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

func createMysqlaasClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := mysql.NewMysqlaasClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientCopyConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "CopyConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CopyConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "CopyConfiguration", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "CopyConfiguration")
	assert.NoError(t, err)

	type CopyConfigurationRequestInfo struct {
		ContainerId string
		Request     mysql.CopyConfigurationRequest
	}

	var requests []CopyConfigurationRequestInfo
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
			response, err := c.CopyConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientCreateConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "CreateConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "CreateConfiguration", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "CreateConfiguration")
	assert.NoError(t, err)

	type CreateConfigurationRequestInfo struct {
		ContainerId string
		Request     mysql.CreateConfigurationRequest
	}

	var requests []CreateConfigurationRequestInfo
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
			response, err := c.CreateConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientDeleteConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "DeleteConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "DeleteConfiguration", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "DeleteConfiguration")
	assert.NoError(t, err)

	type DeleteConfigurationRequestInfo struct {
		ContainerId string
		Request     mysql.DeleteConfigurationRequest
	}

	var requests []DeleteConfigurationRequestInfo
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
			response, err := c.DeleteConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientDeleteInstance(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "DeleteInstance")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteInstance is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "DeleteInstance", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "DeleteInstance")
	assert.NoError(t, err)

	type DeleteInstanceRequestInfo struct {
		ContainerId string
		Request     mysql.DeleteInstanceRequest
	}

	var requests []DeleteInstanceRequestInfo
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
			response, err := c.DeleteInstance(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientEstimateAnalyticsClusterMemory(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "EstimateAnalyticsClusterMemory")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("EstimateAnalyticsClusterMemory is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "EstimateAnalyticsClusterMemory", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "EstimateAnalyticsClusterMemory")
	assert.NoError(t, err)

	type EstimateAnalyticsClusterMemoryRequestInfo struct {
		ContainerId string
		Request     mysql.EstimateAnalyticsClusterMemoryRequest
	}

	var requests []EstimateAnalyticsClusterMemoryRequestInfo
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
			response, err := c.EstimateAnalyticsClusterMemory(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientGetAnalyticsClusterMemoryEstimate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "GetAnalyticsClusterMemoryEstimate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAnalyticsClusterMemoryEstimate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "GetAnalyticsClusterMemoryEstimate", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

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
func TestMysqlaasClientGetConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "GetConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "GetConfiguration", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "GetConfiguration")
	assert.NoError(t, err)

	type GetConfigurationRequestInfo struct {
		ContainerId string
		Request     mysql.GetConfigurationRequest
	}

	var requests []GetConfigurationRequestInfo
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
			response, err := c.GetConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientGetConfigurationRevision(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "GetConfigurationRevision")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConfigurationRevision is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "GetConfigurationRevision", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "GetConfigurationRevision")
	assert.NoError(t, err)

	type GetConfigurationRevisionRequestInfo struct {
		ContainerId string
		Request     mysql.GetConfigurationRevisionRequest
	}

	var requests []GetConfigurationRevisionRequestInfo
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
			response, err := c.GetConfigurationRevision(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientGetInstance(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "GetInstance")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInstance is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "GetInstance", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "GetInstance")
	assert.NoError(t, err)

	type GetInstanceRequestInfo struct {
		ContainerId string
		Request     mysql.GetInstanceRequest
	}

	var requests []GetInstanceRequestInfo
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
			response, err := c.GetInstance(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientGetVersionOptionsMetadata(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "GetVersionOptionsMetadata")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetVersionOptionsMetadata is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "GetVersionOptionsMetadata", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "GetVersionOptionsMetadata")
	assert.NoError(t, err)

	type GetVersionOptionsMetadataRequestInfo struct {
		ContainerId string
		Request     mysql.GetVersionOptionsMetadataRequest
	}

	var requests []GetVersionOptionsMetadataRequestInfo
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
			response, err := c.GetVersionOptionsMetadata(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientListConfigurationRevisions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "ListConfigurationRevisions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListConfigurationRevisions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "ListConfigurationRevisions", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "ListConfigurationRevisions")
	assert.NoError(t, err)

	type ListConfigurationRevisionsRequestInfo struct {
		ContainerId string
		Request     mysql.ListConfigurationRevisionsRequest
	}

	var requests []ListConfigurationRevisionsRequestInfo
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
				r := req.(*mysql.ListConfigurationRevisionsRequest)
				return c.ListConfigurationRevisions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]mysql.ListConfigurationRevisionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(mysql.ListConfigurationRevisionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientListConfigurations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "ListConfigurations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListConfigurations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "ListConfigurations", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "ListConfigurations")
	assert.NoError(t, err)

	type ListConfigurationsRequestInfo struct {
		ContainerId string
		Request     mysql.ListConfigurationsRequest
	}

	var requests []ListConfigurationsRequestInfo
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
				r := req.(*mysql.ListConfigurationsRequest)
				return c.ListConfigurations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]mysql.ListConfigurationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(mysql.ListConfigurationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientListInstances(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "ListInstances")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInstances is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "ListInstances", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "ListInstances")
	assert.NoError(t, err)

	type ListInstancesRequestInfo struct {
		ContainerId string
		Request     mysql.ListInstancesRequest
	}

	var requests []ListInstancesRequestInfo
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
				r := req.(*mysql.ListInstancesRequest)
				return c.ListInstances(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]mysql.ListInstancesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(mysql.ListInstancesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientListShapes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "ListShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListShapes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "ListShapes", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "ListShapes")
	assert.NoError(t, err)

	type ListShapesRequestInfo struct {
		ContainerId string
		Request     mysql.ListShapesRequest
	}

	var requests []ListShapesRequestInfo
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
			response, err := c.ListShapes(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientListVersions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "ListVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListVersions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "ListVersions", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "ListVersions")
	assert.NoError(t, err)

	type ListVersionsRequestInfo struct {
		ContainerId string
		Request     mysql.ListVersionsRequest
	}

	var requests []ListVersionsRequestInfo
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
			response, err := c.ListVersions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientRestartInstance(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "RestartInstance")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestartInstance is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "RestartInstance", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "RestartInstance")
	assert.NoError(t, err)

	type RestartInstanceRequestInfo struct {
		ContainerId string
		Request     mysql.RestartInstanceRequest
	}

	var requests []RestartInstanceRequestInfo
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
			response, err := c.RestartInstance(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientStartInstance(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "StartInstance")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StartInstance is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "StartInstance", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "StartInstance")
	assert.NoError(t, err)

	type StartInstanceRequestInfo struct {
		ContainerId string
		Request     mysql.StartInstanceRequest
	}

	var requests []StartInstanceRequestInfo
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
			response, err := c.StartInstance(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientStopInstance(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "StopInstance")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StopInstance is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "StopInstance", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "StopInstance")
	assert.NoError(t, err)

	type StopInstanceRequestInfo struct {
		ContainerId string
		Request     mysql.StopInstanceRequest
	}

	var requests []StopInstanceRequestInfo
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
			response, err := c.StopInstance(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientUpdateConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "UpdateConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "UpdateConfiguration", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "UpdateConfiguration")
	assert.NoError(t, err)

	type UpdateConfigurationRequestInfo struct {
		ContainerId string
		Request     mysql.UpdateConfigurationRequest
	}

	var requests []UpdateConfigurationRequestInfo
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
			response, err := c.UpdateConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlaasClientUpdateInstance(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "UpdateInstance")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateInstance is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Mysqlaas", "UpdateInstance", createMysqlaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.MysqlaasClient)

	body, err := testClient.getRequests("mysql", "UpdateInstance")
	assert.NoError(t, err)

	type UpdateInstanceRequestInfo struct {
		ContainerId string
		Request     mysql.UpdateInstanceRequest
	}

	var requests []UpdateInstanceRequestInfo
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
			response, err := c.UpdateInstance(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
