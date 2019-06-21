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

func createEventsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := cloudevents.NewEventsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_events_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/CLEV" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/CLEV"
func TestEventsClientChangeRuleCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudevents", "ChangeRuleCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeRuleCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudevents", "Events", "ChangeRuleCompartment", createEventsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudevents.EventsClient)

	body, err := testClient.getRequests("cloudevents", "ChangeRuleCompartment")
	assert.NoError(t, err)

	type ChangeRuleCompartmentRequestInfo struct {
		ContainerId string
		Request     cloudevents.ChangeRuleCompartmentRequest
	}

	var requests []ChangeRuleCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ChangeRuleCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_events_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/CLEV" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/CLEV"
func TestEventsClientCreateRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudevents", "CreateRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudevents", "Events", "CreateRule", createEventsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudevents.EventsClient)

	body, err := testClient.getRequests("cloudevents", "CreateRule")
	assert.NoError(t, err)

	type CreateRuleRequestInfo struct {
		ContainerId string
		Request     cloudevents.CreateRuleRequest
	}

	var requests []CreateRuleRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
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

// IssueRoutingInfo tag="default" email="oci_events_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/CLEV" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/CLEV"
func TestEventsClientDeleteRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudevents", "DeleteRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudevents", "Events", "DeleteRule", createEventsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudevents.EventsClient)

	body, err := testClient.getRequests("cloudevents", "DeleteRule")
	assert.NoError(t, err)

	type DeleteRuleRequestInfo struct {
		ContainerId string
		Request     cloudevents.DeleteRuleRequest
	}

	var requests []DeleteRuleRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
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

// IssueRoutingInfo tag="default" email="oci_events_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/CLEV" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/CLEV"
func TestEventsClientGetRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudevents", "GetRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudevents", "Events", "GetRule", createEventsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudevents.EventsClient)

	body, err := testClient.getRequests("cloudevents", "GetRule")
	assert.NoError(t, err)

	type GetRuleRequestInfo struct {
		ContainerId string
		Request     cloudevents.GetRuleRequest
	}

	var requests []GetRuleRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
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

// IssueRoutingInfo tag="default" email="oci_events_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/CLEV" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/CLEV"
func TestEventsClientListRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudevents", "ListRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudevents", "Events", "ListRules", createEventsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudevents.EventsClient)

	body, err := testClient.getRequests("cloudevents", "ListRules")
	assert.NoError(t, err)

	type ListRulesRequestInfo struct {
		ContainerId string
		Request     cloudevents.ListRulesRequest
	}

	var requests []ListRulesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
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

// IssueRoutingInfo tag="default" email="oci_events_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/CLEV" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/CLEV"
func TestEventsClientUpdateRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cloudevents", "UpdateRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cloudevents", "Events", "UpdateRule", createEventsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cloudevents.EventsClient)

	body, err := testClient.getRequests("cloudevents", "UpdateRule")
	assert.NoError(t, err)

	type UpdateRuleRequestInfo struct {
		ContainerId string
		Request     cloudevents.UpdateRuleRequest
	}

	var requests []UpdateRuleRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
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
