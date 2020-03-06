// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Key Management Service API
//
// API for managing and performing operations with keys and vaults.
//

package keymanagement

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
// the configuration provider will be used for the default signer as well as reading the region
func NewKmsVaultClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client KmsVaultClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newKmsVaultClientFromBaseClient(baseClient, configProvider)
}

// NewKmsVaultClientWithOboToken Creates a new default KmsVault client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewKmsVaultClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client KmsVaultClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newKmsVaultClientFromBaseClient(baseClient, configProvider)
}

func newKmsVaultClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client KmsVaultClient, err error) {
	client = KmsVaultClient{BaseClient: baseClient}
	client.BasePath = ""
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *KmsVaultClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("kms", "https://kms.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *KmsVaultClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *KmsVaultClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// BackupVault Backup an encrypted binary payload that contains only the metadata of the vault so it can be restored.
func (client KmsVaultClient) BackupVault(ctx context.Context, request BackupVaultRequest) (response BackupVaultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.backupVault, policy)
	if err != nil {
		if ociResponse != nil {
			response = BackupVaultResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BackupVaultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BackupVaultResponse")
	}
	return
}

// backupVault implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) backupVault(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20180608/vaults/{vaultId}/actions/backup")
	if err != nil {
		return nil, err
	}

	var response BackupVaultResponse
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

// CancelVaultDeletion Cancels the scheduled deletion of the specified vault. Canceling a scheduled deletion
// restores the vault and all keys in it to their respective states from before their
// scheduled deletion. All keys that were scheduled for deletion prior to vault
// deletion retain their lifecycle state and time of deletion.
// As a provisioning operation, this call is subject to a Key Management limit that applies to
// the total number of requests across all provisioning write operations. Key Management might
// throttle this call to reject an otherwise valid request when the total rate of provisioning
// write operations exceeds 10 requests per second for a given tenancy.
func (client KmsVaultClient) CancelVaultDeletion(ctx context.Context, request CancelVaultDeletionRequest) (response CancelVaultDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.cancelVaultDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			response = CancelVaultDeletionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelVaultDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelVaultDeletionResponse")
	}
	return
}

// cancelVaultDeletion implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) cancelVaultDeletion(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20180608/vaults/{vaultId}/actions/cancelDeletion")
	if err != nil {
		return nil, err
	}

	var response CancelVaultDeletionResponse
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

// ChangeVaultCompartment Moves a vault into a different compartment within the same tenancy. For information about
// moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// When provided, if-match is checked against the ETag values of the resource.
// As a provisioning operation, this call is subject to a Key Management limit that applies to
// the total number of requests across all provisioning write operations. Key Management might
// throttle this call to reject an otherwise valid request when the total rate of provisioning
// write operations exceeds 10 requests per second for a given tenancy.
func (client KmsVaultClient) ChangeVaultCompartment(ctx context.Context, request ChangeVaultCompartmentRequest) (response ChangeVaultCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeVaultCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			response = ChangeVaultCompartmentResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeVaultCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeVaultCompartmentResponse")
	}
	return
}

// changeVaultCompartment implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) changeVaultCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20180608/vaults/{vaultId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeVaultCompartmentResponse
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

