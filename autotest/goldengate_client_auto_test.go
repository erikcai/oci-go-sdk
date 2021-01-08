package autotest

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"github.com/oracle/oci-go-sdk/v31/goldengate"

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
func TestGoldengateGoldenGateClientChangeDeploymentBackupCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "ChangeDeploymentBackupCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeDeploymentBackupCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "ChangeDeploymentBackupCompartment", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "ChangeDeploymentBackupCompartment")
	assert.NoError(t, err)

	type ChangeDeploymentBackupCompartmentRequestInfo struct {
		ContainerId string
		Request     goldengate.ChangeDeploymentBackupCompartmentRequest
	}

	var requests []ChangeDeploymentBackupCompartmentRequestInfo
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
			response, err := c.ChangeDeploymentBackupCompartment(context.Background(), req.Request)
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
func TestGoldengateGoldenGateClientCreateDeploymentBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "CreateDeploymentBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDeploymentBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "CreateDeploymentBackup", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "CreateDeploymentBackup")
	assert.NoError(t, err)

	type CreateDeploymentBackupRequestInfo struct {
		ContainerId string
		Request     goldengate.CreateDeploymentBackupRequest
	}

	var requests []CreateDeploymentBackupRequestInfo
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
			response, err := c.CreateDeploymentBackup(context.Background(), req.Request)
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
func TestGoldengateGoldenGateClientDeleteDeploymentBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "DeleteDeploymentBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDeploymentBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "DeleteDeploymentBackup", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "DeleteDeploymentBackup")
	assert.NoError(t, err)

	type DeleteDeploymentBackupRequestInfo struct {
		ContainerId string
		Request     goldengate.DeleteDeploymentBackupRequest
	}

	var requests []DeleteDeploymentBackupRequestInfo
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
			response, err := c.DeleteDeploymentBackup(context.Background(), req.Request)
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
func TestGoldengateGoldenGateClientGetDeploymentBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "GetDeploymentBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDeploymentBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "GetDeploymentBackup", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "GetDeploymentBackup")
	assert.NoError(t, err)

	type GetDeploymentBackupRequestInfo struct {
		ContainerId string
		Request     goldengate.GetDeploymentBackupRequest
	}

	var requests []GetDeploymentBackupRequestInfo
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
			response, err := c.GetDeploymentBackup(context.Background(), req.Request)
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
func TestGoldengateGoldenGateClientListDeploymentBackups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "ListDeploymentBackups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDeploymentBackups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "ListDeploymentBackups", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "ListDeploymentBackups")
	assert.NoError(t, err)

	type ListDeploymentBackupsRequestInfo struct {
		ContainerId string
		Request     goldengate.ListDeploymentBackupsRequest
	}

	var requests []ListDeploymentBackupsRequestInfo
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
				r := req.(*goldengate.ListDeploymentBackupsRequest)
				return c.ListDeploymentBackups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]goldengate.ListDeploymentBackupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(goldengate.ListDeploymentBackupsResponse)
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
func TestGoldengateGoldenGateClientRestoreDeployment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "RestoreDeployment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestoreDeployment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "RestoreDeployment", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "RestoreDeployment")
	assert.NoError(t, err)

	type RestoreDeploymentRequestInfo struct {
		ContainerId string
		Request     goldengate.RestoreDeploymentRequest
	}

	var requests []RestoreDeploymentRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]RestoreDeploymentRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["RestoreDeploymentDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "restoreDeploymentType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"DEFAULT": &goldengate.DefaultRestoreDeploymentDetails{},
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
			response, err := c.RestoreDeployment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
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
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]StartDeploymentRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["StartDeploymentDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "startDeploymentType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"DEFAULT": &goldengate.DefaultStartDeploymentDetails{},
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
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]StopDeploymentRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["StopDeploymentDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "stopDeploymentType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"DEFAULT": &goldengate.DefaultStopDeploymentDetails{},
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
			response, err := c.StopDeployment(context.Background(), req.Request)
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

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientUpdateDeploymentBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "UpdateDeploymentBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDeploymentBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "UpdateDeploymentBackup", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "UpdateDeploymentBackup")
	assert.NoError(t, err)

	type UpdateDeploymentBackupRequestInfo struct {
		ContainerId string
		Request     goldengate.UpdateDeploymentBackupRequest
	}

	var requests []UpdateDeploymentBackupRequestInfo
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
			response, err := c.UpdateDeploymentBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="ggs_team_ww_grp@oracle.com" jiraProject="GGS" opsJiraProject="GGS"
func TestGoldengateGoldenGateClientUpgradeDeployment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("goldengate", "UpgradeDeployment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpgradeDeployment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("goldengate", "GoldenGate", "UpgradeDeployment", createGoldengateGoldenGateClientWithProvider)
	assert.NoError(t, err)
	c := cc.(goldengate.GoldenGateClient)

	body, err := testClient.getRequests("goldengate", "UpgradeDeployment")
	assert.NoError(t, err)

	type UpgradeDeploymentRequestInfo struct {
		ContainerId string
		Request     goldengate.UpgradeDeploymentRequest
	}

	var requests []UpgradeDeploymentRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]UpgradeDeploymentRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["UpgradeDeploymentDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "upgradeDeploymentType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"CURRENT_RELEASE": &goldengate.UpgradeDeploymentCurrentReleaseDetails{},
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
			response, err := c.UpgradeDeployment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
