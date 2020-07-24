// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Optimizer API
//
// The API for the OCI Optimizer
//

package optimizer

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//EnrollmentStatusClient a client for EnrollmentStatus
type EnrollmentStatusClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewEnrollmentStatusClientWithConfigurationProvider Creates a new default EnrollmentStatus client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewEnrollmentStatusClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client EnrollmentStatusClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newEnrollmentStatusClientFromBaseClient(baseClient, configProvider)
}

// NewEnrollmentStatusClientWithOboToken Creates a new default EnrollmentStatus client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewEnrollmentStatusClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client EnrollmentStatusClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newEnrollmentStatusClientFromBaseClient(baseClient, configProvider)
}

func newEnrollmentStatusClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client EnrollmentStatusClient, err error) {
	client = EnrollmentStatusClient{BaseClient: baseClient}
	client.BasePath = "20200606"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *EnrollmentStatusClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("optimizer", "https://optimizer-api.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *EnrollmentStatusClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *EnrollmentStatusClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetEnrollmentStatus Returns the optimizer enrollment status.
func (client EnrollmentStatusClient) GetEnrollmentStatus(ctx context.Context, request GetEnrollmentStatusRequest) (response GetEnrollmentStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEnrollmentStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEnrollmentStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEnrollmentStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEnrollmentStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEnrollmentStatusResponse")
	}
	return
}

// getEnrollmentStatus implements the OCIOperation interface (enables retrying operations)
func (client EnrollmentStatusClient) getEnrollmentStatus(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/enrollmentStatus/{enrollmentStatusId}")
	if err != nil {
		return nil, err
	}

	var response GetEnrollmentStatusResponse
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

// ListEnrollmentStatuses Returns the list of optimizer enrollment statuses.
func (client EnrollmentStatusClient) ListEnrollmentStatuses(ctx context.Context, request ListEnrollmentStatusesRequest) (response ListEnrollmentStatusesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEnrollmentStatuses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEnrollmentStatusesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEnrollmentStatusesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEnrollmentStatusesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEnrollmentStatusesResponse")
	}
	return
}

// listEnrollmentStatuses implements the OCIOperation interface (enables retrying operations)
func (client EnrollmentStatusClient) listEnrollmentStatuses(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/enrollmentStatus")
	if err != nil {
		return nil, err
	}

	var response ListEnrollmentStatusesResponse
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

// UpdateEnrollmentStatus Update the enrollment status of the tenancy.
func (client EnrollmentStatusClient) UpdateEnrollmentStatus(ctx context.Context, request UpdateEnrollmentStatusRequest) (response UpdateEnrollmentStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEnrollmentStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEnrollmentStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEnrollmentStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEnrollmentStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEnrollmentStatusResponse")
	}
	return
}

// updateEnrollmentStatus implements the OCIOperation interface (enables retrying operations)
func (client EnrollmentStatusClient) updateEnrollmentStatus(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/enrollmentStatus/{enrollmentStatusId}")
	if err != nil {
		return nil, err
	}

	var response UpdateEnrollmentStatusResponse
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
