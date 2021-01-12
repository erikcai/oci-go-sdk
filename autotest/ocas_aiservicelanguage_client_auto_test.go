package autotest

import (
	"github.com/oracle/oci-go-sdk/v32/common"
	"github.com/oracle/oci-go-sdk/v32/ocas"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createOcasAIServiceLanguageClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := ocas.NewAIServiceLanguageClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_ocas_ops_ww_grp@oracle.com" jiraProject="OCAS" opsJiraProject="OCAS"
func TestOcasAIServiceLanguageClientDetectDominantLanguage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocas", "DetectDominantLanguage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DetectDominantLanguage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocas", "AIServiceLanguage", "DetectDominantLanguage", createOcasAIServiceLanguageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocas.AIServiceLanguageClient)

	body, err := testClient.getRequests("ocas", "DetectDominantLanguage")
	assert.NoError(t, err)

	type DetectDominantLanguageRequestInfo struct {
		ContainerId string
		Request     ocas.DetectDominantLanguageRequest
	}

	var requests []DetectDominantLanguageRequestInfo
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
			response, err := c.DetectDominantLanguage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_ocas_ops_ww_grp@oracle.com" jiraProject="OCAS" opsJiraProject="OCAS"
func TestOcasAIServiceLanguageClientDetectLanguageEntities(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocas", "DetectLanguageEntities")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DetectLanguageEntities is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocas", "AIServiceLanguage", "DetectLanguageEntities", createOcasAIServiceLanguageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocas.AIServiceLanguageClient)

	body, err := testClient.getRequests("ocas", "DetectLanguageEntities")
	assert.NoError(t, err)

	type DetectLanguageEntitiesRequestInfo struct {
		ContainerId string
		Request     ocas.DetectLanguageEntitiesRequest
	}

	var requests []DetectLanguageEntitiesRequestInfo
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
			response, err := c.DetectLanguageEntities(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_ocas_ops_ww_grp@oracle.com" jiraProject="OCAS" opsJiraProject="OCAS"
func TestOcasAIServiceLanguageClientDetectLanguageKeyPhrases(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocas", "DetectLanguageKeyPhrases")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DetectLanguageKeyPhrases is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocas", "AIServiceLanguage", "DetectLanguageKeyPhrases", createOcasAIServiceLanguageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocas.AIServiceLanguageClient)

	body, err := testClient.getRequests("ocas", "DetectLanguageKeyPhrases")
	assert.NoError(t, err)

	type DetectLanguageKeyPhrasesRequestInfo struct {
		ContainerId string
		Request     ocas.DetectLanguageKeyPhrasesRequest
	}

	var requests []DetectLanguageKeyPhrasesRequestInfo
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
			response, err := c.DetectLanguageKeyPhrases(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_ocas_ops_ww_grp@oracle.com" jiraProject="OCAS" opsJiraProject="OCAS"
func TestOcasAIServiceLanguageClientDetectLanguageSentiments(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocas", "DetectLanguageSentiments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DetectLanguageSentiments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocas", "AIServiceLanguage", "DetectLanguageSentiments", createOcasAIServiceLanguageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocas.AIServiceLanguageClient)

	body, err := testClient.getRequests("ocas", "DetectLanguageSentiments")
	assert.NoError(t, err)

	type DetectLanguageSentimentsRequestInfo struct {
		ContainerId string
		Request     ocas.DetectLanguageSentimentsRequest
	}

	var requests []DetectLanguageSentimentsRequestInfo
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
			response, err := c.DetectLanguageSentiments(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_ocas_ops_ww_grp@oracle.com" jiraProject="OCAS" opsJiraProject="OCAS"
func TestOcasAIServiceLanguageClientDetectLanguageTopicLabels(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ocas", "DetectLanguageTopicLabels")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DetectLanguageTopicLabels is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ocas", "AIServiceLanguage", "DetectLanguageTopicLabels", createOcasAIServiceLanguageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ocas.AIServiceLanguageClient)

	body, err := testClient.getRequests("ocas", "DetectLanguageTopicLabels")
	assert.NoError(t, err)

	type DetectLanguageTopicLabelsRequestInfo struct {
		ContainerId string
		Request     ocas.DetectLanguageTopicLabelsRequest
	}

	var requests []DetectLanguageTopicLabelsRequestInfo
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
			response, err := c.DetectLanguageTopicLabels(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
