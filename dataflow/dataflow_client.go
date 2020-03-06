// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//DataFlowClient a client for DataFlow
type DataFlowClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDataFlowClientWithConfigurationProvider Creates a new default DataFlow client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDataFlowClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DataFlowClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newDataFlowClientFromBaseClient(baseClient, configProvider)
}

// NewDataFlowClientWithOboToken Creates a new default DataFlow client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDataFlowClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DataFlowClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newDataFlowClientFromBaseClient(baseClient, configProvider)
}

func newDataFlowClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DataFlowClient, err error) {
	client = DataFlowClient{BaseClient: baseClient}
	client.BasePath = "20200129"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DataFlowClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dataflow", "https://dataflow.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DataFlowClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DataFlowClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateApplication Creates an application.
func (client DataFlowClient) CreateApplication(ctx context.Context, request CreateApplicationRequest) (response CreateApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createApplication, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateApplicationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateApplicationResponse")
	}
	return
}

// createApplication implements the OCIOperation interface (enables retrying operations)
func (client DataFlowClient) createApplication(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/applications")
	if err != nil {
		return nil, err
	}

	var response CreateApplicationResponse
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

// CreateRun Creates a run for an application.
func (client DataFlowClient) CreateRun(ctx context.Context, request CreateRunRequest) (response CreateRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createRun, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateRunResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRunResponse")
	}
	return
}

// createRun implements the OCIOperation interface (enables retrying operations)
func (client DataFlowClient) createRun(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/runs")
	if err != nil {
		return nil, err
	}

	var response CreateRunResponse
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

// DeleteApplication Deletes an application using an `applicationId`.
func (client DataFlowClient) DeleteApplication(ctx context.Context, request DeleteApplicationRequest) (response DeleteApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteApplication, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteApplicationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteApplicationResponse")
	}
	return
}

// deleteApplication implements the OCIOperation interface (enables retrying operations)
func (client DataFlowClient) deleteApplication(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/applications/{applicationId}")
	if err != nil {
		return nil, err
	}

	var response DeleteApplicationResponse
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

// DeleteRun Cancels the specified run if it has not already completed or was previously cancelled.
// If a run is in progress, the executing job will be killed.
func (client DataFlowClient) DeleteRun(ctx context.Context, request DeleteRunRequest) (response DeleteRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteRun, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteRunResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRunResponse")
	}
	return
}

// deleteRun implements the OCIOperation interface (enables retrying operations)
func (client DataFlowClient) deleteRun(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/runs/{runId}")
	if err != nil {
		return nil, err
	}

	var response DeleteRunResponse
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

// GetApplication Retrieves an application using an `applicationId`.
func (client DataFlowClient) GetApplication(ctx context.Context, request GetApplicationRequest) (response GetApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getApplication, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetApplicationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetApplicationResponse")
	}
	return
}

// getApplication implements the OCIOperation interface (enables retrying operations)
func (client DataFlowClient) getApplication(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applications/{applicationId}")
	if err != nil {
		return nil, err
	}

	var response GetApplicationResponse
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

// GetRun Retrieves the run for the specified `runId`.
func (client DataFlowClient) GetRun(ctx context.Context, request GetRunRequest) (response GetRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRun, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetRunResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRunResponse")
	}
	return
}

// getRun implements the OCIOperation interface (enables retrying operations)
func (client DataFlowClient) getRun(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/runs/{runId}")
	if err != nil {
		return nil, err
	}

	var response GetRunResponse
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

// GetRunLog Retrieves the content of an run log.
func (client DataFlowClient) GetRunLog(ctx context.Context, request GetRunLogRequest) (response GetRunLogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRunLog, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetRunLogResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRunLogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRunLogResponse")
	}
	return
}

// getRunLog implements the OCIOperation interface (enables retrying operations)
func (client DataFlowClient) getRunLog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/runs/{runId}/logs/{name}")
	if err != nil {
		return nil, err
	}

	var response GetRunLogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApplications Lists all applications in the specified compartment.
func (client DataFlowClient) ListApplications(ctx context.Context, request ListApplicationsRequest) (response ListApplicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listApplications, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListApplicationsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApplicationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApplicationsResponse")
	}
	return
}

// listApplications implements the OCIOperation interface (enables retrying operations)
func (client DataFlowClient) listApplications(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applications")
	if err != nil {
		return nil, err
	}

	var response ListApplicationsResponse
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

// ListRunLogs Retrieves summaries of the run's logs.
func (client DataFlowClient) ListRunLogs(ctx context.Context, request ListRunLogsRequest) (response ListRunLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRunLogs, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListRunLogsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRunLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRunLogsResponse")
	}
	return
}

// listRunLogs implements the OCIOperation interface (enables retrying operations)
func (client DataFlowClient) listRunLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/runs/{runId}/logs")
	if err != nil {
		return nil, err
	}

	var response ListRunLogsResponse
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

// ListRuns Lists all runs of an application in the specified compartment.
func (client DataFlowClient) ListRuns(ctx context.Context, request ListRunsRequest) (response ListRunsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRuns, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListRunsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRunsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRunsResponse")
	}
	return
}

// listRuns implements the OCIOperation interface (enables retrying operations)
func (client DataFlowClient) listRuns(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/runs")
	if err != nil {
		return nil, err
	}

	var response ListRunsResponse
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

// UpdateApplication Updates an application using an `applicationId`.
func (client DataFlowClient) UpdateApplication(ctx context.Context, request UpdateApplicationRequest) (response UpdateApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateApplication, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateApplicationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateApplicationResponse")
	}
	return
}

// updateApplication implements the OCIOperation interface (enables retrying operations)
func (client DataFlowClient) updateApplication(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/applications/{applicationId}")
	if err != nil {
		return nil, err
	}

	var response UpdateApplicationResponse
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

// UpdateRun Updates a run using a `runId`.
func (client DataFlowClient) UpdateRun(ctx context.Context, request UpdateRunRequest) (response UpdateRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRun, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateRunResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRunResponse")
	}
	return
}

// updateRun implements the OCIOperation interface (enables retrying operations)
func (client DataFlowClient) updateRun(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/runs/{runId}")
	if err != nil {
		return nil, err
	}

	var response UpdateRunResponse
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
