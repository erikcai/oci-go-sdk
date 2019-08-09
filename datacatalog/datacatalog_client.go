// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DataCatalog API
//
// A description of the DataCatalog API
//

package datacatalog

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//DataCatalogClient a client for DataCatalog
type DataCatalogClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDataCatalogClientWithConfigurationProvider Creates a new default DataCatalog client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDataCatalogClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DataCatalogClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = DataCatalogClient{BaseClient: baseClient}
	client.BasePath = "20190325"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DataCatalogClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("datacatalog", "https://datacatalog.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DataCatalogClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DataCatalogClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateAttribute Create a new Entity Attribute.
func (client DataCatalogClient) CreateAttribute(ctx context.Context, request CreateAttributeRequest) (response CreateAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateAttributeResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAttributeResponse")
	}
	return
}

// createAttribute implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createAttribute(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes")
	if err != nil {
		return nil, err
	}

	var response CreateAttributeResponse
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

// CreateAttributeTag Create a new Entity Attribute Tag.
func (client DataCatalogClient) CreateAttributeTag(ctx context.Context, request CreateAttributeTagRequest) (response CreateAttributeTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAttributeTag, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateAttributeTagResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAttributeTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAttributeTagResponse")
	}
	return
}

// createAttributeTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createAttributeTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}/tags")
	if err != nil {
		return nil, err
	}

	var response CreateAttributeTagResponse
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

// CreateCatalog Creates a new Data Catalog instance which includes a console and api url for managing metadata operations.
// For more information, please see the documentation.
func (client DataCatalogClient) CreateCatalog(ctx context.Context, request CreateCatalogRequest) (response CreateCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateCatalogResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCatalogResponse")
	}
	return
}

// createCatalog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createCatalog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs")
	if err != nil {
		return nil, err
	}

	var response CreateCatalogResponse
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

// CreateConnection Create a new Connection.
func (client DataCatalogClient) CreateConnection(ctx context.Context, request CreateConnectionRequest) (response CreateConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createConnection, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateConnectionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConnectionResponse")
	}
	return
}

// createConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createConnection(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections")
	if err != nil {
		return nil, err
	}

	var response CreateConnectionResponse
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

// CreateDataAsset Create a new Data Asset.
func (client DataCatalogClient) CreateDataAsset(ctx context.Context, request CreateDataAssetRequest) (response CreateDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateDataAssetResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataAssetResponse")
	}
	return
}

// createDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createDataAsset(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets")
	if err != nil {
		return nil, err
	}

	var response CreateDataAssetResponse
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

// CreateDataAssetTag Create a new Data Asset Tag.
func (client DataCatalogClient) CreateDataAssetTag(ctx context.Context, request CreateDataAssetTagRequest) (response CreateDataAssetTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDataAssetTag, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateDataAssetTagResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataAssetTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataAssetTagResponse")
	}
	return
}

// createDataAssetTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createDataAssetTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/tags")
	if err != nil {
		return nil, err
	}

	var response CreateDataAssetTagResponse
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

// CreateEntity Create a new Entity.
func (client DataCatalogClient) CreateEntity(ctx context.Context, request CreateEntityRequest) (response CreateEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createEntity, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateEntityResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEntityResponse")
	}
	return
}

// createEntity implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createEntity(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities")
	if err != nil {
		return nil, err
	}

	var response CreateEntityResponse
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

// CreateEntityTag Create a new Entity Tag.
func (client DataCatalogClient) CreateEntityTag(ctx context.Context, request CreateEntityTagRequest) (response CreateEntityTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createEntityTag, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateEntityTagResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEntityTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEntityTagResponse")
	}
	return
}

// createEntityTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createEntityTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/tags")
	if err != nil {
		return nil, err
	}

	var response CreateEntityTagResponse
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

// CreateFolder Create a new Folder.
func (client DataCatalogClient) CreateFolder(ctx context.Context, request CreateFolderRequest) (response CreateFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createFolder, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateFolderResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFolderResponse")
	}
	return
}

