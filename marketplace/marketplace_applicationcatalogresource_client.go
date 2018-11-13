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

//ApplicationCatalogResourceClient a client for ApplicationCatalogResource
type ApplicationCatalogResourceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewApplicationCatalogResourceClientWithConfigurationProvider Creates a new default ApplicationCatalogResource client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewApplicationCatalogResourceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ApplicationCatalogResourceClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = ApplicationCatalogResourceClient{BaseClient: baseClient}
	client.BasePath = "20181001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ApplicationCatalogResourceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("marketplace", "https://marketplace.us-ashburn-1.oci.oraclecloud.com")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ApplicationCatalogResourceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ApplicationCatalogResourceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetAppCatalogApplicationMapping This endpoint returns the app catalog listing information for a given marketplaceApplicationId.
func (client ApplicationCatalogResourceClient) GetAppCatalogApplicationMapping(ctx context.Context, request GetAppCatalogApplicationMappingRequest) (response GetAppCatalogApplicationMappingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAppCatalogApplicationMapping, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAppCatalogApplicationMappingResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAppCatalogApplicationMappingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAppCatalogApplicationMappingResponse")
	}
	return
}

// getAppCatalogApplicationMapping implements the OCIOperation interface (enables retrying operations)
func (client ApplicationCatalogResourceClient) getAppCatalogApplicationMapping(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/appCatalogApplicationMapping")
	if err != nil {
		return nil, err
	}

	var response GetAppCatalogApplicationMappingResponse
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
