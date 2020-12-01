// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package keymanagement

import (
	"github.com/oracle/oci-go-sdk/v30/common"
	"net/http"
)

// ListKeysRequest wrapper for the ListKeys operation
type ListKeysRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can specify only one sort order. The default
	// order for `TIMECREATED` is descending. The default order for `DISPLAYNAME`
	// is ascending.
	SortBy ListKeysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListKeysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A key's protection mode indicates how the key persists and where cryptographic operations that use the key are performed. A
	// protection mode of `HSM` means that the key persists on a hardware security module (HSM) and all cryptographic operations are
	// performed inside the HSM. A protection mode of `SOFTWARE` means that the key persists on the server, protected by the vault's
	// RSA wrapping key which persists on the HSM. All cryptographic operations that use a key with a protection mode of
	// `SOFTWARE` are performed on the server.
	ProtectionMode ListKeysProtectionModeEnum `mandatory:"false" contributesTo:"query" name:"protectionMode" omitEmpty:"true"`

	// The algorithm used by a key's key versions to encrypt or decrypt. Currently, only AES, RSA and ECDSA are supported.
	Algorithm ListKeysAlgorithmEnum `mandatory:"false" contributesTo:"query" name:"algorithm" omitEmpty:"true"`

	// The length of the key in bytes, expressed as an integer. Values of 16, 24, 32 are supported.
	Length *int `mandatory:"false" contributesTo:"query" name:"length"`

	// The curve Id of the keys in case of ECDSA keys
	CurveId ListKeysCurveIdEnum `mandatory:"false" contributesTo:"query" name:"curveId" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListKeysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListKeysRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListKeysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListKeysResponse wrapper for the ListKeys operation
type ListKeysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []KeySummary instances
	Items []KeySummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListKeysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListKeysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListKeysSortByEnum Enum with underlying type: string
type ListKeysSortByEnum string

// Set of constants representing the allowable values for ListKeysSortByEnum
const (
	ListKeysSortByTimecreated ListKeysSortByEnum = "TIMECREATED"
	ListKeysSortByDisplayname ListKeysSortByEnum = "DISPLAYNAME"
)

var mappingListKeysSortBy = map[string]ListKeysSortByEnum{
	"TIMECREATED": ListKeysSortByTimecreated,
	"DISPLAYNAME": ListKeysSortByDisplayname,
}

// GetListKeysSortByEnumValues Enumerates the set of values for ListKeysSortByEnum
func GetListKeysSortByEnumValues() []ListKeysSortByEnum {
	values := make([]ListKeysSortByEnum, 0)
	for _, v := range mappingListKeysSortBy {
		values = append(values, v)
	}
	return values
}

// ListKeysSortOrderEnum Enum with underlying type: string
type ListKeysSortOrderEnum string

// Set of constants representing the allowable values for ListKeysSortOrderEnum
const (
	ListKeysSortOrderAsc  ListKeysSortOrderEnum = "ASC"
	ListKeysSortOrderDesc ListKeysSortOrderEnum = "DESC"
)

var mappingListKeysSortOrder = map[string]ListKeysSortOrderEnum{
	"ASC":  ListKeysSortOrderAsc,
	"DESC": ListKeysSortOrderDesc,
}

// GetListKeysSortOrderEnumValues Enumerates the set of values for ListKeysSortOrderEnum
func GetListKeysSortOrderEnumValues() []ListKeysSortOrderEnum {
	values := make([]ListKeysSortOrderEnum, 0)
	for _, v := range mappingListKeysSortOrder {
		values = append(values, v)
	}
	return values
}

// ListKeysProtectionModeEnum Enum with underlying type: string
type ListKeysProtectionModeEnum string

// Set of constants representing the allowable values for ListKeysProtectionModeEnum
const (
	ListKeysProtectionModeHsm      ListKeysProtectionModeEnum = "HSM"
	ListKeysProtectionModeSoftware ListKeysProtectionModeEnum = "SOFTWARE"
)

var mappingListKeysProtectionMode = map[string]ListKeysProtectionModeEnum{
	"HSM":      ListKeysProtectionModeHsm,
	"SOFTWARE": ListKeysProtectionModeSoftware,
}

// GetListKeysProtectionModeEnumValues Enumerates the set of values for ListKeysProtectionModeEnum
func GetListKeysProtectionModeEnumValues() []ListKeysProtectionModeEnum {
	values := make([]ListKeysProtectionModeEnum, 0)
	for _, v := range mappingListKeysProtectionMode {
		values = append(values, v)
	}
	return values
}

// ListKeysAlgorithmEnum Enum with underlying type: string
type ListKeysAlgorithmEnum string

// Set of constants representing the allowable values for ListKeysAlgorithmEnum
const (
	ListKeysAlgorithmAes   ListKeysAlgorithmEnum = "AES"
	ListKeysAlgorithmRsa   ListKeysAlgorithmEnum = "RSA"
	ListKeysAlgorithmEcdsa ListKeysAlgorithmEnum = "ECDSA"
)

var mappingListKeysAlgorithm = map[string]ListKeysAlgorithmEnum{
	"AES":   ListKeysAlgorithmAes,
	"RSA":   ListKeysAlgorithmRsa,
	"ECDSA": ListKeysAlgorithmEcdsa,
}

// GetListKeysAlgorithmEnumValues Enumerates the set of values for ListKeysAlgorithmEnum
func GetListKeysAlgorithmEnumValues() []ListKeysAlgorithmEnum {
	values := make([]ListKeysAlgorithmEnum, 0)
	for _, v := range mappingListKeysAlgorithm {
		values = append(values, v)
	}
	return values
}

// ListKeysCurveIdEnum Enum with underlying type: string
type ListKeysCurveIdEnum string

// Set of constants representing the allowable values for ListKeysCurveIdEnum
const (
	ListKeysCurveIdP256 ListKeysCurveIdEnum = "NIST_P256"
	ListKeysCurveIdP384 ListKeysCurveIdEnum = "NIST_P384"
	ListKeysCurveIdP521 ListKeysCurveIdEnum = "NIST_P521"
)

var mappingListKeysCurveId = map[string]ListKeysCurveIdEnum{
	"NIST_P256": ListKeysCurveIdP256,
	"NIST_P384": ListKeysCurveIdP384,
	"NIST_P521": ListKeysCurveIdP521,
}

// GetListKeysCurveIdEnumValues Enumerates the set of values for ListKeysCurveIdEnum
func GetListKeysCurveIdEnumValues() []ListKeysCurveIdEnum {
	values := make([]ListKeysCurveIdEnum, 0)
	for _, v := range mappingListKeysCurveId {
		values = append(values, v)
	}
	return values
}
