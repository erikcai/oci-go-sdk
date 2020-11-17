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
	"github.com/oracle/oci-go-sdk/v29/common"
	"github.com/oracle/oci-go-sdk/v29/common/auth"
	"net/http"
)

//RoverNodeClient a client for RoverNode
type RoverNodeClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewRoverNodeClientWithConfigurationProvider Creates a new default RoverNode client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewRoverNodeClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client RoverNodeClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newRoverNodeClientFromBaseClient(baseClient, provider)
}

// NewRoverNodeClientWithOboToken Creates a new default RoverNode client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewRoverNodeClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client RoverNodeClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newRoverNodeClientFromBaseClient(baseClient, configProvider)
}

func newRoverNodeClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client RoverNodeClient, err error) {
	client = RoverNodeClient{BaseClient: baseClient}
	client.BasePath = "20180828"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *RoverNodeClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("rover", "https://rover.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *RoverNodeClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *RoverNodeClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeRoverNodeCompartment Moves a cluster into a different compartment.
func (client RoverNodeClient) ChangeRoverNodeCompartment(ctx context.Context, request ChangeRoverNodeCompartmentRequest) (response ChangeRoverNodeCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeRoverNodeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeRoverNodeCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeRoverNodeCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeRoverNodeCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeRoverNodeCompartmentResponse")
	}
	return
}

// changeRoverNodeCompartment implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) changeRoverNodeCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes/{roverNodeId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeRoverNodeCompartmentResponse
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

// CreateRoverNode Creates a new RoverNode.
func (client RoverNodeClient) CreateRoverNode(ctx context.Context, request CreateRoverNodeRequest) (response CreateRoverNodeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createRoverNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRoverNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRoverNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRoverNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRoverNodeResponse")
	}
	return
}

// createRoverNode implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) createRoverNode(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes")
	if err != nil {
		return nil, err
	}

	var response CreateRoverNodeResponse
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

// DeleteRoverNode Deletes a RoverNode resource by identifier
func (client RoverNodeClient) DeleteRoverNode(ctx context.Context, request DeleteRoverNodeRequest) (response DeleteRoverNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteRoverNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRoverNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRoverNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRoverNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRoverNodeResponse")
	}
	return
}

// deleteRoverNode implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) deleteRoverNode(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/roverNodes/{roverNodeId}")
	if err != nil {
		return nil, err
	}

	var response DeleteRoverNodeResponse
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

// GetNodeCertificate Get the certificate for a rover node
func (client RoverNodeClient) GetNodeCertificate(ctx context.Context, request GetNodeCertificateRequest) (response GetNodeCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNodeCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNodeCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNodeCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNodeCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNodeCertificateResponse")
	}
	return
}

// getNodeCertificate implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) getNodeCertificate(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes/{roverNodeId}/certificate")
	if err != nil {
		return nil, err
	}

	var response GetNodeCertificateResponse
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

// GetRoverNode Gets a RoverNode by identifier
func (client RoverNodeClient) GetRoverNode(ctx context.Context, request GetRoverNodeRequest) (response GetRoverNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRoverNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRoverNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRoverNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRoverNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRoverNodeResponse")
	}
	return
}

// getRoverNode implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) getRoverNode(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes/{roverNodeId}")
	if err != nil {
		return nil, err
	}

	var response GetRoverNodeResponse
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

// GetRoverNodeRPT Get the resource principal token for a rover node
func (client RoverNodeClient) GetRoverNodeRPT(ctx context.Context, request GetRoverNodeRPTRequest) (response GetRoverNodeRPTResponse, err error) {
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
func (client RoverNodeClient) getRoverNodeRPT(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// ListRoverNodes Returns a list of RoverNodes.
func (client RoverNodeClient) ListRoverNodes(ctx context.Context, request ListRoverNodesRequest) (response ListRoverNodesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRoverNodes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRoverNodesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRoverNodesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRoverNodesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRoverNodesResponse")
	}
	return
}

// listRoverNodes implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) listRoverNodes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes")
	if err != nil {
		return nil, err
	}

	var response ListRoverNodesResponse
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

// UpdateRoverNode Updates the RoverNode
func (client RoverNodeClient) UpdateRoverNode(ctx context.Context, request UpdateRoverNodeRequest) (response UpdateRoverNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRoverNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRoverNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRoverNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRoverNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRoverNodeResponse")
	}
	return
}

// updateRoverNode implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) updateRoverNode(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/roverNodes/{roverNodeId}")
	if err != nil {
		return nil, err
	}

	var response UpdateRoverNodeResponse
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
