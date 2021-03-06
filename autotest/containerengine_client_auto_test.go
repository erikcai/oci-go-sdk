package autotest

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"github.com/erikcai/oci-go-sdk/v33/containerengine"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createContainerEngineClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := containerengine.NewContainerEngineClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientClusterMigrateToNativeVcn(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "ClusterMigrateToNativeVcn")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ClusterMigrateToNativeVcn is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "ClusterMigrateToNativeVcn", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "ClusterMigrateToNativeVcn")
	assert.NoError(t, err)

	type ClusterMigrateToNativeVcnRequestInfo struct {
		ContainerId string
		Request     containerengine.ClusterMigrateToNativeVcnRequest
	}

	var requests []ClusterMigrateToNativeVcnRequestInfo
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
			response, err := c.ClusterMigrateToNativeVcn(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientCreateCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "CreateCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "CreateCluster", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "CreateCluster")
	assert.NoError(t, err)

	type CreateClusterRequestInfo struct {
		ContainerId string
		Request     containerengine.CreateClusterRequest
	}

	var requests []CreateClusterRequestInfo
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
			response, err := c.CreateCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientCreateKubeconfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "CreateKubeconfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateKubeconfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "CreateKubeconfig", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "CreateKubeconfig")
	assert.NoError(t, err)

	type CreateKubeconfigRequestInfo struct {
		ContainerId string
		Request     containerengine.CreateKubeconfigRequest
	}

	var requests []CreateKubeconfigRequestInfo
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
			response, err := c.CreateKubeconfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientCreateNodePool(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "CreateNodePool")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateNodePool is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "CreateNodePool", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "CreateNodePool")
	assert.NoError(t, err)

	type CreateNodePoolRequestInfo struct {
		ContainerId string
		Request     containerengine.CreateNodePoolRequest
	}

	var requests []CreateNodePoolRequestInfo
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
			response, err := c.CreateNodePool(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientDeleteCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "DeleteCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "DeleteCluster", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "DeleteCluster")
	assert.NoError(t, err)

	type DeleteClusterRequestInfo struct {
		ContainerId string
		Request     containerengine.DeleteClusterRequest
	}

	var requests []DeleteClusterRequestInfo
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
			response, err := c.DeleteCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientDeleteNodePool(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "DeleteNodePool")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteNodePool is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "DeleteNodePool", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "DeleteNodePool")
	assert.NoError(t, err)

	type DeleteNodePoolRequestInfo struct {
		ContainerId string
		Request     containerengine.DeleteNodePoolRequest
	}

	var requests []DeleteNodePoolRequestInfo
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
			response, err := c.DeleteNodePool(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientDeleteWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "DeleteWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "DeleteWorkRequest", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "DeleteWorkRequest")
	assert.NoError(t, err)

	type DeleteWorkRequestRequestInfo struct {
		ContainerId string
		Request     containerengine.DeleteWorkRequestRequest
	}

	var requests []DeleteWorkRequestRequestInfo
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
			response, err := c.DeleteWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientGetCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "GetCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "GetCluster", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "GetCluster")
	assert.NoError(t, err)

	type GetClusterRequestInfo struct {
		ContainerId string
		Request     containerengine.GetClusterRequest
	}

	var requests []GetClusterRequestInfo
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
			response, err := c.GetCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientGetClusterMigrateToNativeVcnStatus(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "GetClusterMigrateToNativeVcnStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetClusterMigrateToNativeVcnStatus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "GetClusterMigrateToNativeVcnStatus", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "GetClusterMigrateToNativeVcnStatus")
	assert.NoError(t, err)

	type GetClusterMigrateToNativeVcnStatusRequestInfo struct {
		ContainerId string
		Request     containerengine.GetClusterMigrateToNativeVcnStatusRequest
	}

	var requests []GetClusterMigrateToNativeVcnStatusRequestInfo
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
			response, err := c.GetClusterMigrateToNativeVcnStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientGetClusterOptions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "GetClusterOptions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetClusterOptions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "GetClusterOptions", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "GetClusterOptions")
	assert.NoError(t, err)

	type GetClusterOptionsRequestInfo struct {
		ContainerId string
		Request     containerengine.GetClusterOptionsRequest
	}

	var requests []GetClusterOptionsRequestInfo
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
			response, err := c.GetClusterOptions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientGetNodePool(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "GetNodePool")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetNodePool is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "GetNodePool", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "GetNodePool")
	assert.NoError(t, err)

	type GetNodePoolRequestInfo struct {
		ContainerId string
		Request     containerengine.GetNodePoolRequest
	}

	var requests []GetNodePoolRequestInfo
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
			response, err := c.GetNodePool(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientGetNodePoolOptions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "GetNodePoolOptions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetNodePoolOptions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "GetNodePoolOptions", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "GetNodePoolOptions")
	assert.NoError(t, err)

	type GetNodePoolOptionsRequestInfo struct {
		ContainerId string
		Request     containerengine.GetNodePoolOptionsRequest
	}

	var requests []GetNodePoolOptionsRequestInfo
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
			response, err := c.GetNodePoolOptions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "GetWorkRequest", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     containerengine.GetWorkRequestRequest
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

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientListClusters(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "ListClusters")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListClusters is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "ListClusters", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "ListClusters")
	assert.NoError(t, err)

	type ListClustersRequestInfo struct {
		ContainerId string
		Request     containerengine.ListClustersRequest
	}

	var requests []ListClustersRequestInfo
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
				r := req.(*containerengine.ListClustersRequest)
				return c.ListClusters(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]containerengine.ListClustersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(containerengine.ListClustersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientListNodePools(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "ListNodePools")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListNodePools is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "ListNodePools", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "ListNodePools")
	assert.NoError(t, err)

	type ListNodePoolsRequestInfo struct {
		ContainerId string
		Request     containerengine.ListNodePoolsRequest
	}

	var requests []ListNodePoolsRequestInfo
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
				r := req.(*containerengine.ListNodePoolsRequest)
				return c.ListNodePools(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]containerengine.ListNodePoolsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(containerengine.ListNodePoolsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "ListWorkRequestErrors", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     containerengine.ListWorkRequestErrorsRequest
	}

	var requests []ListWorkRequestErrorsRequestInfo
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
			response, err := c.ListWorkRequestErrors(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "ListWorkRequestLogs", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     containerengine.ListWorkRequestLogsRequest
	}

	var requests []ListWorkRequestLogsRequestInfo
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
			response, err := c.ListWorkRequestLogs(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "ListWorkRequests", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     containerengine.ListWorkRequestsRequest
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
				r := req.(*containerengine.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]containerengine.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(containerengine.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientUpdateCluster(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "UpdateCluster")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCluster is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "UpdateCluster", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "UpdateCluster")
	assert.NoError(t, err)

	type UpdateClusterRequestInfo struct {
		ContainerId string
		Request     containerengine.UpdateClusterRequest
	}

	var requests []UpdateClusterRequestInfo
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
			response, err := c.UpdateCluster(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientUpdateClusterEndpointConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "UpdateClusterEndpointConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateClusterEndpointConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "UpdateClusterEndpointConfig", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "UpdateClusterEndpointConfig")
	assert.NoError(t, err)

	type UpdateClusterEndpointConfigRequestInfo struct {
		ContainerId string
		Request     containerengine.UpdateClusterEndpointConfigRequest
	}

	var requests []UpdateClusterEndpointConfigRequestInfo
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
			response, err := c.UpdateClusterEndpointConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oke_control_plane_ww_grp@oracle.com" jiraProject="OKE" opsJiraProject="OKE"
func TestContainerEngineClientUpdateNodePool(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerengine", "UpdateNodePool")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateNodePool is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "UpdateNodePool", createContainerEngineClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerengine.ContainerEngineClient)

	body, err := testClient.getRequests("containerengine", "UpdateNodePool")
	assert.NoError(t, err)

	type UpdateNodePoolRequestInfo struct {
		ContainerId string
		Request     containerengine.UpdateNodePoolRequest
	}

	var requests []UpdateNodePoolRequestInfo
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
			response, err := c.UpdateNodePool(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
