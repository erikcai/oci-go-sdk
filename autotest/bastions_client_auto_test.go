package autotest

import (
	"github.com/oracle/oci-go-sdk/bastions"
	"github.com/oracle/oci-go-sdk/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createBastionsBastionsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := bastions.NewBastionsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionsBastionsClientChangeBastionCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastions", "ChangeBastionCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeBastionCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastions", "Bastions", "ChangeBastionCompartment", createBastionsBastionsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastions.BastionsClient)

	body, err := testClient.getRequests("bastions", "ChangeBastionCompartment")
	assert.NoError(t, err)

	type ChangeBastionCompartmentRequestInfo struct {
		ContainerId string
		Request     bastions.ChangeBastionCompartmentRequest
	}

	var requests []ChangeBastionCompartmentRequestInfo
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
			response, err := c.ChangeBastionCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionsBastionsClientCreateBastion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastions", "CreateBastion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBastion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastions", "Bastions", "CreateBastion", createBastionsBastionsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastions.BastionsClient)

	body, err := testClient.getRequests("bastions", "CreateBastion")
	assert.NoError(t, err)

	type CreateBastionRequestInfo struct {
		ContainerId string
		Request     bastions.CreateBastionRequest
	}

	var requests []CreateBastionRequestInfo
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
			response, err := c.CreateBastion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionsBastionsClientDeleteBastion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastions", "DeleteBastion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBastion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastions", "Bastions", "DeleteBastion", createBastionsBastionsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastions.BastionsClient)

	body, err := testClient.getRequests("bastions", "DeleteBastion")
	assert.NoError(t, err)

	type DeleteBastionRequestInfo struct {
		ContainerId string
		Request     bastions.DeleteBastionRequest
	}

	var requests []DeleteBastionRequestInfo
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
			response, err := c.DeleteBastion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionsBastionsClientGetBastion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastions", "GetBastion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBastion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastions", "Bastions", "GetBastion", createBastionsBastionsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastions.BastionsClient)

	body, err := testClient.getRequests("bastions", "GetBastion")
	assert.NoError(t, err)

	type GetBastionRequestInfo struct {
		ContainerId string
		Request     bastions.GetBastionRequest
	}

	var requests []GetBastionRequestInfo
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
			response, err := c.GetBastion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionsBastionsClientListBastions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastions", "ListBastions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBastions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastions", "Bastions", "ListBastions", createBastionsBastionsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastions.BastionsClient)

	body, err := testClient.getRequests("bastions", "ListBastions")
	assert.NoError(t, err)

	type ListBastionsRequestInfo struct {
		ContainerId string
		Request     bastions.ListBastionsRequest
	}

	var requests []ListBastionsRequestInfo
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
				r := req.(*bastions.ListBastionsRequest)
				return c.ListBastions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]bastions.ListBastionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(bastions.ListBastionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_bastions_us_grp@oracle.com" jiraProject="BSTN" opsJiraProject="Bastions"
func TestBastionsBastionsClientUpdateBastion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("bastions", "UpdateBastion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBastion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("bastions", "Bastions", "UpdateBastion", createBastionsBastionsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(bastions.BastionsClient)

	body, err := testClient.getRequests("bastions", "UpdateBastion")
	assert.NoError(t, err)

	type UpdateBastionRequestInfo struct {
		ContainerId string
		Request     bastions.UpdateBastionRequest
	}

	var requests []UpdateBastionRequestInfo
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
			response, err := c.UpdateBastion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
