// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package kms

import (
    "github.com/oracle/oci-go-sdk/common"
    "net/http"
)

// UpdateVaultRequest wrapper for the UpdateVault operation
type UpdateVaultRequest struct {
        
 // The OCID of the Vault. 
        VaultId *string `mandatory:"true" contributesTo:"path" name:"vaultId"`
        
 // UpdateVaultDetails 
         UpdateVaultDetails `contributesTo:"body"`
        
 // For optimistic concurrency control. In the PUT or DELETE call for a
 // resource, set the `if-match` parameter to the value of the etag from a
 // previous GET or POST response for that resource.  The resource will be
 // updated or deleted only if the etag you provide matches the resource's
 // current etag value. 
        IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
        
 // Unique identifier for the request. If provided, the returned request ID
 // will include this value.  Otherwise, a random request ID will be
 // generated by the service. 
        OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

    // Metadata about the request. This information will not be transmitted to the service, but
    // represents information that the SDK will consume to drive retry behavior.
    RequestMetadata common.RequestMetadata
}

func (request UpdateVaultRequest) String() string {
    return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateVaultRequest) HTTPRequest(method, path string) (http.Request, error) {
    return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateVaultRequest) RetryPolicy() *common.RetryPolicy {
    return request.RequestMetadata.RetryPolicy
}

// UpdateVaultResponse wrapper for the UpdateVault operation
type UpdateVaultResponse struct {

    // The underlying http response
    RawResponse *http.Response
    
 // The Vault instance
     Vault `presentIn:"body"`

    
 // For optimistic concurrency control. See `if-match`.
    Etag *string `presentIn:"header" name:"etag"`
    
 // Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
 // a particular request, please provide the request ID.
    OpcRequestId *string `presentIn:"header" name:"opc-request-id"`


}

func (response UpdateVaultResponse) String() string {
    return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateVaultResponse) HTTPResponse() *http.Response {
    return response.RawResponse
}


