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
	"strings"
)

//ApplicationInstanceResourceClient a client for ApplicationInstanceResource
type ApplicationInstanceResourceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewApplicationInstanceResourceClientWithConfigurationProvider Creates a new default ApplicationInstanceResource client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewApplicationInstanceResourceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ApplicationInstanceResourceClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = ApplicationInstanceResourceClient{BaseClient: baseClient}
	client.BasePath = "20181001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ApplicationInstanceResourceClient) SetRegion(region string) {
	client.Host = strings.Replace("https://marketplace.us-ashburn-1.oci.oraclecloud.com", "{region}", region, 1)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ApplicationInstanceResourceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ApplicationInstanceResourceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetInstallConfiguration This endpoint returns a set of configuration properties that are required to install the Application. These properties are configured by the Application publisher. The values for these properties need to be provided by the end user. The payload returned describes each property by providing relevant metadata like data type, mandatory etc.
func (client ApplicationInstanceResourceClient) GetInstallConfiguration(ctx context.Context, request GetInstallConfigurationRequest) (response GetInstallConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInstallConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetInstallConfigurationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInstallConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInstallConfigurationResponse")
	}
	return
}

// getInstallConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ApplicationInstanceResourceClient) getInstallConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applications/{applicationId}/installconfigurations")
	if err != nil {
		return nil, err
	}

	var response GetInstallConfigurationResponse
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

// GetTermsOfUses This API returns the Terms of Use that are configured for the Application. The API returns the Oracle Terms of Use and the Partner Terms of Use if configured. The end user need to accept both the Oracle Terms of Use and Partner Terms of use to initiate the Application install.
func (client ApplicationInstanceResourceClient) GetTermsOfUses(ctx context.Context, request GetTermsOfUsesRequest) (response GetTermsOfUsesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTermsOfUses, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetTermsOfUsesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTermsOfUsesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTermsOfUsesResponse")
	}
	return
}

// getTermsOfUses implements the OCIOperation interface (enables retrying operations)
func (client ApplicationInstanceResourceClient) getTermsOfUses(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applications/{applicationId}/terms")
	if err != nil {
		return nil, err
	}

	var response GetTermsOfUsesResponse
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
