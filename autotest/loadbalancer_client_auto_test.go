package autotest

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"github.com/erikcai/oci-go-sdk/v33/loadbalancer"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createLoadBalancerClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientChangeLoadBalancerCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ChangeLoadBalancerCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeLoadBalancerCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ChangeLoadBalancerCompartment", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ChangeLoadBalancerCompartment")
	assert.NoError(t, err)

	type ChangeLoadBalancerCompartmentRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ChangeLoadBalancerCompartmentRequest
	}

	var requests []ChangeLoadBalancerCompartmentRequestInfo
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
			response, err := c.ChangeLoadBalancerCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientCreateBackend(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateBackend")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBackend is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "CreateBackend", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "CreateBackend")
	assert.NoError(t, err)

	type CreateBackendRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateBackendRequest
	}

	var requests []CreateBackendRequestInfo
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
			response, err := c.CreateBackend(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientCreateBackendSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateBackendSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBackendSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "CreateBackendSet", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "CreateBackendSet")
	assert.NoError(t, err)

	type CreateBackendSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateBackendSetRequest
	}

	var requests []CreateBackendSetRequestInfo
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
			response, err := c.CreateBackendSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientCreateCertificate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateCertificate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCertificate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "CreateCertificate", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "CreateCertificate")
	assert.NoError(t, err)

	type CreateCertificateRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateCertificateRequest
	}

	var requests []CreateCertificateRequestInfo
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
			response, err := c.CreateCertificate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientCreateCidrBlocks(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateCidrBlocks")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCidrBlocks is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "CreateCidrBlocks", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "CreateCidrBlocks")
	assert.NoError(t, err)

	type CreateCidrBlocksRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateCidrBlocksRequest
	}

	var requests []CreateCidrBlocksRequestInfo
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
			response, err := c.CreateCidrBlocks(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientCreateHostname(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateHostname")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateHostname is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "CreateHostname", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "CreateHostname")
	assert.NoError(t, err)

	type CreateHostnameRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateHostnameRequest
	}

	var requests []CreateHostnameRequestInfo
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
			response, err := c.CreateHostname(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientCreateListener(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateListener")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateListener is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "CreateListener", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "CreateListener")
	assert.NoError(t, err)

	type CreateListenerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateListenerRequest
	}

	var requests []CreateListenerRequestInfo
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
			response, err := c.CreateListener(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientCreateLoadBalancer(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateLoadBalancer")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateLoadBalancer is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "CreateLoadBalancer", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "CreateLoadBalancer")
	assert.NoError(t, err)

	type CreateLoadBalancerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateLoadBalancerRequest
	}

	var requests []CreateLoadBalancerRequestInfo
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
			response, err := c.CreateLoadBalancer(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientCreatePathRouteSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "CreatePathRouteSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreatePathRouteSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "CreatePathRouteSet", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "CreatePathRouteSet")
	assert.NoError(t, err)

	type CreatePathRouteSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreatePathRouteSetRequest
	}

	var requests []CreatePathRouteSetRequestInfo
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
			response, err := c.CreatePathRouteSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientCreateRuleSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateRuleSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateRuleSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "CreateRuleSet", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "CreateRuleSet")
	assert.NoError(t, err)

	type CreateRuleSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateRuleSetRequest
	}

	var requests []CreateRuleSetRequestInfo
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
			response, err := c.CreateRuleSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientCreateSSLCipherSuite(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateSSLCipherSuite")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSSLCipherSuite is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "CreateSSLCipherSuite", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "CreateSSLCipherSuite")
	assert.NoError(t, err)

	type CreateSSLCipherSuiteRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateSSLCipherSuiteRequest
	}

	var requests []CreateSSLCipherSuiteRequestInfo
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
			response, err := c.CreateSSLCipherSuite(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientDeleteBackend(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteBackend")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBackend is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "DeleteBackend", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "DeleteBackend")
	assert.NoError(t, err)

	type DeleteBackendRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteBackendRequest
	}

	var requests []DeleteBackendRequestInfo
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
			response, err := c.DeleteBackend(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientDeleteBackendSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteBackendSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBackendSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "DeleteBackendSet", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "DeleteBackendSet")
	assert.NoError(t, err)

	type DeleteBackendSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteBackendSetRequest
	}

	var requests []DeleteBackendSetRequestInfo
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
			response, err := c.DeleteBackendSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientDeleteCertificate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteCertificate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCertificate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "DeleteCertificate", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "DeleteCertificate")
	assert.NoError(t, err)

	type DeleteCertificateRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteCertificateRequest
	}

	var requests []DeleteCertificateRequestInfo
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
			response, err := c.DeleteCertificate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientDeleteCidrBlocks(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteCidrBlocks")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCidrBlocks is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "DeleteCidrBlocks", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "DeleteCidrBlocks")
	assert.NoError(t, err)

	type DeleteCidrBlocksRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteCidrBlocksRequest
	}

	var requests []DeleteCidrBlocksRequestInfo
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
			response, err := c.DeleteCidrBlocks(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientDeleteHostname(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteHostname")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteHostname is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "DeleteHostname", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "DeleteHostname")
	assert.NoError(t, err)

	type DeleteHostnameRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteHostnameRequest
	}

	var requests []DeleteHostnameRequestInfo
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
			response, err := c.DeleteHostname(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientDeleteListener(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteListener")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteListener is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "DeleteListener", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "DeleteListener")
	assert.NoError(t, err)

	type DeleteListenerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteListenerRequest
	}

	var requests []DeleteListenerRequestInfo
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
			response, err := c.DeleteListener(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientDeleteLoadBalancer(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteLoadBalancer")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteLoadBalancer is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "DeleteLoadBalancer", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "DeleteLoadBalancer")
	assert.NoError(t, err)

	type DeleteLoadBalancerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteLoadBalancerRequest
	}

	var requests []DeleteLoadBalancerRequestInfo
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
			response, err := c.DeleteLoadBalancer(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientDeletePathRouteSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "DeletePathRouteSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeletePathRouteSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "DeletePathRouteSet", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "DeletePathRouteSet")
	assert.NoError(t, err)

	type DeletePathRouteSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeletePathRouteSetRequest
	}

	var requests []DeletePathRouteSetRequestInfo
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
			response, err := c.DeletePathRouteSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientDeleteRuleSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteRuleSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteRuleSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "DeleteRuleSet", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "DeleteRuleSet")
	assert.NoError(t, err)

	type DeleteRuleSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteRuleSetRequest
	}

	var requests []DeleteRuleSetRequestInfo
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
			response, err := c.DeleteRuleSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientDeleteSSLCipherSuite(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteSSLCipherSuite")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSSLCipherSuite is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "DeleteSSLCipherSuite", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "DeleteSSLCipherSuite")
	assert.NoError(t, err)

	type DeleteSSLCipherSuiteRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteSSLCipherSuiteRequest
	}

	var requests []DeleteSSLCipherSuiteRequestInfo
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
			response, err := c.DeleteSSLCipherSuite(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientGetBackend(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "GetBackend")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBackend is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "GetBackend", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "GetBackend")
	assert.NoError(t, err)

	type GetBackendRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetBackendRequest
	}

	var requests []GetBackendRequestInfo
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
			response, err := c.GetBackend(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientGetBackendHealth(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "GetBackendHealth")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBackendHealth is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "GetBackendHealth", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "GetBackendHealth")
	assert.NoError(t, err)

	type GetBackendHealthRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetBackendHealthRequest
	}

	var requests []GetBackendHealthRequestInfo
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
			response, err := c.GetBackendHealth(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientGetBackendSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "GetBackendSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBackendSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "GetBackendSet", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "GetBackendSet")
	assert.NoError(t, err)

	type GetBackendSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetBackendSetRequest
	}

	var requests []GetBackendSetRequestInfo
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
			response, err := c.GetBackendSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientGetBackendSetHealth(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "GetBackendSetHealth")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBackendSetHealth is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "GetBackendSetHealth", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "GetBackendSetHealth")
	assert.NoError(t, err)

	type GetBackendSetHealthRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetBackendSetHealthRequest
	}

	var requests []GetBackendSetHealthRequestInfo
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
			response, err := c.GetBackendSetHealth(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientGetCidrBlocks(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "GetCidrBlocks")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCidrBlocks is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "GetCidrBlocks", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "GetCidrBlocks")
	assert.NoError(t, err)

	type GetCidrBlocksRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetCidrBlocksRequest
	}

	var requests []GetCidrBlocksRequestInfo
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
			response, err := c.GetCidrBlocks(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientGetHealthChecker(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "GetHealthChecker")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetHealthChecker is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "GetHealthChecker", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "GetHealthChecker")
	assert.NoError(t, err)

	type GetHealthCheckerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetHealthCheckerRequest
	}

	var requests []GetHealthCheckerRequestInfo
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
			response, err := c.GetHealthChecker(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientGetHostname(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "GetHostname")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetHostname is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "GetHostname", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "GetHostname")
	assert.NoError(t, err)

	type GetHostnameRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetHostnameRequest
	}

	var requests []GetHostnameRequestInfo
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
			response, err := c.GetHostname(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientGetLoadBalancer(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "GetLoadBalancer")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLoadBalancer is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "GetLoadBalancer", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "GetLoadBalancer")
	assert.NoError(t, err)

	type GetLoadBalancerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetLoadBalancerRequest
	}

	var requests []GetLoadBalancerRequestInfo
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
			response, err := c.GetLoadBalancer(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientGetLoadBalancerHealth(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "GetLoadBalancerHealth")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLoadBalancerHealth is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "GetLoadBalancerHealth", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "GetLoadBalancerHealth")
	assert.NoError(t, err)

	type GetLoadBalancerHealthRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetLoadBalancerHealthRequest
	}

	var requests []GetLoadBalancerHealthRequestInfo
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
			response, err := c.GetLoadBalancerHealth(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientGetPathRouteSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "GetPathRouteSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPathRouteSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "GetPathRouteSet", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "GetPathRouteSet")
	assert.NoError(t, err)

	type GetPathRouteSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetPathRouteSetRequest
	}

	var requests []GetPathRouteSetRequestInfo
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
			response, err := c.GetPathRouteSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientGetRuleSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "GetRuleSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRuleSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "GetRuleSet", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "GetRuleSet")
	assert.NoError(t, err)

	type GetRuleSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetRuleSetRequest
	}

	var requests []GetRuleSetRequestInfo
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
			response, err := c.GetRuleSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientGetSSLCipherSuite(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "GetSSLCipherSuite")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSSLCipherSuite is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "GetSSLCipherSuite", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "GetSSLCipherSuite")
	assert.NoError(t, err)

	type GetSSLCipherSuiteRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetSSLCipherSuiteRequest
	}

	var requests []GetSSLCipherSuiteRequestInfo
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
			response, err := c.GetSSLCipherSuite(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "GetWorkRequest", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetWorkRequestRequest
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

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListBackendSets(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListBackendSets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBackendSets is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListBackendSets", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListBackendSets")
	assert.NoError(t, err)

	type ListBackendSetsRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListBackendSetsRequest
	}

	var requests []ListBackendSetsRequestInfo
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
			response, err := c.ListBackendSets(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListBackends(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListBackends")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBackends is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListBackends", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListBackends")
	assert.NoError(t, err)

	type ListBackendsRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListBackendsRequest
	}

	var requests []ListBackendsRequestInfo
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
			response, err := c.ListBackends(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListCertificates(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListCertificates")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCertificates is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListCertificates", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListCertificates")
	assert.NoError(t, err)

	type ListCertificatesRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListCertificatesRequest
	}

	var requests []ListCertificatesRequestInfo
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
			response, err := c.ListCertificates(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListCidrBlocks(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListCidrBlocks")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCidrBlocks is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListCidrBlocks", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListCidrBlocks")
	assert.NoError(t, err)

	type ListCidrBlocksRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListCidrBlocksRequest
	}

	var requests []ListCidrBlocksRequestInfo
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
			response, err := c.ListCidrBlocks(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListHostnames(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListHostnames")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListHostnames is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListHostnames", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListHostnames")
	assert.NoError(t, err)

	type ListHostnamesRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListHostnamesRequest
	}

	var requests []ListHostnamesRequestInfo
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
			response, err := c.ListHostnames(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListListenerRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListListenerRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListListenerRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListListenerRules", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListListenerRules")
	assert.NoError(t, err)

	type ListListenerRulesRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListListenerRulesRequest
	}

	var requests []ListListenerRulesRequestInfo
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
			response, err := c.ListListenerRules(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListLoadBalancerHealths(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListLoadBalancerHealths")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLoadBalancerHealths is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListLoadBalancerHealths", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListLoadBalancerHealths")
	assert.NoError(t, err)

	type ListLoadBalancerHealthsRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListLoadBalancerHealthsRequest
	}

	var requests []ListLoadBalancerHealthsRequestInfo
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
				r := req.(*loadbalancer.ListLoadBalancerHealthsRequest)
				return c.ListLoadBalancerHealths(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loadbalancer.ListLoadBalancerHealthsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loadbalancer.ListLoadBalancerHealthsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListLoadBalancers(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListLoadBalancers")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLoadBalancers is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListLoadBalancers", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListLoadBalancers")
	assert.NoError(t, err)

	type ListLoadBalancersRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListLoadBalancersRequest
	}

	var requests []ListLoadBalancersRequestInfo
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
				r := req.(*loadbalancer.ListLoadBalancersRequest)
				return c.ListLoadBalancers(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loadbalancer.ListLoadBalancersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loadbalancer.ListLoadBalancersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListPathRouteSets(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListPathRouteSets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPathRouteSets is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListPathRouteSets", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListPathRouteSets")
	assert.NoError(t, err)

	type ListPathRouteSetsRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListPathRouteSetsRequest
	}

	var requests []ListPathRouteSetsRequestInfo
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
			response, err := c.ListPathRouteSets(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListPolicies(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListPolicies")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPolicies is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListPolicies", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListPolicies")
	assert.NoError(t, err)

	type ListPoliciesRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListPoliciesRequest
	}

	var requests []ListPoliciesRequestInfo
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
				r := req.(*loadbalancer.ListPoliciesRequest)
				return c.ListPolicies(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loadbalancer.ListPoliciesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loadbalancer.ListPoliciesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListProtocols(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListProtocols")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListProtocols is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListProtocols", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListProtocols")
	assert.NoError(t, err)

	type ListProtocolsRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListProtocolsRequest
	}

	var requests []ListProtocolsRequestInfo
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
				r := req.(*loadbalancer.ListProtocolsRequest)
				return c.ListProtocols(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loadbalancer.ListProtocolsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loadbalancer.ListProtocolsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListRuleSets(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListRuleSets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRuleSets is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListRuleSets", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListRuleSets")
	assert.NoError(t, err)

	type ListRuleSetsRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListRuleSetsRequest
	}

	var requests []ListRuleSetsRequestInfo
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
			response, err := c.ListRuleSets(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListSSLCipherSuites(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListSSLCipherSuites")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSSLCipherSuites is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListSSLCipherSuites", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListSSLCipherSuites")
	assert.NoError(t, err)

	type ListSSLCipherSuitesRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListSSLCipherSuitesRequest
	}

	var requests []ListSSLCipherSuitesRequestInfo
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
			response, err := c.ListSSLCipherSuites(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListShapes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListShapes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListShapes", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListShapes")
	assert.NoError(t, err)

	type ListShapesRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListShapesRequest
	}

	var requests []ListShapesRequestInfo
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
				r := req.(*loadbalancer.ListShapesRequest)
				return c.ListShapes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loadbalancer.ListShapesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loadbalancer.ListShapesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "ListWorkRequests", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListWorkRequestsRequest
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
				r := req.(*loadbalancer.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loadbalancer.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loadbalancer.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientUpdateBackend(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateBackend")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBackend is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "UpdateBackend", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "UpdateBackend")
	assert.NoError(t, err)

	type UpdateBackendRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateBackendRequest
	}

	var requests []UpdateBackendRequestInfo
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
			response, err := c.UpdateBackend(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientUpdateBackendSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateBackendSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBackendSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "UpdateBackendSet", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "UpdateBackendSet")
	assert.NoError(t, err)

	type UpdateBackendSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateBackendSetRequest
	}

	var requests []UpdateBackendSetRequestInfo
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
			response, err := c.UpdateBackendSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientUpdateCidrBlocks(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateCidrBlocks")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCidrBlocks is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "UpdateCidrBlocks", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "UpdateCidrBlocks")
	assert.NoError(t, err)

	type UpdateCidrBlocksRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateCidrBlocksRequest
	}

	var requests []UpdateCidrBlocksRequestInfo
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
			response, err := c.UpdateCidrBlocks(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientUpdateHealthChecker(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateHealthChecker")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateHealthChecker is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "UpdateHealthChecker", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "UpdateHealthChecker")
	assert.NoError(t, err)

	type UpdateHealthCheckerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateHealthCheckerRequest
	}

	var requests []UpdateHealthCheckerRequestInfo
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
			response, err := c.UpdateHealthChecker(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientUpdateHostname(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateHostname")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateHostname is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "UpdateHostname", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "UpdateHostname")
	assert.NoError(t, err)

	type UpdateHostnameRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateHostnameRequest
	}

	var requests []UpdateHostnameRequestInfo
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
			response, err := c.UpdateHostname(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientUpdateListener(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateListener")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateListener is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "UpdateListener", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "UpdateListener")
	assert.NoError(t, err)

	type UpdateListenerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateListenerRequest
	}

	var requests []UpdateListenerRequestInfo
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
			response, err := c.UpdateListener(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientUpdateLoadBalancer(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateLoadBalancer")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateLoadBalancer is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "UpdateLoadBalancer", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "UpdateLoadBalancer")
	assert.NoError(t, err)

	type UpdateLoadBalancerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateLoadBalancerRequest
	}

	var requests []UpdateLoadBalancerRequestInfo
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
			response, err := c.UpdateLoadBalancer(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientUpdateLoadBalancerShape(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateLoadBalancerShape")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateLoadBalancerShape is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "UpdateLoadBalancerShape", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "UpdateLoadBalancerShape")
	assert.NoError(t, err)

	type UpdateLoadBalancerShapeRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateLoadBalancerShapeRequest
	}

	var requests []UpdateLoadBalancerShapeRequestInfo
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
			response, err := c.UpdateLoadBalancerShape(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientUpdateNetworkSecurityGroups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateNetworkSecurityGroups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateNetworkSecurityGroups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "UpdateNetworkSecurityGroups", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "UpdateNetworkSecurityGroups")
	assert.NoError(t, err)

	type UpdateNetworkSecurityGroupsRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateNetworkSecurityGroupsRequest
	}

	var requests []UpdateNetworkSecurityGroupsRequestInfo
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
			response, err := c.UpdateNetworkSecurityGroups(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientUpdatePathRouteSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdatePathRouteSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdatePathRouteSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "UpdatePathRouteSet", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "UpdatePathRouteSet")
	assert.NoError(t, err)

	type UpdatePathRouteSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdatePathRouteSetRequest
	}

	var requests []UpdatePathRouteSetRequestInfo
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
			response, err := c.UpdatePathRouteSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientUpdateRuleSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateRuleSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRuleSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "UpdateRuleSet", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "UpdateRuleSet")
	assert.NoError(t, err)

	type UpdateRuleSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateRuleSetRequest
	}

	var requests []UpdateRuleSetRequestInfo
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
			response, err := c.UpdateRuleSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_lbaas_dev_us_grp@oracle.com" jiraProject="LBCP" opsJiraProject="LBCP"
func TestLoadBalancerClientUpdateSSLCipherSuite(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateSSLCipherSuite")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSSLCipherSuite is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("loadbalancer", "LoadBalancer", "UpdateSSLCipherSuite", createLoadBalancerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(loadbalancer.LoadBalancerClient)

	body, err := testClient.getRequests("loadbalancer", "UpdateSSLCipherSuite")
	assert.NoError(t, err)

	type UpdateSSLCipherSuiteRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateSSLCipherSuiteRequest
	}

	var requests []UpdateSSLCipherSuiteRequestInfo
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
			response, err := c.UpdateSSLCipherSuite(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