// CreateVault Creates a new vault. The type of vault you create determines key placement, pricing, and
// available options. Options include storage isolation, a dedicated service endpoint instead
// of a shared service endpoint for API calls, and either a dedicated hardware security module
// (HSM) or a multitenant HSM.
// As a provisioning operation, this call is subject to a Key Management limit that applies to
// the total number of requests across all provisioning write operations. Key Management might
// throttle this call to reject an otherwise valid request when the total rate of provisioning
// write operations exceeds 10 requests per second for a given tenancy.
func (client KmsVaultClient) CreateVault(ctx context.Context, request CreateVaultRequest) (response CreateVaultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createVault, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateVaultResponse{RawResponse: ociResponse.HTTPResponse()}
		}
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
func (client KmsVaultClient) createVault(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20180608/vaults")
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

// GetVault Gets the specified vault's configuration information.
// As a provisioning operation, this call is subject to a Key Management limit that applies to
// the total number of requests across all provisioning read operations. Key Management might
// throttle this call to reject an otherwise valid request when the total rate of provisioning
// read operations exceeds 10 requests per second for a given tenancy.
func (client KmsVaultClient) GetVault(ctx context.Context, request GetVaultRequest) (response GetVaultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVault, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetVaultResponse{RawResponse: ociResponse.HTTPResponse()}
		}
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
func (client KmsVaultClient) getVault(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20180608/vaults/{vaultId}")
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

// GetVaultUsage Gets the count of keys and key versions in the specified vault to calculate usage against service limits.
func (client KmsVaultClient) GetVaultUsage(ctx context.Context, request GetVaultUsageRequest) (response GetVaultUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVaultUsage, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetVaultUsageResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVaultUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVaultUsageResponse")
	}
	return
}

// getVaultUsage implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) getVaultUsage(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20180608/vaults/{vaultId}/usage")
	if err != nil {
		return nil, err
	}

	var response GetVaultUsageResponse
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

// ListVaults Lists the vaults in the specified compartment.
// As a provisioning operation, this call is subject to a Key Management limit that applies to
// the total number of requests across all provisioning read operations. Key Management might
// throttle this call to reject an otherwise valid request when the total rate of provisioning
// read operations exceeds 10 requests per second for a given tenancy.
func (client KmsVaultClient) ListVaults(ctx context.Context, request ListVaultsRequest) (response ListVaultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVaults, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListVaultsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
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
func (client KmsVaultClient) listVaults(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20180608/vaults")
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

// RestoreVaultFromFile Restore a vault from an encrypted binary backup.  If the vault with the OCID already exists, this operation will return a response with a 409 HTTP status code.
func (client KmsVaultClient) RestoreVaultFromFile(ctx context.Context, request RestoreVaultFromFileRequest) (response RestoreVaultFromFileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.restoreVaultFromFile, policy)
	if err != nil {
		if ociResponse != nil {
			response = RestoreVaultFromFileResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestoreVaultFromFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestoreVaultFromFileResponse")
	}
	return
}

// restoreVaultFromFile implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) restoreVaultFromFile(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20180608/vaults/actions/restoreFromFile")
	if err != nil {
		return nil, err
	}

	var response RestoreVaultFromFileResponse
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

// RestoreVaultFromObjectStore Restore a vault from an encrypted binary backup stored in object store.  If the vault with the OCID already exists, this operation will return a response with a 409 HTTP status code.
func (client KmsVaultClient) RestoreVaultFromObjectStore(ctx context.Context, request RestoreVaultFromObjectStoreRequest) (response RestoreVaultFromObjectStoreResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.restoreVaultFromObjectStore, policy)
	if err != nil {
		if ociResponse != nil {
			response = RestoreVaultFromObjectStoreResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestoreVaultFromObjectStoreResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestoreVaultFromObjectStoreResponse")
	}
	return
}

// restoreVaultFromObjectStore implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) restoreVaultFromObjectStore(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20180608/vaults/actions/restoreFromObjectStore")
	if err != nil {
		return nil, err
	}

	var response RestoreVaultFromObjectStoreResponse
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

// ScheduleVaultDeletion Schedules the deletion of the specified vault. This sets the lifecycle state of the vault and all keys in it
// that are not already scheduled for deletion to `PENDING_DELETION` and then deletes them after the
// retention period ends. The lifecycle state and time of deletion for keys already scheduled for deletion won't
// change. If any keys in the vault are scheduled to be deleted after the specified time of
// deletion for the vault, the call is rejected with the error code 409.
// As a provisioning operation, this call is subject to a Key Management limit that applies to
// the total number of requests across all provisioning write operations. Key Management might
// throttle this call to reject an otherwise valid request when the total rate of provisioning
// write operations exceeds 10 requests per second for a given tenancy.
func (client KmsVaultClient) ScheduleVaultDeletion(ctx context.Context, request ScheduleVaultDeletionRequest) (response ScheduleVaultDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.scheduleVaultDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			response = ScheduleVaultDeletionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScheduleVaultDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScheduleVaultDeletionResponse")
	}
	return
}

// scheduleVaultDeletion implements the OCIOperation interface (enables retrying operations)
func (client KmsVaultClient) scheduleVaultDeletion(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20180608/vaults/{vaultId}/actions/scheduleDeletion")
	if err != nil {
		return nil, err
	}

	var response ScheduleVaultDeletionResponse
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

// UpdateVault Updates the properties of a vault. Specifically, you can update the
// `displayName`, `freeformTags`, and `definedTags` properties. Furthermore,
// the vault must be in an ACTIVE or CREATING state to be updated.
// As a provisioning operation, this call is subject to a Key Management limit that applies to
// the total number of requests across all provisioning write operations. Key Management might
// throttle this call to reject an otherwise valid request when the total rate of provisioning
// write operations exceeds 10 requests per second for a given tenancy.
func (client KmsVaultClient) UpdateVault(ctx context.Context, request UpdateVaultRequest) (response UpdateVaultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVault, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateVaultResponse{RawResponse: ociResponse.HTTPResponse()}
		}
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
func (client KmsVaultClient) updateVault(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/20180608/vaults/{vaultId}")
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
