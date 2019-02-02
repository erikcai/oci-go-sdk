// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//DatabaseClient a client for Database
type DatabaseClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDatabaseClientWithConfigurationProvider Creates a new default Database client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDatabaseClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DatabaseClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = DatabaseClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DatabaseClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("database", "https://database.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DatabaseClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DatabaseClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CompleteExternalBackupJob Changes the status of the standalone backup resource to `ACTIVE` after the backup is created from the on-premises database and placed in Oracle Cloud Infrastructure Object Storage.
// **Note:** This API is used by an Oracle Cloud Infrastructure Python script that is packaged with the Oracle Cloud Infrastructure CLI. Oracle recommends that you use the script instead using the API directly. See Migrating an On-Premises Database to Oracle Cloud Infrastructure by Creating a Backup in the Cloud (https://docs.us-phoenix-1.oraclecloud.com/Content/Database/Tasks/mig-onprembackup.htm) for more information.
func (client DatabaseClient) CompleteExternalBackupJob(ctx context.Context, request CompleteExternalBackupJobRequest) (response CompleteExternalBackupJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.completeExternalBackupJob, policy)
	if err != nil {
		if ociResponse != nil {
			response = CompleteExternalBackupJobResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CompleteExternalBackupJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CompleteExternalBackupJobResponse")
	}
	return
}

// completeExternalBackupJob implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) completeExternalBackupJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalBackupJobs/{backupId}/actions/complete")
	if err != nil {
		return nil, err
	}

	var response CompleteExternalBackupJobResponse
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

// CreateAutonomousContainerDatabase Create a new Autonomous Container Database in the specified Autonomous Exadata Infrastructure.
func (client DatabaseClient) CreateAutonomousContainerDatabase(ctx context.Context, request CreateAutonomousContainerDatabaseRequest) (response CreateAutonomousContainerDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAutonomousContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateAutonomousContainerDatabaseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAutonomousContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAutonomousContainerDatabaseResponse")
	}
	return
}

// createAutonomousContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createAutonomousContainerDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases")
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousContainerDatabaseResponse
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

// CreateAutonomousDataWarehouse Creates a new Autonomous Data Warehouse.
func (client DatabaseClient) CreateAutonomousDataWarehouse(ctx context.Context, request CreateAutonomousDataWarehouseRequest) (response CreateAutonomousDataWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAutonomousDataWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateAutonomousDataWarehouseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAutonomousDataWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAutonomousDataWarehouseResponse")
	}
	return
}

// createAutonomousDataWarehouse implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createAutonomousDataWarehouse(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDataWarehouses")
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousDataWarehouseResponse
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

// CreateAutonomousDataWarehouseBackup Creates a new Autonomous Data Warehouse backup for the specified database based on the provided request parameters.
func (client DatabaseClient) CreateAutonomousDataWarehouseBackup(ctx context.Context, request CreateAutonomousDataWarehouseBackupRequest) (response CreateAutonomousDataWarehouseBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAutonomousDataWarehouseBackup, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateAutonomousDataWarehouseBackupResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAutonomousDataWarehouseBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAutonomousDataWarehouseBackupResponse")
	}
	return
}

// createAutonomousDataWarehouseBackup implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createAutonomousDataWarehouseBackup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDataWarehouseBackups")
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousDataWarehouseBackupResponse
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

// CreateAutonomousDatabase Creates a new Autonomous Database.
func (client DatabaseClient) CreateAutonomousDatabase(ctx context.Context, request CreateAutonomousDatabaseRequest) (response CreateAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateAutonomousDatabaseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAutonomousDatabaseResponse")
	}
	return
}

// createAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createAutonomousDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases")
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousDatabaseResponse
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

// CreateAutonomousDatabaseBackup Creates a new Autonomous Database backup for the specified database based on the provided request parameters.
func (client DatabaseClient) CreateAutonomousDatabaseBackup(ctx context.Context, request CreateAutonomousDatabaseBackupRequest) (response CreateAutonomousDatabaseBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAutonomousDatabaseBackup, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateAutonomousDatabaseBackupResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAutonomousDatabaseBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAutonomousDatabaseBackupResponse")
	}
	return
}

// createAutonomousDatabaseBackup implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createAutonomousDatabaseBackup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabaseBackups")
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousDatabaseBackupResponse
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

// CreateAutonomousPod Create a new Autonomous Pod in the specified Autonomous Exadata Infrastructure.
func (client DatabaseClient) CreateAutonomousPod(ctx context.Context, request CreateAutonomousPodRequest) (response CreateAutonomousPodResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAutonomousPod, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateAutonomousPodResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAutonomousPodResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAutonomousPodResponse")
	}
	return
}

// createAutonomousPod implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createAutonomousPod(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousPods")
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousPodResponse
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

// CreateBackup Creates a new backup in the specified database based on the request parameters you provide. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
func (client DatabaseClient) CreateBackup(ctx context.Context, request CreateBackupRequest) (response CreateBackupResponse, err error) {
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
			response = CreateBackupResponse{RawResponse: ociResponse.HTTPResponse()}
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
func (client DatabaseClient) createBackup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// CreateDataGuardAssociation Creates a new Data Guard association.  A Data Guard association represents the replication relationship between the
// specified database and a peer database. For more information, see Using Oracle Data Guard (https://docs.us-phoenix-1.oraclecloud.com/Content/Database/Tasks/usingdataguard.htm).
// All Oracle Cloud Infrastructure resources, including Data Guard associations, get an Oracle-assigned, unique ID
// called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID in the response.
// You can also retrieve a resource's OCID by using a List API operation on that resource type, or by viewing the
// resource in the Console. For more information, see
// Resource Identifiers (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
func (client DatabaseClient) CreateDataGuardAssociation(ctx context.Context, request CreateDataGuardAssociationRequest) (response CreateDataGuardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDataGuardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateDataGuardAssociationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataGuardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataGuardAssociationResponse")
	}
	return
}

// createDataGuardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createDataGuardAssociation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations")
	if err != nil {
		return nil, err
	}

	var response CreateDataGuardAssociationResponse
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

// CreateDbHome Creates a new database home in the specified DB system based on the request parameters you provide.
func (client DatabaseClient) CreateDbHome(ctx context.Context, request CreateDbHomeRequest) (response CreateDbHomeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDbHome, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateDbHomeResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDbHomeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDbHomeResponse")
	}
	return
}

// createDbHome implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createDbHome(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbHomes")
	if err != nil {
		return nil, err
	}

	var response CreateDbHomeResponse
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

// CreateExternalBackupJob Creates a new backup resource and returns the information the caller needs to back up an on-premises Oracle Database to Oracle Cloud Infrastructure.
// **Note:** This API is used by an Oracle Cloud Infrastructure Python script that is packaged with the Oracle Cloud Infrastructure CLI. Oracle recommends that you use the script instead using the API directly. See Migrating an On-Premises Database to Oracle Cloud Infrastructure by Creating a Backup in the Cloud (https://docs.us-phoenix-1.oraclecloud.com/Content/Database/Tasks/mig-onprembackup.htm) for more information.
func (client DatabaseClient) CreateExternalBackupJob(ctx context.Context, request CreateExternalBackupJobRequest) (response CreateExternalBackupJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createExternalBackupJob, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateExternalBackupJobResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalBackupJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalBackupJobResponse")
	}
	return
}

// createExternalBackupJob implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createExternalBackupJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalBackupJobs")
	if err != nil {
		return nil, err
	}

	var response CreateExternalBackupJobResponse
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

