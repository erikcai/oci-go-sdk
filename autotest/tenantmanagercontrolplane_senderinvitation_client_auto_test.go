package autotest

import (
	"github.com/oracle/oci-go-sdk/v28/common"
	"github.com/oracle/oci-go-sdk/v28/tenantmanagercontrolplane"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createTenantmanagercontrolplaneSenderInvitationClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := tenantmanagercontrolplane.NewSenderInvitationClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="acc_customer_tools_us_grp@oracle.com" jiraProject="ACCMGMT" opsJiraProject="ACCMGMT"
func TestTenantmanagercontrolplaneSenderInvitationClientCancelSenderInvitation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("tenantmanagercontrolplane", "CancelSenderInvitation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CancelSenderInvitation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("tenantmanagercontrolplane", "SenderInvitation", "CancelSenderInvitation", createTenantmanagercontrolplaneSenderInvitationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(tenantmanagercontrolplane.SenderInvitationClient)

	body, err := testClient.getRequests("tenantmanagercontrolplane", "CancelSenderInvitation")
	assert.NoError(t, err)

	type CancelSenderInvitationRequestInfo struct {
		ContainerId string
		Request     tenantmanagercontrolplane.CancelSenderInvitationRequest
	}

	var requests []CancelSenderInvitationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CancelSenderInvitation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="acc_customer_tools_us_grp@oracle.com" jiraProject="ACCMGMT" opsJiraProject="ACCMGMT"
func TestTenantmanagercontrolplaneSenderInvitationClientCreateSenderInvitation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("tenantmanagercontrolplane", "CreateSenderInvitation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSenderInvitation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("tenantmanagercontrolplane", "SenderInvitation", "CreateSenderInvitation", createTenantmanagercontrolplaneSenderInvitationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(tenantmanagercontrolplane.SenderInvitationClient)

	body, err := testClient.getRequests("tenantmanagercontrolplane", "CreateSenderInvitation")
	assert.NoError(t, err)

	type CreateSenderInvitationRequestInfo struct {
		ContainerId string
		Request     tenantmanagercontrolplane.CreateSenderInvitationRequest
	}

	var requests []CreateSenderInvitationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateSenderInvitation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="acc_customer_tools_us_grp@oracle.com" jiraProject="ACCMGMT" opsJiraProject="ACCMGMT"
func TestTenantmanagercontrolplaneSenderInvitationClientGetSenderInvitation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("tenantmanagercontrolplane", "GetSenderInvitation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSenderInvitation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("tenantmanagercontrolplane", "SenderInvitation", "GetSenderInvitation", createTenantmanagercontrolplaneSenderInvitationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(tenantmanagercontrolplane.SenderInvitationClient)

	body, err := testClient.getRequests("tenantmanagercontrolplane", "GetSenderInvitation")
	assert.NoError(t, err)

	type GetSenderInvitationRequestInfo struct {
		ContainerId string
		Request     tenantmanagercontrolplane.GetSenderInvitationRequest
	}

	var requests []GetSenderInvitationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetSenderInvitation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="acc_customer_tools_us_grp@oracle.com" jiraProject="ACCMGMT" opsJiraProject="ACCMGMT"
func TestTenantmanagercontrolplaneSenderInvitationClientListSenderInvitations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("tenantmanagercontrolplane", "ListSenderInvitations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSenderInvitations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("tenantmanagercontrolplane", "SenderInvitation", "ListSenderInvitations", createTenantmanagercontrolplaneSenderInvitationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(tenantmanagercontrolplane.SenderInvitationClient)

	body, err := testClient.getRequests("tenantmanagercontrolplane", "ListSenderInvitations")
	assert.NoError(t, err)

	type ListSenderInvitationsRequestInfo struct {
		ContainerId string
		Request     tenantmanagercontrolplane.ListSenderInvitationsRequest
	}

	var requests []ListSenderInvitationsRequestInfo
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
				r := req.(*tenantmanagercontrolplane.ListSenderInvitationsRequest)
				return c.ListSenderInvitations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]tenantmanagercontrolplane.ListSenderInvitationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(tenantmanagercontrolplane.ListSenderInvitationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="acc_customer_tools_us_grp@oracle.com" jiraProject="ACCMGMT" opsJiraProject="ACCMGMT"
func TestTenantmanagercontrolplaneSenderInvitationClientUpdateSenderInvitation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("tenantmanagercontrolplane", "UpdateSenderInvitation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSenderInvitation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("tenantmanagercontrolplane", "SenderInvitation", "UpdateSenderInvitation", createTenantmanagercontrolplaneSenderInvitationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(tenantmanagercontrolplane.SenderInvitationClient)

	body, err := testClient.getRequests("tenantmanagercontrolplane", "UpdateSenderInvitation")
	assert.NoError(t, err)

	type UpdateSenderInvitationRequestInfo struct {
		ContainerId string
		Request     tenantmanagercontrolplane.UpdateSenderInvitationRequest
	}

	var requests []UpdateSenderInvitationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateSenderInvitation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
