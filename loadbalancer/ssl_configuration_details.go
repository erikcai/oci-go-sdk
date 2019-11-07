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

// SslConfigurationDetails The load balancer's SSL handling configuration details.
// **Note:** User should ensure compatibility between specified SSL protocols and configured ciphers in the cipher suite, or else the request will be rejected
// **Note:** User should ensure compatibility between configured ciphers in the cipher suite and configured certificates (for ex.
// RSA based ciphers require RSA certificate whereas ECDSA based ciphers require ECDSA certificates)
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type SslConfigurationDetails struct {

	// A friendly name for the certificate bundle. It must be unique and it cannot be changed.
	// Valid certificate bundle names include only alphanumeric characters, dashes, and underscores.
	// Certificate bundle names cannot contain spaces. Avoid entering confidential information.
	// Example: `example_certificate_bundle`
	CertificateName *string `mandatory:"true" json:"certificateName"`

	// The maximum depth for peer certificate chain verification.
	// Example: `3`
	VerifyDepth *int `mandatory:"false" json:"verifyDepth"`

	// Whether the load balancer listener should verify peer certificates.
	// Example: `true`
	VerifyPeerCertificate *bool `mandatory:"false" json:"verifyPeerCertificate"`

	// A list of SSL protocols which needs to be supported by the load balancer for HTTPS or SSL connections.
	// SSL protocol is used to establish a secure connection between a client and a server which ensures that all the data passed
	// between the client and server is private.
	// The list of supported protocols:
	//       - "TLSv1"
	//       - "TLSv1.1"
	//       - "TLSv1.2"
	// If this field is not specified, TLSv1.2 is used by default.
	// **WARN:** All the virtual ssl listeners created on any given port should have same the list of SSL protocols defined
	// **Note:** The handshake in SSL connections establishment will fail if the client doesn't support any of the specified
	// protocols.
	// **Note:** User should ensure compatibility between specified SSL protocols and configured ciphers in the cipher suite
	// **Note:** For all existing load balancer listeners and bakcendsets that predates this feature, the GET
	//           operation on such load balancers/listeners will display the current list of ssl protocols being
	//           used by those listeners
	Protocols []string `mandatory:"false" json:"protocols"`

	// Name of the cipher suite to be used for HTTPS or SSL connections
	// If this field is not specified, oci-default-ssl-cipher-suite-v1 is used by default.
	// **Note:** User should ensure compatibility between specified SSL protocols and configured ciphers in the cipher suite,
	//           or else clients will not be able to successfully perform SSL handshake.
	// **Note:** User should ensure compatibility between configured ciphers in the cipher suite and configured certificates (for ex.
	//           RSA based ciphers require RSA certificate whereas ECDSA based ciphers require ECDSA certificates)
	// **Note:** For all existing load balancer listeners that predates this feature, the GET operation on such load
	//           balancers/listeners will display this field as "oci-default-ssl-cipher-suite-v1" inside listener's
	//           SSL configuration if the cipher configuration is not modified after the load balancer creation
	// **Note:** For all existing load balancer listeners that predates this feature, the GET operation on such load
	//           balancers/listeners will display this field as "oci-customized-ssl-cipher-suite" inside listener's
	//           SSL configuration, if the cipher configuration of those load balancers was customized after the load
	//           balancer creation via Oracle operations
	// **Note:** For all existing load balancer backendsets that predates this feature, the GET operation on such load
	//           balancers/backendsets will display this field as "oci-wider-compatible-ssl-cipher-suite-v1" inside
	//           backendset's SSL configuration
	// **Note:** If the GET operation on a load balancer listener display this field as "oci-customized-ssl-cipher-suite",
	//           then user must choose appropriate cipher suite name (either pre-defined or custom defined cipher suites)
	//           while updating such load balancers
	// **Note:** The following Oracle reserved cipher suite names are not accepted as valid input for this field:
	//           - oci-customized-ssl-cipher-suite
	CipherSuiteName *string `mandatory:"false" json:"cipherSuiteName"`

	// Specifies whether the server ciphers should be preferred over the client ciphers.
	// **Note:** This configuration is applicable only when the load balancer is acting as an SSL/HTTPS server i.e. when the SSLConfiguration object is
	// associated with the backend set, then this field is ignored.
	ServerOrderPreference SslConfigurationDetailsServerOrderPreferenceEnum `mandatory:"false" json:"serverOrderPreference,omitempty"`
}

func (m SslConfigurationDetails) String() string {
	return common.PointerString(m)
}

// SslConfigurationDetailsServerOrderPreferenceEnum Enum with underlying type: string
type SslConfigurationDetailsServerOrderPreferenceEnum string

// Set of constants representing the allowable values for SslConfigurationDetailsServerOrderPreferenceEnum
const (
	SslConfigurationDetailsServerOrderPreferenceEnabled  SslConfigurationDetailsServerOrderPreferenceEnum = "ENABLED"
	SslConfigurationDetailsServerOrderPreferenceDisabled SslConfigurationDetailsServerOrderPreferenceEnum = "DISABLED"
)

var mappingSslConfigurationDetailsServerOrderPreference = map[string]SslConfigurationDetailsServerOrderPreferenceEnum{
	"ENABLED":  SslConfigurationDetailsServerOrderPreferenceEnabled,
	"DISABLED": SslConfigurationDetailsServerOrderPreferenceDisabled,
}

// GetSslConfigurationDetailsServerOrderPreferenceEnumValues Enumerates the set of values for SslConfigurationDetailsServerOrderPreferenceEnum
func GetSslConfigurationDetailsServerOrderPreferenceEnumValues() []SslConfigurationDetailsServerOrderPreferenceEnum {
	values := make([]SslConfigurationDetailsServerOrderPreferenceEnum, 0)
	for _, v := range mappingSslConfigurationDetailsServerOrderPreference {
		values = append(values, v)
	}
	return values
}
