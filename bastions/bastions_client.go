// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciBastions API
//
// A description of the OciBastions API
//

package bastions

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//BastionsClient a client for Bastions
type BastionsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewBastionsClientWithConfigurationProvider Creates a new default Bastions client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewBastionsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client BastionsClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newBastionsClientFromBaseClient(baseClient, configProvider)
}

// NewBastionsClientWithOboToken Creates a new default Bastions client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewBastionsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client BastionsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newBastionsClientFromBaseClient(baseClient, configProvider)
}

func newBastionsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client BastionsClient, err error) {
	client = BastionsClient{BaseClient: baseClient}
	client.BasePath = "20201015"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *BastionsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("bastions", "https://bastions.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *BastionsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *BastionsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeBastionCompartment Update bastion compartment
func (client BastionsClient) ChangeBastionCompartment(ctx context.Context, request ChangeBastionCompartmentRequest) (response ChangeBastionCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeBastionCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeBastionCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeBastionCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeBastionCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeBastionCompartmentResponse")
	}
	return
}

// changeBastionCompartment implements the OCIOperation interface (enables retrying operations)
func (client BastionsClient) changeBastionCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bastions/{bastionId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeBastionCompartmentResponse
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

// CreateBastion Creates a new bastion.
func (client BastionsClient) CreateBastion(ctx context.Context, request CreateBastionRequest) (response CreateBastionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createBastion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBastionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBastionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBastionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBastionResponse")
	}
	return
}

// createBastion implements the OCIOperation interface (enables retrying operations)
func (client BastionsClient) createBastion(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bastions")
	if err != nil {
		return nil, err
	}

	var response CreateBastionResponse
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

// DeleteBastion Deletes a bastion resource by identifier
func (client BastionsClient) DeleteBastion(ctx context.Context, request DeleteBastionRequest) (response DeleteBastionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBastion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBastionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBastionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBastionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBastionResponse")
	}
	return
}

// deleteBastion implements the OCIOperation interface (enables retrying operations)
func (client BastionsClient) deleteBastion(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/bastions/{bastionId}")
	if err != nil {
		return nil, err
	}

	var response DeleteBastionResponse
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

// GetBastion Gets a Bastion by identifier
func (client BastionsClient) GetBastion(ctx context.Context, request GetBastionRequest) (response GetBastionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBastion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBastionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBastionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBastionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBastionResponse")
	}
	return
}

// getBastion implements the OCIOperation interface (enables retrying operations)
func (client BastionsClient) getBastion(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bastions/{bastionId}")
	if err != nil {
		return nil, err
	}

	var response GetBastionResponse
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

// ListBastions Returns a list of bastions.
func (client BastionsClient) ListBastions(ctx context.Context, request ListBastionsRequest) (response ListBastionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBastions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBastionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBastionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBastionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBastionsResponse")
	}
	return
}

// listBastions implements the OCIOperation interface (enables retrying operations)
func (client BastionsClient) listBastions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bastions")
	if err != nil {
		return nil, err
	}

	var response ListBastionsResponse
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

// UpdateBastion Updates the bastion
func (client BastionsClient) UpdateBastion(ctx context.Context, request UpdateBastionRequest) (response UpdateBastionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBastion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBastionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBastionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBastionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBastionResponse")
	}
	return
}

// updateBastion implements the OCIOperation interface (enables retrying operations)
func (client BastionsClient) updateBastion(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/bastions/{bastionId}")
	if err != nil {
		return nil, err
	}

	var response UpdateBastionResponse
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