// DbNodeAction Performs one of the following power actions on the specified DB node:
// - start - power on
// - stop - power off
// - softreset - ACPI shutdown and power on
// - reset - power off and power on
// **Note:** Stopping a node affects billing differently, depending on the type of DB system:
// *Bare metal and Exadata DB systems* - The _stop_ state has no effect on the resources you consume.
// Billing continues for DB nodes that you stop, and related resources continue
// to apply against any relevant quotas. You must terminate the DB system
// (TerminateDbSystem)
// to remove its resources from billing and quotas.
// *Virtual machine DB systems* - Stopping a node stops billing for all OCPUs associated with that node, and billing resumes when you restart the node.
func (client DatabaseClient) DbNodeAction(ctx context.Context, request DbNodeActionRequest) (response DbNodeActionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.dbNodeAction, policy)
	if err != nil {
		if ociResponse != nil {
			response = DbNodeActionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DbNodeActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DbNodeActionResponse")
	}
	return
}

// dbNodeAction implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) dbNodeAction(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbNodes/{dbNodeId}")
	if err != nil {
		return nil, err
	}

	var response DbNodeActionResponse
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

// DeleteAutonomousDataWarehouse Deletes the specified Autonomous Data Warehouse.
func (client DatabaseClient) DeleteAutonomousDataWarehouse(ctx context.Context, request DeleteAutonomousDataWarehouseRequest) (response DeleteAutonomousDataWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAutonomousDataWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteAutonomousDataWarehouseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAutonomousDataWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAutonomousDataWarehouseResponse")
	}
	return
}

// deleteAutonomousDataWarehouse implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteAutonomousDataWarehouse(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousDataWarehouses/{autonomousDataWarehouseId}")
	if err != nil {
		return nil, err
	}

	var response DeleteAutonomousDataWarehouseResponse
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

// DeleteAutonomousDatabase Deletes the specified Autonomous Database.
func (client DatabaseClient) DeleteAutonomousDatabase(ctx context.Context, request DeleteAutonomousDatabaseRequest) (response DeleteAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteAutonomousDatabaseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAutonomousDatabaseResponse")
	}
	return
}

// deleteAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteAutonomousDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousDatabases/{autonomousDatabaseId}")
	if err != nil {
		return nil, err
	}

	var response DeleteAutonomousDatabaseResponse
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

// DeleteAutonomousDatabaseBackup Deletes manual backup. You cannot delete automatic backups using this API.
func (client DatabaseClient) DeleteAutonomousDatabaseBackup(ctx context.Context, request DeleteAutonomousDatabaseBackupRequest) (response DeleteAutonomousDatabaseBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAutonomousDatabaseBackup, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteAutonomousDatabaseBackupResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAutonomousDatabaseBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAutonomousDatabaseBackupResponse")
	}
	return
}

// deleteAutonomousDatabaseBackup implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteAutonomousDatabaseBackup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousDatabaseBackups/{autonomousDatabaseBackupId}")
	if err != nil {
		return nil, err
	}

	var response DeleteAutonomousDatabaseBackupResponse
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

// DeleteBackup Deletes a full backup. You cannot delete automatic backups using this API.
func (client DatabaseClient) DeleteBackup(ctx context.Context, request DeleteBackupRequest) (response DeleteBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBackup, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteBackupResponse{RawResponse: ociResponse.HTTPResponse()}
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
func (client DatabaseClient) deleteBackup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// DeleteDbHome Deletes a DB Home. The DB Home and its database data are local to the DB system and will be lost when it is deleted. Oracle recommends that you back up any data in the DB system prior to deleting it.
func (client DatabaseClient) DeleteDbHome(ctx context.Context, request DeleteDbHomeRequest) (response DeleteDbHomeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDbHome, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteDbHomeResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDbHomeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDbHomeResponse")
	}
	return
}

// deleteDbHome implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteDbHome(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbHomes/{dbHomeId}")
	if err != nil {
		return nil, err
	}

	var response DeleteDbHomeResponse
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

// FailoverAutonomousContainerDatabaseMissionCriticalAssociation Performs a failover to transition the standby Container Database identified by the autonomousContainerDatabaseId parameter into the specified Autonomous Container Database Mission Critical Association's primary role after the existing primary Container Database fails or becomes unreachable.
// A failover might result in data loss, depending on the protection mode in effect at the time of the primary Container Database failure.
func (client DatabaseClient) FailoverAutonomousContainerDatabaseMissionCriticalAssociation(ctx context.Context, request FailoverAutonomousContainerDatabaseMissionCriticalAssociationRequest) (response FailoverAutonomousContainerDatabaseMissionCriticalAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.failoverAutonomousContainerDatabaseMissionCriticalAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			response = FailoverAutonomousContainerDatabaseMissionCriticalAssociationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FailoverAutonomousContainerDatabaseMissionCriticalAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FailoverAutonomousContainerDatabaseMissionCriticalAssociationResponse")
	}
	return
}

// failoverAutonomousContainerDatabaseMissionCriticalAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) failoverAutonomousContainerDatabaseMissionCriticalAssociation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseMissionCriticalAssociations/{autonomousContainerDatabaseMissionCriticalAssociationId}/actions/failover")
	if err != nil {
		return nil, err
	}

	var response FailoverAutonomousContainerDatabaseMissionCriticalAssociationResponse
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

// FailoverAutonomousPodMissionCriticalAssociation Performs a failover to transition the standby Pod identified by the autonomousPodId parameter into the specified Autonomous Pod Mission Critical Association's primary role after the existing primary Pod fails or becomes unreachable.
// A failover might result in data loss, depending on the protection mode in effect at the time of the primary Pod failure.
func (client DatabaseClient) FailoverAutonomousPodMissionCriticalAssociation(ctx context.Context, request FailoverAutonomousPodMissionCriticalAssociationRequest) (response FailoverAutonomousPodMissionCriticalAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.failoverAutonomousPodMissionCriticalAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			response = FailoverAutonomousPodMissionCriticalAssociationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FailoverAutonomousPodMissionCriticalAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FailoverAutonomousPodMissionCriticalAssociationResponse")
	}
	return
}

// failoverAutonomousPodMissionCriticalAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) failoverAutonomousPodMissionCriticalAssociation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousPods/{autonomousPodId}/autonomousPodMissionCriticalAssociations/{autonomousPodMissionCriticalAssociationId}/actions/failover")
	if err != nil {
		return nil, err
	}

	var response FailoverAutonomousPodMissionCriticalAssociationResponse
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

// FailoverDataGuardAssociation Performs a failover to transition the standby database identified by the `databaseId` parameter into the
// specified Data Guard association's primary role after the existing primary database fails or becomes unreachable.
// A failover might result in data loss depending on the protection mode in effect at the time of the primary
// database failure.
func (client DatabaseClient) FailoverDataGuardAssociation(ctx context.Context, request FailoverDataGuardAssociationRequest) (response FailoverDataGuardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.failoverDataGuardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			response = FailoverDataGuardAssociationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FailoverDataGuardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FailoverDataGuardAssociationResponse")
	}
	return
}

// failoverDataGuardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) failoverDataGuardAssociation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/failover")
	if err != nil {
		return nil, err
	}

	var response FailoverDataGuardAssociationResponse
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

// GenerateAutonomousDataWarehouseWallet Creates and downloads a wallet for the specified Autonomous Data Warehouse.
func (client DatabaseClient) GenerateAutonomousDataWarehouseWallet(ctx context.Context, request GenerateAutonomousDataWarehouseWalletRequest) (response GenerateAutonomousDataWarehouseWalletResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.generateAutonomousDataWarehouseWallet, policy)
	if err != nil {
		if ociResponse != nil {
			response = GenerateAutonomousDataWarehouseWalletResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateAutonomousDataWarehouseWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateAutonomousDataWarehouseWalletResponse")
	}
	return
}

