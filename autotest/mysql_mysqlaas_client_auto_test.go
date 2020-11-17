package autotest

import (
	"github.com/oracle/oci-go-sdk/v29/common"
	"github.com/oracle/oci-go-sdk/v29/mysql"

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
