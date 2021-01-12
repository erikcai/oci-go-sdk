package autotest

import (
	"github.com/oracle/oci-go-sdk/v32/common"
	"github.com/oracle/oci-go-sdk/v32/dns"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createDnsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := dns.NewDnsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientChangeResolverCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "ChangeResolverCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeResolverCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "ChangeResolverCompartment", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "ChangeResolverCompartment")
	assert.NoError(t, err)

	type ChangeResolverCompartmentRequestInfo struct {
		ContainerId string
		Request     dns.ChangeResolverCompartmentRequest
	}

	var requests []ChangeResolverCompartmentRequestInfo
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
			response, err := c.ChangeResolverCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientChangeSteeringPolicyCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "ChangeSteeringPolicyCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeSteeringPolicyCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "ChangeSteeringPolicyCompartment", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "ChangeSteeringPolicyCompartment")
	assert.NoError(t, err)

	type ChangeSteeringPolicyCompartmentRequestInfo struct {
		ContainerId string
		Request     dns.ChangeSteeringPolicyCompartmentRequest
	}

	var requests []ChangeSteeringPolicyCompartmentRequestInfo
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
			response, err := c.ChangeSteeringPolicyCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientChangeTsigKeyCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "ChangeTsigKeyCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeTsigKeyCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "ChangeTsigKeyCompartment", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "ChangeTsigKeyCompartment")
	assert.NoError(t, err)

	type ChangeTsigKeyCompartmentRequestInfo struct {
		ContainerId string
		Request     dns.ChangeTsigKeyCompartmentRequest
	}

	var requests []ChangeTsigKeyCompartmentRequestInfo
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
			response, err := c.ChangeTsigKeyCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientChangeViewCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "ChangeViewCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeViewCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "ChangeViewCompartment", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "ChangeViewCompartment")
	assert.NoError(t, err)

	type ChangeViewCompartmentRequestInfo struct {
		ContainerId string
		Request     dns.ChangeViewCompartmentRequest
	}

	var requests []ChangeViewCompartmentRequestInfo
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
			response, err := c.ChangeViewCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientChangeZoneCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "ChangeZoneCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeZoneCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "ChangeZoneCompartment", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "ChangeZoneCompartment")
	assert.NoError(t, err)

	type ChangeZoneCompartmentRequestInfo struct {
		ContainerId string
		Request     dns.ChangeZoneCompartmentRequest
	}

	var requests []ChangeZoneCompartmentRequestInfo
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
			response, err := c.ChangeZoneCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientCreateResolverEndpoint(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "CreateResolverEndpoint")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateResolverEndpoint is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "CreateResolverEndpoint", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "CreateResolverEndpoint")
	assert.NoError(t, err)

	type CreateResolverEndpointRequestInfo struct {
		ContainerId string
		Request     dns.CreateResolverEndpointRequest
	}

	var requests []CreateResolverEndpointRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateResolverEndpointRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateResolverEndpointDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "endpointType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"VNIC": &dns.CreateResolverVnicEndpointDetails{},
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
			response, err := c.CreateResolverEndpoint(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientCreateSteeringPolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "CreateSteeringPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSteeringPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "CreateSteeringPolicy", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "CreateSteeringPolicy")
	assert.NoError(t, err)

	type CreateSteeringPolicyRequestInfo struct {
		ContainerId string
		Request     dns.CreateSteeringPolicyRequest
	}

	var requests []CreateSteeringPolicyRequestInfo
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
			response, err := c.CreateSteeringPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientCreateSteeringPolicyAttachment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "CreateSteeringPolicyAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSteeringPolicyAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "CreateSteeringPolicyAttachment", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "CreateSteeringPolicyAttachment")
	assert.NoError(t, err)

	type CreateSteeringPolicyAttachmentRequestInfo struct {
		ContainerId string
		Request     dns.CreateSteeringPolicyAttachmentRequest
	}

	var requests []CreateSteeringPolicyAttachmentRequestInfo
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
			response, err := c.CreateSteeringPolicyAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientCreateTsigKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "CreateTsigKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTsigKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "CreateTsigKey", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "CreateTsigKey")
	assert.NoError(t, err)

	type CreateTsigKeyRequestInfo struct {
		ContainerId string
		Request     dns.CreateTsigKeyRequest
	}

	var requests []CreateTsigKeyRequestInfo
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
			response, err := c.CreateTsigKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientCreateView(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "CreateView")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateView is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "CreateView", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "CreateView")
	assert.NoError(t, err)

	type CreateViewRequestInfo struct {
		ContainerId string
		Request     dns.CreateViewRequest
	}

	var requests []CreateViewRequestInfo
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
			response, err := c.CreateView(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientCreateZone(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "CreateZone")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateZone is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "CreateZone", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "CreateZone")
	assert.NoError(t, err)

	type CreateZoneRequestInfo struct {
		ContainerId string
		Request     dns.CreateZoneRequest
	}

	var requests []CreateZoneRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateZoneRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateZoneBaseDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "migrationSource",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"NONE":   &dns.CreateZoneDetails{},
				"DYNECT": &dns.CreateMigratedDynectZoneDetails{},
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
			response, err := c.CreateZone(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientDeleteDomainRecords(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "DeleteDomainRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDomainRecords is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "DeleteDomainRecords", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "DeleteDomainRecords")
	assert.NoError(t, err)

	type DeleteDomainRecordsRequestInfo struct {
		ContainerId string
		Request     dns.DeleteDomainRecordsRequest
	}

	var requests []DeleteDomainRecordsRequestInfo
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
			response, err := c.DeleteDomainRecords(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientDeleteRRSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "DeleteRRSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteRRSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "DeleteRRSet", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "DeleteRRSet")
	assert.NoError(t, err)

	type DeleteRRSetRequestInfo struct {
		ContainerId string
		Request     dns.DeleteRRSetRequest
	}

	var requests []DeleteRRSetRequestInfo
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
			response, err := c.DeleteRRSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientDeleteResolverEndpoint(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "DeleteResolverEndpoint")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteResolverEndpoint is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "DeleteResolverEndpoint", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "DeleteResolverEndpoint")
	assert.NoError(t, err)

	type DeleteResolverEndpointRequestInfo struct {
		ContainerId string
		Request     dns.DeleteResolverEndpointRequest
	}

	var requests []DeleteResolverEndpointRequestInfo
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
			response, err := c.DeleteResolverEndpoint(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientDeleteSteeringPolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "DeleteSteeringPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSteeringPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "DeleteSteeringPolicy", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "DeleteSteeringPolicy")
	assert.NoError(t, err)

	type DeleteSteeringPolicyRequestInfo struct {
		ContainerId string
		Request     dns.DeleteSteeringPolicyRequest
	}

	var requests []DeleteSteeringPolicyRequestInfo
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
			response, err := c.DeleteSteeringPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientDeleteSteeringPolicyAttachment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "DeleteSteeringPolicyAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSteeringPolicyAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "DeleteSteeringPolicyAttachment", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "DeleteSteeringPolicyAttachment")
	assert.NoError(t, err)

	type DeleteSteeringPolicyAttachmentRequestInfo struct {
		ContainerId string
		Request     dns.DeleteSteeringPolicyAttachmentRequest
	}

	var requests []DeleteSteeringPolicyAttachmentRequestInfo
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
			response, err := c.DeleteSteeringPolicyAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientDeleteTsigKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "DeleteTsigKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteTsigKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "DeleteTsigKey", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "DeleteTsigKey")
	assert.NoError(t, err)

	type DeleteTsigKeyRequestInfo struct {
		ContainerId string
		Request     dns.DeleteTsigKeyRequest
	}

	var requests []DeleteTsigKeyRequestInfo
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
			response, err := c.DeleteTsigKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientDeleteView(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "DeleteView")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteView is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "DeleteView", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "DeleteView")
	assert.NoError(t, err)

	type DeleteViewRequestInfo struct {
		ContainerId string
		Request     dns.DeleteViewRequest
	}

	var requests []DeleteViewRequestInfo
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
			response, err := c.DeleteView(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientDeleteZone(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "DeleteZone")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteZone is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "DeleteZone", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "DeleteZone")
	assert.NoError(t, err)

	type DeleteZoneRequestInfo struct {
		ContainerId string
		Request     dns.DeleteZoneRequest
	}

	var requests []DeleteZoneRequestInfo
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
			response, err := c.DeleteZone(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientGetDomainRecords(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "GetDomainRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDomainRecords is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "GetDomainRecords", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "GetDomainRecords")
	assert.NoError(t, err)

	type GetDomainRecordsRequestInfo struct {
		ContainerId string
		Request     dns.GetDomainRecordsRequest
	}

	var requests []GetDomainRecordsRequestInfo
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
				r := req.(*dns.GetDomainRecordsRequest)
				return c.GetDomainRecords(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.GetDomainRecordsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.GetDomainRecordsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientGetRRSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "GetRRSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRRSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "GetRRSet", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "GetRRSet")
	assert.NoError(t, err)

	type GetRRSetRequestInfo struct {
		ContainerId string
		Request     dns.GetRRSetRequest
	}

	var requests []GetRRSetRequestInfo
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
				r := req.(*dns.GetRRSetRequest)
				return c.GetRRSet(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.GetRRSetResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.GetRRSetResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientGetResolver(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "GetResolver")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetResolver is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "GetResolver", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "GetResolver")
	assert.NoError(t, err)

	type GetResolverRequestInfo struct {
		ContainerId string
		Request     dns.GetResolverRequest
	}

	var requests []GetResolverRequestInfo
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
			response, err := c.GetResolver(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientGetResolverEndpoint(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "GetResolverEndpoint")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetResolverEndpoint is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "GetResolverEndpoint", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "GetResolverEndpoint")
	assert.NoError(t, err)

	type GetResolverEndpointRequestInfo struct {
		ContainerId string
		Request     dns.GetResolverEndpointRequest
	}

	var requests []GetResolverEndpointRequestInfo
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
			response, err := c.GetResolverEndpoint(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientGetSteeringPolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "GetSteeringPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSteeringPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "GetSteeringPolicy", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "GetSteeringPolicy")
	assert.NoError(t, err)

	type GetSteeringPolicyRequestInfo struct {
		ContainerId string
		Request     dns.GetSteeringPolicyRequest
	}

	var requests []GetSteeringPolicyRequestInfo
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
			response, err := c.GetSteeringPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientGetSteeringPolicyAttachment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "GetSteeringPolicyAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSteeringPolicyAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "GetSteeringPolicyAttachment", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "GetSteeringPolicyAttachment")
	assert.NoError(t, err)

	type GetSteeringPolicyAttachmentRequestInfo struct {
		ContainerId string
		Request     dns.GetSteeringPolicyAttachmentRequest
	}

	var requests []GetSteeringPolicyAttachmentRequestInfo
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
			response, err := c.GetSteeringPolicyAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientGetTsigKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "GetTsigKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTsigKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "GetTsigKey", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "GetTsigKey")
	assert.NoError(t, err)

	type GetTsigKeyRequestInfo struct {
		ContainerId string
		Request     dns.GetTsigKeyRequest
	}

	var requests []GetTsigKeyRequestInfo
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
			response, err := c.GetTsigKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientGetView(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "GetView")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetView is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "GetView", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "GetView")
	assert.NoError(t, err)

	type GetViewRequestInfo struct {
		ContainerId string
		Request     dns.GetViewRequest
	}

	var requests []GetViewRequestInfo
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
			response, err := c.GetView(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientGetZone(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "GetZone")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetZone is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "GetZone", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "GetZone")
	assert.NoError(t, err)

	type GetZoneRequestInfo struct {
		ContainerId string
		Request     dns.GetZoneRequest
	}

	var requests []GetZoneRequestInfo
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
			response, err := c.GetZone(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientGetZoneRecords(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "GetZoneRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetZoneRecords is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "GetZoneRecords", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "GetZoneRecords")
	assert.NoError(t, err)

	type GetZoneRecordsRequestInfo struct {
		ContainerId string
		Request     dns.GetZoneRecordsRequest
	}

	var requests []GetZoneRecordsRequestInfo
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
				r := req.(*dns.GetZoneRecordsRequest)
				return c.GetZoneRecords(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.GetZoneRecordsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.GetZoneRecordsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientListResolverEndpoints(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "ListResolverEndpoints")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListResolverEndpoints is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "ListResolverEndpoints", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "ListResolverEndpoints")
	assert.NoError(t, err)

	type ListResolverEndpointsRequestInfo struct {
		ContainerId string
		Request     dns.ListResolverEndpointsRequest
	}

	var requests []ListResolverEndpointsRequestInfo
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
				r := req.(*dns.ListResolverEndpointsRequest)
				return c.ListResolverEndpoints(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.ListResolverEndpointsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.ListResolverEndpointsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientListResolvers(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "ListResolvers")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListResolvers is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "ListResolvers", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "ListResolvers")
	assert.NoError(t, err)

	type ListResolversRequestInfo struct {
		ContainerId string
		Request     dns.ListResolversRequest
	}

	var requests []ListResolversRequestInfo
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
				r := req.(*dns.ListResolversRequest)
				return c.ListResolvers(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.ListResolversResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.ListResolversResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientListSteeringPolicies(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "ListSteeringPolicies")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSteeringPolicies is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "ListSteeringPolicies", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "ListSteeringPolicies")
	assert.NoError(t, err)

	type ListSteeringPoliciesRequestInfo struct {
		ContainerId string
		Request     dns.ListSteeringPoliciesRequest
	}

	var requests []ListSteeringPoliciesRequestInfo
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
				r := req.(*dns.ListSteeringPoliciesRequest)
				return c.ListSteeringPolicies(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.ListSteeringPoliciesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.ListSteeringPoliciesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientListSteeringPolicyAttachments(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "ListSteeringPolicyAttachments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSteeringPolicyAttachments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "ListSteeringPolicyAttachments", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "ListSteeringPolicyAttachments")
	assert.NoError(t, err)

	type ListSteeringPolicyAttachmentsRequestInfo struct {
		ContainerId string
		Request     dns.ListSteeringPolicyAttachmentsRequest
	}

	var requests []ListSteeringPolicyAttachmentsRequestInfo
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
				r := req.(*dns.ListSteeringPolicyAttachmentsRequest)
				return c.ListSteeringPolicyAttachments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.ListSteeringPolicyAttachmentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.ListSteeringPolicyAttachmentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientListTsigKeys(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "ListTsigKeys")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTsigKeys is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "ListTsigKeys", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "ListTsigKeys")
	assert.NoError(t, err)

	type ListTsigKeysRequestInfo struct {
		ContainerId string
		Request     dns.ListTsigKeysRequest
	}

	var requests []ListTsigKeysRequestInfo
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
				r := req.(*dns.ListTsigKeysRequest)
				return c.ListTsigKeys(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.ListTsigKeysResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.ListTsigKeysResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientListViews(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "ListViews")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListViews is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "ListViews", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "ListViews")
	assert.NoError(t, err)

	type ListViewsRequestInfo struct {
		ContainerId string
		Request     dns.ListViewsRequest
	}

	var requests []ListViewsRequestInfo
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
				r := req.(*dns.ListViewsRequest)
				return c.ListViews(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.ListViewsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.ListViewsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientListZones(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "ListZones")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListZones is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "ListZones", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "ListZones")
	assert.NoError(t, err)

	type ListZonesRequestInfo struct {
		ContainerId string
		Request     dns.ListZonesRequest
	}

	var requests []ListZonesRequestInfo
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
				r := req.(*dns.ListZonesRequest)
				return c.ListZones(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.ListZonesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.ListZonesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientPatchDomainRecords(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "PatchDomainRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PatchDomainRecords is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "PatchDomainRecords", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "PatchDomainRecords")
	assert.NoError(t, err)

	type PatchDomainRecordsRequestInfo struct {
		ContainerId string
		Request     dns.PatchDomainRecordsRequest
	}

	var requests []PatchDomainRecordsRequestInfo
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
			response, err := c.PatchDomainRecords(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientPatchRRSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "PatchRRSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PatchRRSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "PatchRRSet", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "PatchRRSet")
	assert.NoError(t, err)

	type PatchRRSetRequestInfo struct {
		ContainerId string
		Request     dns.PatchRRSetRequest
	}

	var requests []PatchRRSetRequestInfo
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
			response, err := c.PatchRRSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientPatchZoneRecords(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "PatchZoneRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PatchZoneRecords is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "PatchZoneRecords", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "PatchZoneRecords")
	assert.NoError(t, err)

	type PatchZoneRecordsRequestInfo struct {
		ContainerId string
		Request     dns.PatchZoneRecordsRequest
	}

	var requests []PatchZoneRecordsRequestInfo
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
			response, err := c.PatchZoneRecords(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientUpdateDomainRecords(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "UpdateDomainRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDomainRecords is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "UpdateDomainRecords", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "UpdateDomainRecords")
	assert.NoError(t, err)

	type UpdateDomainRecordsRequestInfo struct {
		ContainerId string
		Request     dns.UpdateDomainRecordsRequest
	}

	var requests []UpdateDomainRecordsRequestInfo
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
			response, err := c.UpdateDomainRecords(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientUpdateRRSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "UpdateRRSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRRSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "UpdateRRSet", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "UpdateRRSet")
	assert.NoError(t, err)

	type UpdateRRSetRequestInfo struct {
		ContainerId string
		Request     dns.UpdateRRSetRequest
	}

	var requests []UpdateRRSetRequestInfo
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
			response, err := c.UpdateRRSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientUpdateResolver(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "UpdateResolver")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateResolver is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "UpdateResolver", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "UpdateResolver")
	assert.NoError(t, err)

	type UpdateResolverRequestInfo struct {
		ContainerId string
		Request     dns.UpdateResolverRequest
	}

	var requests []UpdateResolverRequestInfo
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
			response, err := c.UpdateResolver(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientUpdateResolverEndpoint(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "UpdateResolverEndpoint")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateResolverEndpoint is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "UpdateResolverEndpoint", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "UpdateResolverEndpoint")
	assert.NoError(t, err)

	type UpdateResolverEndpointRequestInfo struct {
		ContainerId string
		Request     dns.UpdateResolverEndpointRequest
	}

	var requests []UpdateResolverEndpointRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]UpdateResolverEndpointRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["UpdateResolverEndpointDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "endpointType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"VNIC": &dns.UpdateResolverVnicEndpointDetails{},
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
			response, err := c.UpdateResolverEndpoint(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientUpdateSteeringPolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "UpdateSteeringPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSteeringPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "UpdateSteeringPolicy", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "UpdateSteeringPolicy")
	assert.NoError(t, err)

	type UpdateSteeringPolicyRequestInfo struct {
		ContainerId string
		Request     dns.UpdateSteeringPolicyRequest
	}

	var requests []UpdateSteeringPolicyRequestInfo
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
			response, err := c.UpdateSteeringPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientUpdateSteeringPolicyAttachment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "UpdateSteeringPolicyAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSteeringPolicyAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "UpdateSteeringPolicyAttachment", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "UpdateSteeringPolicyAttachment")
	assert.NoError(t, err)

	type UpdateSteeringPolicyAttachmentRequestInfo struct {
		ContainerId string
		Request     dns.UpdateSteeringPolicyAttachmentRequest
	}

	var requests []UpdateSteeringPolicyAttachmentRequestInfo
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
			response, err := c.UpdateSteeringPolicyAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientUpdateTsigKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "UpdateTsigKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTsigKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "UpdateTsigKey", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "UpdateTsigKey")
	assert.NoError(t, err)

	type UpdateTsigKeyRequestInfo struct {
		ContainerId string
		Request     dns.UpdateTsigKeyRequest
	}

	var requests []UpdateTsigKeyRequestInfo
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
			response, err := c.UpdateTsigKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientUpdateView(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "UpdateView")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateView is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "UpdateView", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "UpdateView")
	assert.NoError(t, err)

	type UpdateViewRequestInfo struct {
		ContainerId string
		Request     dns.UpdateViewRequest
	}

	var requests []UpdateViewRequestInfo
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
			response, err := c.UpdateView(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientUpdateZone(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "UpdateZone")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateZone is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "UpdateZone", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "UpdateZone")
	assert.NoError(t, err)

	type UpdateZoneRequestInfo struct {
		ContainerId string
		Request     dns.UpdateZoneRequest
	}

	var requests []UpdateZoneRequestInfo
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
			response, err := c.UpdateZone(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_pubdns_dev_us_grp@oracle.com" jiraProject="PD" opsJiraProject="PUBDNS"
func TestDnsClientUpdateZoneRecords(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("dns", "UpdateZoneRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateZoneRecords is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "UpdateZoneRecords", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "UpdateZoneRecords")
	assert.NoError(t, err)

	type UpdateZoneRecordsRequestInfo struct {
		ContainerId string
		Request     dns.UpdateZoneRecordsRequest
	}

	var requests []UpdateZoneRecordsRequestInfo
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
			response, err := c.UpdateZoneRecords(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