// generateAutonomousDataWarehouseWallet implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) generateAutonomousDataWarehouseWallet(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDataWarehouses/{autonomousDataWarehouseId}/actions/generateWallet")
	if err != nil {
		return nil, err
	}

	var response GenerateAutonomousDataWarehouseWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateAutonomousDatabaseWallet Creates and downloads a wallet for the specified Autonomous Transaction Processing database.
func (client DatabaseClient) GenerateAutonomousDatabaseWallet(ctx context.Context, request GenerateAutonomousDatabaseWalletRequest) (response GenerateAutonomousDatabaseWalletResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.generateAutonomousDatabaseWallet, policy)
	if err != nil {
		if ociResponse != nil {
			response = GenerateAutonomousDatabaseWalletResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateAutonomousDatabaseWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateAutonomousDatabaseWalletResponse")
	}
	return
}

// generateAutonomousDatabaseWallet implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) generateAutonomousDatabaseWallet(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/generateWallet")
	if err != nil {
		return nil, err
	}

	var response GenerateAutonomousDatabaseWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutonomousContainerDatabase Gets information about the specified Autonomous Container Database.
func (client DatabaseClient) GetAutonomousContainerDatabase(ctx context.Context, request GetAutonomousContainerDatabaseRequest) (response GetAutonomousContainerDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAutonomousContainerDatabaseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousContainerDatabaseResponse")
	}
	return
}

// getAutonomousContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousContainerDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}")
	if err != nil {
		return nil, err
	}

	var response GetAutonomousContainerDatabaseResponse
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

// GetAutonomousContainerDatabaseMissionCriticalAssociation Gets an Autonomous Container Database Mission Critical Association for the specified Autonomous Container Database.
func (client DatabaseClient) GetAutonomousContainerDatabaseMissionCriticalAssociation(ctx context.Context, request GetAutonomousContainerDatabaseMissionCriticalAssociationRequest) (response GetAutonomousContainerDatabaseMissionCriticalAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousContainerDatabaseMissionCriticalAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAutonomousContainerDatabaseMissionCriticalAssociationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousContainerDatabaseMissionCriticalAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousContainerDatabaseMissionCriticalAssociationResponse")
	}
	return
}

// getAutonomousContainerDatabaseMissionCriticalAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousContainerDatabaseMissionCriticalAssociation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseMissionCriticalAssociations/{autonomousContainerDatabaseMissionCriticalAssociationId}")
	if err != nil {
		return nil, err
	}

	var response GetAutonomousContainerDatabaseMissionCriticalAssociationResponse
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

// GetAutonomousDataWarehouse Gets the details of the specified Autonomous Data Warehouse.
func (client DatabaseClient) GetAutonomousDataWarehouse(ctx context.Context, request GetAutonomousDataWarehouseRequest) (response GetAutonomousDataWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousDataWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAutonomousDataWarehouseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousDataWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousDataWarehouseResponse")
	}
	return
}

// getAutonomousDataWarehouse implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousDataWarehouse(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDataWarehouses/{autonomousDataWarehouseId}")
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDataWarehouseResponse
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

// GetAutonomousDataWarehouseBackup Gets information about the specified Autonomous Data Warehouse backup.
func (client DatabaseClient) GetAutonomousDataWarehouseBackup(ctx context.Context, request GetAutonomousDataWarehouseBackupRequest) (response GetAutonomousDataWarehouseBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousDataWarehouseBackup, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAutonomousDataWarehouseBackupResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousDataWarehouseBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousDataWarehouseBackupResponse")
	}
	return
}

// getAutonomousDataWarehouseBackup implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousDataWarehouseBackup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDataWarehouseBackups/{autonomousDataWarehouseBackupId}")
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDataWarehouseBackupResponse
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

// GetAutonomousDatabase Gets the details of the specified Autonomous Database.
func (client DatabaseClient) GetAutonomousDatabase(ctx context.Context, request GetAutonomousDatabaseRequest) (response GetAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAutonomousDatabaseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousDatabaseResponse")
	}
	return
}

// getAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/{autonomousDatabaseId}")
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDatabaseResponse
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

// GetAutonomousDatabaseBackup Gets information about the specified Autonomous Database backup.
func (client DatabaseClient) GetAutonomousDatabaseBackup(ctx context.Context, request GetAutonomousDatabaseBackupRequest) (response GetAutonomousDatabaseBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousDatabaseBackup, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAutonomousDatabaseBackupResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousDatabaseBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousDatabaseBackupResponse")
	}
	return
}

// getAutonomousDatabaseBackup implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousDatabaseBackup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabaseBackups/{autonomousDatabaseBackupId}")
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDatabaseBackupResponse
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

// GetAutonomousDatabaseMissionCriticalAssociation Gets an Autonomous Database Mission Critical Assocation associated with the specified Autonomous Database.
func (client DatabaseClient) GetAutonomousDatabaseMissionCriticalAssociation(ctx context.Context, request GetAutonomousDatabaseMissionCriticalAssociationRequest) (response GetAutonomousDatabaseMissionCriticalAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousDatabaseMissionCriticalAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAutonomousDatabaseMissionCriticalAssociationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousDatabaseMissionCriticalAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousDatabaseMissionCriticalAssociationResponse")
	}
	return
}

// getAutonomousDatabaseMissionCriticalAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousDatabaseMissionCriticalAssociation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/{autonomousDatabaseId}/autonomousDatabaseMissionCriticalAssociations/{autonomousDatabaseMissionCriticalAssociationId}")
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDatabaseMissionCriticalAssociationResponse
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

// GetAutonomousDbSystem Gets information about the specified Autonomous DB System.
func (client DatabaseClient) GetAutonomousDbSystem(ctx context.Context, request GetAutonomousDbSystemRequest) (response GetAutonomousDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAutonomousDbSystemResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousDbSystemResponse")
	}
	return
}

// getAutonomousDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousDbSystem(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDbSystems/{autonomousDbSystemId}")
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDbSystemResponse
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

// GetAutonomousPod Gets information about the specified Autonomous Pod.
func (client DatabaseClient) GetAutonomousPod(ctx context.Context, request GetAutonomousPodRequest) (response GetAutonomousPodResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousPod, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAutonomousPodResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousPodResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousPodResponse")
	}
	return
}

// getAutonomousPod implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousPod(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousPods/{autonomousPodId}")
	if err != nil {
		return nil, err
	}

	var response GetAutonomousPodResponse
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

// GetAutonomousPodMissionCriticalAssociation Gets an Autonomous Pod Mission Critical Association for the specified Autonomous Pod.
func (client DatabaseClient) GetAutonomousPodMissionCriticalAssociation(ctx context.Context, request GetAutonomousPodMissionCriticalAssociationRequest) (response GetAutonomousPodMissionCriticalAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousPodMissionCriticalAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAutonomousPodMissionCriticalAssociationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousPodMissionCriticalAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousPodMissionCriticalAssociationResponse")
	}
	return
}

// getAutonomousPodMissionCriticalAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousPodMissionCriticalAssociation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousPods/{autonomousPodId}/autonomousPodMissionCriticalAssociations/{autonomousPodMissionCriticalAssociationId}")
	if err != nil {
		return nil, err
	}

	var response GetAutonomousPodMissionCriticalAssociationResponse
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