// createFolder implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createFolder(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders")
	if err != nil {
		return nil, err
	}

	var response CreateFolderResponse
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

// CreateFolderTag Create a new Folder Tag.
func (client DataCatalogClient) CreateFolderTag(ctx context.Context, request CreateFolderTagRequest) (response CreateFolderTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createFolderTag, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateFolderTagResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFolderTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFolderTagResponse")
	}
	return
}

// createFolderTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createFolderTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}/tags")
	if err != nil {
		return nil, err
	}

	var response CreateFolderTagResponse
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

// CreateGlossary Create a new Glossary.
func (client DataCatalogClient) CreateGlossary(ctx context.Context, request CreateGlossaryRequest) (response CreateGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateGlossaryResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateGlossaryResponse")
	}
	return
}

// createGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createGlossary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries")
	if err != nil {
		return nil, err
	}

	var response CreateGlossaryResponse
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

// CreateGlossaryTerm Create a new Term within a Glossary.
func (client DataCatalogClient) CreateGlossaryTerm(ctx context.Context, request CreateGlossaryTermRequest) (response CreateGlossaryTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createGlossaryTerm, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateGlossaryTermResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateGlossaryTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateGlossaryTermResponse")
	}
	return
}

// createGlossaryTerm implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createGlossaryTerm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms")
	if err != nil {
		return nil, err
	}

	var response CreateGlossaryTermResponse
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

// CreateJob Create a new job.
func (client DataCatalogClient) CreateJob(ctx context.Context, request CreateJobRequest) (response CreateJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createJob, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateJobResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJobResponse")
	}
	return
}

// createJob implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/jobs")
	if err != nil {
		return nil, err
	}

	var response CreateJobResponse
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

// CreateJobDefinition Create a new job definition.
func (client DataCatalogClient) CreateJobDefinition(ctx context.Context, request CreateJobDefinitionRequest) (response CreateJobDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createJobDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateJobDefinitionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJobDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJobDefinitionResponse")
	}
	return
}

// createJobDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createJobDefinition(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/jobDefinitions")
	if err != nil {
		return nil, err
	}

	var response CreateJobDefinitionResponse
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

// DeleteAttribute Delete a specific Entity Attribute.
func (client DataCatalogClient) DeleteAttribute(ctx context.Context, request DeleteAttributeRequest) (response DeleteAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteAttributeResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAttributeResponse")
	}
	return
}

// deleteAttribute implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteAttribute(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteAttributeResponse
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

// DeleteAttributeTag Delete a specific Entity Attribute Tag.
func (client DataCatalogClient) DeleteAttributeTag(ctx context.Context, request DeleteAttributeTagRequest) (response DeleteAttributeTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAttributeTag, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteAttributeTagResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAttributeTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAttributeTagResponse")
	}
	return
}

// deleteAttributeTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteAttributeTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteAttributeTagResponse
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

// DeleteCatalog Deletes a Catalog resource by identifier
func (client DataCatalogClient) DeleteCatalog(ctx context.Context, request DeleteCatalogRequest) (response DeleteCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteCatalogResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCatalogResponse")
	}
	return
}

// deleteCatalog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteCatalog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}")
	if err != nil {
		return nil, err
	}

	var response DeleteCatalogResponse
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

// DeleteConnection Delete a specific Connection of a Data Asset.
func (client DataCatalogClient) DeleteConnection(ctx context.Context, request DeleteConnectionRequest) (response DeleteConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConnection, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteConnectionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConnectionResponse")
	}
	return
}

// deleteConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteConnection(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections/{connectionKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteConnectionResponse
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

// DeleteDataAsset Delete a specific Data Asset identified by it's key
func (client DataCatalogClient) DeleteDataAsset(ctx context.Context, request DeleteDataAssetRequest) (response DeleteDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteDataAssetResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDataAssetResponse")
	}
	return
}

// deleteDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteDataAsset(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteDataAssetResponse
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

