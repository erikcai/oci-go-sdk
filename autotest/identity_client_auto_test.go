package autotest

import (
	"github.com/oracle/oci-go-sdk/v27/common"
	"github.com/oracle/oci-go-sdk/v27/identity"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createIdentityClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := identity.NewIdentityClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientActivateMfaTotpDevice(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ActivateMfaTotpDevice")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ActivateMfaTotpDevice is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ActivateMfaTotpDevice", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ActivateMfaTotpDevice")
	assert.NoError(t, err)

	type ActivateMfaTotpDeviceRequestInfo struct {
		ContainerId string
		Request     identity.ActivateMfaTotpDeviceRequest
	}

	var requests []ActivateMfaTotpDeviceRequestInfo
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
			response, err := c.ActivateMfaTotpDevice(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientAddUserToGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "AddUserToGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AddUserToGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "AddUserToGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "AddUserToGroup")
	assert.NoError(t, err)

	type AddUserToGroupRequestInfo struct {
		ContainerId string
		Request     identity.AddUserToGroupRequest
	}

	var requests []AddUserToGroupRequestInfo
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
			response, err := c.AddUserToGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientAssembleEffectiveTagSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "AssembleEffectiveTagSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AssembleEffectiveTagSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "AssembleEffectiveTagSet", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "AssembleEffectiveTagSet")
	assert.NoError(t, err)

	type AssembleEffectiveTagSetRequestInfo struct {
		ContainerId string
		Request     identity.AssembleEffectiveTagSetRequest
	}

	var requests []AssembleEffectiveTagSetRequestInfo
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
			response, err := c.AssembleEffectiveTagSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientBulkDeleteResources(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "BulkDeleteResources")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("BulkDeleteResources is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "BulkDeleteResources", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "BulkDeleteResources")
	assert.NoError(t, err)

	type BulkDeleteResourcesRequestInfo struct {
		ContainerId string
		Request     identity.BulkDeleteResourcesRequest
	}

	var requests []BulkDeleteResourcesRequestInfo
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
			response, err := c.BulkDeleteResources(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientBulkDeleteTags(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "BulkDeleteTags")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("BulkDeleteTags is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "BulkDeleteTags", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "BulkDeleteTags")
	assert.NoError(t, err)

	type BulkDeleteTagsRequestInfo struct {
		ContainerId string
		Request     identity.BulkDeleteTagsRequest
	}

	var requests []BulkDeleteTagsRequestInfo
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
			response, err := c.BulkDeleteTags(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientBulkEditTags(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "BulkEditTags")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("BulkEditTags is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "BulkEditTags", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "BulkEditTags")
	assert.NoError(t, err)

	type BulkEditTagsRequestInfo struct {
		ContainerId string
		Request     identity.BulkEditTagsRequest
	}

	var requests []BulkEditTagsRequestInfo
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
			response, err := c.BulkEditTags(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientBulkMoveResources(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "BulkMoveResources")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("BulkMoveResources is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "BulkMoveResources", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "BulkMoveResources")
	assert.NoError(t, err)

	type BulkMoveResourcesRequestInfo struct {
		ContainerId string
		Request     identity.BulkMoveResourcesRequest
	}

	var requests []BulkMoveResourcesRequestInfo
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
			response, err := c.BulkMoveResources(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCascadeDeleteTagNamespace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CascadeDeleteTagNamespace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CascadeDeleteTagNamespace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CascadeDeleteTagNamespace", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CascadeDeleteTagNamespace")
	assert.NoError(t, err)

	type CascadeDeleteTagNamespaceRequestInfo struct {
		ContainerId string
		Request     identity.CascadeDeleteTagNamespaceRequest
	}

	var requests []CascadeDeleteTagNamespaceRequestInfo
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
			response, err := c.CascadeDeleteTagNamespace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientChangeTagNamespaceCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ChangeTagNamespaceCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeTagNamespaceCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ChangeTagNamespaceCompartment", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ChangeTagNamespaceCompartment")
	assert.NoError(t, err)

	type ChangeTagNamespaceCompartmentRequestInfo struct {
		ContainerId string
		Request     identity.ChangeTagNamespaceCompartmentRequest
	}

	var requests []ChangeTagNamespaceCompartmentRequestInfo
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
			response, err := c.ChangeTagNamespaceCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateAuthToken(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateAuthToken")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAuthToken is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateAuthToken", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateAuthToken")
	assert.NoError(t, err)

	type CreateAuthTokenRequestInfo struct {
		ContainerId string
		Request     identity.CreateAuthTokenRequest
	}

	var requests []CreateAuthTokenRequestInfo
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
			response, err := c.CreateAuthToken(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateCompartment", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateCompartment")
	assert.NoError(t, err)

	type CreateCompartmentRequestInfo struct {
		ContainerId string
		Request     identity.CreateCompartmentRequest
	}

	var requests []CreateCompartmentRequestInfo
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
			response, err := c.CreateCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateCustomerSecretKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateCustomerSecretKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCustomerSecretKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateCustomerSecretKey", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateCustomerSecretKey")
	assert.NoError(t, err)

	type CreateCustomerSecretKeyRequestInfo struct {
		ContainerId string
		Request     identity.CreateCustomerSecretKeyRequest
	}

	var requests []CreateCustomerSecretKeyRequestInfo
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
			response, err := c.CreateCustomerSecretKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateDynamicGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateDynamicGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDynamicGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateDynamicGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateDynamicGroup")
	assert.NoError(t, err)

	type CreateDynamicGroupRequestInfo struct {
		ContainerId string
		Request     identity.CreateDynamicGroupRequest
	}

	var requests []CreateDynamicGroupRequestInfo
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
			response, err := c.CreateDynamicGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateGroup")
	assert.NoError(t, err)

	type CreateGroupRequestInfo struct {
		ContainerId string
		Request     identity.CreateGroupRequest
	}

	var requests []CreateGroupRequestInfo
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
			response, err := c.CreateGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateIdentityProvider(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateIdentityProvider")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateIdentityProvider is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateIdentityProvider", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateIdentityProvider")
	assert.NoError(t, err)

	type CreateIdentityProviderRequestInfo struct {
		ContainerId string
		Request     identity.CreateIdentityProviderRequest
	}

	var requests []CreateIdentityProviderRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateIdentityProviderRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateIdentityProviderDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "protocol",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"SAML2": &identity.CreateSaml2IdentityProviderDetails{},
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
			response, err := c.CreateIdentityProvider(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateIdpGroupMapping(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateIdpGroupMapping")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateIdpGroupMapping is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateIdpGroupMapping", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateIdpGroupMapping")
	assert.NoError(t, err)

	type CreateIdpGroupMappingRequestInfo struct {
		ContainerId string
		Request     identity.CreateIdpGroupMappingRequest
	}

	var requests []CreateIdpGroupMappingRequestInfo
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
			response, err := c.CreateIdpGroupMapping(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateManagedCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateManagedCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateManagedCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateManagedCompartment", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateManagedCompartment")
	assert.NoError(t, err)

	type CreateManagedCompartmentRequestInfo struct {
		ContainerId string
		Request     identity.CreateManagedCompartmentRequest
	}

	var requests []CreateManagedCompartmentRequestInfo
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
			response, err := c.CreateManagedCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateMfaTotpDevice(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateMfaTotpDevice")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateMfaTotpDevice is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateMfaTotpDevice", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateMfaTotpDevice")
	assert.NoError(t, err)

	type CreateMfaTotpDeviceRequestInfo struct {
		ContainerId string
		Request     identity.CreateMfaTotpDeviceRequest
	}

	var requests []CreateMfaTotpDeviceRequestInfo
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
			response, err := c.CreateMfaTotpDevice(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateNetworkSource(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateNetworkSource")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateNetworkSource is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateNetworkSource", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateNetworkSource")
	assert.NoError(t, err)

	type CreateNetworkSourceRequestInfo struct {
		ContainerId string
		Request     identity.CreateNetworkSourceRequest
	}

	var requests []CreateNetworkSourceRequestInfo
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
			response, err := c.CreateNetworkSource(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateOAuthClientCredential(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateOAuthClientCredential")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateOAuthClientCredential is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateOAuthClientCredential", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateOAuthClientCredential")
	assert.NoError(t, err)

	type CreateOAuthClientCredentialRequestInfo struct {
		ContainerId string
		Request     identity.CreateOAuthClientCredentialRequest
	}

	var requests []CreateOAuthClientCredentialRequestInfo
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
			response, err := c.CreateOAuthClientCredential(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateOrResetUIPassword(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateOrResetUIPassword")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateOrResetUIPassword is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateOrResetUIPassword", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateOrResetUIPassword")
	assert.NoError(t, err)

	type CreateOrResetUIPasswordRequestInfo struct {
		ContainerId string
		Request     identity.CreateOrResetUIPasswordRequest
	}

	var requests []CreateOrResetUIPasswordRequestInfo
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
			response, err := c.CreateOrResetUIPassword(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreatePolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreatePolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreatePolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreatePolicy", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreatePolicy")
	assert.NoError(t, err)

	type CreatePolicyRequestInfo struct {
		ContainerId string
		Request     identity.CreatePolicyRequest
	}

	var requests []CreatePolicyRequestInfo
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
			response, err := c.CreatePolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateRegionSubscription(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateRegionSubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateRegionSubscription is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateRegionSubscription", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateRegionSubscription")
	assert.NoError(t, err)

	type CreateRegionSubscriptionRequestInfo struct {
		ContainerId string
		Request     identity.CreateRegionSubscriptionRequest
	}

	var requests []CreateRegionSubscriptionRequestInfo
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
			response, err := c.CreateRegionSubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateSmtpCredential(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateSmtpCredential")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSmtpCredential is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateSmtpCredential", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateSmtpCredential")
	assert.NoError(t, err)

	type CreateSmtpCredentialRequestInfo struct {
		ContainerId string
		Request     identity.CreateSmtpCredentialRequest
	}

	var requests []CreateSmtpCredentialRequestInfo
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
			response, err := c.CreateSmtpCredential(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateSwiftPassword(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateSwiftPassword")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSwiftPassword is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateSwiftPassword", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateSwiftPassword")
	assert.NoError(t, err)

	type CreateSwiftPasswordRequestInfo struct {
		ContainerId string
		Request     identity.CreateSwiftPasswordRequest
	}

	var requests []CreateSwiftPasswordRequestInfo
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
			response, err := c.CreateSwiftPassword(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateTag(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateTag")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTag is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateTag", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateTag")
	assert.NoError(t, err)

	type CreateTagRequestInfo struct {
		ContainerId string
		Request     identity.CreateTagRequest
	}

	var requests []CreateTagRequestInfo
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
			response, err := c.CreateTag(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateTagDefault(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateTagDefault")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTagDefault is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateTagDefault", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateTagDefault")
	assert.NoError(t, err)

	type CreateTagDefaultRequestInfo struct {
		ContainerId string
		Request     identity.CreateTagDefaultRequest
	}

	var requests []CreateTagDefaultRequestInfo
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
			response, err := c.CreateTagDefault(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateTagNamespace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateTagNamespace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTagNamespace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateTagNamespace", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateTagNamespace")
	assert.NoError(t, err)

	type CreateTagNamespaceRequestInfo struct {
		ContainerId string
		Request     identity.CreateTagNamespaceRequest
	}

	var requests []CreateTagNamespaceRequestInfo
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
			response, err := c.CreateTagNamespace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateTagRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateTagRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTagRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateTagRule", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateTagRule")
	assert.NoError(t, err)

	type CreateTagRuleRequestInfo struct {
		ContainerId string
		Request     identity.CreateTagRuleRequest
	}

	var requests []CreateTagRuleRequestInfo
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
			response, err := c.CreateTagRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientCreateUser(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "CreateUser")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateUser is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateUser", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateUser")
	assert.NoError(t, err)

	type CreateUserRequestInfo struct {
		ContainerId string
		Request     identity.CreateUserRequest
	}

	var requests []CreateUserRequestInfo
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
			response, err := c.CreateUser(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteApiKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteApiKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteApiKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteApiKey", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteApiKey")
	assert.NoError(t, err)

	type DeleteApiKeyRequestInfo struct {
		ContainerId string
		Request     identity.DeleteApiKeyRequest
	}

	var requests []DeleteApiKeyRequestInfo
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
			response, err := c.DeleteApiKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteAuthToken(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteAuthToken")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAuthToken is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteAuthToken", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteAuthToken")
	assert.NoError(t, err)

	type DeleteAuthTokenRequestInfo struct {
		ContainerId string
		Request     identity.DeleteAuthTokenRequest
	}

	var requests []DeleteAuthTokenRequestInfo
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
			response, err := c.DeleteAuthToken(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteCompartment", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteCompartment")
	assert.NoError(t, err)

	type DeleteCompartmentRequestInfo struct {
		ContainerId string
		Request     identity.DeleteCompartmentRequest
	}

	var requests []DeleteCompartmentRequestInfo
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
			response, err := c.DeleteCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteCustomerSecretKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteCustomerSecretKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCustomerSecretKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteCustomerSecretKey", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteCustomerSecretKey")
	assert.NoError(t, err)

	type DeleteCustomerSecretKeyRequestInfo struct {
		ContainerId string
		Request     identity.DeleteCustomerSecretKeyRequest
	}

	var requests []DeleteCustomerSecretKeyRequestInfo
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
			response, err := c.DeleteCustomerSecretKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteDynamicGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteDynamicGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDynamicGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteDynamicGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteDynamicGroup")
	assert.NoError(t, err)

	type DeleteDynamicGroupRequestInfo struct {
		ContainerId string
		Request     identity.DeleteDynamicGroupRequest
	}

	var requests []DeleteDynamicGroupRequestInfo
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
			response, err := c.DeleteDynamicGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteGroup")
	assert.NoError(t, err)

	type DeleteGroupRequestInfo struct {
		ContainerId string
		Request     identity.DeleteGroupRequest
	}

	var requests []DeleteGroupRequestInfo
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
			response, err := c.DeleteGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteIdentityProvider(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteIdentityProvider")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteIdentityProvider is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteIdentityProvider", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteIdentityProvider")
	assert.NoError(t, err)

	type DeleteIdentityProviderRequestInfo struct {
		ContainerId string
		Request     identity.DeleteIdentityProviderRequest
	}

	var requests []DeleteIdentityProviderRequestInfo
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
			response, err := c.DeleteIdentityProvider(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteIdpGroupMapping(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteIdpGroupMapping")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteIdpGroupMapping is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteIdpGroupMapping", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteIdpGroupMapping")
	assert.NoError(t, err)

	type DeleteIdpGroupMappingRequestInfo struct {
		ContainerId string
		Request     identity.DeleteIdpGroupMappingRequest
	}

	var requests []DeleteIdpGroupMappingRequestInfo
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
			response, err := c.DeleteIdpGroupMapping(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteMfaTotpDevice(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteMfaTotpDevice")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteMfaTotpDevice is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteMfaTotpDevice", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteMfaTotpDevice")
	assert.NoError(t, err)

	type DeleteMfaTotpDeviceRequestInfo struct {
		ContainerId string
		Request     identity.DeleteMfaTotpDeviceRequest
	}

	var requests []DeleteMfaTotpDeviceRequestInfo
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
			response, err := c.DeleteMfaTotpDevice(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteNetworkSource(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteNetworkSource")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteNetworkSource is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteNetworkSource", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteNetworkSource")
	assert.NoError(t, err)

	type DeleteNetworkSourceRequestInfo struct {
		ContainerId string
		Request     identity.DeleteNetworkSourceRequest
	}

	var requests []DeleteNetworkSourceRequestInfo
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
			response, err := c.DeleteNetworkSource(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteOAuthClientCredential(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteOAuthClientCredential")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteOAuthClientCredential is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteOAuthClientCredential", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteOAuthClientCredential")
	assert.NoError(t, err)

	type DeleteOAuthClientCredentialRequestInfo struct {
		ContainerId string
		Request     identity.DeleteOAuthClientCredentialRequest
	}

	var requests []DeleteOAuthClientCredentialRequestInfo
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
			response, err := c.DeleteOAuthClientCredential(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeletePolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeletePolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeletePolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeletePolicy", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeletePolicy")
	assert.NoError(t, err)

	type DeletePolicyRequestInfo struct {
		ContainerId string
		Request     identity.DeletePolicyRequest
	}

	var requests []DeletePolicyRequestInfo
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
			response, err := c.DeletePolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteSmtpCredential(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteSmtpCredential")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSmtpCredential is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteSmtpCredential", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteSmtpCredential")
	assert.NoError(t, err)

	type DeleteSmtpCredentialRequestInfo struct {
		ContainerId string
		Request     identity.DeleteSmtpCredentialRequest
	}

	var requests []DeleteSmtpCredentialRequestInfo
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
			response, err := c.DeleteSmtpCredential(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteSwiftPassword(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteSwiftPassword")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSwiftPassword is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteSwiftPassword", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteSwiftPassword")
	assert.NoError(t, err)

	type DeleteSwiftPasswordRequestInfo struct {
		ContainerId string
		Request     identity.DeleteSwiftPasswordRequest
	}

	var requests []DeleteSwiftPasswordRequestInfo
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
			response, err := c.DeleteSwiftPassword(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteTag(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteTag")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteTag is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteTag", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteTag")
	assert.NoError(t, err)

	type DeleteTagRequestInfo struct {
		ContainerId string
		Request     identity.DeleteTagRequest
	}

	var requests []DeleteTagRequestInfo
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
			response, err := c.DeleteTag(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteTagDefault(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteTagDefault")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteTagDefault is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteTagDefault", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteTagDefault")
	assert.NoError(t, err)

	type DeleteTagDefaultRequestInfo struct {
		ContainerId string
		Request     identity.DeleteTagDefaultRequest
	}

	var requests []DeleteTagDefaultRequestInfo
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
			response, err := c.DeleteTagDefault(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteTagNamespace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteTagNamespace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteTagNamespace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteTagNamespace", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteTagNamespace")
	assert.NoError(t, err)

	type DeleteTagNamespaceRequestInfo struct {
		ContainerId string
		Request     identity.DeleteTagNamespaceRequest
	}

	var requests []DeleteTagNamespaceRequestInfo
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
			response, err := c.DeleteTagNamespace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteTagRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteTagRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteTagRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteTagRule", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteTagRule")
	assert.NoError(t, err)

	type DeleteTagRuleRequestInfo struct {
		ContainerId string
		Request     identity.DeleteTagRuleRequest
	}

	var requests []DeleteTagRuleRequestInfo
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
			response, err := c.DeleteTagRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientDeleteUser(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "DeleteUser")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteUser is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteUser", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteUser")
	assert.NoError(t, err)

	type DeleteUserRequestInfo struct {
		ContainerId string
		Request     identity.DeleteUserRequest
	}

	var requests []DeleteUserRequestInfo
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
			response, err := c.DeleteUser(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGenerateTotpSeed(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GenerateTotpSeed")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GenerateTotpSeed is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GenerateTotpSeed", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GenerateTotpSeed")
	assert.NoError(t, err)

	type GenerateTotpSeedRequestInfo struct {
		ContainerId string
		Request     identity.GenerateTotpSeedRequest
	}

	var requests []GenerateTotpSeedRequestInfo
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
			response, err := c.GenerateTotpSeed(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetAccountByEntitlementId(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetAccountByEntitlementId")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAccountByEntitlementId is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetAccountByEntitlementId", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetAccountByEntitlementId")
	assert.NoError(t, err)

	type GetAccountByEntitlementIdRequestInfo struct {
		ContainerId string
		Request     identity.GetAccountByEntitlementIdRequest
	}

	var requests []GetAccountByEntitlementIdRequestInfo
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
			response, err := c.GetAccountByEntitlementId(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetAuthenticationPolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetAuthenticationPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAuthenticationPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetAuthenticationPolicy", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetAuthenticationPolicy")
	assert.NoError(t, err)

	type GetAuthenticationPolicyRequestInfo struct {
		ContainerId string
		Request     identity.GetAuthenticationPolicyRequest
	}

	var requests []GetAuthenticationPolicyRequestInfo
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
			response, err := c.GetAuthenticationPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetCompartment", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetCompartment")
	assert.NoError(t, err)

	type GetCompartmentRequestInfo struct {
		ContainerId string
		Request     identity.GetCompartmentRequest
	}

	var requests []GetCompartmentRequestInfo
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
			response, err := c.GetCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetDynamicGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetDynamicGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDynamicGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetDynamicGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetDynamicGroup")
	assert.NoError(t, err)

	type GetDynamicGroupRequestInfo struct {
		ContainerId string
		Request     identity.GetDynamicGroupRequest
	}

	var requests []GetDynamicGroupRequestInfo
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
			response, err := c.GetDynamicGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetGroup")
	assert.NoError(t, err)

	type GetGroupRequestInfo struct {
		ContainerId string
		Request     identity.GetGroupRequest
	}

	var requests []GetGroupRequestInfo
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
			response, err := c.GetGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetIdentityProvider(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetIdentityProvider")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetIdentityProvider is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetIdentityProvider", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetIdentityProvider")
	assert.NoError(t, err)

	type GetIdentityProviderRequestInfo struct {
		ContainerId string
		Request     identity.GetIdentityProviderRequest
	}

	var requests []GetIdentityProviderRequestInfo
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
			response, err := c.GetIdentityProvider(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetIdpGroupMapping(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetIdpGroupMapping")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetIdpGroupMapping is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetIdpGroupMapping", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetIdpGroupMapping")
	assert.NoError(t, err)

	type GetIdpGroupMappingRequestInfo struct {
		ContainerId string
		Request     identity.GetIdpGroupMappingRequest
	}

	var requests []GetIdpGroupMappingRequestInfo
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
			response, err := c.GetIdpGroupMapping(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetMfaTotpDevice(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetMfaTotpDevice")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetMfaTotpDevice is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetMfaTotpDevice", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetMfaTotpDevice")
	assert.NoError(t, err)

	type GetMfaTotpDeviceRequestInfo struct {
		ContainerId string
		Request     identity.GetMfaTotpDeviceRequest
	}

	var requests []GetMfaTotpDeviceRequestInfo
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
			response, err := c.GetMfaTotpDevice(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetNetworkSource(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetNetworkSource")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetNetworkSource is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetNetworkSource", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetNetworkSource")
	assert.NoError(t, err)

	type GetNetworkSourceRequestInfo struct {
		ContainerId string
		Request     identity.GetNetworkSourceRequest
	}

	var requests []GetNetworkSourceRequestInfo
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
			response, err := c.GetNetworkSource(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetPolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetPolicy", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetPolicy")
	assert.NoError(t, err)

	type GetPolicyRequestInfo struct {
		ContainerId string
		Request     identity.GetPolicyRequest
	}

	var requests []GetPolicyRequestInfo
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
			response, err := c.GetPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetTag(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetTag")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTag is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetTag", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetTag")
	assert.NoError(t, err)

	type GetTagRequestInfo struct {
		ContainerId string
		Request     identity.GetTagRequest
	}

	var requests []GetTagRequestInfo
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
			response, err := c.GetTag(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetTagDefault(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetTagDefault")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTagDefault is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetTagDefault", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetTagDefault")
	assert.NoError(t, err)

	type GetTagDefaultRequestInfo struct {
		ContainerId string
		Request     identity.GetTagDefaultRequest
	}

	var requests []GetTagDefaultRequestInfo
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
			response, err := c.GetTagDefault(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetTagNamespace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetTagNamespace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTagNamespace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetTagNamespace", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetTagNamespace")
	assert.NoError(t, err)

	type GetTagNamespaceRequestInfo struct {
		ContainerId string
		Request     identity.GetTagNamespaceRequest
	}

	var requests []GetTagNamespaceRequestInfo
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
			response, err := c.GetTagNamespace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetTagRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetTagRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTagRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetTagRule", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetTagRule")
	assert.NoError(t, err)

	type GetTagRuleRequestInfo struct {
		ContainerId string
		Request     identity.GetTagRuleRequest
	}

	var requests []GetTagRuleRequestInfo
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
			response, err := c.GetTagRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetTaggingWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetTaggingWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTaggingWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetTaggingWorkRequest", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetTaggingWorkRequest")
	assert.NoError(t, err)

	type GetTaggingWorkRequestRequestInfo struct {
		ContainerId string
		Request     identity.GetTaggingWorkRequestRequest
	}

	var requests []GetTaggingWorkRequestRequestInfo
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
			response, err := c.GetTaggingWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetTenancy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetTenancy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTenancy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetTenancy", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetTenancy")
	assert.NoError(t, err)

	type GetTenancyRequestInfo struct {
		ContainerId string
		Request     identity.GetTenancyRequest
	}

	var requests []GetTenancyRequestInfo
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
			response, err := c.GetTenancy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetUser(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetUser")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetUser is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetUser", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetUser")
	assert.NoError(t, err)

	type GetUserRequestInfo struct {
		ContainerId string
		Request     identity.GetUserRequest
	}

	var requests []GetUserRequestInfo
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
			response, err := c.GetUser(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetUserGroupMembership(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetUserGroupMembership")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetUserGroupMembership is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetUserGroupMembership", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetUserGroupMembership")
	assert.NoError(t, err)

	type GetUserGroupMembershipRequestInfo struct {
		ContainerId string
		Request     identity.GetUserGroupMembershipRequest
	}

	var requests []GetUserGroupMembershipRequestInfo
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
			response, err := c.GetUserGroupMembership(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetUserUIPasswordInformation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetUserUIPasswordInformation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetUserUIPasswordInformation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetUserUIPasswordInformation", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetUserUIPasswordInformation")
	assert.NoError(t, err)

	type GetUserUIPasswordInformationRequestInfo struct {
		ContainerId string
		Request     identity.GetUserUIPasswordInformationRequest
	}

	var requests []GetUserUIPasswordInformationRequestInfo
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
			response, err := c.GetUserUIPasswordInformation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetWorkRequest", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     identity.GetWorkRequestRequest
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

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListApiKeys(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListApiKeys")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListApiKeys is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListApiKeys", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListApiKeys")
	assert.NoError(t, err)

	type ListApiKeysRequestInfo struct {
		ContainerId string
		Request     identity.ListApiKeysRequest
	}

	var requests []ListApiKeysRequestInfo
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
			response, err := c.ListApiKeys(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListAuthTokens(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListAuthTokens")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAuthTokens is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListAuthTokens", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListAuthTokens")
	assert.NoError(t, err)

	type ListAuthTokensRequestInfo struct {
		ContainerId string
		Request     identity.ListAuthTokensRequest
	}

	var requests []ListAuthTokensRequestInfo
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
			response, err := c.ListAuthTokens(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListAvailabilityDomains(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListAvailabilityDomains")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAvailabilityDomains is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListAvailabilityDomains", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListAvailabilityDomains")
	assert.NoError(t, err)

	type ListAvailabilityDomainsRequestInfo struct {
		ContainerId string
		Request     identity.ListAvailabilityDomainsRequest
	}

	var requests []ListAvailabilityDomainsRequestInfo
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
			response, err := c.ListAvailabilityDomains(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListBulkActionResourceTypes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListBulkActionResourceTypes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBulkActionResourceTypes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListBulkActionResourceTypes", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListBulkActionResourceTypes")
	assert.NoError(t, err)

	type ListBulkActionResourceTypesRequestInfo struct {
		ContainerId string
		Request     identity.ListBulkActionResourceTypesRequest
	}

	var requests []ListBulkActionResourceTypesRequestInfo
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
				r := req.(*identity.ListBulkActionResourceTypesRequest)
				return c.ListBulkActionResourceTypes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListBulkActionResourceTypesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListBulkActionResourceTypesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListBulkEditTagsResourceTypes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListBulkEditTagsResourceTypes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBulkEditTagsResourceTypes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListBulkEditTagsResourceTypes", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListBulkEditTagsResourceTypes")
	assert.NoError(t, err)

	type ListBulkEditTagsResourceTypesRequestInfo struct {
		ContainerId string
		Request     identity.ListBulkEditTagsResourceTypesRequest
	}

	var requests []ListBulkEditTagsResourceTypesRequestInfo
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
				r := req.(*identity.ListBulkEditTagsResourceTypesRequest)
				return c.ListBulkEditTagsResourceTypes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListBulkEditTagsResourceTypesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListBulkEditTagsResourceTypesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListCompartments(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListCompartments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCompartments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListCompartments", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListCompartments")
	assert.NoError(t, err)

	type ListCompartmentsRequestInfo struct {
		ContainerId string
		Request     identity.ListCompartmentsRequest
	}

	var requests []ListCompartmentsRequestInfo
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
				r := req.(*identity.ListCompartmentsRequest)
				return c.ListCompartments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListCompartmentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListCompartmentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListCostTrackingTags(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListCostTrackingTags")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCostTrackingTags is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListCostTrackingTags", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListCostTrackingTags")
	assert.NoError(t, err)

	type ListCostTrackingTagsRequestInfo struct {
		ContainerId string
		Request     identity.ListCostTrackingTagsRequest
	}

	var requests []ListCostTrackingTagsRequestInfo
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
				r := req.(*identity.ListCostTrackingTagsRequest)
				return c.ListCostTrackingTags(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListCostTrackingTagsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListCostTrackingTagsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListCustomerSecretKeys(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListCustomerSecretKeys")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCustomerSecretKeys is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListCustomerSecretKeys", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListCustomerSecretKeys")
	assert.NoError(t, err)

	type ListCustomerSecretKeysRequestInfo struct {
		ContainerId string
		Request     identity.ListCustomerSecretKeysRequest
	}

	var requests []ListCustomerSecretKeysRequestInfo
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
			response, err := c.ListCustomerSecretKeys(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListDynamicGroups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListDynamicGroups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDynamicGroups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListDynamicGroups", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListDynamicGroups")
	assert.NoError(t, err)

	type ListDynamicGroupsRequestInfo struct {
		ContainerId string
		Request     identity.ListDynamicGroupsRequest
	}

	var requests []ListDynamicGroupsRequestInfo
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
				r := req.(*identity.ListDynamicGroupsRequest)
				return c.ListDynamicGroups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListDynamicGroupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListDynamicGroupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListFaultDomains(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListFaultDomains")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListFaultDomains is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListFaultDomains", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListFaultDomains")
	assert.NoError(t, err)

	type ListFaultDomainsRequestInfo struct {
		ContainerId string
		Request     identity.ListFaultDomainsRequest
	}

	var requests []ListFaultDomainsRequestInfo
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
			response, err := c.ListFaultDomains(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListGroups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListGroups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListGroups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListGroups", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListGroups")
	assert.NoError(t, err)

	type ListGroupsRequestInfo struct {
		ContainerId string
		Request     identity.ListGroupsRequest
	}

	var requests []ListGroupsRequestInfo
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
				r := req.(*identity.ListGroupsRequest)
				return c.ListGroups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListGroupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListGroupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListIdentityProviderGroups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListIdentityProviderGroups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListIdentityProviderGroups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListIdentityProviderGroups", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListIdentityProviderGroups")
	assert.NoError(t, err)

	type ListIdentityProviderGroupsRequestInfo struct {
		ContainerId string
		Request     identity.ListIdentityProviderGroupsRequest
	}

	var requests []ListIdentityProviderGroupsRequestInfo
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
				r := req.(*identity.ListIdentityProviderGroupsRequest)
				return c.ListIdentityProviderGroups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListIdentityProviderGroupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListIdentityProviderGroupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListIdentityProviders(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListIdentityProviders")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListIdentityProviders is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListIdentityProviders", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListIdentityProviders")
	assert.NoError(t, err)

	type ListIdentityProvidersRequestInfo struct {
		ContainerId string
		Request     identity.ListIdentityProvidersRequest
	}

	var requests []ListIdentityProvidersRequestInfo
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
				r := req.(*identity.ListIdentityProvidersRequest)
				return c.ListIdentityProviders(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListIdentityProvidersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListIdentityProvidersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListIdpGroupMappings(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListIdpGroupMappings")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListIdpGroupMappings is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListIdpGroupMappings", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListIdpGroupMappings")
	assert.NoError(t, err)

	type ListIdpGroupMappingsRequestInfo struct {
		ContainerId string
		Request     identity.ListIdpGroupMappingsRequest
	}

	var requests []ListIdpGroupMappingsRequestInfo
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
				r := req.(*identity.ListIdpGroupMappingsRequest)
				return c.ListIdpGroupMappings(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListIdpGroupMappingsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListIdpGroupMappingsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListManagedCompartments(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListManagedCompartments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListManagedCompartments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListManagedCompartments", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListManagedCompartments")
	assert.NoError(t, err)

	type ListManagedCompartmentsRequestInfo struct {
		ContainerId string
		Request     identity.ListManagedCompartmentsRequest
	}

	var requests []ListManagedCompartmentsRequestInfo
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
				r := req.(*identity.ListManagedCompartmentsRequest)
				return c.ListManagedCompartments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListManagedCompartmentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListManagedCompartmentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListMfaTotpDevices(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListMfaTotpDevices")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListMfaTotpDevices is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListMfaTotpDevices", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListMfaTotpDevices")
	assert.NoError(t, err)

	type ListMfaTotpDevicesRequestInfo struct {
		ContainerId string
		Request     identity.ListMfaTotpDevicesRequest
	}

	var requests []ListMfaTotpDevicesRequestInfo
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
				r := req.(*identity.ListMfaTotpDevicesRequest)
				return c.ListMfaTotpDevices(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListMfaTotpDevicesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListMfaTotpDevicesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListNetworkSources(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListNetworkSources")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListNetworkSources is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListNetworkSources", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListNetworkSources")
	assert.NoError(t, err)

	type ListNetworkSourcesRequestInfo struct {
		ContainerId string
		Request     identity.ListNetworkSourcesRequest
	}

	var requests []ListNetworkSourcesRequestInfo
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
				r := req.(*identity.ListNetworkSourcesRequest)
				return c.ListNetworkSources(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListNetworkSourcesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListNetworkSourcesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListOAuthClientCredentials(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListOAuthClientCredentials")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListOAuthClientCredentials is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListOAuthClientCredentials", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListOAuthClientCredentials")
	assert.NoError(t, err)

	type ListOAuthClientCredentialsRequestInfo struct {
		ContainerId string
		Request     identity.ListOAuthClientCredentialsRequest
	}

	var requests []ListOAuthClientCredentialsRequestInfo
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
				r := req.(*identity.ListOAuthClientCredentialsRequest)
				return c.ListOAuthClientCredentials(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListOAuthClientCredentialsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListOAuthClientCredentialsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListPolicies(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListPolicies")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPolicies is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListPolicies", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListPolicies")
	assert.NoError(t, err)

	type ListPoliciesRequestInfo struct {
		ContainerId string
		Request     identity.ListPoliciesRequest
	}

	var requests []ListPoliciesRequestInfo
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
				r := req.(*identity.ListPoliciesRequest)
				return c.ListPolicies(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListPoliciesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListPoliciesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListRegionSubscriptions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListRegionSubscriptions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRegionSubscriptions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListRegionSubscriptions", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListRegionSubscriptions")
	assert.NoError(t, err)

	type ListRegionSubscriptionsRequestInfo struct {
		ContainerId string
		Request     identity.ListRegionSubscriptionsRequest
	}

	var requests []ListRegionSubscriptionsRequestInfo
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
			response, err := c.ListRegionSubscriptions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListRegions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListRegions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRegions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListRegions", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListRegions")
	assert.NoError(t, err)

	type ListRegionsRequestInfo struct {
		ContainerId string
		Request     interface{}
	}

	var requests []ListRegionsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {

			response, err := c.ListRegions(context.Background())
			message, err := testClient.validateResult(req.ContainerId, nil, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListSmtpCredentials(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListSmtpCredentials")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSmtpCredentials is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListSmtpCredentials", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListSmtpCredentials")
	assert.NoError(t, err)

	type ListSmtpCredentialsRequestInfo struct {
		ContainerId string
		Request     identity.ListSmtpCredentialsRequest
	}

	var requests []ListSmtpCredentialsRequestInfo
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
			response, err := c.ListSmtpCredentials(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListSwiftPasswords(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListSwiftPasswords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSwiftPasswords is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListSwiftPasswords", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListSwiftPasswords")
	assert.NoError(t, err)

	type ListSwiftPasswordsRequestInfo struct {
		ContainerId string
		Request     identity.ListSwiftPasswordsRequest
	}

	var requests []ListSwiftPasswordsRequestInfo
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
			response, err := c.ListSwiftPasswords(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListTagDefaults(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListTagDefaults")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTagDefaults is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListTagDefaults", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListTagDefaults")
	assert.NoError(t, err)

	type ListTagDefaultsRequestInfo struct {
		ContainerId string
		Request     identity.ListTagDefaultsRequest
	}

	var requests []ListTagDefaultsRequestInfo
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
				r := req.(*identity.ListTagDefaultsRequest)
				return c.ListTagDefaults(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListTagDefaultsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListTagDefaultsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListTagNamespaces(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListTagNamespaces")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTagNamespaces is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListTagNamespaces", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListTagNamespaces")
	assert.NoError(t, err)

	type ListTagNamespacesRequestInfo struct {
		ContainerId string
		Request     identity.ListTagNamespacesRequest
	}

	var requests []ListTagNamespacesRequestInfo
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
				r := req.(*identity.ListTagNamespacesRequest)
				return c.ListTagNamespaces(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListTagNamespacesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListTagNamespacesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListTagRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListTagRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTagRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListTagRules", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListTagRules")
	assert.NoError(t, err)

	type ListTagRulesRequestInfo struct {
		ContainerId string
		Request     identity.ListTagRulesRequest
	}

	var requests []ListTagRulesRequestInfo
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
				r := req.(*identity.ListTagRulesRequest)
				return c.ListTagRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListTagRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListTagRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListTaggingWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListTaggingWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTaggingWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListTaggingWorkRequestErrors", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListTaggingWorkRequestErrors")
	assert.NoError(t, err)

	type ListTaggingWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     identity.ListTaggingWorkRequestErrorsRequest
	}

	var requests []ListTaggingWorkRequestErrorsRequestInfo
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
				r := req.(*identity.ListTaggingWorkRequestErrorsRequest)
				return c.ListTaggingWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListTaggingWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListTaggingWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListTaggingWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListTaggingWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTaggingWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListTaggingWorkRequestLogs", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListTaggingWorkRequestLogs")
	assert.NoError(t, err)

	type ListTaggingWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     identity.ListTaggingWorkRequestLogsRequest
	}

	var requests []ListTaggingWorkRequestLogsRequestInfo
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
				r := req.(*identity.ListTaggingWorkRequestLogsRequest)
				return c.ListTaggingWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListTaggingWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListTaggingWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListTaggingWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListTaggingWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTaggingWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListTaggingWorkRequests", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListTaggingWorkRequests")
	assert.NoError(t, err)

	type ListTaggingWorkRequestsRequestInfo struct {
		ContainerId string
		Request     identity.ListTaggingWorkRequestsRequest
	}

	var requests []ListTaggingWorkRequestsRequestInfo
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
				r := req.(*identity.ListTaggingWorkRequestsRequest)
				return c.ListTaggingWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListTaggingWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListTaggingWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListTags(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListTags")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTags is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListTags", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListTags")
	assert.NoError(t, err)

	type ListTagsRequestInfo struct {
		ContainerId string
		Request     identity.ListTagsRequest
	}

	var requests []ListTagsRequestInfo
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
				r := req.(*identity.ListTagsRequest)
				return c.ListTags(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListTagsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListTagsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListTenancies(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListTenancies")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTenancies is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListTenancies", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListTenancies")
	assert.NoError(t, err)

	type ListTenanciesRequestInfo struct {
		ContainerId string
		Request     identity.ListTenanciesRequest
	}

	var requests []ListTenanciesRequestInfo
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
				r := req.(*identity.ListTenanciesRequest)
				return c.ListTenancies(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListTenanciesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListTenanciesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListUserGroupMemberships(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListUserGroupMemberships")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListUserGroupMemberships is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListUserGroupMemberships", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListUserGroupMemberships")
	assert.NoError(t, err)

	type ListUserGroupMembershipsRequestInfo struct {
		ContainerId string
		Request     identity.ListUserGroupMembershipsRequest
	}

	var requests []ListUserGroupMembershipsRequestInfo
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
				r := req.(*identity.ListUserGroupMembershipsRequest)
				return c.ListUserGroupMemberships(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListUserGroupMembershipsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListUserGroupMembershipsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListUsers(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListUsers")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListUsers is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListUsers", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListUsers")
	assert.NoError(t, err)

	type ListUsersRequestInfo struct {
		ContainerId string
		Request     identity.ListUsersRequest
	}

	var requests []ListUsersRequestInfo
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
				r := req.(*identity.ListUsersRequest)
				return c.ListUsers(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListUsersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListUsersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListWorkRequests", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     identity.ListWorkRequestsRequest
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
				r := req.(*identity.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientMoveCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "MoveCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("MoveCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "MoveCompartment", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "MoveCompartment")
	assert.NoError(t, err)

	type MoveCompartmentRequestInfo struct {
		ContainerId string
		Request     identity.MoveCompartmentRequest
	}

	var requests []MoveCompartmentRequestInfo
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
			response, err := c.MoveCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientRecoverCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "RecoverCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RecoverCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "RecoverCompartment", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "RecoverCompartment")
	assert.NoError(t, err)

	type RecoverCompartmentRequestInfo struct {
		ContainerId string
		Request     identity.RecoverCompartmentRequest
	}

	var requests []RecoverCompartmentRequestInfo
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
			response, err := c.RecoverCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientRemoveUserFromGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "RemoveUserFromGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RemoveUserFromGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "RemoveUserFromGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "RemoveUserFromGroup")
	assert.NoError(t, err)

	type RemoveUserFromGroupRequestInfo struct {
		ContainerId string
		Request     identity.RemoveUserFromGroupRequest
	}

	var requests []RemoveUserFromGroupRequestInfo
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
			response, err := c.RemoveUserFromGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientResetIdpScimClient(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "ResetIdpScimClient")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ResetIdpScimClient is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ResetIdpScimClient", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ResetIdpScimClient")
	assert.NoError(t, err)

	type ResetIdpScimClientRequestInfo struct {
		ContainerId string
		Request     identity.ResetIdpScimClientRequest
	}

	var requests []ResetIdpScimClientRequestInfo
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
			response, err := c.ResetIdpScimClient(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateAuthToken(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateAuthToken")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAuthToken is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateAuthToken", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateAuthToken")
	assert.NoError(t, err)

	type UpdateAuthTokenRequestInfo struct {
		ContainerId string
		Request     identity.UpdateAuthTokenRequest
	}

	var requests []UpdateAuthTokenRequestInfo
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
			response, err := c.UpdateAuthToken(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateAuthenticationPolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateAuthenticationPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAuthenticationPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateAuthenticationPolicy", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateAuthenticationPolicy")
	assert.NoError(t, err)

	type UpdateAuthenticationPolicyRequestInfo struct {
		ContainerId string
		Request     identity.UpdateAuthenticationPolicyRequest
	}

	var requests []UpdateAuthenticationPolicyRequestInfo
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
			response, err := c.UpdateAuthenticationPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateCompartment", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateCompartment")
	assert.NoError(t, err)

	type UpdateCompartmentRequestInfo struct {
		ContainerId string
		Request     identity.UpdateCompartmentRequest
	}

	var requests []UpdateCompartmentRequestInfo
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
			response, err := c.UpdateCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateCustomerSecretKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateCustomerSecretKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCustomerSecretKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateCustomerSecretKey", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateCustomerSecretKey")
	assert.NoError(t, err)

	type UpdateCustomerSecretKeyRequestInfo struct {
		ContainerId string
		Request     identity.UpdateCustomerSecretKeyRequest
	}

	var requests []UpdateCustomerSecretKeyRequestInfo
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
			response, err := c.UpdateCustomerSecretKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateDynamicGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateDynamicGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDynamicGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateDynamicGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateDynamicGroup")
	assert.NoError(t, err)

	type UpdateDynamicGroupRequestInfo struct {
		ContainerId string
		Request     identity.UpdateDynamicGroupRequest
	}

	var requests []UpdateDynamicGroupRequestInfo
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
			response, err := c.UpdateDynamicGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateGroup")
	assert.NoError(t, err)

	type UpdateGroupRequestInfo struct {
		ContainerId string
		Request     identity.UpdateGroupRequest
	}

	var requests []UpdateGroupRequestInfo
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
			response, err := c.UpdateGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateIdentityProvider(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateIdentityProvider")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateIdentityProvider is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateIdentityProvider", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateIdentityProvider")
	assert.NoError(t, err)

	type UpdateIdentityProviderRequestInfo struct {
		ContainerId string
		Request     identity.UpdateIdentityProviderRequest
	}

	var requests []UpdateIdentityProviderRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]UpdateIdentityProviderRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["UpdateIdentityProviderDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "protocol",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"SAML2": &identity.UpdateSaml2IdentityProviderDetails{},
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
			response, err := c.UpdateIdentityProvider(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateIdpGroupMapping(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateIdpGroupMapping")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateIdpGroupMapping is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateIdpGroupMapping", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateIdpGroupMapping")
	assert.NoError(t, err)

	type UpdateIdpGroupMappingRequestInfo struct {
		ContainerId string
		Request     identity.UpdateIdpGroupMappingRequest
	}

	var requests []UpdateIdpGroupMappingRequestInfo
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
			response, err := c.UpdateIdpGroupMapping(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateNetworkSource(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateNetworkSource")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateNetworkSource is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateNetworkSource", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateNetworkSource")
	assert.NoError(t, err)

	type UpdateNetworkSourceRequestInfo struct {
		ContainerId string
		Request     identity.UpdateNetworkSourceRequest
	}

	var requests []UpdateNetworkSourceRequestInfo
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
			response, err := c.UpdateNetworkSource(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateOAuthClientCredential(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateOAuthClientCredential")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateOAuthClientCredential is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateOAuthClientCredential", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateOAuthClientCredential")
	assert.NoError(t, err)

	type UpdateOAuthClientCredentialRequestInfo struct {
		ContainerId string
		Request     identity.UpdateOAuthClientCredentialRequest
	}

	var requests []UpdateOAuthClientCredentialRequestInfo
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
			response, err := c.UpdateOAuthClientCredential(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdatePolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdatePolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdatePolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdatePolicy", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdatePolicy")
	assert.NoError(t, err)

	type UpdatePolicyRequestInfo struct {
		ContainerId string
		Request     identity.UpdatePolicyRequest
	}

	var requests []UpdatePolicyRequestInfo
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
			response, err := c.UpdatePolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateSmtpCredential(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateSmtpCredential")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSmtpCredential is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateSmtpCredential", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateSmtpCredential")
	assert.NoError(t, err)

	type UpdateSmtpCredentialRequestInfo struct {
		ContainerId string
		Request     identity.UpdateSmtpCredentialRequest
	}

	var requests []UpdateSmtpCredentialRequestInfo
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
			response, err := c.UpdateSmtpCredential(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateSwiftPassword(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateSwiftPassword")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSwiftPassword is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateSwiftPassword", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateSwiftPassword")
	assert.NoError(t, err)

	type UpdateSwiftPasswordRequestInfo struct {
		ContainerId string
		Request     identity.UpdateSwiftPasswordRequest
	}

	var requests []UpdateSwiftPasswordRequestInfo
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
			response, err := c.UpdateSwiftPassword(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateTag(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateTag")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTag is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateTag", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateTag")
	assert.NoError(t, err)

	type UpdateTagRequestInfo struct {
		ContainerId string
		Request     identity.UpdateTagRequest
	}

	var requests []UpdateTagRequestInfo
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
			response, err := c.UpdateTag(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateTagDefault(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateTagDefault")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTagDefault is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateTagDefault", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateTagDefault")
	assert.NoError(t, err)

	type UpdateTagDefaultRequestInfo struct {
		ContainerId string
		Request     identity.UpdateTagDefaultRequest
	}

	var requests []UpdateTagDefaultRequestInfo
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
			response, err := c.UpdateTagDefault(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateTagNamespace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateTagNamespace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTagNamespace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateTagNamespace", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateTagNamespace")
	assert.NoError(t, err)

	type UpdateTagNamespaceRequestInfo struct {
		ContainerId string
		Request     identity.UpdateTagNamespaceRequest
	}

	var requests []UpdateTagNamespaceRequestInfo
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
			response, err := c.UpdateTagNamespace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateTagRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateTagRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTagRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateTagRule", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateTagRule")
	assert.NoError(t, err)

	type UpdateTagRuleRequestInfo struct {
		ContainerId string
		Request     identity.UpdateTagRuleRequest
	}

	var requests []UpdateTagRuleRequestInfo
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
			response, err := c.UpdateTagRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateUser(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateUser")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateUser is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateUser", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateUser")
	assert.NoError(t, err)

	type UpdateUserRequestInfo struct {
		ContainerId string
		Request     identity.UpdateUserRequest
	}

	var requests []UpdateUserRequestInfo
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
			response, err := c.UpdateUser(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateUserCapabilities(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateUserCapabilities")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateUserCapabilities is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateUserCapabilities", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateUserCapabilities")
	assert.NoError(t, err)

	type UpdateUserCapabilitiesRequestInfo struct {
		ContainerId string
		Request     identity.UpdateUserCapabilitiesRequest
	}

	var requests []UpdateUserCapabilitiesRequestInfo
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
			response, err := c.UpdateUserCapabilities(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUpdateUserState(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UpdateUserState")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateUserState is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateUserState", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateUserState")
	assert.NoError(t, err)

	type UpdateUserStateRequestInfo struct {
		ContainerId string
		Request     identity.UpdateUserStateRequest
	}

	var requests []UpdateUserStateRequestInfo
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
			response, err := c.UpdateUserState(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_identity_team_us_grp@oracle.com" jiraProject="ID" opsJiraProject="ID"
func TestIdentityClientUploadApiKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("identity", "UploadApiKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UploadApiKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UploadApiKey", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UploadApiKey")
	assert.NoError(t, err)

	type UploadApiKeyRequestInfo struct {
		ContainerId string
		Request     identity.UploadApiKeyRequest
	}

	var requests []UploadApiKeyRequestInfo
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
			response, err := c.UploadApiKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
