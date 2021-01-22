package autotest

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"github.com/erikcai/oci-go-sdk/v33/operatoraccesscontrol"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createOperatoraccesscontrolOperatorControlAssignmentClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := operatoraccesscontrol.NewOperatorControlAssignmentClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolOperatorControlAssignmentClientChangeOperatorControlAssignmentCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "ChangeOperatorControlAssignmentCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeOperatorControlAssignmentCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "OperatorControlAssignment", "ChangeOperatorControlAssignmentCompartment", createOperatoraccesscontrolOperatorControlAssignmentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.OperatorControlAssignmentClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "ChangeOperatorControlAssignmentCompartment")
	assert.NoError(t, err)

	type ChangeOperatorControlAssignmentCompartmentRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.ChangeOperatorControlAssignmentCompartmentRequest
	}

	var requests []ChangeOperatorControlAssignmentCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeOperatorControlAssignmentCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolOperatorControlAssignmentClientCreateOperatorControlAssignment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "CreateOperatorControlAssignment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateOperatorControlAssignment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "OperatorControlAssignment", "CreateOperatorControlAssignment", createOperatoraccesscontrolOperatorControlAssignmentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.OperatorControlAssignmentClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "CreateOperatorControlAssignment")
	assert.NoError(t, err)

	type CreateOperatorControlAssignmentRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.CreateOperatorControlAssignmentRequest
	}

	var requests []CreateOperatorControlAssignmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateOperatorControlAssignment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolOperatorControlAssignmentClientDeleteOperatorControlAssignment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "DeleteOperatorControlAssignment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteOperatorControlAssignment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "OperatorControlAssignment", "DeleteOperatorControlAssignment", createOperatoraccesscontrolOperatorControlAssignmentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.OperatorControlAssignmentClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "DeleteOperatorControlAssignment")
	assert.NoError(t, err)

	type DeleteOperatorControlAssignmentRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.DeleteOperatorControlAssignmentRequest
	}

	var requests []DeleteOperatorControlAssignmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteOperatorControlAssignment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolOperatorControlAssignmentClientGetOperatorControlAssignment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "GetOperatorControlAssignment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetOperatorControlAssignment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "OperatorControlAssignment", "GetOperatorControlAssignment", createOperatoraccesscontrolOperatorControlAssignmentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.OperatorControlAssignmentClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "GetOperatorControlAssignment")
	assert.NoError(t, err)

	type GetOperatorControlAssignmentRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.GetOperatorControlAssignmentRequest
	}

	var requests []GetOperatorControlAssignmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetOperatorControlAssignment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolOperatorControlAssignmentClientListOperatorControlAssignments(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "ListOperatorControlAssignments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListOperatorControlAssignments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "OperatorControlAssignment", "ListOperatorControlAssignments", createOperatoraccesscontrolOperatorControlAssignmentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.OperatorControlAssignmentClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "ListOperatorControlAssignments")
	assert.NoError(t, err)

	type ListOperatorControlAssignmentsRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.ListOperatorControlAssignmentsRequest
	}

	var requests []ListOperatorControlAssignmentsRequestInfo
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
				r := req.(*operatoraccesscontrol.ListOperatorControlAssignmentsRequest)
				return c.ListOperatorControlAssignments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operatoraccesscontrol.ListOperatorControlAssignmentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operatoraccesscontrol.ListOperatorControlAssignmentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolOperatorControlAssignmentClientUpdateOperatorControlAssignment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "UpdateOperatorControlAssignment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateOperatorControlAssignment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "OperatorControlAssignment", "UpdateOperatorControlAssignment", createOperatoraccesscontrolOperatorControlAssignmentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.OperatorControlAssignmentClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "UpdateOperatorControlAssignment")
	assert.NoError(t, err)

	type UpdateOperatorControlAssignmentRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.UpdateOperatorControlAssignmentRequest
	}

	var requests []UpdateOperatorControlAssignmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateOperatorControlAssignment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
