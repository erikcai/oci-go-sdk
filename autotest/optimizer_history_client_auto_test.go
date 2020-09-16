package autotest

import (
	"github.com/oracle/oci-go-sdk/v25/common"
	"github.com/oracle/oci-go-sdk/v25/optimizer"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createOptimizerHistoryClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := optimizer.NewHistoryClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerHistoryClientListHistories(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "ListHistories")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListHistories is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "History", "ListHistories", createOptimizerHistoryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.HistoryClient)

	body, err := testClient.getRequests("optimizer", "ListHistories")
	assert.NoError(t, err)

	type ListHistoriesRequestInfo struct {
		ContainerId string
		Request     optimizer.ListHistoriesRequest
	}

	var requests []ListHistoriesRequestInfo
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
				r := req.(*optimizer.ListHistoriesRequest)
				return c.ListHistories(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]optimizer.ListHistoriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(optimizer.ListHistoriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
