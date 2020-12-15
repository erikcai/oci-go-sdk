package autotest

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"github.com/oracle/oci-go-sdk/v31/core"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createCoreArtifactsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := core.NewArtifactsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="ocir" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestCoreArtifactsClientChangeContainerRepositoryCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeContainerRepositoryCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeContainerRepositoryCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Artifacts", "ChangeContainerRepositoryCompartment", createCoreArtifactsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ArtifactsClient)

	body, err := testClient.getRequests("core", "ChangeContainerRepositoryCompartment")
	assert.NoError(t, err)

	type ChangeContainerRepositoryCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeContainerRepositoryCompartmentRequest
	}

	var requests []ChangeContainerRepositoryCompartmentRequestInfo
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
			response, err := c.ChangeContainerRepositoryCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ocir" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestCoreArtifactsClientCreateContainerRepository(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateContainerRepository")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateContainerRepository is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Artifacts", "CreateContainerRepository", createCoreArtifactsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ArtifactsClient)

	body, err := testClient.getRequests("core", "CreateContainerRepository")
	assert.NoError(t, err)

	type CreateContainerRepositoryRequestInfo struct {
		ContainerId string
		Request     core.CreateContainerRepositoryRequest
	}

	var requests []CreateContainerRepositoryRequestInfo
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
			response, err := c.CreateContainerRepository(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ocir" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestCoreArtifactsClientDeleteContainerImage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteContainerImage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteContainerImage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Artifacts", "DeleteContainerImage", createCoreArtifactsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ArtifactsClient)

	body, err := testClient.getRequests("core", "DeleteContainerImage")
	assert.NoError(t, err)

	type DeleteContainerImageRequestInfo struct {
		ContainerId string
		Request     core.DeleteContainerImageRequest
	}

	var requests []DeleteContainerImageRequestInfo
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
			response, err := c.DeleteContainerImage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ocir" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestCoreArtifactsClientDeleteContainerRepository(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteContainerRepository")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteContainerRepository is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Artifacts", "DeleteContainerRepository", createCoreArtifactsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ArtifactsClient)

	body, err := testClient.getRequests("core", "DeleteContainerRepository")
	assert.NoError(t, err)

	type DeleteContainerRepositoryRequestInfo struct {
		ContainerId string
		Request     core.DeleteContainerRepositoryRequest
	}

	var requests []DeleteContainerRepositoryRequestInfo
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
			response, err := c.DeleteContainerRepository(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ocir" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestCoreArtifactsClientGetContainerConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetContainerConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetContainerConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Artifacts", "GetContainerConfiguration", createCoreArtifactsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ArtifactsClient)

	body, err := testClient.getRequests("core", "GetContainerConfiguration")
	assert.NoError(t, err)

	type GetContainerConfigurationRequestInfo struct {
		ContainerId string
		Request     core.GetContainerConfigurationRequest
	}

	var requests []GetContainerConfigurationRequestInfo
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
			response, err := c.GetContainerConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ocir" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestCoreArtifactsClientGetContainerImage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetContainerImage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetContainerImage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Artifacts", "GetContainerImage", createCoreArtifactsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ArtifactsClient)

	body, err := testClient.getRequests("core", "GetContainerImage")
	assert.NoError(t, err)

	type GetContainerImageRequestInfo struct {
		ContainerId string
		Request     core.GetContainerImageRequest
	}

	var requests []GetContainerImageRequestInfo
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
			response, err := c.GetContainerImage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ocir" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestCoreArtifactsClientGetContainerRepository(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetContainerRepository")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetContainerRepository is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Artifacts", "GetContainerRepository", createCoreArtifactsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ArtifactsClient)

	body, err := testClient.getRequests("core", "GetContainerRepository")
	assert.NoError(t, err)

	type GetContainerRepositoryRequestInfo struct {
		ContainerId string
		Request     core.GetContainerRepositoryRequest
	}

	var requests []GetContainerRepositoryRequestInfo
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
			response, err := c.GetContainerRepository(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ocir" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestCoreArtifactsClientListContainerImages(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListContainerImages")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListContainerImages is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Artifacts", "ListContainerImages", createCoreArtifactsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ArtifactsClient)

	body, err := testClient.getRequests("core", "ListContainerImages")
	assert.NoError(t, err)

	type ListContainerImagesRequestInfo struct {
		ContainerId string
		Request     core.ListContainerImagesRequest
	}

	var requests []ListContainerImagesRequestInfo
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
				r := req.(*core.ListContainerImagesRequest)
				return c.ListContainerImages(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListContainerImagesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListContainerImagesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ocir" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestCoreArtifactsClientListContainerRepositories(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListContainerRepositories")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListContainerRepositories is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Artifacts", "ListContainerRepositories", createCoreArtifactsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ArtifactsClient)

	body, err := testClient.getRequests("core", "ListContainerRepositories")
	assert.NoError(t, err)

	type ListContainerRepositoriesRequestInfo struct {
		ContainerId string
		Request     core.ListContainerRepositoriesRequest
	}

	var requests []ListContainerRepositoriesRequestInfo
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
				r := req.(*core.ListContainerRepositoriesRequest)
				return c.ListContainerRepositories(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListContainerRepositoriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListContainerRepositoriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestCoreArtifactsClientRemoveContainerVersion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "RemoveContainerVersion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RemoveContainerVersion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Artifacts", "RemoveContainerVersion", createCoreArtifactsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ArtifactsClient)

	body, err := testClient.getRequests("core", "RemoveContainerVersion")
	assert.NoError(t, err)

	type RemoveContainerVersionRequestInfo struct {
		ContainerId string
		Request     core.RemoveContainerVersionRequest
	}

	var requests []RemoveContainerVersionRequestInfo
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
			response, err := c.RemoveContainerVersion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ocir" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestCoreArtifactsClientRestoreContainerImage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "RestoreContainerImage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestoreContainerImage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Artifacts", "RestoreContainerImage", createCoreArtifactsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ArtifactsClient)

	body, err := testClient.getRequests("core", "RestoreContainerImage")
	assert.NoError(t, err)

	type RestoreContainerImageRequestInfo struct {
		ContainerId string
		Request     core.RestoreContainerImageRequest
	}

	var requests []RestoreContainerImageRequestInfo
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
			response, err := c.RestoreContainerImage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ocir" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestCoreArtifactsClientUpdateContainerConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateContainerConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateContainerConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Artifacts", "UpdateContainerConfiguration", createCoreArtifactsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ArtifactsClient)

	body, err := testClient.getRequests("core", "UpdateContainerConfiguration")
	assert.NoError(t, err)

	type UpdateContainerConfigurationRequestInfo struct {
		ContainerId string
		Request     core.UpdateContainerConfigurationRequest
	}

	var requests []UpdateContainerConfigurationRequestInfo
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
			response, err := c.UpdateContainerConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="ocir" email="odx_registry_grp@oracle.com" jiraProject="OCIR" opsJiraProject="OCIR"
func TestCoreArtifactsClientUpdateContainerRepository(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateContainerRepository")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateContainerRepository is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Artifacts", "UpdateContainerRepository", createCoreArtifactsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ArtifactsClient)

	body, err := testClient.getRequests("core", "UpdateContainerRepository")
	assert.NoError(t, err)

	type UpdateContainerRepositoryRequestInfo struct {
		ContainerId string
		Request     core.UpdateContainerRepositoryRequest
	}

	var requests []UpdateContainerRepositoryRequestInfo
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
			response, err := c.UpdateContainerRepository(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
