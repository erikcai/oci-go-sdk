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

//AnalyticsClient a client for Analytics
type AnalyticsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAnalyticsClientWithConfigurationProvider Creates a new default Analytics client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAnalyticsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AnalyticsClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newAnalyticsClientFromBaseClient(baseClient, configProvider)
}

// NewAnalyticsClientWithOboToken Creates a new default Analytics client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewAnalyticsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AnalyticsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newAnalyticsClientFromBaseClient(baseClient, configProvider)
}

func newAnalyticsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AnalyticsClient, err error) {
	client = AnalyticsClient{BaseClient: baseClient}
	client.BasePath = "20190415"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AnalyticsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("mysql", "https://mysql.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AnalyticsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AnalyticsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateAnalyticsCluster Create and launch a new Analytics Cluster and associate it with an existing MySQLaaS instance.
func (client AnalyticsClient) CreateAnalyticsCluster(ctx context.Context, request CreateAnalyticsClusterRequest) (response CreateAnalyticsClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAnalyticsCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAnalyticsClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAnalyticsClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAnalyticsClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAnalyticsClusterResponse")
	}
	return
}

// createAnalyticsCluster implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) createAnalyticsCluster(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/analyticsClusters")
	if err != nil {
		return nil, err
	}

	var response CreateAnalyticsClusterResponse
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

// DeleteAnalyticsCluster Delete a MySQLaaS Analytics cluster.
func (client AnalyticsClient) DeleteAnalyticsCluster(ctx context.Context, request DeleteAnalyticsClusterRequest) (response DeleteAnalyticsClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAnalyticsCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAnalyticsClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAnalyticsClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAnalyticsClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAnalyticsClusterResponse")
	}
	return
}

// deleteAnalyticsCluster implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) deleteAnalyticsCluster(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/analyticsClusters/{analyticsClusterId}")
	if err != nil {
		return nil, err
	}

	var response DeleteAnalyticsClusterResponse
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

// GetAnalyticsCluster Get information about the specified MySQLaaS Analytics Cluster.
func (client AnalyticsClient) GetAnalyticsCluster(ctx context.Context, request GetAnalyticsClusterRequest) (response GetAnalyticsClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAnalyticsCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAnalyticsClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAnalyticsClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAnalyticsClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAnalyticsClusterResponse")
	}
	return
}

// getAnalyticsCluster implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) getAnalyticsCluster(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/analyticsClusters/{analyticsClusterId}")
	if err != nil {
		return nil, err
	}

	var response GetAnalyticsClusterResponse
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

// ListAnalyticsClusterShapes Gets a list of the shapes you can use to create a new MySQLaaS
// Analytics Cluster. The shape determines the resources allocated to the
// MySQLaaS Analytics Cluster.
func (client AnalyticsClient) ListAnalyticsClusterShapes(ctx context.Context, request ListAnalyticsClusterShapesRequest) (response ListAnalyticsClusterShapesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAnalyticsClusterShapes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAnalyticsClusterShapesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAnalyticsClusterShapesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAnalyticsClusterShapesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAnalyticsClusterShapesResponse")
	}
	return
}

// listAnalyticsClusterShapes implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) listAnalyticsClusterShapes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/analyticsClusterShapes")
	if err != nil {
		return nil, err
	}

	var response ListAnalyticsClusterShapesResponse
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

// ListAnalyticsClusters Get a list of Analytics Clusters in the specified compartment. The default sort order is by timeUpdated, descending.
func (client AnalyticsClient) ListAnalyticsClusters(ctx context.Context, request ListAnalyticsClustersRequest) (response ListAnalyticsClustersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAnalyticsClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAnalyticsClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAnalyticsClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAnalyticsClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAnalyticsClustersResponse")
	}
	return
}

// listAnalyticsClusters implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) listAnalyticsClusters(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/analyticsClusters")
	if err != nil {
		return nil, err
	}

	var response ListAnalyticsClustersResponse
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

// StartAnalyticsCluster Start the specified MySQLaaS Analytics Cluster
func (client AnalyticsClient) StartAnalyticsCluster(ctx context.Context, request StartAnalyticsClusterRequest) (response StartAnalyticsClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.startAnalyticsCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartAnalyticsClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartAnalyticsClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartAnalyticsClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartAnalyticsClusterResponse")
	}
	return
}

// startAnalyticsCluster implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) startAnalyticsCluster(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/analyticsClusters/{analyticsClusterId}/actions/start")
	if err != nil {
		return nil, err
	}

	var response StartAnalyticsClusterResponse
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

// StopAnalyticsCluster Stop the specified MySQLaaS Analytics Cluster.
func (client AnalyticsClient) StopAnalyticsCluster(ctx context.Context, request StopAnalyticsClusterRequest) (response StopAnalyticsClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.stopAnalyticsCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopAnalyticsClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopAnalyticsClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopAnalyticsClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopAnalyticsClusterResponse")
	}
	return
}

// stopAnalyticsCluster implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) stopAnalyticsCluster(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/analyticsClusters/{analyticsClusterId}/actions/stop")
	if err != nil {
		return nil, err
	}

	var response StopAnalyticsClusterResponse
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

// UpdateAnalyticsCluster Update a MySQLaaS Analytics Cluster.
// Updating different fields in the MySQLaaS Analytics Cluster will have
// a variety of different semantics with respect to the timeliness of the
// application of the change.
// Please see the (FIXME: link) documentation for more details on what
// changing the various fields implies.
func (client AnalyticsClient) UpdateAnalyticsCluster(ctx context.Context, request UpdateAnalyticsClusterRequest) (response UpdateAnalyticsClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAnalyticsCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAnalyticsClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAnalyticsClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAnalyticsClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAnalyticsClusterResponse")
	}
	return
}

// updateAnalyticsCluster implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) updateAnalyticsCluster(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/analyticsClusters/{analyticsClusterId}")
	if err != nil {
		return nil, err
	}

	var response UpdateAnalyticsClusterResponse
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
