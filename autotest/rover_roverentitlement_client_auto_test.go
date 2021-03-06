package autotest

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"github.com/erikcai/oci-go-sdk/v33/rover"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createRoverRoverEntitlementClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := rover.NewRoverEntitlementClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverEntitlementClientChangeRoverEntitlementCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "ChangeRoverEntitlementCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeRoverEntitlementCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverEntitlement", "ChangeRoverEntitlementCompartment", createRoverRoverEntitlementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverEntitlementClient)

	body, err := testClient.getRequests("rover", "ChangeRoverEntitlementCompartment")
	assert.NoError(t, err)

	type ChangeRoverEntitlementCompartmentRequestInfo struct {
		ContainerId string
		Request     rover.ChangeRoverEntitlementCompartmentRequest
	}

	var requests []ChangeRoverEntitlementCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeRoverEntitlementCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverEntitlementClientCreateRoverEntitlement(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "CreateRoverEntitlement")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateRoverEntitlement is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverEntitlement", "CreateRoverEntitlement", createRoverRoverEntitlementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverEntitlementClient)

	body, err := testClient.getRequests("rover", "CreateRoverEntitlement")
	assert.NoError(t, err)

	type CreateRoverEntitlementRequestInfo struct {
		ContainerId string
		Request     rover.CreateRoverEntitlementRequest
	}

	var requests []CreateRoverEntitlementRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateRoverEntitlement(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverEntitlementClientDeleteRoverEntitlement(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "DeleteRoverEntitlement")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteRoverEntitlement is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverEntitlement", "DeleteRoverEntitlement", createRoverRoverEntitlementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverEntitlementClient)

	body, err := testClient.getRequests("rover", "DeleteRoverEntitlement")
	assert.NoError(t, err)

	type DeleteRoverEntitlementRequestInfo struct {
		ContainerId string
		Request     rover.DeleteRoverEntitlementRequest
	}

	var requests []DeleteRoverEntitlementRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteRoverEntitlement(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverEntitlementClientGetRoverEntitlement(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "GetRoverEntitlement")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRoverEntitlement is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverEntitlement", "GetRoverEntitlement", createRoverRoverEntitlementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverEntitlementClient)

	body, err := testClient.getRequests("rover", "GetRoverEntitlement")
	assert.NoError(t, err)

	type GetRoverEntitlementRequestInfo struct {
		ContainerId string
		Request     rover.GetRoverEntitlementRequest
	}

	var requests []GetRoverEntitlementRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetRoverEntitlement(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverEntitlementClientListRoverEntitlements(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "ListRoverEntitlements")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRoverEntitlements is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverEntitlement", "ListRoverEntitlements", createRoverRoverEntitlementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverEntitlementClient)

	body, err := testClient.getRequests("rover", "ListRoverEntitlements")
	assert.NoError(t, err)

	type ListRoverEntitlementsRequestInfo struct {
		ContainerId string
		Request     rover.ListRoverEntitlementsRequest
	}

	var requests []ListRoverEntitlementsRequestInfo
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
				r := req.(*rover.ListRoverEntitlementsRequest)
				return c.ListRoverEntitlements(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]rover.ListRoverEntitlementsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(rover.ListRoverEntitlementsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverEntitlementClientUpdateRoverEntitlement(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "UpdateRoverEntitlement")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRoverEntitlement is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverEntitlement", "UpdateRoverEntitlement", createRoverRoverEntitlementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverEntitlementClient)

	body, err := testClient.getRequests("rover", "UpdateRoverEntitlement")
	assert.NoError(t, err)

	type UpdateRoverEntitlementRequestInfo struct {
		ContainerId string
		Request     rover.UpdateRoverEntitlementRequest
	}

	var requests []UpdateRoverEntitlementRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateRoverEntitlement(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