// GetBackup Gets information about the specified backup.
func (client DatabaseClient) GetBackup(ctx context.Context, request GetBackupRequest) (response GetBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBackup, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetBackupResponse{RawResponse: ociResponse.HTTPResponse()}
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
func (client DatabaseClient) getBackup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// GetDataGuardAssociation Gets the specified Data Guard association's configuration information.
func (client DatabaseClient) GetDataGuardAssociation(ctx context.Context, request GetDataGuardAssociationRequest) (response GetDataGuardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataGuardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetDataGuardAssociationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataGuardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataGuardAssociationResponse")
	}
	return
}

// getDataGuardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDataGuardAssociation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}")
	if err != nil {
		return nil, err
	}

	var response GetDataGuardAssociationResponse
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

// GetDatabase Gets information about a specific database.
func (client DatabaseClient) GetDatabase(ctx context.Context, request GetDatabaseRequest) (response GetDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetDatabaseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseResponse")
	}
	return
}

// getDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}")
	if err != nil {
		return nil, err
	}

	var response GetDatabaseResponse
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

// GetDbHome Gets information about the specified database home.
func (client DatabaseClient) GetDbHome(ctx context.Context, request GetDbHomeRequest) (response GetDbHomeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbHome, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetDbHomeResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbHomeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbHomeResponse")
	}
	return
}

// getDbHome implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbHome(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}")
	if err != nil {
		return nil, err
	}

	var response GetDbHomeResponse
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

// GetDbHomePatch Gets information about a specified patch package.
func (client DatabaseClient) GetDbHomePatch(ctx context.Context, request GetDbHomePatchRequest) (response GetDbHomePatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbHomePatch, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetDbHomePatchResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbHomePatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbHomePatchResponse")
	}
	return
}

// getDbHomePatch implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbHomePatch(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patches/{patchId}")
	if err != nil {
		return nil, err
	}

	var response GetDbHomePatchResponse
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

// GetDbHomePatchHistoryEntry Gets the patch history details for the specified patchHistoryEntryId
func (client DatabaseClient) GetDbHomePatchHistoryEntry(ctx context.Context, request GetDbHomePatchHistoryEntryRequest) (response GetDbHomePatchHistoryEntryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbHomePatchHistoryEntry, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetDbHomePatchHistoryEntryResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbHomePatchHistoryEntryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbHomePatchHistoryEntryResponse")
	}
	return
}

// getDbHomePatchHistoryEntry implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbHomePatchHistoryEntry(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patchHistoryEntries/{patchHistoryEntryId}")
	if err != nil {
		return nil, err
	}

	var response GetDbHomePatchHistoryEntryResponse
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

// GetDbNode Gets information about the specified database node.
func (client DatabaseClient) GetDbNode(ctx context.Context, request GetDbNodeRequest) (response GetDbNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbNode, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetDbNodeResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbNodeResponse")
	}
	return
}

// getDbNode implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbNode(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbNodes/{dbNodeId}")
	if err != nil {
		return nil, err
	}

	var response GetDbNodeResponse
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

// GetDbSystem Gets information about the specified DB system.
func (client DatabaseClient) GetDbSystem(ctx context.Context, request GetDbSystemRequest) (response GetDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetDbSystemResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbSystemResponse")
	}
	return
}

// getDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbSystem(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}")
	if err != nil {
		return nil, err
	}

	var response GetDbSystemResponse
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

// GetDbSystemPatch Gets information about a specified patch package.
func (client DatabaseClient) GetDbSystemPatch(ctx context.Context, request GetDbSystemPatchRequest) (response GetDbSystemPatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbSystemPatch, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetDbSystemPatchResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbSystemPatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbSystemPatchResponse")
	}
	return
}

// getDbSystemPatch implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbSystemPatch(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patches/{patchId}")
	if err != nil {
		return nil, err
	}

	var response GetDbSystemPatchResponse
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

// GetDbSystemPatchHistoryEntry Gets the patch history details for the specified patchHistoryEntryId.
func (client DatabaseClient) GetDbSystemPatchHistoryEntry(ctx context.Context, request GetDbSystemPatchHistoryEntryRequest) (response GetDbSystemPatchHistoryEntryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbSystemPatchHistoryEntry, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetDbSystemPatchHistoryEntryResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbSystemPatchHistoryEntryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbSystemPatchHistoryEntryResponse")
	}
	return
}

// getDbSystemPatchHistoryEntry implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbSystemPatchHistoryEntry(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patchHistoryEntries/{patchHistoryEntryId}")
	if err != nil {
		return nil, err
	}

	var response GetDbSystemPatchHistoryEntryResponse
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

// GetExadataIormConfig Gets `IORM` Setting for the requested Exadata DB System.
// The default IORM Settings is pre-created in all the Exadata DB System.
func (client DatabaseClient) GetExadataIormConfig(ctx context.Context, request GetExadataIormConfigRequest) (response GetExadataIormConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExadataIormConfig, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetExadataIormConfigResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExadataIormConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExadataIormConfigResponse")
	}
	return
}

// getExadataIormConfig implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getExadataIormConfig(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/ExadataIormConfig")
	if err != nil {
		return nil, err
	}

	var response GetExadataIormConfigResponse
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

// GetExternalBackupJob Gets information about the specified external backup job.
// **Note:** This API is used by an Oracle Cloud Infrastructure Python script that is packaged with the Oracle Cloud Infrastructure CLI. Oracle recommends that you use the script instead using the API directly. See Migrating an On-Premises Database to Oracle Cloud Infrastructure by Creating a Backup in the Cloud (https://docs.us-phoenix-1.oraclecloud.com/Content/Database/Tasks/mig-onprembackup.htm) for more information.
func (client DatabaseClient) GetExternalBackupJob(ctx context.Context, request GetExternalBackupJobRequest) (response GetExternalBackupJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalBackupJob, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetExternalBackupJobResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalBackupJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalBackupJobResponse")
	}
	return
}

// getExternalBackupJob implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getExternalBackupJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalBackupJobs/{backupId}")
	if err != nil {
		return nil, err
	}

	var response GetExternalBackupJobResponse
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

// GetMaintenanceRun Gets information about the specified Maintenance Run.
func (client DatabaseClient) GetMaintenanceRun(ctx context.Context, request GetMaintenanceRunRequest) (response GetMaintenanceRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMaintenanceRun, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetMaintenanceRunResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMaintenanceRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMaintenanceRunResponse")
	}
	return
}

// getMaintenanceRun implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getMaintenanceRun(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceRuns/{maintenanceRunId}")
	if err != nil {
		return nil, err
	}

	var response GetMaintenanceRunResponse
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

// LaunchAutonomousDbSystem Launches a new Autonomous DB System in the specified compartment and Availability Domain.
func (client DatabaseClient) LaunchAutonomousDbSystem(ctx context.Context, request LaunchAutonomousDbSystemRequest) (response LaunchAutonomousDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.launchAutonomousDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			response = LaunchAutonomousDbSystemResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(LaunchAutonomousDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into LaunchAutonomousDbSystemResponse")
	}
	return
}

// launchAutonomousDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) launchAutonomousDbSystem(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDbSystems")
	if err != nil {
		return nil, err
	}

	var response LaunchAutonomousDbSystemResponse
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

