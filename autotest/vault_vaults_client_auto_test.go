package autotest

import (
	"github.com/oracle/oci-go-sdk/v32/common"
	"github.com/oracle/oci-go-sdk/v32/vault"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createVaultsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := vault.NewVaultsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="team_oci_vault_us_grp@oracle.com" jiraProject="SECSVC" opsJiraProject="SI"
func TestVaultsClientCancelSecretDeletion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("vault", "CancelSecretDeletion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CancelSecretDeletion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("vault", "Vaults", "CancelSecretDeletion", createVaultsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(vault.VaultsClient)

	body, err := testClient.getRequests("vault", "CancelSecretDeletion")
	assert.NoError(t, err)

	type CancelSecretDeletionRequestInfo struct {
		ContainerId string
		Request     vault.CancelSecretDeletionRequest
	}

	var requests []CancelSecretDeletionRequestInfo
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
			response, err := c.CancelSecretDeletion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_vault_us_grp@oracle.com" jiraProject="SECSVC" opsJiraProject="SI"
func TestVaultsClientCancelSecretVersionDeletion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("vault", "CancelSecretVersionDeletion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CancelSecretVersionDeletion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("vault", "Vaults", "CancelSecretVersionDeletion", createVaultsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(vault.VaultsClient)

	body, err := testClient.getRequests("vault", "CancelSecretVersionDeletion")
	assert.NoError(t, err)

	type CancelSecretVersionDeletionRequestInfo struct {
		ContainerId string
		Request     vault.CancelSecretVersionDeletionRequest
	}

	var requests []CancelSecretVersionDeletionRequestInfo
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
			response, err := c.CancelSecretVersionDeletion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_vault_us_grp@oracle.com" jiraProject="SECSVC" opsJiraProject="SI"
func TestVaultsClientChangeSecretCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("vault", "ChangeSecretCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeSecretCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("vault", "Vaults", "ChangeSecretCompartment", createVaultsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(vault.VaultsClient)

	body, err := testClient.getRequests("vault", "ChangeSecretCompartment")
	assert.NoError(t, err)

	type ChangeSecretCompartmentRequestInfo struct {
		ContainerId string
		Request     vault.ChangeSecretCompartmentRequest
	}

	var requests []ChangeSecretCompartmentRequestInfo
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
			response, err := c.ChangeSecretCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_vault_us_grp@oracle.com" jiraProject="SECSVC" opsJiraProject="SI"
func TestVaultsClientCreateSecret(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("vault", "CreateSecret")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSecret is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("vault", "Vaults", "CreateSecret", createVaultsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(vault.VaultsClient)

	body, err := testClient.getRequests("vault", "CreateSecret")
	assert.NoError(t, err)

	type CreateSecretRequestInfo struct {
		ContainerId string
		Request     vault.CreateSecretRequest
	}

	var requests []CreateSecretRequestInfo
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
			response, err := c.CreateSecret(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_vault_us_grp@oracle.com" jiraProject="SECSVC" opsJiraProject="SI"
func TestVaultsClientGetSecret(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("vault", "GetSecret")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSecret is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("vault", "Vaults", "GetSecret", createVaultsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(vault.VaultsClient)

	body, err := testClient.getRequests("vault", "GetSecret")
	assert.NoError(t, err)

	type GetSecretRequestInfo struct {
		ContainerId string
		Request     vault.GetSecretRequest
	}

	var requests []GetSecretRequestInfo
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
			response, err := c.GetSecret(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_vault_us_grp@oracle.com" jiraProject="SECSVC" opsJiraProject="SI"
func TestVaultsClientGetSecretVersion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("vault", "GetSecretVersion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSecretVersion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("vault", "Vaults", "GetSecretVersion", createVaultsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(vault.VaultsClient)

	body, err := testClient.getRequests("vault", "GetSecretVersion")
	assert.NoError(t, err)

	type GetSecretVersionRequestInfo struct {
		ContainerId string
		Request     vault.GetSecretVersionRequest
	}

	var requests []GetSecretVersionRequestInfo
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
			response, err := c.GetSecretVersion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_vault_us_grp@oracle.com" jiraProject="SECSVC" opsJiraProject="SI"
func TestVaultsClientListSecretVersions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("vault", "ListSecretVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSecretVersions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("vault", "Vaults", "ListSecretVersions", createVaultsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(vault.VaultsClient)

	body, err := testClient.getRequests("vault", "ListSecretVersions")
	assert.NoError(t, err)

	type ListSecretVersionsRequestInfo struct {
		ContainerId string
		Request     vault.ListSecretVersionsRequest
	}

	var requests []ListSecretVersionsRequestInfo
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
				r := req.(*vault.ListSecretVersionsRequest)
				return c.ListSecretVersions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]vault.ListSecretVersionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(vault.ListSecretVersionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_vault_us_grp@oracle.com" jiraProject="SECSVC" opsJiraProject="SI"
func TestVaultsClientListSecrets(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("vault", "ListSecrets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSecrets is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("vault", "Vaults", "ListSecrets", createVaultsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(vault.VaultsClient)

	body, err := testClient.getRequests("vault", "ListSecrets")
	assert.NoError(t, err)

	type ListSecretsRequestInfo struct {
		ContainerId string
		Request     vault.ListSecretsRequest
	}

	var requests []ListSecretsRequestInfo
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
				r := req.(*vault.ListSecretsRequest)
				return c.ListSecrets(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]vault.ListSecretsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(vault.ListSecretsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_vault_us_grp@oracle.com" jiraProject="SECSVC" opsJiraProject="SI"
func TestVaultsClientScheduleSecretDeletion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("vault", "ScheduleSecretDeletion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ScheduleSecretDeletion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("vault", "Vaults", "ScheduleSecretDeletion", createVaultsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(vault.VaultsClient)

	body, err := testClient.getRequests("vault", "ScheduleSecretDeletion")
	assert.NoError(t, err)

	type ScheduleSecretDeletionRequestInfo struct {
		ContainerId string
		Request     vault.ScheduleSecretDeletionRequest
	}

	var requests []ScheduleSecretDeletionRequestInfo
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
			response, err := c.ScheduleSecretDeletion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_vault_us_grp@oracle.com" jiraProject="SECSVC" opsJiraProject="SI"
func TestVaultsClientScheduleSecretVersionDeletion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("vault", "ScheduleSecretVersionDeletion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ScheduleSecretVersionDeletion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("vault", "Vaults", "ScheduleSecretVersionDeletion", createVaultsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(vault.VaultsClient)

	body, err := testClient.getRequests("vault", "ScheduleSecretVersionDeletion")
	assert.NoError(t, err)

	type ScheduleSecretVersionDeletionRequestInfo struct {
		ContainerId string
		Request     vault.ScheduleSecretVersionDeletionRequest
	}

	var requests []ScheduleSecretVersionDeletionRequestInfo
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
			response, err := c.ScheduleSecretVersionDeletion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_vault_us_grp@oracle.com" jiraProject="SECSVC" opsJiraProject="SI"
func TestVaultsClientUpdateSecret(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("vault", "UpdateSecret")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSecret is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("vault", "Vaults", "UpdateSecret", createVaultsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(vault.VaultsClient)

	body, err := testClient.getRequests("vault", "UpdateSecret")
	assert.NoError(t, err)

	type UpdateSecretRequestInfo struct {
		ContainerId string
		Request     vault.UpdateSecretRequest
	}

	var requests []UpdateSecretRequestInfo
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
			response, err := c.UpdateSecret(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
