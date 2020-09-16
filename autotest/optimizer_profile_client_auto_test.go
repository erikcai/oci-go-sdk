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

func createOptimizerProfileClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := optimizer.NewProfileClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerProfileClientCreateProfile(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "CreateProfile")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateProfile is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Profile", "CreateProfile", createOptimizerProfileClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.ProfileClient)

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
func TestOptimizerProfileClientDeleteProfile(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "DeleteProfile")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteProfile is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Profile", "DeleteProfile", createOptimizerProfileClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.ProfileClient)

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
func TestOptimizerProfileClientGetProfile(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "GetProfile")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetProfile is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Profile", "GetProfile", createOptimizerProfileClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.ProfileClient)

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
func TestOptimizerProfileClientListProfiles(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "ListProfiles")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListProfiles is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Profile", "ListProfiles", createOptimizerProfileClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.ProfileClient)

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
func TestOptimizerProfileClientUpdateProfile(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "UpdateProfile")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateProfile is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "Profile", "UpdateProfile", createOptimizerProfileClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.ProfileClient)

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
