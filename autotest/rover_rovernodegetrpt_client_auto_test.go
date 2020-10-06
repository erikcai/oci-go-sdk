package autotest

import (
	"github.com/oracle/oci-go-sdk/v26/common"
	"github.com/oracle/oci-go-sdk/v26/rover"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createRoverRoverNodeGetRPTClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := rover.NewRoverNodeGetRPTClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="edge_rover_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="JIRA-OPS"
func TestRoverRoverNodeGetRPTClientGetRoverNodeRPT(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("rover", "GetRoverNodeRPT")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRoverNodeRPT is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("rover", "RoverNodeGetRPT", "GetRoverNodeRPT", createRoverRoverNodeGetRPTClientWithProvider)
	assert.NoError(t, err)
	c := cc.(rover.RoverNodeGetRPTClient)

	body, err := testClient.getRequests("rover", "GetRoverNodeRPT")
	assert.NoError(t, err)

	type GetRoverNodeRPTRequestInfo struct {
		ContainerId string
		Request     rover.GetRoverNodeRPTRequest
	}

	var requests []GetRoverNodeRPTRequestInfo
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
			response, err := c.GetRoverNodeRPT(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
