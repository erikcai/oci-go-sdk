// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Notification API
//
// Use the Notification API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notification Overview (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//NotificationClient a client for Notification
type NotificationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewNotificationClientWithConfigurationProvider Creates a new default Notification client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewNotificationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client NotificationClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = NotificationClient{BaseClient: baseClient}
	client.BasePath = "20180801"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *NotificationClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).Endpoint("notification")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *NotificationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *NotificationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateSubscription Creates a subscription for the specified topic.
func (client NotificationClient) CreateSubscription(ctx context.Context, request CreateSubscriptionRequest) (response CreateSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateSubscriptionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSubscriptionResponse")
	}
	return
}

// createSubscription implements the OCIOperation interface (enables retrying operations)
func (client NotificationClient) createSubscription(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/subscriptions")
	if err != nil {
		return nil, err
	}

	var response CreateSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTopic Creates a topic in the specified compartment. For general information about topics, see
// Managing Topics and Subscriptions (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Notification/Tasks/managingtopicsandsubscriptions.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want the topic to reside.
// For information about access control and compartments, see Overview of the IAM Service (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
// You must specify a display name for the topic.
// All Oracle Cloud Infrastructure resources, including topics, get an Oracle-assigned, unique ID called an
// Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID in the response. You can also
// retrieve a resource's OCID by using a List API operation on that resource type, or by viewing the resource in the
// Console. Fore more information, see Resource Identifiers (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
func (client NotificationClient) CreateTopic(ctx context.Context, request CreateTopicRequest) (response CreateTopicResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTopic, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateTopicResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTopicResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTopicResponse")
	}
	return
}

// createTopic implements the OCIOperation interface (enables retrying operations)
func (client NotificationClient) createTopic(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/topics")
	if err != nil {
		return nil, err
	}

	var response CreateTopicResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSubscription Deletes the specified subscription.
func (client NotificationClient) DeleteSubscription(ctx context.Context, request DeleteSubscriptionRequest) (response DeleteSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteSubscriptionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSubscriptionResponse")
	}
	return
}

// deleteSubscription implements the OCIOperation interface (enables retrying operations)
func (client NotificationClient) deleteSubscription(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/subscriptions/{subscriptionId}")
	if err != nil {
		return nil, err
	}

	var response DeleteSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTopic Deletes the specified topic.
func (client NotificationClient) DeleteTopic(ctx context.Context, request DeleteTopicRequest) (response DeleteTopicResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTopic, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteTopicResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTopicResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTopicResponse")
	}
	return
}

// deleteTopic implements the OCIOperation interface (enables retrying operations)
func (client NotificationClient) deleteTopic(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/topics/{topicOcid}")
	if err != nil {
		return nil, err
	}

	var response DeleteTopicResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConfirmSubscription Gets the confirmation details for the specified subscription.
func (client NotificationClient) GetConfirmSubscription(ctx context.Context, request GetConfirmSubscriptionRequest) (response GetConfirmSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConfirmSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetConfirmSubscriptionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConfirmSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConfirmSubscriptionResponse")
	}
	return
}

// getConfirmSubscription implements the OCIOperation interface (enables retrying operations)
func (client NotificationClient) getConfirmSubscription(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptions/{id}/confirmation")
	if err != nil {
		return nil, err
	}

	var response GetConfirmSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSubscription Gets the specified subscription's configuration information.
func (client NotificationClient) GetSubscription(ctx context.Context, request GetSubscriptionRequest) (response GetSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetSubscriptionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSubscriptionResponse")
	}
	return
}

// getSubscription implements the OCIOperation interface (enables retrying operations)
func (client NotificationClient) getSubscription(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptions/{subscriptionId}")
	if err != nil {
		return nil, err
	}

	var response GetSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTopic Gets the specified topic's configuration information.
func (client NotificationClient) GetTopic(ctx context.Context, request GetTopicRequest) (response GetTopicResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTopic, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetTopicResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTopicResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTopicResponse")
	}
	return
}

