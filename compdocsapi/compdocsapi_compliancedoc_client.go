// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compliance Documents Service API
//
// API for downloading Oracle Cloud Infrastructure compliance documents from the Oracle Cloud Infrastructure Console.
//

package compdocsapi

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v27/common"
	"github.com/oracle/oci-go-sdk/v27/common/auth"
	"net/http"
)

//ComplianceDocClient a client for ComplianceDoc
type ComplianceDocClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewComplianceDocClientWithConfigurationProvider Creates a new default ComplianceDoc client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewComplianceDocClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ComplianceDocClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newComplianceDocClientFromBaseClient(baseClient, provider)
}

// NewComplianceDocClientWithOboToken Creates a new default ComplianceDoc client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewComplianceDocClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ComplianceDocClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newComplianceDocClientFromBaseClient(baseClient, configProvider)
}

func newComplianceDocClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ComplianceDocClient, err error) {
	client = ComplianceDocClient{BaseClient: baseClient}
	client.BasePath = "20200225"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ComplianceDocClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("compdocsapi", "https://compdocs.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ComplianceDocClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ComplianceDocClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateNonDisclosureAgreement Creates a new non-disclosure agreement to associate with a compliance document, according to the request details.
func (client ComplianceDocClient) CreateNonDisclosureAgreement(ctx context.Context, request CreateNonDisclosureAgreementRequest) (response CreateNonDisclosureAgreementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createNonDisclosureAgreement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateNonDisclosureAgreementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateNonDisclosureAgreementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateNonDisclosureAgreementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateNonDisclosureAgreementResponse")
	}
	return
}

// createNonDisclosureAgreement implements the OCIOperation interface (enables retrying operations)
func (client ComplianceDocClient) createNonDisclosureAgreement(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/nonDisclosureAgreements")
	if err != nil {
		return nil, err
	}

	var response CreateNonDisclosureAgreementResponse
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

// GetComplianceDocument Gets the details of a specific compliance document.
func (client ComplianceDocClient) GetComplianceDocument(ctx context.Context, request GetComplianceDocumentRequest) (response GetComplianceDocumentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getComplianceDocument, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetComplianceDocumentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetComplianceDocumentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetComplianceDocumentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetComplianceDocumentResponse")
	}
	return
}

// getComplianceDocument implements the OCIOperation interface (enables retrying operations)
func (client ComplianceDocClient) getComplianceDocument(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/complianceDocuments/{complianceDocumentId}")
	if err != nil {
		return nil, err
	}

	var response GetComplianceDocumentResponse
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

// GetComplianceDocumentContent Downloads the specified compliance document.
func (client ComplianceDocClient) GetComplianceDocumentContent(ctx context.Context, request GetComplianceDocumentContentRequest) (response GetComplianceDocumentContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getComplianceDocumentContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetComplianceDocumentContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetComplianceDocumentContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetComplianceDocumentContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetComplianceDocumentContentResponse")
	}
	return
}

// getComplianceDocumentContent implements the OCIOperation interface (enables retrying operations)
func (client ComplianceDocClient) getComplianceDocumentContent(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/complianceDocuments/{complianceDocumentId}/content")
	if err != nil {
		return nil, err
	}

	var response GetComplianceDocumentContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListComplianceDocuments Gets a list of all compliance documents in the specified compartment.
func (client ComplianceDocClient) ListComplianceDocuments(ctx context.Context, request ListComplianceDocumentsRequest) (response ListComplianceDocumentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listComplianceDocuments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListComplianceDocumentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListComplianceDocumentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListComplianceDocumentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListComplianceDocumentsResponse")
	}
	return
}

// listComplianceDocuments implements the OCIOperation interface (enables retrying operations)
func (client ComplianceDocClient) listComplianceDocuments(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/complianceDocuments")
	if err != nil {
		return nil, err
	}

	var response ListComplianceDocumentsResponse
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
