package autotest

import (
	"github.com/oracle/oci-go-sdk/cloudevents"
	"github.com/oracle/oci-go-sdk/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// IssueRoutingInfo email="oci_events_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/CLEV" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/CLEV"
func TestCloudEventsClientCreateRule(t *testing.T) {
	enabled, err := testClient.isApiEnabled("cloudevents", "CreateRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateRule is not enabled by the testing service")
	}
	c, err := cloudevents.NewCloudEventsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("cloudevents", "CreateRule")
	assert.NoError(t, err)

	type CreateRuleRequestInfo struct {
		ContainerId string
		Request     cloudevents.CreateRuleRequest
	}

	var requests []CreateRuleRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="oci_events_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/CLEV" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/CLEV"
func TestCloudEventsClientDeleteRule(t *testing.T) {
	enabled, err := testClient.isApiEnabled("cloudevents", "DeleteRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteRule is not enabled by the testing service")
	}
	c, err := cloudevents.NewCloudEventsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("cloudevents", "DeleteRule")
	assert.NoError(t, err)

	type DeleteRuleRequestInfo struct {
		ContainerId string
		Request     cloudevents.DeleteRuleRequest
	}

	var requests []DeleteRuleRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="oci_events_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/CLEV" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/CLEV"
func TestCloudEventsClientGetRule(t *testing.T) {
	enabled, err := testClient.isApiEnabled("cloudevents", "GetRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRule is not enabled by the testing service")
	}
	c, err := cloudevents.NewCloudEventsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("cloudevents", "GetRule")
	assert.NoError(t, err)

	type GetRuleRequestInfo struct {
		ContainerId string
		Request     cloudevents.GetRuleRequest
	}

	var requests []GetRuleRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="oci_events_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/CLEV" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/CLEV"
func TestCloudEventsClientListRules(t *testing.T) {
	enabled, err := testClient.isApiEnabled("cloudevents", "ListRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRules is not enabled by the testing service")
	}
	c, err := cloudevents.NewCloudEventsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("cloudevents", "ListRules")
	assert.NoError(t, err)

	type ListRulesRequestInfo struct {
		ContainerId string
		Request     cloudevents.ListRulesRequest
	}

	var requests []ListRulesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*cloudevents.ListRulesRequest)
				return c.ListRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cloudevents.ListRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cloudevents.ListRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="oci_events_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/CLEV" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/CLEV"
func TestCloudEventsClientUpdateRule(t *testing.T) {
	enabled, err := testClient.isApiEnabled("cloudevents", "UpdateRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRule is not enabled by the testing service")
	}
	c, err := cloudevents.NewCloudEventsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("cloudevents", "UpdateRule")
	assert.NoError(t, err)

	type UpdateRuleRequestInfo struct {
		ContainerId string
		Request     cloudevents.UpdateRuleRequest
	}

	var requests []UpdateRuleRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}