// DeleteDataAssetTag Delete a specific Data Asset Tag.
func (client DataCatalogClient) DeleteDataAssetTag(ctx context.Context, request DeleteDataAssetTagRequest) (response DeleteDataAssetTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDataAssetTag, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteDataAssetTagResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDataAssetTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDataAssetTagResponse")
	}
	return
}

// deleteDataAssetTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteDataAssetTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteDataAssetTagResponse
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

// DeleteEntity Delete a specific Entity.
func (client DataCatalogClient) DeleteEntity(ctx context.Context, request DeleteEntityRequest) (response DeleteEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEntity, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteEntityResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEntityResponse")
	}
	return
}

// deleteEntity implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteEntity(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteEntityResponse
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

// DeleteEntityTag Delete a specific Entity Tag.
func (client DataCatalogClient) DeleteEntityTag(ctx context.Context, request DeleteEntityTagRequest) (response DeleteEntityTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEntityTag, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteEntityTagResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEntityTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEntityTagResponse")
	}
	return
}

// deleteEntityTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteEntityTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteEntityTagResponse
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

// DeleteFolder Delete a specific folder of a Data Asset identified by it's key
func (client DataCatalogClient) DeleteFolder(ctx context.Context, request DeleteFolderRequest) (response DeleteFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFolder, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteFolderResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFolderResponse")
	}
	return
}

// deleteFolder implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteFolder(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteFolderResponse
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

// DeleteFolderTag Delete a specific Folder Tag.
func (client DataCatalogClient) DeleteFolderTag(ctx context.Context, request DeleteFolderTagRequest) (response DeleteFolderTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFolderTag, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteFolderTagResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFolderTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFolderTagResponse")
	}
	return
}

// deleteFolderTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteFolderTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteFolderTagResponse
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

// DeleteGlossary Delete a specific Glossary identified by it's key
func (client DataCatalogClient) DeleteGlossary(ctx context.Context, request DeleteGlossaryRequest) (response DeleteGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteGlossaryResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGlossaryResponse")
	}
	return
}

// deleteGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteGlossary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/glossaries/{glossaryKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteGlossaryResponse
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

// DeleteGlossaryTerm Delete a specific Glossary Term.
func (client DataCatalogClient) DeleteGlossaryTerm(ctx context.Context, request DeleteGlossaryTermRequest) (response DeleteGlossaryTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteGlossaryTerm, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteGlossaryTermResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGlossaryTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGlossaryTermResponse")
	}
	return
}

// deleteGlossaryTerm implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteGlossaryTerm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteGlossaryTermResponse
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

// DeleteJob Delete a specific job  identified by it's key
func (client DataCatalogClient) DeleteJob(ctx context.Context, request DeleteJobRequest) (response DeleteJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteJob, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteJobResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteJobResponse")
	}
	return
}

// deleteJob implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/jobs/{jobKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteJobResponse
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

// DeleteJobDefinition Delete a specific job definition identified by it's key
func (client DataCatalogClient) DeleteJobDefinition(ctx context.Context, request DeleteJobDefinitionRequest) (response DeleteJobDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteJobDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteJobDefinitionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteJobDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteJobDefinitionResponse")
	}
	return
}

// deleteJobDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteJobDefinition(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/jobDefinitions/{jobDefinitionKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteJobDefinitionResponse
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

// GetAttribute Get a specific Entity Attribute by key.
func (client DataCatalogClient) GetAttribute(ctx context.Context, request GetAttributeRequest) (response GetAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAttributeResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAttributeResponse")
	}
	return
}

// getAttribute implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getAttribute(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}")
	if err != nil {
		return nil, err
	}

	var response GetAttributeResponse
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

// GetAttributeTag Get a specific Entity Attribute Tag by key.
func (client DataCatalogClient) GetAttributeTag(ctx context.Context, request GetAttributeTagRequest) (response GetAttributeTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAttributeTag, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAttributeTagResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAttributeTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAttributeTagResponse")
	}
	return
}

// getAttributeTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getAttributeTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response GetAttributeTagResponse
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

