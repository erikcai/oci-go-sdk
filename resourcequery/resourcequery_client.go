// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Resource Query Service
//
// Query for resources across your cloud infrastructure
//

package resourcequery

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//ResourceQueryClient a client for ResourceQuery
type ResourceQueryClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewResourceQueryClientWithConfigurationProvider Creates a new default ResourceQuery client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewResourceQueryClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ResourceQueryClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = ResourceQueryClient{BaseClient: baseClient}
	client.BasePath = "20180409"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ResourceQueryClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "query", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ResourceQueryClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ResourceQueryClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetResourceType The GetResourceType API provides a way to get all of the resource type information using either the type name or its OCID.
func (client ResourceQueryClient) GetResourceType(ctx context.Context, request GetResourceTypeRequest) (response GetResourceTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResourceType, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(GetResourceTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResourceTypeResponse")
	}
	return
}

// getResourceType implements the OCIOperation interface (enables retrying operations)
func (client ResourceQueryClient) getResourceType(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceTypes/{name}")
	if err != nil {
		return nil, err
	}

	var response GetResourceTypeResponse
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

// ListResourceTypes The ListResourceTypes API provides a way to discover all resource types that are available for querying.
func (client ResourceQueryClient) ListResourceTypes(ctx context.Context, request ListResourceTypesRequest) (response ListResourceTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResourceTypes, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceTypesResponse")
	}
	return
}

// listResourceTypes implements the OCIOperation interface (enables retrying operations)
func (client ResourceQueryClient) listResourceTypes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceTypes")
	if err != nil {
		return nil, err
	}

	var response ListResourceTypesResponse
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

// SearchResources The query API allows you to search across all of your cloud infrastructure to find resources matching different criteria that you have permissions to access.  Results may be across different types, and across compartments.
func (client ResourceQueryClient) SearchResources(ctx context.Context, request SearchResourcesRequest) (response SearchResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.searchResources, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(SearchResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchResourcesResponse")
	}
	return
}

// searchResources implements the OCIOperation interface (enables retrying operations)
func (client ResourceQueryClient) searchResources(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/resources")
	if err != nil {
		return nil, err
	}

	var response SearchResourcesResponse
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
