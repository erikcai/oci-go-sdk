package autotest

import (
	"github.com/oracle/oci-go-sdk/v29/applicationmigration"
	"github.com/oracle/oci-go-sdk/v29/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createApplicationMigrationClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := applicationmigration.NewApplicationMigrationClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientCancelWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "CancelWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CancelWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "CancelWorkRequest", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "CancelWorkRequest")
	assert.NoError(t, err)

	type CancelWorkRequestRequestInfo struct {
		ContainerId string
		Request     applicationmigration.CancelWorkRequestRequest
	}

	var requests []CancelWorkRequestRequestInfo
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
			response, err := c.CancelWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientChangeMigrationCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "ChangeMigrationCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeMigrationCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "ChangeMigrationCompartment", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "ChangeMigrationCompartment")
	assert.NoError(t, err)

	type ChangeMigrationCompartmentRequestInfo struct {
		ContainerId string
		Request     applicationmigration.ChangeMigrationCompartmentRequest
	}

	var requests []ChangeMigrationCompartmentRequestInfo
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
			response, err := c.ChangeMigrationCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientChangeSourceCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "ChangeSourceCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeSourceCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "ChangeSourceCompartment", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "ChangeSourceCompartment")
	assert.NoError(t, err)

	type ChangeSourceCompartmentRequestInfo struct {
		ContainerId string
		Request     applicationmigration.ChangeSourceCompartmentRequest
	}

	var requests []ChangeSourceCompartmentRequestInfo
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
			response, err := c.ChangeSourceCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientCreateMigration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "CreateMigration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateMigration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "CreateMigration", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "CreateMigration")
	assert.NoError(t, err)

	type CreateMigrationRequestInfo struct {
		ContainerId string
		Request     applicationmigration.CreateMigrationRequest
	}

	var requests []CreateMigrationRequestInfo
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
			response, err := c.CreateMigration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientCreateSource(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "CreateSource")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSource is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "CreateSource", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "CreateSource")
	assert.NoError(t, err)

	type CreateSourceRequestInfo struct {
		ContainerId string
		Request     applicationmigration.CreateSourceRequest
	}

	var requests []CreateSourceRequestInfo
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
			response, err := c.CreateSource(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientDeleteMigration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "DeleteMigration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteMigration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "DeleteMigration", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "DeleteMigration")
	assert.NoError(t, err)

	type DeleteMigrationRequestInfo struct {
		ContainerId string
		Request     applicationmigration.DeleteMigrationRequest
	}

	var requests []DeleteMigrationRequestInfo
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
			response, err := c.DeleteMigration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientDeleteSource(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "DeleteSource")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSource is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "DeleteSource", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "DeleteSource")
	assert.NoError(t, err)

	type DeleteSourceRequestInfo struct {
		ContainerId string
		Request     applicationmigration.DeleteSourceRequest
	}

	var requests []DeleteSourceRequestInfo
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
			response, err := c.DeleteSource(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientGetMigration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "GetMigration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetMigration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "GetMigration", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "GetMigration")
	assert.NoError(t, err)

	type GetMigrationRequestInfo struct {
		ContainerId string
		Request     applicationmigration.GetMigrationRequest
	}

	var requests []GetMigrationRequestInfo
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
			response, err := c.GetMigration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientGetSource(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "GetSource")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSource is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "GetSource", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "GetSource")
	assert.NoError(t, err)

	type GetSourceRequestInfo struct {
		ContainerId string
		Request     applicationmigration.GetSourceRequest
	}

	var requests []GetSourceRequestInfo
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
			response, err := c.GetSource(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "GetWorkRequest", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     applicationmigration.GetWorkRequestRequest
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

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientListMigrations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "ListMigrations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListMigrations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "ListMigrations", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "ListMigrations")
	assert.NoError(t, err)

	type ListMigrationsRequestInfo struct {
		ContainerId string
		Request     applicationmigration.ListMigrationsRequest
	}

	var requests []ListMigrationsRequestInfo
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
				r := req.(*applicationmigration.ListMigrationsRequest)
				return c.ListMigrations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]applicationmigration.ListMigrationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(applicationmigration.ListMigrationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientListSourceApplications(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "ListSourceApplications")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSourceApplications is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "ListSourceApplications", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "ListSourceApplications")
	assert.NoError(t, err)

	type ListSourceApplicationsRequestInfo struct {
		ContainerId string
		Request     applicationmigration.ListSourceApplicationsRequest
	}

	var requests []ListSourceApplicationsRequestInfo
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
				r := req.(*applicationmigration.ListSourceApplicationsRequest)
				return c.ListSourceApplications(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]applicationmigration.ListSourceApplicationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(applicationmigration.ListSourceApplicationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientListSources(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "ListSources")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSources is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "ListSources", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "ListSources")
	assert.NoError(t, err)

	type ListSourcesRequestInfo struct {
		ContainerId string
		Request     applicationmigration.ListSourcesRequest
	}

	var requests []ListSourcesRequestInfo
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
				r := req.(*applicationmigration.ListSourcesRequest)
				return c.ListSources(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]applicationmigration.ListSourcesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(applicationmigration.ListSourcesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "ListWorkRequestErrors", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     applicationmigration.ListWorkRequestErrorsRequest
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
				r := req.(*applicationmigration.ListWorkRequestErrorsRequest)
				return c.ListWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]applicationmigration.ListWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(applicationmigration.ListWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "ListWorkRequestLogs", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     applicationmigration.ListWorkRequestLogsRequest
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
				r := req.(*applicationmigration.ListWorkRequestLogsRequest)
				return c.ListWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]applicationmigration.ListWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(applicationmigration.ListWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "ListWorkRequests", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     applicationmigration.ListWorkRequestsRequest
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
				r := req.(*applicationmigration.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]applicationmigration.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(applicationmigration.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientMigrateApplication(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "MigrateApplication")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("MigrateApplication is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "MigrateApplication", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "MigrateApplication")
	assert.NoError(t, err)

	type MigrateApplicationRequestInfo struct {
		ContainerId string
		Request     applicationmigration.MigrateApplicationRequest
	}

	var requests []MigrateApplicationRequestInfo
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
			response, err := c.MigrateApplication(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientUpdateMigration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "UpdateMigration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateMigration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "UpdateMigration", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "UpdateMigration")
	assert.NoError(t, err)

	type UpdateMigrationRequestInfo struct {
		ContainerId string
		Request     applicationmigration.UpdateMigrationRequest
	}

	var requests []UpdateMigrationRequestInfo
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
			response, err := c.UpdateMigration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci-ams-dev_ww_grp@oracle.com" jiraProject="MIGRATE" opsJiraProject="MIGRATE"
func TestApplicationMigrationClientUpdateSource(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("applicationmigration", "UpdateSource")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSource is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("applicationmigration", "ApplicationMigration", "UpdateSource", createApplicationMigrationClientWithProvider)
	assert.NoError(t, err)
	c := cc.(applicationmigration.ApplicationMigrationClient)

	body, err := testClient.getRequests("applicationmigration", "UpdateSource")
	assert.NoError(t, err)

	type UpdateSourceRequestInfo struct {
		ContainerId string
		Request     applicationmigration.UpdateSourceRequest
	}

	var requests []UpdateSourceRequestInfo
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
			response, err := c.UpdateSource(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