// GetCatalog Gets a Catalog by identifier
func (client DataCatalogClient) GetCatalog(ctx context.Context, request GetCatalogRequest) (response GetCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetCatalogResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCatalogResponse")
	}
	return
}

// getCatalog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getCatalog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}")
	if err != nil {
		return nil, err
	}

	var response GetCatalogResponse
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

// GetConnection Get a specific Data Asset Connection by key.
func (client DataCatalogClient) GetConnection(ctx context.Context, request GetConnectionRequest) (response GetConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConnection, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetConnectionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConnectionResponse")
	}
	return
}

// getConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getConnection(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections/{connectionKey}")
	if err != nil {
		return nil, err
	}

	var response GetConnectionResponse
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

// GetDataAsset Get a specific DataAsset for the given key within a data catalog.
func (client DataCatalogClient) GetDataAsset(ctx context.Context, request GetDataAssetRequest) (response GetDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetDataAssetResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataAssetResponse")
	}
	return
}

// getDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getDataAsset(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}")
	if err != nil {
		return nil, err
	}

	var response GetDataAssetResponse
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

// GetDataAssetTag Get a specific Data Asset Tag by key.
func (client DataCatalogClient) GetDataAssetTag(ctx context.Context, request GetDataAssetTagRequest) (response GetDataAssetTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataAssetTag, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetDataAssetTagResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataAssetTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataAssetTagResponse")
	}
	return
}

// getDataAssetTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getDataAssetTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response GetDataAssetTagResponse
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

// GetEntity Get a specific Entity by key for a Data Asset.
func (client DataCatalogClient) GetEntity(ctx context.Context, request GetEntityRequest) (response GetEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEntity, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetEntityResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEntityResponse")
	}
	return
}

// getEntity implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getEntity(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}")
	if err != nil {
		return nil, err
	}

	var response GetEntityResponse
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

// GetEntityTag Get a specific Entity Tag by key.
func (client DataCatalogClient) GetEntityTag(ctx context.Context, request GetEntityTagRequest) (response GetEntityTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEntityTag, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetEntityTagResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEntityTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEntityTagResponse")
	}
	return
}

// getEntityTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getEntityTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response GetEntityTagResponse
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

// GetFolder Get a specific Data Asset Folder by key.
func (client DataCatalogClient) GetFolder(ctx context.Context, request GetFolderRequest) (response GetFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFolder, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetFolderResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFolderResponse")
	}
	return
}

// getFolder implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getFolder(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}")
	if err != nil {
		return nil, err
	}

	var response GetFolderResponse
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

// GetFolderTag Get a specific Folder Tag by key.
func (client DataCatalogClient) GetFolderTag(ctx context.Context, request GetFolderTagRequest) (response GetFolderTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFolderTag, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetFolderTagResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFolderTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFolderTagResponse")
	}
	return
}

// getFolderTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getFolderTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response GetFolderTagResponse
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

// GetGlossary Get a specific Glossary by key within a data catalog.
func (client DataCatalogClient) GetGlossary(ctx context.Context, request GetGlossaryRequest) (response GetGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetGlossaryResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGlossaryResponse")
	}
	return
}

// getGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getGlossary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries/{glossaryKey}")
	if err != nil {
		return nil, err
	}

	var response GetGlossaryResponse
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

// GetGlossaryTerm Get a specific Glossary Term by key.
func (client DataCatalogClient) GetGlossaryTerm(ctx context.Context, request GetGlossaryTermRequest) (response GetGlossaryTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGlossaryTerm, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetGlossaryTermResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGlossaryTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGlossaryTermResponse")
	}
	return
}

// getGlossaryTerm implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getGlossaryTerm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}")
	if err != nil {
		return nil, err
	}

	var response GetGlossaryTermResponse
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

// GetJob Get a specific Job by key within a data catalog.
func (client DataCatalogClient) GetJob(ctx context.Context, request GetJobRequest) (response GetJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJob, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetJobResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobResponse")
	}
	return
}

