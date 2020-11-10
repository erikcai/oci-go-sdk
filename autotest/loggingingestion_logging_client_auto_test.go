package autotest

import (
	"github.com/oracle/oci-go-sdk/v28/common"
	"github.com/oracle/oci-go-sdk/v28/loggingingestion"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createLoggingingestionLoggingClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := loggingingestion.NewLoggingClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLoggingingestionLoggingClientPutLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loggingingestion", "PutLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PutLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loggingingestion", "Logging", "PutLogs", createLoggingingestionLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loggingingestion.LoggingClient)

	body, err := testClient.getRequests("loggingingestion", "PutLogs")
	assert.NoError(t, err)

	type PutLogsRequestInfo struct {
		ContainerId string
		Request     loggingingestion.PutLogsRequest
	}

	var requests []PutLogsRequestInfo
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
			response, err := c.PutLogs(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
