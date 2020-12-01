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

func createClouddeployStageClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := clouddeploy.NewStageClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployStageClientCreateStage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "CreateStage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateStage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Stage", "CreateStage", createClouddeployStageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.StageClient)

	body, err := testClient.getRequests("clouddeploy", "CreateStage")
	assert.NoError(t, err)

	type CreateStageRequestInfo struct {
		ContainerId string
		Request     clouddeploy.CreateStageRequest
	}

	var requests []CreateStageRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateStageRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateStageDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "stageType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"RUN_VALIDATION_TEST_ON_FUNCTION":           &clouddeploy.CreateRunValidationTestOnFunctionStageDetails{},
				"DEPLOY_FUNCTION":                           &clouddeploy.CreateDeployFunctionStageDetails{},
				"RUN_VALIDATION_TEST_ON_COMPUTE_INSTANCE":   &clouddeploy.CreateRunValidationTestOnComputeInstanceStageDetails{},
				"OKE_DEPLOYMENT":                            &clouddeploy.CreateOkeDeploymentStageDetails{},
				"MANUAL_APPROVAL":                           &clouddeploy.CreateManualApprovalStageDetails{},
				"RUN_DEPLOYMENT_PIPELINE":                   &clouddeploy.CreateRunDeploymentStageDetails{},
				"UPDATE_FUNCTION_APPLICATION":               &clouddeploy.CreateUpdateFunctionApplicationStageDetails{},
				"LOAD_BALANCER_TRAFFIC_SHIFT":               &clouddeploy.CreateLoadBalancerTrafficShiftStageDetails{},
				"INVOKE_FUNCTION":                           &clouddeploy.CreateInvokeFunctionStageDetails{},
				"RUN_JOB_ON_COMPUTE_INSTANCE":               &clouddeploy.CreateRunJobOnComputeInstanceStageDetails{},
				"RUN_JOB_ON_COMPUTE_INSTANCE_GROUP":         &clouddeploy.CreateRunJobOnComputeInstanceGroupStageDetails{},
				"WAIT":                                      &clouddeploy.CreateWaitStageDetails{},
				"COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT": &clouddeploy.CreateComputeInstanceGroupRollingDeploymentStageDetails{},
				"RUN_VALIDATION_TEST_ON_OKE":                &clouddeploy.CreateRunValidationTestOnOkeStageDetails{},
				"RUN_OKE_JOB":                               &clouddeploy.CreateRunOkeJobStageDetails{},
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
			response, err := c.CreateStage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployStageClientDeleteStage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "DeleteStage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteStage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Stage", "DeleteStage", createClouddeployStageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.StageClient)

	body, err := testClient.getRequests("clouddeploy", "DeleteStage")
	assert.NoError(t, err)

	type DeleteStageRequestInfo struct {
		ContainerId string
		Request     clouddeploy.DeleteStageRequest
	}

	var requests []DeleteStageRequestInfo
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
			response, err := c.DeleteStage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployStageClientGetStage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "GetStage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Stage", "GetStage", createClouddeployStageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.StageClient)

	body, err := testClient.getRequests("clouddeploy", "GetStage")
	assert.NoError(t, err)

	type GetStageRequestInfo struct {
		ContainerId string
		Request     clouddeploy.GetStageRequest
	}

	var requests []GetStageRequestInfo
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
			response, err := c.GetStage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployStageClientListStages(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "ListStages")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListStages is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Stage", "ListStages", createClouddeployStageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.StageClient)

	body, err := testClient.getRequests("clouddeploy", "ListStages")
	assert.NoError(t, err)

	type ListStagesRequestInfo struct {
		ContainerId string
		Request     clouddeploy.ListStagesRequest
	}

	var requests []ListStagesRequestInfo
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
				r := req.(*clouddeploy.ListStagesRequest)
				return c.ListStages(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]clouddeploy.ListStagesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(clouddeploy.ListStagesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_devlifecycle_group_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestClouddeployStageClientUpdateStage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("clouddeploy", "UpdateStage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateStage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("clouddeploy", "Stage", "UpdateStage", createClouddeployStageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(clouddeploy.StageClient)

	body, err := testClient.getRequests("clouddeploy", "UpdateStage")
	assert.NoError(t, err)

	type UpdateStageRequestInfo struct {
		ContainerId string
		Request     clouddeploy.UpdateStageRequest
	}

	var requests []UpdateStageRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]UpdateStageRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["UpdateStageDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "stageType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"WAIT":                                      &clouddeploy.UpdateWaitStageDetails{},
				"LOAD_BALANCER_TRAFFIC_SHIFT":               &clouddeploy.UpdateLoadBalancerTrafficShiftStageDetails{},
				"RUN_JOB_ON_COMPUTE_INSTANCE_GROUP":         &clouddeploy.UpdateRunJobOnComputeInstanceGroupStageDetails{},
				"RUN_VALIDATION_TEST_ON_COMPUTE_INSTANCE":   &clouddeploy.UpdateRunValidationTestOnComputeInstanceStageDetails{},
				"UPDATE_FUNCTION_APPLICATION":               &clouddeploy.UpdateFunctionApplicationStageDetails{},
				"RUN_VALIDATION_TEST_ON_FUNCTION":           &clouddeploy.UpdateRunValidationTestOnFunctionStageDetails{},
				"RUN_JOB_ON_COMPUTE_INSTANCE":               &clouddeploy.UpdateRunJobOnComputeInstanceStageDetails{},
				"RUN_DEPLOYMENT_PIPELINE":                   &clouddeploy.UpdateRunDeploymentStageDetails{},
				"OKE_DEPLOYMENT":                            &clouddeploy.UpdateOkeDeploymentStageDetails{},
				"RUN_OKE_JOB":                               &clouddeploy.UpdateRunOkeJobStageDetails{},
				"MANUAL_APPROVAL":                           &clouddeploy.UpdateManualApprovalStageDetails{},
				"COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT": &clouddeploy.UpdateComputeInstanceGroupRollingDeploymentStageDetails{},
				"DEPLOY_FUNCTION":                           &clouddeploy.UpdateDeployFunctionStageDetails{},
				"RUN_VALIDATION_TEST_ON_OKE":                &clouddeploy.UpdateRunValidationTestOnOkeStageDetails{},
				"INVOKE_FUNCTION":                           &clouddeploy.UpdateInvokeFunctionStageDetails{},
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
			response, err := c.UpdateStage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
