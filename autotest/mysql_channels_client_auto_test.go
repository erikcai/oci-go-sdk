package autotest

import (
	"github.com/oracle/oci-go-sdk/v33/common"
	"github.com/oracle/oci-go-sdk/v33/mysql"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createMysqlChannelsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := mysql.NewChannelsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlChannelsClientCreateChannel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "CreateChannel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateChannel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Channels", "CreateChannel", createMysqlChannelsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.ChannelsClient)

	body, err := testClient.getRequests("mysql", "CreateChannel")
	assert.NoError(t, err)

	type CreateChannelRequestInfo struct {
		ContainerId string
		Request     mysql.CreateChannelRequest
	}

	var requests []CreateChannelRequestInfo
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
			response, err := c.CreateChannel(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlChannelsClientDeleteChannel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "DeleteChannel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteChannel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Channels", "DeleteChannel", createMysqlChannelsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.ChannelsClient)

	body, err := testClient.getRequests("mysql", "DeleteChannel")
	assert.NoError(t, err)

	type DeleteChannelRequestInfo struct {
		ContainerId string
		Request     mysql.DeleteChannelRequest
	}

	var requests []DeleteChannelRequestInfo
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
			response, err := c.DeleteChannel(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlChannelsClientGetChannel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "GetChannel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetChannel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Channels", "GetChannel", createMysqlChannelsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.ChannelsClient)

	body, err := testClient.getRequests("mysql", "GetChannel")
	assert.NoError(t, err)

	type GetChannelRequestInfo struct {
		ContainerId string
		Request     mysql.GetChannelRequest
	}

	var requests []GetChannelRequestInfo
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
			response, err := c.GetChannel(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlChannelsClientListChannels(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "ListChannels")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListChannels is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Channels", "ListChannels", createMysqlChannelsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.ChannelsClient)

	body, err := testClient.getRequests("mysql", "ListChannels")
	assert.NoError(t, err)

	type ListChannelsRequestInfo struct {
		ContainerId string
		Request     mysql.ListChannelsRequest
	}

	var requests []ListChannelsRequestInfo
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
				r := req.(*mysql.ListChannelsRequest)
				return c.ListChannels(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]mysql.ListChannelsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(mysql.ListChannelsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlChannelsClientResetChannel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "ResetChannel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ResetChannel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Channels", "ResetChannel", createMysqlChannelsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.ChannelsClient)

	body, err := testClient.getRequests("mysql", "ResetChannel")
	assert.NoError(t, err)

	type ResetChannelRequestInfo struct {
		ContainerId string
		Request     mysql.ResetChannelRequest
	}

	var requests []ResetChannelRequestInfo
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
			response, err := c.ResetChannel(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlChannelsClientResumeChannel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "ResumeChannel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ResumeChannel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Channels", "ResumeChannel", createMysqlChannelsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.ChannelsClient)

	body, err := testClient.getRequests("mysql", "ResumeChannel")
	assert.NoError(t, err)

	type ResumeChannelRequestInfo struct {
		ContainerId string
		Request     mysql.ResumeChannelRequest
	}

	var requests []ResumeChannelRequestInfo
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
			response, err := c.ResumeChannel(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestMysqlChannelsClientUpdateChannel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "UpdateChannel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateChannel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "Channels", "UpdateChannel", createMysqlChannelsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.ChannelsClient)

	body, err := testClient.getRequests("mysql", "UpdateChannel")
	assert.NoError(t, err)

	type UpdateChannelRequestInfo struct {
		ContainerId string
		Request     mysql.UpdateChannelRequest
	}

	var requests []UpdateChannelRequestInfo
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
			response, err := c.UpdateChannel(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
