package autotest

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"github.com/oracle/oci-go-sdk/v31/sch"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createSchServiceConnectorClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := sch.NewServiceConnectorClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="och_us-grp@oracle.com" jiraProject="OCH" opsJiraProject="OCH"
func TestSchServiceConnectorClientActivateServiceConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("sch", "ActivateServiceConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ActivateServiceConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("sch", "ServiceConnector", "ActivateServiceConnector", createSchServiceConnectorClientWithProvider)
	assert.NoError(t, err)
	c := cc.(sch.ServiceConnectorClient)

	body, err := testClient.getRequests("sch", "ActivateServiceConnector")
	assert.NoError(t, err)

	type ActivateServiceConnectorRequestInfo struct {
		ContainerId string
		Request     sch.ActivateServiceConnectorRequest
	}

	var requests []ActivateServiceConnectorRequestInfo
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
			response, err := c.ActivateServiceConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="och_us-grp@oracle.com" jiraProject="OCH" opsJiraProject="OCH"
func TestSchServiceConnectorClientChangeServiceConnectorCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("sch", "ChangeServiceConnectorCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeServiceConnectorCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("sch", "ServiceConnector", "ChangeServiceConnectorCompartment", createSchServiceConnectorClientWithProvider)
	assert.NoError(t, err)
	c := cc.(sch.ServiceConnectorClient)

	body, err := testClient.getRequests("sch", "ChangeServiceConnectorCompartment")
	assert.NoError(t, err)

	type ChangeServiceConnectorCompartmentRequestInfo struct {
		ContainerId string
		Request     sch.ChangeServiceConnectorCompartmentRequest
	}

	var requests []ChangeServiceConnectorCompartmentRequestInfo
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
			response, err := c.ChangeServiceConnectorCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="och_us-grp@oracle.com" jiraProject="OCH" opsJiraProject="OCH"
func TestSchServiceConnectorClientCreateServiceConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("sch", "CreateServiceConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateServiceConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("sch", "ServiceConnector", "CreateServiceConnector", createSchServiceConnectorClientWithProvider)
	assert.NoError(t, err)
	c := cc.(sch.ServiceConnectorClient)

	body, err := testClient.getRequests("sch", "CreateServiceConnector")
	assert.NoError(t, err)

	type CreateServiceConnectorRequestInfo struct {
		ContainerId string
		Request     sch.CreateServiceConnectorRequest
	}

	var requests []CreateServiceConnectorRequestInfo
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
			response, err := c.CreateServiceConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="och_us-grp@oracle.com" jiraProject="OCH" opsJiraProject="OCH"
func TestSchServiceConnectorClientDeactivateServiceConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("sch", "DeactivateServiceConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeactivateServiceConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("sch", "ServiceConnector", "DeactivateServiceConnector", createSchServiceConnectorClientWithProvider)
	assert.NoError(t, err)
	c := cc.(sch.ServiceConnectorClient)

	body, err := testClient.getRequests("sch", "DeactivateServiceConnector")
	assert.NoError(t, err)

	type DeactivateServiceConnectorRequestInfo struct {
		ContainerId string
		Request     sch.DeactivateServiceConnectorRequest
	}

	var requests []DeactivateServiceConnectorRequestInfo
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
			response, err := c.DeactivateServiceConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="och_us-grp@oracle.com" jiraProject="OCH" opsJiraProject="OCH"
func TestSchServiceConnectorClientDeleteServiceConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("sch", "DeleteServiceConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteServiceConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("sch", "ServiceConnector", "DeleteServiceConnector", createSchServiceConnectorClientWithProvider)
	assert.NoError(t, err)
	c := cc.(sch.ServiceConnectorClient)

	body, err := testClient.getRequests("sch", "DeleteServiceConnector")
	assert.NoError(t, err)

	type DeleteServiceConnectorRequestInfo struct {
		ContainerId string
		Request     sch.DeleteServiceConnectorRequest
	}

	var requests []DeleteServiceConnectorRequestInfo
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
			response, err := c.DeleteServiceConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="och_us-grp@oracle.com" jiraProject="OCH" opsJiraProject="OCH"
func TestSchServiceConnectorClientGetServiceConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("sch", "GetServiceConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetServiceConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("sch", "ServiceConnector", "GetServiceConnector", createSchServiceConnectorClientWithProvider)
	assert.NoError(t, err)
	c := cc.(sch.ServiceConnectorClient)

	body, err := testClient.getRequests("sch", "GetServiceConnector")
	assert.NoError(t, err)

	type GetServiceConnectorRequestInfo struct {
		ContainerId string
		Request     sch.GetServiceConnectorRequest
	}

	var requests []GetServiceConnectorRequestInfo
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
			response, err := c.GetServiceConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="och_us-grp@oracle.com" jiraProject="OCH" opsJiraProject="OCH"
func TestSchServiceConnectorClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("sch", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("sch", "ServiceConnector", "GetWorkRequest", createSchServiceConnectorClientWithProvider)
	assert.NoError(t, err)
	c := cc.(sch.ServiceConnectorClient)

	body, err := testClient.getRequests("sch", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     sch.GetWorkRequestRequest
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

// IssueRoutingInfo tag="default" email="och_us-grp@oracle.com" jiraProject="OCH" opsJiraProject="OCH"
func TestSchServiceConnectorClientListServiceConnectors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("sch", "ListServiceConnectors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListServiceConnectors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("sch", "ServiceConnector", "ListServiceConnectors", createSchServiceConnectorClientWithProvider)
	assert.NoError(t, err)
	c := cc.(sch.ServiceConnectorClient)

	body, err := testClient.getRequests("sch", "ListServiceConnectors")
	assert.NoError(t, err)

	type ListServiceConnectorsRequestInfo struct {
		ContainerId string
		Request     sch.ListServiceConnectorsRequest
	}

	var requests []ListServiceConnectorsRequestInfo
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
				r := req.(*sch.ListServiceConnectorsRequest)
				return c.ListServiceConnectors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]sch.ListServiceConnectorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(sch.ListServiceConnectorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="och_us-grp@oracle.com" jiraProject="OCH" opsJiraProject="OCH"
func TestSchServiceConnectorClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("sch", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("sch", "ServiceConnector", "ListWorkRequestErrors", createSchServiceConnectorClientWithProvider)
	assert.NoError(t, err)
	c := cc.(sch.ServiceConnectorClient)

	body, err := testClient.getRequests("sch", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     sch.ListWorkRequestErrorsRequest
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
				r := req.(*sch.ListWorkRequestErrorsRequest)
				return c.ListWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]sch.ListWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(sch.ListWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="och_us-grp@oracle.com" jiraProject="OCH" opsJiraProject="OCH"
func TestSchServiceConnectorClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("sch", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("sch", "ServiceConnector", "ListWorkRequestLogs", createSchServiceConnectorClientWithProvider)
	assert.NoError(t, err)
	c := cc.(sch.ServiceConnectorClient)

	body, err := testClient.getRequests("sch", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     sch.ListWorkRequestLogsRequest
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
				r := req.(*sch.ListWorkRequestLogsRequest)
				return c.ListWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]sch.ListWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(sch.ListWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="och_us-grp@oracle.com" jiraProject="OCH" opsJiraProject="OCH"
func TestSchServiceConnectorClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("sch", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("sch", "ServiceConnector", "ListWorkRequests", createSchServiceConnectorClientWithProvider)
	assert.NoError(t, err)
	c := cc.(sch.ServiceConnectorClient)

	body, err := testClient.getRequests("sch", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     sch.ListWorkRequestsRequest
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
				r := req.(*sch.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]sch.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(sch.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="och_us-grp@oracle.com" jiraProject="OCH" opsJiraProject="OCH"
func TestSchServiceConnectorClientUpdateServiceConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("sch", "UpdateServiceConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateServiceConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("sch", "ServiceConnector", "UpdateServiceConnector", createSchServiceConnectorClientWithProvider)
	assert.NoError(t, err)
	c := cc.(sch.ServiceConnectorClient)

	body, err := testClient.getRequests("sch", "UpdateServiceConnector")
	assert.NoError(t, err)

	type UpdateServiceConnectorRequestInfo struct {
		ContainerId string
		Request     sch.UpdateServiceConnectorRequest
	}

	var requests []UpdateServiceConnectorRequestInfo
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
			response, err := c.UpdateServiceConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
