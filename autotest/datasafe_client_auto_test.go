package autotest

import (
	"github.com/oracle/oci-go-sdk/v32/common"
	"github.com/oracle/oci-go-sdk/v32/datasafe"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createDataSafeClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := datasafe.NewDataSafeClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientChangeDataSafePrivateEndpointCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "ChangeDataSafePrivateEndpointCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeDataSafePrivateEndpointCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "ChangeDataSafePrivateEndpointCompartment", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "ChangeDataSafePrivateEndpointCompartment")
	assert.NoError(t, err)

	type ChangeDataSafePrivateEndpointCompartmentRequestInfo struct {
		ContainerId string
		Request     datasafe.ChangeDataSafePrivateEndpointCompartmentRequest
	}

	var requests []ChangeDataSafePrivateEndpointCompartmentRequestInfo
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
			response, err := c.ChangeDataSafePrivateEndpointCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientChangeOnPremConnectorCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "ChangeOnPremConnectorCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeOnPremConnectorCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "ChangeOnPremConnectorCompartment", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "ChangeOnPremConnectorCompartment")
	assert.NoError(t, err)

	type ChangeOnPremConnectorCompartmentRequestInfo struct {
		ContainerId string
		Request     datasafe.ChangeOnPremConnectorCompartmentRequest
	}

	var requests []ChangeOnPremConnectorCompartmentRequestInfo
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
			response, err := c.ChangeOnPremConnectorCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientCreateDataSafePrivateEndpoint(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "CreateDataSafePrivateEndpoint")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDataSafePrivateEndpoint is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "CreateDataSafePrivateEndpoint", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "CreateDataSafePrivateEndpoint")
	assert.NoError(t, err)

	type CreateDataSafePrivateEndpointRequestInfo struct {
		ContainerId string
		Request     datasafe.CreateDataSafePrivateEndpointRequest
	}

	var requests []CreateDataSafePrivateEndpointRequestInfo
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
			response, err := c.CreateDataSafePrivateEndpoint(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientCreateOnPremConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "CreateOnPremConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateOnPremConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "CreateOnPremConnector", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "CreateOnPremConnector")
	assert.NoError(t, err)

	type CreateOnPremConnectorRequestInfo struct {
		ContainerId string
		Request     datasafe.CreateOnPremConnectorRequest
	}

	var requests []CreateOnPremConnectorRequestInfo
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
			response, err := c.CreateOnPremConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientDeleteDataSafePrivateEndpoint(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "DeleteDataSafePrivateEndpoint")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDataSafePrivateEndpoint is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "DeleteDataSafePrivateEndpoint", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "DeleteDataSafePrivateEndpoint")
	assert.NoError(t, err)

	type DeleteDataSafePrivateEndpointRequestInfo struct {
		ContainerId string
		Request     datasafe.DeleteDataSafePrivateEndpointRequest
	}

	var requests []DeleteDataSafePrivateEndpointRequestInfo
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
			response, err := c.DeleteDataSafePrivateEndpoint(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientDeleteOnPremConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "DeleteOnPremConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteOnPremConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "DeleteOnPremConnector", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "DeleteOnPremConnector")
	assert.NoError(t, err)

	type DeleteOnPremConnectorRequestInfo struct {
		ContainerId string
		Request     datasafe.DeleteOnPremConnectorRequest
	}

	var requests []DeleteOnPremConnectorRequestInfo
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
			response, err := c.DeleteOnPremConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientEnableDataSafeConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "EnableDataSafeConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("EnableDataSafeConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "EnableDataSafeConfiguration", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "EnableDataSafeConfiguration")
	assert.NoError(t, err)

	type EnableDataSafeConfigurationRequestInfo struct {
		ContainerId string
		Request     datasafe.EnableDataSafeConfigurationRequest
	}

	var requests []EnableDataSafeConfigurationRequestInfo
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
			response, err := c.EnableDataSafeConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientGenerateOnPremConnectorConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "GenerateOnPremConnectorConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GenerateOnPremConnectorConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "GenerateOnPremConnectorConfiguration", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "GenerateOnPremConnectorConfiguration")
	assert.NoError(t, err)

	type GenerateOnPremConnectorConfigurationRequestInfo struct {
		ContainerId string
		Request     datasafe.GenerateOnPremConnectorConfigurationRequest
	}

	var requests []GenerateOnPremConnectorConfigurationRequestInfo
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
			response, err := c.GenerateOnPremConnectorConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientGetDataSafeConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "GetDataSafeConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDataSafeConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "GetDataSafeConfiguration", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "GetDataSafeConfiguration")
	assert.NoError(t, err)

	type GetDataSafeConfigurationRequestInfo struct {
		ContainerId string
		Request     datasafe.GetDataSafeConfigurationRequest
	}

	var requests []GetDataSafeConfigurationRequestInfo
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
			response, err := c.GetDataSafeConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientGetDataSafePrivateEndpoint(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "GetDataSafePrivateEndpoint")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDataSafePrivateEndpoint is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "GetDataSafePrivateEndpoint", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "GetDataSafePrivateEndpoint")
	assert.NoError(t, err)

	type GetDataSafePrivateEndpointRequestInfo struct {
		ContainerId string
		Request     datasafe.GetDataSafePrivateEndpointRequest
	}

	var requests []GetDataSafePrivateEndpointRequestInfo
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
			response, err := c.GetDataSafePrivateEndpoint(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientGetOnPremConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "GetOnPremConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetOnPremConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "GetOnPremConnector", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "GetOnPremConnector")
	assert.NoError(t, err)

	type GetOnPremConnectorRequestInfo struct {
		ContainerId string
		Request     datasafe.GetOnPremConnectorRequest
	}

	var requests []GetOnPremConnectorRequestInfo
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
			response, err := c.GetOnPremConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "GetWorkRequest", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     datasafe.GetWorkRequestRequest
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

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientListDataSafePrivateEndpoints(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "ListDataSafePrivateEndpoints")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDataSafePrivateEndpoints is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "ListDataSafePrivateEndpoints", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "ListDataSafePrivateEndpoints")
	assert.NoError(t, err)

	type ListDataSafePrivateEndpointsRequestInfo struct {
		ContainerId string
		Request     datasafe.ListDataSafePrivateEndpointsRequest
	}

	var requests []ListDataSafePrivateEndpointsRequestInfo
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
				r := req.(*datasafe.ListDataSafePrivateEndpointsRequest)
				return c.ListDataSafePrivateEndpoints(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]datasafe.ListDataSafePrivateEndpointsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(datasafe.ListDataSafePrivateEndpointsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientListOnPremConnectors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "ListOnPremConnectors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListOnPremConnectors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "ListOnPremConnectors", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "ListOnPremConnectors")
	assert.NoError(t, err)

	type ListOnPremConnectorsRequestInfo struct {
		ContainerId string
		Request     datasafe.ListOnPremConnectorsRequest
	}

	var requests []ListOnPremConnectorsRequestInfo
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
				r := req.(*datasafe.ListOnPremConnectorsRequest)
				return c.ListOnPremConnectors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]datasafe.ListOnPremConnectorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(datasafe.ListOnPremConnectorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "ListWorkRequestErrors", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     datasafe.ListWorkRequestErrorsRequest
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
				r := req.(*datasafe.ListWorkRequestErrorsRequest)
				return c.ListWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]datasafe.ListWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(datasafe.ListWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "ListWorkRequestLogs", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     datasafe.ListWorkRequestLogsRequest
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
				r := req.(*datasafe.ListWorkRequestLogsRequest)
				return c.ListWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]datasafe.ListWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(datasafe.ListWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "ListWorkRequests", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     datasafe.ListWorkRequestsRequest
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
				r := req.(*datasafe.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]datasafe.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(datasafe.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientUpdateDataSafePrivateEndpoint(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "UpdateDataSafePrivateEndpoint")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDataSafePrivateEndpoint is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "UpdateDataSafePrivateEndpoint", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "UpdateDataSafePrivateEndpoint")
	assert.NoError(t, err)

	type UpdateDataSafePrivateEndpointRequestInfo struct {
		ContainerId string
		Request     datasafe.UpdateDataSafePrivateEndpointRequest
	}

	var requests []UpdateDataSafePrivateEndpointRequestInfo
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
			response, err := c.UpdateDataSafePrivateEndpoint(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientUpdateOnPremConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "UpdateOnPremConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateOnPremConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "UpdateOnPremConnector", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "UpdateOnPremConnector")
	assert.NoError(t, err)

	type UpdateOnPremConnectorRequestInfo struct {
		ContainerId string
		Request     datasafe.UpdateOnPremConnectorRequest
	}

	var requests []UpdateOnPremConnectorRequestInfo
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
			response, err := c.UpdateOnPremConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="datasafe_dex_ww_grp@oracle.com" jiraProject="DS" opsJiraProject="ADS"
func TestDataSafeClientUpdateOnPremConnectorWallet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("datasafe", "UpdateOnPremConnectorWallet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateOnPremConnectorWallet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("datasafe", "DataSafe", "UpdateOnPremConnectorWallet", createDataSafeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(datasafe.DataSafeClient)

	body, err := testClient.getRequests("datasafe", "UpdateOnPremConnectorWallet")
	assert.NoError(t, err)

	type UpdateOnPremConnectorWalletRequestInfo struct {
		ContainerId string
		Request     datasafe.UpdateOnPremConnectorWalletRequest
	}

	var requests []UpdateOnPremConnectorWalletRequestInfo
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
			response, err := c.UpdateOnPremConnectorWallet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
