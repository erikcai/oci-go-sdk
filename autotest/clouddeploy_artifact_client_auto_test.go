package autotest

import (
	"github.com/erikcai/oci-go-sdk/v33/clouddeploy"
	"github.com/erikcai/oci-go-sdk/v33/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createClouddeployArtifactClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := clouddeploy.NewArtifactClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployArtifactClientCreateArtifact(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "CreateArtifact")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateArtifact is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Artifact", "CreateArtifact", createClouddeployArtifactClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.ArtifactClient)

	body, err := testClient.getRequests("clouddeploy", "CreateArtifact")
	assert.NoError(t, err)

	type CreateArtifactRequestInfo struct {
		ContainerId string
		Request     clouddeploy.CreateArtifactRequest
	}

	var requests []CreateArtifactRequestInfo
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
			response, err := c.CreateArtifact(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployArtifactClientDeleteArtifact(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "DeleteArtifact")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteArtifact is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Artifact", "DeleteArtifact", createClouddeployArtifactClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.ArtifactClient)

	body, err := testClient.getRequests("clouddeploy", "DeleteArtifact")
	assert.NoError(t, err)

	type DeleteArtifactRequestInfo struct {
		ContainerId string
		Request     clouddeploy.DeleteArtifactRequest
	}

	var requests []DeleteArtifactRequestInfo
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
			response, err := c.DeleteArtifact(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployArtifactClientGetArtifact(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "GetArtifact")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetArtifact is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Artifact", "GetArtifact", createClouddeployArtifactClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.ArtifactClient)

	body, err := testClient.getRequests("clouddeploy", "GetArtifact")
	assert.NoError(t, err)

	type GetArtifactRequestInfo struct {
		ContainerId string
		Request     clouddeploy.GetArtifactRequest
	}

	var requests []GetArtifactRequestInfo
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
			response, err := c.GetArtifact(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployArtifactClientListArtifacts(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "ListArtifacts")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListArtifacts is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Artifact", "ListArtifacts", createClouddeployArtifactClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.ArtifactClient)

	body, err := testClient.getRequests("clouddeploy", "ListArtifacts")
	assert.NoError(t, err)

	type ListArtifactsRequestInfo struct {
		ContainerId string
		Request     clouddeploy.ListArtifactsRequest
	}

	var requests []ListArtifactsRequestInfo
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
				r := req.(*clouddeploy.ListArtifactsRequest)
				return c.ListArtifacts(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]clouddeploy.ListArtifactsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(clouddeploy.ListArtifactsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployArtifactClientUpdateArtifact(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "UpdateArtifact")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateArtifact is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Artifact", "UpdateArtifact", createClouddeployArtifactClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.ArtifactClient)

	body, err := testClient.getRequests("clouddeploy", "UpdateArtifact")
	assert.NoError(t, err)

	type UpdateArtifactRequestInfo struct {
		ContainerId string
		Request     clouddeploy.UpdateArtifactRequest
	}

	var requests []UpdateArtifactRequestInfo
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
			response, err := c.UpdateArtifact(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
