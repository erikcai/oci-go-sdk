// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
//
// Cloud Incident Management Service
//

package cims

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//IncidentClient a client for Incident
type IncidentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewIncidentClientWithConfigurationProvider Creates a new default Incident client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewIncidentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client IncidentClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = IncidentClient{BaseClient: baseClient}
	client.BasePath = "20181231"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *IncidentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("cims", "https://incidentmanagement.us-phoenix-1.oraclecloud.com/20181231")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *IncidentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *IncidentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateIncident CreateIncident
func (client IncidentClient) CreateIncident(ctx context.Context, request CreateIncidentRequest) (response CreateIncidentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createIncident, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateIncidentResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateIncidentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateIncidentResponse")
	}
	return
}

// createIncident implements the OCIOperation interface (enables retrying operations)
func (client IncidentClient) createIncident(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/v1/incidents")
	if err != nil {
		return nil, err
	}

	var response CreateIncidentResponse
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

// GetIncident GetIncident
func (client IncidentClient) GetIncident(ctx context.Context, request GetIncidentRequest) (response GetIncidentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getIncident, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetIncidentResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetIncidentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetIncidentResponse")
	}
	return
}

// getIncident implements the OCIOperation interface (enables retrying operations)
func (client IncidentClient) getIncident(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v1/incidents/{incidentId}")
	if err != nil {
		return nil, err
	}

	var response GetIncidentResponse
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

// GetStatus GetStatus
func (client IncidentClient) GetStatus(ctx context.Context, request GetStatusRequest) (response GetStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStatus, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetStatusResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStatusResponse")
	}
	return
}

// getStatus implements the OCIOperation interface (enables retrying operations)
func (client IncidentClient) getStatus(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v1/incidents/status/{source}")
	if err != nil {
		return nil, err
	}

	var response GetStatusResponse
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

// ListIncidents ListIncidents
func (client IncidentClient) ListIncidents(ctx context.Context, request ListIncidentsRequest) (response ListIncidentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIncidents, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListIncidentsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIncidentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIncidentsResponse")
	}
	return
}

// listIncidents implements the OCIOperation interface (enables retrying operations)
func (client IncidentClient) listIncidents(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v1/incidents")
	if err != nil {
		return nil, err
	}

	var response ListIncidentsResponse
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

// ListTypes ListTypes
func (client IncidentClient) ListTypes(ctx context.Context, request ListTypesRequest) (response ListTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTypes, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListTypesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTypesResponse")
	}
	return
}

// listTypes implements the OCIOperation interface (enables retrying operations)
func (client IncidentClient) listTypes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v1/incidents/types")
	if err != nil {
		return nil, err
	}

	var response ListTypesResponse
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

// UpdateIncident Update an existing incident
func (client IncidentClient) UpdateIncident(ctx context.Context, request UpdateIncidentRequest) (response UpdateIncidentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateIncident, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateIncidentResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateIncidentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateIncidentResponse")
	}
	return
}

// updateIncident implements the OCIOperation interface (enables retrying operations)
func (client IncidentClient) updateIncident(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/v1/incidents/{incidentId}")
	if err != nil {
		return nil, err
	}

	var response UpdateIncidentResponse
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

// ValidateUser ValidateUser
func (client IncidentClient) ValidateUser(ctx context.Context, request ValidateUserRequest) (response ValidateUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.validateUser, policy)
	if err != nil {
		if ociResponse != nil {
			response = ValidateUserResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateUserResponse")
	}
	return
}

// validateUser implements the OCIOperation interface (enables retrying operations)
func (client IncidentClient) validateUser(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v1/incidents/user/validate")
	if err != nil {
		return nil, err
	}

	var response ValidateUserResponse
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
