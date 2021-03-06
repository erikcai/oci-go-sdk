package autotest

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"github.com/erikcai/oci-go-sdk/v33/resourcemanager"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createResourceManagerClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientCancelJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "CancelJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CancelJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "CancelJob", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "CancelJob")
	assert.NoError(t, err)

	type CancelJobRequestInfo struct {
		ContainerId string
		Request     resourcemanager.CancelJobRequest
	}

	var requests []CancelJobRequestInfo
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
			response, err := c.CancelJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientChangeConfigurationSourceProviderCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ChangeConfigurationSourceProviderCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeConfigurationSourceProviderCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ChangeConfigurationSourceProviderCompartment", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ChangeConfigurationSourceProviderCompartment")
	assert.NoError(t, err)

	type ChangeConfigurationSourceProviderCompartmentRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ChangeConfigurationSourceProviderCompartmentRequest
	}

	var requests []ChangeConfigurationSourceProviderCompartmentRequestInfo
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
			response, err := c.ChangeConfigurationSourceProviderCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientChangeStackCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ChangeStackCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeStackCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ChangeStackCompartment", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ChangeStackCompartment")
	assert.NoError(t, err)

	type ChangeStackCompartmentRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ChangeStackCompartmentRequest
	}

	var requests []ChangeStackCompartmentRequestInfo
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
			response, err := c.ChangeStackCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientChangeTemplateCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ChangeTemplateCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeTemplateCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ChangeTemplateCompartment", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ChangeTemplateCompartment")
	assert.NoError(t, err)

	type ChangeTemplateCompartmentRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ChangeTemplateCompartmentRequest
	}

	var requests []ChangeTemplateCompartmentRequestInfo
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
			response, err := c.ChangeTemplateCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientCreateConfigurationSourceProvider(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "CreateConfigurationSourceProvider")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateConfigurationSourceProvider is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "CreateConfigurationSourceProvider", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "CreateConfigurationSourceProvider")
	assert.NoError(t, err)

	type CreateConfigurationSourceProviderRequestInfo struct {
		ContainerId string
		Request     resourcemanager.CreateConfigurationSourceProviderRequest
	}

	var requests []CreateConfigurationSourceProviderRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateConfigurationSourceProviderRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateConfigurationSourceProviderDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "configSourceProviderType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"GITLAB_ACCESS_TOKEN": &resourcemanager.CreateGitlabAccessTokenConfigurationSourceProviderDetails{},
				"GITHUB_ACCESS_TOKEN": &resourcemanager.CreateGithubAccessTokenConfigurationSourceProviderDetails{},
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
			response, err := c.CreateConfigurationSourceProvider(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientCreateJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "CreateJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "CreateJob", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "CreateJob")
	assert.NoError(t, err)

	type CreateJobRequestInfo struct {
		ContainerId string
		Request     resourcemanager.CreateJobRequest
	}

	var requests []CreateJobRequestInfo
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
			response, err := c.CreateJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientCreateStack(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "CreateStack")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateStack is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "CreateStack", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "CreateStack")
	assert.NoError(t, err)

	type CreateStackRequestInfo struct {
		ContainerId string
		Request     resourcemanager.CreateStackRequest
	}

	var requests []CreateStackRequestInfo
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
			response, err := c.CreateStack(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientCreateTemplate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "CreateTemplate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTemplate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "CreateTemplate", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "CreateTemplate")
	assert.NoError(t, err)

	type CreateTemplateRequestInfo struct {
		ContainerId string
		Request     resourcemanager.CreateTemplateRequest
	}

	var requests []CreateTemplateRequestInfo
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
			response, err := c.CreateTemplate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientDeleteConfigurationSourceProvider(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "DeleteConfigurationSourceProvider")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteConfigurationSourceProvider is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "DeleteConfigurationSourceProvider", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "DeleteConfigurationSourceProvider")
	assert.NoError(t, err)

	type DeleteConfigurationSourceProviderRequestInfo struct {
		ContainerId string
		Request     resourcemanager.DeleteConfigurationSourceProviderRequest
	}

	var requests []DeleteConfigurationSourceProviderRequestInfo
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
			response, err := c.DeleteConfigurationSourceProvider(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientDeleteStack(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "DeleteStack")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteStack is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "DeleteStack", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "DeleteStack")
	assert.NoError(t, err)

	type DeleteStackRequestInfo struct {
		ContainerId string
		Request     resourcemanager.DeleteStackRequest
	}

	var requests []DeleteStackRequestInfo
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
			response, err := c.DeleteStack(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientDeleteTemplate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "DeleteTemplate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteTemplate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "DeleteTemplate", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "DeleteTemplate")
	assert.NoError(t, err)

	type DeleteTemplateRequestInfo struct {
		ContainerId string
		Request     resourcemanager.DeleteTemplateRequest
	}

	var requests []DeleteTemplateRequestInfo
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
			response, err := c.DeleteTemplate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientDetectStackDrift(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "DetectStackDrift")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DetectStackDrift is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "DetectStackDrift", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "DetectStackDrift")
	assert.NoError(t, err)

	type DetectStackDriftRequestInfo struct {
		ContainerId string
		Request     resourcemanager.DetectStackDriftRequest
	}

	var requests []DetectStackDriftRequestInfo
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
			response, err := c.DetectStackDrift(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetConfigurationSourceProvider(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetConfigurationSourceProvider")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConfigurationSourceProvider is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetConfigurationSourceProvider", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetConfigurationSourceProvider")
	assert.NoError(t, err)

	type GetConfigurationSourceProviderRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetConfigurationSourceProviderRequest
	}

	var requests []GetConfigurationSourceProviderRequestInfo
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
			response, err := c.GetConfigurationSourceProvider(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetJob", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetJob")
	assert.NoError(t, err)

	type GetJobRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobRequest
	}

	var requests []GetJobRequestInfo
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
			response, err := c.GetJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetJobLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJobLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJobLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetJobLogs", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetJobLogs")
	assert.NoError(t, err)

	type GetJobLogsRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobLogsRequest
	}

	var requests []GetJobLogsRequestInfo
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
				r := req.(*resourcemanager.GetJobLogsRequest)
				return c.GetJobLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.GetJobLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.GetJobLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetJobLogsContent(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJobLogsContent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJobLogsContent is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetJobLogsContent", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetJobLogsContent")
	assert.NoError(t, err)

	type GetJobLogsContentRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobLogsContentRequest
	}

	var requests []GetJobLogsContentRequestInfo
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
			response, err := c.GetJobLogsContent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetJobTfConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJobTfConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJobTfConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetJobTfConfig", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetJobTfConfig")
	assert.NoError(t, err)

	type GetJobTfConfigRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobTfConfigRequest
	}

	var requests []GetJobTfConfigRequestInfo
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
			response, err := c.GetJobTfConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetJobTfState(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJobTfState")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJobTfState is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetJobTfState", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetJobTfState")
	assert.NoError(t, err)

	type GetJobTfStateRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobTfStateRequest
	}

	var requests []GetJobTfStateRequestInfo
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
			response, err := c.GetJobTfState(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetStack(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetStack")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStack is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetStack", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetStack")
	assert.NoError(t, err)

	type GetStackRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetStackRequest
	}

	var requests []GetStackRequestInfo
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
			response, err := c.GetStack(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetStackTfConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetStackTfConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStackTfConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetStackTfConfig", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetStackTfConfig")
	assert.NoError(t, err)

	type GetStackTfConfigRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetStackTfConfigRequest
	}

	var requests []GetStackTfConfigRequestInfo
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
			response, err := c.GetStackTfConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetStackTfState(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetStackTfState")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStackTfState is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetStackTfState", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetStackTfState")
	assert.NoError(t, err)

	type GetStackTfStateRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetStackTfStateRequest
	}

	var requests []GetStackTfStateRequestInfo
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
			response, err := c.GetStackTfState(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetTemplate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetTemplate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTemplate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetTemplate", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetTemplate")
	assert.NoError(t, err)

	type GetTemplateRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetTemplateRequest
	}

	var requests []GetTemplateRequestInfo
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
			response, err := c.GetTemplate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetTemplateLogo(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetTemplateLogo")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTemplateLogo is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetTemplateLogo", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetTemplateLogo")
	assert.NoError(t, err)

	type GetTemplateLogoRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetTemplateLogoRequest
	}

	var requests []GetTemplateLogoRequestInfo
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
			response, err := c.GetTemplateLogo(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetTemplateTfConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetTemplateTfConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTemplateTfConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetTemplateTfConfig", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetTemplateTfConfig")
	assert.NoError(t, err)

	type GetTemplateTfConfigRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetTemplateTfConfigRequest
	}

	var requests []GetTemplateTfConfigRequestInfo
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
			response, err := c.GetTemplateTfConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetWorkRequest", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetWorkRequestRequest
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

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientListConfigurationSourceProviders(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ListConfigurationSourceProviders")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListConfigurationSourceProviders is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ListConfigurationSourceProviders", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ListConfigurationSourceProviders")
	assert.NoError(t, err)

	type ListConfigurationSourceProvidersRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListConfigurationSourceProvidersRequest
	}

	var requests []ListConfigurationSourceProvidersRequestInfo
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
				r := req.(*resourcemanager.ListConfigurationSourceProvidersRequest)
				return c.ListConfigurationSourceProviders(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.ListConfigurationSourceProvidersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.ListConfigurationSourceProvidersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientListJobs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ListJobs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListJobs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ListJobs", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ListJobs")
	assert.NoError(t, err)

	type ListJobsRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListJobsRequest
	}

	var requests []ListJobsRequestInfo
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
				r := req.(*resourcemanager.ListJobsRequest)
				return c.ListJobs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.ListJobsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.ListJobsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientListResourceDiscoveryServices(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ListResourceDiscoveryServices")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListResourceDiscoveryServices is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ListResourceDiscoveryServices", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ListResourceDiscoveryServices")
	assert.NoError(t, err)

	type ListResourceDiscoveryServicesRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListResourceDiscoveryServicesRequest
	}

	var requests []ListResourceDiscoveryServicesRequestInfo
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
			response, err := c.ListResourceDiscoveryServices(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientListStackResourceDriftDetails(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ListStackResourceDriftDetails")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListStackResourceDriftDetails is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ListStackResourceDriftDetails", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ListStackResourceDriftDetails")
	assert.NoError(t, err)

	type ListStackResourceDriftDetailsRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListStackResourceDriftDetailsRequest
	}

	var requests []ListStackResourceDriftDetailsRequestInfo
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
				r := req.(*resourcemanager.ListStackResourceDriftDetailsRequest)
				return c.ListStackResourceDriftDetails(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.ListStackResourceDriftDetailsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.ListStackResourceDriftDetailsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientListStacks(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ListStacks")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListStacks is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ListStacks", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ListStacks")
	assert.NoError(t, err)

	type ListStacksRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListStacksRequest
	}

	var requests []ListStacksRequestInfo
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
				r := req.(*resourcemanager.ListStacksRequest)
				return c.ListStacks(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.ListStacksResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.ListStacksResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientListTemplateCategories(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ListTemplateCategories")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTemplateCategories is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ListTemplateCategories", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ListTemplateCategories")
	assert.NoError(t, err)

	type ListTemplateCategoriesRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListTemplateCategoriesRequest
	}

	var requests []ListTemplateCategoriesRequestInfo
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
			response, err := c.ListTemplateCategories(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientListTemplates(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ListTemplates")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTemplates is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ListTemplates", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ListTemplates")
	assert.NoError(t, err)

	type ListTemplatesRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListTemplatesRequest
	}

	var requests []ListTemplatesRequestInfo
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
				r := req.(*resourcemanager.ListTemplatesRequest)
				return c.ListTemplates(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.ListTemplatesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.ListTemplatesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientListTerraformVersions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ListTerraformVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTerraformVersions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ListTerraformVersions", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ListTerraformVersions")
	assert.NoError(t, err)

	type ListTerraformVersionsRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListTerraformVersionsRequest
	}

	var requests []ListTerraformVersionsRequestInfo
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
			response, err := c.ListTerraformVersions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ListWorkRequestErrors", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListWorkRequestErrorsRequest
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
				r := req.(*resourcemanager.ListWorkRequestErrorsRequest)
				return c.ListWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.ListWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.ListWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ListWorkRequestLogs", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListWorkRequestLogsRequest
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
				r := req.(*resourcemanager.ListWorkRequestLogsRequest)
				return c.ListWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.ListWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.ListWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ListWorkRequests", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListWorkRequestsRequest
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
				r := req.(*resourcemanager.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientUpdateConfigurationSourceProvider(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "UpdateConfigurationSourceProvider")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateConfigurationSourceProvider is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "UpdateConfigurationSourceProvider", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "UpdateConfigurationSourceProvider")
	assert.NoError(t, err)

	type UpdateConfigurationSourceProviderRequestInfo struct {
		ContainerId string
		Request     resourcemanager.UpdateConfigurationSourceProviderRequest
	}

	var requests []UpdateConfigurationSourceProviderRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]UpdateConfigurationSourceProviderRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["UpdateConfigurationSourceProviderDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "configSourceProviderType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"GITLAB_ACCESS_TOKEN": &resourcemanager.UpdateGitlabAccessTokenConfigurationSourceProviderDetails{},
				"GITHUB_ACCESS_TOKEN": &resourcemanager.UpdateGithubAccessTokenConfigurationSourceProviderDetails{},
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
			response, err := c.UpdateConfigurationSourceProvider(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientUpdateJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "UpdateJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "UpdateJob", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "UpdateJob")
	assert.NoError(t, err)

	type UpdateJobRequestInfo struct {
		ContainerId string
		Request     resourcemanager.UpdateJobRequest
	}

	var requests []UpdateJobRequestInfo
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
			response, err := c.UpdateJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientUpdateStack(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "UpdateStack")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateStack is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "UpdateStack", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "UpdateStack")
	assert.NoError(t, err)

	type UpdateStackRequestInfo struct {
		ContainerId string
		Request     resourcemanager.UpdateStackRequest
	}

	var requests []UpdateStackRequestInfo
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
			response, err := c.UpdateStack(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientUpdateTemplate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "UpdateTemplate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTemplate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "UpdateTemplate", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "UpdateTemplate")
	assert.NoError(t, err)

	type UpdateTemplateRequestInfo struct {
		ContainerId string
		Request     resourcemanager.UpdateTemplateRequest
	}

	var requests []UpdateTemplateRequestInfo
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
			response, err := c.UpdateTemplate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
