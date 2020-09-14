// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GgsControlPlane API
//
// The GoldenGate Control Plane API specifying the operations and metadata for creating and maintaining the control infrastructure
// related to operating an OCI native Oracle managed GoldenGate service.
//

package goldengate

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/common/auth"
	"net/http"
)

//GoldenGateClient a client for GoldenGate
type GoldenGateClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewGoldenGateClientWithConfigurationProvider Creates a new default GoldenGate client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewGoldenGateClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client GoldenGateClient, err error) {
	if provider, err := auth.GetGenericConfigurationProvider(configProvider); err == nil {
		if baseClient, err := common.NewClientWithConfig(provider); err == nil {
			return newGoldenGateClientFromBaseClient(baseClient, configProvider)
		}
	}

	return
}

// NewGoldenGateClientWithOboToken Creates a new default GoldenGate client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewGoldenGateClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client GoldenGateClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newGoldenGateClientFromBaseClient(baseClient, configProvider)
}

func newGoldenGateClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client GoldenGateClient, err error) {
	client = GoldenGateClient{BaseClient: baseClient}
	client.BasePath = "20200407"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *GoldenGateClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("goldengate", "https://ggs-control-plane-api.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *GoldenGateClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *GoldenGateClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeBackupCompartment Moves a Backup into a different compartment within the same tenancy.  When provided, If-Match is checked against ETag values of the resource.
// For information about moving resources between compartments, see Moving Resources Between Compartments (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client GoldenGateClient) ChangeBackupCompartment(ctx context.Context, request ChangeBackupCompartmentRequest) (response ChangeBackupCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeBackupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeBackupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeBackupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeBackupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeBackupCompartmentResponse")
	}
	return
}

// changeBackupCompartment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) changeBackupCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/backups/{backupId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeBackupCompartmentResponse
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

// ChangeDatabaseRegistrationCompartment Moves a DatabaseRegistration into a different compartment within the same tenancy. When provided, If-Match is checked against ETag values of the resource
// For information about moving resources between compartments, see Moving Resources Between Compartments (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client GoldenGateClient) ChangeDatabaseRegistrationCompartment(ctx context.Context, request ChangeDatabaseRegistrationCompartmentRequest) (response ChangeDatabaseRegistrationCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeDatabaseRegistrationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatabaseRegistrationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatabaseRegistrationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatabaseRegistrationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatabaseRegistrationCompartmentResponse")
	}
	return
}

// changeDatabaseRegistrationCompartment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) changeDatabaseRegistrationCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseRegistrations/{databaseRegistrationId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeDatabaseRegistrationCompartmentResponse
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

// ChangeDeploymentCompartment Moves a Deployment into a different compartment within the same tenancy.  When provided, If-Match is checked against ETag values of the resource.
// For information about moving resources between compartments, see Moving Resources Between Compartments (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client GoldenGateClient) ChangeDeploymentCompartment(ctx context.Context, request ChangeDeploymentCompartmentRequest) (response ChangeDeploymentCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeDeploymentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDeploymentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDeploymentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDeploymentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDeploymentCompartmentResponse")
	}
	return
}

// changeDeploymentCompartment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) changeDeploymentCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeDeploymentCompartmentResponse
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

// CreateBackup Creates a new Backup.
func (client GoldenGateClient) CreateBackup(ctx context.Context, request CreateBackupRequest) (response CreateBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBackupResponse")
	}
	return
}

// createBackup implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) createBackup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/backups")
	if err != nil {
		return nil, err
	}

	var response CreateBackupResponse
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

// CreateDatabaseRegistration Creates a new DatabaseRegistration.
func (client GoldenGateClient) CreateDatabaseRegistration(ctx context.Context, request CreateDatabaseRegistrationRequest) (response CreateDatabaseRegistrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDatabaseRegistration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseRegistrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseRegistrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseRegistrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseRegistrationResponse")
	}
	return
}

// createDatabaseRegistration implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) createDatabaseRegistration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseRegistrations")
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseRegistrationResponse
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

