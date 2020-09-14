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
	"github.com/oracle/oci-go-sdk/common/auth"
	"net/http"
)

//HistoryClient a client for History
type HistoryClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewHistoryClientWithConfigurationProvider Creates a new default History client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewHistoryClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client HistoryClient, err error) {
	if provider, err := auth.GetGenericConfigurationProvider(configProvider); err == nil {
		if baseClient, err := common.NewClientWithConfig(provider); err == nil {
			return newHistoryClientFromBaseClient(baseClient, configProvider)
		}
	}

	return
}

// NewHistoryClientWithOboToken Creates a new default History client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewHistoryClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client HistoryClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newHistoryClientFromBaseClient(baseClient, configProvider)
}

func newHistoryClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client HistoryClient, err error) {
	client = HistoryClient{BaseClient: baseClient}
	client.BasePath = "20200606"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *HistoryClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("optimizer", "https://optimizer.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *HistoryClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *HistoryClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListHistories Lists changes to the recommendations based on user activity.
// For example, lists when recommendations have been implemented, dismissed, postponed, or reactivated.
func (client HistoryClient) ListHistories(ctx context.Context, request ListHistoriesRequest) (response ListHistoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listHistories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListHistoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListHistoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListHistoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListHistoriesResponse")
	}
	return
}

// listHistories implements the OCIOperation interface (enables retrying operations)
func (client HistoryClient) listHistories(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/histories")
	if err != nil {
		return nil, err
	}

	var response ListHistoriesResponse
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
