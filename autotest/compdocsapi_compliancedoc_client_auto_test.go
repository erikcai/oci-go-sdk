package autotest

import (
	"github.com/oracle/oci-go-sdk/v26/common"
	"github.com/oracle/oci-go-sdk/v26/compdocsapi"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createComplianceDocClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := compdocsapi.NewComplianceDocClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="compliancedocs_dev_team_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/COMPDOCS" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/COMPDOCS"
func TestComplianceDocClientCreateNonDisclosureAgreement(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("compdocsapi", "CreateNonDisclosureAgreement")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateNonDisclosureAgreement is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("compdocsapi", "ComplianceDoc", "CreateNonDisclosureAgreement", createComplianceDocClientWithProvider)
	assert.NoError(t, err)
	c := cc.(compdocsapi.ComplianceDocClient)

	body, err := testClient.getRequests("compdocsapi", "CreateNonDisclosureAgreement")
	assert.NoError(t, err)

	type CreateNonDisclosureAgreementRequestInfo struct {
		ContainerId string
		Request     compdocsapi.CreateNonDisclosureAgreementRequest
	}

	var requests []CreateNonDisclosureAgreementRequestInfo
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
			response, err := c.CreateNonDisclosureAgreement(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="compliancedocs_dev_team_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/COMPDOCS" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/COMPDOCS"
func TestComplianceDocClientGetComplianceDocument(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("compdocsapi", "GetComplianceDocument")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetComplianceDocument is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("compdocsapi", "ComplianceDoc", "GetComplianceDocument", createComplianceDocClientWithProvider)
	assert.NoError(t, err)
	c := cc.(compdocsapi.ComplianceDocClient)

	body, err := testClient.getRequests("compdocsapi", "GetComplianceDocument")
	assert.NoError(t, err)

	type GetComplianceDocumentRequestInfo struct {
		ContainerId string
		Request     compdocsapi.GetComplianceDocumentRequest
	}

	var requests []GetComplianceDocumentRequestInfo
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
			response, err := c.GetComplianceDocument(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="compliancedocs_dev_team_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/COMPDOCS" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/COMPDOCS"
func TestComplianceDocClientGetComplianceDocumentContent(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("compdocsapi", "GetComplianceDocumentContent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetComplianceDocumentContent is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("compdocsapi", "ComplianceDoc", "GetComplianceDocumentContent", createComplianceDocClientWithProvider)
	assert.NoError(t, err)
	c := cc.(compdocsapi.ComplianceDocClient)

	body, err := testClient.getRequests("compdocsapi", "GetComplianceDocumentContent")
	assert.NoError(t, err)

	type GetComplianceDocumentContentRequestInfo struct {
		ContainerId string
		Request     compdocsapi.GetComplianceDocumentContentRequest
	}

	var requests []GetComplianceDocumentContentRequestInfo
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
			response, err := c.GetComplianceDocumentContent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="compliancedocs_dev_team_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/COMPDOCS" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/COMPDOCS"
func TestComplianceDocClientListComplianceDocuments(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("compdocsapi", "ListComplianceDocuments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListComplianceDocuments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("compdocsapi", "ComplianceDoc", "ListComplianceDocuments", createComplianceDocClientWithProvider)
	assert.NoError(t, err)
	c := cc.(compdocsapi.ComplianceDocClient)

	body, err := testClient.getRequests("compdocsapi", "ListComplianceDocuments")
	assert.NoError(t, err)

	type ListComplianceDocumentsRequestInfo struct {
		ContainerId string
		Request     compdocsapi.ListComplianceDocumentsRequest
	}

	var requests []ListComplianceDocumentsRequestInfo
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
				r := req.(*compdocsapi.ListComplianceDocumentsRequest)
				return c.ListComplianceDocuments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]compdocsapi.ListComplianceDocumentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(compdocsapi.ListComplianceDocumentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
