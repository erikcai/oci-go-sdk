package autotest

import (
	"github.com/oracle/oci-go-sdk/v26/common"
	"github.com/oracle/oci-go-sdk/v26/loggingsearch"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createLoggingsearchLogSearchClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := loggingsearch.NewLogSearchClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLoggingsearchLogSearchClientSearchLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loggingsearch", "SearchLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SearchLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loggingsearch", "LogSearch", "SearchLogs", createLoggingsearchLogSearchClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loggingsearch.LogSearchClient)

	body, err := testClient.getRequests("loggingsearch", "SearchLogs")
	assert.NoError(t, err)

	type SearchLogsRequestInfo struct {
		ContainerId string
		Request     loggingsearch.SearchLogsRequest
	}

	var requests []SearchLogsRequestInfo
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
				r := req.(*loggingsearch.SearchLogsRequest)
				return c.SearchLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loggingsearch.SearchLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loggingsearch.SearchLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