// getJob implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}")
	if err != nil {
		return nil, err
	}

	var response GetJobResponse
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

// GetJobDefinition Get a specific Job Definition by key within a data catalog.
func (client DataCatalogClient) GetJobDefinition(ctx context.Context, request GetJobDefinitionRequest) (response GetJobDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetJobDefinitionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobDefinitionResponse")
	}
	return
}

// getJobDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJobDefinition(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobDefinitions/{jobDefinitionKey}")
	if err != nil {
		return nil, err
	}

	var response GetJobDefinitionResponse
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

// GetJobExecution Get a specific Job Execution by key.
func (client DataCatalogClient) GetJobExecution(ctx context.Context, request GetJobExecutionRequest) (response GetJobExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobExecution, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetJobExecutionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobExecutionResponse")
	}
	return
}

// getJobExecution implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJobExecution(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}")
	if err != nil {
		return nil, err
	}

	var response GetJobExecutionResponse
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

// GetJobLog Get a specific Job Log by key.
func (client DataCatalogClient) GetJobLog(ctx context.Context, request GetJobLogRequest) (response GetJobLogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobLog, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetJobLogResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobLogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobLogResponse")
	}
	return
}

// getJobLog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJobLog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}/logs/{jobLogKey}")
	if err != nil {
		return nil, err
	}

	var response GetJobLogResponse
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

// GetJobMetrics Get a specific Job Metric by key.
func (client DataCatalogClient) GetJobMetrics(ctx context.Context, request GetJobMetricsRequest) (response GetJobMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetJobMetricsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobMetricsResponse")
	}
	return
}

// getJobMetrics implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJobMetrics(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}/metrics/{jobMetricsKey}")
	if err != nil {
		return nil, err
	}

	var response GetJobMetricsResponse
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

// GetType Get a specific Type by key within a data catalog.
func (client DataCatalogClient) GetType(ctx context.Context, request GetTypeRequest) (response GetTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getType, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetTypeResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTypeResponse")
	}
	return
}

// getType implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getType(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/types/{typeKey}")
	if err != nil {
		return nil, err
	}

	var response GetTypeResponse
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

// GetWorkRequest Gets the status of the work request with the given OCID.
func (client DataCatalogClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetWorkRequestResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}")
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
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

// ListAttributeTags Returns a list of all Tags for an Entity Attribute.
func (client DataCatalogClient) ListAttributeTags(ctx context.Context, request ListAttributeTagsRequest) (response ListAttributeTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAttributeTags, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListAttributeTagsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAttributeTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAttributeTagsResponse")
	}
	return
}

// listAttributeTags implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listAttributeTags(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}/tags")
	if err != nil {
		return nil, err
	}

	var response ListAttributeTagsResponse
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

// ListAttributes Returns a list of all Attributes of an Entity.
func (client DataCatalogClient) ListAttributes(ctx context.Context, request ListAttributesRequest) (response ListAttributesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAttributes, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListAttributesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAttributesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAttributesResponse")
	}
	return
}

// listAttributes implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listAttributes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes")
	if err != nil {
		return nil, err
	}

	var response ListAttributesResponse
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

// ListCatalogPermissions List Catalog Permissions within a data catalog for the current user.
func (client DataCatalogClient) ListCatalogPermissions(ctx context.Context, request ListCatalogPermissionsRequest) (response ListCatalogPermissionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCatalogPermissions, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListCatalogPermissionsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCatalogPermissionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCatalogPermissionsResponse")
	}
	return
}

// listCatalogPermissions implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listCatalogPermissions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/permissions")
	if err != nil {
		return nil, err
	}

	var response ListCatalogPermissionsResponse
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

// ListCatalogs Returns a list of Catalogs.
func (client DataCatalogClient) ListCatalogs(ctx context.Context, request ListCatalogsRequest) (response ListCatalogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCatalogs, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListCatalogsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCatalogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCatalogsResponse")
	}
	return
}

// listCatalogs implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listCatalogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs")
	if err != nil {
		return nil, err
	}

	var response ListCatalogsResponse
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

