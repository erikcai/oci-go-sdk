package autotest

import (
	"github.com/oracle/oci-go-sdk/v30/common"
	"github.com/oracle/oci-go-sdk/v30/tenantmanagercontrolplane"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createTenantmanagercontrolplaneRecipientInvitationClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := tenantmanagercontrolplane.NewRecipientInvitationClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="acc_customer_tools_us_grp@oracle.com" jiraProject="ACCMGMT" opsJiraProject="ACCMGMT"
func TestTenantmanagercontrolplaneRecipientInvitationClientAcceptRecipientInvitation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("tenantmanagercontrolplane", "AcceptRecipientInvitation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AcceptRecipientInvitation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("tenantmanagercontrolplane", "RecipientInvitation", "AcceptRecipientInvitation", createTenantmanagercontrolplaneRecipientInvitationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(tenantmanagercontrolplane.RecipientInvitationClient)

	body, err := testClient.getRequests("tenantmanagercontrolplane", "AcceptRecipientInvitation")
	assert.NoError(t, err)

	type AcceptRecipientInvitationRequestInfo struct {
		ContainerId string
		Request     tenantmanagercontrolplane.AcceptRecipientInvitationRequest
	}

	var requests []AcceptRecipientInvitationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.AcceptRecipientInvitation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="acc_customer_tools_us_grp@oracle.com" jiraProject="ACCMGMT" opsJiraProject="ACCMGMT"
func TestTenantmanagercontrolplaneRecipientInvitationClientGetRecipientInvitation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("tenantmanagercontrolplane", "GetRecipientInvitation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRecipientInvitation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("tenantmanagercontrolplane", "RecipientInvitation", "GetRecipientInvitation", createTenantmanagercontrolplaneRecipientInvitationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(tenantmanagercontrolplane.RecipientInvitationClient)

	body, err := testClient.getRequests("tenantmanagercontrolplane", "GetRecipientInvitation")
	assert.NoError(t, err)

	type GetRecipientInvitationRequestInfo struct {
		ContainerId string
		Request     tenantmanagercontrolplane.GetRecipientInvitationRequest
	}

	var requests []GetRecipientInvitationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetRecipientInvitation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="acc_customer_tools_us_grp@oracle.com" jiraProject="ACCMGMT" opsJiraProject="ACCMGMT"
func TestTenantmanagercontrolplaneRecipientInvitationClientIgnoreRecipientInvitation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("tenantmanagercontrolplane", "IgnoreRecipientInvitation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("IgnoreRecipientInvitation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("tenantmanagercontrolplane", "RecipientInvitation", "IgnoreRecipientInvitation", createTenantmanagercontrolplaneRecipientInvitationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(tenantmanagercontrolplane.RecipientInvitationClient)

	body, err := testClient.getRequests("tenantmanagercontrolplane", "IgnoreRecipientInvitation")
	assert.NoError(t, err)

	type IgnoreRecipientInvitationRequestInfo struct {
		ContainerId string
		Request     tenantmanagercontrolplane.IgnoreRecipientInvitationRequest
	}

	var requests []IgnoreRecipientInvitationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.IgnoreRecipientInvitation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="acc_customer_tools_us_grp@oracle.com" jiraProject="ACCMGMT" opsJiraProject="ACCMGMT"
func TestTenantmanagercontrolplaneRecipientInvitationClientListRecipientInvitations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("tenantmanagercontrolplane", "ListRecipientInvitations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRecipientInvitations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("tenantmanagercontrolplane", "RecipientInvitation", "ListRecipientInvitations", createTenantmanagercontrolplaneRecipientInvitationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(tenantmanagercontrolplane.RecipientInvitationClient)

	body, err := testClient.getRequests("tenantmanagercontrolplane", "ListRecipientInvitations")
	assert.NoError(t, err)

	type ListRecipientInvitationsRequestInfo struct {
		ContainerId string
		Request     tenantmanagercontrolplane.ListRecipientInvitationsRequest
	}

	var requests []ListRecipientInvitationsRequestInfo
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
				r := req.(*tenantmanagercontrolplane.ListRecipientInvitationsRequest)
				return c.ListRecipientInvitations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]tenantmanagercontrolplane.ListRecipientInvitationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(tenantmanagercontrolplane.ListRecipientInvitationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="acc_customer_tools_us_grp@oracle.com" jiraProject="ACCMGMT" opsJiraProject="ACCMGMT"
func TestTenantmanagercontrolplaneRecipientInvitationClientUpdateRecipientInvitation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("tenantmanagercontrolplane", "UpdateRecipientInvitation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRecipientInvitation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("tenantmanagercontrolplane", "RecipientInvitation", "UpdateRecipientInvitation", createTenantmanagercontrolplaneRecipientInvitationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(tenantmanagercontrolplane.RecipientInvitationClient)

	body, err := testClient.getRequests("tenantmanagercontrolplane", "UpdateRecipientInvitation")
	assert.NoError(t, err)

	type UpdateRecipientInvitationRequestInfo struct {
		ContainerId string
		Request     tenantmanagercontrolplane.UpdateRecipientInvitationRequest
	}

	var requests []UpdateRecipientInvitationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateRecipientInvitation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
