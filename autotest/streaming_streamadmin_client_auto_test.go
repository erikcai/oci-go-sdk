package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/streaming"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createStreamAdminClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := streaming.NewStreamAdminClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientChangeConnectHarnessCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "ChangeConnectHarnessCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeConnectHarnessCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "ChangeConnectHarnessCompartment", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "ChangeConnectHarnessCompartment")
	assert.NoError(t, err)

	type ChangeConnectHarnessCompartmentRequestInfo struct {
		ContainerId string
		Request     streaming.ChangeConnectHarnessCompartmentRequest
	}

	var requests []ChangeConnectHarnessCompartmentRequestInfo
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
			response, err := c.ChangeConnectHarnessCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientChangeStreamCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "ChangeStreamCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeStreamCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "ChangeStreamCompartment", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "ChangeStreamCompartment")
	assert.NoError(t, err)

	type ChangeStreamCompartmentRequestInfo struct {
		ContainerId string
		Request     streaming.ChangeStreamCompartmentRequest
	}

	var requests []ChangeStreamCompartmentRequestInfo
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
			response, err := c.ChangeStreamCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientChangeStreamPoolCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "ChangeStreamPoolCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeStreamPoolCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "ChangeStreamPoolCompartment", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "ChangeStreamPoolCompartment")
	assert.NoError(t, err)

	type ChangeStreamPoolCompartmentRequestInfo struct {
		ContainerId string
		Request     streaming.ChangeStreamPoolCompartmentRequest
	}

	var requests []ChangeStreamPoolCompartmentRequestInfo
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
			response, err := c.ChangeStreamPoolCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientCreateConnectHarness(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "CreateConnectHarness")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateConnectHarness is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "CreateConnectHarness", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "CreateConnectHarness")
	assert.NoError(t, err)

	type CreateConnectHarnessRequestInfo struct {
		ContainerId string
		Request     streaming.CreateConnectHarnessRequest
	}

	var requests []CreateConnectHarnessRequestInfo
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
			response, err := c.CreateConnectHarness(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientCreateStream(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "CreateStream")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateStream is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "CreateStream", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "CreateStream")
	assert.NoError(t, err)

	type CreateStreamRequestInfo struct {
		ContainerId string
		Request     streaming.CreateStreamRequest
	}

	var requests []CreateStreamRequestInfo
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
			response, err := c.CreateStream(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientCreateStreamPool(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "CreateStreamPool")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateStreamPool is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "CreateStreamPool", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "CreateStreamPool")
	assert.NoError(t, err)

	type CreateStreamPoolRequestInfo struct {
		ContainerId string
		Request     streaming.CreateStreamPoolRequest
	}

	var requests []CreateStreamPoolRequestInfo
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
			response, err := c.CreateStreamPool(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientDeleteConnectHarness(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "DeleteConnectHarness")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteConnectHarness is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "DeleteConnectHarness", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "DeleteConnectHarness")
	assert.NoError(t, err)

	type DeleteConnectHarnessRequestInfo struct {
		ContainerId string
		Request     streaming.DeleteConnectHarnessRequest
	}

	var requests []DeleteConnectHarnessRequestInfo
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
			response, err := c.DeleteConnectHarness(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientDeleteStream(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "DeleteStream")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteStream is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "DeleteStream", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "DeleteStream")
	assert.NoError(t, err)

	type DeleteStreamRequestInfo struct {
		ContainerId string
		Request     streaming.DeleteStreamRequest
	}

	var requests []DeleteStreamRequestInfo
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
			response, err := c.DeleteStream(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientDeleteStreamPool(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "DeleteStreamPool")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteStreamPool is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "DeleteStreamPool", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "DeleteStreamPool")
	assert.NoError(t, err)

	type DeleteStreamPoolRequestInfo struct {
		ContainerId string
		Request     streaming.DeleteStreamPoolRequest
	}

	var requests []DeleteStreamPoolRequestInfo
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
			response, err := c.DeleteStreamPool(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientGetConnectHarness(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "GetConnectHarness")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConnectHarness is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "GetConnectHarness", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "GetConnectHarness")
	assert.NoError(t, err)

	type GetConnectHarnessRequestInfo struct {
		ContainerId string
		Request     streaming.GetConnectHarnessRequest
	}

	var requests []GetConnectHarnessRequestInfo
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
			response, err := c.GetConnectHarness(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientGetStream(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "GetStream")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStream is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "GetStream", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "GetStream")
	assert.NoError(t, err)

	type GetStreamRequestInfo struct {
		ContainerId string
		Request     streaming.GetStreamRequest
	}

	var requests []GetStreamRequestInfo
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
			response, err := c.GetStream(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientGetStreamPool(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "GetStreamPool")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStreamPool is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "GetStreamPool", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "GetStreamPool")
	assert.NoError(t, err)

	type GetStreamPoolRequestInfo struct {
		ContainerId string
		Request     streaming.GetStreamPoolRequest
	}

	var requests []GetStreamPoolRequestInfo
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
			response, err := c.GetStreamPool(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientListConnectHarnesses(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "ListConnectHarnesses")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListConnectHarnesses is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "ListConnectHarnesses", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "ListConnectHarnesses")
	assert.NoError(t, err)

	type ListConnectHarnessesRequestInfo struct {
		ContainerId string
		Request     streaming.ListConnectHarnessesRequest
	}

	var requests []ListConnectHarnessesRequestInfo
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
				r := req.(*streaming.ListConnectHarnessesRequest)
				return c.ListConnectHarnesses(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]streaming.ListConnectHarnessesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(streaming.ListConnectHarnessesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientListStreamPools(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "ListStreamPools")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListStreamPools is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "ListStreamPools", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "ListStreamPools")
	assert.NoError(t, err)

	type ListStreamPoolsRequestInfo struct {
		ContainerId string
		Request     streaming.ListStreamPoolsRequest
	}

	var requests []ListStreamPoolsRequestInfo
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
				r := req.(*streaming.ListStreamPoolsRequest)
				return c.ListStreamPools(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]streaming.ListStreamPoolsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(streaming.ListStreamPoolsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientListStreams(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "ListStreams")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListStreams is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "ListStreams", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "ListStreams")
	assert.NoError(t, err)

	type ListStreamsRequestInfo struct {
		ContainerId string
		Request     streaming.ListStreamsRequest
	}

	var requests []ListStreamsRequestInfo
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
				r := req.(*streaming.ListStreamsRequest)
				return c.ListStreams(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]streaming.ListStreamsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(streaming.ListStreamsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientUpdateConnectHarness(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "UpdateConnectHarness")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateConnectHarness is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "UpdateConnectHarness", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "UpdateConnectHarness")
	assert.NoError(t, err)

	type UpdateConnectHarnessRequestInfo struct {
		ContainerId string
		Request     streaming.UpdateConnectHarnessRequest
	}

	var requests []UpdateConnectHarnessRequestInfo
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
			response, err := c.UpdateConnectHarness(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientUpdateStream(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "UpdateStream")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateStream is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "UpdateStream", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "UpdateStream")
	assert.NoError(t, err)

	type UpdateStreamRequestInfo struct {
		ContainerId string
		Request     streaming.UpdateStreamRequest
	}

	var requests []UpdateStreamRequestInfo
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
			response, err := c.UpdateStream(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_streaming_us_grp@oracle.com" jiraProject="STREAMSTR" opsJiraProject="STREAMOSS"
func TestStreamAdminClientUpdateStreamPool(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "UpdateStreamPool")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateStreamPool is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "UpdateStreamPool", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "UpdateStreamPool")
	assert.NoError(t, err)

	type UpdateStreamPoolRequestInfo struct {
		ContainerId string
		Request     streaming.UpdateStreamPoolRequest
	}

	var requests []UpdateStreamPoolRequestInfo
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
			response, err := c.UpdateStreamPool(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
