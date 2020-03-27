package autotest

import (
	"github.com/oracle/oci-go-sdk/apigateway"
	"github.com/oracle/oci-go-sdk/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createApiGatewayClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := apigateway.NewApiGatewayClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_apigw_ww_grp@oracle.com" jiraProject="APIGW" opsJiraProject="APIGW"
func TestApiGatewayClientChangeCertificateCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "ChangeCertificateCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeCertificateCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "ChangeCertificateCompartment", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "ChangeCertificateCompartment")
	assert.NoError(t, err)

	type ChangeCertificateCompartmentRequestInfo struct {
		ContainerId string
		Request     apigateway.ChangeCertificateCompartmentRequest
	}

	var requests []ChangeCertificateCompartmentRequestInfo
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
			response, err := c.ChangeCertificateCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_apigw_ww_grp@oracle.com" jiraProject="APIGW" opsJiraProject="APIGW"
func TestApiGatewayClientCreateCertificate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "CreateCertificate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCertificate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "CreateCertificate", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "CreateCertificate")
	assert.NoError(t, err)

	type CreateCertificateRequestInfo struct {
		ContainerId string
		Request     apigateway.CreateCertificateRequest
	}

	var requests []CreateCertificateRequestInfo
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
			response, err := c.CreateCertificate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_apigw_ww_grp@oracle.com" jiraProject="APIGW" opsJiraProject="APIGW"
func TestApiGatewayClientDeleteCertificate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "DeleteCertificate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCertificate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "DeleteCertificate", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "DeleteCertificate")
	assert.NoError(t, err)

	type DeleteCertificateRequestInfo struct {
		ContainerId string
		Request     apigateway.DeleteCertificateRequest
	}

	var requests []DeleteCertificateRequestInfo
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
			response, err := c.DeleteCertificate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_apigw_ww_grp@oracle.com" jiraProject="APIGW" opsJiraProject="APIGW"
func TestApiGatewayClientGetCertificate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "GetCertificate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCertificate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "GetCertificate", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "GetCertificate")
	assert.NoError(t, err)

	type GetCertificateRequestInfo struct {
		ContainerId string
		Request     apigateway.GetCertificateRequest
	}

	var requests []GetCertificateRequestInfo
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
			response, err := c.GetCertificate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_apigw_ww_grp@oracle.com" jiraProject="APIGW" opsJiraProject="APIGW"
func TestApiGatewayClientListCertificates(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "ListCertificates")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCertificates is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "ListCertificates", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "ListCertificates")
	assert.NoError(t, err)

	type ListCertificatesRequestInfo struct {
		ContainerId string
		Request     apigateway.ListCertificatesRequest
	}

	var requests []ListCertificatesRequestInfo
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
				r := req.(*apigateway.ListCertificatesRequest)
				return c.ListCertificates(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]apigateway.ListCertificatesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(apigateway.ListCertificatesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_apigw_ww_grp@oracle.com" jiraProject="APIGW" opsJiraProject="APIGW"
func TestApiGatewayClientUpdateCertificate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "UpdateCertificate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCertificate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "UpdateCertificate", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "UpdateCertificate")
	assert.NoError(t, err)

	type UpdateCertificateRequestInfo struct {
		ContainerId string
		Request     apigateway.UpdateCertificateRequest
	}

	var requests []UpdateCertificateRequestInfo
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
			response, err := c.UpdateCertificate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