// getTopic implements the OCIOperation interface (enables retrying operations)
func (client NotificationClient) getTopic(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/topics/{topicOcid}")
	if err != nil {
		return nil, err
	}

	var response GetTopicResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUnsubscription Gets the unsubscription details for the specified subscription.
func (client NotificationClient) GetUnsubscription(ctx context.Context, request GetUnsubscriptionRequest) (response GetUnsubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUnsubscription, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetUnsubscriptionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUnsubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUnsubscriptionResponse")
	}
	return
}

// getUnsubscription implements the OCIOperation interface (enables retrying operations)
func (client NotificationClient) getUnsubscription(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptions/{id}/unsubscription")
	if err != nil {
		return nil, err
	}

	var response GetUnsubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSubscriptions Lists the subscriptions in the specified compartment or for the specified topic.
func (client NotificationClient) ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) (response ListSubscriptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSubscriptions, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListSubscriptionsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSubscriptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSubscriptionsResponse")
	}
	return
}

// listSubscriptions implements the OCIOperation interface (enables retrying operations)
func (client NotificationClient) listSubscriptions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptions")
	if err != nil {
		return nil, err
	}

	var response ListSubscriptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTopics Lists topics in the specified compartment.
func (client NotificationClient) ListTopics(ctx context.Context, request ListTopicsRequest) (response ListTopicsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTopics, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListTopicsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTopicsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTopicsResponse")
	}
	return
}

// listTopics implements the OCIOperation interface (enables retrying operations)
func (client NotificationClient) listTopics(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/topics")
	if err != nil {
		return nil, err
	}

	var response ListTopicsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PublishMessage Publishes a message to the specified topic. For more information about publishing messages, see Publishing Messages (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Notification/Tasks/publishingmessages.htm).
func (client NotificationClient) PublishMessage(ctx context.Context, request PublishMessageRequest) (response PublishMessageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.publishMessage, policy)
	if err != nil {
		if ociResponse != nil {
			response = PublishMessageResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PublishMessageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PublishMessageResponse")
	}
	return
}

// publishMessage implements the OCIOperation interface (enables retrying operations)
func (client NotificationClient) publishMessage(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/topics/{topicOcid}/messages")
	if err != nil {
		return nil, err
	}

	var response PublishMessageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResendSubscriptionConfirmation Resends the confirmation details for the specified subscription.
func (client NotificationClient) ResendSubscriptionConfirmation(ctx context.Context, request ResendSubscriptionConfirmationRequest) (response ResendSubscriptionConfirmationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.resendSubscriptionConfirmation, policy)
	if err != nil {
		if ociResponse != nil {
			response = ResendSubscriptionConfirmationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResendSubscriptionConfirmationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResendSubscriptionConfirmationResponse")
	}
	return
}

// resendSubscriptionConfirmation implements the OCIOperation interface (enables retrying operations)
func (client NotificationClient) resendSubscriptionConfirmation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/subscriptions/{id}/resendConfirmation")
	if err != nil {
		return nil, err
	}

	var response ResendSubscriptionConfirmationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSubscription Updates the specified subscription's configuration.
func (client NotificationClient) UpdateSubscription(ctx context.Context, request UpdateSubscriptionRequest) (response UpdateSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateSubscriptionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSubscriptionResponse")
	}
	return
}

// updateSubscription implements the OCIOperation interface (enables retrying operations)
func (client NotificationClient) updateSubscription(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/subscriptions/{subscriptionId}")
	if err != nil {
		return nil, err
	}

	var response UpdateSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTopic Updates the specified topic's configuration.
func (client NotificationClient) UpdateTopic(ctx context.Context, request UpdateTopicRequest) (response UpdateTopicResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTopic, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateTopicResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTopicResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTopicResponse")
	}
	return
}

// updateTopic implements the OCIOperation interface (enables retrying operations)
func (client NotificationClient) updateTopic(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/topics/{topicOcid}")
	if err != nil {
		return nil, err
	}

	var response UpdateTopicResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
