package autotest

import (
	"github.com/oracle/oci-go-sdk/v27/common"
	"github.com/oracle/oci-go-sdk/v27/goldengate"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createGoldengateGoldenGateClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := goldengate.NewGoldenGateClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientChangeBackupCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "ChangeBackupCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeBackupCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "ChangeBackupCompartment", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "ChangeBackupCompartment")
	assert.NoError(t, err)

	type ChangeBackupCompartmentRequestInfo struct {
		ContainerId string
		Request     goldengate.ChangeBackupCompartmentRequest
	}

	var requests []ChangeBackupCompartmentRequestInfo
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
			response, err := c.ChangeBackupCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientChangeDatabaseRegistrationCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "ChangeDatabaseRegistrationCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeDatabaseRegistrationCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "ChangeDatabaseRegistrationCompartment", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "ChangeDatabaseRegistrationCompartment")
	assert.NoError(t, err)

	type ChangeDatabaseRegistrationCompartmentRequestInfo struct {
		ContainerId string
		Request     goldengate.ChangeDatabaseRegistrationCompartmentRequest
	}

	var requests []ChangeDatabaseRegistrationCompartmentRequestInfo
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
			response, err := c.ChangeDatabaseRegistrationCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientChangeDeploymentCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "ChangeDeploymentCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeDeploymentCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "ChangeDeploymentCompartment", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "ChangeDeploymentCompartment")
	assert.NoError(t, err)

	type ChangeDeploymentCompartmentRequestInfo struct {
		ContainerId string
		Request     goldengate.ChangeDeploymentCompartmentRequest
	}

	var requests []ChangeDeploymentCompartmentRequestInfo
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
			response, err := c.ChangeDeploymentCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientCreateBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "CreateBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "CreateBackup", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "CreateBackup")
	assert.NoError(t, err)

	type CreateBackupRequestInfo struct {
		ContainerId string
		Request     goldengate.CreateBackupRequest
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

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientCreateDatabaseRegistration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "CreateDatabaseRegistration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDatabaseRegistration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "CreateDatabaseRegistration", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "CreateDatabaseRegistration")
	assert.NoError(t, err)

	type CreateDatabaseRegistrationRequestInfo struct {
		ContainerId string
		Request     goldengate.CreateDatabaseRegistrationRequest
	}

	var requests []CreateDatabaseRegistrationRequestInfo
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
			response, err := c.CreateDatabaseRegistration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientCreateDeployment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "CreateDeployment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDeployment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "CreateDeployment", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "CreateDeployment")
	assert.NoError(t, err)

	type CreateDeploymentRequestInfo struct {
		ContainerId string
		Request     goldengate.CreateDeploymentRequest
	}

	var requests []CreateDeploymentRequestInfo
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
			response, err := c.CreateDeployment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientDeleteBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "DeleteBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "DeleteBackup", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "DeleteBackup")
	assert.NoError(t, err)

	type DeleteBackupRequestInfo struct {
		ContainerId string
		Request     goldengate.DeleteBackupRequest
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

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientDeleteDatabaseRegistration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "DeleteDatabaseRegistration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDatabaseRegistration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "DeleteDatabaseRegistration", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "DeleteDatabaseRegistration")
	assert.NoError(t, err)

	type DeleteDatabaseRegistrationRequestInfo struct {
		ContainerId string
		Request     goldengate.DeleteDatabaseRegistrationRequest
	}

	var requests []DeleteDatabaseRegistrationRequestInfo
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
			response, err := c.DeleteDatabaseRegistration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientDeleteDeployment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "DeleteDeployment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDeployment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "DeleteDeployment", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "DeleteDeployment")
	assert.NoError(t, err)

	type DeleteDeploymentRequestInfo struct {
		ContainerId string
		Request     goldengate.DeleteDeploymentRequest
	}

	var requests []DeleteDeploymentRequestInfo
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
			response, err := c.DeleteDeployment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientGetBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "GetBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "GetBackup", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "GetBackup")
	assert.NoError(t, err)

	type GetBackupRequestInfo struct {
		ContainerId string
		Request     goldengate.GetBackupRequest
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

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientGetDatabaseRegistration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "GetDatabaseRegistration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDatabaseRegistration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "GetDatabaseRegistration", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "GetDatabaseRegistration")
	assert.NoError(t, err)

	type GetDatabaseRegistrationRequestInfo struct {
		ContainerId string
		Request     goldengate.GetDatabaseRegistrationRequest
	}

	var requests []GetDatabaseRegistrationRequestInfo
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
			response, err := c.GetDatabaseRegistration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientGetDeployment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "GetDeployment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDeployment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "GetDeployment", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "GetDeployment")
	assert.NoError(t, err)

	type GetDeploymentRequestInfo struct {
		ContainerId string
		Request     goldengate.GetDeploymentRequest
	}

	var requests []GetDeploymentRequestInfo
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
			response, err := c.GetDeployment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "GetWorkRequest", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     goldengate.GetWorkRequestRequest
	}

	var requests []GetWorkRequestRequestInfo
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
			response, err := c.GetWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientListBackups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "ListBackups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBackups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "ListBackups", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "ListBackups")
	assert.NoError(t, err)

	type ListBackupsRequestInfo struct {
		ContainerId string
		Request     goldengate.ListBackupsRequest
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
				r := req.(*goldengate.ListBackupsRequest)
				return c.ListBackups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]goldengate.ListBackupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(goldengate.ListBackupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientListDatabaseRegistrations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "ListDatabaseRegistrations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDatabaseRegistrations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "ListDatabaseRegistrations", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "ListDatabaseRegistrations")
	assert.NoError(t, err)

	type ListDatabaseRegistrationsRequestInfo struct {
		ContainerId string
		Request     goldengate.ListDatabaseRegistrationsRequest
	}

	var requests []ListDatabaseRegistrationsRequestInfo
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
				r := req.(*goldengate.ListDatabaseRegistrationsRequest)
				return c.ListDatabaseRegistrations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]goldengate.ListDatabaseRegistrationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(goldengate.ListDatabaseRegistrationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientListDeployments(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "ListDeployments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDeployments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "ListDeployments", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "ListDeployments")
	assert.NoError(t, err)

	type ListDeploymentsRequestInfo struct {
		ContainerId string
		Request     goldengate.ListDeploymentsRequest
	}

	var requests []ListDeploymentsRequestInfo
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
				r := req.(*goldengate.ListDeploymentsRequest)
				return c.ListDeployments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]goldengate.ListDeploymentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(goldengate.ListDeploymentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "ListWorkRequestErrors", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     goldengate.ListWorkRequestErrorsRequest
	}

	var requests []ListWorkRequestErrorsRequestInfo
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
				r := req.(*goldengate.ListWorkRequestErrorsRequest)
				return c.ListWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]goldengate.ListWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(goldengate.ListWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "ListWorkRequestLogs", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     goldengate.ListWorkRequestLogsRequest
	}

	var requests []ListWorkRequestLogsRequestInfo
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
				r := req.(*goldengate.ListWorkRequestLogsRequest)
				return c.ListWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]goldengate.ListWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(goldengate.ListWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "ListWorkRequests", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     goldengate.ListWorkRequestsRequest
	}

	var requests []ListWorkRequestsRequestInfo
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
				r := req.(*goldengate.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]goldengate.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(goldengate.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientStartDeployment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "StartDeployment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StartDeployment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "StartDeployment", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "StartDeployment")
	assert.NoError(t, err)

	type StartDeploymentRequestInfo struct {
		ContainerId string
		Request     goldengate.StartDeploymentRequest
	}

	var requests []StartDeploymentRequestInfo
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
			response, err := c.StartDeployment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientStopDeployment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "StopDeployment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StopDeployment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "StopDeployment", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "StopDeployment")
	assert.NoError(t, err)

	type StopDeploymentRequestInfo struct {
		ContainerId string
		Request     goldengate.StopDeploymentRequest
	}

	var requests []StopDeploymentRequestInfo
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
			response, err := c.StopDeployment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientSynchronizeDatabaseRegistration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "SynchronizeDatabaseRegistration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SynchronizeDatabaseRegistration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "SynchronizeDatabaseRegistration", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "SynchronizeDatabaseRegistration")
	assert.NoError(t, err)

	type SynchronizeDatabaseRegistrationRequestInfo struct {
		ContainerId string
		Request     goldengate.SynchronizeDatabaseRegistrationRequest
	}

	var requests []SynchronizeDatabaseRegistrationRequestInfo
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
			response, err := c.SynchronizeDatabaseRegistration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientUpdateBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "UpdateBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "UpdateBackup", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "UpdateBackup")
	assert.NoError(t, err)

	type UpdateBackupRequestInfo struct {
		ContainerId string
		Request     goldengate.UpdateBackupRequest
	}

	var requests []UpdateBackupRequestInfo
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
			response, err := c.UpdateBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientUpdateDatabaseRegistration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "UpdateDatabaseRegistration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDatabaseRegistration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "UpdateDatabaseRegistration", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "UpdateDatabaseRegistration")
	assert.NoError(t, err)

	type UpdateDatabaseRegistrationRequestInfo struct {
		ContainerId string
		Request     goldengate.UpdateDatabaseRegistrationRequest
	}

	var requests []UpdateDatabaseRegistrationRequestInfo
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
			response, err := c.UpdateDatabaseRegistration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientUpdateDeployment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "UpdateDeployment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDeployment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "UpdateDeployment", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "UpdateDeployment")
	assert.NoError(t, err)

	type UpdateDeploymentRequestInfo struct {
		ContainerId string
		Request     goldengate.UpdateDeploymentRequest
	}

	var requests []UpdateDeploymentRequestInfo
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
			response, err := c.UpdateDeployment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
