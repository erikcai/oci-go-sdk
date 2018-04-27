// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateListenerDetails The configuration details for adding a listener to a backend set.
// For more information on listener configuration, see
// Managing Load Balancer Listeners (https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/tasks/managinglisteners.htm).
type CreateListenerDetails struct {

	// The name of the associated backend set.
	// Example: `example_backend_set`
	DefaultBackendSetName *string `mandatory:"true" json:"defaultBackendSetName"`

	// A friendly name for the listener. It must be unique and it cannot be changed.
	// Avoid entering confidential information.
	// Example: `example_listener`
	Name *string `mandatory:"true" json:"name"`

	// The communication port for the listener.
	// Example: `80`
	Port *int `mandatory:"true" json:"port"`

	// The protocol on which the listener accepts connection requests.
	// To get a list of valid protocols, use the ListProtocols
	// operation.
	// Example: `HTTP`
	Protocol *string `mandatory:"true" json:"protocol"`

	ConnectionConfiguration *ConnectionConfiguration `mandatory:"false" json:"connectionConfiguration"`

	// An array of hostname resource names.
	HostnameNames []string `mandatory:"false" json:"hostnameNames"`

	// The name of the set of path-based routing rules, PathRouteSet,
	// applied to this listener's traffic.
	// Example: `example_path_route_set`
	PathRouteSetName *string `mandatory:"false" json:"pathRouteSetName"`

	// The name of the rule sets to apply to the listener.
	// Example: ["example_http_rule_list"]
	RuleSetNames []string `mandatory:"false" json:"ruleSetNames"`

	// Deprecated. Use `hostnames` instead.
	// Specifies a virtual host name for this listener.
	// Example: `app.example.com`
	// This feature supports HTTP and HTTPS listeners only. It does not support TCP listeners.
	// You can define an exact virtual hostname, as in the preceding example, or you can use a wildcard name. A
	// wildcard name includes an asterisk (*) in place of the first or last part of the name. When searching for a
	// virtual hostname, the service chooses the first matching variant in the following priority order:
	// 1. Exact name match (no asterisk), such as `app.example.com`.
	// 2. Longest wildcard name that begins with an asterisk, such as `*.example.com`.
	//    Note: Prefix wildcard names might require a wildcard certificate for HTTPS sites.
	// 3. Longest wildcard name that ends with an asterisk, such as `app.example.*`.
	//    Note: Suffix wildcard names might require a multi-domain Subject Alternative Name (SAN) certificate for HTTPS sites.
	// You do not need to specify the matching pattern to apply. It is inherent in the `serverName` asterisk position,
	// that is, starting, ending, or none. Server name selection priority is not related to the virtual server's
	// configuration order.
	// You cannot use regular expressions.
	// If a listener has no `serverName` specified, it is the default listener on the assigned port.
	// If all listeners on a port have virtual host names, the first virtual host configured for the port serves as
	// the default listener.
	// For information on the interaction between virtual servers and path route rules, see
	// Managing Request Routing (https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Tasks/managingrequest.htm).
	ServerName *string `mandatory:"false" json:"serverName"`

	SslConfiguration *SslConfigurationDetails `mandatory:"false" json:"sslConfiguration"`
}

func (m CreateListenerDetails) String() string {
	return common.PointerString(m)
}
