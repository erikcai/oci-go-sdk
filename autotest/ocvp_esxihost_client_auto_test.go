package autotest

import (
	"github.com/oracle/oci-go-sdk/v26/common"
	"github.com/oracle/oci-go-sdk/v26/ocvp"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createEsxiHostClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := ocvp.NewEsxiHostClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="sic_ocvp_us_grp@oracle.com" jiraProject="OCVP" opsJiraProject="OCVP"
func TestEsxiHostClientCreateEsxiHost(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocvp", "CreateEsxiHost")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateEsxiHost is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocvp", "EsxiHost", "CreateEsxiHost", createEsxiHostClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocvp.EsxiHostClient)

	body, err := testClient.getRequests("ocvp", "CreateEsxiHost")
	assert.NoError(t, err)

	type CreateEsxiHostRequestInfo struct {
		ContainerId string
		Request     ocvp.CreateEsxiHostRequest
	}

	var requests []CreateEsxiHostRequestInfo
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
			response, err := c.CreateEsxiHost(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ocvp_us_grp@oracle.com" jiraProject="OCVP" opsJiraProject="OCVP"
func TestEsxiHostClientDeleteEsxiHost(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocvp", "DeleteEsxiHost")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteEsxiHost is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocvp", "EsxiHost", "DeleteEsxiHost", createEsxiHostClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocvp.EsxiHostClient)

	body, err := testClient.getRequests("ocvp", "DeleteEsxiHost")
	assert.NoError(t, err)

	type DeleteEsxiHostRequestInfo struct {
		ContainerId string
		Request     ocvp.DeleteEsxiHostRequest
	}

	var requests []DeleteEsxiHostRequestInfo
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
			response, err := c.DeleteEsxiHost(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ocvp_us_grp@oracle.com" jiraProject="OCVP" opsJiraProject="OCVP"
func TestEsxiHostClientGetEsxiHost(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocvp", "GetEsxiHost")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetEsxiHost is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocvp", "EsxiHost", "GetEsxiHost", createEsxiHostClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocvp.EsxiHostClient)

	body, err := testClient.getRequests("ocvp", "GetEsxiHost")
	assert.NoError(t, err)

	type GetEsxiHostRequestInfo struct {
		ContainerId string
		Request     ocvp.GetEsxiHostRequest
	}

	var requests []GetEsxiHostRequestInfo
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
			response, err := c.GetEsxiHost(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ocvp_us_grp@oracle.com" jiraProject="OCVP" opsJiraProject="OCVP"
func TestEsxiHostClientListEsxiHosts(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocvp", "ListEsxiHosts")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListEsxiHosts is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocvp", "EsxiHost", "ListEsxiHosts", createEsxiHostClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocvp.EsxiHostClient)

	body, err := testClient.getRequests("ocvp", "ListEsxiHosts")
	assert.NoError(t, err)

	type ListEsxiHostsRequestInfo struct {
		ContainerId string
		Request     ocvp.ListEsxiHostsRequest
	}

	var requests []ListEsxiHostsRequestInfo
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
				r := req.(*ocvp.ListEsxiHostsRequest)
				return c.ListEsxiHosts(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]ocvp.ListEsxiHostsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(ocvp.ListEsxiHostsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ocvp_us_grp@oracle.com" jiraProject="OCVP" opsJiraProject="OCVP"
func TestEsxiHostClientUpdateEsxiHost(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocvp", "UpdateEsxiHost")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateEsxiHost is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocvp", "EsxiHost", "UpdateEsxiHost", createEsxiHostClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocvp.EsxiHostClient)

	body, err := testClient.getRequests("ocvp", "UpdateEsxiHost")
	assert.NoError(t, err)

	type UpdateEsxiHostRequestInfo struct {
		ContainerId string
		Request     ocvp.UpdateEsxiHostRequest
	}

	var requests []UpdateEsxiHostRequestInfo
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
			response, err := c.UpdateEsxiHost(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
