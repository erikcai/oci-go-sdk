// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// UsageApi API
//
// A description of the UsageApi API.
//

package usage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v28/common"
	"github.com/oracle/oci-go-sdk/v28/common/auth"
	"net/http"
)

//UsageClient a client for Usage
type UsageClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewUsageClientWithConfigurationProvider Creates a new default Usage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewUsageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client UsageClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newUsageClientFromBaseClient(baseClient, provider)
}

// NewUsageClientWithOboToken Creates a new default Usage client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewUsageClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client UsageClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newUsageClientFromBaseClient(baseClient, configProvider)
}

func newUsageClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client UsageClient, err error) {
	client = UsageClient{BaseClient: baseClient}
	client.BasePath = "20190111"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *UsageClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("identity", "https://identity.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *UsageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *UsageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetSubscriptionInfo Returns the subscription information for the specified tenancy.
// > **Important**: Call to this API will only succeed against the endpoint in the home region or the tenancy.
func (client UsageClient) GetSubscriptionInfo(ctx context.Context, request GetSubscriptionInfoRequest) (response GetSubscriptionInfoResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSubscriptionInfo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSubscriptionInfoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSubscriptionInfoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSubscriptionInfoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSubscriptionInfoResponse")
	}
	return
}

// getSubscriptionInfo implements the OCIOperation interface (enables retrying operations)
func (client UsageClient) getSubscriptionInfo(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/usagecost/{tenancyId}/subscriptionInfo")
	if err != nil {
		return nil, err
	}

	var response GetSubscriptionInfoResponse
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

// ListUsageRecords Returns the usage data and date range for the given tenancy, including cost information broken out by a specified granularity.
// Depending on the granularity, the date range cannot be higher than
//   - for DAILY/MONTHLY => 1 month
//   - for HOURLY => 1 day
// > **Important**: Call to this API will only succeed against the endpoint in the home region.
func (client UsageClient) ListUsageRecords(ctx context.Context, request ListUsageRecordsRequest) (response ListUsageRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUsageRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUsageRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUsageRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUsageRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUsageRecordsResponse")
	}
	return
}

// listUsageRecords implements the OCIOperation interface (enables retrying operations)
func (client UsageClient) listUsageRecords(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/usagecost/{tenancyId}")
	if err != nil {
		return nil, err
	}

	var response ListUsageRecordsResponse
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
