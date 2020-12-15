package autotest

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"github.com/oracle/oci-go-sdk/v31/managementdashboard"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createManagementdashboardDashxApisClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := managementdashboard.NewDashxApisClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="em_target_analytics_grp@oracle.com" jiraProject="MGMTUI" opsJiraProject="LOGAN"
func TestManagementdashboardDashxApisClientChangeManagementDashboardsCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementdashboard", "ChangeManagementDashboardsCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeManagementDashboardsCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementdashboard", "DashxApis", "ChangeManagementDashboardsCompartment", createManagementdashboardDashxApisClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementdashboard.DashxApisClient)

	body, err := testClient.getRequests("managementdashboard", "ChangeManagementDashboardsCompartment")
	assert.NoError(t, err)

	type ChangeManagementDashboardsCompartmentRequestInfo struct {
		ContainerId string
		Request     managementdashboard.ChangeManagementDashboardsCompartmentRequest
	}

	var requests []ChangeManagementDashboardsCompartmentRequestInfo
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
			response, err := c.ChangeManagementDashboardsCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="em_target_analytics_grp@oracle.com" jiraProject="MGMTUI" opsJiraProject="LOGAN"
func TestManagementdashboardDashxApisClientChangeManagementSavedSearchesCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementdashboard", "ChangeManagementSavedSearchesCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeManagementSavedSearchesCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementdashboard", "DashxApis", "ChangeManagementSavedSearchesCompartment", createManagementdashboardDashxApisClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementdashboard.DashxApisClient)

	body, err := testClient.getRequests("managementdashboard", "ChangeManagementSavedSearchesCompartment")
	assert.NoError(t, err)

	type ChangeManagementSavedSearchesCompartmentRequestInfo struct {
		ContainerId string
		Request     managementdashboard.ChangeManagementSavedSearchesCompartmentRequest
	}

	var requests []ChangeManagementSavedSearchesCompartmentRequestInfo
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
			response, err := c.ChangeManagementSavedSearchesCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="em_target_analytics_grp@oracle.com" jiraProject="MGMTUI" opsJiraProject="LOGAN"