// CreateDeployment Creates a new Deployment.
func (client GoldenGateClient) CreateDeployment(ctx context.Context, request CreateDeploymentRequest) (response CreateDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDeploymentResponse")
	}
	return
}

// createDeployment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) createDeployment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments")
	if err != nil {
		return nil, err
	}

	var response CreateDeploymentResponse
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

// DeleteBackup Deletes a Backup.
func (client GoldenGateClient) DeleteBackup(ctx context.Context, request DeleteBackupRequest) (response DeleteBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBackupResponse")
	}
	return
}

// deleteBackup implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) deleteBackup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/backups/{backupId}")
	if err != nil {
		return nil, err
	}

	var response DeleteBackupResponse
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

// DeleteDatabaseRegistration Deletes a DatabaseRegistration.
func (client GoldenGateClient) DeleteDatabaseRegistration(ctx context.Context, request DeleteDatabaseRegistrationRequest) (response DeleteDatabaseRegistrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabaseRegistration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseRegistrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseRegistrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseRegistrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseRegistrationResponse")
	}
	return
}

// deleteDatabaseRegistration implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) deleteDatabaseRegistration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseRegistrations/{databaseRegistrationId}")
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseRegistrationResponse
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

// DeleteDeployment Deletes a Deployment.
func (client GoldenGateClient) DeleteDeployment(ctx context.Context, request DeleteDeploymentRequest) (response DeleteDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDeploymentResponse")
	}
	return
}

// deleteDeployment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) deleteDeployment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/deployments/{deploymentId}")
	if err != nil {
		return nil, err
	}

	var response DeleteDeploymentResponse
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

// GetBackup Retrieves a Backup.
func (client GoldenGateClient) GetBackup(ctx context.Context, request GetBackupRequest) (response GetBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBackupResponse")
	}
	return
}

// getBackup implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) getBackup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/backups/{backupId}")
	if err != nil {
		return nil, err
	}

	var response GetBackupResponse
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

// GetDatabaseRegistration Retrieves a DatabaseRegistration.
func (client GoldenGateClient) GetDatabaseRegistration(ctx context.Context, request GetDatabaseRegistrationRequest) (response GetDatabaseRegistrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseRegistration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseRegistrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseRegistrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseRegistrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseRegistrationResponse")
	}
	return
}

// getDatabaseRegistration implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) getDatabaseRegistration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseRegistrations/{databaseRegistrationId}")
	if err != nil {
		return nil, err
	}

	var response GetDatabaseRegistrationResponse
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

// GetDeployment Retrieves a Deployment.
func (client GoldenGateClient) GetDeployment(ctx context.Context, request GetDeploymentRequest) (response GetDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDeploymentResponse")
	}
	return
}

// getDeployment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) getDeployment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployments/{deploymentId}")
	if err != nil {
		return nil, err
	}

	var response GetDeploymentResponse
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

// GetWorkRequest Gets the workrequest with the given OCID.
func (client GoldenGateClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
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
func (client GoldenGateClient) getWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// ListBackups Lists the Backups for a compartment.
func (client GoldenGateClient) ListBackups(ctx context.Context, request ListBackupsRequest) (response ListBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBackups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBackupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBackupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBackupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBackupsResponse")
	}
	return
}

// listBackups implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listBackups(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/backups")
	if err != nil {
		return nil, err
	}

	var response ListBackupsResponse
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

// ListDatabaseRegistrations Lists the DatabaseRegistrations for a compartment.
func (client GoldenGateClient) ListDatabaseRegistrations(ctx context.Context, request ListDatabaseRegistrationsRequest) (response ListDatabaseRegistrationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseRegistrations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseRegistrationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseRegistrationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseRegistrationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseRegistrationsResponse")
	}
	return
}

// listDatabaseRegistrations implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listDatabaseRegistrations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseRegistrations")
	if err != nil {
		return nil, err
	}

	var response ListDatabaseRegistrationsResponse
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

