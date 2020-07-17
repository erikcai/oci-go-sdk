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

func createDatabaseClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := database.NewDatabaseClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientActivateExadataInfrastructure(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ActivateExadataInfrastructure")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ActivateExadataInfrastructure is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ActivateExadataInfrastructure", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ActivateExadataInfrastructure")
	assert.NoError(t, err)

	type ActivateExadataInfrastructureRequestInfo struct {
		ContainerId string
		Request     database.ActivateExadataInfrastructureRequest
	}

	var requests []ActivateExadataInfrastructureRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ActivateExadataInfrastructure(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientAutonomousDatabaseManualRefresh(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "AutonomousDatabaseManualRefresh")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AutonomousDatabaseManualRefresh is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "AutonomousDatabaseManualRefresh", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "AutonomousDatabaseManualRefresh")
	assert.NoError(t, err)

	type AutonomousDatabaseManualRefreshRequestInfo struct {
		ContainerId string
		Request     database.AutonomousDatabaseManualRefreshRequest
	}

	var requests []AutonomousDatabaseManualRefreshRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.AutonomousDatabaseManualRefresh(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientChangeAutonomousContainerDatabaseCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ChangeAutonomousContainerDatabaseCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeAutonomousContainerDatabaseCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ChangeAutonomousContainerDatabaseCompartment", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ChangeAutonomousContainerDatabaseCompartment")
	assert.NoError(t, err)

	type ChangeAutonomousContainerDatabaseCompartmentRequestInfo struct {
		ContainerId string
		Request     database.ChangeAutonomousContainerDatabaseCompartmentRequest
	}

	var requests []ChangeAutonomousContainerDatabaseCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeAutonomousContainerDatabaseCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientChangeAutonomousDatabaseCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ChangeAutonomousDatabaseCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeAutonomousDatabaseCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ChangeAutonomousDatabaseCompartment", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ChangeAutonomousDatabaseCompartment")
	assert.NoError(t, err)

	type ChangeAutonomousDatabaseCompartmentRequestInfo struct {
		ContainerId string
		Request     database.ChangeAutonomousDatabaseCompartmentRequest
	}

	var requests []ChangeAutonomousDatabaseCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeAutonomousDatabaseCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientChangeAutonomousExadataInfrastructureCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ChangeAutonomousExadataInfrastructureCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeAutonomousExadataInfrastructureCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ChangeAutonomousExadataInfrastructureCompartment", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ChangeAutonomousExadataInfrastructureCompartment")
	assert.NoError(t, err)

	type ChangeAutonomousExadataInfrastructureCompartmentRequestInfo struct {
		ContainerId string
		Request     database.ChangeAutonomousExadataInfrastructureCompartmentRequest
	}

	var requests []ChangeAutonomousExadataInfrastructureCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeAutonomousExadataInfrastructureCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientChangeAutonomousVmClusterCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ChangeAutonomousVmClusterCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeAutonomousVmClusterCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ChangeAutonomousVmClusterCompartment", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ChangeAutonomousVmClusterCompartment")
	assert.NoError(t, err)

	type ChangeAutonomousVmClusterCompartmentRequestInfo struct {
		ContainerId string
		Request     database.ChangeAutonomousVmClusterCompartmentRequest
	}

	var requests []ChangeAutonomousVmClusterCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeAutonomousVmClusterCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientChangeBackupDestinationCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ChangeBackupDestinationCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeBackupDestinationCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ChangeBackupDestinationCompartment", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ChangeBackupDestinationCompartment")
	assert.NoError(t, err)

	type ChangeBackupDestinationCompartmentRequestInfo struct {
		ContainerId string
		Request     database.ChangeBackupDestinationCompartmentRequest
	}

	var requests []ChangeBackupDestinationCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeBackupDestinationCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCS" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientChangeCloudExadataInfrastructureCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ChangeCloudExadataInfrastructureCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeCloudExadataInfrastructureCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ChangeCloudExadataInfrastructureCompartment", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ChangeCloudExadataInfrastructureCompartment")
	assert.NoError(t, err)

	type ChangeCloudExadataInfrastructureCompartmentRequestInfo struct {
		ContainerId string
		Request     database.ChangeCloudExadataInfrastructureCompartmentRequest
	}

	var requests []ChangeCloudExadataInfrastructureCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeCloudExadataInfrastructureCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCS" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientChangeCloudVmClusterCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ChangeCloudVmClusterCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeCloudVmClusterCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ChangeCloudVmClusterCompartment", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ChangeCloudVmClusterCompartment")
	assert.NoError(t, err)

	type ChangeCloudVmClusterCompartmentRequestInfo struct {
		ContainerId string
		Request     database.ChangeCloudVmClusterCompartmentRequest
	}

	var requests []ChangeCloudVmClusterCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeCloudVmClusterCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientChangeDatabaseSoftwareImageCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ChangeDatabaseSoftwareImageCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeDatabaseSoftwareImageCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ChangeDatabaseSoftwareImageCompartment", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ChangeDatabaseSoftwareImageCompartment")
	assert.NoError(t, err)

	type ChangeDatabaseSoftwareImageCompartmentRequestInfo struct {
		ContainerId string
		Request     database.ChangeDatabaseSoftwareImageCompartmentRequest
	}

	var requests []ChangeDatabaseSoftwareImageCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeDatabaseSoftwareImageCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientChangeDbSystemCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ChangeDbSystemCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeDbSystemCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ChangeDbSystemCompartment", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ChangeDbSystemCompartment")
	assert.NoError(t, err)

	type ChangeDbSystemCompartmentRequestInfo struct {
		ContainerId string
		Request     database.ChangeDbSystemCompartmentRequest
	}

	var requests []ChangeDbSystemCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeDbSystemCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientChangeExadataInfrastructureCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ChangeExadataInfrastructureCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeExadataInfrastructureCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ChangeExadataInfrastructureCompartment", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ChangeExadataInfrastructureCompartment")
	assert.NoError(t, err)

	type ChangeExadataInfrastructureCompartmentRequestInfo struct {
		ContainerId string
		Request     database.ChangeExadataInfrastructureCompartmentRequest
	}

	var requests []ChangeExadataInfrastructureCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeExadataInfrastructureCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientChangeVmClusterCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ChangeVmClusterCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeVmClusterCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ChangeVmClusterCompartment", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ChangeVmClusterCompartment")
	assert.NoError(t, err)

	type ChangeVmClusterCompartmentRequestInfo struct {
		ContainerId string
		Request     database.ChangeVmClusterCompartmentRequest
	}

	var requests []ChangeVmClusterCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeVmClusterCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCompleteExternalBackupJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CompleteExternalBackupJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CompleteExternalBackupJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CompleteExternalBackupJob", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CompleteExternalBackupJob")
	assert.NoError(t, err)

	type CompleteExternalBackupJobRequestInfo struct {
		ContainerId string
		Request     database.CompleteExternalBackupJobRequest
	}

	var requests []CompleteExternalBackupJobRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CompleteExternalBackupJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousContainerDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousContainerDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousContainerDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateAutonomousContainerDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateAutonomousContainerDatabase")
	assert.NoError(t, err)

	type CreateAutonomousContainerDatabaseRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousContainerDatabaseRequest
	}

	var requests []CreateAutonomousContainerDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateAutonomousContainerDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousDataWarehouse(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousDataWarehouse is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateAutonomousDataWarehouse", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateAutonomousDataWarehouse")
	assert.NoError(t, err)

	type CreateAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousDataWarehouseRequest
	}

	var requests []CreateAutonomousDataWarehouseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousDataWarehouseBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousDataWarehouseBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousDataWarehouseBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateAutonomousDataWarehouseBackup", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateAutonomousDataWarehouseBackup")
	assert.NoError(t, err)

	type CreateAutonomousDataWarehouseBackupRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousDataWarehouseBackupRequest
	}

	var requests []CreateAutonomousDataWarehouseBackupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateAutonomousDataWarehouseBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateAutonomousDatabase")
	assert.NoError(t, err)

	type CreateAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousDatabaseRequest
	}

	var requests []CreateAutonomousDatabaseRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateAutonomousDatabaseRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateAutonomousDatabaseBase"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "source",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"DATABASE":              &database.CreateAutonomousDatabaseCloneDetails{},
				"CLONE_TO_REFRESHABLE":  &database.CreateRefreshableAutonomousDatabaseCloneDetails{},
				"BACKUP_FROM_ID":        &database.CreateAutonomousDatabaseFromBackupDetails{},
				"BACKUP_FROM_TIMESTAMP": &database.CreateAutonomousDatabaseFromBackupTimestampDetails{},
				"NONE":                  &database.CreateAutonomousDatabaseDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousDatabaseBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousDatabaseBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousDatabaseBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateAutonomousDatabaseBackup", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateAutonomousDatabaseBackup")
	assert.NoError(t, err)

	type CreateAutonomousDatabaseBackupRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousDatabaseBackupRequest
	}

	var requests []CreateAutonomousDatabaseBackupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateAutonomousDatabaseBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateAutonomousVmCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateAutonomousVmCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutonomousVmCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateAutonomousVmCluster", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateAutonomousVmCluster")
	assert.NoError(t, err)

	type CreateAutonomousVmClusterRequestInfo struct {
		ContainerId string
		Request     database.CreateAutonomousVmClusterRequest
	}

	var requests []CreateAutonomousVmClusterRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateAutonomousVmCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateBackup", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateBackup")
	assert.NoError(t, err)

	type CreateBackupRequestInfo struct {
		ContainerId string
		Request     database.CreateBackupRequest
	}

	var requests []CreateBackupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateBackupDestination(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateBackupDestination")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBackupDestination is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateBackupDestination", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateBackupDestination")
	assert.NoError(t, err)

	type CreateBackupDestinationRequestInfo struct {
		ContainerId string
		Request     database.CreateBackupDestinationRequest
	}

	var requests []CreateBackupDestinationRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateBackupDestinationRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateBackupDestinationDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "type",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"NFS":                &database.CreateNfsBackupDestinationDetails{},
				"RECOVERY_APPLIANCE": &database.CreateRecoveryApplianceBackupDestinationDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateBackupDestination(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCS" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateCloudExadataInfrastructure(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateCloudExadataInfrastructure")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCloudExadataInfrastructure is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateCloudExadataInfrastructure", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateCloudExadataInfrastructure")
	assert.NoError(t, err)

	type CreateCloudExadataInfrastructureRequestInfo struct {
		ContainerId string
		Request     database.CreateCloudExadataInfrastructureRequest
	}

	var requests []CreateCloudExadataInfrastructureRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateCloudExadataInfrastructure(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCS" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateCloudVmCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateCloudVmCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCloudVmCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateCloudVmCluster", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateCloudVmCluster")
	assert.NoError(t, err)

	type CreateCloudVmClusterRequestInfo struct {
		ContainerId string
		Request     database.CreateCloudVmClusterRequest
	}

	var requests []CreateCloudVmClusterRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateCloudVmCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateConsoleConnection(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateConsoleConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateConsoleConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateConsoleConnection", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateConsoleConnection")
	assert.NoError(t, err)

	type CreateConsoleConnectionRequestInfo struct {
		ContainerId string
		Request     database.CreateConsoleConnectionRequest
	}

	var requests []CreateConsoleConnectionRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateConsoleConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateDataGuardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDataGuardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateDataGuardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateDataGuardAssociation")
	assert.NoError(t, err)

	type CreateDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.CreateDataGuardAssociationRequest
	}

	var requests []CreateDataGuardAssociationRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateDataGuardAssociationRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateDataGuardAssociationDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "creationType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"NewDbSystem":       &database.CreateDataGuardAssociationWithNewDbSystemDetails{},
				"ExistingVmCluster": &database.CreateDataGuardAssociationToExistingVmClusterDetails{},
				"ExistingDbSystem":  &database.CreateDataGuardAssociationToExistingDbSystemDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateDatabase")
	assert.NoError(t, err)

	type CreateDatabaseRequestInfo struct {
		ContainerId string
		Request     database.CreateDatabaseRequest
	}

	var requests []CreateDatabaseRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateDatabaseRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateDatabaseBase"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "source",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"NONE":      &database.CreateNewDatabaseDetails{},
				"DB_BACKUP": &database.CreateDatabaseFromBackup{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateDatabaseSoftwareImage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateDatabaseSoftwareImage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDatabaseSoftwareImage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateDatabaseSoftwareImage", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateDatabaseSoftwareImage")
	assert.NoError(t, err)

	type CreateDatabaseSoftwareImageRequestInfo struct {
		ContainerId string
		Request     database.CreateDatabaseSoftwareImageRequest
	}

	var requests []CreateDatabaseSoftwareImageRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateDatabaseSoftwareImage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateDbHome(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateDbHome")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDbHome is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateDbHome", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateDbHome")
	assert.NoError(t, err)

	type CreateDbHomeRequestInfo struct {
		ContainerId string
		Request     database.CreateDbHomeRequest
	}

	var requests []CreateDbHomeRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateDbHomeRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateDbHomeBase"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "source",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"DATABASE":          &database.CreateDbHomeWithDbSystemIdFromDatabaseDetails{},
				"DB_BACKUP":         &database.CreateDbHomeWithDbSystemIdFromBackupDetails{},
				"VM_CLUSTER_BACKUP": &database.CreateDbHomeWithVmClusterIdFromBackupDetails{},
				"NONE":              &database.CreateDbHomeWithDbSystemIdDetails{},
				"VM_CLUSTER_NEW":    &database.CreateDbHomeWithVmClusterIdDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateDbHome(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateExadataInfrastructure(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateExadataInfrastructure")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateExadataInfrastructure is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateExadataInfrastructure", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateExadataInfrastructure")
	assert.NoError(t, err)

	type CreateExadataInfrastructureRequestInfo struct {
		ContainerId string
		Request     database.CreateExadataInfrastructureRequest
	}

	var requests []CreateExadataInfrastructureRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateExadataInfrastructure(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateExternalBackupJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateExternalBackupJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateExternalBackupJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateExternalBackupJob", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateExternalBackupJob")
	assert.NoError(t, err)

	type CreateExternalBackupJobRequestInfo struct {
		ContainerId string
		Request     database.CreateExternalBackupJobRequest
	}

	var requests []CreateExternalBackupJobRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateExternalBackupJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateVmCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateVmCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateVmCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateVmCluster", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateVmCluster")
	assert.NoError(t, err)

	type CreateVmClusterRequestInfo struct {
		ContainerId string
		Request     database.CreateVmClusterRequest
	}

	var requests []CreateVmClusterRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateVmCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientCreateVmClusterNetwork(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "CreateVmClusterNetwork")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateVmClusterNetwork is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "CreateVmClusterNetwork", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "CreateVmClusterNetwork")
	assert.NoError(t, err)

	type CreateVmClusterNetworkRequestInfo struct {
		ContainerId string
		Request     database.CreateVmClusterNetworkRequest
	}

	var requests []CreateVmClusterNetworkRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateVmClusterNetwork(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDbNodeAction(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DbNodeAction")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DbNodeAction is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DbNodeAction", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DbNodeAction")
	assert.NoError(t, err)

	type DbNodeActionRequestInfo struct {
		ContainerId string
		Request     database.DbNodeActionRequest
	}

	var requests []DbNodeActionRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DbNodeAction(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteAutonomousDataWarehouse(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAutonomousDataWarehouse is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteAutonomousDataWarehouse", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteAutonomousDataWarehouse")
	assert.NoError(t, err)

	type DeleteAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.DeleteAutonomousDataWarehouseRequest
	}

	var requests []DeleteAutonomousDataWarehouseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteAutonomousDatabase")
	assert.NoError(t, err)

	type DeleteAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.DeleteAutonomousDatabaseRequest
	}

	var requests []DeleteAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteAutonomousVmCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteAutonomousVmCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAutonomousVmCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteAutonomousVmCluster", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteAutonomousVmCluster")
	assert.NoError(t, err)

	type DeleteAutonomousVmClusterRequestInfo struct {
		ContainerId string
		Request     database.DeleteAutonomousVmClusterRequest
	}

	var requests []DeleteAutonomousVmClusterRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteAutonomousVmCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteBackup", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteBackup")
	assert.NoError(t, err)

	type DeleteBackupRequestInfo struct {
		ContainerId string
		Request     database.DeleteBackupRequest
	}

	var requests []DeleteBackupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteBackupDestination(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteBackupDestination")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBackupDestination is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteBackupDestination", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteBackupDestination")
	assert.NoError(t, err)

	type DeleteBackupDestinationRequestInfo struct {
		ContainerId string
		Request     database.DeleteBackupDestinationRequest
	}

	var requests []DeleteBackupDestinationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteBackupDestination(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCS" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteCloudExadataInfrastructure(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteCloudExadataInfrastructure")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCloudExadataInfrastructure is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteCloudExadataInfrastructure", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteCloudExadataInfrastructure")
	assert.NoError(t, err)

	type DeleteCloudExadataInfrastructureRequestInfo struct {
		ContainerId string
		Request     database.DeleteCloudExadataInfrastructureRequest
	}

	var requests []DeleteCloudExadataInfrastructureRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteCloudExadataInfrastructure(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCS" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteCloudVmCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteCloudVmCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCloudVmCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteCloudVmCluster", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteCloudVmCluster")
	assert.NoError(t, err)

	type DeleteCloudVmClusterRequestInfo struct {
		ContainerId string
		Request     database.DeleteCloudVmClusterRequest
	}

	var requests []DeleteCloudVmClusterRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteCloudVmCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteConsoleConnection(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteConsoleConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteConsoleConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteConsoleConnection", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteConsoleConnection")
	assert.NoError(t, err)

	type DeleteConsoleConnectionRequestInfo struct {
		ContainerId string
		Request     database.DeleteConsoleConnectionRequest
	}

	var requests []DeleteConsoleConnectionRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteConsoleConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteDatabase")
	assert.NoError(t, err)

	type DeleteDatabaseRequestInfo struct {
		ContainerId string
		Request     database.DeleteDatabaseRequest
	}

	var requests []DeleteDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteDatabaseSoftwareImage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteDatabaseSoftwareImage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDatabaseSoftwareImage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteDatabaseSoftwareImage", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteDatabaseSoftwareImage")
	assert.NoError(t, err)

	type DeleteDatabaseSoftwareImageRequestInfo struct {
		ContainerId string
		Request     database.DeleteDatabaseSoftwareImageRequest
	}

	var requests []DeleteDatabaseSoftwareImageRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteDatabaseSoftwareImage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteDbHome(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteDbHome")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDbHome is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteDbHome", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteDbHome")
	assert.NoError(t, err)

	type DeleteDbHomeRequestInfo struct {
		ContainerId string
		Request     database.DeleteDbHomeRequest
	}

	var requests []DeleteDbHomeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteDbHome(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteExadataInfrastructure(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteExadataInfrastructure")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteExadataInfrastructure is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteExadataInfrastructure", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteExadataInfrastructure")
	assert.NoError(t, err)

	type DeleteExadataInfrastructureRequestInfo struct {
		ContainerId string
		Request     database.DeleteExadataInfrastructureRequest
	}

	var requests []DeleteExadataInfrastructureRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteExadataInfrastructure(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteVmCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteVmCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteVmCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteVmCluster", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteVmCluster")
	assert.NoError(t, err)

	type DeleteVmClusterRequestInfo struct {
		ContainerId string
		Request     database.DeleteVmClusterRequest
	}

	var requests []DeleteVmClusterRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteVmCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeleteVmClusterNetwork(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeleteVmClusterNetwork")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteVmClusterNetwork is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeleteVmClusterNetwork", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeleteVmClusterNetwork")
	assert.NoError(t, err)

	type DeleteVmClusterNetworkRequestInfo struct {
		ContainerId string
		Request     database.DeleteVmClusterNetworkRequest
	}

	var requests []DeleteVmClusterNetworkRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteVmClusterNetwork(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDeregisterAutonomousDatabaseDataSafe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DeregisterAutonomousDatabaseDataSafe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeregisterAutonomousDatabaseDataSafe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DeregisterAutonomousDatabaseDataSafe", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DeregisterAutonomousDatabaseDataSafe")
	assert.NoError(t, err)

	type DeregisterAutonomousDatabaseDataSafeRequestInfo struct {
		ContainerId string
		Request     database.DeregisterAutonomousDatabaseDataSafeRequest
	}

	var requests []DeregisterAutonomousDatabaseDataSafeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeregisterAutonomousDatabaseDataSafe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDisableAutonomousDatabaseOperationsInsights(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DisableAutonomousDatabaseOperationsInsights")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DisableAutonomousDatabaseOperationsInsights is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DisableAutonomousDatabaseOperationsInsights", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DisableAutonomousDatabaseOperationsInsights")
	assert.NoError(t, err)

	type DisableAutonomousDatabaseOperationsInsightsRequestInfo struct {
		ContainerId string
		Request     database.DisableAutonomousDatabaseOperationsInsightsRequest
	}

	var requests []DisableAutonomousDatabaseOperationsInsightsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DisableAutonomousDatabaseOperationsInsights(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDownloadExadataInfrastructureConfigFile(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DownloadExadataInfrastructureConfigFile")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DownloadExadataInfrastructureConfigFile is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DownloadExadataInfrastructureConfigFile", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DownloadExadataInfrastructureConfigFile")
	assert.NoError(t, err)

	type DownloadExadataInfrastructureConfigFileRequestInfo struct {
		ContainerId string
		Request     database.DownloadExadataInfrastructureConfigFileRequest
	}

	var requests []DownloadExadataInfrastructureConfigFileRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DownloadExadataInfrastructureConfigFile(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientDownloadVmClusterNetworkConfigFile(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "DownloadVmClusterNetworkConfigFile")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DownloadVmClusterNetworkConfigFile is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "DownloadVmClusterNetworkConfigFile", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "DownloadVmClusterNetworkConfigFile")
	assert.NoError(t, err)

	type DownloadVmClusterNetworkConfigFileRequestInfo struct {
		ContainerId string
		Request     database.DownloadVmClusterNetworkConfigFileRequest
	}

	var requests []DownloadVmClusterNetworkConfigFileRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DownloadVmClusterNetworkConfigFile(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientEnableAutonomousDatabaseOperationsInsights(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "EnableAutonomousDatabaseOperationsInsights")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("EnableAutonomousDatabaseOperationsInsights is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "EnableAutonomousDatabaseOperationsInsights", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "EnableAutonomousDatabaseOperationsInsights")
	assert.NoError(t, err)

	type EnableAutonomousDatabaseOperationsInsightsRequestInfo struct {
		ContainerId string
		Request     database.EnableAutonomousDatabaseOperationsInsightsRequest
	}

	var requests []EnableAutonomousDatabaseOperationsInsightsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.EnableAutonomousDatabaseOperationsInsights(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientFailOverAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "FailOverAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("FailOverAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "FailOverAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "FailOverAutonomousDatabase")
	assert.NoError(t, err)

	type FailOverAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.FailOverAutonomousDatabaseRequest
	}

	var requests []FailOverAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.FailOverAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientFailoverAutonomousContainerDatabaseDataguardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "FailoverAutonomousContainerDatabaseDataguardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("FailoverAutonomousContainerDatabaseDataguardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "FailoverAutonomousContainerDatabaseDataguardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "FailoverAutonomousContainerDatabaseDataguardAssociation")
	assert.NoError(t, err)

	type FailoverAutonomousContainerDatabaseDataguardAssociationRequestInfo struct {
		ContainerId string
		Request     database.FailoverAutonomousContainerDatabaseDataguardAssociationRequest
	}

	var requests []FailoverAutonomousContainerDatabaseDataguardAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.FailoverAutonomousContainerDatabaseDataguardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientFailoverDataGuardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "FailoverDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("FailoverDataGuardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "FailoverDataGuardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "FailoverDataGuardAssociation")
	assert.NoError(t, err)

	type FailoverDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.FailoverDataGuardAssociationRequest
	}

	var requests []FailoverDataGuardAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.FailoverDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGenerateAutonomousDataWarehouseWallet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GenerateAutonomousDataWarehouseWallet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GenerateAutonomousDataWarehouseWallet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GenerateAutonomousDataWarehouseWallet", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GenerateAutonomousDataWarehouseWallet")
	assert.NoError(t, err)

	type GenerateAutonomousDataWarehouseWalletRequestInfo struct {
		ContainerId string
		Request     database.GenerateAutonomousDataWarehouseWalletRequest
	}

	var requests []GenerateAutonomousDataWarehouseWalletRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GenerateAutonomousDataWarehouseWallet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGenerateAutonomousDatabaseWallet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GenerateAutonomousDatabaseWallet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GenerateAutonomousDatabaseWallet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GenerateAutonomousDatabaseWallet", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GenerateAutonomousDatabaseWallet")
	assert.NoError(t, err)

	type GenerateAutonomousDatabaseWalletRequestInfo struct {
		ContainerId string
		Request     database.GenerateAutonomousDatabaseWalletRequest
	}

	var requests []GenerateAutonomousDatabaseWalletRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GenerateAutonomousDatabaseWallet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGenerateRecommendedVmClusterNetwork(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GenerateRecommendedVmClusterNetwork")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GenerateRecommendedVmClusterNetwork is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GenerateRecommendedVmClusterNetwork", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GenerateRecommendedVmClusterNetwork")
	assert.NoError(t, err)

	type GenerateRecommendedVmClusterNetworkRequestInfo struct {
		ContainerId string
		Request     database.GenerateRecommendedVmClusterNetworkRequest
	}

	var requests []GenerateRecommendedVmClusterNetworkRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GenerateRecommendedVmClusterNetwork(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousContainerDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousContainerDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousContainerDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousContainerDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousContainerDatabase")
	assert.NoError(t, err)

	type GetAutonomousContainerDatabaseRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousContainerDatabaseRequest
	}

	var requests []GetAutonomousContainerDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetAutonomousContainerDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousContainerDatabaseDataguardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousContainerDatabaseDataguardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousContainerDatabaseDataguardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousContainerDatabaseDataguardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousContainerDatabaseDataguardAssociation")
	assert.NoError(t, err)

	type GetAutonomousContainerDatabaseDataguardAssociationRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousContainerDatabaseDataguardAssociationRequest
	}

	var requests []GetAutonomousContainerDatabaseDataguardAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetAutonomousContainerDatabaseDataguardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDataWarehouse(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDataWarehouse is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousDataWarehouse", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousDataWarehouse")
	assert.NoError(t, err)

	type GetAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDataWarehouseRequest
	}

	var requests []GetAutonomousDataWarehouseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDataWarehouseBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDataWarehouseBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDataWarehouseBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousDataWarehouseBackup", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousDataWarehouseBackup")
	assert.NoError(t, err)

	type GetAutonomousDataWarehouseBackupRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDataWarehouseBackupRequest
	}

	var requests []GetAutonomousDataWarehouseBackupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetAutonomousDataWarehouseBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousDatabase")
	assert.NoError(t, err)

	type GetAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDatabaseRequest
	}

	var requests []GetAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDatabaseBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDatabaseBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDatabaseBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousDatabaseBackup", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousDatabaseBackup")
	assert.NoError(t, err)

	type GetAutonomousDatabaseBackupRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDatabaseBackupRequest
	}

	var requests []GetAutonomousDatabaseBackupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetAutonomousDatabaseBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDatabaseDataguardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDatabaseDataguardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDatabaseDataguardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousDatabaseDataguardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousDatabaseDataguardAssociation")
	assert.NoError(t, err)

	type GetAutonomousDatabaseDataguardAssociationRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDatabaseDataguardAssociationRequest
	}

	var requests []GetAutonomousDatabaseDataguardAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetAutonomousDatabaseDataguardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDatabaseRegionalWallet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDatabaseRegionalWallet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDatabaseRegionalWallet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousDatabaseRegionalWallet", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousDatabaseRegionalWallet")
	assert.NoError(t, err)

	type GetAutonomousDatabaseRegionalWalletRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDatabaseRegionalWalletRequest
	}

	var requests []GetAutonomousDatabaseRegionalWalletRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetAutonomousDatabaseRegionalWallet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousDatabaseWallet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousDatabaseWallet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousDatabaseWallet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousDatabaseWallet", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousDatabaseWallet")
	assert.NoError(t, err)

	type GetAutonomousDatabaseWalletRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousDatabaseWalletRequest
	}

	var requests []GetAutonomousDatabaseWalletRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetAutonomousDatabaseWallet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousExadataInfrastructure(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousExadataInfrastructure")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousExadataInfrastructure is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousExadataInfrastructure", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousExadataInfrastructure")
	assert.NoError(t, err)

	type GetAutonomousExadataInfrastructureRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousExadataInfrastructureRequest
	}

	var requests []GetAutonomousExadataInfrastructureRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetAutonomousExadataInfrastructure(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetAutonomousVmCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetAutonomousVmCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutonomousVmCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetAutonomousVmCluster", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetAutonomousVmCluster")
	assert.NoError(t, err)

	type GetAutonomousVmClusterRequestInfo struct {
		ContainerId string
		Request     database.GetAutonomousVmClusterRequest
	}

	var requests []GetAutonomousVmClusterRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetAutonomousVmCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetBackup", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetBackup")
	assert.NoError(t, err)

	type GetBackupRequestInfo struct {
		ContainerId string
		Request     database.GetBackupRequest
	}

	var requests []GetBackupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetBackupDestination(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetBackupDestination")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBackupDestination is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetBackupDestination", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetBackupDestination")
	assert.NoError(t, err)

	type GetBackupDestinationRequestInfo struct {
		ContainerId string
		Request     database.GetBackupDestinationRequest
	}

	var requests []GetBackupDestinationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetBackupDestination(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCS" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetCloudExadataInfrastructure(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetCloudExadataInfrastructure")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCloudExadataInfrastructure is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetCloudExadataInfrastructure", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetCloudExadataInfrastructure")
	assert.NoError(t, err)

	type GetCloudExadataInfrastructureRequestInfo struct {
		ContainerId string
		Request     database.GetCloudExadataInfrastructureRequest
	}

	var requests []GetCloudExadataInfrastructureRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetCloudExadataInfrastructure(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCS" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetCloudVmCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetCloudVmCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCloudVmCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetCloudVmCluster", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetCloudVmCluster")
	assert.NoError(t, err)

	type GetCloudVmClusterRequestInfo struct {
		ContainerId string
		Request     database.GetCloudVmClusterRequest
	}

	var requests []GetCloudVmClusterRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetCloudVmCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetCloudVmClusterIormConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetCloudVmClusterIormConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCloudVmClusterIormConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetCloudVmClusterIormConfig", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetCloudVmClusterIormConfig")
	assert.NoError(t, err)

	type GetCloudVmClusterIormConfigRequestInfo struct {
		ContainerId string
		Request     database.GetCloudVmClusterIormConfigRequest
	}

	var requests []GetCloudVmClusterIormConfigRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetCloudVmClusterIormConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetCloudVmClusterUpdate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetCloudVmClusterUpdate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCloudVmClusterUpdate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetCloudVmClusterUpdate", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetCloudVmClusterUpdate")
	assert.NoError(t, err)

	type GetCloudVmClusterUpdateRequestInfo struct {
		ContainerId string
		Request     database.GetCloudVmClusterUpdateRequest
	}

	var requests []GetCloudVmClusterUpdateRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetCloudVmClusterUpdate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetCloudVmClusterUpdateHistoryEntry(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetCloudVmClusterUpdateHistoryEntry")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCloudVmClusterUpdateHistoryEntry is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetCloudVmClusterUpdateHistoryEntry", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetCloudVmClusterUpdateHistoryEntry")
	assert.NoError(t, err)

	type GetCloudVmClusterUpdateHistoryEntryRequestInfo struct {
		ContainerId string
		Request     database.GetCloudVmClusterUpdateHistoryEntryRequest
	}

	var requests []GetCloudVmClusterUpdateHistoryEntryRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetCloudVmClusterUpdateHistoryEntry(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetConsoleConnection(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetConsoleConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConsoleConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetConsoleConnection", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetConsoleConnection")
	assert.NoError(t, err)

	type GetConsoleConnectionRequestInfo struct {
		ContainerId string
		Request     database.GetConsoleConnectionRequest
	}

	var requests []GetConsoleConnectionRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetConsoleConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDataGuardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDataGuardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDataGuardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDataGuardAssociation")
	assert.NoError(t, err)

	type GetDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.GetDataGuardAssociationRequest
	}

	var requests []GetDataGuardAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDatabase")
	assert.NoError(t, err)

	type GetDatabaseRequestInfo struct {
		ContainerId string
		Request     database.GetDatabaseRequest
	}

	var requests []GetDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDatabaseSoftwareImage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDatabaseSoftwareImage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDatabaseSoftwareImage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDatabaseSoftwareImage", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDatabaseSoftwareImage")
	assert.NoError(t, err)

	type GetDatabaseSoftwareImageRequestInfo struct {
		ContainerId string
		Request     database.GetDatabaseSoftwareImageRequest
	}

	var requests []GetDatabaseSoftwareImageRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDatabaseSoftwareImage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbHome(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDbHome")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbHome is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDbHome", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDbHome")
	assert.NoError(t, err)

	type GetDbHomeRequestInfo struct {
		ContainerId string
		Request     database.GetDbHomeRequest
	}

	var requests []GetDbHomeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDbHome(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbHomePatch(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDbHomePatch")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbHomePatch is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDbHomePatch", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDbHomePatch")
	assert.NoError(t, err)

	type GetDbHomePatchRequestInfo struct {
		ContainerId string
		Request     database.GetDbHomePatchRequest
	}

	var requests []GetDbHomePatchRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDbHomePatch(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbHomePatchHistoryEntry(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDbHomePatchHistoryEntry")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbHomePatchHistoryEntry is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDbHomePatchHistoryEntry", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDbHomePatchHistoryEntry")
	assert.NoError(t, err)

	type GetDbHomePatchHistoryEntryRequestInfo struct {
		ContainerId string
		Request     database.GetDbHomePatchHistoryEntryRequest
	}

	var requests []GetDbHomePatchHistoryEntryRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDbHomePatchHistoryEntry(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbNode(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDbNode")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbNode is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDbNode", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDbNode")
	assert.NoError(t, err)

	type GetDbNodeRequestInfo struct {
		ContainerId string
		Request     database.GetDbNodeRequest
	}

	var requests []GetDbNodeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDbNode(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDbSystem", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDbSystem")
	assert.NoError(t, err)

	type GetDbSystemRequestInfo struct {
		ContainerId string
		Request     database.GetDbSystemRequest
	}

	var requests []GetDbSystemRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbSystemPatch(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDbSystemPatch")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbSystemPatch is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDbSystemPatch", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDbSystemPatch")
	assert.NoError(t, err)

	type GetDbSystemPatchRequestInfo struct {
		ContainerId string
		Request     database.GetDbSystemPatchRequest
	}

	var requests []GetDbSystemPatchRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDbSystemPatch(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetDbSystemPatchHistoryEntry(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetDbSystemPatchHistoryEntry")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDbSystemPatchHistoryEntry is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetDbSystemPatchHistoryEntry", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetDbSystemPatchHistoryEntry")
	assert.NoError(t, err)

	type GetDbSystemPatchHistoryEntryRequestInfo struct {
		ContainerId string
		Request     database.GetDbSystemPatchHistoryEntryRequest
	}

	var requests []GetDbSystemPatchHistoryEntryRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDbSystemPatchHistoryEntry(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetExadataInfrastructure(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetExadataInfrastructure")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetExadataInfrastructure is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetExadataInfrastructure", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetExadataInfrastructure")
	assert.NoError(t, err)

	type GetExadataInfrastructureRequestInfo struct {
		ContainerId string
		Request     database.GetExadataInfrastructureRequest
	}

	var requests []GetExadataInfrastructureRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetExadataInfrastructure(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetExadataInfrastructureOcpus(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetExadataInfrastructureOcpus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetExadataInfrastructureOcpus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetExadataInfrastructureOcpus", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetExadataInfrastructureOcpus")
	assert.NoError(t, err)

	type GetExadataInfrastructureOcpusRequestInfo struct {
		ContainerId string
		Request     database.GetExadataInfrastructureOcpusRequest
	}

	var requests []GetExadataInfrastructureOcpusRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetExadataInfrastructureOcpus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetExadataIormConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetExadataIormConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetExadataIormConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetExadataIormConfig", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetExadataIormConfig")
	assert.NoError(t, err)

	type GetExadataIormConfigRequestInfo struct {
		ContainerId string
		Request     database.GetExadataIormConfigRequest
	}

	var requests []GetExadataIormConfigRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetExadataIormConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetExternalBackupJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetExternalBackupJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetExternalBackupJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetExternalBackupJob", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetExternalBackupJob")
	assert.NoError(t, err)

	type GetExternalBackupJobRequestInfo struct {
		ContainerId string
		Request     database.GetExternalBackupJobRequest
	}

	var requests []GetExternalBackupJobRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetExternalBackupJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetMaintenanceRun(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetMaintenanceRun")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetMaintenanceRun is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetMaintenanceRun", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetMaintenanceRun")
	assert.NoError(t, err)

	type GetMaintenanceRunRequestInfo struct {
		ContainerId string
		Request     database.GetMaintenanceRunRequest
	}

	var requests []GetMaintenanceRunRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetMaintenanceRun(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetTargetDbHome(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetTargetDbHome")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTargetDbHome is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetTargetDbHome", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetTargetDbHome")
	assert.NoError(t, err)

	type GetTargetDbHomeRequestInfo struct {
		ContainerId string
		Request     database.GetTargetDbHomeRequest
	}

	var requests []GetTargetDbHomeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetTargetDbHome(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetVmCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetVmCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetVmCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetVmCluster", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetVmCluster")
	assert.NoError(t, err)

	type GetVmClusterRequestInfo struct {
		ContainerId string
		Request     database.GetVmClusterRequest
	}

	var requests []GetVmClusterRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetVmCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetVmClusterNetwork(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetVmClusterNetwork")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetVmClusterNetwork is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetVmClusterNetwork", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetVmClusterNetwork")
	assert.NoError(t, err)

	type GetVmClusterNetworkRequestInfo struct {
		ContainerId string
		Request     database.GetVmClusterNetworkRequest
	}

	var requests []GetVmClusterNetworkRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetVmClusterNetwork(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetVmClusterPatch(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetVmClusterPatch")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetVmClusterPatch is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetVmClusterPatch", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetVmClusterPatch")
	assert.NoError(t, err)

	type GetVmClusterPatchRequestInfo struct {
		ContainerId string
		Request     database.GetVmClusterPatchRequest
	}

	var requests []GetVmClusterPatchRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetVmClusterPatch(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientGetVmClusterPatchHistoryEntry(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "GetVmClusterPatchHistoryEntry")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetVmClusterPatchHistoryEntry is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "GetVmClusterPatchHistoryEntry", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "GetVmClusterPatchHistoryEntry")
	assert.NoError(t, err)

	type GetVmClusterPatchHistoryEntryRequestInfo struct {
		ContainerId string
		Request     database.GetVmClusterPatchHistoryEntryRequest
	}

	var requests []GetVmClusterPatchHistoryEntryRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetVmClusterPatchHistoryEntry(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientLaunchAutonomousExadataInfrastructure(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "LaunchAutonomousExadataInfrastructure")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("LaunchAutonomousExadataInfrastructure is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "LaunchAutonomousExadataInfrastructure", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "LaunchAutonomousExadataInfrastructure")
	assert.NoError(t, err)

	type LaunchAutonomousExadataInfrastructureRequestInfo struct {
		ContainerId string
		Request     database.LaunchAutonomousExadataInfrastructureRequest
	}

	var requests []LaunchAutonomousExadataInfrastructureRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.LaunchAutonomousExadataInfrastructure(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientLaunchDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "LaunchDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("LaunchDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "LaunchDbSystem", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "LaunchDbSystem")
	assert.NoError(t, err)

	type LaunchDbSystemRequestInfo struct {
		ContainerId string
		Request     database.LaunchDbSystemRequest
	}

	var requests []LaunchDbSystemRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]LaunchDbSystemRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["LaunchDbSystemBase"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "source",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"NONE":      &database.LaunchDbSystemDetails{},
				"DB_SYSTEM": &database.LaunchDbSystemFromDbSystemDetails{},
				"DATABASE":  &database.LaunchDbSystemFromDatabaseDetails{},
				"DB_BACKUP": &database.LaunchDbSystemFromBackupDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.LaunchDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousContainerDatabaseDataguardAssociations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousContainerDatabaseDataguardAssociations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousContainerDatabaseDataguardAssociations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousContainerDatabaseDataguardAssociations", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousContainerDatabaseDataguardAssociations")
	assert.NoError(t, err)

	type ListAutonomousContainerDatabaseDataguardAssociationsRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousContainerDatabaseDataguardAssociationsRequest
	}

	var requests []ListAutonomousContainerDatabaseDataguardAssociationsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousContainerDatabaseDataguardAssociationsRequest)
				return c.ListAutonomousContainerDatabaseDataguardAssociations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousContainerDatabaseDataguardAssociationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousContainerDatabaseDataguardAssociationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousContainerDatabases(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousContainerDatabases")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousContainerDatabases is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousContainerDatabases", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousContainerDatabases")
	assert.NoError(t, err)

	type ListAutonomousContainerDatabasesRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousContainerDatabasesRequest
	}

	var requests []ListAutonomousContainerDatabasesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousContainerDatabasesRequest)
				return c.ListAutonomousContainerDatabases(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousContainerDatabasesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousContainerDatabasesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDataWarehouseBackups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDataWarehouseBackups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDataWarehouseBackups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousDataWarehouseBackups", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousDataWarehouseBackups")
	assert.NoError(t, err)

	type ListAutonomousDataWarehouseBackupsRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDataWarehouseBackupsRequest
	}

	var requests []ListAutonomousDataWarehouseBackupsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDataWarehouses(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDataWarehouses")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDataWarehouses is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousDataWarehouses", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousDataWarehouses")
	assert.NoError(t, err)

	type ListAutonomousDataWarehousesRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDataWarehousesRequest
	}

	var requests []ListAutonomousDataWarehousesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDatabaseBackups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDatabaseBackups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDatabaseBackups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousDatabaseBackups", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousDatabaseBackups")
	assert.NoError(t, err)

	type ListAutonomousDatabaseBackupsRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDatabaseBackupsRequest
	}

	var requests []ListAutonomousDatabaseBackupsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDatabaseDataguardAssociations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDatabaseDataguardAssociations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDatabaseDataguardAssociations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousDatabaseDataguardAssociations", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousDatabaseDataguardAssociations")
	assert.NoError(t, err)

	type ListAutonomousDatabaseDataguardAssociationsRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDatabaseDataguardAssociationsRequest
	}

	var requests []ListAutonomousDatabaseDataguardAssociationsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDatabaseDataguardAssociationsRequest)
				return c.ListAutonomousDatabaseDataguardAssociations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDatabaseDataguardAssociationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDatabaseDataguardAssociationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDatabaseRefreshableClones(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDatabaseRefreshableClones")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDatabaseRefreshableClones is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousDatabaseRefreshableClones", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousDatabaseRefreshableClones")
	assert.NoError(t, err)

	type ListAutonomousDatabaseRefreshableClonesRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDatabaseRefreshableClonesRequest
	}

	var requests []ListAutonomousDatabaseRefreshableClonesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDatabaseRefreshableClonesRequest)
				return c.ListAutonomousDatabaseRefreshableClones(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDatabaseRefreshableClonesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDatabaseRefreshableClonesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDatabases(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDatabases")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDatabases is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousDatabases", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousDatabases")
	assert.NoError(t, err)

	type ListAutonomousDatabasesRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDatabasesRequest
	}

	var requests []ListAutonomousDatabasesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDbPreviewVersions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDbPreviewVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDbPreviewVersions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousDbPreviewVersions", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousDbPreviewVersions")
	assert.NoError(t, err)

	type ListAutonomousDbPreviewVersionsRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDbPreviewVersionsRequest
	}

	var requests []ListAutonomousDbPreviewVersionsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDbPreviewVersionsRequest)
				return c.ListAutonomousDbPreviewVersions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDbPreviewVersionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDbPreviewVersionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousDbVersions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousDbVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousDbVersions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousDbVersions", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousDbVersions")
	assert.NoError(t, err)

	type ListAutonomousDbVersionsRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousDbVersionsRequest
	}

	var requests []ListAutonomousDbVersionsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousDbVersionsRequest)
				return c.ListAutonomousDbVersions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousDbVersionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousDbVersionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousExadataInfrastructureShapes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousExadataInfrastructureShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousExadataInfrastructureShapes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousExadataInfrastructureShapes", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousExadataInfrastructureShapes")
	assert.NoError(t, err)

	type ListAutonomousExadataInfrastructureShapesRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousExadataInfrastructureShapesRequest
	}

	var requests []ListAutonomousExadataInfrastructureShapesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousExadataInfrastructureShapesRequest)
				return c.ListAutonomousExadataInfrastructureShapes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousExadataInfrastructureShapesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousExadataInfrastructureShapesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousExadataInfrastructures(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousExadataInfrastructures")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousExadataInfrastructures is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousExadataInfrastructures", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousExadataInfrastructures")
	assert.NoError(t, err)

	type ListAutonomousExadataInfrastructuresRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousExadataInfrastructuresRequest
	}

	var requests []ListAutonomousExadataInfrastructuresRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousExadataInfrastructuresRequest)
				return c.ListAutonomousExadataInfrastructures(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousExadataInfrastructuresResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousExadataInfrastructuresResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListAutonomousVmClusters(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListAutonomousVmClusters")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutonomousVmClusters is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListAutonomousVmClusters", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListAutonomousVmClusters")
	assert.NoError(t, err)

	type ListAutonomousVmClustersRequestInfo struct {
		ContainerId string
		Request     database.ListAutonomousVmClustersRequest
	}

	var requests []ListAutonomousVmClustersRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListAutonomousVmClustersRequest)
				return c.ListAutonomousVmClusters(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListAutonomousVmClustersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListAutonomousVmClustersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListBackupDestination(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListBackupDestination")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBackupDestination is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListBackupDestination", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListBackupDestination")
	assert.NoError(t, err)

	type ListBackupDestinationRequestInfo struct {
		ContainerId string
		Request     database.ListBackupDestinationRequest
	}

	var requests []ListBackupDestinationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListBackupDestinationRequest)
				return c.ListBackupDestination(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListBackupDestinationResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListBackupDestinationResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListBackups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListBackups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBackups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListBackups", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListBackups")
	assert.NoError(t, err)

	type ListBackupsRequestInfo struct {
		ContainerId string
		Request     database.ListBackupsRequest
	}

	var requests []ListBackupsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="ExaCS" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListCloudExadataInfrastructures(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListCloudExadataInfrastructures")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCloudExadataInfrastructures is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListCloudExadataInfrastructures", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListCloudExadataInfrastructures")
	assert.NoError(t, err)

	type ListCloudExadataInfrastructuresRequestInfo struct {
		ContainerId string
		Request     database.ListCloudExadataInfrastructuresRequest
	}

	var requests []ListCloudExadataInfrastructuresRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListCloudExadataInfrastructuresRequest)
				return c.ListCloudExadataInfrastructures(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListCloudExadataInfrastructuresResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListCloudExadataInfrastructuresResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListCloudVmClusterUpdateHistoryEntries(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListCloudVmClusterUpdateHistoryEntries")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCloudVmClusterUpdateHistoryEntries is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListCloudVmClusterUpdateHistoryEntries", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListCloudVmClusterUpdateHistoryEntries")
	assert.NoError(t, err)

	type ListCloudVmClusterUpdateHistoryEntriesRequestInfo struct {
		ContainerId string
		Request     database.ListCloudVmClusterUpdateHistoryEntriesRequest
	}

	var requests []ListCloudVmClusterUpdateHistoryEntriesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListCloudVmClusterUpdateHistoryEntriesRequest)
				return c.ListCloudVmClusterUpdateHistoryEntries(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListCloudVmClusterUpdateHistoryEntriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListCloudVmClusterUpdateHistoryEntriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListCloudVmClusterUpdates(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListCloudVmClusterUpdates")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCloudVmClusterUpdates is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListCloudVmClusterUpdates", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListCloudVmClusterUpdates")
	assert.NoError(t, err)

	type ListCloudVmClusterUpdatesRequestInfo struct {
		ContainerId string
		Request     database.ListCloudVmClusterUpdatesRequest
	}

	var requests []ListCloudVmClusterUpdatesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListCloudVmClusterUpdatesRequest)
				return c.ListCloudVmClusterUpdates(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListCloudVmClusterUpdatesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListCloudVmClusterUpdatesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCS" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListCloudVmClusters(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListCloudVmClusters")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCloudVmClusters is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListCloudVmClusters", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListCloudVmClusters")
	assert.NoError(t, err)

	type ListCloudVmClustersRequestInfo struct {
		ContainerId string
		Request     database.ListCloudVmClustersRequest
	}

	var requests []ListCloudVmClustersRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListCloudVmClustersRequest)
				return c.ListCloudVmClusters(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListCloudVmClustersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListCloudVmClustersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListConsoleConnections(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListConsoleConnections")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListConsoleConnections is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListConsoleConnections", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListConsoleConnections")
	assert.NoError(t, err)

	type ListConsoleConnectionsRequestInfo struct {
		ContainerId string
		Request     database.ListConsoleConnectionsRequest
	}

	var requests []ListConsoleConnectionsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ListConsoleConnections(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListContainerDatabasePatches(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListContainerDatabasePatches")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListContainerDatabasePatches is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListContainerDatabasePatches", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListContainerDatabasePatches")
	assert.NoError(t, err)

	type ListContainerDatabasePatchesRequestInfo struct {
		ContainerId string
		Request     database.ListContainerDatabasePatchesRequest
	}

	var requests []ListContainerDatabasePatchesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListContainerDatabasePatchesRequest)
				return c.ListContainerDatabasePatches(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListContainerDatabasePatchesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListContainerDatabasePatchesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDataGuardAssociations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDataGuardAssociations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDataGuardAssociations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDataGuardAssociations", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDataGuardAssociations")
	assert.NoError(t, err)

	type ListDataGuardAssociationsRequestInfo struct {
		ContainerId string
		Request     database.ListDataGuardAssociationsRequest
	}

	var requests []ListDataGuardAssociationsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDatabaseSoftwareImages(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDatabaseSoftwareImages")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDatabaseSoftwareImages is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDatabaseSoftwareImages", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDatabaseSoftwareImages")
	assert.NoError(t, err)

	type ListDatabaseSoftwareImagesRequestInfo struct {
		ContainerId string
		Request     database.ListDatabaseSoftwareImagesRequest
	}

	var requests []ListDatabaseSoftwareImagesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListDatabaseSoftwareImagesRequest)
				return c.ListDatabaseSoftwareImages(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListDatabaseSoftwareImagesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListDatabaseSoftwareImagesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDatabases(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDatabases")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDatabases is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDatabases", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDatabases")
	assert.NoError(t, err)

	type ListDatabasesRequestInfo struct {
		ContainerId string
		Request     database.ListDatabasesRequest
	}

	var requests []ListDatabasesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbHomePatchHistoryEntries(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbHomePatchHistoryEntries")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbHomePatchHistoryEntries is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbHomePatchHistoryEntries", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbHomePatchHistoryEntries")
	assert.NoError(t, err)

	type ListDbHomePatchHistoryEntriesRequestInfo struct {
		ContainerId string
		Request     database.ListDbHomePatchHistoryEntriesRequest
	}

	var requests []ListDbHomePatchHistoryEntriesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbHomePatches(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbHomePatches")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbHomePatches is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbHomePatches", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbHomePatches")
	assert.NoError(t, err)

	type ListDbHomePatchesRequestInfo struct {
		ContainerId string
		Request     database.ListDbHomePatchesRequest
	}

	var requests []ListDbHomePatchesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbHomes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbHomes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbHomes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbHomes", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbHomes")
	assert.NoError(t, err)

	type ListDbHomesRequestInfo struct {
		ContainerId string
		Request     database.ListDbHomesRequest
	}

	var requests []ListDbHomesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbNodes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbNodes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbNodes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbNodes", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbNodes")
	assert.NoError(t, err)

	type ListDbNodesRequestInfo struct {
		ContainerId string
		Request     database.ListDbNodesRequest
	}

	var requests []ListDbNodesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbSystemPatchHistoryEntries(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbSystemPatchHistoryEntries")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbSystemPatchHistoryEntries is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbSystemPatchHistoryEntries", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbSystemPatchHistoryEntries")
	assert.NoError(t, err)

	type ListDbSystemPatchHistoryEntriesRequestInfo struct {
		ContainerId string
		Request     database.ListDbSystemPatchHistoryEntriesRequest
	}

	var requests []ListDbSystemPatchHistoryEntriesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbSystemPatches(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbSystemPatches")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbSystemPatches is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbSystemPatches", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbSystemPatches")
	assert.NoError(t, err)

	type ListDbSystemPatchesRequestInfo struct {
		ContainerId string
		Request     database.ListDbSystemPatchesRequest
	}

	var requests []ListDbSystemPatchesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbSystemShapes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbSystemShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbSystemShapes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbSystemShapes", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbSystemShapes")
	assert.NoError(t, err)

	type ListDbSystemShapesRequestInfo struct {
		ContainerId string
		Request     database.ListDbSystemShapesRequest
	}

	var requests []ListDbSystemShapesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbSystems(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbSystems")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbSystems is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbSystems", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbSystems")
	assert.NoError(t, err)

	type ListDbSystemsRequestInfo struct {
		ContainerId string
		Request     database.ListDbSystemsRequest
	}

	var requests []ListDbSystemsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListDbVersions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListDbVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDbVersions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListDbVersions", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListDbVersions")
	assert.NoError(t, err)

	type ListDbVersionsRequestInfo struct {
		ContainerId string
		Request     database.ListDbVersionsRequest
	}

	var requests []ListDbVersionsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListExadataInfrastructures(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListExadataInfrastructures")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListExadataInfrastructures is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListExadataInfrastructures", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListExadataInfrastructures")
	assert.NoError(t, err)

	type ListExadataInfrastructuresRequestInfo struct {
		ContainerId string
		Request     database.ListExadataInfrastructuresRequest
	}

	var requests []ListExadataInfrastructuresRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListExadataInfrastructuresRequest)
				return c.ListExadataInfrastructures(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListExadataInfrastructuresResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListExadataInfrastructuresResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListGiVersions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListGiVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListGiVersions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListGiVersions", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListGiVersions")
	assert.NoError(t, err)

	type ListGiVersionsRequestInfo struct {
		ContainerId string
		Request     database.ListGiVersionsRequest
	}

	var requests []ListGiVersionsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListGiVersionsRequest)
				return c.ListGiVersions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListGiVersionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListGiVersionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListMaintenanceRuns(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListMaintenanceRuns")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListMaintenanceRuns is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListMaintenanceRuns", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListMaintenanceRuns")
	assert.NoError(t, err)

	type ListMaintenanceRunsRequestInfo struct {
		ContainerId string
		Request     database.ListMaintenanceRunsRequest
	}

	var requests []ListMaintenanceRunsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListVmClusterNetworks(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListVmClusterNetworks")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListVmClusterNetworks is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListVmClusterNetworks", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListVmClusterNetworks")
	assert.NoError(t, err)

	type ListVmClusterNetworksRequestInfo struct {
		ContainerId string
		Request     database.ListVmClusterNetworksRequest
	}

	var requests []ListVmClusterNetworksRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListVmClusterNetworksRequest)
				return c.ListVmClusterNetworks(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListVmClusterNetworksResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListVmClusterNetworksResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListVmClusterPatchHistoryEntries(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListVmClusterPatchHistoryEntries")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListVmClusterPatchHistoryEntries is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListVmClusterPatchHistoryEntries", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListVmClusterPatchHistoryEntries")
	assert.NoError(t, err)

	type ListVmClusterPatchHistoryEntriesRequestInfo struct {
		ContainerId string
		Request     database.ListVmClusterPatchHistoryEntriesRequest
	}

	var requests []ListVmClusterPatchHistoryEntriesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListVmClusterPatchHistoryEntriesRequest)
				return c.ListVmClusterPatchHistoryEntries(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListVmClusterPatchHistoryEntriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListVmClusterPatchHistoryEntriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListVmClusterPatches(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListVmClusterPatches")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListVmClusterPatches is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListVmClusterPatches", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListVmClusterPatches")
	assert.NoError(t, err)

	type ListVmClusterPatchesRequestInfo struct {
		ContainerId string
		Request     database.ListVmClusterPatchesRequest
	}

	var requests []ListVmClusterPatchesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListVmClusterPatchesRequest)
				return c.ListVmClusterPatches(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListVmClusterPatchesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListVmClusterPatchesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientListVmClusters(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ListVmClusters")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListVmClusters is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ListVmClusters", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ListVmClusters")
	assert.NoError(t, err)

	type ListVmClustersRequestInfo struct {
		ContainerId string
		Request     database.ListVmClustersRequest
	}

	var requests []ListVmClustersRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*database.ListVmClustersRequest)
				return c.ListVmClusters(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]database.ListVmClustersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(database.ListVmClustersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientMigrateExadataDbSystemResourceModel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "MigrateExadataDbSystemResourceModel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("MigrateExadataDbSystemResourceModel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "MigrateExadataDbSystemResourceModel", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "MigrateExadataDbSystemResourceModel")
	assert.NoError(t, err)

	type MigrateExadataDbSystemResourceModelRequestInfo struct {
		ContainerId string
		Request     database.MigrateExadataDbSystemResourceModelRequest
	}

	var requests []MigrateExadataDbSystemResourceModelRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.MigrateExadataDbSystemResourceModel(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientMigrateVaultKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "MigrateVaultKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("MigrateVaultKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "MigrateVaultKey", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "MigrateVaultKey")
	assert.NoError(t, err)

	type MigrateVaultKeyRequestInfo struct {
		ContainerId string
		Request     database.MigrateVaultKeyRequest
	}

	var requests []MigrateVaultKeyRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.MigrateVaultKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientMoveDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "MoveDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("MoveDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "MoveDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "MoveDatabase")
	assert.NoError(t, err)

	type MoveDatabaseRequestInfo struct {
		ContainerId string
		Request     database.MoveDatabaseRequest
	}

	var requests []MoveDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.MoveDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientRegisterAutonomousDatabaseDataSafe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "RegisterAutonomousDatabaseDataSafe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RegisterAutonomousDatabaseDataSafe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "RegisterAutonomousDatabaseDataSafe", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "RegisterAutonomousDatabaseDataSafe")
	assert.NoError(t, err)

	type RegisterAutonomousDatabaseDataSafeRequestInfo struct {
		ContainerId string
		Request     database.RegisterAutonomousDatabaseDataSafeRequest
	}

	var requests []RegisterAutonomousDatabaseDataSafeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.RegisterAutonomousDatabaseDataSafe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientReinstateAutonomousContainerDatabaseDataguardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ReinstateAutonomousContainerDatabaseDataguardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ReinstateAutonomousContainerDatabaseDataguardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ReinstateAutonomousContainerDatabaseDataguardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ReinstateAutonomousContainerDatabaseDataguardAssociation")
	assert.NoError(t, err)

	type ReinstateAutonomousContainerDatabaseDataguardAssociationRequestInfo struct {
		ContainerId string
		Request     database.ReinstateAutonomousContainerDatabaseDataguardAssociationRequest
	}

	var requests []ReinstateAutonomousContainerDatabaseDataguardAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ReinstateAutonomousContainerDatabaseDataguardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientReinstateDataGuardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ReinstateDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ReinstateDataGuardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ReinstateDataGuardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ReinstateDataGuardAssociation")
	assert.NoError(t, err)

	type ReinstateDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.ReinstateDataGuardAssociationRequest
	}

	var requests []ReinstateDataGuardAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ReinstateDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientRestartAutonomousContainerDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "RestartAutonomousContainerDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestartAutonomousContainerDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "RestartAutonomousContainerDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "RestartAutonomousContainerDatabase")
	assert.NoError(t, err)

	type RestartAutonomousContainerDatabaseRequestInfo struct {
		ContainerId string
		Request     database.RestartAutonomousContainerDatabaseRequest
	}

	var requests []RestartAutonomousContainerDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.RestartAutonomousContainerDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientRestartAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "RestartAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestartAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "RestartAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "RestartAutonomousDatabase")
	assert.NoError(t, err)

	type RestartAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.RestartAutonomousDatabaseRequest
	}

	var requests []RestartAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.RestartAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientRestoreAutonomousDataWarehouse(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "RestoreAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestoreAutonomousDataWarehouse is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "RestoreAutonomousDataWarehouse", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "RestoreAutonomousDataWarehouse")
	assert.NoError(t, err)

	type RestoreAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.RestoreAutonomousDataWarehouseRequest
	}

	var requests []RestoreAutonomousDataWarehouseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.RestoreAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientRestoreAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "RestoreAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestoreAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "RestoreAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "RestoreAutonomousDatabase")
	assert.NoError(t, err)

	type RestoreAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.RestoreAutonomousDatabaseRequest
	}

	var requests []RestoreAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.RestoreAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientRestoreDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "RestoreDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestoreDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "RestoreDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "RestoreDatabase")
	assert.NoError(t, err)

	type RestoreDatabaseRequestInfo struct {
		ContainerId string
		Request     database.RestoreDatabaseRequest
	}

	var requests []RestoreDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.RestoreDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientRotateAutonomousContainerDatabaseVaultKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "RotateAutonomousContainerDatabaseVaultKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RotateAutonomousContainerDatabaseVaultKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "RotateAutonomousContainerDatabaseVaultKey", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "RotateAutonomousContainerDatabaseVaultKey")
	assert.NoError(t, err)

	type RotateAutonomousContainerDatabaseVaultKeyRequestInfo struct {
		ContainerId string
		Request     database.RotateAutonomousContainerDatabaseVaultKeyRequest
	}

	var requests []RotateAutonomousContainerDatabaseVaultKeyRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.RotateAutonomousContainerDatabaseVaultKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientRotateVaultKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "RotateVaultKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RotateVaultKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "RotateVaultKey", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "RotateVaultKey")
	assert.NoError(t, err)

	type RotateVaultKeyRequestInfo struct {
		ContainerId string
		Request     database.RotateVaultKeyRequest
	}

	var requests []RotateVaultKeyRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.RotateVaultKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientStartAutonomousDataWarehouse(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "StartAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StartAutonomousDataWarehouse is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "StartAutonomousDataWarehouse", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "StartAutonomousDataWarehouse")
	assert.NoError(t, err)

	type StartAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.StartAutonomousDataWarehouseRequest
	}

	var requests []StartAutonomousDataWarehouseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.StartAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientStartAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "StartAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StartAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "StartAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "StartAutonomousDatabase")
	assert.NoError(t, err)

	type StartAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.StartAutonomousDatabaseRequest
	}

	var requests []StartAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.StartAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientStopAutonomousDataWarehouse(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "StopAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StopAutonomousDataWarehouse is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "StopAutonomousDataWarehouse", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "StopAutonomousDataWarehouse")
	assert.NoError(t, err)

	type StopAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.StopAutonomousDataWarehouseRequest
	}

	var requests []StopAutonomousDataWarehouseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.StopAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientStopAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "StopAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StopAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "StopAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "StopAutonomousDatabase")
	assert.NoError(t, err)

	type StopAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.StopAutonomousDatabaseRequest
	}

	var requests []StopAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.StopAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientSwitchoverAutonomousContainerDatabaseDataguardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "SwitchoverAutonomousContainerDatabaseDataguardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SwitchoverAutonomousContainerDatabaseDataguardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "SwitchoverAutonomousContainerDatabaseDataguardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "SwitchoverAutonomousContainerDatabaseDataguardAssociation")
	assert.NoError(t, err)

	type SwitchoverAutonomousContainerDatabaseDataguardAssociationRequestInfo struct {
		ContainerId string
		Request     database.SwitchoverAutonomousContainerDatabaseDataguardAssociationRequest
	}

	var requests []SwitchoverAutonomousContainerDatabaseDataguardAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.SwitchoverAutonomousContainerDatabaseDataguardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientSwitchoverAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "SwitchoverAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SwitchoverAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "SwitchoverAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "SwitchoverAutonomousDatabase")
	assert.NoError(t, err)

	type SwitchoverAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.SwitchoverAutonomousDatabaseRequest
	}

	var requests []SwitchoverAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.SwitchoverAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientSwitchoverDataGuardAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "SwitchoverDataGuardAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SwitchoverDataGuardAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "SwitchoverDataGuardAssociation", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "SwitchoverDataGuardAssociation")
	assert.NoError(t, err)

	type SwitchoverDataGuardAssociationRequestInfo struct {
		ContainerId string
		Request     database.SwitchoverDataGuardAssociationRequest
	}

	var requests []SwitchoverDataGuardAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.SwitchoverDataGuardAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientTerminateAutonomousContainerDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "TerminateAutonomousContainerDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("TerminateAutonomousContainerDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "TerminateAutonomousContainerDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "TerminateAutonomousContainerDatabase")
	assert.NoError(t, err)

	type TerminateAutonomousContainerDatabaseRequestInfo struct {
		ContainerId string
		Request     database.TerminateAutonomousContainerDatabaseRequest
	}

	var requests []TerminateAutonomousContainerDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.TerminateAutonomousContainerDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientTerminateAutonomousExadataInfrastructure(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "TerminateAutonomousExadataInfrastructure")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("TerminateAutonomousExadataInfrastructure is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "TerminateAutonomousExadataInfrastructure", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "TerminateAutonomousExadataInfrastructure")
	assert.NoError(t, err)

	type TerminateAutonomousExadataInfrastructureRequestInfo struct {
		ContainerId string
		Request     database.TerminateAutonomousExadataInfrastructureRequest
	}

	var requests []TerminateAutonomousExadataInfrastructureRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.TerminateAutonomousExadataInfrastructure(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientTerminateDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "TerminateDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("TerminateDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "TerminateDbSystem", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "TerminateDbSystem")
	assert.NoError(t, err)

	type TerminateDbSystemRequestInfo struct {
		ContainerId string
		Request     database.TerminateDbSystemRequest
	}

	var requests []TerminateDbSystemRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.TerminateDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateAutonomousContainerDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateAutonomousContainerDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutonomousContainerDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateAutonomousContainerDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateAutonomousContainerDatabase")
	assert.NoError(t, err)

	type UpdateAutonomousContainerDatabaseRequestInfo struct {
		ContainerId string
		Request     database.UpdateAutonomousContainerDatabaseRequest
	}

	var requests []UpdateAutonomousContainerDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateAutonomousContainerDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateAutonomousDataWarehouse(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateAutonomousDataWarehouse")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutonomousDataWarehouse is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateAutonomousDataWarehouse", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateAutonomousDataWarehouse")
	assert.NoError(t, err)

	type UpdateAutonomousDataWarehouseRequestInfo struct {
		ContainerId string
		Request     database.UpdateAutonomousDataWarehouseRequest
	}

	var requests []UpdateAutonomousDataWarehouseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateAutonomousDataWarehouse(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateAutonomousDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateAutonomousDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutonomousDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateAutonomousDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateAutonomousDatabase")
	assert.NoError(t, err)

	type UpdateAutonomousDatabaseRequestInfo struct {
		ContainerId string
		Request     database.UpdateAutonomousDatabaseRequest
	}

	var requests []UpdateAutonomousDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateAutonomousDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateAutonomousDatabaseRegionalWallet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateAutonomousDatabaseRegionalWallet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutonomousDatabaseRegionalWallet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateAutonomousDatabaseRegionalWallet", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateAutonomousDatabaseRegionalWallet")
	assert.NoError(t, err)

	type UpdateAutonomousDatabaseRegionalWalletRequestInfo struct {
		ContainerId string
		Request     database.UpdateAutonomousDatabaseRegionalWalletRequest
	}

	var requests []UpdateAutonomousDatabaseRegionalWalletRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateAutonomousDatabaseRegionalWallet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-adb" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateAutonomousDatabaseWallet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateAutonomousDatabaseWallet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutonomousDatabaseWallet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateAutonomousDatabaseWallet", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateAutonomousDatabaseWallet")
	assert.NoError(t, err)

	type UpdateAutonomousDatabaseWalletRequestInfo struct {
		ContainerId string
		Request     database.UpdateAutonomousDatabaseWalletRequest
	}

	var requests []UpdateAutonomousDatabaseWalletRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateAutonomousDatabaseWallet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateAutonomousExadataInfrastructure(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateAutonomousExadataInfrastructure")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutonomousExadataInfrastructure is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateAutonomousExadataInfrastructure", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateAutonomousExadataInfrastructure")
	assert.NoError(t, err)

	type UpdateAutonomousExadataInfrastructureRequestInfo struct {
		ContainerId string
		Request     database.UpdateAutonomousExadataInfrastructureRequest
	}

	var requests []UpdateAutonomousExadataInfrastructureRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateAutonomousExadataInfrastructure(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateAutonomousVmCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateAutonomousVmCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutonomousVmCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateAutonomousVmCluster", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateAutonomousVmCluster")
	assert.NoError(t, err)

	type UpdateAutonomousVmClusterRequestInfo struct {
		ContainerId string
		Request     database.UpdateAutonomousVmClusterRequest
	}

	var requests []UpdateAutonomousVmClusterRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateAutonomousVmCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateBackupDestination(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateBackupDestination")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBackupDestination is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateBackupDestination", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateBackupDestination")
	assert.NoError(t, err)

	type UpdateBackupDestinationRequestInfo struct {
		ContainerId string
		Request     database.UpdateBackupDestinationRequest
	}

	var requests []UpdateBackupDestinationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateBackupDestination(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCS" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateCloudExadataInfrastructure(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateCloudExadataInfrastructure")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCloudExadataInfrastructure is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateCloudExadataInfrastructure", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateCloudExadataInfrastructure")
	assert.NoError(t, err)

	type UpdateCloudExadataInfrastructureRequestInfo struct {
		ContainerId string
		Request     database.UpdateCloudExadataInfrastructureRequest
	}

	var requests []UpdateCloudExadataInfrastructureRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateCloudExadataInfrastructure(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCS" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateCloudVmCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateCloudVmCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCloudVmCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateCloudVmCluster", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateCloudVmCluster")
	assert.NoError(t, err)

	type UpdateCloudVmClusterRequestInfo struct {
		ContainerId string
		Request     database.UpdateCloudVmClusterRequest
	}

	var requests []UpdateCloudVmClusterRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateCloudVmCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateCloudVmClusterIormConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateCloudVmClusterIormConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCloudVmClusterIormConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateCloudVmClusterIormConfig", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateCloudVmClusterIormConfig")
	assert.NoError(t, err)

	type UpdateCloudVmClusterIormConfigRequestInfo struct {
		ContainerId string
		Request     database.UpdateCloudVmClusterIormConfigRequest
	}

	var requests []UpdateCloudVmClusterIormConfigRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateCloudVmClusterIormConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateDatabase", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateDatabase")
	assert.NoError(t, err)

	type UpdateDatabaseRequestInfo struct {
		ContainerId string
		Request     database.UpdateDatabaseRequest
	}

	var requests []UpdateDatabaseRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateDatabaseSoftwareImage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateDatabaseSoftwareImage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDatabaseSoftwareImage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateDatabaseSoftwareImage", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateDatabaseSoftwareImage")
	assert.NoError(t, err)

	type UpdateDatabaseSoftwareImageRequestInfo struct {
		ContainerId string
		Request     database.UpdateDatabaseSoftwareImageRequest
	}

	var requests []UpdateDatabaseSoftwareImageRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateDatabaseSoftwareImage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateDbHome(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateDbHome")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDbHome is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateDbHome", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateDbHome")
	assert.NoError(t, err)

	type UpdateDbHomeRequestInfo struct {
		ContainerId string
		Request     database.UpdateDbHomeRequest
	}

	var requests []UpdateDbHomeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateDbHome(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateDbSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateDbSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDbSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateDbSystem", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateDbSystem")
	assert.NoError(t, err)

	type UpdateDbSystemRequestInfo struct {
		ContainerId string
		Request     database.UpdateDbSystemRequest
	}

	var requests []UpdateDbSystemRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateDbSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateExadataInfrastructure(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateExadataInfrastructure")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateExadataInfrastructure is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateExadataInfrastructure", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateExadataInfrastructure")
	assert.NoError(t, err)

	type UpdateExadataInfrastructureRequestInfo struct {
		ContainerId string
		Request     database.UpdateExadataInfrastructureRequest
	}

	var requests []UpdateExadataInfrastructureRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateExadataInfrastructure(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateExadataIormConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateExadataIormConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateExadataIormConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateExadataIormConfig", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateExadataIormConfig")
	assert.NoError(t, err)

	type UpdateExadataIormConfigRequestInfo struct {
		ContainerId string
		Request     database.UpdateExadataIormConfigRequest
	}

	var requests []UpdateExadataIormConfigRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateExadataIormConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="dbaas-atp-d" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateMaintenanceRun(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateMaintenanceRun")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateMaintenanceRun is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateMaintenanceRun", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateMaintenanceRun")
	assert.NoError(t, err)

	type UpdateMaintenanceRunRequestInfo struct {
		ContainerId string
		Request     database.UpdateMaintenanceRunRequest
	}

	var requests []UpdateMaintenanceRunRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateMaintenanceRun(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateVmCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateVmCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateVmCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateVmCluster", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateVmCluster")
	assert.NoError(t, err)

	type UpdateVmClusterRequestInfo struct {
		ContainerId string
		Request     database.UpdateVmClusterRequest
	}

	var requests []UpdateVmClusterRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateVmCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientUpdateVmClusterNetwork(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "UpdateVmClusterNetwork")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateVmClusterNetwork is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "UpdateVmClusterNetwork", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "UpdateVmClusterNetwork")
	assert.NoError(t, err)

	type UpdateVmClusterNetworkRequestInfo struct {
		ContainerId string
		Request     database.UpdateVmClusterNetworkRequest
	}

	var requests []UpdateVmClusterNetworkRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateVmClusterNetwork(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ExaCC" email="sic_dbaas_cp_us_grp@oracle.com" jiraProject="DBAAS" opsJiraProject="DBAASOPS"
func TestDatabaseClientValidateVmClusterNetwork(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("database", "ValidateVmClusterNetwork")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ValidateVmClusterNetwork is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("database", "Database", "ValidateVmClusterNetwork", createDatabaseClientWithProvider)
	assert.NoError(t, err)
	c := cc.(database.DatabaseClient)

	body, err := testClient.getRequests("database", "ValidateVmClusterNetwork")
	assert.NoError(t, err)

	type ValidateVmClusterNetworkRequestInfo struct {
		ContainerId string
		Request     database.ValidateVmClusterNetworkRequest
	}

	var requests []ValidateVmClusterNetworkRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ValidateVmClusterNetwork(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