// LaunchDbSystem Launches a new DB system in the specified compartment and availability domain. The Oracle
// Database edition that you specify applies to all the databases on that DB system. The selected edition cannot be changed.
// An initial database is created on the DB system based on the request parameters you provide and some default
// options. For more information,
// see Default Options for the Initial Database (https://docs.us-phoenix-1.oraclecloud.com/Content/Database/Tasks/launchingDB.htm#DefaultOptionsfortheInitialDatabase).
func (client DatabaseClient) LaunchDbSystem(ctx context.Context, request LaunchDbSystemRequest) (response LaunchDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.launchDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			response = LaunchDbSystemResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(LaunchDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into LaunchDbSystemResponse")
	}
	return
}

// launchDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) launchDbSystem(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems")
	if err != nil {
		return nil, err
	}

	var response LaunchDbSystemResponse
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

// ListAutonomousContainerDatabaseMissionCriticalAssociations Gets a list of the Autonomous Container Database Mission Critical Associations for the specified Autonomous Container Database.
func (client DatabaseClient) ListAutonomousContainerDatabaseMissionCriticalAssociations(ctx context.Context, request ListAutonomousContainerDatabaseMissionCriticalAssociationsRequest) (response ListAutonomousContainerDatabaseMissionCriticalAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousContainerDatabaseMissionCriticalAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListAutonomousContainerDatabaseMissionCriticalAssociationsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousContainerDatabaseMissionCriticalAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousContainerDatabaseMissionCriticalAssociationsResponse")
	}
	return
}

// listAutonomousContainerDatabaseMissionCriticalAssociations implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousContainerDatabaseMissionCriticalAssociations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseMissionCriticalAssociations")
	if err != nil {
		return nil, err
	}

	var response ListAutonomousContainerDatabaseMissionCriticalAssociationsResponse
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

// ListAutonomousContainerDatabases Gets a list of the Autonomous Container Databases in the specified compartment.
func (client DatabaseClient) ListAutonomousContainerDatabases(ctx context.Context, request ListAutonomousContainerDatabasesRequest) (response ListAutonomousContainerDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousContainerDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListAutonomousContainerDatabasesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousContainerDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousContainerDatabasesResponse")
	}
	return
}

// listAutonomousContainerDatabases implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousContainerDatabases(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases")
	if err != nil {
		return nil, err
	}

	var response ListAutonomousContainerDatabasesResponse
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

// ListAutonomousDataWarehouseBackups Gets a list of Autonomous Data Warehouse backups based on either the `autonomousDataWarehouseId` or `compartmentId` specified as a query parameter.
func (client DatabaseClient) ListAutonomousDataWarehouseBackups(ctx context.Context, request ListAutonomousDataWarehouseBackupsRequest) (response ListAutonomousDataWarehouseBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDataWarehouseBackups, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListAutonomousDataWarehouseBackupsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDataWarehouseBackupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDataWarehouseBackupsResponse")
	}
	return
}

// listAutonomousDataWarehouseBackups implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDataWarehouseBackups(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDataWarehouseBackups")
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDataWarehouseBackupsResponse
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

// ListAutonomousDataWarehouses Gets a list of Autonomous Data Warehouses.
func (client DatabaseClient) ListAutonomousDataWarehouses(ctx context.Context, request ListAutonomousDataWarehousesRequest) (response ListAutonomousDataWarehousesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDataWarehouses, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListAutonomousDataWarehousesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDataWarehousesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDataWarehousesResponse")
	}
	return
}

// listAutonomousDataWarehouses implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDataWarehouses(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDataWarehouses")
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDataWarehousesResponse
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

// ListAutonomousDatabaseBackups Gets a list of Autonomous Database backups based on either the `autonomousDatabaseId` or `compartmentId` specified as a query parameter.
func (client DatabaseClient) ListAutonomousDatabaseBackups(ctx context.Context, request ListAutonomousDatabaseBackupsRequest) (response ListAutonomousDatabaseBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDatabaseBackups, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListAutonomousDatabaseBackupsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDatabaseBackupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDatabaseBackupsResponse")
	}
	return
}

// listAutonomousDatabaseBackups implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDatabaseBackups(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabaseBackups")
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDatabaseBackupsResponse
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

// ListAutonomousDatabaseMissionCriticalAssociations Gets a list of the Autonomous Database Mission Critical Assocations associated with the given Autonomous Database.
func (client DatabaseClient) ListAutonomousDatabaseMissionCriticalAssociations(ctx context.Context, request ListAutonomousDatabaseMissionCriticalAssociationsRequest) (response ListAutonomousDatabaseMissionCriticalAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDatabaseMissionCriticalAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListAutonomousDatabaseMissionCriticalAssociationsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDatabaseMissionCriticalAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDatabaseMissionCriticalAssociationsResponse")
	}
	return
}

// listAutonomousDatabaseMissionCriticalAssociations implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDatabaseMissionCriticalAssociations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/{autonomousDatabaseId}/autonomousDatabaseMissionCriticalAssociations")
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDatabaseMissionCriticalAssociationsResponse
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

// ListAutonomousDatabases Gets a list of Autonomous Databases.
func (client DatabaseClient) ListAutonomousDatabases(ctx context.Context, request ListAutonomousDatabasesRequest) (response ListAutonomousDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListAutonomousDatabasesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDatabasesResponse")
	}
	return
}

// listAutonomousDatabases implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDatabases(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases")
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDatabasesResponse
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

// ListAutonomousDbSystemShapes Gets a list of the shapes that can be used to launch a new Autonomous DB System. The shape determines resources to allocate to the DB system -  CPU cores, memory and storage for shapes.
func (client DatabaseClient) ListAutonomousDbSystemShapes(ctx context.Context, request ListAutonomousDbSystemShapesRequest) (response ListAutonomousDbSystemShapesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDbSystemShapes, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListAutonomousDbSystemShapesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDbSystemShapesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDbSystemShapesResponse")
	}
	return
}

// listAutonomousDbSystemShapes implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDbSystemShapes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDbSystemShapes")
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDbSystemShapesResponse
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

// ListAutonomousDbSystems Gets a list of the Autonomous DB Systems in the specified compartment.
func (client DatabaseClient) ListAutonomousDbSystems(ctx context.Context, request ListAutonomousDbSystemsRequest) (response ListAutonomousDbSystemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDbSystems, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListAutonomousDbSystemsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDbSystemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDbSystemsResponse")
	}
	return
}

// listAutonomousDbSystems implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDbSystems(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDbSystems")
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDbSystemsResponse
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

// ListAutonomousPodMissionCriticalAssociations Gets a list of the Autonomous Pod Mission Critical Associations for the specified Autonomous Pod.
func (client DatabaseClient) ListAutonomousPodMissionCriticalAssociations(ctx context.Context, request ListAutonomousPodMissionCriticalAssociationsRequest) (response ListAutonomousPodMissionCriticalAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousPodMissionCriticalAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListAutonomousPodMissionCriticalAssociationsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousPodMissionCriticalAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousPodMissionCriticalAssociationsResponse")
	}
	return
}

// listAutonomousPodMissionCriticalAssociations implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousPodMissionCriticalAssociations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousPods/{autonomousPodId}/autonomousPodMissionCriticalAssociations")
	if err != nil {
		return nil, err
	}

	var response ListAutonomousPodMissionCriticalAssociationsResponse
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

// ListAutonomousPods Gets a list of the Autonomous Pods in the specified compartment.
func (client DatabaseClient) ListAutonomousPods(ctx context.Context, request ListAutonomousPodsRequest) (response ListAutonomousPodsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousPods, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListAutonomousPodsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousPodsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousPodsResponse")
	}
	return
}

// listAutonomousPods implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousPods(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousPods")
	if err != nil {
		return nil, err
	}

	var response ListAutonomousPodsResponse
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

