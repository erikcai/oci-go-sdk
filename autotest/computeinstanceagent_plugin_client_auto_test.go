package autotest

import (
	"github.com/oracle/oci-go-sdk/v29/common"
	"github.com/oracle/oci-go-sdk/v29/computeinstanceagent"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createComputeinstanceagentPluginClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := computeinstanceagent.NewPluginClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_osi_grp@oracle.com" jiraProject="OSI" opsJiraProject="IMAGE"
func TestComputeinstanceagentPluginClientGetInstanceAgentPlugin(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("computeinstanceagent", "GetInstanceAgentPlugin")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInstanceAgentPlugin is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("computeinstanceagent", "Plugin", "GetInstanceAgentPlugin", createComputeinstanceagentPluginClientWithProvider)
	assert.NoError(t, err)
	c := cc.(computeinstanceagent.PluginClient)

	body, err := testClient.getRequests("computeinstanceagent", "GetInstanceAgentPlugin")
	assert.NoError(t, err)

	type GetInstanceAgentPluginRequestInfo struct {
		ContainerId string
		Request     computeinstanceagent.GetInstanceAgentPluginRequest
	}

	var requests []GetInstanceAgentPluginRequestInfo
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
			response, err := c.GetInstanceAgentPlugin(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_osi_grp@oracle.com" jiraProject="OSI" opsJiraProject="IMAGE"
func TestComputeinstanceagentPluginClientListInstanceAgentPlugins(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("computeinstanceagent", "ListInstanceAgentPlugins")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInstanceAgentPlugins is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("computeinstanceagent", "Plugin", "ListInstanceAgentPlugins", createComputeinstanceagentPluginClientWithProvider)
	assert.NoError(t, err)
	c := cc.(computeinstanceagent.PluginClient)

	body, err := testClient.getRequests("computeinstanceagent", "ListInstanceAgentPlugins")
	assert.NoError(t, err)

	type ListInstanceAgentPluginsRequestInfo struct {
		ContainerId string
		Request     computeinstanceagent.ListInstanceAgentPluginsRequest
	}

	var requests []ListInstanceAgentPluginsRequestInfo
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
				r := req.(*computeinstanceagent.ListInstanceAgentPluginsRequest)
				return c.ListInstanceAgentPlugins(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]computeinstanceagent.ListInstanceAgentPluginsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(computeinstanceagent.ListInstanceAgentPluginsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