// ListConnections Returns a list of all Connections of a Data Asset.
func (client DataCatalogClient) ListConnections(ctx context.Context, request ListConnectionsRequest) (response ListConnectionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConnections, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListConnectionsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConnectionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConnectionsResponse")
	}
	return
}

// listConnections implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listConnections(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections")
	if err != nil {
		return nil, err
	}

	var response ListConnectionsResponse
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

// ListDataAssetPermissions Returns a list of Data Asset permissions.
func (client DataCatalogClient) ListDataAssetPermissions(ctx context.Context, request ListDataAssetPermissionsRequest) (response ListDataAssetPermissionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataAssetPermissions, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListDataAssetPermissionsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataAssetPermissionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataAssetPermissionsResponse")
	}
	return
}

// listDataAssetPermissions implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listDataAssetPermissions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/permissions")
	if err != nil {
		return nil, err
	}

	var response ListDataAssetPermissionsResponse
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

// ListDataAssetTags Returns a list of all Tags for a Data Asset.
func (client DataCatalogClient) ListDataAssetTags(ctx context.Context, request ListDataAssetTagsRequest) (response ListDataAssetTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataAssetTags, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListDataAssetTagsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataAssetTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataAssetTagsResponse")
	}
	return
}

// listDataAssetTags implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listDataAssetTags(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/tags")
	if err != nil {
		return nil, err
	}

	var response ListDataAssetTagsResponse
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

// ListDataAssets Returns a list of Data Assets within a data catalog.
func (client DataCatalogClient) ListDataAssets(ctx context.Context, request ListDataAssetsRequest) (response ListDataAssetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataAssets, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListDataAssetsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataAssetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataAssetsResponse")
	}
	return
}

// listDataAssets implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listDataAssets(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets")
	if err != nil {
		return nil, err
	}

	var response ListDataAssetsResponse
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

// ListEntities Returns a list of all Entities of a Data Asset.
func (client DataCatalogClient) ListEntities(ctx context.Context, request ListEntitiesRequest) (response ListEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEntities, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListEntitiesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEntitiesResponse")
	}
	return
}

// listEntities implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listEntities(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities")
	if err != nil {
		return nil, err
	}

	var response ListEntitiesResponse
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

// ListEntityTags Returns a list of all Tags for an Entity.
func (client DataCatalogClient) ListEntityTags(ctx context.Context, request ListEntityTagsRequest) (response ListEntityTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEntityTags, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListEntityTagsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEntityTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEntityTagsResponse")
	}
	return
}

// listEntityTags implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listEntityTags(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/tags")
	if err != nil {
		return nil, err
	}

	var response ListEntityTagsResponse
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

// ListFolderTags Returns a list of all Tags for a Folder.
func (client DataCatalogClient) ListFolderTags(ctx context.Context, request ListFolderTagsRequest) (response ListFolderTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFolderTags, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListFolderTagsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFolderTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFolderTagsResponse")
	}
	return
}

// listFolderTags implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listFolderTags(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}/tags")
	if err != nil {
		return nil, err
	}

	var response ListFolderTagsResponse
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

// ListFolders Returns a list of all folders.
func (client DataCatalogClient) ListFolders(ctx context.Context, request ListFoldersRequest) (response ListFoldersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFolders, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListFoldersResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFoldersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFoldersResponse")
	}
	return
}

// listFolders implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listFolders(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders")
	if err != nil {
		return nil, err
	}

	var response ListFoldersResponse
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

// ListGlossaries Returns a list of all Glossaries within a data catalog.
func (client DataCatalogClient) ListGlossaries(ctx context.Context, request ListGlossariesRequest) (response ListGlossariesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGlossaries, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListGlossariesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGlossariesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGlossariesResponse")
	}
	return
}

// listGlossaries implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listGlossaries(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries")
	if err != nil {
		return nil, err
	}

	var response ListGlossariesResponse
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