// ListBackups Gets a list of backups based on the databaseId or compartmentId specified. Either one of the query parameters must be provided.
func (client DatabaseClient) ListBackups(ctx context.Context, request ListBackupsRequest) (response ListBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBackups, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListBackupsResponse{RawResponse: ociResponse.HTTPResponse()}
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
func (client DatabaseClient) listBackups(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// ListDataGuardAssociations Lists all Data Guard associations for the specified database.
func (client DatabaseClient) ListDataGuardAssociations(ctx context.Context, request ListDataGuardAssociationsRequest) (response ListDataGuardAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataGuardAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListDataGuardAssociationsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataGuardAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataGuardAssociationsResponse")
	}
	return
}

// listDataGuardAssociations implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDataGuardAssociations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}/dataGuardAssociations")
	if err != nil {
		return nil, err
	}

	var response ListDataGuardAssociationsResponse
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

// ListDatabases Gets a list of the databases in the specified database home.
func (client DatabaseClient) ListDatabases(ctx context.Context, request ListDatabasesRequest) (response ListDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListDatabasesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabasesResponse")
	}
	return
}

// listDatabases implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDatabases(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases")
	if err != nil {
		return nil, err
	}

	var response ListDatabasesResponse
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

// ListDbHomePatchHistoryEntries Gets history of the actions taken for patches for the specified database home.
func (client DatabaseClient) ListDbHomePatchHistoryEntries(ctx context.Context, request ListDbHomePatchHistoryEntriesRequest) (response ListDbHomePatchHistoryEntriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbHomePatchHistoryEntries, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListDbHomePatchHistoryEntriesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbHomePatchHistoryEntriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbHomePatchHistoryEntriesResponse")
	}
	return
}

// listDbHomePatchHistoryEntries implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbHomePatchHistoryEntries(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patchHistoryEntries")
	if err != nil {
		return nil, err
	}

	var response ListDbHomePatchHistoryEntriesResponse
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

// ListDbHomePatches Lists patches applicable to the requested database home.
func (client DatabaseClient) ListDbHomePatches(ctx context.Context, request ListDbHomePatchesRequest) (response ListDbHomePatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbHomePatches, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListDbHomePatchesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbHomePatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbHomePatchesResponse")
	}
	return
}

// listDbHomePatches implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbHomePatches(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patches")
	if err != nil {
		return nil, err
	}

	var response ListDbHomePatchesResponse
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

// ListDbHomes Gets a list of database homes in the specified DB system and compartment. A database home is a directory where Oracle Database software is installed.
func (client DatabaseClient) ListDbHomes(ctx context.Context, request ListDbHomesRequest) (response ListDbHomesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbHomes, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListDbHomesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbHomesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbHomesResponse")
	}
	return
}

// listDbHomes implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbHomes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes")
	if err != nil {
		return nil, err
	}

	var response ListDbHomesResponse
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

// ListDbNodes Gets a list of database nodes in the specified DB system and compartment. A database node is a server running database software.
func (client DatabaseClient) ListDbNodes(ctx context.Context, request ListDbNodesRequest) (response ListDbNodesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbNodes, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListDbNodesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbNodesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbNodesResponse")
	}
	return
}

// listDbNodes implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbNodes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbNodes")
	if err != nil {
		return nil, err
	}

	var response ListDbNodesResponse
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

// ListDbSystemPatchHistoryEntries Gets the history of the patch actions performed on the specified DB system.
func (client DatabaseClient) ListDbSystemPatchHistoryEntries(ctx context.Context, request ListDbSystemPatchHistoryEntriesRequest) (response ListDbSystemPatchHistoryEntriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbSystemPatchHistoryEntries, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListDbSystemPatchHistoryEntriesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbSystemPatchHistoryEntriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbSystemPatchHistoryEntriesResponse")
	}
	return
}

// listDbSystemPatchHistoryEntries implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbSystemPatchHistoryEntries(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patchHistoryEntries")
	if err != nil {
		return nil, err
	}

	var response ListDbSystemPatchHistoryEntriesResponse
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

// ListDbSystemPatches Lists the patches applicable to the requested DB system.
func (client DatabaseClient) ListDbSystemPatches(ctx context.Context, request ListDbSystemPatchesRequest) (response ListDbSystemPatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbSystemPatches, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListDbSystemPatchesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbSystemPatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbSystemPatchesResponse")
	}
	return
}

// listDbSystemPatches implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbSystemPatches(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patches")
	if err != nil {
		return nil, err
	}

	var response ListDbSystemPatchesResponse
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

// ListDbSystemShapes Gets a list of the shapes that can be used to launch a new DB system. The shape determines resources to allocate to the DB system - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes.
func (client DatabaseClient) ListDbSystemShapes(ctx context.Context, request ListDbSystemShapesRequest) (response ListDbSystemShapesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbSystemShapes, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListDbSystemShapesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbSystemShapesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbSystemShapesResponse")
	}
	return
}

// listDbSystemShapes implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbSystemShapes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystemShapes")
	if err != nil {
		return nil, err
	}

	var response ListDbSystemShapesResponse
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

// ListDbSystems Gets a list of the DB systems in the specified compartment. You can specify a backupId to list only the DB systems that support creating a database using this backup in this compartment.
//
func (client DatabaseClient) ListDbSystems(ctx context.Context, request ListDbSystemsRequest) (response ListDbSystemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbSystems, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListDbSystemsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbSystemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbSystemsResponse")
	}
	return
}

// listDbSystems implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbSystems(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems")
	if err != nil {
		return nil, err
	}

	var response ListDbSystemsResponse
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

// ListDbVersions Gets a list of supported Oracle Database versions.
func (client DatabaseClient) ListDbVersions(ctx context.Context, request ListDbVersionsRequest) (response ListDbVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbVersions, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListDbVersionsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbVersionsResponse")
	}
	return
}

// listDbVersions implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbVersions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbVersions")
	if err != nil {
		return nil, err
	}

	var response ListDbVersionsResponse
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

// ListMaintenanceRuns Gets a list of the Maintenance Runs in the specified compartment.
func (client DatabaseClient) ListMaintenanceRuns(ctx context.Context, request ListMaintenanceRunsRequest) (response ListMaintenanceRunsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaintenanceRuns, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListMaintenanceRunsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaintenanceRunsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaintenanceRunsResponse")
	}
	return
}

// listMaintenanceRuns implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listMaintenanceRuns(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceRuns")
	if err != nil {
		return nil, err
	}

	var response ListMaintenanceRunsResponse
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

// ReinstateDataGuardAssociation Reinstates the database identified by the `databaseId` parameter into the standby role in a Data Guard association.
func (client DatabaseClient) ReinstateDataGuardAssociation(ctx context.Context, request ReinstateDataGuardAssociationRequest) (response ReinstateDataGuardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.reinstateDataGuardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			response = ReinstateDataGuardAssociationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReinstateDataGuardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReinstateDataGuardAssociationResponse")
	}
	return
}

// reinstateDataGuardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) reinstateDataGuardAssociation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/reinstate")
	if err != nil {
		return nil, err
	}

	var response ReinstateDataGuardAssociationResponse
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

// RestoreAutonomousDataWarehouse Restores an Autonomous Data Warehouse based on the provided request parameters.
func (client DatabaseClient) RestoreAutonomousDataWarehouse(ctx context.Context, request RestoreAutonomousDataWarehouseRequest) (response RestoreAutonomousDataWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.restoreAutonomousDataWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			response = RestoreAutonomousDataWarehouseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestoreAutonomousDataWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestoreAutonomousDataWarehouseResponse")
	}
	return
}

