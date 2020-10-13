package autotest

import (
	"github.com/oracle/oci-go-sdk/v27/apigateway"
	"github.com/oracle/oci-go-sdk/v27/common"

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
func TestApiGatewayClientChangeApiCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "ChangeApiCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeApiCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "ChangeApiCompartment", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "ChangeApiCompartment")
	assert.NoError(t, err)

	type ChangeApiCompartmentRequestInfo struct {
		ContainerId string
		Request     apigateway.ChangeApiCompartmentRequest
	}

	var requests []ChangeApiCompartmentRequestInfo
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
			response, err := c.ChangeApiCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
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
func TestApiGatewayClientCreateApi(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "CreateApi")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateApi is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "CreateApi", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "CreateApi")
	assert.NoError(t, err)

	type CreateApiRequestInfo struct {
		ContainerId string
		Request     apigateway.CreateApiRequest
	}

	var requests []CreateApiRequestInfo
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
			response, err := c.CreateApi(context.Background(), req.Request)
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
func TestApiGatewayClientDeleteApi(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "DeleteApi")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteApi is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "DeleteApi", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "DeleteApi")
	assert.NoError(t, err)

	type DeleteApiRequestInfo struct {
		ContainerId string
		Request     apigateway.DeleteApiRequest
	}

	var requests []DeleteApiRequestInfo
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
			response, err := c.DeleteApi(context.Background(), req.Request)
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
func TestApiGatewayClientGetApi(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "GetApi")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetApi is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "GetApi", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "GetApi")
	assert.NoError(t, err)

	type GetApiRequestInfo struct {
		ContainerId string
		Request     apigateway.GetApiRequest
	}

	var requests []GetApiRequestInfo
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
			response, err := c.GetApi(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_apigw_ww_grp@oracle.com" jiraProject="APIGW" opsJiraProject="APIGW"
func TestApiGatewayClientGetApiContent(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "GetApiContent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetApiContent is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "GetApiContent", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "GetApiContent")
	assert.NoError(t, err)

	type GetApiContentRequestInfo struct {
		ContainerId string
		Request     apigateway.GetApiContentRequest
	}

	var requests []GetApiContentRequestInfo
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
			response, err := c.GetApiContent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_apigw_ww_grp@oracle.com" jiraProject="APIGW" opsJiraProject="APIGW"
func TestApiGatewayClientGetApiDeploymentSpecification(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "GetApiDeploymentSpecification")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetApiDeploymentSpecification is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "GetApiDeploymentSpecification", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "GetApiDeploymentSpecification")
	assert.NoError(t, err)

	type GetApiDeploymentSpecificationRequestInfo struct {
		ContainerId string
		Request     apigateway.GetApiDeploymentSpecificationRequest
	}

	var requests []GetApiDeploymentSpecificationRequestInfo
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
			response, err := c.GetApiDeploymentSpecification(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_apigw_ww_grp@oracle.com" jiraProject="APIGW" opsJiraProject="APIGW"
func TestApiGatewayClientGetApiValidations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "GetApiValidations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetApiValidations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "GetApiValidations", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "GetApiValidations")
	assert.NoError(t, err)

	type GetApiValidationsRequestInfo struct {
		ContainerId string
		Request     apigateway.GetApiValidationsRequest
	}

	var requests []GetApiValidationsRequestInfo
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
			response, err := c.GetApiValidations(context.Background(), req.Request)
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
func TestApiGatewayClientListApis(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "ListApis")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListApis is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "ListApis", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "ListApis")
	assert.NoError(t, err)

	type ListApisRequestInfo struct {
		ContainerId string
		Request     apigateway.ListApisRequest
	}

	var requests []ListApisRequestInfo
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
				r := req.(*apigateway.ListApisRequest)
				return c.ListApis(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]apigateway.ListApisResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(apigateway.ListApisResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
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
func TestApiGatewayClientUpdateApi(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("apigateway", "UpdateApi")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateApi is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("apigateway", "ApiGateway", "UpdateApi", createApiGatewayClientWithProvider)
	assert.NoError(t, err)
	c := cc.(apigateway.ApiGatewayClient)

	body, err := testClient.getRequests("apigateway", "UpdateApi")
	assert.NoError(t, err)

	type UpdateApiRequestInfo struct {
		ContainerId string
		Request     apigateway.UpdateApiRequest
	}

	var requests []UpdateApiRequestInfo
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
			response, err := c.UpdateApi(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
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
