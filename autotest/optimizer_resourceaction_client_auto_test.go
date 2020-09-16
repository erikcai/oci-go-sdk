package autotest

import (
	"github.com/oracle/oci-go-sdk/v25/common"
	"github.com/oracle/oci-go-sdk/v25/optimizer"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createOptimizerResourceActionClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := optimizer.NewResourceActionClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerResourceActionClientGetResourceAction(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "GetResourceAction")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetResourceAction is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "ResourceAction", "GetResourceAction", createOptimizerResourceActionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.ResourceActionClient)

	body, err := testClient.getRequests("optimizer", "GetResourceAction")
	assert.NoError(t, err)

	type GetResourceActionRequestInfo struct {
		ContainerId string
		Request     optimizer.GetResourceActionRequest
	}

	var requests []GetResourceActionRequestInfo
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
			response, err := c.GetResourceAction(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerResourceActionClientListResourceActions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "ListResourceActions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListResourceActions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "ResourceAction", "ListResourceActions", createOptimizerResourceActionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.ResourceActionClient)

	body, err := testClient.getRequests("optimizer", "ListResourceActions")
	assert.NoError(t, err)

	type ListResourceActionsRequestInfo struct {
		ContainerId string
		Request     optimizer.ListResourceActionsRequest
	}

	var requests []ListResourceActionsRequestInfo
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
				r := req.(*optimizer.ListResourceActionsRequest)
				return c.ListResourceActions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]optimizer.ListResourceActionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(optimizer.ListResourceActionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerResourceActionClientUpdateResourceAction(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "UpdateResourceAction")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateResourceAction is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "ResourceAction", "UpdateResourceAction", createOptimizerResourceActionClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.ResourceActionClient)

	body, err := testClient.getRequests("optimizer", "UpdateResourceAction")
	assert.NoError(t, err)

	type UpdateResourceActionRequestInfo struct {
		ContainerId string
		Request     optimizer.UpdateResourceActionRequest
	}

	var requests []UpdateResourceActionRequestInfo
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
			response, err := c.UpdateResourceAction(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