// restoreAutonomousDataWarehouse implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) restoreAutonomousDataWarehouse(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDataWarehouses/{autonomousDataWarehouseId}/actions/restore")
	if err != nil {
		return nil, err
	}

	var response RestoreAutonomousDataWarehouseResponse
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

// RestoreAutonomousDatabase Restores an Autonomous Database based on the provided request parameters.
func (client DatabaseClient) RestoreAutonomousDatabase(ctx context.Context, request RestoreAutonomousDatabaseRequest) (response RestoreAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.restoreAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			response = RestoreAutonomousDatabaseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestoreAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestoreAutonomousDatabaseResponse")
	}
	return
}

// restoreAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) restoreAutonomousDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/restore")
	if err != nil {
		return nil, err
	}

	var response RestoreAutonomousDatabaseResponse
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

// RestoreDatabase Restore a Database based on the request parameters you provide.
func (client DatabaseClient) RestoreDatabase(ctx context.Context, request RestoreDatabaseRequest) (response RestoreDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.restoreDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			response = RestoreDatabaseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestoreDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestoreDatabaseResponse")
	}
	return
}

// restoreDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) restoreDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/actions/restore")
	if err != nil {
		return nil, err
	}

	var response RestoreDatabaseResponse
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

// StartAutonomousDataWarehouse Starts the specified autonomous Data Warehouse.
func (client DatabaseClient) StartAutonomousDataWarehouse(ctx context.Context, request StartAutonomousDataWarehouseRequest) (response StartAutonomousDataWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.startAutonomousDataWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			response = StartAutonomousDataWarehouseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartAutonomousDataWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartAutonomousDataWarehouseResponse")
	}
	return
}

// startAutonomousDataWarehouse implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) startAutonomousDataWarehouse(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDataWarehouses/{autonomousDataWarehouseId}/actions/start")
	if err != nil {
		return nil, err
	}

	var response StartAutonomousDataWarehouseResponse
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

// StartAutonomousDatabase Starts the specified Autonomous Database.
func (client DatabaseClient) StartAutonomousDatabase(ctx context.Context, request StartAutonomousDatabaseRequest) (response StartAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.startAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			response = StartAutonomousDatabaseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartAutonomousDatabaseResponse")
	}
	return
}

// startAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) startAutonomousDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/start")
	if err != nil {
		return nil, err
	}

	var response StartAutonomousDatabaseResponse
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

// StopAutonomousDataWarehouse Stops the specified Autonomous Data Warehouse.
func (client DatabaseClient) StopAutonomousDataWarehouse(ctx context.Context, request StopAutonomousDataWarehouseRequest) (response StopAutonomousDataWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.stopAutonomousDataWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			response = StopAutonomousDataWarehouseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopAutonomousDataWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopAutonomousDataWarehouseResponse")
	}
	return
}

// stopAutonomousDataWarehouse implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) stopAutonomousDataWarehouse(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDataWarehouses/{autonomousDataWarehouseId}/actions/stop")
	if err != nil {
		return nil, err
	}

	var response StopAutonomousDataWarehouseResponse
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

// StopAutonomousDatabase Stops the specified Autonomous Database.
func (client DatabaseClient) StopAutonomousDatabase(ctx context.Context, request StopAutonomousDatabaseRequest) (response StopAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.stopAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			response = StopAutonomousDatabaseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopAutonomousDatabaseResponse")
	}
	return
}

// stopAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) stopAutonomousDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/stop")
	if err != nil {
		return nil, err
	}

	var response StopAutonomousDatabaseResponse
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

// SwitchoverAutonomousContainerDatabaseMissionCriticalAssociation Performs a switchover to transition the primary Container Database of an Autonomous Container Database Mission Critical Association into a standby role. The standby Container Database associated with autonomousContainerDatabaseMissionCriticalAssociationId assumes the primary Container Database role.
// A switchover guarantees no data loss.
func (client DatabaseClient) SwitchoverAutonomousContainerDatabaseMissionCriticalAssociation(ctx context.Context, request SwitchoverAutonomousContainerDatabaseMissionCriticalAssociationRequest) (response SwitchoverAutonomousContainerDatabaseMissionCriticalAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.switchoverAutonomousContainerDatabaseMissionCriticalAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			response = SwitchoverAutonomousContainerDatabaseMissionCriticalAssociationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SwitchoverAutonomousContainerDatabaseMissionCriticalAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SwitchoverAutonomousContainerDatabaseMissionCriticalAssociationResponse")
	}
	return
}

// switchoverAutonomousContainerDatabaseMissionCriticalAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) switchoverAutonomousContainerDatabaseMissionCriticalAssociation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseMissionCriticalAssociations/{autonomousContainerDatabaseMissionCriticalAssociationId}/actions/switchover")
	if err != nil {
		return nil, err
	}

	var response SwitchoverAutonomousContainerDatabaseMissionCriticalAssociationResponse
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

// SwitchoverAutonomousPodMissionCriticalAssociation Performs a switchover to transition the primary Pod of an Autonomous Pod Mission Critical Association into a standby role. The standby Pod associated with autonomousPodMissionCriticalAssociationId assumes the primary Pod role.
// A switchover guarantees no data loss.
func (client DatabaseClient) SwitchoverAutonomousPodMissionCriticalAssociation(ctx context.Context, request SwitchoverAutonomousPodMissionCriticalAssociationRequest) (response SwitchoverAutonomousPodMissionCriticalAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.switchoverAutonomousPodMissionCriticalAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			response = SwitchoverAutonomousPodMissionCriticalAssociationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SwitchoverAutonomousPodMissionCriticalAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SwitchoverAutonomousPodMissionCriticalAssociationResponse")
	}
	return
}

// switchoverAutonomousPodMissionCriticalAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) switchoverAutonomousPodMissionCriticalAssociation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousPods/{autonomousPodId}/autonomousPodMissionCriticalAssociations/{autonomousPodMissionCriticalAssociationId}/actions/switchover")
	if err != nil {
		return nil, err
	}

	var response SwitchoverAutonomousPodMissionCriticalAssociationResponse
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

// SwitchoverDataGuardAssociation Performs a switchover to transition the primary database of a Data Guard association into a standby role. The
// standby database associated with the `dataGuardAssociationId` assumes the primary database role.
// A switchover guarantees no data loss.
func (client DatabaseClient) SwitchoverDataGuardAssociation(ctx context.Context, request SwitchoverDataGuardAssociationRequest) (response SwitchoverDataGuardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.switchoverDataGuardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			response = SwitchoverDataGuardAssociationResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SwitchoverDataGuardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SwitchoverDataGuardAssociationResponse")
	}
	return
}

// switchoverDataGuardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) switchoverDataGuardAssociation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/switchover")
	if err != nil {
		return nil, err
	}

	var response SwitchoverDataGuardAssociationResponse
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

// TerminateAutonomousContainerDatabase Terminates an Autonomous Container Database, which permanently deletes the container database and any databases within the container database. The database data is local to the Autonomous Exadata Infrastructure and will be lost when the container database is terminated. Oracle recommends that you back up any data in the Autonomous Container Database prior to terminating it.
func (client DatabaseClient) TerminateAutonomousContainerDatabase(ctx context.Context, request TerminateAutonomousContainerDatabaseRequest) (response TerminateAutonomousContainerDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.terminateAutonomousContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			response = TerminateAutonomousContainerDatabaseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TerminateAutonomousContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TerminateAutonomousContainerDatabaseResponse")
	}
	return
}

// terminateAutonomousContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) terminateAutonomousContainerDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}")
	if err != nil {
		return nil, err
	}

	var response TerminateAutonomousContainerDatabaseResponse
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

