package autotest

import (
	"github.com/oracle/oci-go-sdk/v32/common"
	"github.com/oracle/oci-go-sdk/v32/databasemanagement"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createDatabasemanagementDbManagementClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := databasemanagement.NewDbManagementClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientAddManagedDatabaseToManagedDatabaseGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "AddManagedDatabaseToManagedDatabaseGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AddManagedDatabaseToManagedDatabaseGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "AddManagedDatabaseToManagedDatabaseGroup", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "AddManagedDatabaseToManagedDatabaseGroup")
	assert.NoError(t, err)

	type AddManagedDatabaseToManagedDatabaseGroupRequestInfo struct {
		ContainerId string
		Request     databasemanagement.AddManagedDatabaseToManagedDatabaseGroupRequest
	}

	var requests []AddManagedDatabaseToManagedDatabaseGroupRequestInfo
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
			response, err := c.AddManagedDatabaseToManagedDatabaseGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientChangeJobCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "ChangeJobCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeJobCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "ChangeJobCompartment", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "ChangeJobCompartment")
	assert.NoError(t, err)

	type ChangeJobCompartmentRequestInfo struct {
		ContainerId string
		Request     databasemanagement.ChangeJobCompartmentRequest
	}

	var requests []ChangeJobCompartmentRequestInfo
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
			response, err := c.ChangeJobCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientChangeManagedDatabaseGroupCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "ChangeManagedDatabaseGroupCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeManagedDatabaseGroupCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "ChangeManagedDatabaseGroupCompartment", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "ChangeManagedDatabaseGroupCompartment")
	assert.NoError(t, err)

	type ChangeManagedDatabaseGroupCompartmentRequestInfo struct {
		ContainerId string
		Request     databasemanagement.ChangeManagedDatabaseGroupCompartmentRequest
	}

	var requests []ChangeManagedDatabaseGroupCompartmentRequestInfo
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
			response, err := c.ChangeManagedDatabaseGroupCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientCreateJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "CreateJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "CreateJob", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "CreateJob")
	assert.NoError(t, err)

	type CreateJobRequestInfo struct {
		ContainerId string
		Request     databasemanagement.CreateJobRequest
	}

	var requests []CreateJobRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateJobRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateJobDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "jobType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"SQL": &databasemanagement.CreateSqlJobDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientCreateManagedDatabaseGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "CreateManagedDatabaseGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateManagedDatabaseGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "CreateManagedDatabaseGroup", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "CreateManagedDatabaseGroup")
	assert.NoError(t, err)

	type CreateManagedDatabaseGroupRequestInfo struct {
		ContainerId string
		Request     databasemanagement.CreateManagedDatabaseGroupRequest
	}

	var requests []CreateManagedDatabaseGroupRequestInfo
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
			response, err := c.CreateManagedDatabaseGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientDeleteJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "DeleteJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "DeleteJob", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "DeleteJob")
	assert.NoError(t, err)

	type DeleteJobRequestInfo struct {
		ContainerId string
		Request     databasemanagement.DeleteJobRequest
	}

	var requests []DeleteJobRequestInfo
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
			response, err := c.DeleteJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientDeleteManagedDatabaseGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "DeleteManagedDatabaseGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteManagedDatabaseGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "DeleteManagedDatabaseGroup", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "DeleteManagedDatabaseGroup")
	assert.NoError(t, err)

	type DeleteManagedDatabaseGroupRequestInfo struct {
		ContainerId string
		Request     databasemanagement.DeleteManagedDatabaseGroupRequest
	}

	var requests []DeleteManagedDatabaseGroupRequestInfo
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
			response, err := c.DeleteManagedDatabaseGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientGetDatabaseFleetHealthMetrics(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "GetDatabaseFleetHealthMetrics")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDatabaseFleetHealthMetrics is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "GetDatabaseFleetHealthMetrics", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "GetDatabaseFleetHealthMetrics")
	assert.NoError(t, err)

	type GetDatabaseFleetHealthMetricsRequestInfo struct {
		ContainerId string
		Request     databasemanagement.GetDatabaseFleetHealthMetricsRequest
	}

	var requests []GetDatabaseFleetHealthMetricsRequestInfo
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
			response, err := c.GetDatabaseFleetHealthMetrics(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientGetDatabaseHomeMetrics(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "GetDatabaseHomeMetrics")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDatabaseHomeMetrics is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "GetDatabaseHomeMetrics", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "GetDatabaseHomeMetrics")
	assert.NoError(t, err)

	type GetDatabaseHomeMetricsRequestInfo struct {
		ContainerId string
		Request     databasemanagement.GetDatabaseHomeMetricsRequest
	}

	var requests []GetDatabaseHomeMetricsRequestInfo
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
			response, err := c.GetDatabaseHomeMetrics(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientGetJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "GetJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "GetJob", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "GetJob")
	assert.NoError(t, err)

	type GetJobRequestInfo struct {
		ContainerId string
		Request     databasemanagement.GetJobRequest
	}

	var requests []GetJobRequestInfo
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
			response, err := c.GetJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientGetJobExecution(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "GetJobExecution")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJobExecution is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "GetJobExecution", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "GetJobExecution")
	assert.NoError(t, err)

	type GetJobExecutionRequestInfo struct {
		ContainerId string
		Request     databasemanagement.GetJobExecutionRequest
	}

	var requests []GetJobExecutionRequestInfo
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
			response, err := c.GetJobExecution(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientGetJobRun(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "GetJobRun")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJobRun is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "GetJobRun", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "GetJobRun")
	assert.NoError(t, err)

	type GetJobRunRequestInfo struct {
		ContainerId string
		Request     databasemanagement.GetJobRunRequest
	}

	var requests []GetJobRunRequestInfo
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
			response, err := c.GetJobRun(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientGetManagedDatabase(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "GetManagedDatabase")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetManagedDatabase is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "GetManagedDatabase", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "GetManagedDatabase")
	assert.NoError(t, err)

	type GetManagedDatabaseRequestInfo struct {
		ContainerId string
		Request     databasemanagement.GetManagedDatabaseRequest
	}

	var requests []GetManagedDatabaseRequestInfo
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
			response, err := c.GetManagedDatabase(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientGetManagedDatabaseGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "GetManagedDatabaseGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetManagedDatabaseGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "GetManagedDatabaseGroup", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "GetManagedDatabaseGroup")
	assert.NoError(t, err)

	type GetManagedDatabaseGroupRequestInfo struct {
		ContainerId string
		Request     databasemanagement.GetManagedDatabaseGroupRequest
	}

	var requests []GetManagedDatabaseGroupRequestInfo
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
			response, err := c.GetManagedDatabaseGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientListJobExecutions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "ListJobExecutions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListJobExecutions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "ListJobExecutions", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "ListJobExecutions")
	assert.NoError(t, err)

	type ListJobExecutionsRequestInfo struct {
		ContainerId string
		Request     databasemanagement.ListJobExecutionsRequest
	}

	var requests []ListJobExecutionsRequestInfo
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
				r := req.(*databasemanagement.ListJobExecutionsRequest)
				return c.ListJobExecutions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]databasemanagement.ListJobExecutionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(databasemanagement.ListJobExecutionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientListJobRuns(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "ListJobRuns")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListJobRuns is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "ListJobRuns", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "ListJobRuns")
	assert.NoError(t, err)

	type ListJobRunsRequestInfo struct {
		ContainerId string
		Request     databasemanagement.ListJobRunsRequest
	}

	var requests []ListJobRunsRequestInfo
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
				r := req.(*databasemanagement.ListJobRunsRequest)
				return c.ListJobRuns(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]databasemanagement.ListJobRunsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(databasemanagement.ListJobRunsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientListJobs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "ListJobs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListJobs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "ListJobs", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "ListJobs")
	assert.NoError(t, err)

	type ListJobsRequestInfo struct {
		ContainerId string
		Request     databasemanagement.ListJobsRequest
	}

	var requests []ListJobsRequestInfo
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
				r := req.(*databasemanagement.ListJobsRequest)
				return c.ListJobs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]databasemanagement.ListJobsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(databasemanagement.ListJobsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientListManagedDatabaseGroups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "ListManagedDatabaseGroups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListManagedDatabaseGroups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "ListManagedDatabaseGroups", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "ListManagedDatabaseGroups")
	assert.NoError(t, err)

	type ListManagedDatabaseGroupsRequestInfo struct {
		ContainerId string
		Request     databasemanagement.ListManagedDatabaseGroupsRequest
	}

	var requests []ListManagedDatabaseGroupsRequestInfo
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
				r := req.(*databasemanagement.ListManagedDatabaseGroupsRequest)
				return c.ListManagedDatabaseGroups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]databasemanagement.ListManagedDatabaseGroupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(databasemanagement.ListManagedDatabaseGroupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientListManagedDatabases(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "ListManagedDatabases")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListManagedDatabases is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "ListManagedDatabases", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "ListManagedDatabases")
	assert.NoError(t, err)

	type ListManagedDatabasesRequestInfo struct {
		ContainerId string
		Request     databasemanagement.ListManagedDatabasesRequest
	}

	var requests []ListManagedDatabasesRequestInfo
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
				r := req.(*databasemanagement.ListManagedDatabasesRequest)
				return c.ListManagedDatabases(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]databasemanagement.ListManagedDatabasesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(databasemanagement.ListManagedDatabasesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientRemoveManagedDatabaseFromManagedDatabaseGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "RemoveManagedDatabaseFromManagedDatabaseGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RemoveManagedDatabaseFromManagedDatabaseGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "RemoveManagedDatabaseFromManagedDatabaseGroup", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "RemoveManagedDatabaseFromManagedDatabaseGroup")
	assert.NoError(t, err)

	type RemoveManagedDatabaseFromManagedDatabaseGroupRequestInfo struct {
		ContainerId string
		Request     databasemanagement.RemoveManagedDatabaseFromManagedDatabaseGroupRequest
	}

	var requests []RemoveManagedDatabaseFromManagedDatabaseGroupRequestInfo
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
			response, err := c.RemoveManagedDatabaseFromManagedDatabaseGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="dpd_dev_grp@oracle.com" jiraProject="DPD" opsJiraProject="DPD"
func TestDatabasemanagementDbManagementClientUpdateManagedDatabaseGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("databasemanagement", "UpdateManagedDatabaseGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateManagedDatabaseGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("databasemanagement", "DbManagement", "UpdateManagedDatabaseGroup", createDatabasemanagementDbManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(databasemanagement.DbManagementClient)

	body, err := testClient.getRequests("databasemanagement", "UpdateManagedDatabaseGroup")
	assert.NoError(t, err)

	type UpdateManagedDatabaseGroupRequestInfo struct {
		ContainerId string
		Request     databasemanagement.UpdateManagedDatabaseGroupRequest
	}

	var requests []UpdateManagedDatabaseGroupRequestInfo
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
			response, err := c.UpdateManagedDatabaseGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
