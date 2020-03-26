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

//LimitsChecksClient a client for LimitsChecks
type LimitsChecksClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLimitsChecksClientWithConfigurationProvider Creates a new default LimitsChecks client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLimitsChecksClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LimitsChecksClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newLimitsChecksClientFromBaseClient(baseClient, configProvider)
}

// NewLimitsChecksClientWithOboToken Creates a new default LimitsChecks client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewLimitsChecksClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LimitsChecksClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newLimitsChecksClientFromBaseClient(baseClient, configProvider)
}

func newLimitsChecksClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LimitsChecksClient, err error) {
	client = LimitsChecksClient{BaseClient: baseClient}
	client.BasePath = "20190415"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LimitsChecksClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("mysql", "https://mysql.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LimitsChecksClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LimitsChecksClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetMysqlaasInstanceLimitsCheck Gets an object that is either empty or describes an error (if there are not enough available resources to launch a MySQLaaS
// Instance), according to the given instance details. It does not reserve the resources, just provides a way to
// fail fast if there are no resources at the moment, not preventing the case of concurrent launches that will pass
// the checks and fail later because resources are no longer available when actually launching the instances.
func (client LimitsChecksClient) GetMysqlaasInstanceLimitsCheck(ctx context.Context, request GetMysqlaasInstanceLimitsCheckRequest) (response GetMysqlaasInstanceLimitsCheckResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMysqlaasInstanceLimitsCheck, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = GetMysqlaasInstanceLimitsCheckResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMysqlaasInstanceLimitsCheckResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMysqlaasInstanceLimitsCheckResponse")
	}
	return
}

// getMysqlaasInstanceLimitsCheck implements the OCIOperation interface (enables retrying operations)
func (client LimitsChecksClient) getMysqlaasInstanceLimitsCheck(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/limits/mysqlaasInstance")
	if err != nil {
		return nil, err
	}

	var response GetMysqlaasInstanceLimitsCheckResponse
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
