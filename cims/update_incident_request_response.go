// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cims

import (
    "github.com/oracle/oci-go-sdk/common"
    "net/http"
)

// UpdateIncidentRequest wrapper for the UpdateIncident operation
type UpdateIncidentRequest struct {
        
 // Unique identifier for the support ticket. 
        IncidentKey *string `mandatory:"true" contributesTo:"path" name:"incidentKey"`
        
 // The Customer Support Identifier associated with the support account. 
        Csi *string `mandatory:"true" contributesTo:"header" name:"csi"`
        
 // Details about the support ticket being updated. 
        UpdateIncidentDetails UpdateIncident `contributesTo:"body"`
        
 // User OCID for Oracle Identity Cloud Service (IDCS) users who also have a federated Oracle Cloud Infrastructure account. 
        Ocid *string `mandatory:"true" contributesTo:"header" name:"ocid"`
        
 // A token that uniquely identifies a request so it can be retried in case of a timeout or server error without risk of executing that same action again. Retry tokens expire after 24 hours, but can be invalidated before then due to conflicting operations. For example, if a resource has been deleted and purged from the system, then a retry of the original creation request might be rejected. 
        OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`
        
 // Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID. 
        OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`
        
 // For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match` parameter to the value of the etag from a previous GET or POST response for that resource. The resource will be updated or deleted only if the etag you provide matches the resource's current etag value. 
        IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`


    // Metadata about the request. This information will not be transmitted to the service, but
    // represents information that the SDK will consume to drive retry behavior.
    RequestMetadata common.RequestMetadata
}

func (request UpdateIncidentRequest) String() string {
    return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateIncidentRequest) HTTPRequest(method, path string) (http.Request, error) {
    return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateIncidentRequest) RetryPolicy() *common.RetryPolicy {
    return request.RequestMetadata.RetryPolicy
}

// UpdateIncidentResponse wrapper for the UpdateIncident operation
type UpdateIncidentResponse struct {

    // The underlying http response
    RawResponse *http.Response
    
 // The Incident instance
     Incident `presentIn:"body"`

    
 // Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
    OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
    
 // The entity tag that allows optimistic concurrency control. For more information, see REST APIs (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#eleven).
    Etag *string `presentIn:"header" name:"etag"`


}

func (response UpdateIncidentResponse) String() string {
    return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateIncidentResponse) HTTPResponse() *http.Response {
    return response.RawResponse
}


