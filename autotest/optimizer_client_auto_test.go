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

func createOptimizerOptimizerClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := optimizer.NewOptimizerClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientBulkApplyRecommendations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "BulkApplyRecommendations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("BulkApplyRecommendations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "BulkApplyRecommendations", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

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
func TestOptimizerOptimizerClientCreateProfile(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "CreateProfile")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateProfile is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "CreateProfile", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "CreateProfile")
	assert.NoError(t, err)

	type CreateProfileRequestInfo struct {
		ContainerId string
		Request     optimizer.CreateProfileRequest
	}

	var requests []CreateProfileRequestInfo
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
			response, err := c.CreateProfile(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientDeleteProfile(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "DeleteProfile")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteProfile is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "DeleteProfile", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "DeleteProfile")
	assert.NoError(t, err)

	type DeleteProfileRequestInfo struct {
		ContainerId string
		Request     optimizer.DeleteProfileRequest
	}

	var requests []DeleteProfileRequestInfo
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
			response, err := c.DeleteProfile(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientGetCategory(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "GetCategory")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCategory is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "GetCategory", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "GetCategory")
	assert.NoError(t, err)

	type GetCategoryRequestInfo struct {
		ContainerId string
		Request     optimizer.GetCategoryRequest
	}

	var requests []GetCategoryRequestInfo
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
			response, err := c.GetCategory(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientGetEnrollmentStatus(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "GetEnrollmentStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetEnrollmentStatus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "GetEnrollmentStatus", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "GetEnrollmentStatus")
	assert.NoError(t, err)

	type GetEnrollmentStatusRequestInfo struct {
		ContainerId string
		Request     optimizer.GetEnrollmentStatusRequest
	}

	var requests []GetEnrollmentStatusRequestInfo
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
			response, err := c.GetEnrollmentStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientGetProfile(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "GetProfile")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetProfile is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "GetProfile", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "GetProfile")
	assert.NoError(t, err)

	type GetProfileRequestInfo struct {
		ContainerId string
		Request     optimizer.GetProfileRequest
	}

	var requests []GetProfileRequestInfo
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
			response, err := c.GetProfile(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientGetRecommendation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "GetRecommendation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRecommendation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "GetRecommendation", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

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
func TestOptimizerOptimizerClientGetResourceAction(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "GetResourceAction")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetResourceAction is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "GetResourceAction", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

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
func TestOptimizerOptimizerClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "GetWorkRequest", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     optimizer.GetWorkRequestRequest
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

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientListCategories(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "ListCategories")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCategories is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "ListCategories", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "ListCategories")
	assert.NoError(t, err)

	type ListCategoriesRequestInfo struct {
		ContainerId string
		Request     optimizer.ListCategoriesRequest
	}

	var requests []ListCategoriesRequestInfo
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
				r := req.(*optimizer.ListCategoriesRequest)
				return c.ListCategories(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]optimizer.ListCategoriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(optimizer.ListCategoriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientListEnrollmentStatuses(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "ListEnrollmentStatuses")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListEnrollmentStatuses is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "ListEnrollmentStatuses", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "ListEnrollmentStatuses")
	assert.NoError(t, err)

	type ListEnrollmentStatusesRequestInfo struct {
		ContainerId string
		Request     optimizer.ListEnrollmentStatusesRequest
	}

	var requests []ListEnrollmentStatusesRequestInfo
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
				r := req.(*optimizer.ListEnrollmentStatusesRequest)
				return c.ListEnrollmentStatuses(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]optimizer.ListEnrollmentStatusesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(optimizer.ListEnrollmentStatusesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientListHistories(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "ListHistories")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListHistories is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "ListHistories", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "ListHistories")
	assert.NoError(t, err)

	type ListHistoriesRequestInfo struct {
		ContainerId string
		Request     optimizer.ListHistoriesRequest
	}

	var requests []ListHistoriesRequestInfo
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
				r := req.(*optimizer.ListHistoriesRequest)
				return c.ListHistories(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]optimizer.ListHistoriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(optimizer.ListHistoriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientListProfiles(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "ListProfiles")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListProfiles is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "ListProfiles", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "ListProfiles")
	assert.NoError(t, err)

	type ListProfilesRequestInfo struct {
		ContainerId string
		Request     optimizer.ListProfilesRequest
	}

	var requests []ListProfilesRequestInfo
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
				r := req.(*optimizer.ListProfilesRequest)
				return c.ListProfiles(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]optimizer.ListProfilesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(optimizer.ListProfilesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientListRecommendations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "ListRecommendations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRecommendations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "ListRecommendations", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

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
func TestOptimizerOptimizerClientListResourceActions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "ListResourceActions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListResourceActions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "ListResourceActions", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

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
func TestOptimizerOptimizerClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "ListWorkRequestErrors", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     optimizer.ListWorkRequestErrorsRequest
	}

	var requests []ListWorkRequestErrorsRequestInfo
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
				r := req.(*optimizer.ListWorkRequestErrorsRequest)
				return c.ListWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]optimizer.ListWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(optimizer.ListWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "ListWorkRequestLogs", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     optimizer.ListWorkRequestLogsRequest
	}

	var requests []ListWorkRequestLogsRequestInfo
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
				r := req.(*optimizer.ListWorkRequestLogsRequest)
				return c.ListWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]optimizer.ListWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(optimizer.ListWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "ListWorkRequests", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     optimizer.ListWorkRequestsRequest
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
				r := req.(*optimizer.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]optimizer.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(optimizer.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientUpdateEnrollmentStatus(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "UpdateEnrollmentStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateEnrollmentStatus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "UpdateEnrollmentStatus", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "UpdateEnrollmentStatus")
	assert.NoError(t, err)

	type UpdateEnrollmentStatusRequestInfo struct {
		ContainerId string
		Request     optimizer.UpdateEnrollmentStatusRequest
	}

	var requests []UpdateEnrollmentStatusRequestInfo
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
			response, err := c.UpdateEnrollmentStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientUpdateProfile(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "UpdateProfile")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateProfile is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "UpdateProfile", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

	body, err := testClient.getRequests("optimizer", "UpdateProfile")
	assert.NoError(t, err)

	type UpdateProfileRequestInfo struct {
		ContainerId string
		Request     optimizer.UpdateProfileRequest
	}

	var requests []UpdateProfileRequestInfo
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
			response, err := c.UpdateProfile(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientUpdateRecommendation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "UpdateRecommendation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRecommendation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "UpdateRecommendation", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

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

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerOptimizerClientUpdateResourceAction(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "UpdateResourceAction")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateResourceAction is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Optimizer", "UpdateResourceAction", createOptimizerOptimizerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.OptimizerClient)

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