// ListDeployments Lists the Deployments for a compartment.
func (client GoldenGateClient) ListDeployments(ctx context.Context, request ListDeploymentsRequest) (response ListDeploymentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDeployments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDeploymentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDeploymentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDeploymentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDeploymentsResponse")
	}
	return
}

// listDeployments implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listDeployments(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployments")
	if err != nil {
		return nil, err
	}

	var response ListDeploymentsResponse
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

// ListWorkRequestErrors Lists work request errors.
func (client GoldenGateClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
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
func (client GoldenGateClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// ListWorkRequestLogs Lists work request logs.
func (client GoldenGateClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
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
func (client GoldenGateClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// ListWorkRequests Lists work requests in a compartment.
func (client GoldenGateClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
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
func (client GoldenGateClient) listWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// StartDeployment Starts a Deployment.  When provided, If-Match is checked against ETag values of the resource.
func (client GoldenGateClient) StartDeployment(ctx context.Context, request StartDeploymentRequest) (response StartDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.startDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartDeploymentResponse")
	}
	return
}

// startDeployment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) startDeployment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/start")
	if err != nil {
		return nil, err
	}

	var response StartDeploymentResponse
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

// StopDeployment Stops a Deployment.  When provided, If-Match is checked against ETag values of the resource.
func (client GoldenGateClient) StopDeployment(ctx context.Context, request StopDeploymentRequest) (response StopDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.stopDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopDeploymentResponse")
	}
	return
}

// stopDeployment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) stopDeployment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/stop")
	if err != nil {
		return nil, err
	}

	var response StopDeploymentResponse
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

// SynchronizeDatabaseRegistration Cause the DatabaseRegistration to notify each Deployment in the same compartment of its latest properties.  This action may be useful to ensure each Deployment has
// the latest values which it may have missed if it were not availabile for its update when the DatabaseRegistrion was updated previously.
// As this is an asyncronous action, the workrequest is available in the header, and may be used to follow the progress of the synchronization.
func (client GoldenGateClient) SynchronizeDatabaseRegistration(ctx context.Context, request SynchronizeDatabaseRegistrationRequest) (response SynchronizeDatabaseRegistrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.synchronizeDatabaseRegistration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SynchronizeDatabaseRegistrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SynchronizeDatabaseRegistrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SynchronizeDatabaseRegistrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SynchronizeDatabaseRegistrationResponse")
	}
	return
}

// synchronizeDatabaseRegistration implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) synchronizeDatabaseRegistration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseRegistrations/{databaseRegistrationId}/actions/synchronize")
	if err != nil {
		return nil, err
	}

	var response SynchronizeDatabaseRegistrationResponse
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

// UpdateBackup Modifies a Backup.
func (client GoldenGateClient) UpdateBackup(ctx context.Context, request UpdateBackupRequest) (response UpdateBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBackupResponse")
	}
	return
}

// updateBackup implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) updateBackup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/backups/{backupId}")
	if err != nil {
		return nil, err
	}

	var response UpdateBackupResponse
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

// UpdateDatabaseRegistration Updates the DatabaseRegistration.
func (client GoldenGateClient) UpdateDatabaseRegistration(ctx context.Context, request UpdateDatabaseRegistrationRequest) (response UpdateDatabaseRegistrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseRegistration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseRegistrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseRegistrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseRegistrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseRegistrationResponse")
	}
	return
}

// updateDatabaseRegistration implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) updateDatabaseRegistration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseRegistrations/{databaseRegistrationId}")
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseRegistrationResponse
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

// UpdateDeployment Modifies a Deployment.
func (client GoldenGateClient) UpdateDeployment(ctx context.Context, request UpdateDeploymentRequest) (response UpdateDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDeploymentResponse")
	}
	return
}

// updateDeployment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) updateDeployment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/deployments/{deploymentId}")
	if err != nil {
		return nil, err
	}

	var response UpdateDeploymentResponse
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
