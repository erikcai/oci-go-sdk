package autotest

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"github.com/erikcai/oci-go-sdk/v33/email"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createEmailClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := email.NewEmailClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="email-dev_us_grp@oracle.com" jiraProject="ED" opsJiraProject="ED"
func TestEmailClientChangeSenderCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("email", "ChangeSenderCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeSenderCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("email", "Email", "ChangeSenderCompartment", createEmailClientWithProvider)
	assert.NoError(t, err)
	c := cc.(email.EmailClient)

	body, err := testClient.getRequests("email", "ChangeSenderCompartment")
	assert.NoError(t, err)

	type ChangeSenderCompartmentRequestInfo struct {
		ContainerId string
		Request     email.ChangeSenderCompartmentRequest
	}

	var requests []ChangeSenderCompartmentRequestInfo
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
			response, err := c.ChangeSenderCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="email-dev_us_grp@oracle.com" jiraProject="ED" opsJiraProject="ED"
func TestEmailClientCreateSender(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("email", "CreateSender")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSender is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("email", "Email", "CreateSender", createEmailClientWithProvider)
	assert.NoError(t, err)
	c := cc.(email.EmailClient)

	body, err := testClient.getRequests("email", "CreateSender")
	assert.NoError(t, err)

	type CreateSenderRequestInfo struct {
		ContainerId string
		Request     email.CreateSenderRequest
	}

	var requests []CreateSenderRequestInfo
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
			response, err := c.CreateSender(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="email-dev_us_grp@oracle.com" jiraProject="ED" opsJiraProject="ED"
func TestEmailClientCreateSuppression(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("email", "CreateSuppression")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSuppression is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("email", "Email", "CreateSuppression", createEmailClientWithProvider)
	assert.NoError(t, err)
	c := cc.(email.EmailClient)

	body, err := testClient.getRequests("email", "CreateSuppression")
	assert.NoError(t, err)

	type CreateSuppressionRequestInfo struct {
		ContainerId string
		Request     email.CreateSuppressionRequest
	}

	var requests []CreateSuppressionRequestInfo
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
			response, err := c.CreateSuppression(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="email-dev_us_grp@oracle.com" jiraProject="ED" opsJiraProject="ED"
func TestEmailClientDeleteSender(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("email", "DeleteSender")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSender is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("email", "Email", "DeleteSender", createEmailClientWithProvider)
	assert.NoError(t, err)
	c := cc.(email.EmailClient)

	body, err := testClient.getRequests("email", "DeleteSender")
	assert.NoError(t, err)

	type DeleteSenderRequestInfo struct {
		ContainerId string
		Request     email.DeleteSenderRequest
	}

	var requests []DeleteSenderRequestInfo
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
			response, err := c.DeleteSender(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="email-dev_us_grp@oracle.com" jiraProject="ED" opsJiraProject="ED"
func TestEmailClientDeleteSuppression(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("email", "DeleteSuppression")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSuppression is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("email", "Email", "DeleteSuppression", createEmailClientWithProvider)
	assert.NoError(t, err)
	c := cc.(email.EmailClient)

	body, err := testClient.getRequests("email", "DeleteSuppression")
	assert.NoError(t, err)

	type DeleteSuppressionRequestInfo struct {
		ContainerId string
		Request     email.DeleteSuppressionRequest
	}

	var requests []DeleteSuppressionRequestInfo
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
			response, err := c.DeleteSuppression(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="email-dev_us_grp@oracle.com" jiraProject="ED" opsJiraProject="ED"
func TestEmailClientGetSender(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("email", "GetSender")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSender is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("email", "Email", "GetSender", createEmailClientWithProvider)
	assert.NoError(t, err)
	c := cc.(email.EmailClient)

	body, err := testClient.getRequests("email", "GetSender")
	assert.NoError(t, err)

	type GetSenderRequestInfo struct {
		ContainerId string
		Request     email.GetSenderRequest
	}

	var requests []GetSenderRequestInfo
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
			response, err := c.GetSender(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="email-dev_us_grp@oracle.com" jiraProject="ED" opsJiraProject="ED"
func TestEmailClientGetSuppression(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("email", "GetSuppression")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSuppression is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("email", "Email", "GetSuppression", createEmailClientWithProvider)
	assert.NoError(t, err)
	c := cc.(email.EmailClient)

	body, err := testClient.getRequests("email", "GetSuppression")
	assert.NoError(t, err)

	type GetSuppressionRequestInfo struct {
		ContainerId string
		Request     email.GetSuppressionRequest
	}

	var requests []GetSuppressionRequestInfo
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
			response, err := c.GetSuppression(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="email-dev_us_grp@oracle.com" jiraProject="ED" opsJiraProject="ED"
func TestEmailClientListSenders(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("email", "ListSenders")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSenders is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("email", "Email", "ListSenders", createEmailClientWithProvider)
	assert.NoError(t, err)
	c := cc.(email.EmailClient)

	body, err := testClient.getRequests("email", "ListSenders")
	assert.NoError(t, err)

	type ListSendersRequestInfo struct {
		ContainerId string
		Request     email.ListSendersRequest
	}

	var requests []ListSendersRequestInfo
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
				r := req.(*email.ListSendersRequest)
				return c.ListSenders(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]email.ListSendersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(email.ListSendersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="email-dev_us_grp@oracle.com" jiraProject="ED" opsJiraProject="ED"
func TestEmailClientListSuppressions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("email", "ListSuppressions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSuppressions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("email", "Email", "ListSuppressions", createEmailClientWithProvider)
	assert.NoError(t, err)
	c := cc.(email.EmailClient)

	body, err := testClient.getRequests("email", "ListSuppressions")
	assert.NoError(t, err)

	type ListSuppressionsRequestInfo struct {
		ContainerId string
		Request     email.ListSuppressionsRequest
	}

	var requests []ListSuppressionsRequestInfo
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
				r := req.(*email.ListSuppressionsRequest)
				return c.ListSuppressions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]email.ListSuppressionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(email.ListSuppressionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="email-dev_us_grp@oracle.com" jiraProject="ED" opsJiraProject="ED"
func TestEmailClientUpdateSender(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("email", "UpdateSender")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSender is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("email", "Email", "UpdateSender", createEmailClientWithProvider)
	assert.NoError(t, err)
	c := cc.(email.EmailClient)

	body, err := testClient.getRequests("email", "UpdateSender")
	assert.NoError(t, err)

	type UpdateSenderRequestInfo struct {
		ContainerId string
		Request     email.UpdateSenderRequest
	}

	var requests []UpdateSenderRequestInfo
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
			response, err := c.UpdateSender(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
