package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/functions"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createFunctionsInvokeClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := functions.NewFunctionsInvokeClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="serverless_grp@oracle.com" jiraProject="FAAS" opsJiraProject="FAAS"
func TestFunctionsInvokeClientInvokeFunction(t *testing.T) {
	enabled, err := testClient.isApiEnabled("functions", "InvokeFunction")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("InvokeFunction is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("functions", "FunctionsInvoke", "InvokeFunction", createFunctionsInvokeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(functions.FunctionsInvokeClient)

	body, err := testClient.getRequests("functions", "InvokeFunction")
	assert.NoError(t, err)

	type InvokeFunctionRequestInfo struct {
		ContainerId string
		Request     functions.InvokeFunctionRequest
	}

	var requests []InvokeFunctionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.InvokeFunction(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
