package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/containerregistry"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createContainerRegistryClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := containerregistry.NewContainerRegistryClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestContainerRegistryClientChangeDockerRepositoryCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerregistry", "ChangeDockerRepositoryCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeDockerRepositoryCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerregistry", "ContainerRegistry", "ChangeDockerRepositoryCompartment", createContainerRegistryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerregistry.ContainerRegistryClient)

	body, err := testClient.getRequests("containerregistry", "ChangeDockerRepositoryCompartment")
	assert.NoError(t, err)

	type ChangeDockerRepositoryCompartmentRequestInfo struct {
		ContainerId string
		Request     containerregistry.ChangeDockerRepositoryCompartmentRequest
	}

	var requests []ChangeDockerRepositoryCompartmentRequestInfo
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
			response, err := c.ChangeDockerRepositoryCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestContainerRegistryClientCreateDockerRepository(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerregistry", "CreateDockerRepository")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDockerRepository is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerregistry", "ContainerRegistry", "CreateDockerRepository", createContainerRegistryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerregistry.ContainerRegistryClient)

	body, err := testClient.getRequests("containerregistry", "CreateDockerRepository")
	assert.NoError(t, err)

	type CreateDockerRepositoryRequestInfo struct {
		ContainerId string
		Request     containerregistry.CreateDockerRepositoryRequest
	}

	var requests []CreateDockerRepositoryRequestInfo
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
			response, err := c.CreateDockerRepository(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestContainerRegistryClientDeleteDockerRepository(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerregistry", "DeleteDockerRepository")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDockerRepository is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerregistry", "ContainerRegistry", "DeleteDockerRepository", createContainerRegistryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerregistry.ContainerRegistryClient)

	body, err := testClient.getRequests("containerregistry", "DeleteDockerRepository")
	assert.NoError(t, err)

	type DeleteDockerRepositoryRequestInfo struct {
		ContainerId string
		Request     containerregistry.DeleteDockerRepositoryRequest
	}

	var requests []DeleteDockerRepositoryRequestInfo
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
			response, err := c.DeleteDockerRepository(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestContainerRegistryClientDeleteDockerRepositoryContents(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerregistry", "DeleteDockerRepositoryContents")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDockerRepositoryContents is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerregistry", "ContainerRegistry", "DeleteDockerRepositoryContents", createContainerRegistryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerregistry.ContainerRegistryClient)

	body, err := testClient.getRequests("containerregistry", "DeleteDockerRepositoryContents")
	assert.NoError(t, err)

	type DeleteDockerRepositoryContentsRequestInfo struct {
		ContainerId string
		Request     containerregistry.DeleteDockerRepositoryContentsRequest
	}

	var requests []DeleteDockerRepositoryContentsRequestInfo
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
			response, err := c.DeleteDockerRepositoryContents(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestContainerRegistryClientDeleteDockerTag(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerregistry", "DeleteDockerTag")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDockerTag is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerregistry", "ContainerRegistry", "DeleteDockerTag", createContainerRegistryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerregistry.ContainerRegistryClient)

	body, err := testClient.getRequests("containerregistry", "DeleteDockerTag")
	assert.NoError(t, err)

	type DeleteDockerTagRequestInfo struct {
		ContainerId string
		Request     containerregistry.DeleteDockerTagRequest
	}

	var requests []DeleteDockerTagRequestInfo
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
			response, err := c.DeleteDockerTag(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestContainerRegistryClientGetDockerImageMetadata(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerregistry", "GetDockerImageMetadata")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDockerImageMetadata is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerregistry", "ContainerRegistry", "GetDockerImageMetadata", createContainerRegistryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerregistry.ContainerRegistryClient)

	body, err := testClient.getRequests("containerregistry", "GetDockerImageMetadata")
	assert.NoError(t, err)

	type GetDockerImageMetadataRequestInfo struct {
		ContainerId string
		Request     containerregistry.GetDockerImageMetadataRequest
	}

	var requests []GetDockerImageMetadataRequestInfo
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
			response, err := c.GetDockerImageMetadata(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestContainerRegistryClientGetDockerRepository(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerregistry", "GetDockerRepository")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDockerRepository is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerregistry", "ContainerRegistry", "GetDockerRepository", createContainerRegistryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerregistry.ContainerRegistryClient)

	body, err := testClient.getRequests("containerregistry", "GetDockerRepository")
	assert.NoError(t, err)

	type GetDockerRepositoryRequestInfo struct {
		ContainerId string
		Request     containerregistry.GetDockerRepositoryRequest
	}

	var requests []GetDockerRepositoryRequestInfo
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
			response, err := c.GetDockerRepository(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestContainerRegistryClientGetDockerToken(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerregistry", "GetDockerToken")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDockerToken is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerregistry", "ContainerRegistry", "GetDockerToken", createContainerRegistryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerregistry.ContainerRegistryClient)

	body, err := testClient.getRequests("containerregistry", "GetDockerToken")
	assert.NoError(t, err)

	type GetDockerTokenRequestInfo struct {
		ContainerId string
		Request     containerregistry.GetDockerTokenRequest
	}

	var requests []GetDockerTokenRequestInfo
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
			response, err := c.GetDockerToken(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestContainerRegistryClientGetDockerTokenWithOAuth(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerregistry", "GetDockerTokenWithOAuth")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDockerTokenWithOAuth is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerregistry", "ContainerRegistry", "GetDockerTokenWithOAuth", createContainerRegistryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerregistry.ContainerRegistryClient)

	body, err := testClient.getRequests("containerregistry", "GetDockerTokenWithOAuth")
	assert.NoError(t, err)

	type GetDockerTokenWithOAuthRequestInfo struct {
		ContainerId string
		Request     containerregistry.GetDockerTokenWithOAuthRequest
	}

	var requests []GetDockerTokenWithOAuthRequestInfo
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
			response, err := c.GetDockerTokenWithOAuth(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestContainerRegistryClientGetNamespace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerregistry", "GetNamespace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetNamespace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerregistry", "ContainerRegistry", "GetNamespace", createContainerRegistryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerregistry.ContainerRegistryClient)

	body, err := testClient.getRequests("containerregistry", "GetNamespace")
	assert.NoError(t, err)

	type GetNamespaceRequestInfo struct {
		ContainerId string
		Request     containerregistry.GetNamespaceRequest
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

// IssueRoutingInfo tag="default" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestContainerRegistryClientListDockerRepositories(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerregistry", "ListDockerRepositories")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDockerRepositories is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerregistry", "ContainerRegistry", "ListDockerRepositories", createContainerRegistryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerregistry.ContainerRegistryClient)

	body, err := testClient.getRequests("containerregistry", "ListDockerRepositories")
	assert.NoError(t, err)

	type ListDockerRepositoriesRequestInfo struct {
		ContainerId string
		Request     containerregistry.ListDockerRepositoriesRequest
	}

	var requests []ListDockerRepositoriesRequestInfo
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
				r := req.(*containerregistry.ListDockerRepositoriesRequest)
				return c.ListDockerRepositories(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]containerregistry.ListDockerRepositoriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(containerregistry.ListDockerRepositoriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestContainerRegistryClientListDockerTags(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerregistry", "ListDockerTags")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDockerTags is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerregistry", "ContainerRegistry", "ListDockerTags", createContainerRegistryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerregistry.ContainerRegistryClient)

	body, err := testClient.getRequests("containerregistry", "ListDockerTags")
	assert.NoError(t, err)

	type ListDockerTagsRequestInfo struct {
		ContainerId string
		Request     containerregistry.ListDockerTagsRequest
	}

	var requests []ListDockerTagsRequestInfo
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
				r := req.(*containerregistry.ListDockerTagsRequest)
				return c.ListDockerTags(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]containerregistry.ListDockerTagsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(containerregistry.ListDockerTagsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestContainerRegistryClientUndeleteManifest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerregistry", "UndeleteManifest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UndeleteManifest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerregistry", "ContainerRegistry", "UndeleteManifest", createContainerRegistryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerregistry.ContainerRegistryClient)

	body, err := testClient.getRequests("containerregistry", "UndeleteManifest")
	assert.NoError(t, err)

	type UndeleteManifestRequestInfo struct {
		ContainerId string
		Request     containerregistry.UndeleteManifestRequest
	}

	var requests []UndeleteManifestRequestInfo
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
			response, err := c.UndeleteManifest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestContainerRegistryClientUpdateDockerRepository(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("containerregistry", "UpdateDockerRepository")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDockerRepository is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("containerregistry", "ContainerRegistry", "UpdateDockerRepository", createContainerRegistryClientWithProvider)
	assert.NoError(t, err)
	c := cc.(containerregistry.ContainerRegistryClient)

	body, err := testClient.getRequests("containerregistry", "UpdateDockerRepository")
	assert.NoError(t, err)

	type UpdateDockerRepositoryRequestInfo struct {
		ContainerId string
		Request     containerregistry.UpdateDockerRepositoryRequest
	}

	var requests []UpdateDockerRepositoryRequestInfo
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
			response, err := c.UpdateDockerRepository(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
