package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/ons"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestNotificationClientCreateSubscription(t *testing.T) {
	enabled, err := testClient.isApiEnabled("ons", "CreateSubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSubscription is not enabled by the testing service")
	}
	c, err := ons.NewNotificationClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("ons", "CreateSubscription")
	assert.NoError(t, err)

	type CreateSubscriptionRequestInfo struct {
		ContainerId string
		Request     ons.CreateSubscriptionRequest
	}

	var requests []CreateSubscriptionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateSubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestNotificationClientCreateTopic(t *testing.T) {
	enabled, err := testClient.isApiEnabled("ons", "CreateTopic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTopic is not enabled by the testing service")
	}
	c, err := ons.NewNotificationClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("ons", "CreateTopic")
	assert.NoError(t, err)

	type CreateTopicRequestInfo struct {
		ContainerId string
		Request     ons.CreateTopicRequest
	}

	var requests []CreateTopicRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateTopic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestNotificationClientDeleteSubscription(t *testing.T) {
	enabled, err := testClient.isApiEnabled("ons", "DeleteSubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSubscription is not enabled by the testing service")
	}
	c, err := ons.NewNotificationClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("ons", "DeleteSubscription")
	assert.NoError(t, err)

	type DeleteSubscriptionRequestInfo struct {
		ContainerId string
		Request     ons.DeleteSubscriptionRequest
	}

	var requests []DeleteSubscriptionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteSubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestNotificationClientDeleteTopic(t *testing.T) {
	enabled, err := testClient.isApiEnabled("ons", "DeleteTopic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteTopic is not enabled by the testing service")
	}
	c, err := ons.NewNotificationClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("ons", "DeleteTopic")
	assert.NoError(t, err)

	type DeleteTopicRequestInfo struct {
		ContainerId string
		Request     ons.DeleteTopicRequest
	}

	var requests []DeleteTopicRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteTopic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestNotificationClientGetConfirmSubscription(t *testing.T) {
	enabled, err := testClient.isApiEnabled("ons", "GetConfirmSubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConfirmSubscription is not enabled by the testing service")
	}
	c, err := ons.NewNotificationClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("ons", "GetConfirmSubscription")
	assert.NoError(t, err)

	type GetConfirmSubscriptionRequestInfo struct {
		ContainerId string
		Request     ons.GetConfirmSubscriptionRequest
	}

	var requests []GetConfirmSubscriptionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetConfirmSubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestNotificationClientGetSubscription(t *testing.T) {
	enabled, err := testClient.isApiEnabled("ons", "GetSubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSubscription is not enabled by the testing service")
	}
	c, err := ons.NewNotificationClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("ons", "GetSubscription")
	assert.NoError(t, err)

	type GetSubscriptionRequestInfo struct {
		ContainerId string
		Request     ons.GetSubscriptionRequest
	}

	var requests []GetSubscriptionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetSubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestNotificationClientGetTopic(t *testing.T) {
	enabled, err := testClient.isApiEnabled("ons", "GetTopic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTopic is not enabled by the testing service")
	}
	c, err := ons.NewNotificationClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("ons", "GetTopic")
	assert.NoError(t, err)

	type GetTopicRequestInfo struct {
		ContainerId string
		Request     ons.GetTopicRequest
	}

	var requests []GetTopicRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetTopic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestNotificationClientGetUnsubscription(t *testing.T) {
	enabled, err := testClient.isApiEnabled("ons", "GetUnsubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetUnsubscription is not enabled by the testing service")
	}
	c, err := ons.NewNotificationClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("ons", "GetUnsubscription")
	assert.NoError(t, err)

	type GetUnsubscriptionRequestInfo struct {
		ContainerId string
		Request     ons.GetUnsubscriptionRequest
	}

	var requests []GetUnsubscriptionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetUnsubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestNotificationClientListSubscriptions(t *testing.T) {
	enabled, err := testClient.isApiEnabled("ons", "ListSubscriptions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSubscriptions is not enabled by the testing service")
	}
	c, err := ons.NewNotificationClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("ons", "ListSubscriptions")
	assert.NoError(t, err)

	type ListSubscriptionsRequestInfo struct {
		ContainerId string
		Request     ons.ListSubscriptionsRequest
	}

	var requests []ListSubscriptionsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*ons.ListSubscriptionsRequest)
				return c.ListSubscriptions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]ons.ListSubscriptionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(ons.ListSubscriptionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestNotificationClientListTopics(t *testing.T) {
	enabled, err := testClient.isApiEnabled("ons", "ListTopics")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTopics is not enabled by the testing service")
	}
	c, err := ons.NewNotificationClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("ons", "ListTopics")
	assert.NoError(t, err)

	type ListTopicsRequestInfo struct {
		ContainerId string
		Request     ons.ListTopicsRequest
	}

	var requests []ListTopicsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*ons.ListTopicsRequest)
				return c.ListTopics(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]ons.ListTopicsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(ons.ListTopicsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestNotificationClientPublishMessage(t *testing.T) {
	enabled, err := testClient.isApiEnabled("ons", "PublishMessage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PublishMessage is not enabled by the testing service")
	}
	c, err := ons.NewNotificationClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("ons", "PublishMessage")
	assert.NoError(t, err)

	type PublishMessageRequestInfo struct {
		ContainerId string
		Request     ons.PublishMessageRequest
	}

	var requests []PublishMessageRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.PublishMessage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestNotificationClientResendSubscriptionConfirmation(t *testing.T) {
	enabled, err := testClient.isApiEnabled("ons", "ResendSubscriptionConfirmation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ResendSubscriptionConfirmation is not enabled by the testing service")
	}
	c, err := ons.NewNotificationClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("ons", "ResendSubscriptionConfirmation")
	assert.NoError(t, err)

	type ResendSubscriptionConfirmationRequestInfo struct {
		ContainerId string
		Request     ons.ResendSubscriptionConfirmationRequest
	}

	var requests []ResendSubscriptionConfirmationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ResendSubscriptionConfirmation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestNotificationClientUpdateSubscription(t *testing.T) {
	enabled, err := testClient.isApiEnabled("ons", "UpdateSubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSubscription is not enabled by the testing service")
	}
	c, err := ons.NewNotificationClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("ons", "UpdateSubscription")
	assert.NoError(t, err)

	type UpdateSubscriptionRequestInfo struct {
		ContainerId string
		Request     ons.UpdateSubscriptionRequest
	}

	var requests []UpdateSubscriptionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateSubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestNotificationClientUpdateTopic(t *testing.T) {
	enabled, err := testClient.isApiEnabled("ons", "UpdateTopic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTopic is not enabled by the testing service")
	}
	c, err := ons.NewNotificationClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("ons", "UpdateTopic")
	assert.NoError(t, err)

	type UpdateTopicRequestInfo struct {
		ContainerId string
		Request     ons.UpdateTopicRequest
	}

	var requests []UpdateTopicRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateTopic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
