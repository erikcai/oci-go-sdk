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
func TestDbSystemClientCreateInBoundChannel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "CreateInBoundChannel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateInBoundChannel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "CreateInBoundChannel", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "CreateInBoundChannel")
	assert.NoError(t, err)

	type CreateInBoundChannelRequestInfo struct {
		ContainerId string
		Request     mysql.CreateInBoundChannelRequest
	}

	var requests []CreateInBoundChannelRequestInfo
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
			response, err := c.CreateInBoundChannel(context.Background(), req.Request)
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
func TestDbSystemClientDeleteInBoundChannel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "DeleteInBoundChannel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteInBoundChannel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "DeleteInBoundChannel", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "DeleteInBoundChannel")
	assert.NoError(t, err)

	type DeleteInBoundChannelRequestInfo struct {
		ContainerId string
		Request     mysql.DeleteInBoundChannelRequest
	}

	var requests []DeleteInBoundChannelRequestInfo
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
			response, err := c.DeleteInBoundChannel(context.Background(), req.Request)
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
func TestDbSystemClientGetInBoundChannel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "GetInBoundChannel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInBoundChannel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "GetInBoundChannel", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "GetInBoundChannel")
	assert.NoError(t, err)

	type GetInBoundChannelRequestInfo struct {
		ContainerId string
		Request     mysql.GetInBoundChannelRequest
	}

	var requests []GetInBoundChannelRequestInfo
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
			response, err := c.GetInBoundChannel(context.Background(), req.Request)
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
func TestDbSystemClientListInBoundChannels(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "ListInBoundChannels")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInBoundChannels is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "ListInBoundChannels", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "ListInBoundChannels")
	assert.NoError(t, err)

	type ListInBoundChannelsRequestInfo struct {
		ContainerId string
		Request     mysql.ListInBoundChannelsRequest
	}

	var requests []ListInBoundChannelsRequestInfo
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
			response, err := c.ListInBoundChannels(context.Background(), req.Request)
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
func TestDbSystemClientRestoreDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "RestoreDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestoreDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbSystem", "RestoreDbSystem", createDbSystemClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbSystemClient)

	body, err := testClient.getRequests("mysql", "RestoreDbSystem")
	assert.NoError(t, err)

	type RestoreDbSystemRequestInfo struct {
		ContainerId string
		Request     mysql.RestoreDbSystemRequest
	}

	var requests []RestoreDbSystemRequestInfo
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
			response, err := c.RestoreDbSystem(context.Background(), req.Request)
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
