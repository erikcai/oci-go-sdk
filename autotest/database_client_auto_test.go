package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/database"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCompleteExternalBackupJob(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "CompleteExternalBackupJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CompleteExternalBackupJob is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "CompleteExternalBackupJob")
	assert.NoError(t, err)

	type CompleteExternalBackupJobRequestInfo struct {
		ContainerId string
		Request     database.CompleteExternalBackupJobRequest
	}

	var requests []CompleteExternalBackupJobRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CompleteExternalBackupJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousDataWarehouse(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousDataWarehouse is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "CreateAutonomousDataWarehouse")
	assert.NoError(t, err)

	type CreateAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousDataWarehouseRequest
	}

	var requests []CreateAutonomousDataWarehouseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousDataWarehouseBackup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousDataWarehouseBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousDataWarehouseBackup is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "CreateAutonomousDataWarehouseBackup")
	assert.NoError(t, err)

	type CreateAutonomousDataWarehouseBackupRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousDataWarehouseBackupRequest
	}

	var requests []CreateAutonomousDataWarehouseBackupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateAutonomousDataWarehouseBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousDatabase(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousDatabase is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "CreateAutonomousDatabase")
	assert.NoError(t, err)

	type CreateAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousDatabaseRequest
	}

	var requests []CreateAutonomousDatabaseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousDatabaseBackup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousDatabaseBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousDatabaseBackup is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "CreateAutonomousDatabaseBackup")
	assert.NoError(t, err)

	type CreateAutonomousDatabaseBackupRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousDatabaseBackupRequest
	}

	var requests []CreateAutonomousDatabaseBackupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateAutonomousDatabaseBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousPod(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousPod")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousPod is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "CreateAutonomousPod")
	assert.NoError(t, err)

	type CreateAutonomousPodRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousPodRequest
	}

	var requests []CreateAutonomousPodRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateAutonomousPod(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateBackup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "CreateBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBackup is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "CreateBackup")
	assert.NoError(t, err)

	type CreateBackupRequestInfo struct {
		ContainerId string
		Request     database.CreateBackupRequest
	}

	var requests []CreateBackupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateDataGuardAssociation(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "CreateDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDataGuardAssociation is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "CreateDataGuardAssociation")
	assert.NoError(t, err)

	type CreateDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.CreateDataGuardAssociationRequest
	}

	var requests []CreateDataGuardAssociationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateDbHome(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "CreateDbHome")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDbHome is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "CreateDbHome")
	assert.NoError(t, err)

	type CreateDbHomeRequestInfo struct {
		ContainerId string
		Request     database.CreateDbHomeRequest
	}

	var requests []CreateDbHomeRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateDbHome(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateExternalBackupJob(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "CreateExternalBackupJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateExternalBackupJob is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "CreateExternalBackupJob")
	assert.NoError(t, err)

	type CreateExternalBackupJobRequestInfo struct {
		ContainerId string
		Request     database.CreateExternalBackupJobRequest
	}

	var requests []CreateExternalBackupJobRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateExternalBackupJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDbNodeAction(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "DbNodeAction")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DbNodeAction is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "DbNodeAction")
	assert.NoError(t, err)

	type DbNodeActionRequestInfo struct {
		ContainerId string
		Request     database.DbNodeActionRequest
	}

	var requests []DbNodeActionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DbNodeAction(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteAutonomousDataWarehouse(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "DeleteAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAutonomousDataWarehouse is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "DeleteAutonomousDataWarehouse")
	assert.NoError(t, err)

	type DeleteAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.DeleteAutonomousDataWarehouseRequest
	}

	var requests []DeleteAutonomousDataWarehouseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteAutonomousDatabase(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "DeleteAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAutonomousDatabase is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "DeleteAutonomousDatabase")
	assert.NoError(t, err)

	type DeleteAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.DeleteAutonomousDatabaseRequest
	}

	var requests []DeleteAutonomousDatabaseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteAutonomousDatabaseBackup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "DeleteAutonomousDatabaseBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAutonomousDatabaseBackup is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "DeleteAutonomousDatabaseBackup")
	assert.NoError(t, err)

	type DeleteAutonomousDatabaseBackupRequestInfo struct {
		ContainerId string
		Request     database.DeleteAutonomousDatabaseBackupRequest
	}

	var requests []DeleteAutonomousDatabaseBackupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteAutonomousDatabaseBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteBackup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "DeleteBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBackup is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "DeleteBackup")
	assert.NoError(t, err)

	type DeleteBackupRequestInfo struct {
		ContainerId string
		Request     database.DeleteBackupRequest
	}

	var requests []DeleteBackupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteDbHome(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "DeleteDbHome")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDbHome is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "DeleteDbHome")
	assert.NoError(t, err)

	type DeleteDbHomeRequestInfo struct {
		ContainerId string
		Request     database.DeleteDbHomeRequest
	}

	var requests []DeleteDbHomeRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteDbHome(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientFailoverAutonomousPodMissionCriticalAssociation(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "FailoverAutonomousPodMissionCriticalAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("FailoverAutonomousPodMissionCriticalAssociation is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "FailoverAutonomousPodMissionCriticalAssociation")
	assert.NoError(t, err)

	type FailoverAutonomousPodMissionCriticalAssociationRequestInfo struct {
		ContainerId string
		Request     database.FailoverAutonomousPodMissionCriticalAssociationRequest
	}

	var requests []FailoverAutonomousPodMissionCriticalAssociationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.FailoverAutonomousPodMissionCriticalAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientFailoverDataGuardAssociation(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "FailoverDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("FailoverDataGuardAssociation is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "FailoverDataGuardAssociation")
	assert.NoError(t, err)

	type FailoverDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.FailoverDataGuardAssociationRequest
	}

	var requests []FailoverDataGuardAssociationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.FailoverDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGenerateAutonomousDataWarehouseWallet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GenerateAutonomousDataWarehouseWallet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GenerateAutonomousDataWarehouseWallet is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GenerateAutonomousDataWarehouseWallet")
	assert.NoError(t, err)

	type GenerateAutonomousDataWarehouseWalletRequestInfo struct {
		ContainerId string
		Request     database.GenerateAutonomousDataWarehouseWalletRequest
	}

	var requests []GenerateAutonomousDataWarehouseWalletRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GenerateAutonomousDataWarehouseWallet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGenerateAutonomousDatabaseWallet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GenerateAutonomousDatabaseWallet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GenerateAutonomousDatabaseWallet is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GenerateAutonomousDatabaseWallet")
	assert.NoError(t, err)

	type GenerateAutonomousDatabaseWalletRequestInfo struct {
		ContainerId string
		Request     database.GenerateAutonomousDatabaseWalletRequest
	}

	var requests []GenerateAutonomousDatabaseWalletRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GenerateAutonomousDatabaseWallet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDataWarehouse(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDataWarehouse is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetAutonomousDataWarehouse")
	assert.NoError(t, err)

	type GetAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDataWarehouseRequest
	}

	var requests []GetAutonomousDataWarehouseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDataWarehouseBackup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDataWarehouseBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDataWarehouseBackup is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetAutonomousDataWarehouseBackup")
	assert.NoError(t, err)

	type GetAutonomousDataWarehouseBackupRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDataWarehouseBackupRequest
	}

	var requests []GetAutonomousDataWarehouseBackupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAutonomousDataWarehouseBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDatabase(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDatabase is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetAutonomousDatabase")
	assert.NoError(t, err)

	type GetAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDatabaseRequest
	}

	var requests []GetAutonomousDatabaseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDatabaseBackup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDatabaseBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDatabaseBackup is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetAutonomousDatabaseBackup")
	assert.NoError(t, err)

	type GetAutonomousDatabaseBackupRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDatabaseBackupRequest
	}

	var requests []GetAutonomousDatabaseBackupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAutonomousDatabaseBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDatabaseMissionCriticalAssociation(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDatabaseMissionCriticalAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDatabaseMissionCriticalAssociation is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetAutonomousDatabaseMissionCriticalAssociation")
	assert.NoError(t, err)

	type GetAutonomousDatabaseMissionCriticalAssociationRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDatabaseMissionCriticalAssociationRequest
	}

	var requests []GetAutonomousDatabaseMissionCriticalAssociationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAutonomousDatabaseMissionCriticalAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDbSystem(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDbSystem is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetAutonomousDbSystem")
	assert.NoError(t, err)

	type GetAutonomousDbSystemRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDbSystemRequest
	}

	var requests []GetAutonomousDbSystemRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAutonomousDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousPod(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetAutonomousPod")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousPod is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetAutonomousPod")
	assert.NoError(t, err)

	type GetAutonomousPodRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousPodRequest
	}

	var requests []GetAutonomousPodRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAutonomousPod(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousPodMissionCriticalAssociation(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetAutonomousPodMissionCriticalAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousPodMissionCriticalAssociation is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetAutonomousPodMissionCriticalAssociation")
	assert.NoError(t, err)

	type GetAutonomousPodMissionCriticalAssociationRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousPodMissionCriticalAssociationRequest
	}

	var requests []GetAutonomousPodMissionCriticalAssociationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAutonomousPodMissionCriticalAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetBackup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBackup is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetBackup")
	assert.NoError(t, err)

	type GetBackupRequestInfo struct {
		ContainerId string
		Request     database.GetBackupRequest
	}

	var requests []GetBackupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDataGuardAssociation(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDataGuardAssociation is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetDataGuardAssociation")
	assert.NoError(t, err)

	type GetDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.GetDataGuardAssociationRequest
	}

	var requests []GetDataGuardAssociationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDatabase(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDatabase is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetDatabase")
	assert.NoError(t, err)

	type GetDatabaseRequestInfo struct {
		ContainerId string
		Request     database.GetDatabaseRequest
	}

	var requests []GetDatabaseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbHome(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetDbHome")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbHome is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetDbHome")
	assert.NoError(t, err)

	type GetDbHomeRequestInfo struct {
		ContainerId string
		Request     database.GetDbHomeRequest
	}

	var requests []GetDbHomeRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDbHome(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbHomePatch(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetDbHomePatch")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbHomePatch is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetDbHomePatch")
	assert.NoError(t, err)

	type GetDbHomePatchRequestInfo struct {
		ContainerId string
		Request     database.GetDbHomePatchRequest
	}

	var requests []GetDbHomePatchRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDbHomePatch(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbHomePatchHistoryEntry(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetDbHomePatchHistoryEntry")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbHomePatchHistoryEntry is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetDbHomePatchHistoryEntry")
	assert.NoError(t, err)

	type GetDbHomePatchHistoryEntryRequestInfo struct {
		ContainerId string
		Request     database.GetDbHomePatchHistoryEntryRequest
	}

	var requests []GetDbHomePatchHistoryEntryRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDbHomePatchHistoryEntry(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbNode(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetDbNode")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbNode is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetDbNode")
	assert.NoError(t, err)

	type GetDbNodeRequestInfo struct {
		ContainerId string
		Request     database.GetDbNodeRequest
	}

	var requests []GetDbNodeRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDbNode(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbSystem(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbSystem is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetDbSystem")
	assert.NoError(t, err)

	type GetDbSystemRequestInfo struct {
		ContainerId string
		Request     database.GetDbSystemRequest
	}

	var requests []GetDbSystemRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbSystemPatch(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetDbSystemPatch")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbSystemPatch is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetDbSystemPatch")
	assert.NoError(t, err)

	type GetDbSystemPatchRequestInfo struct {
		ContainerId string
		Request     database.GetDbSystemPatchRequest
	}

	var requests []GetDbSystemPatchRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDbSystemPatch(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbSystemPatchHistoryEntry(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetDbSystemPatchHistoryEntry")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbSystemPatchHistoryEntry is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetDbSystemPatchHistoryEntry")
	assert.NoError(t, err)

	type GetDbSystemPatchHistoryEntryRequestInfo struct {
		ContainerId string
		Request     database.GetDbSystemPatchHistoryEntryRequest
	}

	var requests []GetDbSystemPatchHistoryEntryRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDbSystemPatchHistoryEntry(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetExadataIormConfig(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetExadataIormConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetExadataIormConfig is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetExadataIormConfig")
	assert.NoError(t, err)

	type GetExadataIormConfigRequestInfo struct {
		ContainerId string
		Request     database.GetExadataIormConfigRequest
	}

	var requests []GetExadataIormConfigRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetExadataIormConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetExternalBackupJob(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetExternalBackupJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetExternalBackupJob is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetExternalBackupJob")
	assert.NoError(t, err)

	type GetExternalBackupJobRequestInfo struct {
		ContainerId string
		Request     database.GetExternalBackupJobRequest
	}

	var requests []GetExternalBackupJobRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetExternalBackupJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetMaintenanceRun(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "GetMaintenanceRun")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetMaintenanceRun is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "GetMaintenanceRun")
	assert.NoError(t, err)

	type GetMaintenanceRunRequestInfo struct {
		ContainerId string
		Request     database.GetMaintenanceRunRequest
	}

	var requests []GetMaintenanceRunRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetMaintenanceRun(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientLaunchAutonomousDbSystem(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "LaunchAutonomousDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("LaunchAutonomousDbSystem is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "LaunchAutonomousDbSystem")
	assert.NoError(t, err)

	type LaunchAutonomousDbSystemRequestInfo struct {
		ContainerId string
		Request     database.LaunchAutonomousDbSystemRequest
	}

	var requests []LaunchAutonomousDbSystemRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.LaunchAutonomousDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientLaunchDbSystem(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "LaunchDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("LaunchDbSystem is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "LaunchDbSystem")
	assert.NoError(t, err)

	type LaunchDbSystemRequestInfo struct {
		ContainerId string
		Request     database.LaunchDbSystemRequest
	}

	var requests []LaunchDbSystemRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.LaunchDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDataWarehouseBackups(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDataWarehouseBackups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDataWarehouseBackups is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListAutonomousDataWarehouseBackups")
	assert.NoError(t, err)

	type ListAutonomousDataWarehouseBackupsRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDataWarehouseBackupsRequest
	}

	var requests []ListAutonomousDataWarehouseBackupsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDataWarehouseBackupsRequest)
				return c.ListAutonomousDataWarehouseBackups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDataWarehouseBackupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDataWarehouseBackupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDataWarehouses(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDataWarehouses")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDataWarehouses is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListAutonomousDataWarehouses")
	assert.NoError(t, err)

	type ListAutonomousDataWarehousesRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDataWarehousesRequest
	}

	var requests []ListAutonomousDataWarehousesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDataWarehousesRequest)
				return c.ListAutonomousDataWarehouses(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDataWarehousesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDataWarehousesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDatabaseBackups(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDatabaseBackups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDatabaseBackups is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListAutonomousDatabaseBackups")
	assert.NoError(t, err)

	type ListAutonomousDatabaseBackupsRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDatabaseBackupsRequest
	}

	var requests []ListAutonomousDatabaseBackupsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDatabaseBackupsRequest)
				return c.ListAutonomousDatabaseBackups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDatabaseBackupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDatabaseBackupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDatabaseMissionCriticalAssociations(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDatabaseMissionCriticalAssociations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDatabaseMissionCriticalAssociations is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListAutonomousDatabaseMissionCriticalAssociations")
	assert.NoError(t, err)

	type ListAutonomousDatabaseMissionCriticalAssociationsRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDatabaseMissionCriticalAssociationsRequest
	}

	var requests []ListAutonomousDatabaseMissionCriticalAssociationsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDatabaseMissionCriticalAssociationsRequest)
				return c.ListAutonomousDatabaseMissionCriticalAssociations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDatabaseMissionCriticalAssociationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDatabaseMissionCriticalAssociationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDatabases(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDatabases")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDatabases is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListAutonomousDatabases")
	assert.NoError(t, err)

	type ListAutonomousDatabasesRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDatabasesRequest
	}

	var requests []ListAutonomousDatabasesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDatabasesRequest)
				return c.ListAutonomousDatabases(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDatabasesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDatabasesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDbSystemShapes(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDbSystemShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDbSystemShapes is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListAutonomousDbSystemShapes")
	assert.NoError(t, err)

	type ListAutonomousDbSystemShapesRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDbSystemShapesRequest
	}

	var requests []ListAutonomousDbSystemShapesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDbSystemShapesRequest)
				return c.ListAutonomousDbSystemShapes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDbSystemShapesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDbSystemShapesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDbSystems(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDbSystems")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDbSystems is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListAutonomousDbSystems")
	assert.NoError(t, err)

	type ListAutonomousDbSystemsRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDbSystemsRequest
	}

	var requests []ListAutonomousDbSystemsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDbSystemsRequest)
				return c.ListAutonomousDbSystems(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDbSystemsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDbSystemsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousPodMissionCriticalAssociations(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListAutonomousPodMissionCriticalAssociations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousPodMissionCriticalAssociations is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListAutonomousPodMissionCriticalAssociations")
	assert.NoError(t, err)

	type ListAutonomousPodMissionCriticalAssociationsRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousPodMissionCriticalAssociationsRequest
	}

	var requests []ListAutonomousPodMissionCriticalAssociationsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousPodMissionCriticalAssociationsRequest)
				return c.ListAutonomousPodMissionCriticalAssociations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousPodMissionCriticalAssociationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousPodMissionCriticalAssociationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousPods(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListAutonomousPods")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousPods is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListAutonomousPods")
	assert.NoError(t, err)

	type ListAutonomousPodsRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousPodsRequest
	}

	var requests []ListAutonomousPodsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousPodsRequest)
				return c.ListAutonomousPods(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousPodsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousPodsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListBackups(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListBackups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBackups is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListBackups")
	assert.NoError(t, err)

	type ListBackupsRequestInfo struct {
		ContainerId string
		Request     database.ListBackupsRequest
	}

	var requests []ListBackupsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListBackupsRequest)
				return c.ListBackups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListBackupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListBackupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDataGuardAssociations(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListDataGuardAssociations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDataGuardAssociations is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListDataGuardAssociations")
	assert.NoError(t, err)

	type ListDataGuardAssociationsRequestInfo struct {
		ContainerId string
		Request     database.ListDataGuardAssociationsRequest
	}

	var requests []ListDataGuardAssociationsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDataGuardAssociationsRequest)
				return c.ListDataGuardAssociations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDataGuardAssociationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDataGuardAssociationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDatabases(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListDatabases")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDatabases is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListDatabases")
	assert.NoError(t, err)

	type ListDatabasesRequestInfo struct {
		ContainerId string
		Request     database.ListDatabasesRequest
	}

	var requests []ListDatabasesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDatabasesRequest)
				return c.ListDatabases(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDatabasesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDatabasesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbHomePatchHistoryEntries(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListDbHomePatchHistoryEntries")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbHomePatchHistoryEntries is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListDbHomePatchHistoryEntries")
	assert.NoError(t, err)

	type ListDbHomePatchHistoryEntriesRequestInfo struct {
		ContainerId string
		Request     database.ListDbHomePatchHistoryEntriesRequest
	}

	var requests []ListDbHomePatchHistoryEntriesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbHomePatchHistoryEntriesRequest)
				return c.ListDbHomePatchHistoryEntries(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbHomePatchHistoryEntriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbHomePatchHistoryEntriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbHomePatches(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListDbHomePatches")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbHomePatches is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListDbHomePatches")
	assert.NoError(t, err)

	type ListDbHomePatchesRequestInfo struct {
		ContainerId string
		Request     database.ListDbHomePatchesRequest
	}

	var requests []ListDbHomePatchesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbHomePatchesRequest)
				return c.ListDbHomePatches(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbHomePatchesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbHomePatchesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbHomes(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListDbHomes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbHomes is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListDbHomes")
	assert.NoError(t, err)

	type ListDbHomesRequestInfo struct {
		ContainerId string
		Request     database.ListDbHomesRequest
	}

	var requests []ListDbHomesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbHomesRequest)
				return c.ListDbHomes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbHomesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbHomesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbNodes(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListDbNodes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbNodes is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListDbNodes")
	assert.NoError(t, err)

	type ListDbNodesRequestInfo struct {
		ContainerId string
		Request     database.ListDbNodesRequest
	}

	var requests []ListDbNodesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbNodesRequest)
				return c.ListDbNodes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbNodesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbNodesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbSystemPatchHistoryEntries(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListDbSystemPatchHistoryEntries")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbSystemPatchHistoryEntries is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListDbSystemPatchHistoryEntries")
	assert.NoError(t, err)

	type ListDbSystemPatchHistoryEntriesRequestInfo struct {
		ContainerId string
		Request     database.ListDbSystemPatchHistoryEntriesRequest
	}

	var requests []ListDbSystemPatchHistoryEntriesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbSystemPatchHistoryEntriesRequest)
				return c.ListDbSystemPatchHistoryEntries(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbSystemPatchHistoryEntriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbSystemPatchHistoryEntriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbSystemPatches(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListDbSystemPatches")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbSystemPatches is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListDbSystemPatches")
	assert.NoError(t, err)

	type ListDbSystemPatchesRequestInfo struct {
		ContainerId string
		Request     database.ListDbSystemPatchesRequest
	}

	var requests []ListDbSystemPatchesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbSystemPatchesRequest)
				return c.ListDbSystemPatches(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbSystemPatchesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbSystemPatchesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbSystemShapes(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListDbSystemShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbSystemShapes is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListDbSystemShapes")
	assert.NoError(t, err)

	type ListDbSystemShapesRequestInfo struct {
		ContainerId string
		Request     database.ListDbSystemShapesRequest
	}

	var requests []ListDbSystemShapesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbSystemShapesRequest)
				return c.ListDbSystemShapes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbSystemShapesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbSystemShapesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbSystems(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListDbSystems")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbSystems is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListDbSystems")
	assert.NoError(t, err)

	type ListDbSystemsRequestInfo struct {
		ContainerId string
		Request     database.ListDbSystemsRequest
	}

	var requests []ListDbSystemsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbSystemsRequest)
				return c.ListDbSystems(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbSystemsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbSystemsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbVersions(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListDbVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbVersions is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListDbVersions")
	assert.NoError(t, err)

	type ListDbVersionsRequestInfo struct {
		ContainerId string
		Request     database.ListDbVersionsRequest
	}

	var requests []ListDbVersionsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDbVersionsRequest)
				return c.ListDbVersions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDbVersionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDbVersionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListMaintenanceRuns(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ListMaintenanceRuns")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListMaintenanceRuns is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ListMaintenanceRuns")
	assert.NoError(t, err)

	type ListMaintenanceRunsRequestInfo struct {
		ContainerId string
		Request     database.ListMaintenanceRunsRequest
	}

	var requests []ListMaintenanceRunsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListMaintenanceRunsRequest)
				return c.ListMaintenanceRuns(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListMaintenanceRunsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListMaintenanceRunsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientReinstateDataGuardAssociation(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "ReinstateDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ReinstateDataGuardAssociation is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "ReinstateDataGuardAssociation")
	assert.NoError(t, err)

	type ReinstateDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.ReinstateDataGuardAssociationRequest
	}

	var requests []ReinstateDataGuardAssociationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ReinstateDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientRestoreAutonomousDataWarehouse(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "RestoreAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestoreAutonomousDataWarehouse is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "RestoreAutonomousDataWarehouse")
	assert.NoError(t, err)

	type RestoreAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.RestoreAutonomousDataWarehouseRequest
	}

	var requests []RestoreAutonomousDataWarehouseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.RestoreAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientRestoreAutonomousDatabase(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "RestoreAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestoreAutonomousDatabase is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "RestoreAutonomousDatabase")
	assert.NoError(t, err)

	type RestoreAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.RestoreAutonomousDatabaseRequest
	}

	var requests []RestoreAutonomousDatabaseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.RestoreAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientRestoreDatabase(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "RestoreDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestoreDatabase is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "RestoreDatabase")
	assert.NoError(t, err)

	type RestoreDatabaseRequestInfo struct {
		ContainerId string
		Request     database.RestoreDatabaseRequest
	}

	var requests []RestoreDatabaseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.RestoreDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientStartAutonomousDataWarehouse(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "StartAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StartAutonomousDataWarehouse is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "StartAutonomousDataWarehouse")
	assert.NoError(t, err)

	type StartAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.StartAutonomousDataWarehouseRequest
	}

	var requests []StartAutonomousDataWarehouseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.StartAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientStartAutonomousDatabase(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "StartAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StartAutonomousDatabase is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "StartAutonomousDatabase")
	assert.NoError(t, err)

	type StartAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.StartAutonomousDatabaseRequest
	}

	var requests []StartAutonomousDatabaseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.StartAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientStopAutonomousDataWarehouse(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "StopAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StopAutonomousDataWarehouse is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "StopAutonomousDataWarehouse")
	assert.NoError(t, err)

	type StopAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.StopAutonomousDataWarehouseRequest
	}

	var requests []StopAutonomousDataWarehouseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.StopAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientStopAutonomousDatabase(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "StopAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StopAutonomousDatabase is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "StopAutonomousDatabase")
	assert.NoError(t, err)

	type StopAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.StopAutonomousDatabaseRequest
	}

	var requests []StopAutonomousDatabaseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.StopAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientSwitchoverAutonomousPodMissionCriticalAssociation(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "SwitchoverAutonomousPodMissionCriticalAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SwitchoverAutonomousPodMissionCriticalAssociation is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "SwitchoverAutonomousPodMissionCriticalAssociation")
	assert.NoError(t, err)

	type SwitchoverAutonomousPodMissionCriticalAssociationRequestInfo struct {
		ContainerId string
		Request     database.SwitchoverAutonomousPodMissionCriticalAssociationRequest
	}

	var requests []SwitchoverAutonomousPodMissionCriticalAssociationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.SwitchoverAutonomousPodMissionCriticalAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientSwitchoverDataGuardAssociation(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "SwitchoverDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SwitchoverDataGuardAssociation is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "SwitchoverDataGuardAssociation")
	assert.NoError(t, err)

	type SwitchoverDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.SwitchoverDataGuardAssociationRequest
	}

	var requests []SwitchoverDataGuardAssociationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.SwitchoverDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientTerminateAutonomousDbSystem(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "TerminateAutonomousDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("TerminateAutonomousDbSystem is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "TerminateAutonomousDbSystem")
	assert.NoError(t, err)

	type TerminateAutonomousDbSystemRequestInfo struct {
		ContainerId string
		Request     database.TerminateAutonomousDbSystemRequest
	}

	var requests []TerminateAutonomousDbSystemRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.TerminateAutonomousDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientTerminateAutonomousPod(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "TerminateAutonomousPod")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("TerminateAutonomousPod is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "TerminateAutonomousPod")
	assert.NoError(t, err)

	type TerminateAutonomousPodRequestInfo struct {
		ContainerId string
		Request     database.TerminateAutonomousPodRequest
	}

	var requests []TerminateAutonomousPodRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.TerminateAutonomousPod(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientTerminateDbSystem(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "TerminateDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("TerminateDbSystem is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "TerminateDbSystem")
	assert.NoError(t, err)

	type TerminateDbSystemRequestInfo struct {
		ContainerId string
		Request     database.TerminateDbSystemRequest
	}

	var requests []TerminateDbSystemRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.TerminateDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateAutonomousDataWarehouse(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "UpdateAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutonomousDataWarehouse is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "UpdateAutonomousDataWarehouse")
	assert.NoError(t, err)

	type UpdateAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.UpdateAutonomousDataWarehouseRequest
	}

	var requests []UpdateAutonomousDataWarehouseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateAutonomousDatabase(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "UpdateAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutonomousDatabase is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "UpdateAutonomousDatabase")
	assert.NoError(t, err)

	type UpdateAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.UpdateAutonomousDatabaseRequest
	}

	var requests []UpdateAutonomousDatabaseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateAutonomousDbSystem(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "UpdateAutonomousDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutonomousDbSystem is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "UpdateAutonomousDbSystem")
	assert.NoError(t, err)

	type UpdateAutonomousDbSystemRequestInfo struct {
		ContainerId string
		Request     database.UpdateAutonomousDbSystemRequest
	}

	var requests []UpdateAutonomousDbSystemRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateAutonomousDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateAutonomousPod(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "UpdateAutonomousPod")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutonomousPod is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "UpdateAutonomousPod")
	assert.NoError(t, err)

	type UpdateAutonomousPodRequestInfo struct {
		ContainerId string
		Request     database.UpdateAutonomousPodRequest
	}

	var requests []UpdateAutonomousPodRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateAutonomousPod(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateDatabase(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "UpdateDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDatabase is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "UpdateDatabase")
	assert.NoError(t, err)

	type UpdateDatabaseRequestInfo struct {
		ContainerId string
		Request     database.UpdateDatabaseRequest
	}

	var requests []UpdateDatabaseRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateDbHome(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "UpdateDbHome")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDbHome is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "UpdateDbHome")
	assert.NoError(t, err)

	type UpdateDbHomeRequestInfo struct {
		ContainerId string
		Request     database.UpdateDbHomeRequest
	}

	var requests []UpdateDbHomeRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateDbHome(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateDbSystem(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "UpdateDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDbSystem is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "UpdateDbSystem")
	assert.NoError(t, err)

	type UpdateDbSystemRequestInfo struct {
		ContainerId string
		Request     database.UpdateDbSystemRequest
	}

	var requests []UpdateDbSystemRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateExadataIormConfig(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "UpdateExadataIormConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateExadataIormConfig is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "UpdateExadataIormConfig")
	assert.NoError(t, err)

	type UpdateExadataIormConfigRequestInfo struct {
		ContainerId string
		Request     database.UpdateExadataIormConfigRequest
	}

	var requests []UpdateExadataIormConfigRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateExadataIormConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateMaintenanceRun(t *testing.T) {
	enabled, err := testClient.isApiEnabled("database", "UpdateMaintenanceRun")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateMaintenanceRun is not enabled by the testing service")
	}
	c, err := database.NewDatabaseClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("database", "UpdateMaintenanceRun")
	assert.NoError(t, err)

	type UpdateMaintenanceRunRequestInfo struct {
		ContainerId string
		Request     database.UpdateMaintenanceRunRequest
	}

	var requests []UpdateMaintenanceRunRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateMaintenanceRun(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}