package autotest

import (
	"github.com/oracle/oci-go-sdk/v26/common"
	"github.com/oracle/oci-go-sdk/v26/operatoraccesscontrol"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createOperatoraccesscontrolOperatorActionsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := operatoraccesscontrol.NewOperatorActionsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolOperatorActionsClientGetOperatorAction(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "GetOperatorAction")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetOperatorAction is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "OperatorActions", "GetOperatorAction", createOperatoraccesscontrolOperatorActionsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.OperatorActionsClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "GetOperatorAction")
	assert.NoError(t, err)

	type GetOperatorActionRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.GetOperatorActionRequest
	}

	var requests []GetOperatorActionRequestInfo
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
			response, err := c.GetOperatorAction(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="adbd_admins_us_grp@oracle.com" jiraProject="CCA" opsJiraProject="OPCTL"
func TestOperatoraccesscontrolOperatorActionsClientListOperatorActions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("operatoraccesscontrol", "ListOperatorActions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListOperatorActions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("operatoraccesscontrol", "OperatorActions", "ListOperatorActions", createOperatoraccesscontrolOperatorActionsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(operatoraccesscontrol.OperatorActionsClient)

	body, err := testClient.getRequests("operatoraccesscontrol", "ListOperatorActions")
	assert.NoError(t, err)

	type ListOperatorActionsRequestInfo struct {
		ContainerId string
		Request     operatoraccesscontrol.ListOperatorActionsRequest
	}

	var requests []ListOperatorActionsRequestInfo
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
				r := req.(*operatoraccesscontrol.ListOperatorActionsRequest)
				return c.ListOperatorActions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]operatoraccesscontrol.ListOperatorActionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(operatoraccesscontrol.ListOperatorActionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
