package autotest

import (
	"github.com/oracle/oci-go-sdk/v29/common"
	"github.com/oracle/oci-go-sdk/v29/managementagent"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createManagementagentManagementAgentClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := managementagent.NewManagementAgentClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientCreateManagementAgentInstallKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "CreateManagementAgentInstallKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateManagementAgentInstallKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "CreateManagementAgentInstallKey", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "CreateManagementAgentInstallKey")
	assert.NoError(t, err)

	type CreateManagementAgentInstallKeyRequestInfo struct {
		ContainerId string
		Request     managementagent.CreateManagementAgentInstallKeyRequest
	}

	var requests []CreateManagementAgentInstallKeyRequestInfo
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
			response, err := c.CreateManagementAgentInstallKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientDeleteManagementAgent(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "DeleteManagementAgent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteManagementAgent is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "DeleteManagementAgent", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "DeleteManagementAgent")
	assert.NoError(t, err)

	type DeleteManagementAgentRequestInfo struct {
		ContainerId string
		Request     managementagent.DeleteManagementAgentRequest
	}

	var requests []DeleteManagementAgentRequestInfo
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
			response, err := c.DeleteManagementAgent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientDeleteManagementAgentInstallKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "DeleteManagementAgentInstallKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteManagementAgentInstallKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "DeleteManagementAgentInstallKey", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "DeleteManagementAgentInstallKey")
	assert.NoError(t, err)

	type DeleteManagementAgentInstallKeyRequestInfo struct {
		ContainerId string
		Request     managementagent.DeleteManagementAgentInstallKeyRequest
	}

	var requests []DeleteManagementAgentInstallKeyRequestInfo
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
			response, err := c.DeleteManagementAgentInstallKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientDeleteWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "DeleteWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "DeleteWorkRequest", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "DeleteWorkRequest")
	assert.NoError(t, err)

	type DeleteWorkRequestRequestInfo struct {
		ContainerId string
		Request     managementagent.DeleteWorkRequestRequest
	}

	var requests []DeleteWorkRequestRequestInfo
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
			response, err := c.DeleteWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientDeployPlugins(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "DeployPlugins")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeployPlugins is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "DeployPlugins", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "DeployPlugins")
	assert.NoError(t, err)

	type DeployPluginsRequestInfo struct {
		ContainerId string
		Request     managementagent.DeployPluginsRequest
	}

	var requests []DeployPluginsRequestInfo
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
			response, err := c.DeployPlugins(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientGetManagementAgent(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "GetManagementAgent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetManagementAgent is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "GetManagementAgent", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "GetManagementAgent")
	assert.NoError(t, err)

	type GetManagementAgentRequestInfo struct {
		ContainerId string
		Request     managementagent.GetManagementAgentRequest
	}

	var requests []GetManagementAgentRequestInfo
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
			response, err := c.GetManagementAgent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientGetManagementAgentInstallKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "GetManagementAgentInstallKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetManagementAgentInstallKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "GetManagementAgentInstallKey", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "GetManagementAgentInstallKey")
	assert.NoError(t, err)

	type GetManagementAgentInstallKeyRequestInfo struct {
		ContainerId string
		Request     managementagent.GetManagementAgentInstallKeyRequest
	}

	var requests []GetManagementAgentInstallKeyRequestInfo
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
			response, err := c.GetManagementAgentInstallKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientGetManagementAgentInstallKeyContent(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "GetManagementAgentInstallKeyContent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetManagementAgentInstallKeyContent is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "GetManagementAgentInstallKeyContent", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "GetManagementAgentInstallKeyContent")
	assert.NoError(t, err)

	type GetManagementAgentInstallKeyContentRequestInfo struct {
		ContainerId string
		Request     managementagent.GetManagementAgentInstallKeyContentRequest
	}

	var requests []GetManagementAgentInstallKeyContentRequestInfo
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
			response, err := c.GetManagementAgentInstallKeyContent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "GetWorkRequest", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     managementagent.GetWorkRequestRequest
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

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientListAvailabilityHistories(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "ListAvailabilityHistories")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAvailabilityHistories is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "ListAvailabilityHistories", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "ListAvailabilityHistories")
	assert.NoError(t, err)

	type ListAvailabilityHistoriesRequestInfo struct {
		ContainerId string
		Request     managementagent.ListAvailabilityHistoriesRequest
	}

	var requests []ListAvailabilityHistoriesRequestInfo
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
				r := req.(*managementagent.ListAvailabilityHistoriesRequest)
				return c.ListAvailabilityHistories(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]managementagent.ListAvailabilityHistoriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(managementagent.ListAvailabilityHistoriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientListManagementAgentImages(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "ListManagementAgentImages")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListManagementAgentImages is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "ListManagementAgentImages", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "ListManagementAgentImages")
	assert.NoError(t, err)

	type ListManagementAgentImagesRequestInfo struct {
		ContainerId string
		Request     managementagent.ListManagementAgentImagesRequest
	}

	var requests []ListManagementAgentImagesRequestInfo
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
				r := req.(*managementagent.ListManagementAgentImagesRequest)
				return c.ListManagementAgentImages(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]managementagent.ListManagementAgentImagesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(managementagent.ListManagementAgentImagesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientListManagementAgentInstallKeys(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "ListManagementAgentInstallKeys")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListManagementAgentInstallKeys is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "ListManagementAgentInstallKeys", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "ListManagementAgentInstallKeys")
	assert.NoError(t, err)

	type ListManagementAgentInstallKeysRequestInfo struct {
		ContainerId string
		Request     managementagent.ListManagementAgentInstallKeysRequest
	}

	var requests []ListManagementAgentInstallKeysRequestInfo
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
				r := req.(*managementagent.ListManagementAgentInstallKeysRequest)
				return c.ListManagementAgentInstallKeys(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]managementagent.ListManagementAgentInstallKeysResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(managementagent.ListManagementAgentInstallKeysResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientListManagementAgentPlugins(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "ListManagementAgentPlugins")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListManagementAgentPlugins is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "ListManagementAgentPlugins", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "ListManagementAgentPlugins")
	assert.NoError(t, err)

	type ListManagementAgentPluginsRequestInfo struct {
		ContainerId string
		Request     managementagent.ListManagementAgentPluginsRequest
	}

	var requests []ListManagementAgentPluginsRequestInfo
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
				r := req.(*managementagent.ListManagementAgentPluginsRequest)
				return c.ListManagementAgentPlugins(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]managementagent.ListManagementAgentPluginsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(managementagent.ListManagementAgentPluginsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientListManagementAgents(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "ListManagementAgents")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListManagementAgents is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "ListManagementAgents", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "ListManagementAgents")
	assert.NoError(t, err)

	type ListManagementAgentsRequestInfo struct {
		ContainerId string
		Request     managementagent.ListManagementAgentsRequest
	}

	var requests []ListManagementAgentsRequestInfo
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
				r := req.(*managementagent.ListManagementAgentsRequest)
				return c.ListManagementAgents(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]managementagent.ListManagementAgentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(managementagent.ListManagementAgentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "ListWorkRequestErrors", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     managementagent.ListWorkRequestErrorsRequest
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
				r := req.(*managementagent.ListWorkRequestErrorsRequest)
				return c.ListWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]managementagent.ListWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(managementagent.ListWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "ListWorkRequestLogs", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     managementagent.ListWorkRequestLogsRequest
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
				r := req.(*managementagent.ListWorkRequestLogsRequest)
				return c.ListWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]managementagent.ListWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(managementagent.ListWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "ListWorkRequests", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     managementagent.ListWorkRequestsRequest
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
				r := req.(*managementagent.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]managementagent.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(managementagent.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientUpdateManagementAgent(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "UpdateManagementAgent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateManagementAgent is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "UpdateManagementAgent", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "UpdateManagementAgent")
	assert.NoError(t, err)

	type UpdateManagementAgentRequestInfo struct {
		ContainerId string
		Request     managementagent.UpdateManagementAgentRequest
	}

	var requests []UpdateManagementAgentRequestInfo
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
			response, err := c.UpdateManagementAgent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_mgmtagent_macs_ww_grp@oracle.com" jiraProject="MGMTAGENT" opsJiraProject="MGMTAGENT"
func TestManagementagentManagementAgentClientUpdateManagementAgentInstallKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementagent", "UpdateManagementAgentInstallKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateManagementAgentInstallKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementagent", "ManagementAgent", "UpdateManagementAgentInstallKey", createManagementagentManagementAgentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementagent.ManagementAgentClient)

	body, err := testClient.getRequests("managementagent", "UpdateManagementAgentInstallKey")
	assert.NoError(t, err)

	type UpdateManagementAgentInstallKeyRequestInfo struct {
		ContainerId string
		Request     managementagent.UpdateManagementAgentInstallKeyRequest
	}

	var requests []UpdateManagementAgentInstallKeyRequestInfo
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
			response, err := c.UpdateManagementAgentInstallKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
