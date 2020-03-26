// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Registry Extension API
//
// API for managing images and repositories in Oracle Cloud Infrastructure Registry.
//

package containerregistry

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//ContainerRegistryClient a client for ContainerRegistry
type ContainerRegistryClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewContainerRegistryClientWithConfigurationProvider Creates a new default ContainerRegistry client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewContainerRegistryClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ContainerRegistryClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newContainerRegistryClientFromBaseClient(baseClient, configProvider)
}

// NewContainerRegistryClientWithOboToken Creates a new default ContainerRegistry client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewContainerRegistryClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ContainerRegistryClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newContainerRegistryClientFromBaseClient(baseClient, configProvider)
}

func newContainerRegistryClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ContainerRegistryClient, err error) {
	client = ContainerRegistryClient{BaseClient: baseClient}
	client.BasePath = "20200202"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ContainerRegistryClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("containerregistry", "https://ocir.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ContainerRegistryClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ContainerRegistryClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeDockerRepositoryCompartment Moves a docker repository into a different compartment. When provided, If-Match is checked against ETag values of the resource.
func (client ContainerRegistryClient) ChangeDockerRepositoryCompartment(ctx context.Context, request ChangeDockerRepositoryCompartmentRequest) (response ChangeDockerRepositoryCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeDockerRepositoryCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = ChangeDockerRepositoryCompartmentResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDockerRepositoryCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDockerRepositoryCompartmentResponse")
	}
	return
}

// changeDockerRepositoryCompartment implements the OCIOperation interface (enables retrying operations)
func (client ContainerRegistryClient) changeDockerRepositoryCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dockerRepositories/{dockerRepositoryId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeDockerRepositoryCompartmentResponse
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

// CreateDockerRepository Create a new empty docker repository in the tenancy. Avoid entering confidential information.
func (client ContainerRegistryClient) CreateDockerRepository(ctx context.Context, request CreateDockerRepositoryRequest) (response CreateDockerRepositoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDockerRepository, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = CreateDockerRepositoryResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDockerRepositoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDockerRepositoryResponse")
	}
	return
}

// createDockerRepository implements the OCIOperation interface (enables retrying operations)
func (client ContainerRegistryClient) createDockerRepository(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dockerRepositories")
	if err != nil {
		return nil, err
	}

	var response CreateDockerRepositoryResponse
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

// DeleteDockerRepository Delete docker repository
func (client ContainerRegistryClient) DeleteDockerRepository(ctx context.Context, request DeleteDockerRepositoryRequest) (response DeleteDockerRepositoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDockerRepository, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = DeleteDockerRepositoryResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDockerRepositoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDockerRepositoryResponse")
	}
	return
}

// deleteDockerRepository implements the OCIOperation interface (enables retrying operations)
func (client ContainerRegistryClient) deleteDockerRepository(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dockerRepositories/{dockerRepositoryId}")
	if err != nil {
		return nil, err
	}

	var response DeleteDockerRepositoryResponse
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

// DeleteDockerRepositoryContents Delete only the contents for a docker repository
func (client ContainerRegistryClient) DeleteDockerRepositoryContents(ctx context.Context, request DeleteDockerRepositoryContentsRequest) (response DeleteDockerRepositoryContentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDockerRepositoryContents, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = DeleteDockerRepositoryContentsResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDockerRepositoryContentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDockerRepositoryContentsResponse")
	}
	return
}

// deleteDockerRepositoryContents implements the OCIOperation interface (enables retrying operations)
func (client ContainerRegistryClient) deleteDockerRepositoryContents(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dockerRepositories/{dockerRepositoryId}/actions/emptyRepo")
	if err != nil {
		return nil, err
	}

	var response DeleteDockerRepositoryContentsResponse
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

// DeleteDockerTag Delete a tag from a docker repository
func (client ContainerRegistryClient) DeleteDockerTag(ctx context.Context, request DeleteDockerTagRequest) (response DeleteDockerTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDockerTag, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = DeleteDockerTagResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDockerTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDockerTagResponse")
	}
	return
}

// deleteDockerTag implements the OCIOperation interface (enables retrying operations)
func (client ContainerRegistryClient) deleteDockerTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dockerRepositories/{dockerRepositoryId}/tags/{tagName}")
	if err != nil {
		return nil, err
	}

	var response DeleteDockerTagResponse
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

// GetDockerImageMetadata Get details for the specified docker repository image
func (client ContainerRegistryClient) GetDockerImageMetadata(ctx context.Context, request GetDockerImageMetadataRequest) (response GetDockerImageMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDockerImageMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = GetDockerImageMetadataResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDockerImageMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDockerImageMetadataResponse")
	}
	return
}

// getDockerImageMetadata implements the OCIOperation interface (enables retrying operations)
func (client ContainerRegistryClient) getDockerImageMetadata(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dockerRepositories/{dockerRepositoryId}/images/{imageDigestOrTagName}")
	if err != nil {
		return nil, err
	}

	var response GetDockerImageMetadataResponse
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

// GetDockerRepository Get summary for specified docker repository
func (client ContainerRegistryClient) GetDockerRepository(ctx context.Context, request GetDockerRepositoryRequest) (response GetDockerRepositoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDockerRepository, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = GetDockerRepositoryResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDockerRepositoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDockerRepositoryResponse")
	}
	return
}

