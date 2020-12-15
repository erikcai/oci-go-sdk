package autotest

import (
	"github.com/oracle/oci-go-sdk/v31/cloudguard"
	"github.com/oracle/oci-go-sdk/v31/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createCloudguardCloudGuardClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := cloudguard.NewCloudGuardClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientChangeDetectorRecipeCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ChangeDetectorRecipeCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeDetectorRecipeCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ChangeDetectorRecipeCompartment", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ChangeDetectorRecipeCompartment")
	assert.NoError(t, err)

	type ChangeDetectorRecipeCompartmentRequestInfo struct {
		ContainerId string
		Request     cloudguard.ChangeDetectorRecipeCompartmentRequest
	}

	var requests []ChangeDetectorRecipeCompartmentRequestInfo
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
			response, err := c.ChangeDetectorRecipeCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientChangeManagedListCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ChangeManagedListCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeManagedListCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ChangeManagedListCompartment", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ChangeManagedListCompartment")
	assert.NoError(t, err)

	type ChangeManagedListCompartmentRequestInfo struct {
		ContainerId string
		Request     cloudguard.ChangeManagedListCompartmentRequest
	}

	var requests []ChangeManagedListCompartmentRequestInfo
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
			response, err := c.ChangeManagedListCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientChangeResponderRecipeCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ChangeResponderRecipeCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeResponderRecipeCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ChangeResponderRecipeCompartment", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ChangeResponderRecipeCompartment")
	assert.NoError(t, err)

	type ChangeResponderRecipeCompartmentRequestInfo struct {
		ContainerId string
		Request     cloudguard.ChangeResponderRecipeCompartmentRequest
	}

	var requests []ChangeResponderRecipeCompartmentRequestInfo
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
			response, err := c.ChangeResponderRecipeCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientCreateDetectorRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "CreateDetectorRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDetectorRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "CreateDetectorRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "CreateDetectorRecipe")
	assert.NoError(t, err)

	type CreateDetectorRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.CreateDetectorRecipeRequest
	}

	var requests []CreateDetectorRecipeRequestInfo
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
			response, err := c.CreateDetectorRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientCreateManagedList(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "CreateManagedList")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateManagedList is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "CreateManagedList", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "CreateManagedList")
	assert.NoError(t, err)

	type CreateManagedListRequestInfo struct {
		ContainerId string
		Request     cloudguard.CreateManagedListRequest
	}

	var requests []CreateManagedListRequestInfo
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
			response, err := c.CreateManagedList(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientCreateResponderRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "CreateResponderRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateResponderRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "CreateResponderRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "CreateResponderRecipe")
	assert.NoError(t, err)

	type CreateResponderRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.CreateResponderRecipeRequest
	}

	var requests []CreateResponderRecipeRequestInfo
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
			response, err := c.CreateResponderRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientCreateTarget(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "CreateTarget")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTarget is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "CreateTarget", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "CreateTarget")
	assert.NoError(t, err)

	type CreateTargetRequestInfo struct {
		ContainerId string
		Request     cloudguard.CreateTargetRequest
	}

	var requests []CreateTargetRequestInfo
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
			response, err := c.CreateTarget(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientCreateTargetDetectorRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "CreateTargetDetectorRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTargetDetectorRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "CreateTargetDetectorRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "CreateTargetDetectorRecipe")
	assert.NoError(t, err)

	type CreateTargetDetectorRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.CreateTargetDetectorRecipeRequest
	}

	var requests []CreateTargetDetectorRecipeRequestInfo
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
			response, err := c.CreateTargetDetectorRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientCreateTargetResponderRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "CreateTargetResponderRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTargetResponderRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "CreateTargetResponderRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "CreateTargetResponderRecipe")
	assert.NoError(t, err)

	type CreateTargetResponderRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.CreateTargetResponderRecipeRequest
	}

	var requests []CreateTargetResponderRecipeRequestInfo
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
			response, err := c.CreateTargetResponderRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientDeleteDetectorRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "DeleteDetectorRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDetectorRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "DeleteDetectorRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "DeleteDetectorRecipe")
	assert.NoError(t, err)

	type DeleteDetectorRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.DeleteDetectorRecipeRequest
	}

	var requests []DeleteDetectorRecipeRequestInfo
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
			response, err := c.DeleteDetectorRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientDeleteManagedList(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "DeleteManagedList")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteManagedList is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "DeleteManagedList", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "DeleteManagedList")
	assert.NoError(t, err)

	type DeleteManagedListRequestInfo struct {
		ContainerId string
		Request     cloudguard.DeleteManagedListRequest
	}

	var requests []DeleteManagedListRequestInfo
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
			response, err := c.DeleteManagedList(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientDeleteResponderRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "DeleteResponderRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteResponderRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "DeleteResponderRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "DeleteResponderRecipe")
	assert.NoError(t, err)

	type DeleteResponderRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.DeleteResponderRecipeRequest
	}

	var requests []DeleteResponderRecipeRequestInfo
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
			response, err := c.DeleteResponderRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientDeleteTarget(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "DeleteTarget")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteTarget is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "DeleteTarget", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "DeleteTarget")
	assert.NoError(t, err)

	type DeleteTargetRequestInfo struct {
		ContainerId string
		Request     cloudguard.DeleteTargetRequest
	}

	var requests []DeleteTargetRequestInfo
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
			response, err := c.DeleteTarget(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientDeleteTargetDetectorRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "DeleteTargetDetectorRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteTargetDetectorRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "DeleteTargetDetectorRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "DeleteTargetDetectorRecipe")
	assert.NoError(t, err)

	type DeleteTargetDetectorRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.DeleteTargetDetectorRecipeRequest
	}

	var requests []DeleteTargetDetectorRecipeRequestInfo
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
			response, err := c.DeleteTargetDetectorRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientDeleteTargetResponderRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "DeleteTargetResponderRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteTargetResponderRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "DeleteTargetResponderRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "DeleteTargetResponderRecipe")
	assert.NoError(t, err)

	type DeleteTargetResponderRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.DeleteTargetResponderRecipeRequest
	}

	var requests []DeleteTargetResponderRecipeRequestInfo
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
			response, err := c.DeleteTargetResponderRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientExecuteResponderExecution(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ExecuteResponderExecution")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ExecuteResponderExecution is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ExecuteResponderExecution", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ExecuteResponderExecution")
	assert.NoError(t, err)

	type ExecuteResponderExecutionRequestInfo struct {
		ContainerId string
		Request     cloudguard.ExecuteResponderExecutionRequest
	}

	var requests []ExecuteResponderExecutionRequestInfo
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
			response, err := c.ExecuteResponderExecution(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetConditionMetadataType(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetConditionMetadataType")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConditionMetadataType is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetConditionMetadataType", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetConditionMetadataType")
	assert.NoError(t, err)

	type GetConditionMetadataTypeRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetConditionMetadataTypeRequest
	}

	var requests []GetConditionMetadataTypeRequestInfo
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
			response, err := c.GetConditionMetadataType(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetConfiguration", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetConfiguration")
	assert.NoError(t, err)

	type GetConfigurationRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetConfigurationRequest
	}

	var requests []GetConfigurationRequestInfo
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
			response, err := c.GetConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetDetector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetDetector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDetector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetDetector", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetDetector")
	assert.NoError(t, err)

	type GetDetectorRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetDetectorRequest
	}

	var requests []GetDetectorRequestInfo
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
			response, err := c.GetDetector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetDetectorRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetDetectorRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDetectorRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetDetectorRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetDetectorRecipe")
	assert.NoError(t, err)

	type GetDetectorRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetDetectorRecipeRequest
	}

	var requests []GetDetectorRecipeRequestInfo
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
			response, err := c.GetDetectorRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetDetectorRecipeDetectorRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetDetectorRecipeDetectorRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDetectorRecipeDetectorRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetDetectorRecipeDetectorRule", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetDetectorRecipeDetectorRule")
	assert.NoError(t, err)

	type GetDetectorRecipeDetectorRuleRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetDetectorRecipeDetectorRuleRequest
	}

	var requests []GetDetectorRecipeDetectorRuleRequestInfo
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
			response, err := c.GetDetectorRecipeDetectorRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetDetectorRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetDetectorRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDetectorRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetDetectorRule", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetDetectorRule")
	assert.NoError(t, err)

	type GetDetectorRuleRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetDetectorRuleRequest
	}

	var requests []GetDetectorRuleRequestInfo
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
			response, err := c.GetDetectorRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetManagedList(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetManagedList")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetManagedList is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetManagedList", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetManagedList")
	assert.NoError(t, err)

	type GetManagedListRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetManagedListRequest
	}

	var requests []GetManagedListRequestInfo
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
			response, err := c.GetManagedList(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetProblem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetProblem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetProblem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetProblem", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetProblem")
	assert.NoError(t, err)

	type GetProblemRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetProblemRequest
	}

	var requests []GetProblemRequestInfo
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
			response, err := c.GetProblem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetResponderExecution(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetResponderExecution")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetResponderExecution is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetResponderExecution", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetResponderExecution")
	assert.NoError(t, err)

	type GetResponderExecutionRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetResponderExecutionRequest
	}

	var requests []GetResponderExecutionRequestInfo
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
			response, err := c.GetResponderExecution(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetResponderRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetResponderRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetResponderRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetResponderRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetResponderRecipe")
	assert.NoError(t, err)

	type GetResponderRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetResponderRecipeRequest
	}

	var requests []GetResponderRecipeRequestInfo
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
			response, err := c.GetResponderRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetResponderRecipeResponderRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetResponderRecipeResponderRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetResponderRecipeResponderRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetResponderRecipeResponderRule", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetResponderRecipeResponderRule")
	assert.NoError(t, err)

	type GetResponderRecipeResponderRuleRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetResponderRecipeResponderRuleRequest
	}

	var requests []GetResponderRecipeResponderRuleRequestInfo
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
			response, err := c.GetResponderRecipeResponderRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetResponderRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetResponderRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetResponderRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetResponderRule", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetResponderRule")
	assert.NoError(t, err)

	type GetResponderRuleRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetResponderRuleRequest
	}

	var requests []GetResponderRuleRequestInfo
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
			response, err := c.GetResponderRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetTarget(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetTarget")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTarget is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetTarget", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetTarget")
	assert.NoError(t, err)

	type GetTargetRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetTargetRequest
	}

	var requests []GetTargetRequestInfo
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
			response, err := c.GetTarget(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetTargetDetectorRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetTargetDetectorRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTargetDetectorRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetTargetDetectorRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetTargetDetectorRecipe")
	assert.NoError(t, err)

	type GetTargetDetectorRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetTargetDetectorRecipeRequest
	}

	var requests []GetTargetDetectorRecipeRequestInfo
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
			response, err := c.GetTargetDetectorRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetTargetDetectorRecipeDetectorRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetTargetDetectorRecipeDetectorRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTargetDetectorRecipeDetectorRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetTargetDetectorRecipeDetectorRule", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetTargetDetectorRecipeDetectorRule")
	assert.NoError(t, err)

	type GetTargetDetectorRecipeDetectorRuleRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetTargetDetectorRecipeDetectorRuleRequest
	}

	var requests []GetTargetDetectorRecipeDetectorRuleRequestInfo
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
			response, err := c.GetTargetDetectorRecipeDetectorRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetTargetResponderRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetTargetResponderRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTargetResponderRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetTargetResponderRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetTargetResponderRecipe")
	assert.NoError(t, err)

	type GetTargetResponderRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetTargetResponderRecipeRequest
	}

	var requests []GetTargetResponderRecipeRequestInfo
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
			response, err := c.GetTargetResponderRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientGetTargetResponderRecipeResponderRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "GetTargetResponderRecipeResponderRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTargetResponderRecipeResponderRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "GetTargetResponderRecipeResponderRule", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "GetTargetResponderRecipeResponderRule")
	assert.NoError(t, err)

	type GetTargetResponderRecipeResponderRuleRequestInfo struct {
		ContainerId string
		Request     cloudguard.GetTargetResponderRecipeResponderRuleRequest
	}

	var requests []GetTargetResponderRecipeResponderRuleRequestInfo
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
			response, err := c.GetTargetResponderRecipeResponderRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListConditionMetadataTypes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListConditionMetadataTypes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListConditionMetadataTypes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListConditionMetadataTypes", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListConditionMetadataTypes")
	assert.NoError(t, err)

	type ListConditionMetadataTypesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListConditionMetadataTypesRequest
	}

	var requests []ListConditionMetadataTypesRequestInfo
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
				r := req.(*cloudguard.ListConditionMetadataTypesRequest)
				return c.ListConditionMetadataTypes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListConditionMetadataTypesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListConditionMetadataTypesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListDetectorRecipeDetectorRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListDetectorRecipeDetectorRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDetectorRecipeDetectorRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListDetectorRecipeDetectorRules", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListDetectorRecipeDetectorRules")
	assert.NoError(t, err)

	type ListDetectorRecipeDetectorRulesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListDetectorRecipeDetectorRulesRequest
	}

	var requests []ListDetectorRecipeDetectorRulesRequestInfo
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
				r := req.(*cloudguard.ListDetectorRecipeDetectorRulesRequest)
				return c.ListDetectorRecipeDetectorRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListDetectorRecipeDetectorRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListDetectorRecipeDetectorRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListDetectorRecipes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListDetectorRecipes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDetectorRecipes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListDetectorRecipes", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListDetectorRecipes")
	assert.NoError(t, err)

	type ListDetectorRecipesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListDetectorRecipesRequest
	}

	var requests []ListDetectorRecipesRequestInfo
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
				r := req.(*cloudguard.ListDetectorRecipesRequest)
				return c.ListDetectorRecipes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListDetectorRecipesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListDetectorRecipesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListDetectorRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListDetectorRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDetectorRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListDetectorRules", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListDetectorRules")
	assert.NoError(t, err)

	type ListDetectorRulesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListDetectorRulesRequest
	}

	var requests []ListDetectorRulesRequestInfo
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
				r := req.(*cloudguard.ListDetectorRulesRequest)
				return c.ListDetectorRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListDetectorRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListDetectorRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListDetectors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListDetectors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDetectors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListDetectors", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListDetectors")
	assert.NoError(t, err)

	type ListDetectorsRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListDetectorsRequest
	}

	var requests []ListDetectorsRequestInfo
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
				r := req.(*cloudguard.ListDetectorsRequest)
				return c.ListDetectors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListDetectorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListDetectorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListImpactedResources(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListImpactedResources")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListImpactedResources is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListImpactedResources", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListImpactedResources")
	assert.NoError(t, err)

	type ListImpactedResourcesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListImpactedResourcesRequest
	}

	var requests []ListImpactedResourcesRequestInfo
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
				r := req.(*cloudguard.ListImpactedResourcesRequest)
				return c.ListImpactedResources(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListImpactedResourcesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListImpactedResourcesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListManagedListTypes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListManagedListTypes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListManagedListTypes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListManagedListTypes", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListManagedListTypes")
	assert.NoError(t, err)

	type ListManagedListTypesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListManagedListTypesRequest
	}

	var requests []ListManagedListTypesRequestInfo
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
				r := req.(*cloudguard.ListManagedListTypesRequest)
				return c.ListManagedListTypes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListManagedListTypesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListManagedListTypesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListManagedLists(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListManagedLists")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListManagedLists is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListManagedLists", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListManagedLists")
	assert.NoError(t, err)

	type ListManagedListsRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListManagedListsRequest
	}

	var requests []ListManagedListsRequestInfo
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
				r := req.(*cloudguard.ListManagedListsRequest)
				return c.ListManagedLists(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListManagedListsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListManagedListsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListProblemHistories(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListProblemHistories")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListProblemHistories is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListProblemHistories", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListProblemHistories")
	assert.NoError(t, err)

	type ListProblemHistoriesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListProblemHistoriesRequest
	}

	var requests []ListProblemHistoriesRequestInfo
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
				r := req.(*cloudguard.ListProblemHistoriesRequest)
				return c.ListProblemHistories(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListProblemHistoriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListProblemHistoriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListProblems(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListProblems")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListProblems is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListProblems", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListProblems")
	assert.NoError(t, err)

	type ListProblemsRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListProblemsRequest
	}

	var requests []ListProblemsRequestInfo
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
				r := req.(*cloudguard.ListProblemsRequest)
				return c.ListProblems(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListProblemsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListProblemsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListRecommendations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListRecommendations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRecommendations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListRecommendations", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListRecommendations")
	assert.NoError(t, err)

	type ListRecommendationsRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListRecommendationsRequest
	}

	var requests []ListRecommendationsRequestInfo
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
				r := req.(*cloudguard.ListRecommendationsRequest)
				return c.ListRecommendations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListRecommendationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListRecommendationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListResourceTypes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListResourceTypes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListResourceTypes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListResourceTypes", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListResourceTypes")
	assert.NoError(t, err)

	type ListResourceTypesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListResourceTypesRequest
	}

	var requests []ListResourceTypesRequestInfo
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
				r := req.(*cloudguard.ListResourceTypesRequest)
				return c.ListResourceTypes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListResourceTypesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListResourceTypesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListResponderActivities(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListResponderActivities")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListResponderActivities is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListResponderActivities", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListResponderActivities")
	assert.NoError(t, err)

	type ListResponderActivitiesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListResponderActivitiesRequest
	}

	var requests []ListResponderActivitiesRequestInfo
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
				r := req.(*cloudguard.ListResponderActivitiesRequest)
				return c.ListResponderActivities(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListResponderActivitiesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListResponderActivitiesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListResponderExecutions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListResponderExecutions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListResponderExecutions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListResponderExecutions", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListResponderExecutions")
	assert.NoError(t, err)

	type ListResponderExecutionsRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListResponderExecutionsRequest
	}

	var requests []ListResponderExecutionsRequestInfo
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
				r := req.(*cloudguard.ListResponderExecutionsRequest)
				return c.ListResponderExecutions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListResponderExecutionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListResponderExecutionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListResponderRecipeResponderRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListResponderRecipeResponderRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListResponderRecipeResponderRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListResponderRecipeResponderRules", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListResponderRecipeResponderRules")
	assert.NoError(t, err)

	type ListResponderRecipeResponderRulesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListResponderRecipeResponderRulesRequest
	}

	var requests []ListResponderRecipeResponderRulesRequestInfo
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
				r := req.(*cloudguard.ListResponderRecipeResponderRulesRequest)
				return c.ListResponderRecipeResponderRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListResponderRecipeResponderRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListResponderRecipeResponderRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListResponderRecipes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListResponderRecipes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListResponderRecipes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListResponderRecipes", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListResponderRecipes")
	assert.NoError(t, err)

	type ListResponderRecipesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListResponderRecipesRequest
	}

	var requests []ListResponderRecipesRequestInfo
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
				r := req.(*cloudguard.ListResponderRecipesRequest)
				return c.ListResponderRecipes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListResponderRecipesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListResponderRecipesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListResponderRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListResponderRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListResponderRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListResponderRules", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListResponderRules")
	assert.NoError(t, err)

	type ListResponderRulesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListResponderRulesRequest
	}

	var requests []ListResponderRulesRequestInfo
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
				r := req.(*cloudguard.ListResponderRulesRequest)
				return c.ListResponderRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListResponderRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListResponderRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListTargetDetectorRecipeDetectorRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListTargetDetectorRecipeDetectorRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTargetDetectorRecipeDetectorRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListTargetDetectorRecipeDetectorRules", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListTargetDetectorRecipeDetectorRules")
	assert.NoError(t, err)

	type ListTargetDetectorRecipeDetectorRulesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListTargetDetectorRecipeDetectorRulesRequest
	}

	var requests []ListTargetDetectorRecipeDetectorRulesRequestInfo
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
				r := req.(*cloudguard.ListTargetDetectorRecipeDetectorRulesRequest)
				return c.ListTargetDetectorRecipeDetectorRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListTargetDetectorRecipeDetectorRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListTargetDetectorRecipeDetectorRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListTargetDetectorRecipes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListTargetDetectorRecipes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTargetDetectorRecipes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListTargetDetectorRecipes", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListTargetDetectorRecipes")
	assert.NoError(t, err)

	type ListTargetDetectorRecipesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListTargetDetectorRecipesRequest
	}

	var requests []ListTargetDetectorRecipesRequestInfo
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
				r := req.(*cloudguard.ListTargetDetectorRecipesRequest)
				return c.ListTargetDetectorRecipes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListTargetDetectorRecipesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListTargetDetectorRecipesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListTargetResponderRecipeResponderRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListTargetResponderRecipeResponderRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTargetResponderRecipeResponderRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListTargetResponderRecipeResponderRules", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListTargetResponderRecipeResponderRules")
	assert.NoError(t, err)

	type ListTargetResponderRecipeResponderRulesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListTargetResponderRecipeResponderRulesRequest
	}

	var requests []ListTargetResponderRecipeResponderRulesRequestInfo
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
				r := req.(*cloudguard.ListTargetResponderRecipeResponderRulesRequest)
				return c.ListTargetResponderRecipeResponderRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListTargetResponderRecipeResponderRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListTargetResponderRecipeResponderRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListTargetResponderRecipes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListTargetResponderRecipes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTargetResponderRecipes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListTargetResponderRecipes", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListTargetResponderRecipes")
	assert.NoError(t, err)

	type ListTargetResponderRecipesRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListTargetResponderRecipesRequest
	}

	var requests []ListTargetResponderRecipesRequestInfo
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
				r := req.(*cloudguard.ListTargetResponderRecipesRequest)
				return c.ListTargetResponderRecipes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListTargetResponderRecipesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListTargetResponderRecipesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientListTargets(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "ListTargets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTargets is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "ListTargets", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "ListTargets")
	assert.NoError(t, err)

	type ListTargetsRequestInfo struct {
		ContainerId string
		Request     cloudguard.ListTargetsRequest
	}

	var requests []ListTargetsRequestInfo
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
				r := req.(*cloudguard.ListTargetsRequest)
				return c.ListTargets(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.ListTargetsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.ListTargetsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientRequestRiskScores(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "RequestRiskScores")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestRiskScores is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "RequestRiskScores", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "RequestRiskScores")
	assert.NoError(t, err)

	type RequestRiskScoresRequestInfo struct {
		ContainerId string
		Request     cloudguard.RequestRiskScoresRequest
	}

	var requests []RequestRiskScoresRequestInfo
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
				r := req.(*cloudguard.RequestRiskScoresRequest)
				return c.RequestRiskScores(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.RequestRiskScoresResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.RequestRiskScoresResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientRequestSecurityScoreSummarizedTrend(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "RequestSecurityScoreSummarizedTrend")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSecurityScoreSummarizedTrend is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "RequestSecurityScoreSummarizedTrend", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "RequestSecurityScoreSummarizedTrend")
	assert.NoError(t, err)

	type RequestSecurityScoreSummarizedTrendRequestInfo struct {
		ContainerId string
		Request     cloudguard.RequestSecurityScoreSummarizedTrendRequest
	}

	var requests []RequestSecurityScoreSummarizedTrendRequestInfo
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
				r := req.(*cloudguard.RequestSecurityScoreSummarizedTrendRequest)
				return c.RequestSecurityScoreSummarizedTrend(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.RequestSecurityScoreSummarizedTrendResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.RequestSecurityScoreSummarizedTrendResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientRequestSecurityScores(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "RequestSecurityScores")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSecurityScores is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "RequestSecurityScores", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "RequestSecurityScores")
	assert.NoError(t, err)

	type RequestSecurityScoresRequestInfo struct {
		ContainerId string
		Request     cloudguard.RequestSecurityScoresRequest
	}

	var requests []RequestSecurityScoresRequestInfo
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
				r := req.(*cloudguard.RequestSecurityScoresRequest)
				return c.RequestSecurityScores(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.RequestSecurityScoresResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.RequestSecurityScoresResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientRequestSummarizedActivityProblems(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "RequestSummarizedActivityProblems")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSummarizedActivityProblems is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "RequestSummarizedActivityProblems", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "RequestSummarizedActivityProblems")
	assert.NoError(t, err)

	type RequestSummarizedActivityProblemsRequestInfo struct {
		ContainerId string
		Request     cloudguard.RequestSummarizedActivityProblemsRequest
	}

	var requests []RequestSummarizedActivityProblemsRequestInfo
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
				r := req.(*cloudguard.RequestSummarizedActivityProblemsRequest)
				return c.RequestSummarizedActivityProblems(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.RequestSummarizedActivityProblemsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.RequestSummarizedActivityProblemsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientRequestSummarizedProblems(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "RequestSummarizedProblems")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSummarizedProblems is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "RequestSummarizedProblems", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "RequestSummarizedProblems")
	assert.NoError(t, err)

	type RequestSummarizedProblemsRequestInfo struct {
		ContainerId string
		Request     cloudguard.RequestSummarizedProblemsRequest
	}

	var requests []RequestSummarizedProblemsRequestInfo
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
				r := req.(*cloudguard.RequestSummarizedProblemsRequest)
				return c.RequestSummarizedProblems(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.RequestSummarizedProblemsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.RequestSummarizedProblemsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientRequestSummarizedResponderExecutions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "RequestSummarizedResponderExecutions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSummarizedResponderExecutions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "RequestSummarizedResponderExecutions", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "RequestSummarizedResponderExecutions")
	assert.NoError(t, err)

	type RequestSummarizedResponderExecutionsRequestInfo struct {
		ContainerId string
		Request     cloudguard.RequestSummarizedResponderExecutionsRequest
	}

	var requests []RequestSummarizedResponderExecutionsRequestInfo
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
				r := req.(*cloudguard.RequestSummarizedResponderExecutionsRequest)
				return c.RequestSummarizedResponderExecutions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.RequestSummarizedResponderExecutionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.RequestSummarizedResponderExecutionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientRequestSummarizedRiskScores(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "RequestSummarizedRiskScores")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSummarizedRiskScores is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "RequestSummarizedRiskScores", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "RequestSummarizedRiskScores")
	assert.NoError(t, err)

	type RequestSummarizedRiskScoresRequestInfo struct {
		ContainerId string
		Request     cloudguard.RequestSummarizedRiskScoresRequest
	}

	var requests []RequestSummarizedRiskScoresRequestInfo
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
				r := req.(*cloudguard.RequestSummarizedRiskScoresRequest)
				return c.RequestSummarizedRiskScores(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.RequestSummarizedRiskScoresResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.RequestSummarizedRiskScoresResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientRequestSummarizedSecurityScores(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "RequestSummarizedSecurityScores")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSummarizedSecurityScores is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "RequestSummarizedSecurityScores", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "RequestSummarizedSecurityScores")
	assert.NoError(t, err)

	type RequestSummarizedSecurityScoresRequestInfo struct {
		ContainerId string
		Request     cloudguard.RequestSummarizedSecurityScoresRequest
	}

	var requests []RequestSummarizedSecurityScoresRequestInfo
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
				r := req.(*cloudguard.RequestSummarizedSecurityScoresRequest)
				return c.RequestSummarizedSecurityScores(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.RequestSummarizedSecurityScoresResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.RequestSummarizedSecurityScoresResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientRequestSummarizedTrendProblems(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "RequestSummarizedTrendProblems")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSummarizedTrendProblems is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "RequestSummarizedTrendProblems", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "RequestSummarizedTrendProblems")
	assert.NoError(t, err)

	type RequestSummarizedTrendProblemsRequestInfo struct {
		ContainerId string
		Request     cloudguard.RequestSummarizedTrendProblemsRequest
	}

	var requests []RequestSummarizedTrendProblemsRequestInfo
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
				r := req.(*cloudguard.RequestSummarizedTrendProblemsRequest)
				return c.RequestSummarizedTrendProblems(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.RequestSummarizedTrendProblemsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.RequestSummarizedTrendProblemsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientRequestSummarizedTrendResponderExecutions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "RequestSummarizedTrendResponderExecutions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSummarizedTrendResponderExecutions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "RequestSummarizedTrendResponderExecutions", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "RequestSummarizedTrendResponderExecutions")
	assert.NoError(t, err)

	type RequestSummarizedTrendResponderExecutionsRequestInfo struct {
		ContainerId string
		Request     cloudguard.RequestSummarizedTrendResponderExecutionsRequest
	}

	var requests []RequestSummarizedTrendResponderExecutionsRequestInfo
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
				r := req.(*cloudguard.RequestSummarizedTrendResponderExecutionsRequest)
				return c.RequestSummarizedTrendResponderExecutions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.RequestSummarizedTrendResponderExecutionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.RequestSummarizedTrendResponderExecutionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientRequestSummarizedTrendSecurityScores(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "RequestSummarizedTrendSecurityScores")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSummarizedTrendSecurityScores is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "RequestSummarizedTrendSecurityScores", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "RequestSummarizedTrendSecurityScores")
	assert.NoError(t, err)

	type RequestSummarizedTrendSecurityScoresRequestInfo struct {
		ContainerId string
		Request     cloudguard.RequestSummarizedTrendSecurityScoresRequest
	}

	var requests []RequestSummarizedTrendSecurityScoresRequestInfo
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
				r := req.(*cloudguard.RequestSummarizedTrendSecurityScoresRequest)
				return c.RequestSummarizedTrendSecurityScores(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudguard.RequestSummarizedTrendSecurityScoresResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudguard.RequestSummarizedTrendSecurityScoresResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientSkipBulkResponderExecution(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "SkipBulkResponderExecution")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SkipBulkResponderExecution is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "SkipBulkResponderExecution", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "SkipBulkResponderExecution")
	assert.NoError(t, err)

	type SkipBulkResponderExecutionRequestInfo struct {
		ContainerId string
		Request     cloudguard.SkipBulkResponderExecutionRequest
	}

	var requests []SkipBulkResponderExecutionRequestInfo
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
			response, err := c.SkipBulkResponderExecution(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientSkipResponderExecution(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "SkipResponderExecution")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SkipResponderExecution is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "SkipResponderExecution", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "SkipResponderExecution")
	assert.NoError(t, err)

	type SkipResponderExecutionRequestInfo struct {
		ContainerId string
		Request     cloudguard.SkipResponderExecutionRequest
	}

	var requests []SkipResponderExecutionRequestInfo
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
			response, err := c.SkipResponderExecution(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientTriggerResponder(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "TriggerResponder")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("TriggerResponder is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "TriggerResponder", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "TriggerResponder")
	assert.NoError(t, err)

	type TriggerResponderRequestInfo struct {
		ContainerId string
		Request     cloudguard.TriggerResponderRequest
	}

	var requests []TriggerResponderRequestInfo
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
			response, err := c.TriggerResponder(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientUpdateBulkProblemStatus(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "UpdateBulkProblemStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBulkProblemStatus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "UpdateBulkProblemStatus", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "UpdateBulkProblemStatus")
	assert.NoError(t, err)

	type UpdateBulkProblemStatusRequestInfo struct {
		ContainerId string
		Request     cloudguard.UpdateBulkProblemStatusRequest
	}

	var requests []UpdateBulkProblemStatusRequestInfo
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
			response, err := c.UpdateBulkProblemStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientUpdateConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "UpdateConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "UpdateConfiguration", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "UpdateConfiguration")
	assert.NoError(t, err)

	type UpdateConfigurationRequestInfo struct {
		ContainerId string
		Request     cloudguard.UpdateConfigurationRequest
	}

	var requests []UpdateConfigurationRequestInfo
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
			response, err := c.UpdateConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientUpdateDetectorRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "UpdateDetectorRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDetectorRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "UpdateDetectorRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "UpdateDetectorRecipe")
	assert.NoError(t, err)

	type UpdateDetectorRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.UpdateDetectorRecipeRequest
	}

	var requests []UpdateDetectorRecipeRequestInfo
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
			response, err := c.UpdateDetectorRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientUpdateDetectorRecipeDetectorRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "UpdateDetectorRecipeDetectorRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDetectorRecipeDetectorRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "UpdateDetectorRecipeDetectorRule", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "UpdateDetectorRecipeDetectorRule")
	assert.NoError(t, err)

	type UpdateDetectorRecipeDetectorRuleRequestInfo struct {
		ContainerId string
		Request     cloudguard.UpdateDetectorRecipeDetectorRuleRequest
	}

	var requests []UpdateDetectorRecipeDetectorRuleRequestInfo
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
			response, err := c.UpdateDetectorRecipeDetectorRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientUpdateManagedList(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "UpdateManagedList")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateManagedList is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "UpdateManagedList", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "UpdateManagedList")
	assert.NoError(t, err)

	type UpdateManagedListRequestInfo struct {
		ContainerId string
		Request     cloudguard.UpdateManagedListRequest
	}

	var requests []UpdateManagedListRequestInfo
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
			response, err := c.UpdateManagedList(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientUpdateProblemStatus(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "UpdateProblemStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateProblemStatus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "UpdateProblemStatus", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "UpdateProblemStatus")
	assert.NoError(t, err)

	type UpdateProblemStatusRequestInfo struct {
		ContainerId string
		Request     cloudguard.UpdateProblemStatusRequest
	}

	var requests []UpdateProblemStatusRequestInfo
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
			response, err := c.UpdateProblemStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientUpdateResponderRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "UpdateResponderRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateResponderRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "UpdateResponderRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "UpdateResponderRecipe")
	assert.NoError(t, err)

	type UpdateResponderRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.UpdateResponderRecipeRequest
	}

	var requests []UpdateResponderRecipeRequestInfo
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
			response, err := c.UpdateResponderRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientUpdateResponderRecipeResponderRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "UpdateResponderRecipeResponderRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateResponderRecipeResponderRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "UpdateResponderRecipeResponderRule", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "UpdateResponderRecipeResponderRule")
	assert.NoError(t, err)

	type UpdateResponderRecipeResponderRuleRequestInfo struct {
		ContainerId string
		Request     cloudguard.UpdateResponderRecipeResponderRuleRequest
	}

	var requests []UpdateResponderRecipeResponderRuleRequestInfo
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
			response, err := c.UpdateResponderRecipeResponderRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientUpdateTarget(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "UpdateTarget")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTarget is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "UpdateTarget", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "UpdateTarget")
	assert.NoError(t, err)

	type UpdateTargetRequestInfo struct {
		ContainerId string
		Request     cloudguard.UpdateTargetRequest
	}

	var requests []UpdateTargetRequestInfo
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
			response, err := c.UpdateTarget(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientUpdateTargetDetectorRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "UpdateTargetDetectorRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTargetDetectorRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "UpdateTargetDetectorRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "UpdateTargetDetectorRecipe")
	assert.NoError(t, err)

	type UpdateTargetDetectorRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.UpdateTargetDetectorRecipeRequest
	}

	var requests []UpdateTargetDetectorRecipeRequestInfo
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
			response, err := c.UpdateTargetDetectorRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientUpdateTargetDetectorRecipeDetectorRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "UpdateTargetDetectorRecipeDetectorRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTargetDetectorRecipeDetectorRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "UpdateTargetDetectorRecipeDetectorRule", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "UpdateTargetDetectorRecipeDetectorRule")
	assert.NoError(t, err)

	type UpdateTargetDetectorRecipeDetectorRuleRequestInfo struct {
		ContainerId string
		Request     cloudguard.UpdateTargetDetectorRecipeDetectorRuleRequest
	}

	var requests []UpdateTargetDetectorRecipeDetectorRuleRequestInfo
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
			response, err := c.UpdateTargetDetectorRecipeDetectorRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientUpdateTargetResponderRecipe(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "UpdateTargetResponderRecipe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTargetResponderRecipe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "UpdateTargetResponderRecipe", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "UpdateTargetResponderRecipe")
	assert.NoError(t, err)

	type UpdateTargetResponderRecipeRequestInfo struct {
		ContainerId string
		Request     cloudguard.UpdateTargetResponderRecipeRequest
	}

	var requests []UpdateTargetResponderRecipeRequestInfo
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
			response, err := c.UpdateTargetResponderRecipe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="seccen-engg_ww_grp@oracle.com" jiraProject="SECCEN" opsJiraProject="SECCEN-OPS"
func TestCloudguardCloudGuardClientUpdateTargetResponderRecipeResponderRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudguard", "UpdateTargetResponderRecipeResponderRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTargetResponderRecipeResponderRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudguard", "CloudGuard", "UpdateTargetResponderRecipeResponderRule", createCloudguardCloudGuardClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudguard.CloudGuardClient)

	body, err := testClient.getRequests("cloudguard", "UpdateTargetResponderRecipeResponderRule")
	assert.NoError(t, err)

	type UpdateTargetResponderRecipeResponderRuleRequestInfo struct {
		ContainerId string
		Request     cloudguard.UpdateTargetResponderRecipeResponderRuleRequest
	}

	var requests []UpdateTargetResponderRecipeResponderRuleRequestInfo
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
			response, err := c.UpdateTargetResponderRecipeResponderRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
