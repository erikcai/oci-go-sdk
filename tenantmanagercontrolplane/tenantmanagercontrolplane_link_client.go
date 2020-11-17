// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// TenantManager API
//
// A description of the TenantManager API
//

package tenantmanagercontrolplane

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v29/common"
	"github.com/oracle/oci-go-sdk/v29/common/auth"
	"net/http"
)

//LinkClient a client for Link
type LinkClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLinkClientWithConfigurationProvider Creates a new default Link client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLinkClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LinkClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newLinkClientFromBaseClient(baseClient, provider)
}

// NewLinkClientWithOboToken Creates a new default Link client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewLinkClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LinkClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newLinkClientFromBaseClient(baseClient, configProvider)
}

func newLinkClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LinkClient, err error) {
	client = LinkClient{BaseClient: baseClient}
	client.BasePath = "20200801"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LinkClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("tenantmanagercontrolplane", "https://organizations.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LinkClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LinkClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// DeleteLink Terminate the link.
func (client LinkClient) DeleteLink(ctx context.Context, request DeleteLinkRequest) (response DeleteLinkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLink, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLinkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLinkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLinkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLinkResponse")
	}
	return
}

// deleteLink implements the OCIOperation interface (enables retrying operations)
func (client LinkClient) deleteLink(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/links/{linkId}")
	if err != nil {
		return nil, err
	}

	var response DeleteLinkResponse
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

// GetLink Gets information about the link.
func (client LinkClient) GetLink(ctx context.Context, request GetLinkRequest) (response GetLinkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLink, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLinkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLinkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLinkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLinkResponse")
	}
	return
}

// getLink implements the OCIOperation interface (enables retrying operations)
func (client LinkClient) getLink(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/links/{linkId}")
	if err != nil {
		return nil, err
	}

	var response GetLinkResponse
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

// ListLinks Return a (paginated) list of links.
func (client LinkClient) ListLinks(ctx context.Context, request ListLinksRequest) (response ListLinksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLinks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLinksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLinksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLinksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLinksResponse")
	}
	return
}

// listLinks implements the OCIOperation interface (enables retrying operations)
func (client LinkClient) listLinks(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/links")
	if err != nil {
		return nil, err
	}

	var response ListLinksResponse
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
