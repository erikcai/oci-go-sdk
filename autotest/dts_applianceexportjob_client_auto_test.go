package autotest

import (
	"github.com/oracle/oci-go-sdk/v25/common"
	"github.com/oracle/oci-go-sdk/v25/dts"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createApplianceExportJobClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := dts.NewApplianceExportJobClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="data_transfer_platform_dev_ww_grp@oracle.com" jiraProject="BDTS" opsJiraProject="DTS"
func TestApplianceExportJobClientChangeApplianceExportJobCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dts", "ChangeApplianceExportJobCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeApplianceExportJobCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dts", "ApplianceExportJob", "ChangeApplianceExportJobCompartment", createApplianceExportJobClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dts.ApplianceExportJobClient)

	body, err := testClient.getRequests("dts", "ChangeApplianceExportJobCompartment")
	assert.NoError(t, err)

	type ChangeApplianceExportJobCompartmentRequestInfo struct {
		ContainerId string
		Request     dts.ChangeApplianceExportJobCompartmentRequest
	}

	var requests []ChangeApplianceExportJobCompartmentRequestInfo
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
			response, err := c.ChangeApplianceExportJobCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="data_transfer_platform_dev_ww_grp@oracle.com" jiraProject="BDTS" opsJiraProject="DTS"
func TestApplianceExportJobClientCreateApplianceExportJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dts", "CreateApplianceExportJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateApplianceExportJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dts", "ApplianceExportJob", "CreateApplianceExportJob", createApplianceExportJobClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dts.ApplianceExportJobClient)

	body, err := testClient.getRequests("dts", "CreateApplianceExportJob")
	assert.NoError(t, err)

	type CreateApplianceExportJobRequestInfo struct {
		ContainerId string
		Request     dts.CreateApplianceExportJobRequest
	}

	var requests []CreateApplianceExportJobRequestInfo
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
			response, err := c.CreateApplianceExportJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="data_transfer_platform_dev_ww_grp@oracle.com" jiraProject="BDTS" opsJiraProject="DTS"
func TestApplianceExportJobClientDeleteApplianceExportJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dts", "DeleteApplianceExportJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteApplianceExportJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dts", "ApplianceExportJob", "DeleteApplianceExportJob", createApplianceExportJobClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dts.ApplianceExportJobClient)

	body, err := testClient.getRequests("dts", "DeleteApplianceExportJob")
	assert.NoError(t, err)

	type DeleteApplianceExportJobRequestInfo struct {
		ContainerId string
		Request     dts.DeleteApplianceExportJobRequest
	}

	var requests []DeleteApplianceExportJobRequestInfo
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
			response, err := c.DeleteApplianceExportJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="data_transfer_platform_dev_ww_grp@oracle.com" jiraProject="BDTS" opsJiraProject="DTS"
func TestApplianceExportJobClientGetApplianceExportJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dts", "GetApplianceExportJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetApplianceExportJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dts", "ApplianceExportJob", "GetApplianceExportJob", createApplianceExportJobClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dts.ApplianceExportJobClient)

	body, err := testClient.getRequests("dts", "GetApplianceExportJob")
	assert.NoError(t, err)

	type GetApplianceExportJobRequestInfo struct {
		ContainerId string
		Request     dts.GetApplianceExportJobRequest
	}

	var requests []GetApplianceExportJobRequestInfo
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
			response, err := c.GetApplianceExportJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="data_transfer_platform_dev_ww_grp@oracle.com" jiraProject="BDTS" opsJiraProject="DTS"
func TestApplianceExportJobClientListApplianceExportJobs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dts", "ListApplianceExportJobs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListApplianceExportJobs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dts", "ApplianceExportJob", "ListApplianceExportJobs", createApplianceExportJobClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dts.ApplianceExportJobClient)

	body, err := testClient.getRequests("dts", "ListApplianceExportJobs")
	assert.NoError(t, err)

	type ListApplianceExportJobsRequestInfo struct {
		ContainerId string
		Request     dts.ListApplianceExportJobsRequest
	}

	var requests []ListApplianceExportJobsRequestInfo
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
				r := req.(*dts.ListApplianceExportJobsRequest)
				return c.ListApplianceExportJobs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dts.ListApplianceExportJobsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dts.ListApplianceExportJobsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="data_transfer_platform_dev_ww_grp@oracle.com" jiraProject="BDTS" opsJiraProject="DTS"
func TestApplianceExportJobClientUpdateApplianceExportJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dts", "UpdateApplianceExportJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateApplianceExportJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dts", "ApplianceExportJob", "UpdateApplianceExportJob", createApplianceExportJobClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dts.ApplianceExportJobClient)

	body, err := testClient.getRequests("dts", "UpdateApplianceExportJob")
	assert.NoError(t, err)

	type UpdateApplianceExportJobRequestInfo struct {
		ContainerId string
		Request     dts.UpdateApplianceExportJobRequest
	}

	var requests []UpdateApplianceExportJobRequestInfo
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
			response, err := c.UpdateApplianceExportJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