// ListGlossaryPermissions Returns a list of Glossary Permissions.
func (client DataCatalogClient) ListGlossaryPermissions(ctx context.Context, request ListGlossaryPermissionsRequest) (response ListGlossaryPermissionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGlossaryPermissions, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListGlossaryPermissionsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGlossaryPermissionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGlossaryPermissionsResponse")
	}
	return
}

// listGlossaryPermissions implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listGlossaryPermissions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries/{glossaryKey}/permissions")
	if err != nil {
		return nil, err
	}

	var response ListGlossaryPermissionsResponse
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

// ListGlossaryTerms Returns a list of all terms within a glossary.
func (client DataCatalogClient) ListGlossaryTerms(ctx context.Context, request ListGlossaryTermsRequest) (response ListGlossaryTermsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGlossaryTerms, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListGlossaryTermsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGlossaryTermsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGlossaryTermsResponse")
	}
	return
}

// listGlossaryTerms implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listGlossaryTerms(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms")
	if err != nil {
		return nil, err
	}

	var response ListGlossaryTermsResponse
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

// ListJobDefinitionPermissions Returns a list of job definition permissions.
func (client DataCatalogClient) ListJobDefinitionPermissions(ctx context.Context, request ListJobDefinitionPermissionsRequest) (response ListJobDefinitionPermissionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobDefinitionPermissions, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListJobDefinitionPermissionsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobDefinitionPermissionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobDefinitionPermissionsResponse")
	}
	return
}

// listJobDefinitionPermissions implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobDefinitionPermissions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobDefinitions/{jobDefinitionKey}/permissions")
	if err != nil {
		return nil, err
	}

	var response ListJobDefinitionPermissionsResponse
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

// ListJobDefinitions Returns a list of Job Definitions within a data catalog.
func (client DataCatalogClient) ListJobDefinitions(ctx context.Context, request ListJobDefinitionsRequest) (response ListJobDefinitionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobDefinitions, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListJobDefinitionsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobDefinitionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobDefinitionsResponse")
	}
	return
}

// listJobDefinitions implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobDefinitions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobDefinitions")
	if err != nil {
		return nil, err
	}

	var response ListJobDefinitionsResponse
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

// ListJobExecutions Returns a list of Job Executions for a Job.
func (client DataCatalogClient) ListJobExecutions(ctx context.Context, request ListJobExecutionsRequest) (response ListJobExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListJobExecutionsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobExecutionsResponse")
	}
	return
}

// listJobExecutions implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobExecutions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions")
	if err != nil {
		return nil, err
	}

	var response ListJobExecutionsResponse
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

// ListJobLogs Returns a list of job logs.
func (client DataCatalogClient) ListJobLogs(ctx context.Context, request ListJobLogsRequest) (response ListJobLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobLogs, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListJobLogsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobLogsResponse")
	}
	return
}

// listJobLogs implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}/logs")
	if err != nil {
		return nil, err
	}

	var response ListJobLogsResponse
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

// ListJobMetrics Returns a list of job metrics.
func (client DataCatalogClient) ListJobMetrics(ctx context.Context, request ListJobMetricsRequest) (response ListJobMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListJobMetricsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobMetricsResponse")
	}
	return
}

// listJobMetrics implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobMetrics(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}/metrics")
	if err != nil {
		return nil, err
	}

	var response ListJobMetricsResponse
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

// ListJobs Returns a list of Jobs within a data catalog.
func (client DataCatalogClient) ListJobs(ctx context.Context, request ListJobsRequest) (response ListJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobs, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListJobsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobsResponse")
	}
	return
}

// listJobs implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs")
	if err != nil {
		return nil, err
	}

	var response ListJobsResponse
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

// ListSearchResults Returns a list of search results within a data catalog.
func (client DataCatalogClient) ListSearchResults(ctx context.Context, request ListSearchResultsRequest) (response ListSearchResultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSearchResults, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListSearchResultsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSearchResultsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSearchResultsResponse")
	}
	return
}

// listSearchResults implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listSearchResults(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/search")
	if err != nil {
		return nil, err
	}

	var response ListSearchResultsResponse
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

