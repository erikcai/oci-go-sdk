// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
// 
 // Cloud Incident Management Service
//

package cims


    // ErrorCodeEnum Enum with underlying type: string
type ErrorCodeEnum string

// Set of constants representing the allowable values for ErrorCodeEnum
const (
    ErrorCodeContentEmpty ErrorCodeEnum = "CONTENT_EMPTY"
    ErrorCodeClientException ErrorCodeEnum = "CLIENT_EXCEPTION"
    ErrorCodeInvalidFormat ErrorCodeEnum = "INVALID_FORMAT"
    ErrorCodeInvalidJsonInput ErrorCodeEnum = "INVALID_JSON_INPUT"
    ErrorCodeSslAuthorization ErrorCodeEnum = "SSL_AUTHORIZATION"
    ErrorCodeAuthFailed ErrorCodeEnum = "AUTH_FAILED"
    ErrorCodeCsiNotAuthorized ErrorCodeEnum = "CSI_NOT_AUTHORIZED"
    ErrorCodeUserPolicyNotAuthorized ErrorCodeEnum = "USER_POLICY_NOT_AUTHORIZED"
    ErrorCodeEmailNotVerified ErrorCodeEnum = "EMAIL_NOT_VERIFIED"
    ErrorCodeEmailNotFound ErrorCodeEnum = "EMAIL_NOT_FOUND"
    ErrorCodeIdcsEmailNotValid ErrorCodeEnum = "IDCS_EMAIL_NOT_VALID"
    ErrorCodeInvalidPath ErrorCodeEnum = "INVALID_PATH"
    ErrorCodeMethodNotAllowed ErrorCodeEnum = "METHOD_NOT_ALLOWED"
    ErrorCodeJsonProcessing ErrorCodeEnum = "JSON_PROCESSING"
    ErrorCodeGenericException ErrorCodeEnum = "GENERIC_EXCEPTION"
    ErrorCodeExternalServiceProviderUnavailable ErrorCodeEnum = "EXTERNAL_SERVICE_PROVIDER_UNAVAILABLE"
    ErrorCodeExternalServiceProviderTimeout ErrorCodeEnum = "EXTERNAL_SERVICE_PROVIDER_TIMEOUT"
)

var mappingErrorCode = map[string]ErrorCodeEnum { 
    "CONTENT_EMPTY": ErrorCodeContentEmpty,
    "CLIENT_EXCEPTION": ErrorCodeClientException,
    "INVALID_FORMAT": ErrorCodeInvalidFormat,
    "INVALID_JSON_INPUT": ErrorCodeInvalidJsonInput,
    "SSL_AUTHORIZATION": ErrorCodeSslAuthorization,
    "AUTH_FAILED": ErrorCodeAuthFailed,
    "CSI_NOT_AUTHORIZED": ErrorCodeCsiNotAuthorized,
    "USER_POLICY_NOT_AUTHORIZED": ErrorCodeUserPolicyNotAuthorized,
    "EMAIL_NOT_VERIFIED": ErrorCodeEmailNotVerified,
    "EMAIL_NOT_FOUND": ErrorCodeEmailNotFound,
    "IDCS_EMAIL_NOT_VALID": ErrorCodeIdcsEmailNotValid,
    "INVALID_PATH": ErrorCodeInvalidPath,
    "METHOD_NOT_ALLOWED": ErrorCodeMethodNotAllowed,
    "JSON_PROCESSING": ErrorCodeJsonProcessing,
    "GENERIC_EXCEPTION": ErrorCodeGenericException,
    "EXTERNAL_SERVICE_PROVIDER_UNAVAILABLE": ErrorCodeExternalServiceProviderUnavailable,
    "EXTERNAL_SERVICE_PROVIDER_TIMEOUT": ErrorCodeExternalServiceProviderTimeout,
}

// GetErrorCodeEnumValues Enumerates the set of values for ErrorCodeEnum
func GetErrorCodeEnumValues() []ErrorCodeEnum {
   values := make([]ErrorCodeEnum, 0)
   for _, v := range mappingErrorCode {
       values = append(values, v)
   }
   return values
}


