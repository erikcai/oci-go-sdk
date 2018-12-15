// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace REST API endpoint
//
// Marketplace REST API for loom plugin
//

package marketplace

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//PublisherResourcesClient a client for PublisherResources
type PublisherResourcesClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewPublisherResourcesClientWithConfigurationProvider Creates a new default PublisherResources client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewPublisherResourcesClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client PublisherResourcesClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = PublisherResourcesClient{BaseClient: baseClient}
	client.BasePath = "20181001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *PublisherResourcesClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("marketplace", "https://marketplace.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *PublisherResourcesClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *PublisherResourcesClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetPublisher Get the details of the specified publisher including the description, headquarter location, year founded, etc
func (client PublisherResourcesClient) GetPublisher(ctx context.Context, request GetPublisherRequest) (response GetPublisherResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPublisher, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetPublisherResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPublisherResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPublisherResponse")
	}
	return
}

// getPublisher implements the OCIOperation interface (enables retrying operations)
func (client PublisherResourcesClient) getPublisher(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/publishers/{publisherId}")
	if err != nil {
		return nil, err
	}

	var response GetPublisherResponse
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

// ListPublisherApplications Get the list of applications by the specified publisher ID.
func (client PublisherResourcesClient) ListPublisherApplications(ctx context.Context, request ListPublisherApplicationsRequest) (response ListPublisherApplicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPublisherApplications, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListPublisherApplicationsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPublisherApplicationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPublisherApplicationsResponse")
	}
	return
}

// listPublisherApplications implements the OCIOperation interface (enables retrying operations)
func (client PublisherResourcesClient) listPublisherApplications(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/publishers/{publisherId}/applications")
	if err != nil {
		return nil, err
	}

	var response ListPublisherApplicationsResponse
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
