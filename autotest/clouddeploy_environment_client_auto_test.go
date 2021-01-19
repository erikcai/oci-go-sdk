package autotest

import (
	"github.com/oracle/oci-go-sdk/v33/clouddeploy"
	"github.com/oracle/oci-go-sdk/v33/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createClouddeployEnvironmentClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := clouddeploy.NewEnvironmentClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployEnvironmentClientCreateEnvironment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "CreateEnvironment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateEnvironment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Environment", "CreateEnvironment", createClouddeployEnvironmentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.EnvironmentClient)

	body, err := testClient.getRequests("clouddeploy", "CreateEnvironment")
	assert.NoError(t, err)

	type CreateEnvironmentRequestInfo struct {
		ContainerId string
		Request     clouddeploy.CreateEnvironmentRequest
	}

	var requests []CreateEnvironmentRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateEnvironmentRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateEnvironmentDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "environmentType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"FUNCTION":               &clouddeploy.CreateFunctionEnvironmentDetails{},
				"LOAD_BALANCE_LISTENER":  &clouddeploy.CreateLoadBalancerListenerEnvironmentDetails{},
				"COMPUTE_INSTANCE_GROUP": &clouddeploy.CreateComputeInstanceGroupEnvironmentDetails{},
				"OKE_CLUSTER":            &clouddeploy.CreateOkeClusterEnvironmentDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateEnvironment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployEnvironmentClientDeleteEnvironment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "DeleteEnvironment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteEnvironment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Environment", "DeleteEnvironment", createClouddeployEnvironmentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.EnvironmentClient)

	body, err := testClient.getRequests("clouddeploy", "DeleteEnvironment")
	assert.NoError(t, err)

	type DeleteEnvironmentRequestInfo struct {
		ContainerId string
		Request     clouddeploy.DeleteEnvironmentRequest
	}

	var requests []DeleteEnvironmentRequestInfo
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
			response, err := c.DeleteEnvironment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployEnvironmentClientGetEnvironment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "GetEnvironment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetEnvironment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Environment", "GetEnvironment", createClouddeployEnvironmentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.EnvironmentClient)

	body, err := testClient.getRequests("clouddeploy", "GetEnvironment")
	assert.NoError(t, err)

	type GetEnvironmentRequestInfo struct {
		ContainerId string
		Request     clouddeploy.GetEnvironmentRequest
	}

	var requests []GetEnvironmentRequestInfo
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
			response, err := c.GetEnvironment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployEnvironmentClientListEnvironments(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "ListEnvironments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListEnvironments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Environment", "ListEnvironments", createClouddeployEnvironmentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.EnvironmentClient)

	body, err := testClient.getRequests("clouddeploy", "ListEnvironments")
	assert.NoError(t, err)

	type ListEnvironmentsRequestInfo struct {
		ContainerId string
		Request     clouddeploy.ListEnvironmentsRequest
	}

	var requests []ListEnvironmentsRequestInfo
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
				r := req.(*clouddeploy.ListEnvironmentsRequest)
				return c.ListEnvironments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]clouddeploy.ListEnvironmentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(clouddeploy.ListEnvironmentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployEnvironmentClientUpdateEnvironment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "UpdateEnvironment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateEnvironment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Environment", "UpdateEnvironment", createClouddeployEnvironmentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.EnvironmentClient)

	body, err := testClient.getRequests("clouddeploy", "UpdateEnvironment")
	assert.NoError(t, err)

	type UpdateEnvironmentRequestInfo struct {
		ContainerId string
		Request     clouddeploy.UpdateEnvironmentRequest
	}

	var requests []UpdateEnvironmentRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]UpdateEnvironmentRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["UpdateEnvironmentDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "environmentType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"COMPUTE_INSTANCE_GROUP": &clouddeploy.UpdateComputeInstanceGroupEnvironmentDetails{},
				"OKE_CLUSTER":            &clouddeploy.UpdateOkeClusterEnvironmentDetails{},
				"FUNCTION":               &clouddeploy.UpdateFunctionEnvironmentDetails{},
				"LOAD_BALANCE_LISTENER":  &clouddeploy.UpdateLoadBalancerListenerEnvironmentDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateEnvironment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
