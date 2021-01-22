package autotest

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"github.com/erikcai/oci-go-sdk/v33/jms"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createJmsJavaManagementServiceClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := jms.NewJavaManagementServiceClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientChangeFleetCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "ChangeFleetCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeFleetCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "ChangeFleetCompartment", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "ChangeFleetCompartment")
	assert.NoError(t, err)

	type ChangeFleetCompartmentRequestInfo struct {
		ContainerId string
		Request     jms.ChangeFleetCompartmentRequest
	}

	var requests []ChangeFleetCompartmentRequestInfo
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
			response, err := c.ChangeFleetCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientCreateFleet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "CreateFleet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateFleet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "CreateFleet", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "CreateFleet")
	assert.NoError(t, err)

	type CreateFleetRequestInfo struct {
		ContainerId string
		Request     jms.CreateFleetRequest
	}

	var requests []CreateFleetRequestInfo
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
			response, err := c.CreateFleet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientDeleteFleet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "DeleteFleet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteFleet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "DeleteFleet", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "DeleteFleet")
	assert.NoError(t, err)

	type DeleteFleetRequestInfo struct {
		ContainerId string
		Request     jms.DeleteFleetRequest
	}

	var requests []DeleteFleetRequestInfo
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
			response, err := c.DeleteFleet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientGetFleet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "GetFleet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetFleet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "GetFleet", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "GetFleet")
	assert.NoError(t, err)

	type GetFleetRequestInfo struct {
		ContainerId string
		Request     jms.GetFleetRequest
	}

	var requests []GetFleetRequestInfo
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
			response, err := c.GetFleet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "GetWorkRequest", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     jms.GetWorkRequestRequest
	}

	var requests []GetWorkRequestRequestInfo
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
			response, err := c.GetWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientListFleets(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "ListFleets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListFleets is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "ListFleets", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "ListFleets")
	assert.NoError(t, err)

	type ListFleetsRequestInfo struct {
		ContainerId string
		Request     jms.ListFleetsRequest
	}

	var requests []ListFleetsRequestInfo
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
				r := req.(*jms.ListFleetsRequest)
				return c.ListFleets(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]jms.ListFleetsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(jms.ListFleetsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "ListWorkRequestErrors", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     jms.ListWorkRequestErrorsRequest
	}

	var requests []ListWorkRequestErrorsRequestInfo
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
				r := req.(*jms.ListWorkRequestErrorsRequest)
				return c.ListWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]jms.ListWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(jms.ListWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "ListWorkRequestLogs", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     jms.ListWorkRequestLogsRequest
	}

	var requests []ListWorkRequestLogsRequestInfo
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
				r := req.(*jms.ListWorkRequestLogsRequest)
				return c.ListWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]jms.ListWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(jms.ListWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "ListWorkRequests", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     jms.ListWorkRequestsRequest
	}

	var requests []ListWorkRequestsRequestInfo
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
				r := req.(*jms.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]jms.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(jms.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientRequestSummarizedApplications(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "RequestSummarizedApplications")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSummarizedApplications is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "RequestSummarizedApplications", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "RequestSummarizedApplications")
	assert.NoError(t, err)

	type RequestSummarizedApplicationsRequestInfo struct {
		ContainerId string
		Request     jms.RequestSummarizedApplicationsRequest
	}

	var requests []RequestSummarizedApplicationsRequestInfo
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
				r := req.(*jms.RequestSummarizedApplicationsRequest)
				return c.RequestSummarizedApplications(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]jms.RequestSummarizedApplicationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(jms.RequestSummarizedApplicationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientRequestSummarizedInstallations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "RequestSummarizedInstallations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSummarizedInstallations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "RequestSummarizedInstallations", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "RequestSummarizedInstallations")
	assert.NoError(t, err)

	type RequestSummarizedInstallationsRequestInfo struct {
		ContainerId string
		Request     jms.RequestSummarizedInstallationsRequest
	}

	var requests []RequestSummarizedInstallationsRequestInfo
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
				r := req.(*jms.RequestSummarizedInstallationsRequest)
				return c.RequestSummarizedInstallations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]jms.RequestSummarizedInstallationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(jms.RequestSummarizedInstallationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientRequestSummarizedJres(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "RequestSummarizedJres")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSummarizedJres is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "RequestSummarizedJres", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "RequestSummarizedJres")
	assert.NoError(t, err)

	type RequestSummarizedJresRequestInfo struct {
		ContainerId string
		Request     jms.RequestSummarizedJresRequest
	}

	var requests []RequestSummarizedJresRequestInfo
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
				r := req.(*jms.RequestSummarizedJresRequest)
				return c.RequestSummarizedJres(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]jms.RequestSummarizedJresResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(jms.RequestSummarizedJresResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientRequestSummarizedManagedInstances(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "RequestSummarizedManagedInstances")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RequestSummarizedManagedInstances is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "RequestSummarizedManagedInstances", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "RequestSummarizedManagedInstances")
	assert.NoError(t, err)

	type RequestSummarizedManagedInstancesRequestInfo struct {
		ContainerId string
		Request     jms.RequestSummarizedManagedInstancesRequest
	}

	var requests []RequestSummarizedManagedInstancesRequestInfo
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
				r := req.(*jms.RequestSummarizedManagedInstancesRequest)
				return c.RequestSummarizedManagedInstances(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]jms.RequestSummarizedManagedInstancesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(jms.RequestSummarizedManagedInstancesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientSummarizeApplications(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "SummarizeApplications")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeApplications is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "SummarizeApplications", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "SummarizeApplications")
	assert.NoError(t, err)

	type SummarizeApplicationsRequestInfo struct {
		ContainerId string
		Request     jms.SummarizeApplicationsRequest
	}

	var requests []SummarizeApplicationsRequestInfo
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
				r := req.(*jms.SummarizeApplicationsRequest)
				return c.SummarizeApplications(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]jms.SummarizeApplicationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(jms.SummarizeApplicationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientSummarizeInstallations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "SummarizeInstallations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeInstallations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "SummarizeInstallations", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "SummarizeInstallations")
	assert.NoError(t, err)

	type SummarizeInstallationsRequestInfo struct {
		ContainerId string
		Request     jms.SummarizeInstallationsRequest
	}

	var requests []SummarizeInstallationsRequestInfo
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
				r := req.(*jms.SummarizeInstallationsRequest)
				return c.SummarizeInstallations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]jms.SummarizeInstallationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(jms.SummarizeInstallationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientSummarizeJres(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "SummarizeJres")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeJres is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "SummarizeJres", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "SummarizeJres")
	assert.NoError(t, err)

	type SummarizeJresRequestInfo struct {
		ContainerId string
		Request     jms.SummarizeJresRequest
	}

	var requests []SummarizeJresRequestInfo
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
				r := req.(*jms.SummarizeJresRequest)
				return c.SummarizeJres(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]jms.SummarizeJresResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(jms.SummarizeJresResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientSummarizeManagedInstances(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "SummarizeManagedInstances")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeManagedInstances is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "SummarizeManagedInstances", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "SummarizeManagedInstances")
	assert.NoError(t, err)

	type SummarizeManagedInstancesRequestInfo struct {
		ContainerId string
		Request     jms.SummarizeManagedInstancesRequest
	}

	var requests []SummarizeManagedInstancesRequestInfo
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
				r := req.(*jms.SummarizeManagedInstancesRequest)
				return c.SummarizeManagedInstances(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]jms.SummarizeManagedInstancesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(jms.SummarizeManagedInstancesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="autonomous_java-dev_us_grp@oracle.com" jiraProject="AJ" opsJiraProject="AJ"
func TestJmsJavaManagementServiceClientUpdateFleet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("jms", "UpdateFleet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateFleet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("jms", "JavaManagementService", "UpdateFleet", createJmsJavaManagementServiceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(jms.JavaManagementServiceClient)

	body, err := testClient.getRequests("jms", "UpdateFleet")
	assert.NoError(t, err)

	type UpdateFleetRequestInfo struct {
		ContainerId string
		Request     jms.UpdateFleetRequest
	}

	var requests []UpdateFleetRequestInfo
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
			response, err := c.UpdateFleet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
