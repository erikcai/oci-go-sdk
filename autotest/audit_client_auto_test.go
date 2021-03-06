package autotest

import (
	"github.com/erikcai/oci-go-sdk/v33/audit"
	"github.com/erikcai/oci-go-sdk/v33/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createAuditClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := audit.NewAuditClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_events_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/SA" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/SA"
func TestAuditClientGetConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("audit", "GetConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("audit", "Audit", "GetConfiguration", createAuditClientWithProvider)
	assert.NoError(t, err)
	c := cc.(audit.AuditClient)

	body, err := testClient.getRequests("audit", "GetConfiguration")
	assert.NoError(t, err)

	type GetConfigurationRequestInfo struct {
		ContainerId string
		Request     audit.GetConfigurationRequest
	}

	var requests []GetConfigurationRequestInfo
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
			response, err := c.GetConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_events_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/SA" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/SA"
func TestAuditClientListEvents(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("audit", "ListEvents")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListEvents is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("audit", "Audit", "ListEvents", createAuditClientWithProvider)
	assert.NoError(t, err)
	c := cc.(audit.AuditClient)

	body, err := testClient.getRequests("audit", "ListEvents")
	assert.NoError(t, err)

	type ListEventsRequestInfo struct {
		ContainerId string
		Request     audit.ListEventsRequest
	}

	var requests []ListEventsRequestInfo
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
				r := req.(*audit.ListEventsRequest)
				return c.ListEvents(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]audit.ListEventsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(audit.ListEventsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_events_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/SA" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/SA"
func TestAuditClientUpdateConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("audit", "UpdateConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("audit", "Audit", "UpdateConfiguration", createAuditClientWithProvider)
	assert.NoError(t, err)
	c := cc.(audit.AuditClient)

	body, err := testClient.getRequests("audit", "UpdateConfiguration")
	assert.NoError(t, err)

	type UpdateConfigurationRequestInfo struct {
		ContainerId string
		Request     audit.UpdateConfigurationRequest
	}

	var requests []UpdateConfigurationRequestInfo
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
			response, err := c.UpdateConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
