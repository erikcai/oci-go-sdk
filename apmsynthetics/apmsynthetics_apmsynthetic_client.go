// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Synthetic
//
// API for APM Synthetic service. Use this API to query Synthethetic scripts and monitors.
//

package apmsynthetics

import (
	"context"
	"fmt"
	"github.com/erikcai/oci-go-sdk/v33/common"
	"github.com/erikcai/oci-go-sdk/v33/common/auth"
	"net/http"
)

//ApmSyntheticClient a client for ApmSynthetic
type ApmSyntheticClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewApmSyntheticClientWithConfigurationProvider Creates a new default ApmSynthetic client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewApmSyntheticClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ApmSyntheticClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newApmSyntheticClientFromBaseClient(baseClient, provider)
}

// NewApmSyntheticClientWithOboToken Creates a new default ApmSynthetic client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewApmSyntheticClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ApmSyntheticClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newApmSyntheticClientFromBaseClient(baseClient, configProvider)
}

func newApmSyntheticClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ApmSyntheticClient, err error) {
	client = ApmSyntheticClient{BaseClient: baseClient}
	client.BasePath = "20200630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ApmSyntheticClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apmsynthetics", "https://apm-synthetic.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ApmSyntheticClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ApmSyntheticClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateMonitor Creates a new Monitor.
func (client ApmSyntheticClient) CreateMonitor(ctx context.Context, request CreateMonitorRequest) (response CreateMonitorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMonitorResponse")
	}
	return
}

// createMonitor implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) createMonitor(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitors")
	if err != nil {
		return nil, err
	}

	var response CreateMonitorResponse
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

// CreateScript Creates a new Script.
func (client ApmSyntheticClient) CreateScript(ctx context.Context, request CreateScriptRequest) (response CreateScriptResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateScriptResponse")
	}
	return
}

// createScript implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) createScript(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/scripts")
	if err != nil {
		return nil, err
	}

	var response CreateScriptResponse
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

// DeleteMonitor Deletes the specified monitor
func (client ApmSyntheticClient) DeleteMonitor(ctx context.Context, request DeleteMonitorRequest) (response DeleteMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMonitorResponse")
	}
	return
}

// deleteMonitor implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) deleteMonitor(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/monitors/{monitorId}")
	if err != nil {
		return nil, err
	}

	var response DeleteMonitorResponse
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

// DeleteScript Deletes the specified script.
func (client ApmSyntheticClient) DeleteScript(ctx context.Context, request DeleteScriptRequest) (response DeleteScriptResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteScriptResponse")
	}
	return
}

// deleteScript implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) deleteScript(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/scripts/{scriptId}")
	if err != nil {
		return nil, err
	}

	var response DeleteScriptResponse
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

// GetMonitor Get the configuration of the monitor identified by the OCID.
func (client ApmSyntheticClient) GetMonitor(ctx context.Context, request GetMonitorRequest) (response GetMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMonitorResponse")
	}
	return
}

// getMonitor implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) getMonitor(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/monitors/{monitorId}")
	if err != nil {
		return nil, err
	}

	var response GetMonitorResponse
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

// GetMonitorResult Get result(har or screenshot or consolelog) identified by the execution time of the monitor identified by the OCID.
func (client ApmSyntheticClient) GetMonitorResult(ctx context.Context, request GetMonitorResultRequest) (response GetMonitorResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMonitorResult, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMonitorResultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMonitorResultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMonitorResultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMonitorResultResponse")
	}
	return
}

// getMonitorResult implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) getMonitorResult(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/monitors/{monitorId}/results/{executionTime}")
	if err != nil {
		return nil, err
	}

	var response GetMonitorResultResponse
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

// GetScript Get the configuration of the Script identified by the OCID.
func (client ApmSyntheticClient) GetScript(ctx context.Context, request GetScriptRequest) (response GetScriptResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetScriptResponse")
	}
	return
}

// getScript implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) getScript(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scripts/{scriptId}")
	if err != nil {
		return nil, err
	}

	var response GetScriptResponse
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

// ListMonitors Returns a list of Monitor.
func (client ApmSyntheticClient) ListMonitors(ctx context.Context, request ListMonitorsRequest) (response ListMonitorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMonitors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMonitorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMonitorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMonitorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMonitorsResponse")
	}
	return
}

// listMonitors implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) listMonitors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/monitors")
	if err != nil {
		return nil, err
	}

	var response ListMonitorsResponse
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

// ListPublicVantagePoints Returns a list of Public Vantage Point.
func (client ApmSyntheticClient) ListPublicVantagePoints(ctx context.Context, request ListPublicVantagePointsRequest) (response ListPublicVantagePointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPublicVantagePoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPublicVantagePointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPublicVantagePointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPublicVantagePointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPublicVantagePointsResponse")
	}
	return
}

// listPublicVantagePoints implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) listPublicVantagePoints(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/publicVantagePoints")
	if err != nil {
		return nil, err
	}

	var response ListPublicVantagePointsResponse
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

// ListScripts Returns a list of script resources.
func (client ApmSyntheticClient) ListScripts(ctx context.Context, request ListScriptsRequest) (response ListScriptsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listScripts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListScriptsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListScriptsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListScriptsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListScriptsResponse")
	}
	return
}

// listScripts implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) listScripts(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scripts")
	if err != nil {
		return nil, err
	}

	var response ListScriptsResponse
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

// UpdateMonitor Updates the Monitor.
func (client ApmSyntheticClient) UpdateMonitor(ctx context.Context, request UpdateMonitorRequest) (response UpdateMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMonitorResponse")
	}
	return
}

// updateMonitor implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) updateMonitor(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/monitors/{monitorId}")
	if err != nil {
		return nil, err
	}

	var response UpdateMonitorResponse
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

// UpdateScript Updates the Script.
func (client ApmSyntheticClient) UpdateScript(ctx context.Context, request UpdateScriptRequest) (response UpdateScriptResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateScriptResponse")
	}
	return
}

// updateScript implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) updateScript(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/scripts/{scriptId}")
	if err != nil {
		return nil, err
	}

	var response UpdateScriptResponse
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
