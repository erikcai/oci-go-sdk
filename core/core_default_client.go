// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//DefaultClient a client for Default
type DefaultClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDefaultClientWithConfigurationProvider Creates a new default Default client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDefaultClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DefaultClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = DefaultClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DefaultClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "iaas", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DefaultClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DefaultClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListBlockstorageConsumption Lists the corresponding usage for the limit group in the specified tenancy, targeting a Region or an Availability Domain.
func (client DefaultClient) ListBlockstorageConsumption(ctx context.Context, request ListBlockstorageConsumptionRequest) (response ListBlockstorageConsumptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBlockstorageConsumption, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListBlockstorageConsumptionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBlockstorageConsumptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBlockstorageConsumptionResponse")
	}
	return
}

// listBlockstorageConsumption implements the OCIOperation interface (enables retrying operations)
func (client DefaultClient) listBlockstorageConsumption(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/blockstorageConsumption")
	if err != nil {
		return nil, err
	}

	var response ListBlockstorageConsumptionResponse
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
