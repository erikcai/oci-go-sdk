// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Key Management Service API
// 
 // APIs for managing and performing operations with keys and vaults.
//

package kms

import(
    "github.com/oracle/oci-go-sdk/common"
    "context"
    "fmt"
    "net/http"
)

//KmsProvisioningClient a client for KmsProvisioning
type KmsProvisioningClient struct {
    common.BaseClient
    config *common.ConfigurationProvider
}


// NewKmsProvisioningClientWithConfigurationProvider Creates a new default KmsProvisioning client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewKmsProvisioningClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client KmsProvisioningClient, err error){
    baseClient, err := common.NewClientWithConfig(configProvider)
    if err != nil {
        return
    }

    client = KmsProvisioningClient{BaseClient: baseClient}
    client.BasePath = "20180201"
    err = client.setConfigurationProvider(configProvider)
    return
}

// SetRegion overrides the region of this client.
func (client *KmsProvisioningClient) SetRegion(region string)  {
    client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "kms", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *KmsProvisioningClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *KmsProvisioningClient) ConfigurationProvider() *common.ConfigurationProvider {
    return client.config
}






 // CreateVault Create a new vault.  TODO description
func(client KmsProvisioningClient) CreateVault(ctx context.Context, request CreateVaultRequest) (response CreateVaultResponse, err error) {
    var ociResponse common.OCIResponse
    policy := common.NoRetryPolicy()
    if request.RetryPolicy() != nil {
        policy = *request.RetryPolicy()
    }
    ociResponse, err = common.Retry(ctx, request, client.createVault, policy)
    if err != nil {
        return
    }
    if convertedResponse, ok := ociResponse.(CreateVaultResponse); ok {
        response = convertedResponse
    } else {
        err = fmt.Errorf("failed to convert OCIResponse into CreateVaultResponse")
    }
    return
}

// createVault implements the OCIOperation interface (enables retrying operations)
func(client KmsProvisioningClient) createVault(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
    httpRequest, err := request.HTTPRequest(http.MethodPost, "/vaults")
    if err != nil {
        return nil, err
    }

    var response CreateVaultResponse
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




 // GetVault Get information about the specified Vault.  TODO description
func(client KmsProvisioningClient) GetVault(ctx context.Context, request GetVaultRequest) (response GetVaultResponse, err error) {
    var ociResponse common.OCIResponse
    policy := common.NoRetryPolicy()
    if request.RetryPolicy() != nil {
        policy = *request.RetryPolicy()
    }
    ociResponse, err = common.Retry(ctx, request, client.getVault, policy)
    if err != nil {
        return
    }
    if convertedResponse, ok := ociResponse.(GetVaultResponse); ok {
        response = convertedResponse
    } else {
        err = fmt.Errorf("failed to convert OCIResponse into GetVaultResponse")
    }
    return
}

// getVault implements the OCIOperation interface (enables retrying operations)
func(client KmsProvisioningClient) getVault(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
    httpRequest, err := request.HTTPRequest(http.MethodGet, "/vaults/{vaultId}")
    if err != nil {
        return nil, err
    }

    var response GetVaultResponse
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




 // ListVaults Lists the Vaults in the specified compartment.
func(client KmsProvisioningClient) ListVaults(ctx context.Context, request ListVaultsRequest) (response ListVaultsResponse, err error) {
    var ociResponse common.OCIResponse
    policy := common.NoRetryPolicy()
    if request.RetryPolicy() != nil {
        policy = *request.RetryPolicy()
    }
    ociResponse, err = common.Retry(ctx, request, client.listVaults, policy)
    if err != nil {
        return
    }
    if convertedResponse, ok := ociResponse.(ListVaultsResponse); ok {
        response = convertedResponse
    } else {
        err = fmt.Errorf("failed to convert OCIResponse into ListVaultsResponse")
    }
    return
}

// listVaults implements the OCIOperation interface (enables retrying operations)
func(client KmsProvisioningClient) listVaults(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
    httpRequest, err := request.HTTPRequest(http.MethodGet, "/vaults")
    if err != nil {
        return nil, err
    }

    var response ListVaultsResponse
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




 // UpdateVault Updates the properties of a Vault, such as the displayName.
func(client KmsProvisioningClient) UpdateVault(ctx context.Context, request UpdateVaultRequest) (response UpdateVaultResponse, err error) {
    var ociResponse common.OCIResponse
    policy := common.NoRetryPolicy()
    if request.RetryPolicy() != nil {
        policy = *request.RetryPolicy()
    }
    ociResponse, err = common.Retry(ctx, request, client.updateVault, policy)
    if err != nil {
        return
    }
    if convertedResponse, ok := ociResponse.(UpdateVaultResponse); ok {
        response = convertedResponse
    } else {
        err = fmt.Errorf("failed to convert OCIResponse into UpdateVaultResponse")
    }
    return
}

// updateVault implements the OCIOperation interface (enables retrying operations)
func(client KmsProvisioningClient) updateVault(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
    httpRequest, err := request.HTTPRequest(http.MethodPut, "/vaults/{vaultId}")
    if err != nil {
        return nil, err
    }

    var response UpdateVaultResponse
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

