package autotest

import (
	"github.com/oracle/oci-go-sdk/v30/clouddeploy"
	"github.com/oracle/oci-go-sdk/v30/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createClouddeployDeploymentClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := clouddeploy.NewDeploymentClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployDeploymentClientApproveDeployment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "ApproveDeployment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ApproveDeployment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Deployment", "ApproveDeployment", createClouddeployDeploymentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.DeploymentClient)

	body, err := testClient.getRequests("clouddeploy", "ApproveDeployment")
	assert.NoError(t, err)

	type ApproveDeploymentRequestInfo struct {
		ContainerId string
		Request     clouddeploy.ApproveDeploymentRequest
	}

	var requests []ApproveDeploymentRequestInfo
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
			response, err := c.ApproveDeployment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployDeploymentClientCreateDeployment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "CreateDeployment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDeployment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Deployment", "CreateDeployment", createClouddeployDeploymentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.DeploymentClient)

	body, err := testClient.getRequests("clouddeploy", "CreateDeployment")
	assert.NoError(t, err)

	type CreateDeploymentRequestInfo struct {
		ContainerId string
		Request     clouddeploy.CreateDeploymentRequest
	}

	var requests []CreateDeploymentRequestInfo
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
			response, err := c.CreateDeployment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployDeploymentClientDeleteDeployment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "DeleteDeployment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDeployment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Deployment", "DeleteDeployment", createClouddeployDeploymentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.DeploymentClient)

	body, err := testClient.getRequests("clouddeploy", "DeleteDeployment")
	assert.NoError(t, err)

	type DeleteDeploymentRequestInfo struct {
		ContainerId string
		Request     clouddeploy.DeleteDeploymentRequest
	}

	var requests []DeleteDeploymentRequestInfo
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
			response, err := c.DeleteDeployment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployDeploymentClientGetDeployment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "GetDeployment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDeployment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Deployment", "GetDeployment", createClouddeployDeploymentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.DeploymentClient)

	body, err := testClient.getRequests("clouddeploy", "GetDeployment")
	assert.NoError(t, err)

	type GetDeploymentRequestInfo struct {
		ContainerId string
		Request     clouddeploy.GetDeploymentRequest
	}

	var requests []GetDeploymentRequestInfo
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
			response, err := c.GetDeployment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployDeploymentClientListDeployments(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "ListDeployments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDeployments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Deployment", "ListDeployments", createClouddeployDeploymentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.DeploymentClient)

	body, err := testClient.getRequests("clouddeploy", "ListDeployments")
	assert.NoError(t, err)

	type ListDeploymentsRequestInfo struct {
		ContainerId string
		Request     clouddeploy.ListDeploymentsRequest
	}

	var requests []ListDeploymentsRequestInfo
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
				r := req.(*clouddeploy.ListDeploymentsRequest)
				return c.ListDeployments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]clouddeploy.ListDeploymentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(clouddeploy.ListDeploymentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployDeploymentClientUpdateDeployment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "UpdateDeployment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDeployment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Deployment", "UpdateDeployment", createClouddeployDeploymentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.DeploymentClient)

	body, err := testClient.getRequests("clouddeploy", "UpdateDeployment")
	assert.NoError(t, err)

	type UpdateDeploymentRequestInfo struct {
		ContainerId string
		Request     clouddeploy.UpdateDeploymentRequest
	}

	var requests []UpdateDeploymentRequestInfo
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
			response, err := c.UpdateDeployment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
