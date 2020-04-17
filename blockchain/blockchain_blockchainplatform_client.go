// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Plane API
//

package blockchain

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//BlockchainPlatformClient a client for BlockchainPlatform
type BlockchainPlatformClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewBlockchainPlatformClientWithConfigurationProvider Creates a new default BlockchainPlatform client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewBlockchainPlatformClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client BlockchainPlatformClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newBlockchainPlatformClientFromBaseClient(baseClient, configProvider)
}

// NewBlockchainPlatformClientWithOboToken Creates a new default BlockchainPlatform client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewBlockchainPlatformClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client BlockchainPlatformClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newBlockchainPlatformClientFromBaseClient(baseClient, configProvider)
}

func newBlockchainPlatformClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client BlockchainPlatformClient, err error) {
	client = BlockchainPlatformClient{BaseClient: baseClient}
	client.BasePath = "20191010"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *BlockchainPlatformClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("blockchain", "https://blockchain.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *BlockchainPlatformClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *BlockchainPlatformClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeBlockchainPlatformCompartment Change Blockchain Platform Compartment
func (client BlockchainPlatformClient) ChangeBlockchainPlatformCompartment(ctx context.Context, request ChangeBlockchainPlatformCompartmentRequest) (response ChangeBlockchainPlatformCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeBlockchainPlatformCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeBlockchainPlatformCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeBlockchainPlatformCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeBlockchainPlatformCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeBlockchainPlatformCompartmentResponse")
	}
	return
}

// changeBlockchainPlatformCompartment implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) changeBlockchainPlatformCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blockchainPlatforms/{blockchainPlatformId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeBlockchainPlatformCompartmentResponse
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

// CreateBlockchainPlatform Creates a new Blockchain Platform.
func (client BlockchainPlatformClient) CreateBlockchainPlatform(ctx context.Context, request CreateBlockchainPlatformRequest) (response CreateBlockchainPlatformResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBlockchainPlatformResponse")
	}
	return
}

// createBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) createBlockchainPlatform(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blockchainPlatforms")
	if err != nil {
		return nil, err
	}

	var response CreateBlockchainPlatformResponse
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

// DeleteBlockchainPlatform Delete a particular of a Blockchain Platform
func (client BlockchainPlatformClient) DeleteBlockchainPlatform(ctx context.Context, request DeleteBlockchainPlatformRequest) (response DeleteBlockchainPlatformResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBlockchainPlatformResponse")
	}
	return
}

// deleteBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) deleteBlockchainPlatform(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/blockchainPlatforms/{blockchainPlatformId}")
	if err != nil {
		return nil, err
	}

	var response DeleteBlockchainPlatformResponse
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

// GetBlockchainPlatform Gets information about a Blockchain Platform identified by the specific id
func (client BlockchainPlatformClient) GetBlockchainPlatform(ctx context.Context, request GetBlockchainPlatformRequest) (response GetBlockchainPlatformResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBlockchainPlatformResponse")
	}
	return
}

// getBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) getBlockchainPlatform(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/blockchainPlatforms/{blockchainPlatformId}")
	if err != nil {
		return nil, err
	}

	var response GetBlockchainPlatformResponse
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

// GetWorkRequest Gets the status of the work request with the given ID.
func (client BlockchainPlatformClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) getWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}")
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
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

// ListBlockchainPlatforms Returns a list Blockchain Platform Instances in a compartment
func (client BlockchainPlatformClient) ListBlockchainPlatforms(ctx context.Context, request ListBlockchainPlatformsRequest) (response ListBlockchainPlatformsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBlockchainPlatforms, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBlockchainPlatformsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBlockchainPlatformsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBlockchainPlatformsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBlockchainPlatformsResponse")
	}
	return
}

// listBlockchainPlatforms implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) listBlockchainPlatforms(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/blockchainPlatforms")
	if err != nil {
		return nil, err
	}

	var response ListBlockchainPlatformsResponse
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

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
func (client BlockchainPlatformClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
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

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
func (client BlockchainPlatformClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
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

// ListWorkRequests Lists the work requests in a compartment.
func (client BlockchainPlatformClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) listWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
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

// ScaleBlockchainPlatform Scale Blockchain Platform
func (client BlockchainPlatformClient) ScaleBlockchainPlatform(ctx context.Context, request ScaleBlockchainPlatformRequest) (response ScaleBlockchainPlatformResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.scaleBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScaleBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScaleBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScaleBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScaleBlockchainPlatformResponse")
	}
	return
}

// scaleBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) scaleBlockchainPlatform(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blockchainPlatforms/{blockchainPlatformId}/actions/scale")
	if err != nil {
		return nil, err
	}

	var response ScaleBlockchainPlatformResponse
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

// StartBlockchainPlatform Start a Blockchain Platform
func (client BlockchainPlatformClient) StartBlockchainPlatform(ctx context.Context, request StartBlockchainPlatformRequest) (response StartBlockchainPlatformResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.startBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartBlockchainPlatformResponse")
	}
	return
}

// startBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) startBlockchainPlatform(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blockchainPlatforms/{blockchainPlatformId}/actions/start")
	if err != nil {
		return nil, err
	}

	var response StartBlockchainPlatformResponse
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

// StopBlockchainPlatform Stop a Blockchain Platform
func (client BlockchainPlatformClient) StopBlockchainPlatform(ctx context.Context, request StopBlockchainPlatformRequest) (response StopBlockchainPlatformResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.stopBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopBlockchainPlatformResponse")
	}
	return
}

// stopBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) stopBlockchainPlatform(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blockchainPlatforms/{blockchainPlatformId}/actions/stop")
	if err != nil {
		return nil, err
	}

	var response StopBlockchainPlatformResponse
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

// UpdateBlockchainPlatform Update a particular of a Blockchain Platform
func (client BlockchainPlatformClient) UpdateBlockchainPlatform(ctx context.Context, request UpdateBlockchainPlatformRequest) (response UpdateBlockchainPlatformResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBlockchainPlatformResponse")
	}
	return
}

// updateBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) updateBlockchainPlatform(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/blockchainPlatforms/{blockchainPlatformId}")
	if err != nil {
		return nil, err
	}

	var response UpdateBlockchainPlatformResponse
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
