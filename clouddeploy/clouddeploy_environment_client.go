// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v30/common"
	"github.com/oracle/oci-go-sdk/v30/common/auth"
	"net/http"
)

//EnvironmentClient a client for Environment
type EnvironmentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewEnvironmentClientWithConfigurationProvider Creates a new default Environment client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewEnvironmentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client EnvironmentClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newEnvironmentClientFromBaseClient(baseClient, provider)
}

// NewEnvironmentClientWithOboToken Creates a new default Environment client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewEnvironmentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client EnvironmentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newEnvironmentClientFromBaseClient(baseClient, configProvider)
}

func newEnvironmentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client EnvironmentClient, err error) {
	client = EnvironmentClient{BaseClient: baseClient}
	client.BasePath = "20180828"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *EnvironmentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("clouddeploy", "https://cloud-deploy.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *EnvironmentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *EnvironmentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateEnvironment Creates a new Environment.
func (client EnvironmentClient) CreateEnvironment(ctx context.Context, request CreateEnvironmentRequest) (response CreateEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEnvironmentResponse")
	}
	return
}

// createEnvironment implements the OCIOperation interface (enables retrying operations)
func (client EnvironmentClient) createEnvironment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/environments")
	if err != nil {
		return nil, err
	}

	var response CreateEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &environment{})
	return response, err
}

// DeleteEnvironment Deletes a Environment resource by identifier
func (client EnvironmentClient) DeleteEnvironment(ctx context.Context, request DeleteEnvironmentRequest) (response DeleteEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEnvironmentResponse")
	}
	return
}

// deleteEnvironment implements the OCIOperation interface (enables retrying operations)
func (client EnvironmentClient) deleteEnvironment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/environments/{environmentId}")
	if err != nil {
		return nil, err
	}

	var response DeleteEnvironmentResponse
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

// GetEnvironment Gets a Environment by identifier
func (client EnvironmentClient) GetEnvironment(ctx context.Context, request GetEnvironmentRequest) (response GetEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEnvironmentResponse")
	}
	return
}

// getEnvironment implements the OCIOperation interface (enables retrying operations)
func (client EnvironmentClient) getEnvironment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/environments/{environmentId}")
	if err != nil {
		return nil, err
	}

	var response GetEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &environment{})
	return response, err
}

// ListEnvironments Returns a list of Environments.
func (client EnvironmentClient) ListEnvironments(ctx context.Context, request ListEnvironmentsRequest) (response ListEnvironmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEnvironments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEnvironmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEnvironmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEnvironmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEnvironmentsResponse")
	}
	return
}

// listEnvironments implements the OCIOperation interface (enables retrying operations)
func (client EnvironmentClient) listEnvironments(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/environments")
	if err != nil {
		return nil, err
	}

	var response ListEnvironmentsResponse
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

// UpdateEnvironment Updates the Environment
func (client EnvironmentClient) UpdateEnvironment(ctx context.Context, request UpdateEnvironmentRequest) (response UpdateEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEnvironmentResponse")
	}
	return
}

// updateEnvironment implements the OCIOperation interface (enables retrying operations)
func (client EnvironmentClient) updateEnvironment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/environments/{environmentId}")
	if err != nil {
		return nil, err
	}

	var response UpdateEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &environment{})
	return response, err
}
