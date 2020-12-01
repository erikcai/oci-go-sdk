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

func createClouddeployApplicationClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := clouddeploy.NewApplicationClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployApplicationClientChangeApplicationCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "ChangeApplicationCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeApplicationCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Application", "ChangeApplicationCompartment", createClouddeployApplicationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.ApplicationClient)

	body, err := testClient.getRequests("clouddeploy", "ChangeApplicationCompartment")
	assert.NoError(t, err)

	type ChangeApplicationCompartmentRequestInfo struct {
		ContainerId string
		Request     clouddeploy.ChangeApplicationCompartmentRequest
	}

	var requests []ChangeApplicationCompartmentRequestInfo
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
			response, err := c.ChangeApplicationCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployApplicationClientCreateApplication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "CreateApplication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateApplication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Application", "CreateApplication", createClouddeployApplicationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.ApplicationClient)

	body, err := testClient.getRequests("clouddeploy", "CreateApplication")
	assert.NoError(t, err)

	type CreateApplicationRequestInfo struct {
		ContainerId string
		Request     clouddeploy.CreateApplicationRequest
	}

	var requests []CreateApplicationRequestInfo
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
			response, err := c.CreateApplication(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployApplicationClientDeleteApplication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "DeleteApplication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteApplication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Application", "DeleteApplication", createClouddeployApplicationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.ApplicationClient)

	body, err := testClient.getRequests("clouddeploy", "DeleteApplication")
	assert.NoError(t, err)

	type DeleteApplicationRequestInfo struct {
		ContainerId string
		Request     clouddeploy.DeleteApplicationRequest
	}

	var requests []DeleteApplicationRequestInfo
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
			response, err := c.DeleteApplication(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployApplicationClientGetApplication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "GetApplication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetApplication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Application", "GetApplication", createClouddeployApplicationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.ApplicationClient)

	body, err := testClient.getRequests("clouddeploy", "GetApplication")
	assert.NoError(t, err)

	type GetApplicationRequestInfo struct {
		ContainerId string
		Request     clouddeploy.GetApplicationRequest
	}

	var requests []GetApplicationRequestInfo
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
			response, err := c.GetApplication(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployApplicationClientListApplications(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "ListApplications")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListApplications is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Application", "ListApplications", createClouddeployApplicationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.ApplicationClient)

	body, err := testClient.getRequests("clouddeploy", "ListApplications")
	assert.NoError(t, err)

	type ListApplicationsRequestInfo struct {
		ContainerId string
		Request     clouddeploy.ListApplicationsRequest
	}

	var requests []ListApplicationsRequestInfo
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
				r := req.(*clouddeploy.ListApplicationsRequest)
				return c.ListApplications(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]clouddeploy.ListApplicationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(clouddeploy.ListApplicationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployApplicationClientUpdateApplication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "UpdateApplication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateApplication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Application", "UpdateApplication", createClouddeployApplicationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.ApplicationClient)

	body, err := testClient.getRequests("clouddeploy", "UpdateApplication")
	assert.NoError(t, err)

	type UpdateApplicationRequestInfo struct {
		ContainerId string
		Request     clouddeploy.UpdateApplicationRequest
	}

	var requests []UpdateApplicationRequestInfo
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
			response, err := c.UpdateApplication(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
