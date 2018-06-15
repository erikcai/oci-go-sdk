// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Orchestration
//
// Orchestration Maestro API
//

package orchestration

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//OrchestrationClient a client for Orchestration
type OrchestrationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOrchestrationClientWithConfigurationProvider Creates a new default Orchestration client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOrchestrationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OrchestrationClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = OrchestrationClient{BaseClient: baseClient}
	client.BasePath = "20170630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OrchestrationClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "orchestration", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OrchestrationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OrchestrationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelJob Cancel the job
func (client OrchestrationClient) CancelJob(ctx context.Context, request CancelJobRequest) (response CancelJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelJob, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(CancelJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelJobResponse")
	}
	return
}

// cancelJob implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) cancelJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/jobs/{jobId}")
	if err != nil {
		return nil, err
	}

	var response CancelJobResponse
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

// CreateJob Creates a Job
func (client OrchestrationClient) CreateJob(ctx context.Context, request CreateJobRequest) (response CreateJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createJob, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJobResponse")
	}
	return
}

// createJob implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) createJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobs")
	if err != nil {
		return nil, err
	}

	var response CreateJobResponse
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

// CreateStack Creates a Stack
func (client OrchestrationClient) CreateStack(ctx context.Context, request CreateStackRequest) (response CreateStackResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createStack, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(CreateStackResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateStackResponse")
	}
	return
}

// createStack implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) createStack(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/stacks")
	if err != nil {
		return nil, err
	}

	var response CreateStackResponse
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

// DeleteStack Delete the Stack object
func (client OrchestrationClient) DeleteStack(ctx context.Context, request DeleteStackRequest) (response DeleteStackResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteStack, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteStackResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteStackResponse")
	}
	return
}

// deleteStack implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) deleteStack(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/stacks/{stackId}")
	if err != nil {
		return nil, err
	}

	var response DeleteStackResponse
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

// GetJob Returns a job
func (client OrchestrationClient) GetJob(ctx context.Context, request GetJobRequest) (response GetJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJob, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobResponse")
	}
	return
}

// getJob implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) getJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}")
	if err != nil {
		return nil, err
	}

	var response GetJobResponse
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

// GetJobLogs Returns a Job's logs
func (client OrchestrationClient) GetJobLogs(ctx context.Context, request GetJobLogsRequest) (response GetJobLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobLogs, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobLogsResponse")
	}
	return
}

// getJobLogs implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) getJobLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/logs")
	if err != nil {
		return nil, err
	}

	var response GetJobLogsResponse
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

// GetJobLogsContent Returns a Job's raw logs. If the logs contain more than 100000 entries, only 100000 are returned.
func (client OrchestrationClient) GetJobLogsContent(ctx context.Context, request GetJobLogsContentRequest) (response GetJobLogsContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobLogsContent, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobLogsContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobLogsContentResponse")
	}
	return
}

// getJobLogsContent implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) getJobLogsContent(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/logs/content")
	if err != nil {
		return nil, err
	}

	var response GetJobLogsContentResponse
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

// GetJobTfConfig Returns a job's Terraform configuration in a zip file
func (client OrchestrationClient) GetJobTfConfig(ctx context.Context, request GetJobTfConfigRequest) (response GetJobTfConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobTfConfig, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobTfConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobTfConfigResponse")
	}
	return
}

// getJobTfConfig implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) getJobTfConfig(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/tfConfig")
	if err != nil {
		return nil, err
	}

	var response GetJobTfConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJobTfExecutionPlan Returns a job's Terraform execution plan in a binary file
func (client OrchestrationClient) GetJobTfExecutionPlan(ctx context.Context, request GetJobTfExecutionPlanRequest) (response GetJobTfExecutionPlanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobTfExecutionPlan, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobTfExecutionPlanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobTfExecutionPlanResponse")
	}
	return
}

// getJobTfExecutionPlan implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) getJobTfExecutionPlan(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/tfExecutionPlan")
	if err != nil {
		return nil, err
	}

	var response GetJobTfExecutionPlanResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJobTfState Returns a Job's Terraform state
func (client OrchestrationClient) GetJobTfState(ctx context.Context, request GetJobTfStateRequest) (response GetJobTfStateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobTfState, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobTfStateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobTfStateResponse")
	}
	return
}

// getJobTfState implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) getJobTfState(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/tfState")
	if err != nil {
		return nil, err
	}

	var response GetJobTfStateResponse
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

// GetStack Gets a Stack
func (client OrchestrationClient) GetStack(ctx context.Context, request GetStackRequest) (response GetStackResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStack, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(GetStackResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStackResponse")
	}
	return
}

// getStack implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) getStack(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/stacks/{stackId}")
	if err != nil {
		return nil, err
	}

	var response GetStackResponse
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

// GetStackTfConfig Returns a stack's Terraform configuration in a zip file if one has been provided
func (client OrchestrationClient) GetStackTfConfig(ctx context.Context, request GetStackTfConfigRequest) (response GetStackTfConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStackTfConfig, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(GetStackTfConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStackTfConfigResponse")
	}
	return
}

// getStackTfConfig implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) getStackTfConfig(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/stacks/{stackId}/tfConfig")
	if err != nil {
		return nil, err
	}

	var response GetStackTfConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJobs Returns a list of jobs. Either stack or compartment must be provided to indicate which jobs to list.
func (client OrchestrationClient) ListJobs(ctx context.Context, request ListJobsRequest) (response ListJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobs, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobsResponse")
	}
	return
}

// listJobs implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) listJobs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs")
	if err != nil {
		return nil, err
	}

	var response ListJobsResponse
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

// ListStacks Returns a list of stacks
func (client OrchestrationClient) ListStacks(ctx context.Context, request ListStacksRequest) (response ListStacksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listStacks, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(ListStacksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListStacksResponse")
	}
	return
}

// listStacks implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) listStacks(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/stacks")
	if err != nil {
		return nil, err
	}

	var response ListStacksResponse
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

// UpdateJob Update the Job
func (client OrchestrationClient) UpdateJob(ctx context.Context, request UpdateJobRequest) (response UpdateJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateJob, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateJobResponse")
	}
	return
}

// updateJob implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) updateJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/jobs/{jobId}")
	if err != nil {
		return nil, err
	}

	var response UpdateJobResponse
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

// UpdateStack Update the Stack object
func (client OrchestrationClient) UpdateStack(ctx context.Context, request UpdateStackRequest) (response UpdateStackResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateStack, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateStackResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateStackResponse")
	}
	return
}

// updateStack implements the OCIOperation interface (enables retrying operations)
func (client OrchestrationClient) updateStack(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/stacks/{stackId}")
	if err != nil {
		return nil, err
	}

	var response UpdateStackResponse
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
