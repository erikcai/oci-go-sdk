// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//MysqlaasClient a client for Mysqlaas
type MysqlaasClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMysqlaasClientWithConfigurationProvider Creates a new default Mysqlaas client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMysqlaasClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MysqlaasClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newMysqlaasClientFromBaseClient(baseClient, configProvider)
}

// NewMysqlaasClientWithOboToken Creates a new default Mysqlaas client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewMysqlaasClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MysqlaasClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newMysqlaasClientFromBaseClient(baseClient, configProvider)
}

func newMysqlaasClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MysqlaasClient, err error) {
	client = MysqlaasClient{BaseClient: baseClient}
	client.BasePath = "20190415"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MysqlaasClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("mysql", "https://mysql.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MysqlaasClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *MysqlaasClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CopyConfiguration Copy an existing MySQLaaS Configuration under a new name.
// No MySQLaaS Instances are associated with this new Configuration.
func (client MysqlaasClient) CopyConfiguration(ctx context.Context, request CopyConfigurationRequest) (response CopyConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.copyConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = CopyConfigurationResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CopyConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CopyConfigurationResponse")
	}
	return
}

// copyConfiguration implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) copyConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/configurations/{configurationId}/actions/copy")
	if err != nil {
		return nil, err
	}

	var response CopyConfigurationResponse
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

// CreateConfiguration Creates a new MySQLaaS Configuration.
func (client MysqlaasClient) CreateConfiguration(ctx context.Context, request CreateConfigurationRequest) (response CreateConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = CreateConfigurationResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConfigurationResponse")
	}
	return
}

// createConfiguration implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) createConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/configurations")
	if err != nil {
		return nil, err
	}

	var response CreateConfigurationResponse
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

// DeleteConfiguration Deletes a MySQLaaS Configuration.
// The Configuration must not be in use by any MySQLaaS Instances.
func (client MysqlaasClient) DeleteConfiguration(ctx context.Context, request DeleteConfigurationRequest) (response DeleteConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = DeleteConfigurationResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConfigurationResponse")
	}
	return
}

// deleteConfiguration implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) deleteConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/configurations/{configurationId}")
	if err != nil {
		return nil, err
	}

	var response DeleteConfigurationResponse
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

// DeleteInstance DEPRECATED -- this operation will be only available on DbSystem once implementation catches up.
// Delete a MySQLaaS instance, including terminating, detaching,
// removing, finalizing and otherwise deleting all related resources.
func (client MysqlaasClient) DeleteInstance(ctx context.Context, request DeleteInstanceRequest) (response DeleteInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteInstance, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = DeleteInstanceResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteInstanceResponse")
	}
	return
}

// deleteInstance implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) deleteInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/instances/{instanceId}")
	if err != nil {
		return nil, err
	}

	var response DeleteInstanceResponse
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

// EstimateAnalyticsClusterMemory Post a request to estimate the memory footprints of user tables when loaded to Analytics Cluster memory.
func (client MysqlaasClient) EstimateAnalyticsClusterMemory(ctx context.Context, request EstimateAnalyticsClusterMemoryRequest) (response EstimateAnalyticsClusterMemoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.estimateAnalyticsClusterMemory, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = EstimateAnalyticsClusterMemoryResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EstimateAnalyticsClusterMemoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EstimateAnalyticsClusterMemoryResponse")
	}
	return
}

// estimateAnalyticsClusterMemory implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) estimateAnalyticsClusterMemory(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instances/{instanceId}/analyticsClusterMemoryEstimate")
	if err != nil {
		return nil, err
	}

	var response EstimateAnalyticsClusterMemoryResponse
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

// GetAnalyticsClusterMemoryEstimate Get the most recent Analytics Cluster memory estimate that can be used to provision the Analytics Cluster size.
func (client MysqlaasClient) GetAnalyticsClusterMemoryEstimate(ctx context.Context, request GetAnalyticsClusterMemoryEstimateRequest) (response GetAnalyticsClusterMemoryEstimateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAnalyticsClusterMemoryEstimate, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = GetAnalyticsClusterMemoryEstimateResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAnalyticsClusterMemoryEstimateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAnalyticsClusterMemoryEstimateResponse")
	}
	return
}

