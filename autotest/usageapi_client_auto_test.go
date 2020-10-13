package autotest

import (
	"github.com/oracle/oci-go-sdk/v27/common"
	"github.com/oracle/oci-go-sdk/v27/usageapi"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createUsageapiUsageapiClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := usageapi.NewUsageapiClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_metering_team_us_grp@oracle.com" jiraProject="METER" opsJiraProject="MTRC"
func TestUsageapiUsageapiClientRequestSummarizedConfigurations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("usageapi", "RequestSummarizedConfigurations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSummarizedConfigurations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("usageapi", "Usageapi", "RequestSummarizedConfigurations", createUsageapiUsageapiClientWithProvider)
	assert.NoError(t, err)
	c := cc.(usageapi.UsageapiClient)

	body, err := testClient.getRequests("usageapi", "RequestSummarizedConfigurations")
	assert.NoError(t, err)

	type RequestSummarizedConfigurationsRequestInfo struct {
		ContainerId string
		Request     usageapi.RequestSummarizedConfigurationsRequest
	}

	var requests []RequestSummarizedConfigurationsRequestInfo
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
			response, err := c.RequestSummarizedConfigurations(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_metering_team_us_grp@oracle.com" jiraProject="METER" opsJiraProject="MTRC"
func TestUsageapiUsageapiClientRequestSummarizedUsages(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("usageapi", "RequestSummarizedUsages")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSummarizedUsages is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("usageapi", "Usageapi", "RequestSummarizedUsages", createUsageapiUsageapiClientWithProvider)
	assert.NoError(t, err)
	c := cc.(usageapi.UsageapiClient)

	body, err := testClient.getRequests("usageapi", "RequestSummarizedUsages")
	assert.NoError(t, err)

	type RequestSummarizedUsagesRequestInfo struct {
		ContainerId string
		Request     usageapi.RequestSummarizedUsagesRequest
	}

	var requests []RequestSummarizedUsagesRequestInfo
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
				r := req.(*usageapi.RequestSummarizedUsagesRequest)
				return c.RequestSummarizedUsages(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]usageapi.RequestSummarizedUsagesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(usageapi.RequestSummarizedUsagesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
