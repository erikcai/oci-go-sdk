// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Key Management Service API
//
// APIs for managing and performing operations with keys and vaults.
//

package kms

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//KmsVaultClient a client for KmsVault
type KmsVaultClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewKmsVaultClientWithConfigurationProvider Creates a new default KmsVault client with the given configuration provider.
// the configuration provider will be used for the default signer
func NewKmsVaultClientWithConfigurationProvider(configProvider common.ConfigurationProvider, endpoint string) (client KmsVaultClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = KmsVaultClient{BaseClient: baseClient}
	client.BasePath = "20180201"
	client.Host = endpoint
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *KmsVaultClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *KmsVaultClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateKey Creates a new Key.  TODO description
func (client KmsVaultClient) CreateKey(ctx context.Context, request CreateKeyRequest) (response CreateKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createKey, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(CreateKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateKeyResponse")
	}
	return
}

// createKey implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) createKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/keys")
	if err != nil {
		return nil, err
	}

	var response CreateKeyResponse
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

// DisableKey null
func (client KmsVaultClient) DisableKey(ctx context.Context, request DisableKeyRequest) (response DisableKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.disableKey, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(DisableKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableKeyResponse")
	}
	return
}

// disableKey implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) disableKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/keys/{keyId}/actions/disable")
	if err != nil {
		return nil, err
	}

	var response DisableKeyResponse
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

// EnableKey null
func (client KmsVaultClient) EnableKey(ctx context.Context, request EnableKeyRequest) (response EnableKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.enableKey, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(EnableKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableKeyResponse")
	}
	return
}

// enableKey implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) enableKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/keys/{keyId}/actions/enable")
	if err != nil {
		return nil, err
	}

	var response EnableKeyResponse
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

// GetKey Get information about the specified Key.  TODO description
func (client KmsVaultClient) GetKey(ctx context.Context, request GetKeyRequest) (response GetKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getKey, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(GetKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetKeyResponse")
	}
	return
}

// getKey implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) getKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/keys/{keyId}")
	if err != nil {
		return nil, err
	}

	var response GetKeyResponse
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

// ListKeyVersions null
func (client KmsVaultClient) ListKeyVersions(ctx context.Context, request ListKeyVersionsRequest) (response ListKeyVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listKeyVersions, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(ListKeyVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListKeyVersionsResponse")
	}
	return
}

// listKeyVersions implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) listKeyVersions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/keys/{keyId}/keyVersions")
	if err != nil {
		return nil, err
	}

	var response ListKeyVersionsResponse
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

// ListKeys Lists the Keys in this vault in this compartment.
func (client KmsVaultClient) ListKeys(ctx context.Context, request ListKeysRequest) (response ListKeysResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listKeys, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(ListKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListKeysResponse")
	}
	return
}

// listKeys implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) listKeys(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/keys")
	if err != nil {
		return nil, err
	}

	var response ListKeysResponse
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

// RotateKey null
func (client KmsVaultClient) RotateKey(ctx context.Context, request RotateKeyRequest) (response RotateKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.rotateKey, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(RotateKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateKeyResponse")
	}
	return
}

// rotateKey implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) rotateKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/keys/{keyId}/actions/rotate")
	if err != nil {
		return nil, err
	}

	var response RotateKeyResponse
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

// UpdateKey Updates the properties of a Key, such as the displayName.
func (client KmsVaultClient) UpdateKey(ctx context.Context, request UpdateKeyRequest) (response UpdateKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateKey, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateKeyResponse")
	}
	return
}

// updateKey implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) updateKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/keys/{keyId}")
	if err != nil {
		return nil, err
	}

	var response UpdateKeyResponse
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
