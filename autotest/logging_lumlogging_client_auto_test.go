package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/logging"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createLumLoggingClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := logging.NewLumLoggingClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientChangeLogGroupCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "ChangeLogGroupCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeLogGroupCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "ChangeLogGroupCompartment", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "ChangeLogGroupCompartment")
	assert.NoError(t, err)

	type ChangeLogGroupCompartmentRequestInfo struct {
		ContainerId string
		Request     logging.ChangeLogGroupCompartmentRequest
	}

	var requests []ChangeLogGroupCompartmentRequestInfo
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
			response, err := c.ChangeLogGroupCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientChangeLogLogGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "ChangeLogLogGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeLogLogGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "ChangeLogLogGroup", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "ChangeLogLogGroup")
	assert.NoError(t, err)

	type ChangeLogLogGroupRequestInfo struct {
		ContainerId string
		Request     logging.ChangeLogLogGroupRequest
	}

	var requests []ChangeLogLogGroupRequestInfo
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
			response, err := c.ChangeLogLogGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientChangeLogRuleCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "ChangeLogRuleCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeLogRuleCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "ChangeLogRuleCompartment", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "ChangeLogRuleCompartment")
	assert.NoError(t, err)

	type ChangeLogRuleCompartmentRequestInfo struct {
		ContainerId string
		Request     logging.ChangeLogRuleCompartmentRequest
	}

	var requests []ChangeLogRuleCompartmentRequestInfo
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
			response, err := c.ChangeLogRuleCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientCreateLog(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "CreateLog")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateLog is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "CreateLog", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "CreateLog")
	assert.NoError(t, err)

	type CreateLogRequestInfo struct {
		ContainerId string
		Request     logging.CreateLogRequest
	}

	var requests []CreateLogRequestInfo
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
			response, err := c.CreateLog(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientCreateLogGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "CreateLogGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateLogGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "CreateLogGroup", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "CreateLogGroup")
	assert.NoError(t, err)

	type CreateLogGroupRequestInfo struct {
		ContainerId string
		Request     logging.CreateLogGroupRequest
	}

	var requests []CreateLogGroupRequestInfo
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
			response, err := c.CreateLogGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientCreateLogRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "CreateLogRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateLogRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "CreateLogRule", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "CreateLogRule")
	assert.NoError(t, err)

	type CreateLogRuleRequestInfo struct {
		ContainerId string
		Request     logging.CreateLogRuleRequest
	}

	var requests []CreateLogRuleRequestInfo
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
			response, err := c.CreateLogRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientDeleteLog(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "DeleteLog")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteLog is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "DeleteLog", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "DeleteLog")
	assert.NoError(t, err)

	type DeleteLogRequestInfo struct {
		ContainerId string
		Request     logging.DeleteLogRequest
	}

	var requests []DeleteLogRequestInfo
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
			response, err := c.DeleteLog(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientDeleteLogGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "DeleteLogGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteLogGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "DeleteLogGroup", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "DeleteLogGroup")
	assert.NoError(t, err)

	type DeleteLogGroupRequestInfo struct {
		ContainerId string
		Request     logging.DeleteLogGroupRequest
	}

	var requests []DeleteLogGroupRequestInfo
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
			response, err := c.DeleteLogGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientDeleteLogRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "DeleteLogRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteLogRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "DeleteLogRule", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "DeleteLogRule")
	assert.NoError(t, err)

	type DeleteLogRuleRequestInfo struct {
		ContainerId string
		Request     logging.DeleteLogRuleRequest
	}

	var requests []DeleteLogRuleRequestInfo
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
			response, err := c.DeleteLogRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientGetLog(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "GetLog")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLog is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "GetLog", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "GetLog")
	assert.NoError(t, err)

	type GetLogRequestInfo struct {
		ContainerId string
		Request     logging.GetLogRequest
	}

	var requests []GetLogRequestInfo
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
			response, err := c.GetLog(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientGetLogGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "GetLogGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLogGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "GetLogGroup", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "GetLogGroup")
	assert.NoError(t, err)

	type GetLogGroupRequestInfo struct {
		ContainerId string
		Request     logging.GetLogGroupRequest
	}

	var requests []GetLogGroupRequestInfo
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
			response, err := c.GetLogGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientGetLogRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "GetLogRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLogRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "GetLogRule", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "GetLogRule")
	assert.NoError(t, err)

	type GetLogRuleRequestInfo struct {
		ContainerId string
		Request     logging.GetLogRuleRequest
	}

	var requests []GetLogRuleRequestInfo
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
			response, err := c.GetLogRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientListLogGroups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "ListLogGroups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLogGroups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "ListLogGroups", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "ListLogGroups")
	assert.NoError(t, err)

	type ListLogGroupsRequestInfo struct {
		ContainerId string
		Request     logging.ListLogGroupsRequest
	}

	var requests []ListLogGroupsRequestInfo
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
				r := req.(*logging.ListLogGroupsRequest)
				return c.ListLogGroups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]logging.ListLogGroupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(logging.ListLogGroupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientListLogRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "ListLogRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLogRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "ListLogRules", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "ListLogRules")
	assert.NoError(t, err)

	type ListLogRulesRequestInfo struct {
		ContainerId string
		Request     logging.ListLogRulesRequest
	}

	var requests []ListLogRulesRequestInfo
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
				r := req.(*logging.ListLogRulesRequest)
				return c.ListLogRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]logging.ListLogRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(logging.ListLogRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientListLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "ListLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "ListLogs", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "ListLogs")
	assert.NoError(t, err)

	type ListLogsRequestInfo struct {
		ContainerId string
		Request     logging.ListLogsRequest
	}

	var requests []ListLogsRequestInfo
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
				r := req.(*logging.ListLogsRequest)
				return c.ListLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]logging.ListLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(logging.ListLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientListServices(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "ListServices")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListServices is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "ListServices", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "ListServices")
	assert.NoError(t, err)

	type ListServicesRequestInfo struct {
		ContainerId string
		Request     logging.ListServicesRequest
	}

	var requests []ListServicesRequestInfo
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
			response, err := c.ListServices(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientUpdateLog(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "UpdateLog")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateLog is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "UpdateLog", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "UpdateLog")
	assert.NoError(t, err)

	type UpdateLogRequestInfo struct {
		ContainerId string
		Request     logging.UpdateLogRequest
	}

	var requests []UpdateLogRequestInfo
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
			response, err := c.UpdateLog(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientUpdateLogGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "UpdateLogGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateLogGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "UpdateLogGroup", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "UpdateLogGroup")
	assert.NoError(t, err)

	type UpdateLogGroupRequestInfo struct {
		ContainerId string
		Request     logging.UpdateLogGroupRequest
	}

	var requests []UpdateLogGroupRequestInfo
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
			response, err := c.UpdateLogGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="hydra_dev_us_grp@oracle.com" jiraProject="HYD" opsJiraProject="HYD"
func TestLumLoggingClientUpdateLogRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("logging", "UpdateLogRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateLogRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("logging", "LumLogging", "UpdateLogRule", createLumLoggingClientWithProvider)
	assert.NoError(t, err)
	c := cc.(logging.LumLoggingClient)

	body, err := testClient.getRequests("logging", "UpdateLogRule")
	assert.NoError(t, err)

	type UpdateLogRuleRequestInfo struct {
		ContainerId string
		Request     logging.UpdateLogRuleRequest
	}

	var requests []UpdateLogRuleRequestInfo
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
			response, err := c.UpdateLogRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
