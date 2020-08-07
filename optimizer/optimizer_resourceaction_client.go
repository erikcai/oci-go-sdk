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

//ResourceActionClient a client for ResourceAction
type ResourceActionClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewResourceActionClientWithConfigurationProvider Creates a new default ResourceAction client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewResourceActionClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ResourceActionClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newResourceActionClientFromBaseClient(baseClient, configProvider)
}

// NewResourceActionClientWithOboToken Creates a new default ResourceAction client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewResourceActionClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ResourceActionClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newResourceActionClientFromBaseClient(baseClient, configProvider)
}

func newResourceActionClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ResourceActionClient, err error) {
	client = ResourceActionClient{BaseClient: baseClient}
	client.BasePath = "20200606"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ResourceActionClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("optimizer", "https://optimizer-api.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ResourceActionClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ResourceActionClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetResourceAction Gets the resource action that corresponds to the specified OCID.
func (client ResourceActionClient) GetResourceAction(ctx context.Context, request GetResourceActionRequest) (response GetResourceActionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResourceAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResourceActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResourceActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResourceActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResourceActionResponse")
	}
	return
}

// getResourceAction implements the OCIOperation interface (enables retrying operations)
func (client ResourceActionClient) getResourceAction(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceActions/{resourceActionId}")
	if err != nil {
		return nil, err
	}

	var response GetResourceActionResponse
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

// ListResourceActions Included a list of resource actions that belong to the given recommendation.
func (client ResourceActionClient) ListResourceActions(ctx context.Context, request ListResourceActionsRequest) (response ListResourceActionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResourceActions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourceActionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourceActionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceActionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceActionsResponse")
	}
	return
}

// listResourceActions implements the OCIOperation interface (enables retrying operations)
func (client ResourceActionClient) listResourceActions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceActions")
	if err != nil {
		return nil, err
	}

	var response ListResourceActionsResponse
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

// UpdateResourceAction Updates the resource action that corresponds to the specified OCID.
// Use this operation to implement the following actions:
//   * Postpone resource action
//   * Ignore resource action
//   * Reactivate resource action
func (client ResourceActionClient) UpdateResourceAction(ctx context.Context, request UpdateResourceActionRequest) (response UpdateResourceActionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateResourceAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateResourceActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateResourceActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateResourceActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateResourceActionResponse")
	}
	return
}

// updateResourceAction implements the OCIOperation interface (enables retrying operations)
func (client ResourceActionClient) updateResourceAction(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/resourceActions/{resourceActionId}")
	if err != nil {
		return nil, err
	}

	var response UpdateResourceActionResponse
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
