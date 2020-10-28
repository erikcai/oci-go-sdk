// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v27/common"
	"github.com/oracle/oci-go-sdk/v27/common/auth"
	"net/http"
)

//RoverNodeGetRPTClient a client for RoverNodeGetRPT
type RoverNodeGetRPTClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewRoverNodeGetRPTClientWithConfigurationProvider Creates a new default RoverNodeGetRPT client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewRoverNodeGetRPTClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client RoverNodeGetRPTClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newRoverNodeGetRPTClientFromBaseClient(baseClient, provider)
}

// NewRoverNodeGetRPTClientWithOboToken Creates a new default RoverNodeGetRPT client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewRoverNodeGetRPTClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client RoverNodeGetRPTClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newRoverNodeGetRPTClientFromBaseClient(baseClient, configProvider)
}

func newRoverNodeGetRPTClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client RoverNodeGetRPTClient, err error) {
	client = RoverNodeGetRPTClient{BaseClient: baseClient}
	client.BasePath = "20180828"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *RoverNodeGetRPTClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("rover", "https://rover-cloud-service.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *RoverNodeGetRPTClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *RoverNodeGetRPTClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetRoverNodeRPT Get the resource principal token for a rover node
func (client RoverNodeGetRPTClient) GetRoverNodeRPT(ctx context.Context, request GetRoverNodeRPTRequest) (response GetRoverNodeRPTResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRoverNodeRPT, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRoverNodeRPTResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRoverNodeRPTResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRoverNodeRPTResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRoverNodeRPTResponse")
	}
	return
}

// getRoverNodeRPT implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeGetRPTClient) getRoverNodeRPT(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes/{roverNodeId}/getRPT")
	if err != nil {
		return nil, err
	}

	var response GetRoverNodeRPTResponse
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