func TestManagementdashboardDashxApisClientCreateManagementDashboard(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementdashboard", "CreateManagementDashboard")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateManagementDashboard is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementdashboard", "DashxApis", "CreateManagementDashboard", createManagementdashboardDashxApisClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementdashboard.DashxApisClient)

	body, err := testClient.getRequests("managementdashboard", "CreateManagementDashboard")
	assert.NoError(t, err)

	type CreateManagementDashboardRequestInfo struct {
		ContainerId string
		Request     managementdashboard.CreateManagementDashboardRequest
	}

	var requests []CreateManagementDashboardRequestInfo
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
			response, err := c.CreateManagementDashboard(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="em_target_analytics_grp@oracle.com" jiraProject="MGMTUI" opsJiraProject="LOGAN"
func TestManagementdashboardDashxApisClientCreateManagementSavedSearch(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementdashboard", "CreateManagementSavedSearch")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateManagementSavedSearch is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementdashboard", "DashxApis", "CreateManagementSavedSearch", createManagementdashboardDashxApisClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementdashboard.DashxApisClient)

	body, err := testClient.getRequests("managementdashboard", "CreateManagementSavedSearch")
	assert.NoError(t, err)

	type CreateManagementSavedSearchRequestInfo struct {
		ContainerId string
		Request     managementdashboard.CreateManagementSavedSearchRequest
	}

	var requests []CreateManagementSavedSearchRequestInfo
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
			response, err := c.CreateManagementSavedSearch(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="em_target_analytics_grp@oracle.com" jiraProject="MGMTUI" opsJiraProject="LOGAN"
func TestManagementdashboardDashxApisClientDeleteManagementDashboard(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementdashboard", "DeleteManagementDashboard")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteManagementDashboard is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementdashboard", "DashxApis", "DeleteManagementDashboard", createManagementdashboardDashxApisClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementdashboard.DashxApisClient)

	body, err := testClient.getRequests("managementdashboard", "DeleteManagementDashboard")
	assert.NoError(t, err)

	type DeleteManagementDashboardRequestInfo struct {
		ContainerId string
		Request     managementdashboard.DeleteManagementDashboardRequest
	}

	var requests []DeleteManagementDashboardRequestInfo
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
			response, err := c.DeleteManagementDashboard(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="em_target_analytics_grp@oracle.com" jiraProject="MGMTUI" opsJiraProject="LOGAN"
func TestManagementdashboardDashxApisClientDeleteManagementSavedSearch(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementdashboard", "DeleteManagementSavedSearch")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteManagementSavedSearch is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementdashboard", "DashxApis", "DeleteManagementSavedSearch", createManagementdashboardDashxApisClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementdashboard.DashxApisClient)

	body, err := testClient.getRequests("managementdashboard", "DeleteManagementSavedSearch")
	assert.NoError(t, err)

	type DeleteManagementSavedSearchRequestInfo struct {
		ContainerId string
		Request     managementdashboard.DeleteManagementSavedSearchRequest
	}

	var requests []DeleteManagementSavedSearchRequestInfo
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
			response, err := c.DeleteManagementSavedSearch(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="em_target_analytics_grp@oracle.com" jiraProject="MGMTUI" opsJiraProject="LOGAN"
func TestManagementdashboardDashxApisClientExportDashboard(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementdashboard", "ExportDashboard")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ExportDashboard is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementdashboard", "DashxApis", "ExportDashboard", createManagementdashboardDashxApisClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementdashboard.DashxApisClient)

	body, err := testClient.getRequests("managementdashboard", "ExportDashboard")
	assert.NoError(t, err)

	type ExportDashboardRequestInfo struct {
		ContainerId string
		Request     managementdashboard.ExportDashboardRequest
	}

	var requests []ExportDashboardRequestInfo
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
			response, err := c.ExportDashboard(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="em_target_analytics_grp@oracle.com" jiraProject="MGMTUI" opsJiraProject="LOGAN"
func TestManagementdashboardDashxApisClientGetManagementDashboard(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementdashboard", "GetManagementDashboard")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetManagementDashboard is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementdashboard", "DashxApis", "GetManagementDashboard", createManagementdashboardDashxApisClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementdashboard.DashxApisClient)

	body, err := testClient.getRequests("managementdashboard", "GetManagementDashboard")
	assert.NoError(t, err)

	type GetManagementDashboardRequestInfo struct {
		ContainerId string
		Request     managementdashboard.GetManagementDashboardRequest
	}

	var requests []GetManagementDashboardRequestInfo
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
			response, err := c.GetManagementDashboard(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="em_target_analytics_grp@oracle.com" jiraProject="MGMTUI" opsJiraProject="LOGAN"
func TestManagementdashboardDashxApisClientGetManagementSavedSearch(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementdashboard", "GetManagementSavedSearch")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetManagementSavedSearch is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementdashboard", "DashxApis", "GetManagementSavedSearch", createManagementdashboardDashxApisClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementdashboard.DashxApisClient)

	body, err := testClient.getRequests("managementdashboard", "GetManagementSavedSearch")
	assert.NoError(t, err)

	type GetManagementSavedSearchRequestInfo struct {
		ContainerId string
		Request     managementdashboard.GetManagementSavedSearchRequest
	}

	var requests []GetManagementSavedSearchRequestInfo
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
			response, err := c.GetManagementSavedSearch(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="em_target_analytics_grp@oracle.com" jiraProject="MGMTUI" opsJiraProject="LOGAN"
func TestManagementdashboardDashxApisClientImportDashboard(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementdashboard", "ImportDashboard")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ImportDashboard is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementdashboard", "DashxApis", "ImportDashboard", createManagementdashboardDashxApisClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementdashboard.DashxApisClient)

	body, err := testClient.getRequests("managementdashboard", "ImportDashboard")
	assert.NoError(t, err)

	type ImportDashboardRequestInfo struct {
		ContainerId string
		Request     managementdashboard.ImportDashboardRequest
	}

	var requests []ImportDashboardRequestInfo
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
			response, err := c.ImportDashboard(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="em_target_analytics_grp@oracle.com" jiraProject="MGMTUI" opsJiraProject="LOGAN"
func TestManagementdashboardDashxApisClientListManagementDashboards(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementdashboard", "ListManagementDashboards")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListManagementDashboards is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementdashboard", "DashxApis", "ListManagementDashboards", createManagementdashboardDashxApisClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementdashboard.DashxApisClient)

	body, err := testClient.getRequests("managementdashboard", "ListManagementDashboards")
	assert.NoError(t, err)

	type ListManagementDashboardsRequestInfo struct {
		ContainerId string
		Request     managementdashboard.ListManagementDashboardsRequest
	}

	var requests []ListManagementDashboardsRequestInfo
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
				r := req.(*managementdashboard.ListManagementDashboardsRequest)
				return c.ListManagementDashboards(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]managementdashboard.ListManagementDashboardsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(managementdashboard.ListManagementDashboardsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="em_target_analytics_grp@oracle.com" jiraProject="MGMTUI" opsJiraProject="LOGAN"
func TestManagementdashboardDashxApisClientListManagementSavedSearches(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementdashboard", "ListManagementSavedSearches")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListManagementSavedSearches is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementdashboard", "DashxApis", "ListManagementSavedSearches", createManagementdashboardDashxApisClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementdashboard.DashxApisClient)

	body, err := testClient.getRequests("managementdashboard", "ListManagementSavedSearches")
	assert.NoError(t, err)

	type ListManagementSavedSearchesRequestInfo struct {
		ContainerId string
		Request     managementdashboard.ListManagementSavedSearchesRequest
	}

	var requests []ListManagementSavedSearchesRequestInfo
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
				r := req.(*managementdashboard.ListManagementSavedSearchesRequest)
				return c.ListManagementSavedSearches(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]managementdashboard.ListManagementSavedSearchesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(managementdashboard.ListManagementSavedSearchesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="em_target_analytics_grp@oracle.com" jiraProject="MGMTUI" opsJiraProject="LOGAN"
func TestManagementdashboardDashxApisClientUpdateManagementDashboard(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementdashboard", "UpdateManagementDashboard")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateManagementDashboard is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementdashboard", "DashxApis", "UpdateManagementDashboard", createManagementdashboardDashxApisClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementdashboard.DashxApisClient)

	body, err := testClient.getRequests("managementdashboard", "UpdateManagementDashboard")
	assert.NoError(t, err)

	type UpdateManagementDashboardRequestInfo struct {
		ContainerId string
		Request     managementdashboard.UpdateManagementDashboardRequest
	}

	var requests []UpdateManagementDashboardRequestInfo
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
			response, err := c.UpdateManagementDashboard(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="em_target_analytics_grp@oracle.com" jiraProject="MGMTUI" opsJiraProject="LOGAN"
func TestManagementdashboardDashxApisClientUpdateManagementSavedSearch(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("managementdashboard", "UpdateManagementSavedSearch")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateManagementSavedSearch is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("managementdashboard", "DashxApis", "UpdateManagementSavedSearch", createManagementdashboardDashxApisClientWithProvider)
	assert.NoError(t, err)
	c := cc.(managementdashboard.DashxApisClient)

	body, err := testClient.getRequests("managementdashboard", "UpdateManagementSavedSearch")
	assert.NoError(t, err)

	type UpdateManagementSavedSearchRequestInfo struct {
		ContainerId string
		Request     managementdashboard.UpdateManagementSavedSearchRequest
	}

	var requests []UpdateManagementSavedSearchRequestInfo
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
			response, err := c.UpdateManagementSavedSearch(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
