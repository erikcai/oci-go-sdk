// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.

package auth

import (
	"crypto/rsa"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
)

type servicePrincipalKeyProvider struct {
	federationClient federationClient
}

func newServicePrincipalKeyProvider(tenancyID, region string, cert, key []byte, intermediates [][]byte, passphrase []byte) (provider *servicePrincipalKeyProvider, err error) {
	leafCertificateRetriever := newStaticX509CertificateRetriever(cert, key, passphrase)

	var intermediateCertificateRetrievers []x509CertificateRetriever
	for _, intermediate := range intermediates {
		intermediateCertificateRetrievers =
			append(intermediateCertificateRetrievers, newStaticX509CertificateRetriever(intermediate, key, passphrase))
	}

	federationClient := newX509FederationClient(
		common.Region(region), tenancyID, leafCertificateRetriever, intermediateCertificateRetrievers, true)

	provider = &servicePrincipalKeyProvider{federationClient: federationClient}
	return
}

func (p *servicePrincipalKeyProvider) PrivateRSAKey() (privateKey *rsa.PrivateKey, err error) {
	if privateKey, err = p.federationClient.PrivateKey(); err != nil {
		err = fmt.Errorf("failed to get private key: %s", err.Error())
		return nil, err
	}
	return privateKey, nil
}

func (p *servicePrincipalKeyProvider) KeyID() (string, error) {
	var securityToken string
	var err error
	if securityToken, err = p.federationClient.SecurityToken(); err != nil {
		return "", fmt.Errorf("failed to get security token: %s", err.Error())
	}

	return fmt.Sprintf("ST$%s", securityToken), nil
}


type servicePrincipalConfigurationProvider struct {
	keyProvider *servicePrincipalKeyProvider
	tenancyID, region      string
}

// NewServicePrincipalConfigurationProvider will create a new service principal configuration provider
func NewServicePrincipalConfigurationProvider(tenancyID, region string, cert, key []byte, intermediates [][]byte, passphrase []byte) (common.ConfigurationProvider, error) {
	var err error
	var keyProvider *servicePrincipalKeyProvider
	if keyProvider, err = newServicePrincipalKeyProvider(tenancyID, region, cert, key, intermediates, passphrase); err != nil {
		return nil, fmt.Errorf("failed to create a new key provider: %s", err.Error())
	}
	return servicePrincipalConfigurationProvider{keyProvider: keyProvider, region: region, tenancyID:tenancyID}, nil
}

func (p servicePrincipalConfigurationProvider) PrivateRSAKey() (*rsa.PrivateKey, error) {
	return p.keyProvider.PrivateRSAKey()
}

func (p servicePrincipalConfigurationProvider) KeyID() (string, error) {
	return p.keyProvider.KeyID()
}

func (p servicePrincipalConfigurationProvider) TenancyOCID() (string, error) {
	return p.tenancyID , nil
}

func (p servicePrincipalConfigurationProvider) UserOCID() (string, error) {
	return "", nil
}

func (p servicePrincipalConfigurationProvider) KeyFingerprint() (string, error) {
	return "", nil
}

func (p servicePrincipalConfigurationProvider) Region() (string, error) {
	return p.region, nil
}
