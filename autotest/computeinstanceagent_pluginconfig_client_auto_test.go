package autotest

import (
	"github.com/oracle/oci-go-sdk/v30/common"
	"github.com/oracle/oci-go-sdk/v30/computeinstanceagent"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createComputeinstanceagentPluginconfigClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := computeinstanceagent.NewPluginconfigClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_osi_grp@oracle.com" jiraProject="OSI" opsJiraProject="IMAGE"
func TestComputeinstanceagentPluginconfigClientListInstanceagentAvailablePlugins(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("computeinstanceagent", "ListInstanceagentAvailablePlugins")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInstanceagentAvailablePlugins is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("computeinstanceagent", "Pluginconfig", "ListInstanceagentAvailablePlugins", createComputeinstanceagentPluginconfigClientWithProvider)
	assert.NoError(t, err)
	c := cc.(computeinstanceagent.PluginconfigClient)

	body, err := testClient.getRequests("computeinstanceagent", "ListInstanceagentAvailablePlugins")
	assert.NoError(t, err)

	type ListInstanceagentAvailablePluginsRequestInfo struct {
		ContainerId string
		Request     computeinstanceagent.ListInstanceagentAvailablePluginsRequest
	}

	var requests []ListInstanceagentAvailablePluginsRequestInfo
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
				r := req.(*computeinstanceagent.ListInstanceagentAvailablePluginsRequest)
				return c.ListInstanceagentAvailablePlugins(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]computeinstanceagent.ListInstanceagentAvailablePluginsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(computeinstanceagent.ListInstanceagentAvailablePluginsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
