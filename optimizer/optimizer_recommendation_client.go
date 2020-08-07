// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//RecommendationClient a client for Recommendation
type RecommendationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewRecommendationClientWithConfigurationProvider Creates a new default Recommendation client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewRecommendationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client RecommendationClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newRecommendationClientFromBaseClient(baseClient, configProvider)
}

// NewRecommendationClientWithOboToken Creates a new default Recommendation client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewRecommendationClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client RecommendationClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newRecommendationClientFromBaseClient(baseClient, configProvider)
}

func newRecommendationClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client RecommendationClient, err error) {
	client = RecommendationClient{BaseClient: baseClient}
	client.BasePath = "20200606"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *RecommendationClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("optimizer", "https://optimizer-api.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *RecommendationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *RecommendationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// BulkApplyRecommendations Applies the specified recommendations to the resources.
func (client RecommendationClient) BulkApplyRecommendations(ctx context.Context, request BulkApplyRecommendationsRequest) (response BulkApplyRecommendationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.bulkApplyRecommendations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkApplyRecommendationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkApplyRecommendationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkApplyRecommendationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkApplyRecommendationsResponse")
	}
	return
}

// bulkApplyRecommendations implements the OCIOperation interface (enables retrying operations)
func (client RecommendationClient) bulkApplyRecommendations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/recommendations/{recommendationId}/actions/bulkApplyRecommendations")
	if err != nil {
		return nil, err
	}

	var response BulkApplyRecommendationsResponse
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

// GetRecommendation Gets the recommendation for the specified OCID.
func (client RecommendationClient) GetRecommendation(ctx context.Context, request GetRecommendationRequest) (response GetRecommendationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRecommendation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRecommendationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRecommendationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRecommendationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRecommendationResponse")
	}
	return
}

// getRecommendation implements the OCIOperation interface (enables retrying operations)
func (client RecommendationClient) getRecommendation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/recommendations/{recommendationId}")
	if err != nil {
		return nil, err
	}

	var response GetRecommendationResponse
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

// ListRecommendations Lists the Cloud Advisor recommendations that are currently supported in the specified category.
func (client RecommendationClient) ListRecommendations(ctx context.Context, request ListRecommendationsRequest) (response ListRecommendationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRecommendations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRecommendationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRecommendationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRecommendationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRecommendationsResponse")
	}
	return
}

// listRecommendations implements the OCIOperation interface (enables retrying operations)
func (client RecommendationClient) listRecommendations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/recommendations")
	if err != nil {
		return nil, err
	}

	var response ListRecommendationsResponse
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

// UpdateRecommendation Updates the recommendation that corresponds to the specified OCID.
// Use this operation to implement the following actions:
//   * Postpone recommendation
//   * Dismiss recommendation
//   * Reactivate recommendation
func (client RecommendationClient) UpdateRecommendation(ctx context.Context, request UpdateRecommendationRequest) (response UpdateRecommendationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRecommendation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRecommendationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRecommendationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRecommendationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRecommendationResponse")
	}
	return
}

// updateRecommendation implements the OCIOperation interface (enables retrying operations)
func (client RecommendationClient) updateRecommendation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/recommendations/{recommendationId}")
	if err != nil {
		return nil, err
	}

	var response UpdateRecommendationResponse
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