// ListTypes Returns a list of all Types within a data catalog.
func (client DataCatalogClient) ListTypes(ctx context.Context, request ListTypesRequest) (response ListTypesResponse, err error) {
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
func (client DataCatalogClient) listTypes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/types")
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

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
func (client DataCatalogClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListWorkRequestErrorsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
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

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
func (client DataCatalogClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListWorkRequestLogsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
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

// ListWorkRequests Lists the work requests in a compartment.
func (client DataCatalogClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListWorkRequestsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
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

// UpdateAttribute Update a specific Data Asset Attribute.
func (client DataCatalogClient) UpdateAttribute(ctx context.Context, request UpdateAttributeRequest) (response UpdateAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateAttributeResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAttributeResponse")
	}
	return
}

// updateAttribute implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateAttribute(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateAttributeResponse
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

// UpdateCatalog Updates the Catalog
func (client DataCatalogClient) UpdateCatalog(ctx context.Context, request UpdateCatalogRequest) (response UpdateCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateCatalogResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCatalogResponse")
	}
	return
}

// updateCatalog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateCatalog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}")
	if err != nil {
		return nil, err
	}

	var response UpdateCatalogResponse
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

// UpdateConnection Update a specific Connection of a Data Asset.
func (client DataCatalogClient) UpdateConnection(ctx context.Context, request UpdateConnectionRequest) (response UpdateConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConnection, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateConnectionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConnectionResponse")
	}
	return
}

// updateConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateConnection(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections/{connectionKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateConnectionResponse
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

// UpdateDataAsset Update a specific Data Asset identified by the given key
func (client DataCatalogClient) UpdateDataAsset(ctx context.Context, request UpdateDataAssetRequest) (response UpdateDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateDataAssetResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDataAssetResponse")
	}
	return
}

// updateDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateDataAsset(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateDataAssetResponse
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

// UpdateEntity Update a specific entity.
func (client DataCatalogClient) UpdateEntity(ctx context.Context, request UpdateEntityRequest) (response UpdateEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEntity, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateEntityResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEntityResponse")
	}
	return
}

// updateEntity implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateEntity(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateEntityResponse
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

// UpdateFolder Update a specific folder of a Data Asset.
func (client DataCatalogClient) UpdateFolder(ctx context.Context, request UpdateFolderRequest) (response UpdateFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFolder, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateFolderResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFolderResponse")
	}
	return
}

// updateFolder implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateFolder(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateFolderResponse
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

// UpdateGlossary Update a specific glossary identified by the given key
func (client DataCatalogClient) UpdateGlossary(ctx context.Context, request UpdateGlossaryRequest) (response UpdateGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateGlossaryResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGlossaryResponse")
	}
	return
}

// updateGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateGlossary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/glossaries/{glossaryKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateGlossaryResponse
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

// UpdateGlossaryTerm Update a specific Glossary Term.
func (client DataCatalogClient) UpdateGlossaryTerm(ctx context.Context, request UpdateGlossaryTermRequest) (response UpdateGlossaryTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateGlossaryTerm, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateGlossaryTermResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGlossaryTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGlossaryTermResponse")
	}
	return
}

// updateGlossaryTerm implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateGlossaryTerm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateGlossaryTermResponse
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

// UpdateJob Update a specific job identified by the given key
func (client DataCatalogClient) UpdateJob(ctx context.Context, request UpdateJobRequest) (response UpdateJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateJob, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateJobResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateJobResponse")
	}
	return
}

// updateJob implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/jobs/{jobKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateJobResponse
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

// UpdateJobDefinition Update a specific job definition identified by the given key
func (client DataCatalogClient) UpdateJobDefinition(ctx context.Context, request UpdateJobDefinitionRequest) (response UpdateJobDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateJobDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateJobDefinitionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateJobDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateJobDefinitionResponse")
	}
	return
}

// updateJobDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateJobDefinition(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/jobDefinitions/{jobDefinitionKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateJobDefinitionResponse
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
