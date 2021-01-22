package autotest

import (
	"github.com/erikcai/oci-go-sdk/v33/apmsynthetics"
	"github.com/erikcai/oci-go-sdk/v33/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createApmsyntheticsApmSyntheticClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := apmsynthetics.NewApmSyntheticClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="rchandok_org_ww@oracle.com" jiraProject="APM" opsJiraProject="APMSDC"
func TestApmsyntheticsApmSyntheticClientCreateMonitor(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apmsynthetics", "CreateMonitor")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateMonitor is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apmsynthetics", "ApmSynthetic", "CreateMonitor", createApmsyntheticsApmSyntheticClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apmsynthetics.ApmSyntheticClient)

	body, err := testClient.getRequests("apmsynthetics", "CreateMonitor")
	assert.NoError(t, err)

	type CreateMonitorRequestInfo struct {
		ContainerId string
		Request     apmsynthetics.CreateMonitorRequest
	}

	var requests []CreateMonitorRequestInfo
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
			response, err := c.CreateMonitor(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="rchandok_org_ww@oracle.com" jiraProject="APM" opsJiraProject="APMSDC"
func TestApmsyntheticsApmSyntheticClientCreateScript(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apmsynthetics", "CreateScript")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateScript is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apmsynthetics", "ApmSynthetic", "CreateScript", createApmsyntheticsApmSyntheticClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apmsynthetics.ApmSyntheticClient)

	body, err := testClient.getRequests("apmsynthetics", "CreateScript")
	assert.NoError(t, err)

	type CreateScriptRequestInfo struct {
		ContainerId string
		Request     apmsynthetics.CreateScriptRequest
	}

	var requests []CreateScriptRequestInfo
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
			response, err := c.CreateScript(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="rchandok_org_ww@oracle.com" jiraProject="APM" opsJiraProject="APMSDC"
func TestApmsyntheticsApmSyntheticClientDeleteMonitor(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apmsynthetics", "DeleteMonitor")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteMonitor is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apmsynthetics", "ApmSynthetic", "DeleteMonitor", createApmsyntheticsApmSyntheticClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apmsynthetics.ApmSyntheticClient)

	body, err := testClient.getRequests("apmsynthetics", "DeleteMonitor")
	assert.NoError(t, err)

	type DeleteMonitorRequestInfo struct {
		ContainerId string
		Request     apmsynthetics.DeleteMonitorRequest
	}

	var requests []DeleteMonitorRequestInfo
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
			response, err := c.DeleteMonitor(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="rchandok_org_ww@oracle.com" jiraProject="APM" opsJiraProject="APMSDC"
func TestApmsyntheticsApmSyntheticClientDeleteScript(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apmsynthetics", "DeleteScript")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteScript is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apmsynthetics", "ApmSynthetic", "DeleteScript", createApmsyntheticsApmSyntheticClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apmsynthetics.ApmSyntheticClient)

	body, err := testClient.getRequests("apmsynthetics", "DeleteScript")
	assert.NoError(t, err)

	type DeleteScriptRequestInfo struct {
		ContainerId string
		Request     apmsynthetics.DeleteScriptRequest
	}

	var requests []DeleteScriptRequestInfo
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
			response, err := c.DeleteScript(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="rchandok_org_ww@oracle.com" jiraProject="APM" opsJiraProject="APMSDC"
func TestApmsyntheticsApmSyntheticClientGetMonitor(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apmsynthetics", "GetMonitor")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetMonitor is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apmsynthetics", "ApmSynthetic", "GetMonitor", createApmsyntheticsApmSyntheticClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apmsynthetics.ApmSyntheticClient)

	body, err := testClient.getRequests("apmsynthetics", "GetMonitor")
	assert.NoError(t, err)

	type GetMonitorRequestInfo struct {
		ContainerId string
		Request     apmsynthetics.GetMonitorRequest
	}

	var requests []GetMonitorRequestInfo
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
			response, err := c.GetMonitor(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="rchandok_org_ww@oracle.com" jiraProject="APM" opsJiraProject="APMSDC"
func TestApmsyntheticsApmSyntheticClientGetMonitorResult(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apmsynthetics", "GetMonitorResult")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetMonitorResult is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apmsynthetics", "ApmSynthetic", "GetMonitorResult", createApmsyntheticsApmSyntheticClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apmsynthetics.ApmSyntheticClient)

	body, err := testClient.getRequests("apmsynthetics", "GetMonitorResult")
	assert.NoError(t, err)

	type GetMonitorResultRequestInfo struct {
		ContainerId string
		Request     apmsynthetics.GetMonitorResultRequest
	}

	var requests []GetMonitorResultRequestInfo
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
			response, err := c.GetMonitorResult(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="rchandok_org_ww@oracle.com" jiraProject="APM" opsJiraProject="APMSDC"
func TestApmsyntheticsApmSyntheticClientGetScript(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apmsynthetics", "GetScript")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetScript is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apmsynthetics", "ApmSynthetic", "GetScript", createApmsyntheticsApmSyntheticClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apmsynthetics.ApmSyntheticClient)

	body, err := testClient.getRequests("apmsynthetics", "GetScript")
	assert.NoError(t, err)

	type GetScriptRequestInfo struct {
		ContainerId string
		Request     apmsynthetics.GetScriptRequest
	}

	var requests []GetScriptRequestInfo
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
			response, err := c.GetScript(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="rchandok_org_ww@oracle.com" jiraProject="APM" opsJiraProject="APMSDC"
func TestApmsyntheticsApmSyntheticClientListMonitors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apmsynthetics", "ListMonitors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListMonitors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apmsynthetics", "ApmSynthetic", "ListMonitors", createApmsyntheticsApmSyntheticClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apmsynthetics.ApmSyntheticClient)

	body, err := testClient.getRequests("apmsynthetics", "ListMonitors")
	assert.NoError(t, err)

	type ListMonitorsRequestInfo struct {
		ContainerId string
		Request     apmsynthetics.ListMonitorsRequest
	}

	var requests []ListMonitorsRequestInfo
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
				r := req.(*apmsynthetics.ListMonitorsRequest)
				return c.ListMonitors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]apmsynthetics.ListMonitorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(apmsynthetics.ListMonitorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="rchandok_org_ww@oracle.com" jiraProject="APM" opsJiraProject="APMSDC"
func TestApmsyntheticsApmSyntheticClientListPublicVantagePoints(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apmsynthetics", "ListPublicVantagePoints")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPublicVantagePoints is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apmsynthetics", "ApmSynthetic", "ListPublicVantagePoints", createApmsyntheticsApmSyntheticClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apmsynthetics.ApmSyntheticClient)

	body, err := testClient.getRequests("apmsynthetics", "ListPublicVantagePoints")
	assert.NoError(t, err)

	type ListPublicVantagePointsRequestInfo struct {
		ContainerId string
		Request     apmsynthetics.ListPublicVantagePointsRequest
	}

	var requests []ListPublicVantagePointsRequestInfo
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
				r := req.(*apmsynthetics.ListPublicVantagePointsRequest)
				return c.ListPublicVantagePoints(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]apmsynthetics.ListPublicVantagePointsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(apmsynthetics.ListPublicVantagePointsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="rchandok_org_ww@oracle.com" jiraProject="APM" opsJiraProject="APMSDC"
func TestApmsyntheticsApmSyntheticClientListScripts(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apmsynthetics", "ListScripts")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListScripts is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apmsynthetics", "ApmSynthetic", "ListScripts", createApmsyntheticsApmSyntheticClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apmsynthetics.ApmSyntheticClient)

	body, err := testClient.getRequests("apmsynthetics", "ListScripts")
	assert.NoError(t, err)

	type ListScriptsRequestInfo struct {
		ContainerId string
		Request     apmsynthetics.ListScriptsRequest
	}

	var requests []ListScriptsRequestInfo
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
				r := req.(*apmsynthetics.ListScriptsRequest)
				return c.ListScripts(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]apmsynthetics.ListScriptsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(apmsynthetics.ListScriptsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="rchandok_org_ww@oracle.com" jiraProject="APM" opsJiraProject="APMSDC"
func TestApmsyntheticsApmSyntheticClientUpdateMonitor(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apmsynthetics", "UpdateMonitor")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateMonitor is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apmsynthetics", "ApmSynthetic", "UpdateMonitor", createApmsyntheticsApmSyntheticClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apmsynthetics.ApmSyntheticClient)

	body, err := testClient.getRequests("apmsynthetics", "UpdateMonitor")
	assert.NoError(t, err)

	type UpdateMonitorRequestInfo struct {
		ContainerId string
		Request     apmsynthetics.UpdateMonitorRequest
	}

	var requests []UpdateMonitorRequestInfo
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
			response, err := c.UpdateMonitor(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="rchandok_org_ww@oracle.com" jiraProject="APM" opsJiraProject="APMSDC"
func TestApmsyntheticsApmSyntheticClientUpdateScript(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apmsynthetics", "UpdateScript")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateScript is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apmsynthetics", "ApmSynthetic", "UpdateScript", createApmsyntheticsApmSyntheticClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apmsynthetics.ApmSyntheticClient)

	body, err := testClient.getRequests("apmsynthetics", "UpdateScript")
	assert.NoError(t, err)

	type UpdateScriptRequestInfo struct {
		ContainerId string
		Request     apmsynthetics.UpdateScriptRequest
	}

	var requests []UpdateScriptRequestInfo
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
			response, err := c.UpdateScript(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
