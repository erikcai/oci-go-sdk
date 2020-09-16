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

func createOptimizerRecommendationClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := optimizer.NewRecommendationClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerRecommendationClientBulkApplyRecommendations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "BulkApplyRecommendations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("BulkApplyRecommendations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Recommendation", "BulkApplyRecommendations", createOptimizerRecommendationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.RecommendationClient)

	body, err := testClient.getRequests("optimizer", "BulkApplyRecommendations")
	assert.NoError(t, err)

	type BulkApplyRecommendationsRequestInfo struct {
		ContainerId string
		Request     optimizer.BulkApplyRecommendationsRequest
	}

	var requests []BulkApplyRecommendationsRequestInfo
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
			response, err := c.BulkApplyRecommendations(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerRecommendationClientGetRecommendation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "GetRecommendation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRecommendation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Recommendation", "GetRecommendation", createOptimizerRecommendationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.RecommendationClient)

	body, err := testClient.getRequests("optimizer", "GetRecommendation")
	assert.NoError(t, err)

	type GetRecommendationRequestInfo struct {
		ContainerId string
		Request     optimizer.GetRecommendationRequest
	}

	var requests []GetRecommendationRequestInfo
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
			response, err := c.GetRecommendation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerRecommendationClientListRecommendations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "ListRecommendations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRecommendations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Recommendation", "ListRecommendations", createOptimizerRecommendationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.RecommendationClient)

	body, err := testClient.getRequests("optimizer", "ListRecommendations")
	assert.NoError(t, err)

	type ListRecommendationsRequestInfo struct {
		ContainerId string
		Request     optimizer.ListRecommendationsRequest
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
				r := req.(*optimizer.ListRecommendationsRequest)
				return c.ListRecommendations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]optimizer.ListRecommendationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(optimizer.ListRecommendationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerRecommendationClientUpdateRecommendation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "UpdateRecommendation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRecommendation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Recommendation", "UpdateRecommendation", createOptimizerRecommendationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.RecommendationClient)

	body, err := testClient.getRequests("optimizer", "UpdateRecommendation")
	assert.NoError(t, err)

	type UpdateRecommendationRequestInfo struct {
		ContainerId string
		Request     optimizer.UpdateRecommendationRequest
	}

	var requests []UpdateRecommendationRequestInfo
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
			response, err := c.UpdateRecommendation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
