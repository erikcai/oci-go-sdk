package autotest

import (
	"github.com/oracle/oci-go-sdk/v28/common"
	"github.com/oracle/oci-go-sdk/v28/loganalytics"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createLoganalyticsLogAnalyticsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := loganalytics.NewLogAnalyticsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientAddEntityAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "AddEntityAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AddEntityAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "AddEntityAssociation", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "AddEntityAssociation")
	assert.NoError(t, err)

	type AddEntityAssociationRequestInfo struct {
		ContainerId string
		Request     loganalytics.AddEntityAssociationRequest
	}

	var requests []AddEntityAssociationRequestInfo
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
			response, err := c.AddEntityAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientBatchGetBasicInfo(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "BatchGetBasicInfo")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("BatchGetBasicInfo is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "BatchGetBasicInfo", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "BatchGetBasicInfo")
	assert.NoError(t, err)

	type BatchGetBasicInfoRequestInfo struct {
		ContainerId string
		Request     loganalytics.BatchGetBasicInfoRequest
	}

	var requests []BatchGetBasicInfoRequestInfo
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
				r := req.(*loganalytics.BatchGetBasicInfoRequest)
				return c.BatchGetBasicInfo(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.BatchGetBasicInfoResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.BatchGetBasicInfoResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientCancelQueryWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "CancelQueryWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CancelQueryWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "CancelQueryWorkRequest", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "CancelQueryWorkRequest")
	assert.NoError(t, err)

	type CancelQueryWorkRequestRequestInfo struct {
		ContainerId string
		Request     loganalytics.CancelQueryWorkRequestRequest
	}

	var requests []CancelQueryWorkRequestRequestInfo
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
			response, err := c.CancelQueryWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientChangeLogAnalyticsEntityCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ChangeLogAnalyticsEntityCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeLogAnalyticsEntityCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ChangeLogAnalyticsEntityCompartment", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ChangeLogAnalyticsEntityCompartment")
	assert.NoError(t, err)

	type ChangeLogAnalyticsEntityCompartmentRequestInfo struct {
		ContainerId string
		Request     loganalytics.ChangeLogAnalyticsEntityCompartmentRequest
	}

	var requests []ChangeLogAnalyticsEntityCompartmentRequestInfo
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
			response, err := c.ChangeLogAnalyticsEntityCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientChangeLogAnalyticsLogGroupCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ChangeLogAnalyticsLogGroupCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeLogAnalyticsLogGroupCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ChangeLogAnalyticsLogGroupCompartment", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ChangeLogAnalyticsLogGroupCompartment")
	assert.NoError(t, err)

	type ChangeLogAnalyticsLogGroupCompartmentRequestInfo struct {
		ContainerId string
		Request     loganalytics.ChangeLogAnalyticsLogGroupCompartmentRequest
	}

	var requests []ChangeLogAnalyticsLogGroupCompartmentRequestInfo
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
			response, err := c.ChangeLogAnalyticsLogGroupCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientChangeLogAnalyticsObjectCollectionRuleCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ChangeLogAnalyticsObjectCollectionRuleCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeLogAnalyticsObjectCollectionRuleCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ChangeLogAnalyticsObjectCollectionRuleCompartment", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ChangeLogAnalyticsObjectCollectionRuleCompartment")
	assert.NoError(t, err)

	type ChangeLogAnalyticsObjectCollectionRuleCompartmentRequestInfo struct {
		ContainerId string
		Request     loganalytics.ChangeLogAnalyticsObjectCollectionRuleCompartmentRequest
	}

	var requests []ChangeLogAnalyticsObjectCollectionRuleCompartmentRequestInfo
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
			response, err := c.ChangeLogAnalyticsObjectCollectionRuleCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientChangeScheduledTaskCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ChangeScheduledTaskCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeScheduledTaskCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ChangeScheduledTaskCompartment", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ChangeScheduledTaskCompartment")
	assert.NoError(t, err)

	type ChangeScheduledTaskCompartmentRequestInfo struct {
		ContainerId string
		Request     loganalytics.ChangeScheduledTaskCompartmentRequest
	}

	var requests []ChangeScheduledTaskCompartmentRequestInfo
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
			response, err := c.ChangeScheduledTaskCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientClean(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "Clean")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("Clean is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "Clean", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "Clean")
	assert.NoError(t, err)

	type CleanRequestInfo struct {
		ContainerId string
		Request     loganalytics.CleanRequest
	}

	var requests []CleanRequestInfo
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
			response, err := c.Clean(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientCreateLogAnalyticsEntity(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "CreateLogAnalyticsEntity")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateLogAnalyticsEntity is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "CreateLogAnalyticsEntity", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "CreateLogAnalyticsEntity")
	assert.NoError(t, err)

	type CreateLogAnalyticsEntityRequestInfo struct {
		ContainerId string
		Request     loganalytics.CreateLogAnalyticsEntityRequest
	}

	var requests []CreateLogAnalyticsEntityRequestInfo
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
			response, err := c.CreateLogAnalyticsEntity(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientCreateLogAnalyticsEntityType(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "CreateLogAnalyticsEntityType")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateLogAnalyticsEntityType is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "CreateLogAnalyticsEntityType", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "CreateLogAnalyticsEntityType")
	assert.NoError(t, err)

	type CreateLogAnalyticsEntityTypeRequestInfo struct {
		ContainerId string
		Request     loganalytics.CreateLogAnalyticsEntityTypeRequest
	}

	var requests []CreateLogAnalyticsEntityTypeRequestInfo
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
			response, err := c.CreateLogAnalyticsEntityType(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientCreateLogAnalyticsLogGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "CreateLogAnalyticsLogGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateLogAnalyticsLogGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "CreateLogAnalyticsLogGroup", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "CreateLogAnalyticsLogGroup")
	assert.NoError(t, err)

	type CreateLogAnalyticsLogGroupRequestInfo struct {
		ContainerId string
		Request     loganalytics.CreateLogAnalyticsLogGroupRequest
	}

	var requests []CreateLogAnalyticsLogGroupRequestInfo
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
			response, err := c.CreateLogAnalyticsLogGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientCreateLogAnalyticsObjectCollectionRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "CreateLogAnalyticsObjectCollectionRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateLogAnalyticsObjectCollectionRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "CreateLogAnalyticsObjectCollectionRule", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "CreateLogAnalyticsObjectCollectionRule")
	assert.NoError(t, err)

	type CreateLogAnalyticsObjectCollectionRuleRequestInfo struct {
		ContainerId string
		Request     loganalytics.CreateLogAnalyticsObjectCollectionRuleRequest
	}

	var requests []CreateLogAnalyticsObjectCollectionRuleRequestInfo
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
			response, err := c.CreateLogAnalyticsObjectCollectionRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientCreateScheduledTask(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "CreateScheduledTask")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateScheduledTask is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "CreateScheduledTask", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "CreateScheduledTask")
	assert.NoError(t, err)

	type CreateScheduledTaskRequestInfo struct {
		ContainerId string
		Request     loganalytics.CreateScheduledTaskRequest
	}

	var requests []CreateScheduledTaskRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateScheduledTaskRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateScheduledTaskDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "kind",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"STANDARD":     &loganalytics.CreateStandardTaskDetails{},
				"ACCELERATION": &loganalytics.CreateAccelerationTaskDetails{},
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
			response, err := c.CreateScheduledTask(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientDeleteAssociations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "DeleteAssociations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAssociations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "DeleteAssociations", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "DeleteAssociations")
	assert.NoError(t, err)

	type DeleteAssociationsRequestInfo struct {
		ContainerId string
		Request     loganalytics.DeleteAssociationsRequest
	}

	var requests []DeleteAssociationsRequestInfo
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
			response, err := c.DeleteAssociations(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientDeleteField(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "DeleteField")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteField is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "DeleteField", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "DeleteField")
	assert.NoError(t, err)

	type DeleteFieldRequestInfo struct {
		ContainerId string
		Request     loganalytics.DeleteFieldRequest
	}

	var requests []DeleteFieldRequestInfo
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
			response, err := c.DeleteField(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientDeleteLabel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "DeleteLabel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteLabel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "DeleteLabel", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "DeleteLabel")
	assert.NoError(t, err)

	type DeleteLabelRequestInfo struct {
		ContainerId string
		Request     loganalytics.DeleteLabelRequest
	}

	var requests []DeleteLabelRequestInfo
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
			response, err := c.DeleteLabel(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientDeleteLogAnalyticsEntity(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "DeleteLogAnalyticsEntity")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteLogAnalyticsEntity is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "DeleteLogAnalyticsEntity", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "DeleteLogAnalyticsEntity")
	assert.NoError(t, err)

	type DeleteLogAnalyticsEntityRequestInfo struct {
		ContainerId string
		Request     loganalytics.DeleteLogAnalyticsEntityRequest
	}

	var requests []DeleteLogAnalyticsEntityRequestInfo
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
			response, err := c.DeleteLogAnalyticsEntity(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientDeleteLogAnalyticsEntityType(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "DeleteLogAnalyticsEntityType")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteLogAnalyticsEntityType is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "DeleteLogAnalyticsEntityType", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "DeleteLogAnalyticsEntityType")
	assert.NoError(t, err)

	type DeleteLogAnalyticsEntityTypeRequestInfo struct {
		ContainerId string
		Request     loganalytics.DeleteLogAnalyticsEntityTypeRequest
	}

	var requests []DeleteLogAnalyticsEntityTypeRequestInfo
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
			response, err := c.DeleteLogAnalyticsEntityType(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientDeleteLogAnalyticsLogGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "DeleteLogAnalyticsLogGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteLogAnalyticsLogGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "DeleteLogAnalyticsLogGroup", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "DeleteLogAnalyticsLogGroup")
	assert.NoError(t, err)

	type DeleteLogAnalyticsLogGroupRequestInfo struct {
		ContainerId string
		Request     loganalytics.DeleteLogAnalyticsLogGroupRequest
	}

	var requests []DeleteLogAnalyticsLogGroupRequestInfo
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
			response, err := c.DeleteLogAnalyticsLogGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientDeleteLogAnalyticsObjectCollectionRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "DeleteLogAnalyticsObjectCollectionRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteLogAnalyticsObjectCollectionRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "DeleteLogAnalyticsObjectCollectionRule", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "DeleteLogAnalyticsObjectCollectionRule")
	assert.NoError(t, err)

	type DeleteLogAnalyticsObjectCollectionRuleRequestInfo struct {
		ContainerId string
		Request     loganalytics.DeleteLogAnalyticsObjectCollectionRuleRequest
	}

	var requests []DeleteLogAnalyticsObjectCollectionRuleRequestInfo
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
			response, err := c.DeleteLogAnalyticsObjectCollectionRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientDeleteParser(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "DeleteParser")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteParser is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "DeleteParser", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "DeleteParser")
	assert.NoError(t, err)

	type DeleteParserRequestInfo struct {
		ContainerId string
		Request     loganalytics.DeleteParserRequest
	}

	var requests []DeleteParserRequestInfo
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
			response, err := c.DeleteParser(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientDeleteScheduledTask(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "DeleteScheduledTask")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteScheduledTask is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "DeleteScheduledTask", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "DeleteScheduledTask")
	assert.NoError(t, err)

	type DeleteScheduledTaskRequestInfo struct {
		ContainerId string
		Request     loganalytics.DeleteScheduledTaskRequest
	}

	var requests []DeleteScheduledTaskRequestInfo
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
			response, err := c.DeleteScheduledTask(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientDeleteSource(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "DeleteSource")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSource is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "DeleteSource", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "DeleteSource")
	assert.NoError(t, err)

	type DeleteSourceRequestInfo struct {
		ContainerId string
		Request     loganalytics.DeleteSourceRequest
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

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientDeleteUpload(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "DeleteUpload")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteUpload is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "DeleteUpload", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "DeleteUpload")
	assert.NoError(t, err)

	type DeleteUploadRequestInfo struct {
		ContainerId string
		Request     loganalytics.DeleteUploadRequest
	}

	var requests []DeleteUploadRequestInfo
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
			response, err := c.DeleteUpload(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientDeleteUploadFile(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "DeleteUploadFile")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteUploadFile is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "DeleteUploadFile", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "DeleteUploadFile")
	assert.NoError(t, err)

	type DeleteUploadFileRequestInfo struct {
		ContainerId string
		Request     loganalytics.DeleteUploadFileRequest
	}

	var requests []DeleteUploadFileRequestInfo
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
			response, err := c.DeleteUploadFile(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientDeleteUploadWarning(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "DeleteUploadWarning")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteUploadWarning is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "DeleteUploadWarning", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "DeleteUploadWarning")
	assert.NoError(t, err)

	type DeleteUploadWarningRequestInfo struct {
		ContainerId string
		Request     loganalytics.DeleteUploadWarningRequest
	}

	var requests []DeleteUploadWarningRequestInfo
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
			response, err := c.DeleteUploadWarning(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientDisableArchiving(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "DisableArchiving")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DisableArchiving is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "DisableArchiving", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "DisableArchiving")
	assert.NoError(t, err)

	type DisableArchivingRequestInfo struct {
		ContainerId string
		Request     loganalytics.DisableArchivingRequest
	}

	var requests []DisableArchivingRequestInfo
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
			response, err := c.DisableArchiving(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientEnableArchiving(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "EnableArchiving")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("EnableArchiving is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "EnableArchiving", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "EnableArchiving")
	assert.NoError(t, err)

	type EnableArchivingRequestInfo struct {
		ContainerId string
		Request     loganalytics.EnableArchivingRequest
	}

	var requests []EnableArchivingRequestInfo
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
			response, err := c.EnableArchiving(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientEstimatePurgeDataSize(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "EstimatePurgeDataSize")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("EstimatePurgeDataSize is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "EstimatePurgeDataSize", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "EstimatePurgeDataSize")
	assert.NoError(t, err)

	type EstimatePurgeDataSizeRequestInfo struct {
		ContainerId string
		Request     loganalytics.EstimatePurgeDataSizeRequest
	}

	var requests []EstimatePurgeDataSizeRequestInfo
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
			response, err := c.EstimatePurgeDataSize(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientExportCustomContent(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ExportCustomContent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ExportCustomContent is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ExportCustomContent", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ExportCustomContent")
	assert.NoError(t, err)

	type ExportCustomContentRequestInfo struct {
		ContainerId string
		Request     loganalytics.ExportCustomContentRequest
	}

	var requests []ExportCustomContentRequestInfo
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
			response, err := c.ExportCustomContent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientExportQueryResult(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ExportQueryResult")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ExportQueryResult is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ExportQueryResult", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ExportQueryResult")
	assert.NoError(t, err)

	type ExportQueryResultRequestInfo struct {
		ContainerId string
		Request     loganalytics.ExportQueryResultRequest
	}

	var requests []ExportQueryResultRequestInfo
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
			response, err := c.ExportQueryResult(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientExtractStructuredLogFieldPaths(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ExtractStructuredLogFieldPaths")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ExtractStructuredLogFieldPaths is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ExtractStructuredLogFieldPaths", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ExtractStructuredLogFieldPaths")
	assert.NoError(t, err)

	type ExtractStructuredLogFieldPathsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ExtractStructuredLogFieldPathsRequest
	}

	var requests []ExtractStructuredLogFieldPathsRequestInfo
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
			response, err := c.ExtractStructuredLogFieldPaths(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientExtractStructuredLogHeaderPaths(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ExtractStructuredLogHeaderPaths")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ExtractStructuredLogHeaderPaths is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ExtractStructuredLogHeaderPaths", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ExtractStructuredLogHeaderPaths")
	assert.NoError(t, err)

	type ExtractStructuredLogHeaderPathsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ExtractStructuredLogHeaderPathsRequest
	}

	var requests []ExtractStructuredLogHeaderPathsRequestInfo
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
			response, err := c.ExtractStructuredLogHeaderPaths(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientFilter(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "Filter")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("Filter is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "Filter", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "Filter")
	assert.NoError(t, err)

	type FilterRequestInfo struct {
		ContainerId string
		Request     loganalytics.FilterRequest
	}

	var requests []FilterRequestInfo
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
			response, err := c.Filter(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetAssociationSummary(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetAssociationSummary")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAssociationSummary is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetAssociationSummary", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetAssociationSummary")
	assert.NoError(t, err)

	type GetAssociationSummaryRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetAssociationSummaryRequest
	}

	var requests []GetAssociationSummaryRequestInfo
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
			response, err := c.GetAssociationSummary(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetColumnNames(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetColumnNames")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetColumnNames is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetColumnNames", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetColumnNames")
	assert.NoError(t, err)

	type GetColumnNamesRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetColumnNamesRequest
	}

	var requests []GetColumnNamesRequestInfo
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
			response, err := c.GetColumnNames(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetConfigWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetConfigWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConfigWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetConfigWorkRequest", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetConfigWorkRequest")
	assert.NoError(t, err)

	type GetConfigWorkRequestRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetConfigWorkRequestRequest
	}

	var requests []GetConfigWorkRequestRequestInfo
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
			response, err := c.GetConfigWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetField(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetField")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetField is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetField", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetField")
	assert.NoError(t, err)

	type GetFieldRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetFieldRequest
	}

	var requests []GetFieldRequestInfo
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
			response, err := c.GetField(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetFieldsSummary(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetFieldsSummary")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetFieldsSummary is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetFieldsSummary", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetFieldsSummary")
	assert.NoError(t, err)

	type GetFieldsSummaryRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetFieldsSummaryRequest
	}

	var requests []GetFieldsSummaryRequestInfo
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
			response, err := c.GetFieldsSummary(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetLabel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetLabel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLabel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetLabel", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetLabel")
	assert.NoError(t, err)

	type GetLabelRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetLabelRequest
	}

	var requests []GetLabelRequestInfo
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
			response, err := c.GetLabel(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetLabelSummary(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetLabelSummary")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLabelSummary is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetLabelSummary", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetLabelSummary")
	assert.NoError(t, err)

	type GetLabelSummaryRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetLabelSummaryRequest
	}

	var requests []GetLabelSummaryRequestInfo
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
			response, err := c.GetLabelSummary(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetLogAnalyticsEntitiesSummary(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetLogAnalyticsEntitiesSummary")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLogAnalyticsEntitiesSummary is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetLogAnalyticsEntitiesSummary", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetLogAnalyticsEntitiesSummary")
	assert.NoError(t, err)

	type GetLogAnalyticsEntitiesSummaryRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetLogAnalyticsEntitiesSummaryRequest
	}

	var requests []GetLogAnalyticsEntitiesSummaryRequestInfo
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
			response, err := c.GetLogAnalyticsEntitiesSummary(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetLogAnalyticsEntity(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetLogAnalyticsEntity")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLogAnalyticsEntity is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetLogAnalyticsEntity", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetLogAnalyticsEntity")
	assert.NoError(t, err)

	type GetLogAnalyticsEntityRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetLogAnalyticsEntityRequest
	}

	var requests []GetLogAnalyticsEntityRequestInfo
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
			response, err := c.GetLogAnalyticsEntity(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetLogAnalyticsEntityType(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetLogAnalyticsEntityType")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLogAnalyticsEntityType is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetLogAnalyticsEntityType", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetLogAnalyticsEntityType")
	assert.NoError(t, err)

	type GetLogAnalyticsEntityTypeRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetLogAnalyticsEntityTypeRequest
	}

	var requests []GetLogAnalyticsEntityTypeRequestInfo
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
			response, err := c.GetLogAnalyticsEntityType(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetLogAnalyticsLogGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetLogAnalyticsLogGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLogAnalyticsLogGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetLogAnalyticsLogGroup", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetLogAnalyticsLogGroup")
	assert.NoError(t, err)

	type GetLogAnalyticsLogGroupRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetLogAnalyticsLogGroupRequest
	}

	var requests []GetLogAnalyticsLogGroupRequestInfo
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
			response, err := c.GetLogAnalyticsLogGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetLogAnalyticsLogGroupsSummary(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetLogAnalyticsLogGroupsSummary")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLogAnalyticsLogGroupsSummary is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetLogAnalyticsLogGroupsSummary", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetLogAnalyticsLogGroupsSummary")
	assert.NoError(t, err)

	type GetLogAnalyticsLogGroupsSummaryRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetLogAnalyticsLogGroupsSummaryRequest
	}

	var requests []GetLogAnalyticsLogGroupsSummaryRequestInfo
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
			response, err := c.GetLogAnalyticsLogGroupsSummary(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetLogAnalyticsObjectCollectionRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetLogAnalyticsObjectCollectionRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLogAnalyticsObjectCollectionRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetLogAnalyticsObjectCollectionRule", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetLogAnalyticsObjectCollectionRule")
	assert.NoError(t, err)

	type GetLogAnalyticsObjectCollectionRuleRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetLogAnalyticsObjectCollectionRuleRequest
	}

	var requests []GetLogAnalyticsObjectCollectionRuleRequestInfo
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
			response, err := c.GetLogAnalyticsObjectCollectionRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetNamespace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetNamespace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetNamespace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetNamespace", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetNamespace")
	assert.NoError(t, err)

	type GetNamespaceRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetNamespaceRequest
	}

	var requests []GetNamespaceRequestInfo
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
			response, err := c.GetNamespace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetParser(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetParser")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetParser is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetParser", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetParser")
	assert.NoError(t, err)

	type GetParserRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetParserRequest
	}

	var requests []GetParserRequestInfo
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
			response, err := c.GetParser(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetParserSummary(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetParserSummary")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetParserSummary is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetParserSummary", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetParserSummary")
	assert.NoError(t, err)

	type GetParserSummaryRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetParserSummaryRequest
	}

	var requests []GetParserSummaryRequestInfo
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
			response, err := c.GetParserSummary(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetQueryResult(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetQueryResult")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetQueryResult is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetQueryResult", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetQueryResult")
	assert.NoError(t, err)

	type GetQueryResultRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetQueryResultRequest
	}

	var requests []GetQueryResultRequestInfo
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
				r := req.(*loganalytics.GetQueryResultRequest)
				return c.GetQueryResult(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.GetQueryResultResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.GetQueryResultResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetQueryWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetQueryWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetQueryWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetQueryWorkRequest", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetQueryWorkRequest")
	assert.NoError(t, err)

	type GetQueryWorkRequestRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetQueryWorkRequestRequest
	}

	var requests []GetQueryWorkRequestRequestInfo
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
			response, err := c.GetQueryWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetScheduledTask(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetScheduledTask")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetScheduledTask is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetScheduledTask", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetScheduledTask")
	assert.NoError(t, err)

	type GetScheduledTaskRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetScheduledTaskRequest
	}

	var requests []GetScheduledTaskRequestInfo
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
			response, err := c.GetScheduledTask(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetSource(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetSource")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSource is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetSource", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetSource")
	assert.NoError(t, err)

	type GetSourceRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetSourceRequest
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

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetSourceSummary(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetSourceSummary")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSourceSummary is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetSourceSummary", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetSourceSummary")
	assert.NoError(t, err)

	type GetSourceSummaryRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetSourceSummaryRequest
	}

	var requests []GetSourceSummaryRequestInfo
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
			response, err := c.GetSourceSummary(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetStorage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetStorage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStorage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetStorage", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetStorage")
	assert.NoError(t, err)

	type GetStorageRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetStorageRequest
	}

	var requests []GetStorageRequestInfo
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
			response, err := c.GetStorage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetStorageUsage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetStorageUsage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStorageUsage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetStorageUsage", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetStorageUsage")
	assert.NoError(t, err)

	type GetStorageUsageRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetStorageUsageRequest
	}

	var requests []GetStorageUsageRequestInfo
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
			response, err := c.GetStorageUsage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetStorageWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetStorageWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStorageWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetStorageWorkRequest", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetStorageWorkRequest")
	assert.NoError(t, err)

	type GetStorageWorkRequestRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetStorageWorkRequestRequest
	}

	var requests []GetStorageWorkRequestRequestInfo
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
			response, err := c.GetStorageWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetUpload(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetUpload")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetUpload is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetUpload", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetUpload")
	assert.NoError(t, err)

	type GetUploadRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetUploadRequest
	}

	var requests []GetUploadRequestInfo
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
			response, err := c.GetUpload(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "GetWorkRequest", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     loganalytics.GetWorkRequestRequest
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

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientImportCustomContent(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ImportCustomContent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ImportCustomContent is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ImportCustomContent", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ImportCustomContent")
	assert.NoError(t, err)

	type ImportCustomContentRequestInfo struct {
		ContainerId string
		Request     loganalytics.ImportCustomContentRequest
	}

	var requests []ImportCustomContentRequestInfo
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
			response, err := c.ImportCustomContent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListAssociatedEntities(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListAssociatedEntities")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAssociatedEntities is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListAssociatedEntities", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListAssociatedEntities")
	assert.NoError(t, err)

	type ListAssociatedEntitiesRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListAssociatedEntitiesRequest
	}

	var requests []ListAssociatedEntitiesRequestInfo
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
				r := req.(*loganalytics.ListAssociatedEntitiesRequest)
				return c.ListAssociatedEntities(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListAssociatedEntitiesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListAssociatedEntitiesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListConfigWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListConfigWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListConfigWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListConfigWorkRequests", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListConfigWorkRequests")
	assert.NoError(t, err)

	type ListConfigWorkRequestsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListConfigWorkRequestsRequest
	}

	var requests []ListConfigWorkRequestsRequestInfo
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
				r := req.(*loganalytics.ListConfigWorkRequestsRequest)
				return c.ListConfigWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListConfigWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListConfigWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListEntityAssociations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListEntityAssociations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListEntityAssociations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListEntityAssociations", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListEntityAssociations")
	assert.NoError(t, err)

	type ListEntityAssociationsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListEntityAssociationsRequest
	}

	var requests []ListEntityAssociationsRequestInfo
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
				r := req.(*loganalytics.ListEntityAssociationsRequest)
				return c.ListEntityAssociations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListEntityAssociationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListEntityAssociationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListEntitySourceAssociations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListEntitySourceAssociations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListEntitySourceAssociations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListEntitySourceAssociations", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListEntitySourceAssociations")
	assert.NoError(t, err)

	type ListEntitySourceAssociationsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListEntitySourceAssociationsRequest
	}

	var requests []ListEntitySourceAssociationsRequestInfo
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
				r := req.(*loganalytics.ListEntitySourceAssociationsRequest)
				return c.ListEntitySourceAssociations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListEntitySourceAssociationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListEntitySourceAssociationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListFields(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListFields")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListFields is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListFields", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListFields")
	assert.NoError(t, err)

	type ListFieldsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListFieldsRequest
	}

	var requests []ListFieldsRequestInfo
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
				r := req.(*loganalytics.ListFieldsRequest)
				return c.ListFields(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListFieldsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListFieldsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListLabelPriorities(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListLabelPriorities")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLabelPriorities is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListLabelPriorities", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListLabelPriorities")
	assert.NoError(t, err)

	type ListLabelPrioritiesRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListLabelPrioritiesRequest
	}

	var requests []ListLabelPrioritiesRequestInfo
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
				r := req.(*loganalytics.ListLabelPrioritiesRequest)
				return c.ListLabelPriorities(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListLabelPrioritiesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListLabelPrioritiesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListLabelSourceDetails(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListLabelSourceDetails")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLabelSourceDetails is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListLabelSourceDetails", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListLabelSourceDetails")
	assert.NoError(t, err)

	type ListLabelSourceDetailsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListLabelSourceDetailsRequest
	}

	var requests []ListLabelSourceDetailsRequestInfo
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
				r := req.(*loganalytics.ListLabelSourceDetailsRequest)
				return c.ListLabelSourceDetails(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListLabelSourceDetailsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListLabelSourceDetailsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListLabels(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListLabels")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLabels is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListLabels", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListLabels")
	assert.NoError(t, err)

	type ListLabelsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListLabelsRequest
	}

	var requests []ListLabelsRequestInfo
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
				r := req.(*loganalytics.ListLabelsRequest)
				return c.ListLabels(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListLabelsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListLabelsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListLogAnalyticsEntities(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListLogAnalyticsEntities")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLogAnalyticsEntities is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListLogAnalyticsEntities", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListLogAnalyticsEntities")
	assert.NoError(t, err)

	type ListLogAnalyticsEntitiesRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListLogAnalyticsEntitiesRequest
	}

	var requests []ListLogAnalyticsEntitiesRequestInfo
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
				r := req.(*loganalytics.ListLogAnalyticsEntitiesRequest)
				return c.ListLogAnalyticsEntities(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListLogAnalyticsEntitiesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListLogAnalyticsEntitiesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListLogAnalyticsEntityTypes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListLogAnalyticsEntityTypes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLogAnalyticsEntityTypes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListLogAnalyticsEntityTypes", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListLogAnalyticsEntityTypes")
	assert.NoError(t, err)

	type ListLogAnalyticsEntityTypesRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListLogAnalyticsEntityTypesRequest
	}

	var requests []ListLogAnalyticsEntityTypesRequestInfo
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
				r := req.(*loganalytics.ListLogAnalyticsEntityTypesRequest)
				return c.ListLogAnalyticsEntityTypes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListLogAnalyticsEntityTypesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListLogAnalyticsEntityTypesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListLogAnalyticsLogGroups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListLogAnalyticsLogGroups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLogAnalyticsLogGroups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListLogAnalyticsLogGroups", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListLogAnalyticsLogGroups")
	assert.NoError(t, err)

	type ListLogAnalyticsLogGroupsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListLogAnalyticsLogGroupsRequest
	}

	var requests []ListLogAnalyticsLogGroupsRequestInfo
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
				r := req.(*loganalytics.ListLogAnalyticsLogGroupsRequest)
				return c.ListLogAnalyticsLogGroups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListLogAnalyticsLogGroupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListLogAnalyticsLogGroupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListLogAnalyticsObjectCollectionRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListLogAnalyticsObjectCollectionRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLogAnalyticsObjectCollectionRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListLogAnalyticsObjectCollectionRules", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListLogAnalyticsObjectCollectionRules")
	assert.NoError(t, err)

	type ListLogAnalyticsObjectCollectionRulesRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListLogAnalyticsObjectCollectionRulesRequest
	}

	var requests []ListLogAnalyticsObjectCollectionRulesRequestInfo
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
				r := req.(*loganalytics.ListLogAnalyticsObjectCollectionRulesRequest)
				return c.ListLogAnalyticsObjectCollectionRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListLogAnalyticsObjectCollectionRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListLogAnalyticsObjectCollectionRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListMetaSourceTypes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListMetaSourceTypes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListMetaSourceTypes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListMetaSourceTypes", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListMetaSourceTypes")
	assert.NoError(t, err)

	type ListMetaSourceTypesRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListMetaSourceTypesRequest
	}

	var requests []ListMetaSourceTypesRequestInfo
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
				r := req.(*loganalytics.ListMetaSourceTypesRequest)
				return c.ListMetaSourceTypes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListMetaSourceTypesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListMetaSourceTypesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListNamespaces(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListNamespaces")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListNamespaces is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListNamespaces", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListNamespaces")
	assert.NoError(t, err)

	type ListNamespacesRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListNamespacesRequest
	}

	var requests []ListNamespacesRequestInfo
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
			response, err := c.ListNamespaces(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListParserFunctions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListParserFunctions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListParserFunctions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListParserFunctions", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListParserFunctions")
	assert.NoError(t, err)

	type ListParserFunctionsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListParserFunctionsRequest
	}

	var requests []ListParserFunctionsRequestInfo
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
				r := req.(*loganalytics.ListParserFunctionsRequest)
				return c.ListParserFunctions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListParserFunctionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListParserFunctionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListParserMetaPlugins(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListParserMetaPlugins")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListParserMetaPlugins is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListParserMetaPlugins", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListParserMetaPlugins")
	assert.NoError(t, err)

	type ListParserMetaPluginsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListParserMetaPluginsRequest
	}

	var requests []ListParserMetaPluginsRequestInfo
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
				r := req.(*loganalytics.ListParserMetaPluginsRequest)
				return c.ListParserMetaPlugins(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListParserMetaPluginsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListParserMetaPluginsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListParsers(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListParsers")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListParsers is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListParsers", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListParsers")
	assert.NoError(t, err)

	type ListParsersRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListParsersRequest
	}

	var requests []ListParsersRequestInfo
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
				r := req.(*loganalytics.ListParsersRequest)
				return c.ListParsers(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListParsersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListParsersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListQueryWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListQueryWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListQueryWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListQueryWorkRequests", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListQueryWorkRequests")
	assert.NoError(t, err)

	type ListQueryWorkRequestsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListQueryWorkRequestsRequest
	}

	var requests []ListQueryWorkRequestsRequestInfo
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
				r := req.(*loganalytics.ListQueryWorkRequestsRequest)
				return c.ListQueryWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListQueryWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListQueryWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListScheduledTasks(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListScheduledTasks")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListScheduledTasks is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListScheduledTasks", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListScheduledTasks")
	assert.NoError(t, err)

	type ListScheduledTasksRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListScheduledTasksRequest
	}

	var requests []ListScheduledTasksRequestInfo
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
				r := req.(*loganalytics.ListScheduledTasksRequest)
				return c.ListScheduledTasks(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListScheduledTasksResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListScheduledTasksResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListSourceAssociations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListSourceAssociations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSourceAssociations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListSourceAssociations", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListSourceAssociations")
	assert.NoError(t, err)

	type ListSourceAssociationsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListSourceAssociationsRequest
	}

	var requests []ListSourceAssociationsRequestInfo
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
				r := req.(*loganalytics.ListSourceAssociationsRequest)
				return c.ListSourceAssociations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListSourceAssociationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListSourceAssociationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListSourceExtendedFieldDefinitions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListSourceExtendedFieldDefinitions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSourceExtendedFieldDefinitions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListSourceExtendedFieldDefinitions", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListSourceExtendedFieldDefinitions")
	assert.NoError(t, err)

	type ListSourceExtendedFieldDefinitionsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListSourceExtendedFieldDefinitionsRequest
	}

	var requests []ListSourceExtendedFieldDefinitionsRequestInfo
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
				r := req.(*loganalytics.ListSourceExtendedFieldDefinitionsRequest)
				return c.ListSourceExtendedFieldDefinitions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListSourceExtendedFieldDefinitionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListSourceExtendedFieldDefinitionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListSourceLabelOperators(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListSourceLabelOperators")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSourceLabelOperators is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListSourceLabelOperators", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListSourceLabelOperators")
	assert.NoError(t, err)

	type ListSourceLabelOperatorsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListSourceLabelOperatorsRequest
	}

	var requests []ListSourceLabelOperatorsRequestInfo
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
				r := req.(*loganalytics.ListSourceLabelOperatorsRequest)
				return c.ListSourceLabelOperators(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListSourceLabelOperatorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListSourceLabelOperatorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListSourceMetaFunctions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListSourceMetaFunctions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSourceMetaFunctions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListSourceMetaFunctions", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListSourceMetaFunctions")
	assert.NoError(t, err)

	type ListSourceMetaFunctionsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListSourceMetaFunctionsRequest
	}

	var requests []ListSourceMetaFunctionsRequestInfo
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
				r := req.(*loganalytics.ListSourceMetaFunctionsRequest)
				return c.ListSourceMetaFunctions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListSourceMetaFunctionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListSourceMetaFunctionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListSourcePatterns(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListSourcePatterns")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSourcePatterns is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListSourcePatterns", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListSourcePatterns")
	assert.NoError(t, err)

	type ListSourcePatternsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListSourcePatternsRequest
	}

	var requests []ListSourcePatternsRequestInfo
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
				r := req.(*loganalytics.ListSourcePatternsRequest)
				return c.ListSourcePatterns(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListSourcePatternsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListSourcePatternsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListSources(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListSources")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSources is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListSources", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListSources")
	assert.NoError(t, err)

	type ListSourcesRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListSourcesRequest
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
				r := req.(*loganalytics.ListSourcesRequest)
				return c.ListSources(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListSourcesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListSourcesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListStorageWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListStorageWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListStorageWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListStorageWorkRequestErrors", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListStorageWorkRequestErrors")
	assert.NoError(t, err)

	type ListStorageWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListStorageWorkRequestErrorsRequest
	}

	var requests []ListStorageWorkRequestErrorsRequestInfo
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
				r := req.(*loganalytics.ListStorageWorkRequestErrorsRequest)
				return c.ListStorageWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListStorageWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListStorageWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListStorageWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListStorageWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListStorageWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListStorageWorkRequests", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListStorageWorkRequests")
	assert.NoError(t, err)

	type ListStorageWorkRequestsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListStorageWorkRequestsRequest
	}

	var requests []ListStorageWorkRequestsRequestInfo
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
				r := req.(*loganalytics.ListStorageWorkRequestsRequest)
				return c.ListStorageWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListStorageWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListStorageWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListSupportedCharEncodings(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListSupportedCharEncodings")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSupportedCharEncodings is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListSupportedCharEncodings", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListSupportedCharEncodings")
	assert.NoError(t, err)

	type ListSupportedCharEncodingsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListSupportedCharEncodingsRequest
	}

	var requests []ListSupportedCharEncodingsRequestInfo
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
				r := req.(*loganalytics.ListSupportedCharEncodingsRequest)
				return c.ListSupportedCharEncodings(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListSupportedCharEncodingsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListSupportedCharEncodingsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListSupportedTimezones(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListSupportedTimezones")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSupportedTimezones is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListSupportedTimezones", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListSupportedTimezones")
	assert.NoError(t, err)

	type ListSupportedTimezonesRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListSupportedTimezonesRequest
	}

	var requests []ListSupportedTimezonesRequestInfo
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
				r := req.(*loganalytics.ListSupportedTimezonesRequest)
				return c.ListSupportedTimezones(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListSupportedTimezonesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListSupportedTimezonesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListUploadFiles(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListUploadFiles")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListUploadFiles is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListUploadFiles", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListUploadFiles")
	assert.NoError(t, err)

	type ListUploadFilesRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListUploadFilesRequest
	}

	var requests []ListUploadFilesRequestInfo
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
				r := req.(*loganalytics.ListUploadFilesRequest)
				return c.ListUploadFiles(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListUploadFilesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListUploadFilesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListUploadWarnings(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListUploadWarnings")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListUploadWarnings is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListUploadWarnings", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListUploadWarnings")
	assert.NoError(t, err)

	type ListUploadWarningsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListUploadWarningsRequest
	}

	var requests []ListUploadWarningsRequestInfo
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
				r := req.(*loganalytics.ListUploadWarningsRequest)
				return c.ListUploadWarnings(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListUploadWarningsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListUploadWarningsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListUploads(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListUploads")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListUploads is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListUploads", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListUploads")
	assert.NoError(t, err)

	type ListUploadsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListUploadsRequest
	}

	var requests []ListUploadsRequestInfo
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
				r := req.(*loganalytics.ListUploadsRequest)
				return c.ListUploads(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListUploadsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListUploadsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListWorkRequestErrors", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListWorkRequestErrorsRequest
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
				r := req.(*loganalytics.ListWorkRequestErrorsRequest)
				return c.ListWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListWorkRequestLogs", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListWorkRequestLogsRequest
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
				r := req.(*loganalytics.ListWorkRequestLogsRequest)
				return c.ListWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ListWorkRequests", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ListWorkRequestsRequest
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
				r := req.(*loganalytics.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientOffboardNamespace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "OffboardNamespace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("OffboardNamespace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "OffboardNamespace", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "OffboardNamespace")
	assert.NoError(t, err)

	type OffboardNamespaceRequestInfo struct {
		ContainerId string
		Request     loganalytics.OffboardNamespaceRequest
	}

	var requests []OffboardNamespaceRequestInfo
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
			response, err := c.OffboardNamespace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientOnboardNamespace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "OnboardNamespace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("OnboardNamespace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "OnboardNamespace", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "OnboardNamespace")
	assert.NoError(t, err)

	type OnboardNamespaceRequestInfo struct {
		ContainerId string
		Request     loganalytics.OnboardNamespaceRequest
	}

	var requests []OnboardNamespaceRequestInfo
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
			response, err := c.OnboardNamespace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientParseQuery(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ParseQuery")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ParseQuery is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ParseQuery", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ParseQuery")
	assert.NoError(t, err)

	type ParseQueryRequestInfo struct {
		ContainerId string
		Request     loganalytics.ParseQueryRequest
	}

	var requests []ParseQueryRequestInfo
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
			response, err := c.ParseQuery(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientPurgeStorageData(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "PurgeStorageData")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PurgeStorageData is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "PurgeStorageData", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "PurgeStorageData")
	assert.NoError(t, err)

	type PurgeStorageDataRequestInfo struct {
		ContainerId string
		Request     loganalytics.PurgeStorageDataRequest
	}

	var requests []PurgeStorageDataRequestInfo
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
			response, err := c.PurgeStorageData(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientPutQueryWorkRequestBackground(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "PutQueryWorkRequestBackground")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PutQueryWorkRequestBackground is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "PutQueryWorkRequestBackground", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "PutQueryWorkRequestBackground")
	assert.NoError(t, err)

	type PutQueryWorkRequestBackgroundRequestInfo struct {
		ContainerId string
		Request     loganalytics.PutQueryWorkRequestBackgroundRequest
	}

	var requests []PutQueryWorkRequestBackgroundRequestInfo
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
			response, err := c.PutQueryWorkRequestBackground(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientQuery(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "Query")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("Query is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "Query", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "Query")
	assert.NoError(t, err)

	type QueryRequestInfo struct {
		ContainerId string
		Request     loganalytics.QueryRequest
	}

	var requests []QueryRequestInfo
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
				r := req.(*loganalytics.QueryRequest)
				return c.Query(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.QueryResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.QueryResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientRecallArchivedData(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "RecallArchivedData")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RecallArchivedData is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "RecallArchivedData", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "RecallArchivedData")
	assert.NoError(t, err)

	type RecallArchivedDataRequestInfo struct {
		ContainerId string
		Request     loganalytics.RecallArchivedDataRequest
	}

	var requests []RecallArchivedDataRequestInfo
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
			response, err := c.RecallArchivedData(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientRegisterLookup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "RegisterLookup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RegisterLookup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "RegisterLookup", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "RegisterLookup")
	assert.NoError(t, err)

	type RegisterLookupRequestInfo struct {
		ContainerId string
		Request     loganalytics.RegisterLookupRequest
	}

	var requests []RegisterLookupRequestInfo
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
			response, err := c.RegisterLookup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientReleaseRecalledData(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ReleaseRecalledData")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ReleaseRecalledData is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ReleaseRecalledData", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ReleaseRecalledData")
	assert.NoError(t, err)

	type ReleaseRecalledDataRequestInfo struct {
		ContainerId string
		Request     loganalytics.ReleaseRecalledDataRequest
	}

	var requests []ReleaseRecalledDataRequestInfo
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
			response, err := c.ReleaseRecalledData(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientRemoveEntityAssociations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "RemoveEntityAssociations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RemoveEntityAssociations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "RemoveEntityAssociations", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "RemoveEntityAssociations")
	assert.NoError(t, err)

	type RemoveEntityAssociationsRequestInfo struct {
		ContainerId string
		Request     loganalytics.RemoveEntityAssociationsRequest
	}

	var requests []RemoveEntityAssociationsRequestInfo
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
			response, err := c.RemoveEntityAssociations(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientRun(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "Run")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("Run is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "Run", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "Run")
	assert.NoError(t, err)

	type RunRequestInfo struct {
		ContainerId string
		Request     loganalytics.RunRequest
	}

	var requests []RunRequestInfo
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
			response, err := c.Run(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientSuggest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "Suggest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("Suggest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "Suggest", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "Suggest")
	assert.NoError(t, err)

	type SuggestRequestInfo struct {
		ContainerId string
		Request     loganalytics.SuggestRequest
	}

	var requests []SuggestRequestInfo
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
			response, err := c.Suggest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientTestParser(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "TestParser")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("TestParser is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "TestParser", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "TestParser")
	assert.NoError(t, err)

	type TestParserRequestInfo struct {
		ContainerId string
		Request     loganalytics.TestParserRequest
	}

	var requests []TestParserRequestInfo
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
			response, err := c.TestParser(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientUpdateLogAnalyticsEntity(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "UpdateLogAnalyticsEntity")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateLogAnalyticsEntity is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "UpdateLogAnalyticsEntity", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "UpdateLogAnalyticsEntity")
	assert.NoError(t, err)

	type UpdateLogAnalyticsEntityRequestInfo struct {
		ContainerId string
		Request     loganalytics.UpdateLogAnalyticsEntityRequest
	}

	var requests []UpdateLogAnalyticsEntityRequestInfo
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
			response, err := c.UpdateLogAnalyticsEntity(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientUpdateLogAnalyticsEntityType(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "UpdateLogAnalyticsEntityType")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateLogAnalyticsEntityType is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "UpdateLogAnalyticsEntityType", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "UpdateLogAnalyticsEntityType")
	assert.NoError(t, err)

	type UpdateLogAnalyticsEntityTypeRequestInfo struct {
		ContainerId string
		Request     loganalytics.UpdateLogAnalyticsEntityTypeRequest
	}

	var requests []UpdateLogAnalyticsEntityTypeRequestInfo
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
			response, err := c.UpdateLogAnalyticsEntityType(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientUpdateLogAnalyticsLogGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "UpdateLogAnalyticsLogGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateLogAnalyticsLogGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "UpdateLogAnalyticsLogGroup", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "UpdateLogAnalyticsLogGroup")
	assert.NoError(t, err)

	type UpdateLogAnalyticsLogGroupRequestInfo struct {
		ContainerId string
		Request     loganalytics.UpdateLogAnalyticsLogGroupRequest
	}

	var requests []UpdateLogAnalyticsLogGroupRequestInfo
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
			response, err := c.UpdateLogAnalyticsLogGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientUpdateLogAnalyticsObjectCollectionRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "UpdateLogAnalyticsObjectCollectionRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateLogAnalyticsObjectCollectionRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "UpdateLogAnalyticsObjectCollectionRule", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "UpdateLogAnalyticsObjectCollectionRule")
	assert.NoError(t, err)

	type UpdateLogAnalyticsObjectCollectionRuleRequestInfo struct {
		ContainerId string
		Request     loganalytics.UpdateLogAnalyticsObjectCollectionRuleRequest
	}

	var requests []UpdateLogAnalyticsObjectCollectionRuleRequestInfo
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
			response, err := c.UpdateLogAnalyticsObjectCollectionRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientUpdateScheduledTask(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "UpdateScheduledTask")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateScheduledTask is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "UpdateScheduledTask", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "UpdateScheduledTask")
	assert.NoError(t, err)

	type UpdateScheduledTaskRequestInfo struct {
		ContainerId string
		Request     loganalytics.UpdateScheduledTaskRequest
	}

	var requests []UpdateScheduledTaskRequestInfo
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
			response, err := c.UpdateScheduledTask(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientUpdateStorage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "UpdateStorage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateStorage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "UpdateStorage", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "UpdateStorage")
	assert.NoError(t, err)

	type UpdateStorageRequestInfo struct {
		ContainerId string
		Request     loganalytics.UpdateStorageRequest
	}

	var requests []UpdateStorageRequestInfo
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
			response, err := c.UpdateStorage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientUploadLogFile(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "UploadLogFile")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UploadLogFile is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "UploadLogFile", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "UploadLogFile")
	assert.NoError(t, err)

	type UploadLogFileRequestInfo struct {
		ContainerId string
		Request     loganalytics.UploadLogFileRequest
	}

	var requests []UploadLogFileRequestInfo
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
			response, err := c.UploadLogFile(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientUpsertAssociations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "UpsertAssociations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpsertAssociations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "UpsertAssociations", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "UpsertAssociations")
	assert.NoError(t, err)

	type UpsertAssociationsRequestInfo struct {
		ContainerId string
		Request     loganalytics.UpsertAssociationsRequest
	}

	var requests []UpsertAssociationsRequestInfo
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
			response, err := c.UpsertAssociations(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientUpsertField(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "UpsertField")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpsertField is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "UpsertField", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "UpsertField")
	assert.NoError(t, err)

	type UpsertFieldRequestInfo struct {
		ContainerId string
		Request     loganalytics.UpsertFieldRequest
	}

	var requests []UpsertFieldRequestInfo
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
			response, err := c.UpsertField(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientUpsertLabel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "UpsertLabel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpsertLabel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "UpsertLabel", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "UpsertLabel")
	assert.NoError(t, err)

	type UpsertLabelRequestInfo struct {
		ContainerId string
		Request     loganalytics.UpsertLabelRequest
	}

	var requests []UpsertLabelRequestInfo
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
			response, err := c.UpsertLabel(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientUpsertParser(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "UpsertParser")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpsertParser is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "UpsertParser", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "UpsertParser")
	assert.NoError(t, err)

	type UpsertParserRequestInfo struct {
		ContainerId string
		Request     loganalytics.UpsertParserRequest
	}

	var requests []UpsertParserRequestInfo
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
			response, err := c.UpsertParser(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientUpsertSource(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "UpsertSource")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpsertSource is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "UpsertSource", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "UpsertSource")
	assert.NoError(t, err)

	type UpsertSourceRequestInfo struct {
		ContainerId string
		Request     loganalytics.UpsertSourceRequest
	}

	var requests []UpsertSourceRequestInfo
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
			response, err := c.UpsertSource(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientValidateAssociationParameters(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ValidateAssociationParameters")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ValidateAssociationParameters is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ValidateAssociationParameters", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ValidateAssociationParameters")
	assert.NoError(t, err)

	type ValidateAssociationParametersRequestInfo struct {
		ContainerId string
		Request     loganalytics.ValidateAssociationParametersRequest
	}

	var requests []ValidateAssociationParametersRequestInfo
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
				r := req.(*loganalytics.ValidateAssociationParametersRequest)
				return c.ValidateAssociationParameters(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loganalytics.ValidateAssociationParametersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loganalytics.ValidateAssociationParametersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientValidateFile(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ValidateFile")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ValidateFile is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ValidateFile", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ValidateFile")
	assert.NoError(t, err)

	type ValidateFileRequestInfo struct {
		ContainerId string
		Request     loganalytics.ValidateFileRequest
	}

	var requests []ValidateFileRequestInfo
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
			response, err := c.ValidateFile(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientValidateSource(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ValidateSource")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ValidateSource is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ValidateSource", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ValidateSource")
	assert.NoError(t, err)

	type ValidateSourceRequestInfo struct {
		ContainerId string
		Request     loganalytics.ValidateSourceRequest
	}

	var requests []ValidateSourceRequestInfo
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
			response, err := c.ValidateSource(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientValidateSourceExtendedFieldDetails(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ValidateSourceExtendedFieldDetails")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ValidateSourceExtendedFieldDetails is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ValidateSourceExtendedFieldDetails", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ValidateSourceExtendedFieldDetails")
	assert.NoError(t, err)

	type ValidateSourceExtendedFieldDetailsRequestInfo struct {
		ContainerId string
		Request     loganalytics.ValidateSourceExtendedFieldDetailsRequest
	}

	var requests []ValidateSourceExtendedFieldDetailsRequestInfo
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
			response, err := c.ValidateSourceExtendedFieldDetails(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="omc_loganalytics_dev_ww_grp@oracle.com" jiraProject="LOGAN" opsJiraProject="LOGAN"
func TestLoganalyticsLogAnalyticsClientValidateSourceMapping(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loganalytics", "ValidateSourceMapping")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ValidateSourceMapping is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loganalytics", "LogAnalytics", "ValidateSourceMapping", createLoganalyticsLogAnalyticsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loganalytics.LogAnalyticsClient)

	body, err := testClient.getRequests("loganalytics", "ValidateSourceMapping")
	assert.NoError(t, err)

	type ValidateSourceMappingRequestInfo struct {
		ContainerId string
		Request     loganalytics.ValidateSourceMappingRequest
	}

	var requests []ValidateSourceMappingRequestInfo
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
			response, err := c.ValidateSourceMapping(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
