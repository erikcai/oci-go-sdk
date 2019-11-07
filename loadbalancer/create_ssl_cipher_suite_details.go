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

// CreateSslCipherSuiteDetails The configuration details for adding a SSL cipher suite to a load balancer.
// A cipher suite is a set of algorithms that help secure a network connection that uses Transport Layer Security (TLS) or Secure Socket Layer (SSL).
// Cipher suites gives you the ability to control the list of security algorithms that your load balancer uses to negotiate with the peers while
// sending and receiving information on a network connection using TLS or SSL.
// To ease the user configuration, Oracle has pre-defined the following three cipher suites:
//              - "oci-default-ssl-cipher-suite-v1"
//                         "ECDHE-RSA-AES128-GCM-SHA256"
//                         "ECDHE-RSA-AES128-SHA256"
//                         "ECDHE-RSA-AES256-GCM-SHA384"
//                         "ECDHE-RSA-AES256-SHA384"
//                         "DHE-RSA-AES256-GCM-SHA384"
//                         "DHE-RSA-AES256-SHA256"
//                         "DHE-RSA-AES128-GCM-SHA256"
//                         "DHE-RSA-AES128-SHA256"
//              - "oci-modern-ssl-cipher-suite-v1"
//                         "ECDHE-ECDSA-AES128-GCM-SHA256"
//                         "ECDHE-RSA-AES128-GCM-SHA256"
//                         "ECDHE-ECDSA-AES128-SHA256"
//                         "ECDHE-RSA-AES128-SHA256"
//                         "ECDHE-ECDSA-AES256-GCM-SHA384"
//                         "ECDHE-RSA-AES256-GCM-SHA384"
//                         "ECDHE-ECDSA-AES256-SHA384"
//                         "ECDHE-RSA-AES256-SHA384"
//                         "AES128-GCM-SHA256"
//                         "AES128-SHA256"
//                         "AES256-GCM-SHA384"
//                         "AES256-SHA256"
//                         "DHE-RSA-AES256-GCM-SHA384"
//                         "DHE-RSA-AES256-SHA256"
//                         "DHE-RSA-AES128-GCM-SHA256"
//                         "DHE-RSA-AES128-SHA256"
//              - "oci-compatible-ssl-cipher-suite-v1"
//                         "ECDHE-ECDSA-AES128-GCM-SHA256"
//                         "ECDHE-RSA-AES128-GCM-SHA256"
//                         "ECDHE-ECDSA-AES128-SHA256"
//                         "ECDHE-RSA-AES128-SHA256"
//                         "ECDHE-ECDSA-AES128-SHA"
//                         "ECDHE-RSA-AES128-SHA"
//                         "ECDHE-ECDSA-AES256-GCM-SHA384"
//                         "ECDHE-RSA-AES256-GCM-SHA384"
//                         "ECDHE-ECDSA-AES256-SHA384"
//                         "ECDHE-RSA-AES256-SHA384"
//                         "ECDHE-RSA-AES256-SHA"
//                         "ECDHE-ECDSA-AES256-SHA"
//                         "AES128-GCM-SHA256"
//                         "AES128-SHA256"
//                         "AES128-SHA"
//                         "AES256-GCM-SHA384"
//                         "AES256-SHA256"
//                         "AES256-SHA"
//                         "DHE-RSA-AES256-GCM-SHA384"
//                         "DHE-RSA-AES256-SHA256"
//                         "DHE-RSA-AES128-GCM-SHA256"
//                         "DHE-RSA-AES128-SHA256"
//              - "oci-wider-compatible-ssl-cipher-suite-v1"
//                         "ECDHE-ECDSA-AES128-GCM-SHA256"
//                         "ECDHE-RSA-AES128-GCM-SHA256"
//                         "ECDHE-ECDSA-AES128-SHA256"
//                         "ECDHE-RSA-AES128-SHA256"
//                         "ECDHE-ECDSA-AES256-GCM-SHA384"
//                         "ECDHE-RSA-AES256-GCM-SHA384"
//                         "ECDHE-ECDSA-AES256-SHA384"
//                         "ECDHE-RSA-AES256-SHA384"
//                         "AES128-GCM-SHA256"
//                         "AES128-SHA256"
//                         "AES256-GCM-SHA384"
//                         "AES256-SHA256"
//                         "DHE-RSA-AES256-GCM-SHA384"
//                         "DHE-RSA-AES256-SHA256"
//                         "DHE-RSA-AES128-GCM-SHA256"
//                         "DHE-RSA-AES128-SHA256"
//                         "DH-DSS-AES256-GCM-SHA384"
//                         "DHE-DSS-AES256-GCM-SHA384"
//                         "DH-RSA-AES256-GCM-SHA384"
//                         "DHE-DSS-AES256-SHA256"
//                         "DH-RSA-AES256-SHA256"
//                         "DH-DSS-AES256-SHA256"
//                         "ECDH-RSA-AES256-GCM-SHA384"
//                         "ECDH-ECDSA-AES256-GCM-SHA384"
//                         "ECDH-RSA-AES256-SHA384"
//                         "ECDH-ECDSA-AES256-SHA384"
//                         "DH-DSS-AES128-GCM-SHA256"
//                         "DHE-DSS-AES128-GCM-SHA256"
//                         "DH-RSA-AES128-GCM-SHA256"
//                         "DHE-DSS-AES128-SHA256"
//                         "DH-RSA-AES128-SHA256"
//                         "DH-DSS-AES128-SHA256"
//                         "ECDH-RSA-AES128-GCM-SHA256"
//                         "ECDH-ECDSA-AES128-GCM-SHA256"
//                         "ECDH-RSA-AES128-SHA256"
//                         "ECDH-ECDSA-AES128-SHA256"
//                         "ECDHE-ECDSA-AES128-SHA"
//                         "ECDHE-RSA-AES128-SHA"
//                         "ECDHE-RSA-AES256-SHA"
//                         "ECDHE-ECDSA-AES256-SHA"
//                         "AES128-SHA"
//                         "AES256-SHA"
//                         "DHE-RSA-AES128-SHA"
//                         "DHE-RSA-CAMELLIA256-SHA"
//                         "DHE-RSA-CAMELLIA128-SHA"
//                         "DHE-DSS-CAMELLIA256-SHA"
//                         "DHE-DSS-CAMELLIA128-SHA"
//                         "DHE-RSA-SEED-SHA"
//                         "DHE-DSS-SEED-SHA"
//                         "DH-RSA-SEED-SHA"
//                         "DH-DSS-SEED-SHA"
//                         "DHE-RSA-AES256-SHA"
//                         "DHE-DSS-AES256-SHA"
//                         "DH-RSA-AES256-SHA"
//                         "DH-DSS-AES256-SHA"
//                         "DH-RSA-CAMELLIA256-SHA"
//                         "DH-DSS-CAMELLIA256-SHA"
//                         "ECDH-RSA-AES256-SHA"
//                         "ECDH-ECDSA-AES256-SHA"
//                         "CAMELLIA256-SHA"
//                         "PSK-AES256-CBC-SHA"
//                         "DHE-DSS-AES128-SHA"
//                         "DH-RSA-AES128-SHA"
//                         "DH-DSS-AES128-SHA"
//                         "DH-RSA-CAMELLIA128-SHA"
//                         "DH-DSS-CAMELLIA128-SHA"
//                         "ECDH-RSA-AES128-SHA"
//                         "ECDH-ECDSA-AES128-SHA"
//                         "SEED-SHA"
//                         "CAMELLIA128-SHA"
//                         "PSK-AES128-CBC-SHA"
//                         "DES-CBC3-SHA"
//                         "IDEA-CBC-SHA"
//                         "ECDHE-RSA-DES-CBC3-SHA"
//                         "ECDHE-ECDSA-DES-CBC3-SHA"
//                         "DHE-RSA-DES-CBC3-SHA"
//                         "DHE-DSS-DES-CBC3-SHA"
//                         "DH-RSA-DES-CBC3-SHA"
//                         "DH-DSS-DES-CBC3-SHAv"
//                         "ECDH-RSA-DES-CBC3-SHA"
//                         "ECDH-ECDSA-DES-CBC3-SHA"
//                         "PSK-3DES-EDE-CBC-SHA"
//                         "KRB5-IDEA-CBC-SHA"
//                         "KRB5-DES-CBC3-SHA"
//                         "KRB5-IDEA-CBC-MD5"
//                         "KRB5-DES-CBC3-MD5"
//                         "ECDHE-RSA-RC4-SHA"
//                         "ECDHE-ECDSA-RC4-SHA"
//                         "ECDH-RSA-RC4-SHA"
//                         "ECDH-ECDSA-RC4-SHA"
//                         "RC4-SHA"
//                         "RC4-MD5"
//                         "PSK-RC4-SHA"
//                         "KRB5-RC4-SHA"
//                         "KRB5-RC4-MD5"
//
// You can create custom cipher suites if the pre-defined cipher suites doesn't meet the required criteria.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateSslCipherSuiteDetails struct {

	// A friendly name for the SSL cipher suite. It must be unique and it cannot be changed.
	// **Note:** The name of user defined cipher suite should not be same as any of the following Oracle's pre-defined or reserved ssl cipher suite names:
	//           - oci-default-ssl-cipher-suite-v1
	//           - oci-modern-ssl-cipher-suite-v1
	//           - oci-compatible-ssl-cipher-suite-v1
	//           - oci-wider-compatible-ssl-cipher-suite-v1
	//           - oci-customized-ssl-cipher-suite
	Name *string `mandatory:"true" json:"name"`

	// A list of SSL ciphers which needs to be supported by the load balancer for HTTPS or SSL connections.
	// The list of supported ciphers:
	//       #### TLSv1.2 ciphers
	//       - "ECDHE-ECDSA-AES128-GCM-SHA256"
	//       - "ECDHE-RSA-AES128-GCM-SHA256"
	//       - "ECDHE-ECDSA-AES128-SHA256"
	//       - "ECDHE-RSA-AES128-SHA256"
	//       - "ECDHE-ECDSA-AES256-GCM-SHA384"
	//       - "ECDHE-RSA-AES256-GCM-SHA384"
	//       - "ECDHE-ECDSA-AES256-SHA384"
	//       - "ECDHE-RSA-AES256-SHA384"
	//       - "AES128-GCM-SHA256"
	//       - "AES128-SHA256"
	//       - "AES256-GCM-SHA384"
	//       - "AES256-SHA256"
	//       - "DHE-RSA-AES256-GCM-SHA384"
	//       - "DHE-RSA-AES256-SHA256"
	//       - "DHE-RSA-AES128-GCM-SHA256"
	//       - "DHE-RSA-AES128-SHA256"
	//       - "DH-DSS-AES256-GCM-SHA384"
	//       - "DHE-DSS-AES256-GCM-SHA384"
	//       - "DH-RSA-AES256-GCM-SHA384"
	//       - "DHE-DSS-AES256-SHA256"
	//       - "DH-RSA-AES256-SHA256"
	//       - "DH-DSS-AES256-SHA256"
	//       - "ECDH-RSA-AES256-GCM-SHA384"
	//       - "ECDH-ECDSA-AES256-GCM-SHA384"
	//       - "ECDH-RSA-AES256-SHA384"
	//       - "ECDH-ECDSA-AES256-SHA384"
	//       - "DH-DSS-AES128-GCM-SHA256"
	//       - "DHE-DSS-AES128-GCM-SHA256"
	//       - "DH-RSA-AES128-GCM-SHA256"
	//       - "DHE-DSS-AES128-SHA256"
	//       - "DH-RSA-AES128-SHA256"
	//       - "DH-DSS-AES128-SHA256"
	//       - "ECDH-RSA-AES128-GCM-SHA256"
	//       - "ECDH-ECDSA-AES128-GCM-SHA256"
	//       - "ECDH-RSA-AES128-SHA256"
	//       - "ECDH-ECDSA-AES128-SHA256"
	//       #### TLSv1 ciphers which are also supported by TLSv1.2
	//       - "ECDHE-ECDSA-AES128-SHA"
	//       - "ECDHE-RSA-AES128-SHA"
	//       - "ECDHE-RSA-AES256-SHA"
	//       - "ECDHE-ECDSA-AES256-SHA"
	//       - "AES128-SHA"
	//       - "AES256-SHA"
	//       - "DHE-RSA-AES128-SHA"
	//       - "DHE-RSA-CAMELLIA256-SHA"
	//       - "DHE-RSA-CAMELLIA128-SHA"
	//       - "DHE-DSS-CAMELLIA256-SHA"
	//       - "DHE-DSS-CAMELLIA128-SHA"
	//       - "DHE-RSA-SEED-SHA"
	//       - "DHE-DSS-SEED-SHA"
	//       - "DH-RSA-SEED-SHA"
	//       - "DH-DSS-SEED-SHA"
	//       - "DHE-RSA-AES256-SHA"
	//       - "DHE-DSS-AES256-SHA"
	//       - "DH-RSA-AES256-SHA"
	//       - "DH-DSS-AES256-SHA"
	//       - "DH-RSA-CAMELLIA256-SHA"
	//       - "DH-DSS-CAMELLIA256-SHA"
	//       - "ECDH-RSA-AES256-SHA"
	//       - "ECDH-ECDSA-AES256-SHA"
	//       - "CAMELLIA256-SHA"
	//       - "PSK-AES256-CBC-SHA"
	//       - "DHE-DSS-AES128-SHA"
	//       - "DH-RSA-AES128-SHA"
	//       - "DH-DSS-AES128-SHA"
	//       - "DH-RSA-CAMELLIA128-SHA"
	//       - "DH-DSS-CAMELLIA128-SHA"
	//       - "ECDH-RSA-AES128-SHA"
	//       - "ECDH-ECDSA-AES128-SHA"
	//       - "SEED-SHA"
	//       - "CAMELLIA128-SHA"
	//       - "PSK-AES128-CBC-SHA"
	//       - "DES-CBC3-SHA"
	//       - "IDEA-CBC-SHA"
	//       - "ECDHE-RSA-DES-CBC3-SHA"
	//       - "ECDHE-ECDSA-DES-CBC3-SHA"
	//       - "DHE-RSA-DES-CBC3-SHA"
	//       - "DHE-DSS-DES-CBC3-SHA"
	//       - "DH-RSA-DES-CBC3-SHA"
	//       - "DH-DSS-DES-CBC3-SHAv"
	//       - "ECDH-RSA-DES-CBC3-SHA"
	//       - "ECDH-ECDSA-DES-CBC3-SHA"
	//       - "PSK-3DES-EDE-CBC-SHA"
	//       - "KRB5-IDEA-CBC-SHA"
	//       - "KRB5-DES-CBC3-SHA"
	//       - "KRB5-IDEA-CBC-MD5"
	//       - "KRB5-DES-CBC3-MD5"
	//       - "ECDHE-RSA-RC4-SHA"
	//       - "ECDHE-ECDSA-RC4-SHA"
	//       - "ECDH-RSA-RC4-SHA"
	//       - "ECDH-ECDSA-RC4-SHA"
	//       - "RC4-SHA"
	//       - "RC4-MD5"
	//       - "PSK-RC4-SHA"
	//       - "KRB5-RC4-SHA"
	//       - "KRB5-RC4-MD5"
	Ciphers []string `mandatory:"true" json:"ciphers"`
}

func (m CreateSslCipherSuiteDetails) String() string {
	return common.PointerString(m)
}
