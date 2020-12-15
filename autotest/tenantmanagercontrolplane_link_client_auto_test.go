package autotest

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"github.com/oracle/oci-go-sdk/v31/tenantmanagercontrolplane"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createTenantmanagercontrolplaneLinkClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := tenantmanagercontrolplane.NewLinkClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="acc_customer_tools_us_grp@oracle.com" jiraProject="ACCMGMT" opsJiraProject="ACCMGMT"
func TestTenantmanagercontrolplaneLinkClientDeleteLink(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("tenantmanagercontrolplane", "DeleteLink")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteLink is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("tenantmanagercontrolplane", "Link", "DeleteLink", createTenantmanagercontrolplaneLinkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(tenantmanagercontrolplane.LinkClient)

	body, err := testClient.getRequests("tenantmanagercontrolplane", "DeleteLink")
	assert.NoError(t, err)

	type DeleteLinkRequestInfo struct {
		ContainerId string
		Request     tenantmanagercontrolplane.DeleteLinkRequest
	}

	var requests []DeleteLinkRequestInfo
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
			response, err := c.DeleteLink(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="acc_customer_tools_us_grp@oracle.com" jiraProject="ACCMGMT" opsJiraProject="ACCMGMT"
func TestTenantmanagercontrolplaneLinkClientGetLink(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("tenantmanagercontrolplane", "GetLink")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLink is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("tenantmanagercontrolplane", "Link", "GetLink", createTenantmanagercontrolplaneLinkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(tenantmanagercontrolplane.LinkClient)

	body, err := testClient.getRequests("tenantmanagercontrolplane", "GetLink")
	assert.NoError(t, err)

	type GetLinkRequestInfo struct {
		ContainerId string
		Request     tenantmanagercontrolplane.GetLinkRequest
	}

	var requests []GetLinkRequestInfo
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
			response, err := c.GetLink(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="acc_customer_tools_us_grp@oracle.com" jiraProject="ACCMGMT" opsJiraProject="ACCMGMT"
func TestTenantmanagercontrolplaneLinkClientListLinks(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("tenantmanagercontrolplane", "ListLinks")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLinks is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("tenantmanagercontrolplane", "Link", "ListLinks", createTenantmanagercontrolplaneLinkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(tenantmanagercontrolplane.LinkClient)

	body, err := testClient.getRequests("tenantmanagercontrolplane", "ListLinks")
	assert.NoError(t, err)

	type ListLinksRequestInfo struct {
		ContainerId string
		Request     tenantmanagercontrolplane.ListLinksRequest
	}

	var requests []ListLinksRequestInfo
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
				r := req.(*tenantmanagercontrolplane.ListLinksRequest)
				return c.ListLinks(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]tenantmanagercontrolplane.ListLinksResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(tenantmanagercontrolplane.ListLinksResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
