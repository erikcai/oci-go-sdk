// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RedirectUri The redirect URI to which the original request URI should be redirected. Redirect URI is constructed based
// on the values provided for protocol, host, port, path and query properties. The Load Balancer service,
// can not automatically detect or avoid infinite redirects, so it is user's responsibility to provide meaningful
// and correct values for these fields.
// In addition to static string values, tokens below can be used to fetch the runtime values from the runtime HTTP
// request URI and use them in formulating the redirect URI.
// Tokens are case sensitive ({host} will be treated as token while {HOST} is not a token).
//   {protocol} : corresponds to protocol from the runtime HTTP request URI.
//   {host}     : corresponds to the domain name from the runtime HTTP request URI.
//   {port}     : corresponds to the port from the runtime HTTP request URI.
//   {path}     : corresponds to the path from the runtime HTTP request URI.
//   {query}    : corresponds to the query string from the runtime HTTP request URI.
// '\' should be used as escape character for escaping \, {, } characters while specifying values for path and query
// properties to retain the token as is in redirect URI.
// For example, when '/example{path}123\{path\}' is provided as value for the path property, it gets translated as
// '/example/video123{path}' in redirect URI when the path in runtime HTTP request URI is /video.
type RedirectUri struct {

	// Protocol is the HTTP protocol to be configured in the redirect URI.
	// When the value is not set, or set to null or {protocol}, original protocol in the runtime HTTP request URI will
	// be preserved. Allowed values are HTTP, HTTPS, {protocol}
	// No other token is allowed here except {protocol} and the token can be mentioned only once.
	// Example: HTTPS
	Protocol *string `mandatory:"false" json:"protocol"`

	// Host is the valid domain name (hostname) or IP address to be configured in the redirect URI.
	// When the value is not set, or set to null or {host}, the original domain name in the runtime HTTP request URI
	// will be preserved. All the tokens, even more than once, are allowed here. Curly braces are allowed for this
	// parameter only when surrounding tokens (like {host}).
	// Examples: example.com   translates to example.com in redirect URI.
	//          in{host}      translates to inexample.com in redirect URI when example.com is the host name
	//                        in runtime HTTP request URI.
	//          {port}{host}  translates to 8081example.com in redirect URI when example.com is the host name
	//                        and port is 8081 in runtime HTTP request URI.
	Host *string `mandatory:"false" json:"host"`

	// Port is the communication port to be configured in the redirect URI.
	// Values should be integer in the range of 1 to 65535.
	// When the value is null, the original port in the runtime HTTP request URI will be preserved.
	// Example: 8081          translates to 8081 in redirect URI.
	Port *int `mandatory:"false" json:"port"`

	// Path is the HTTP URI path to be configured in the redirect URI.
	// When the value is not set, or set to null or {path}, the original path in the runtime HTTP request URI will be
	// preserved. If the path needs to be dropped in the redirect URI, then this value must be set to "".
	// All the tokens, even more than once, are allowed here.
	// Path must be starting with '/' if it is not starting with {path} token.
	// Examples: /example/video/123  translates to /example/video/123 in redirect URI.
	//          /example{path}      translates to /example/video/123 in redirect URI when /video/123 is the
	//                              path in runtime HTTP request URI.
	//          {path}/123          translates to /example/video/123 in redirect URI when /example/video is
	//                              the path in runtime HTTP request URI.
	//          {path}123           translates to /example/video123 in redirect URI when /example/video is
	//                              the path in runtime HTTP request URI.
	//          /{host}/123         translates to /example.com/video/123 in redirect URI when example.com is
	//                              the host name in runtime HTTP request URI.
	//          /{host}/{port}      translates to /example.com/123 in redirect URI when example.com is
	//                              the host name and 123 is the port in runtime HTTP request URI.
	//          /{query}            translates to /lang=en in redirect URI when  query is lang=en in runtime
	//                              HTTP request URI.
	Path *string `mandatory:"false" json:"path"`

	// The query is the query string to be configured in the redirect URI.
	// When the value is not set, or set to null or {query}, original query parameters in the runtime HTTP request
	// URI will be preserved. All the tokens, even more than once, are allowed here.
	// Multiple query parameters can be specified as single string and they must be separated with '&'.
	// In order to drop all query parameters from input URI, set this attribute to an empty string i.e. ""
	// Examples: lang=en&time_zone=PST  translates to lang=en&time_zone=PST in redirect URI.
	//          {query}   translates to lang=en&time_zone=PST when the query string in runtime HTTP request is
	//                    lang=en&time_zone=PST , and translates to empty string when the runtime HTTP request does
	//                    not have any query parameters.
	//          lang=en&{query}&time_zone=PST translates to lang=en&country=us&time_zone=PST when the query string
	//                                        in runtime HTTP request is country=us , and translates to
	//                                        lang=en&time_zone=PST when the runtime HTTP request does not have any
	//                                        query parameters.
	//          protocol={protocol}&hostname={host}  translates to
	//                                               protocol=http&hostname=example.com when protocol is HTTP
	//                                               and hostname is example.com in the runtime HTTP request
	//                                               URI.
	//          port={port}&hostname={host} translates to port=8080&hostname=example.com when port is 8080
	//                                      and hostname is example.com in the runtime HTTP request URI.
	Query *string `mandatory:"false" json:"query"`
}

func (m RedirectUri) String() string {
	return common.PointerString(m)
}
