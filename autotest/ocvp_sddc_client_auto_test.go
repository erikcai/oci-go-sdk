package autotest

import (
	"github.com/oracle/oci-go-sdk/v32/common"
	"github.com/oracle/oci-go-sdk/v32/ocvp"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createSddcClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := ocvp.NewSddcClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="sic_ocvp_us_grp@oracle.com" jiraProject="OCVP" opsJiraProject="OCVP"
func TestSddcClientChangeSddcCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocvp", "ChangeSddcCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeSddcCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocvp", "Sddc", "ChangeSddcCompartment", createSddcClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocvp.SddcClient)

	body, err := testClient.getRequests("ocvp", "ChangeSddcCompartment")
	assert.NoError(t, err)

	type ChangeSddcCompartmentRequestInfo struct {
		ContainerId string
		Request     ocvp.ChangeSddcCompartmentRequest
	}

	var requests []ChangeSddcCompartmentRequestInfo
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
			response, err := c.ChangeSddcCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ocvp_us_grp@oracle.com" jiraProject="OCVP" opsJiraProject="OCVP"
func TestSddcClientCreateSddc(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocvp", "CreateSddc")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSddc is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocvp", "Sddc", "CreateSddc", createSddcClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocvp.SddcClient)

	body, err := testClient.getRequests("ocvp", "CreateSddc")
	assert.NoError(t, err)

	type CreateSddcRequestInfo struct {
		ContainerId string
		Request     ocvp.CreateSddcRequest
	}

	var requests []CreateSddcRequestInfo
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
			response, err := c.CreateSddc(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ocvp_us_grp@oracle.com" jiraProject="OCVP" opsJiraProject="OCVP"
func TestSddcClientDeleteSddc(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocvp", "DeleteSddc")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSddc is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocvp", "Sddc", "DeleteSddc", createSddcClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocvp.SddcClient)

	body, err := testClient.getRequests("ocvp", "DeleteSddc")
	assert.NoError(t, err)

	type DeleteSddcRequestInfo struct {
		ContainerId string
		Request     ocvp.DeleteSddcRequest
	}

	var requests []DeleteSddcRequestInfo
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
			response, err := c.DeleteSddc(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ocvp_us_grp@oracle.com" jiraProject="OCVP" opsJiraProject="OCVP"
func TestSddcClientGetSddc(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocvp", "GetSddc")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSddc is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocvp", "Sddc", "GetSddc", createSddcClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocvp.SddcClient)

	body, err := testClient.getRequests("ocvp", "GetSddc")
	assert.NoError(t, err)

	type GetSddcRequestInfo struct {
		ContainerId string
		Request     ocvp.GetSddcRequest
	}

	var requests []GetSddcRequestInfo
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
			response, err := c.GetSddc(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ocvp_us_grp@oracle.com" jiraProject="OCVP" opsJiraProject="OCVP"
func TestSddcClientListSddcs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocvp", "ListSddcs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSddcs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocvp", "Sddc", "ListSddcs", createSddcClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocvp.SddcClient)

	body, err := testClient.getRequests("ocvp", "ListSddcs")
	assert.NoError(t, err)

	type ListSddcsRequestInfo struct {
		ContainerId string
		Request     ocvp.ListSddcsRequest
	}

	var requests []ListSddcsRequestInfo
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
				r := req.(*ocvp.ListSddcsRequest)
				return c.ListSddcs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]ocvp.ListSddcsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(ocvp.ListSddcsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ocvp_us_grp@oracle.com" jiraProject="OCVP" opsJiraProject="OCVP"
func TestSddcClientListSupportedVmwareSoftwareVersions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocvp", "ListSupportedVmwareSoftwareVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSupportedVmwareSoftwareVersions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocvp", "Sddc", "ListSupportedVmwareSoftwareVersions", createSddcClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocvp.SddcClient)

	body, err := testClient.getRequests("ocvp", "ListSupportedVmwareSoftwareVersions")
	assert.NoError(t, err)

	type ListSupportedVmwareSoftwareVersionsRequestInfo struct {
		ContainerId string
		Request     ocvp.ListSupportedVmwareSoftwareVersionsRequest
	}

	var requests []ListSupportedVmwareSoftwareVersionsRequestInfo
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
				r := req.(*ocvp.ListSupportedVmwareSoftwareVersionsRequest)
				return c.ListSupportedVmwareSoftwareVersions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]ocvp.ListSupportedVmwareSoftwareVersionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(ocvp.ListSupportedVmwareSoftwareVersionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ocvp_us_grp@oracle.com" jiraProject="OCVP" opsJiraProject="OCVP"
func TestSddcClientUpdateSddc(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocvp", "UpdateSddc")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSddc is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocvp", "Sddc", "UpdateSddc", createSddcClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocvp.SddcClient)

	body, err := testClient.getRequests("ocvp", "UpdateSddc")
	assert.NoError(t, err)

	type UpdateSddcRequestInfo struct {
		ContainerId string
		Request     ocvp.UpdateSddcRequest
	}

	var requests []UpdateSddcRequestInfo
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
			response, err := c.UpdateSddc(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