// getAnalyticsClusterMemoryEstimate implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) getAnalyticsClusterMemoryEstimate(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instances/{instanceId}/analyticsClusterMemoryEstimate")
	if err != nil {
		return nil, err
	}

	var response GetAnalyticsClusterMemoryEstimateResponse
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

// GetConfiguration Get the full details of the specified Configuration, including the list of MySQL Options and their values.
func (client MysqlaasClient) GetConfiguration(ctx context.Context, request GetConfigurationRequest) (response GetConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = GetConfigurationResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConfigurationResponse")
	}
	return
}

// getConfiguration implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) getConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configurations/{configurationId}")
	if err != nil {
		return nil, err
	}

	var response GetConfigurationResponse
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

// GetConfigurationRevision Get a specific MySQLaaS Configuration Revision.
func (client MysqlaasClient) GetConfigurationRevision(ctx context.Context, request GetConfigurationRevisionRequest) (response GetConfigurationRevisionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConfigurationRevision, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = GetConfigurationRevisionResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConfigurationRevisionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConfigurationRevisionResponse")
	}
	return
}

// getConfigurationRevision implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) getConfigurationRevision(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configurations/{configurationId}/revisions/{revisionId}")
	if err != nil {
		return nil, err
	}

	var response GetConfigurationRevisionResponse
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

// GetInstance Get information about the specified MySQLaaS Instance.
func (client MysqlaasClient) GetInstance(ctx context.Context, request GetInstanceRequest) (response GetInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInstance, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = GetInstanceResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInstanceResponse")
	}
	return
}

// getInstance implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) getInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instances/{instanceId}")
	if err != nil {
		return nil, err
	}

	var response GetInstanceResponse
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

// GetVersionOptionsMetadata List the options that are valid for the specified MySQL version, including
// the datatype and range of values allowed.
func (client MysqlaasClient) GetVersionOptionsMetadata(ctx context.Context, request GetVersionOptionsMetadataRequest) (response GetVersionOptionsMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVersionOptionsMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = GetVersionOptionsMetadataResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVersionOptionsMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVersionOptionsMetadataResponse")
	}
	return
}

// getVersionOptionsMetadata implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) getVersionOptionsMetadata(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/versions/{versionId}/options")
	if err != nil {
		return nil, err
	}

	var response GetVersionOptionsMetadataResponse
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

// ListConfigurationRevisions Get the MySQLaaS Configuration Revision history.
func (client MysqlaasClient) ListConfigurationRevisions(ctx context.Context, request ListConfigurationRevisionsRequest) (response ListConfigurationRevisionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConfigurationRevisions, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = ListConfigurationRevisionsResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConfigurationRevisionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConfigurationRevisionsResponse")
	}
	return
}

// listConfigurationRevisions implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) listConfigurationRevisions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configurations/{configurationId}/revisions")
	if err != nil {
		return nil, err
	}

	var response ListConfigurationRevisionsResponse
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

// ListConfigurations Lists the Configurations you can use when deploying a MySQLaaS instance.
// This may include DEFAULT configurations per MySQL Shape and CUSTOM configurations.
// The default sort order is a multi-part sort by:
//   - shapeName, ascending
//   - DEFAULT-before-CUSTOM
//   - displayName ascending
func (client MysqlaasClient) ListConfigurations(ctx context.Context, request ListConfigurationsRequest) (response ListConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = ListConfigurationsResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConfigurationsResponse")
	}
	return
}

// listConfigurations implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) listConfigurations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configurations")
	if err != nil {
		return nil, err
	}

	var response ListConfigurationsResponse
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

// ListInstances Get a list of MySQLaaS Instances in the specified compartment. The default sort order is by timeUpdated, descending.
func (client MysqlaasClient) ListInstances(ctx context.Context, request ListInstancesRequest) (response ListInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInstances, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = ListInstancesResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInstancesResponse")
	}
	return
}

