package autotest

import (
	"github.com/oracle/oci-go-sdk/cims"
	"github.com/oracle/oci-go-sdk/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createIncidentClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := cims.NewIncidentClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_ops_cims_dev_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="CIMS"
func TestIncidentClientCreateIncident(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cims", "CreateIncident")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateIncident is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cims", "Incident", "CreateIncident", createIncidentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cims.IncidentClient)

	body, err := testClient.getRequests("cims", "CreateIncident")
	assert.NoError(t, err)

	type CreateIncidentRequestInfo struct {
		ContainerId string
		Request     cims.CreateIncidentRequest
	}

	var requests []CreateIncidentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateIncident(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_ops_cims_dev_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="CIMS"
func TestIncidentClientGetIncident(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cims", "GetIncident")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetIncident is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cims", "Incident", "GetIncident", createIncidentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cims.IncidentClient)

	body, err := testClient.getRequests("cims", "GetIncident")
	assert.NoError(t, err)

	type GetIncidentRequestInfo struct {
		ContainerId string
		Request     cims.GetIncidentRequest
	}

	var requests []GetIncidentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetIncident(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_ops_cims_dev_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="CIMS"
func TestIncidentClientGetStatus(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cims", "GetStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStatus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cims", "Incident", "GetStatus", createIncidentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cims.IncidentClient)

	body, err := testClient.getRequests("cims", "GetStatus")
	assert.NoError(t, err)

	type GetStatusRequestInfo struct {
		ContainerId string
		Request     cims.GetStatusRequest
	}

	var requests []GetStatusRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_ops_cims_dev_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="CIMS"
func TestIncidentClientListIncidents(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cims", "ListIncidents")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListIncidents is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cims", "Incident", "ListIncidents", createIncidentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cims.IncidentClient)

	body, err := testClient.getRequests("cims", "ListIncidents")
	assert.NoError(t, err)

	type ListIncidentsRequestInfo struct {
		ContainerId string
		Request     cims.ListIncidentsRequest
	}

	var requests []ListIncidentsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*cims.ListIncidentsRequest)
				return c.ListIncidents(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cims.ListIncidentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cims.ListIncidentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_ops_cims_dev_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="CIMS"
func TestIncidentClientListTypes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cims", "ListTypes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTypes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cims", "Incident", "ListTypes", createIncidentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cims.IncidentClient)

	body, err := testClient.getRequests("cims", "ListTypes")
	assert.NoError(t, err)

	type ListTypesRequestInfo struct {
		ContainerId string
		Request     cims.ListTypesRequest
	}

	var requests []ListTypesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*cims.ListTypesRequest)
				return c.ListTypes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]cims.ListTypesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(cims.ListTypesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_ops_cims_dev_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="CIMS"
func TestIncidentClientUpdateIncident(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cims", "UpdateIncident")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateIncident is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cims", "Incident", "UpdateIncident", createIncidentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cims.IncidentClient)

	body, err := testClient.getRequests("cims", "UpdateIncident")
	assert.NoError(t, err)

	type UpdateIncidentRequestInfo struct {
		ContainerId string
		Request     cims.UpdateIncidentRequest
	}

	var requests []UpdateIncidentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateIncident(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_ops_cims_dev_us_grp@oracle.com" jiraProject="JIRA" opsJiraProject="CIMS"
func TestIncidentClientValidateUser(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cims", "ValidateUser")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ValidateUser is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cims", "Incident", "ValidateUser", createIncidentClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cims.IncidentClient)

	body, err := testClient.getRequests("cims", "ValidateUser")
	assert.NoError(t, err)

	type ValidateUserRequestInfo struct {
		ContainerId string
		Request     cims.ValidateUserRequest
	}

	var requests []ValidateUserRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ValidateUser(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