// getDockerRepository implements the OCIOperation interface (enables retrying operations)
func (client ContainerRegistryClient) getDockerRepository(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dockerRepositories/{dockerRepositoryId}")
	if err != nil {
		return nil, err
	}

	var response GetDockerRepositoryResponse
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

// GetDockerToken Basic or signature header-based access authentication
func (client ContainerRegistryClient) GetDockerToken(ctx context.Context, request GetDockerTokenRequest) (response GetDockerTokenResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDockerToken, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = GetDockerTokenResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDockerTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDockerTokenResponse")
	}
	return
}

// getDockerToken implements the OCIOperation interface (enables retrying operations)
func (client ContainerRegistryClient) getDockerToken(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dockerToken")
	if err != nil {
		return nil, err
	}

	var response GetDockerTokenResponse
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

// GetDockerTokenWithOAuth OAuth authentication as described at https://docs.docker.com/registry/spec/auth/oauth/
func (client ContainerRegistryClient) GetDockerTokenWithOAuth(ctx context.Context, request GetDockerTokenWithOAuthRequest) (response GetDockerTokenWithOAuthResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDockerTokenWithOAuth, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = GetDockerTokenWithOAuthResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDockerTokenWithOAuthResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDockerTokenWithOAuthResponse")
	}
	return
}

// getDockerTokenWithOAuth implements the OCIOperation interface (enables retrying operations)
func (client ContainerRegistryClient) getDockerTokenWithOAuth(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dockerToken")
	if err != nil {
		return nil, err
	}

	var response GetDockerTokenWithOAuthResponse
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

// GetNamespace Get the namespace of the compartment
func (client ContainerRegistryClient) GetNamespace(ctx context.Context, request GetNamespaceRequest) (response GetNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = GetNamespaceResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNamespaceResponse")
	}
	return
}

// getNamespace implements the OCIOperation interface (enables retrying operations)
func (client ContainerRegistryClient) getNamespace(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespace")
	if err != nil {
		return nil, err
	}

	var response GetNamespaceResponse
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

// ListDockerRepositories List docker repositories
func (client ContainerRegistryClient) ListDockerRepositories(ctx context.Context, request ListDockerRepositoriesRequest) (response ListDockerRepositoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDockerRepositories, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = ListDockerRepositoriesResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDockerRepositoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDockerRepositoriesResponse")
	}
	return
}

// listDockerRepositories implements the OCIOperation interface (enables retrying operations)
func (client ContainerRegistryClient) listDockerRepositories(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dockerRepositories")
	if err != nil {
		return nil, err
	}

	var response ListDockerRepositoriesResponse
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

// ListDockerTags List docker tags across docker repositories in a compartment
func (client ContainerRegistryClient) ListDockerTags(ctx context.Context, request ListDockerTagsRequest) (response ListDockerTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDockerTags, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = ListDockerTagsResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDockerTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDockerTagsResponse")
	}
	return
}

// listDockerTags implements the OCIOperation interface (enables retrying operations)
func (client ContainerRegistryClient) listDockerTags(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dockerTags")
	if err != nil {
		return nil, err
	}

	var response ListDockerTagsResponse
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

// UndeleteManifest Undelete a deleted manifest for a docker repository
func (client ContainerRegistryClient) UndeleteManifest(ctx context.Context, request UndeleteManifestRequest) (response UndeleteManifestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.undeleteManifest, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = UndeleteManifestResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UndeleteManifestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UndeleteManifestResponse")
	}
	return
}

// undeleteManifest implements the OCIOperation interface (enables retrying operations)
func (client ContainerRegistryClient) undeleteManifest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dockerRepositories/{dockerRepositoryId}/actions/undeleteManifest")
	if err != nil {
		return nil, err
	}

	var response UndeleteManifestResponse
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

// UpdateDockerRepository This endpoint allows the caller to modify the properties of a docker repository. Avoid entering confidential information.
func (client ContainerRegistryClient) UpdateDockerRepository(ctx context.Context, request UpdateDockerRepositoryRequest) (response UpdateDockerRepositoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDockerRepository, policy)
	if err != nil {
		if ociResponse != nil {
			opcRequestId := ociResponse.HTTPResponse().Header.Get("opc-request-id")
			response = UpdateDockerRepositoryResponse{RawResponse: ociResponse.HTTPResponse(), OpcRequestId: &opcRequestId}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDockerRepositoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDockerRepositoryResponse")
	}
	return
}

// updateDockerRepository implements the OCIOperation interface (enables retrying operations)
func (client ContainerRegistryClient) updateDockerRepository(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dockerRepositories/{dockerRepositoryId}")
	if err != nil {
		return nil, err
	}

	var response UpdateDockerRepositoryResponse
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