// listInstances implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) listInstances(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instances")
	if err != nil {
		return nil, err
	}

	var response ListInstancesResponse
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

// ListShapes Gets a list of the shapes you can use to create a new MySQLaaS
// Instance. The shape determines the resources allocated to the MySQLaaS
// instance - CPU cores and memory for VM shapes; CPU cores, memory and
// storage for non-VM (or bare metal) shapes.
func (client MysqlaasClient) ListShapes(ctx context.Context, request ListShapesRequest) (response ListShapesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listShapes, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = ListShapesResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListShapesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListShapesResponse")
	}
	return
}

// listShapes implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) listShapes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/shapes")
	if err != nil {
		return nil, err
	}

	var response ListShapesResponse
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

// ListVersions Get a list of supported and available MySQL database major versions.
// The list is sorted by version family.
func (client MysqlaasClient) ListVersions(ctx context.Context, request ListVersionsRequest) (response ListVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVersions, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = ListVersionsResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVersionsResponse")
	}
	return
}

// listVersions implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) listVersions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/versions")
	if err != nil {
		return nil, err
	}

	var response ListVersionsResponse
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

// RestartInstance Restart the specified MySQLaaS instance.
func (client MysqlaasClient) RestartInstance(ctx context.Context, request RestartInstanceRequest) (response RestartInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.restartInstance, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = RestartInstanceResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestartInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestartInstanceResponse")
	}
	return
}

// restartInstance implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) restartInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instances/{instanceId}/actions/restart")
	if err != nil {
		return nil, err
	}

	var response RestartInstanceResponse
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

// StartInstance DEPRECATED -- this operation will be only available on DbSystem once implementation catches up.
// Post a start action to be taken on the selected MySQLaaS instance.
func (client MysqlaasClient) StartInstance(ctx context.Context, request StartInstanceRequest) (response StartInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.startInstance, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = StartInstanceResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartInstanceResponse")
	}
	return
}

// startInstance implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) startInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instances/{instanceId}/actions/start")
	if err != nil {
		return nil, err
	}

	var response StartInstanceResponse
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

// StopInstance DEPRECATED -- this operation will be only available on DbSystem once implementation catches up.
// Stop the specified MySQLaaS instance.
// A stopped MySQLaaS Instance is not billed.
func (client MysqlaasClient) StopInstance(ctx context.Context, request StopInstanceRequest) (response StopInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.stopInstance, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = StopInstanceResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopInstanceResponse")
	}
	return
}

// stopInstance implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) stopInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instances/{instanceId}/actions/stop")
	if err != nil {
		return nil, err
	}

	var response StopInstanceResponse
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

// UpdateConfiguration Updates the MySQLaaS Configuration details.
func (client MysqlaasClient) UpdateConfiguration(ctx context.Context, request UpdateConfigurationRequest) (response UpdateConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = UpdateConfigurationResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConfigurationResponse")
	}
	return
}

// updateConfiguration implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) updateConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/configurations/{configurationId}")
	if err != nil {
		return nil, err
	}

	var response UpdateConfigurationResponse
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

// UpdateInstance Updating different fields in the MySQLaaS Instance will have a variety
// of different semantics with respect to the timeliness of the
// application of the change. For instance, changing the displayName of
// an instance will take effect immediately, but changing the shape of an
// instance is an asynchronous operation that involves provisioning new
// Compute resources, pausing the MySQLaaS instance and migrating storage
// before making the instance available again.
// Please see the (FIXME: link) documentation for more details on what
// changing the various fields implies.
func (client MysqlaasClient) UpdateInstance(ctx context.Context, request UpdateInstanceRequest) (response UpdateInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateInstance, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = UpdateInstanceResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateInstanceResponse")
	}
	return
}

// updateInstance implements the OCIOperation interface (enables retrying operations)
func (client MysqlaasClient) updateInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/instances/{instanceId}")
	if err != nil {
		return nil, err
	}

	var response UpdateInstanceResponse
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
