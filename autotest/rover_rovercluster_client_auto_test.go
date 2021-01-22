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

func createRoverRoverClusterClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := rover.NewRoverClusterClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverClusterClientChangeRoverClusterCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "ChangeRoverClusterCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeRoverClusterCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverCluster", "ChangeRoverClusterCompartment", createRoverRoverClusterClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverClusterClient)

	body, err := testClient.getRequests("rover", "ChangeRoverClusterCompartment")
	assert.NoError(t, err)

	type ChangeRoverClusterCompartmentRequestInfo struct {
		ContainerId string
		Request     rover.ChangeRoverClusterCompartmentRequest
	}

	var requests []ChangeRoverClusterCompartmentRequestInfo
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
			response, err := c.ChangeRoverClusterCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverClusterClientCreateRoverCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "CreateRoverCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateRoverCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverCluster", "CreateRoverCluster", createRoverRoverClusterClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverClusterClient)

	body, err := testClient.getRequests("rover", "CreateRoverCluster")
	assert.NoError(t, err)

	type CreateRoverClusterRequestInfo struct {
		ContainerId string
		Request     rover.CreateRoverClusterRequest
	}

	var requests []CreateRoverClusterRequestInfo
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
			response, err := c.CreateRoverCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverClusterClientDeleteRoverCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "DeleteRoverCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteRoverCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverCluster", "DeleteRoverCluster", createRoverRoverClusterClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverClusterClient)

	body, err := testClient.getRequests("rover", "DeleteRoverCluster")
	assert.NoError(t, err)

	type DeleteRoverClusterRequestInfo struct {
		ContainerId string
		Request     rover.DeleteRoverClusterRequest
	}

	var requests []DeleteRoverClusterRequestInfo
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
			response, err := c.DeleteRoverCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverClusterClientGetRoverCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "GetRoverCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRoverCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverCluster", "GetRoverCluster", createRoverRoverClusterClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverClusterClient)

	body, err := testClient.getRequests("rover", "GetRoverCluster")
	assert.NoError(t, err)

	type GetRoverClusterRequestInfo struct {
		ContainerId string
		Request     rover.GetRoverClusterRequest
	}

	var requests []GetRoverClusterRequestInfo
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
			response, err := c.GetRoverCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverClusterClientGetRoverClusterCertificate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "GetRoverClusterCertificate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRoverClusterCertificate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverCluster", "GetRoverClusterCertificate", createRoverRoverClusterClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverClusterClient)

	body, err := testClient.getRequests("rover", "GetRoverClusterCertificate")
	assert.NoError(t, err)

	type GetRoverClusterCertificateRequestInfo struct {
		ContainerId string
		Request     rover.GetRoverClusterCertificateRequest
	}

	var requests []GetRoverClusterCertificateRequestInfo
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
			response, err := c.GetRoverClusterCertificate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverClusterClientListRoverClusters(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "ListRoverClusters")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRoverClusters is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverCluster", "ListRoverClusters", createRoverRoverClusterClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverClusterClient)

	body, err := testClient.getRequests("rover", "ListRoverClusters")
	assert.NoError(t, err)

	type ListRoverClustersRequestInfo struct {
		ContainerId string
		Request     rover.ListRoverClustersRequest
	}

	var requests []ListRoverClustersRequestInfo
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
				r := req.(*rover.ListRoverClustersRequest)
				return c.ListRoverClusters(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]rover.ListRoverClustersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(rover.ListRoverClustersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverClusterClientUpdateRoverCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "UpdateRoverCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRoverCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverCluster", "UpdateRoverCluster", createRoverRoverClusterClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverClusterClient)

	body, err := testClient.getRequests("rover", "UpdateRoverCluster")
	assert.NoError(t, err)

	type UpdateRoverClusterRequestInfo struct {
		ContainerId string
		Request     rover.UpdateRoverClusterRequest
	}

	var requests []UpdateRoverClusterRequestInfo
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
			response, err := c.UpdateRoverCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
