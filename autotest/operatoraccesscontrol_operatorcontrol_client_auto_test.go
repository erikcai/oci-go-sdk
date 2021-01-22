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

func createOperatoraccesscontrolOperatorControlClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := operatoraccesscontrol.NewOperatorControlClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolOperatorControlClientChangeOperatorControlCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "ChangeOperatorControlCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeOperatorControlCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "OperatorControl", "ChangeOperatorControlCompartment", createOperatoraccesscontrolOperatorControlClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.OperatorControlClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "ChangeOperatorControlCompartment")
	assert.NoError(t, err)

	type ChangeOperatorControlCompartmentRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.ChangeOperatorControlCompartmentRequest
	}

	var requests []ChangeOperatorControlCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeOperatorControlCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolOperatorControlClientCreateOperatorControl(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "CreateOperatorControl")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateOperatorControl is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "OperatorControl", "CreateOperatorControl", createOperatoraccesscontrolOperatorControlClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.OperatorControlClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "CreateOperatorControl")
	assert.NoError(t, err)

	type CreateOperatorControlRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.CreateOperatorControlRequest
	}

	var requests []CreateOperatorControlRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateOperatorControl(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolOperatorControlClientDeleteOperatorControl(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "DeleteOperatorControl")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteOperatorControl is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "OperatorControl", "DeleteOperatorControl", createOperatoraccesscontrolOperatorControlClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.OperatorControlClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "DeleteOperatorControl")
	assert.NoError(t, err)

	type DeleteOperatorControlRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.DeleteOperatorControlRequest
	}

	var requests []DeleteOperatorControlRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteOperatorControl(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolOperatorControlClientGetOperatorControl(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "GetOperatorControl")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetOperatorControl is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "OperatorControl", "GetOperatorControl", createOperatoraccesscontrolOperatorControlClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.OperatorControlClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "GetOperatorControl")
	assert.NoError(t, err)

	type GetOperatorControlRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.GetOperatorControlRequest
	}

	var requests []GetOperatorControlRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetOperatorControl(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolOperatorControlClientListOperatorControls(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "ListOperatorControls")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListOperatorControls is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "OperatorControl", "ListOperatorControls", createOperatoraccesscontrolOperatorControlClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.OperatorControlClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "ListOperatorControls")
	assert.NoError(t, err)

	type ListOperatorControlsRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.ListOperatorControlsRequest
	}

	var requests []ListOperatorControlsRequestInfo
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
				r := req.(*operatoraccesscontrol.ListOperatorControlsRequest)
				return c.ListOperatorControls(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operatoraccesscontrol.ListOperatorControlsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operatoraccesscontrol.ListOperatorControlsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolOperatorControlClientUpdateOperatorControl(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "UpdateOperatorControl")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateOperatorControl is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "OperatorControl", "UpdateOperatorControl", createOperatoraccesscontrolOperatorControlClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.OperatorControlClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "UpdateOperatorControl")
	assert.NoError(t, err)

	type UpdateOperatorControlRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.UpdateOperatorControlRequest
	}

	var requests []UpdateOperatorControlRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateOperatorControl(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
