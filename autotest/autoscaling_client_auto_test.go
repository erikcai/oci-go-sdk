package autotest

import (
	"github.com/oracle/oci-go-sdk/autoscaling"
	"github.com/oracle/oci-go-sdk/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// IssueRoutingInfo tag="default" email="instance_dev_us_grp@oracle.com" jiraProject="CIM" opsJiraProject="COM"
func TestAutoScalingClientCreateAutoScalingConfiguration(t *testing.T) {
	enabled, err := testClient.isApiEnabled("autoscaling", "CreateAutoScalingConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutoScalingConfiguration is not enabled by the testing service")
	}
	c, err := autoscaling.NewAutoScalingClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("autoscaling", "CreateAutoScalingConfiguration")
	assert.NoError(t, err)

	type CreateAutoScalingConfigurationRequestInfo struct {
		ContainerId string
		Request     autoscaling.CreateAutoScalingConfigurationRequest
	}

	var requests []CreateAutoScalingConfigurationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateAutoScalingConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="instance_dev_us_grp@oracle.com" jiraProject="CIM" opsJiraProject="COM"
func TestAutoScalingClientCreateAutoScalingPolicy(t *testing.T) {
	enabled, err := testClient.isApiEnabled("autoscaling", "CreateAutoScalingPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAutoScalingPolicy is not enabled by the testing service")
	}
	c, err := autoscaling.NewAutoScalingClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("autoscaling", "CreateAutoScalingPolicy")
	assert.NoError(t, err)

	type CreateAutoScalingPolicyRequestInfo struct {
		ContainerId string
		Request     autoscaling.CreateAutoScalingPolicyRequest
	}

	var requests []CreateAutoScalingPolicyRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateAutoScalingPolicyRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateAutoScalingPolicyDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "policyType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"threshold": &autoscaling.CreateThresholdPolicyDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateAutoScalingPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="instance_dev_us_grp@oracle.com" jiraProject="CIM" opsJiraProject="COM"
func TestAutoScalingClientDeleteAutoScalingConfiguration(t *testing.T) {
	enabled, err := testClient.isApiEnabled("autoscaling", "DeleteAutoScalingConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAutoScalingConfiguration is not enabled by the testing service")
	}
	c, err := autoscaling.NewAutoScalingClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("autoscaling", "DeleteAutoScalingConfiguration")
	assert.NoError(t, err)

	type DeleteAutoScalingConfigurationRequestInfo struct {
		ContainerId string
		Request     autoscaling.DeleteAutoScalingConfigurationRequest
	}

	var requests []DeleteAutoScalingConfigurationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteAutoScalingConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="instance_dev_us_grp@oracle.com" jiraProject="CIM" opsJiraProject="COM"
func TestAutoScalingClientDeleteAutoScalingPolicy(t *testing.T) {
	enabled, err := testClient.isApiEnabled("autoscaling", "DeleteAutoScalingPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAutoScalingPolicy is not enabled by the testing service")
	}
	c, err := autoscaling.NewAutoScalingClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("autoscaling", "DeleteAutoScalingPolicy")
	assert.NoError(t, err)

	type DeleteAutoScalingPolicyRequestInfo struct {
		ContainerId string
		Request     autoscaling.DeleteAutoScalingPolicyRequest
	}

	var requests []DeleteAutoScalingPolicyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteAutoScalingPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="instance_dev_us_grp@oracle.com" jiraProject="CIM" opsJiraProject="COM"
func TestAutoScalingClientGetAutoScalingConfiguration(t *testing.T) {
	enabled, err := testClient.isApiEnabled("autoscaling", "GetAutoScalingConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutoScalingConfiguration is not enabled by the testing service")
	}
	c, err := autoscaling.NewAutoScalingClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("autoscaling", "GetAutoScalingConfiguration")
	assert.NoError(t, err)

	type GetAutoScalingConfigurationRequestInfo struct {
		ContainerId string
		Request     autoscaling.GetAutoScalingConfigurationRequest
	}

	var requests []GetAutoScalingConfigurationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAutoScalingConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="instance_dev_us_grp@oracle.com" jiraProject="CIM" opsJiraProject="COM"
func TestAutoScalingClientGetAutoScalingPolicy(t *testing.T) {
	enabled, err := testClient.isApiEnabled("autoscaling", "GetAutoScalingPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAutoScalingPolicy is not enabled by the testing service")
	}
	c, err := autoscaling.NewAutoScalingClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("autoscaling", "GetAutoScalingPolicy")
	assert.NoError(t, err)

	type GetAutoScalingPolicyRequestInfo struct {
		ContainerId string
		Request     autoscaling.GetAutoScalingPolicyRequest
	}

	var requests []GetAutoScalingPolicyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetAutoScalingPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="instance_dev_us_grp@oracle.com" jiraProject="CIM" opsJiraProject="COM"
func TestAutoScalingClientListAutoScalingConfigurations(t *testing.T) {
	enabled, err := testClient.isApiEnabled("autoscaling", "ListAutoScalingConfigurations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutoScalingConfigurations is not enabled by the testing service")
	}
	c, err := autoscaling.NewAutoScalingClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("autoscaling", "ListAutoScalingConfigurations")
	assert.NoError(t, err)

	type ListAutoScalingConfigurationsRequestInfo struct {
		ContainerId string
		Request     autoscaling.ListAutoScalingConfigurationsRequest
	}

	var requests []ListAutoScalingConfigurationsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*autoscaling.ListAutoScalingConfigurationsRequest)
				return c.ListAutoScalingConfigurations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]autoscaling.ListAutoScalingConfigurationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(autoscaling.ListAutoScalingConfigurationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="instance_dev_us_grp@oracle.com" jiraProject="CIM" opsJiraProject="COM"
func TestAutoScalingClientListAutoScalingPolicies(t *testing.T) {
	enabled, err := testClient.isApiEnabled("autoscaling", "ListAutoScalingPolicies")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAutoScalingPolicies is not enabled by the testing service")
	}
	c, err := autoscaling.NewAutoScalingClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("autoscaling", "ListAutoScalingPolicies")
	assert.NoError(t, err)

	type ListAutoScalingPoliciesRequestInfo struct {
		ContainerId string
		Request     autoscaling.ListAutoScalingPoliciesRequest
	}

	var requests []ListAutoScalingPoliciesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*autoscaling.ListAutoScalingPoliciesRequest)
				return c.ListAutoScalingPolicies(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]autoscaling.ListAutoScalingPoliciesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(autoscaling.ListAutoScalingPoliciesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="instance_dev_us_grp@oracle.com" jiraProject="CIM" opsJiraProject="COM"
func TestAutoScalingClientUpdateAutoScalingConfiguration(t *testing.T) {
	enabled, err := testClient.isApiEnabled("autoscaling", "UpdateAutoScalingConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutoScalingConfiguration is not enabled by the testing service")
	}
	c, err := autoscaling.NewAutoScalingClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("autoscaling", "UpdateAutoScalingConfiguration")
	assert.NoError(t, err)

	type UpdateAutoScalingConfigurationRequestInfo struct {
		ContainerId string
		Request     autoscaling.UpdateAutoScalingConfigurationRequest
	}

	var requests []UpdateAutoScalingConfigurationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateAutoScalingConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="instance_dev_us_grp@oracle.com" jiraProject="CIM" opsJiraProject="COM"
func TestAutoScalingClientUpdateAutoScalingPolicy(t *testing.T) {
	enabled, err := testClient.isApiEnabled("autoscaling", "UpdateAutoScalingPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAutoScalingPolicy is not enabled by the testing service")
	}
	c, err := autoscaling.NewAutoScalingClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("autoscaling", "UpdateAutoScalingPolicy")
	assert.NoError(t, err)

	type UpdateAutoScalingPolicyRequestInfo struct {
		ContainerId string
		Request     autoscaling.UpdateAutoScalingPolicyRequest
	}

	var requests []UpdateAutoScalingPolicyRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]UpdateAutoScalingPolicyRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["UpdateAutoScalingPolicyDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "policyType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"threshold": &autoscaling.UpdateThresholdPolicyDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateAutoScalingPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