// TerminateAutonomousDbSystem Terminates an Autonomous Exadata Infrastructure, which permanently deletes the Exadata Infrastructure and any container databases and databases contained in the Exadata Infrastructure. The database data is local to the Autonomous Exadata Infrastructure and will be lost when the system is terminated. Oracle recommends that you back up any data in the Autonomous Exadata Infrastructure prior to terminating it.
func (client DatabaseClient) TerminateAutonomousDbSystem(ctx context.Context, request TerminateAutonomousDbSystemRequest) (response TerminateAutonomousDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.terminateAutonomousDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			response = TerminateAutonomousDbSystemResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TerminateAutonomousDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TerminateAutonomousDbSystemResponse")
	}
	return
}

// terminateAutonomousDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) terminateAutonomousDbSystem(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousDbSystems/{autonomousDbSystemId}")
	if err != nil {
		return nil, err
	}

	var response TerminateAutonomousDbSystemResponse
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

// TerminateAutonomousPod Terminates an Autonomous Pod, which permanently deletes the container database and any databases within the container database. The database data is local to the Autonomous Exadata Infrastructure and will be lost when the container database is terminated. Oracle recommends that you back up any data in the Autonomous Pod prior to terminating it.
func (client DatabaseClient) TerminateAutonomousPod(ctx context.Context, request TerminateAutonomousPodRequest) (response TerminateAutonomousPodResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.terminateAutonomousPod, policy)
	if err != nil {
		if ociResponse != nil {
			response = TerminateAutonomousPodResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TerminateAutonomousPodResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TerminateAutonomousPodResponse")
	}
	return
}

// terminateAutonomousPod implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) terminateAutonomousPod(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousPods/{autonomousPodId}")
	if err != nil {
		return nil, err
	}

	var response TerminateAutonomousPodResponse
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

// TerminateDbSystem Terminates a DB system and permanently deletes it and any databases running on it, and any storage volumes attached to it. The database data is local to the DB system and will be lost when the system is terminated. Oracle recommends that you back up any data in the DB system prior to terminating it.
func (client DatabaseClient) TerminateDbSystem(ctx context.Context, request TerminateDbSystemRequest) (response TerminateDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.terminateDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			response = TerminateDbSystemResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TerminateDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TerminateDbSystemResponse")
	}
	return
}

// terminateDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) terminateDbSystem(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbSystems/{dbSystemId}")
	if err != nil {
		return nil, err
	}

	var response TerminateDbSystemResponse
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

// UpdateAutonomousContainerDatabase Updates the properties of an Autonomous Container Database, such as the CPU core count and storage size.
func (client DatabaseClient) UpdateAutonomousContainerDatabase(ctx context.Context, request UpdateAutonomousContainerDatabaseRequest) (response UpdateAutonomousContainerDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAutonomousContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateAutonomousContainerDatabaseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutonomousContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutonomousContainerDatabaseResponse")
	}
	return
}

// updateAutonomousContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateAutonomousContainerDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}")
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousContainerDatabaseResponse
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

// UpdateAutonomousDataWarehouse Updates the specified Autonomous Data Warehouse with a new CPU core count and size.
func (client DatabaseClient) UpdateAutonomousDataWarehouse(ctx context.Context, request UpdateAutonomousDataWarehouseRequest) (response UpdateAutonomousDataWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAutonomousDataWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateAutonomousDataWarehouseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutonomousDataWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutonomousDataWarehouseResponse")
	}
	return
}

// updateAutonomousDataWarehouse implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateAutonomousDataWarehouse(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousDataWarehouses/{autonomousDataWarehouseId}")
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousDataWarehouseResponse
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

// UpdateAutonomousDatabase Updates the specified Autonomous Database with a new CPU core count and size.
func (client DatabaseClient) UpdateAutonomousDatabase(ctx context.Context, request UpdateAutonomousDatabaseRequest) (response UpdateAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateAutonomousDatabaseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutonomousDatabaseResponse")
	}
	return
}

// updateAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateAutonomousDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousDatabases/{autonomousDatabaseId}")
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousDatabaseResponse
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

// UpdateAutonomousDbSystem Updates the properties of a Autonomous DB System, such as the CPU core count.
func (client DatabaseClient) UpdateAutonomousDbSystem(ctx context.Context, request UpdateAutonomousDbSystemRequest) (response UpdateAutonomousDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAutonomousDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateAutonomousDbSystemResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutonomousDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutonomousDbSystemResponse")
	}
	return
}

// updateAutonomousDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateAutonomousDbSystem(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousDbSystems/{autonomousDbSystemId}")
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousDbSystemResponse
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

// UpdateAutonomousPod Updates the properties of an Autonomous Pod, such as the CPU core count and storage size.
func (client DatabaseClient) UpdateAutonomousPod(ctx context.Context, request UpdateAutonomousPodRequest) (response UpdateAutonomousPodResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAutonomousPod, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateAutonomousPodResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutonomousPodResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutonomousPodResponse")
	}
	return
}

// updateAutonomousPod implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateAutonomousPod(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousPods/{autonomousPodId}")
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousPodResponse
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

// UpdateDatabase Update a Database based on the request parameters you provide.
func (client DatabaseClient) UpdateDatabase(ctx context.Context, request UpdateDatabaseRequest) (response UpdateDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateDatabaseResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseResponse")
	}
	return
}

// updateDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databases/{databaseId}")
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseResponse
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

// UpdateDbHome Patches the specified dbHome.
func (client DatabaseClient) UpdateDbHome(ctx context.Context, request UpdateDbHomeRequest) (response UpdateDbHomeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDbHome, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateDbHomeResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDbHomeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDbHomeResponse")
	}
	return
}

// updateDbHome implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateDbHome(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbHomes/{dbHomeId}")
	if err != nil {
		return nil, err
	}

	var response UpdateDbHomeResponse
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

// UpdateDbSystem Updates the properties of a DB system, such as the CPU core count.
func (client DatabaseClient) UpdateDbSystem(ctx context.Context, request UpdateDbSystemRequest) (response UpdateDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateDbSystemResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDbSystemResponse")
	}
	return
}

// updateDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateDbSystem(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbSystems/{dbSystemId}")
	if err != nil {
		return nil, err
	}

	var response UpdateDbSystemResponse
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

// UpdateExadataIormConfig Update `IORM` Settings for the requested Exadata DB System.
func (client DatabaseClient) UpdateExadataIormConfig(ctx context.Context, request UpdateExadataIormConfigRequest) (response UpdateExadataIormConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExadataIormConfig, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateExadataIormConfigResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExadataIormConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExadataIormConfigResponse")
	}
	return
}

// updateExadataIormConfig implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateExadataIormConfig(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbSystems/{dbSystemId}/ExadataIormConfig")
	if err != nil {
		return nil, err
	}

	var response UpdateExadataIormConfigResponse
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

// UpdateMaintenanceRun Updates the properties of a Maintenance Run, such as the state of a Maintenance Run.
func (client DatabaseClient) UpdateMaintenanceRun(ctx context.Context, request UpdateMaintenanceRunRequest) (response UpdateMaintenanceRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMaintenanceRun, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateMaintenanceRunResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMaintenanceRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMaintenanceRunResponse")
	}
	return
}

// updateMaintenanceRun implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateMaintenanceRun(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/maintenanceRuns/{maintenanceRunId}")
	if err != nil {
		return nil, err
	}

	var response UpdateMaintenanceRunResponse
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
