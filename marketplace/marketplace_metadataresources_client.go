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

//MetadataResourcesClient a client for MetadataResources
type MetadataResourcesClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMetadataResourcesClientWithConfigurationProvider Creates a new default MetadataResources client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMetadataResourcesClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MetadataResourcesClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = MetadataResourcesClient{BaseClient: baseClient}
	client.BasePath = "20181001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MetadataResourcesClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("marketplace", "https://marketplace.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MetadataResourcesClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *MetadataResourcesClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListCategories This endpoint is used for getting the list of all available categories for the specified product type. It is used to populate the category filter selections of the search bar.
func (client MetadataResourcesClient) ListCategories(ctx context.Context, request ListCategoriesRequest) (response ListCategoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCategories, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListCategoriesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCategoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCategoriesResponse")
	}
	return
}

// listCategories implements the OCIOperation interface (enables retrying operations)
func (client MetadataResourcesClient) listCategories(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/products/{productCode}/categories")
	if err != nil {
		return nil, err
	}

	var response ListCategoriesResponse
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

// ListProductFilters The endpoint will get the list of all applicable filters for the specified product type. The filters can be further used for adding more restrictions on the result set of applications.
func (client MetadataResourcesClient) ListProductFilters(ctx context.Context, request ListProductFiltersRequest) (response ListProductFiltersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProductFilters, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListProductFiltersResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProductFiltersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProductFiltersResponse")
	}
	return
}

// listProductFilters implements the OCIOperation interface (enables retrying operations)
func (client MetadataResourcesClient) listProductFilters(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/products/{productCode}/filters")
	if err != nil {
		return nil, err
	}

	var response ListProductFiltersResponse
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
