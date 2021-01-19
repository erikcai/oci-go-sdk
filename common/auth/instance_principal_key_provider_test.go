// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package auth

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v33/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInstancePrincipalKeyProvider_getRegionForFederationClient(t *testing.T) {
	regionServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "phx")
	}))
	defer regionServer.Close()

	actualRegion, err := getRegionForFederationClient(&http.Client{}, regionServer.URL)

	assert.NoError(t, err)
	assert.Equal(t, common.RegionPHX, actualRegion)
}

func TestInstancePrincipalKeyProvider_getRegionForFederationClientNotFound(t *testing.T) {
	regionServer := httptest.NewServer(http.NotFoundHandler())
	defer regionServer.Close()

	_, err := getRegionForFederationClient(&http.Client{}, regionServer.URL)

	assert.Error(t, err)
}

func TestInstancePrincipalKeyProvider_getRegionForFederationClientInternalServerError(t *testing.T) {
	regionServer := httptest.NewServer(http.HandlerFunc(internalServerError))
	defer regionServer.Close()

	_, err := getRegionForFederationClient(&http.Client{}, regionServer.URL)

	assert.Error(t, err)
}

func TestInstancePrincipalKeyProvider_RegionForFederationClient(t *testing.T) {
	expectedRegion := common.StringToRegion("sea")
	keyProvider := &instancePrincipalKeyProvider{Region: expectedRegion}
	returnedRegion := keyProvider.RegionForFederationClient()
	assert.Equal(t, returnedRegion, expectedRegion)
}

func TestInstancePrincipalKeyProvider_PrivateRSAKey(t *testing.T) {
	mockFederationClient := new(mockFederationClient)
	expectedPrivateKey := new(rsa.PrivateKey)
	mockFederationClient.On("PrivateKey").Return(expectedPrivateKey, nil).Once()

	keyProvider := &instancePrincipalKeyProvider{FederationClient: mockFederationClient}

	actualPrivateKey, err := keyProvider.PrivateRSAKey()

	assert.NoError(t, err)
	assert.Equal(t, expectedPrivateKey, actualPrivateKey)
	mockFederationClient.AssertExpectations(t)
}

func TestInstancePrincipalKeyProvider_PrivateRSAKeyError(t *testing.T) {
	mockFederationClient := new(mockFederationClient)
	var nilPtr *rsa.PrivateKey
	expectedErrorMessage := "TestPrivateRSAKeyError"
	mockFederationClient.On("PrivateKey").Return(nilPtr, fmt.Errorf(expectedErrorMessage)).Once()

	keyProvider := &instancePrincipalKeyProvider{FederationClient: mockFederationClient}

	actualPrivateKey, actualError := keyProvider.PrivateRSAKey()

	assert.Nil(t, actualPrivateKey)
	assert.EqualError(t, actualError, fmt.Sprintf("failed to get private key: %s", expectedErrorMessage))
	mockFederationClient.AssertExpectations(t)
}

func TestInstancePrincipalKeyProvider_KeyID(t *testing.T) {
	mockFederationClient := new(mockFederationClient)
	mockFederationClient.On("SecurityToken").Return("TestSecurityTokenString", nil).Once()

	keyProvider := &instancePrincipalKeyProvider{FederationClient: mockFederationClient}

	actualKeyID, err := keyProvider.KeyID()

	assert.NoError(t, err)
	assert.Equal(t, "ST$TestSecurityTokenString", actualKeyID)
}

type requestVerifier struct {
	t               *testing.T
	expectedPurpose string
}

func (r requestVerifier) Do(req *http.Request) (*http.Response, error) {
	bts, _ := ioutil.ReadAll(req.Body)
	fedRequest := X509FederationDetails{}
	err := json.Unmarshal(bts, &fedRequest)
	if err != nil {
		return nil, err
	}

	assert.Equal(r.t, r.expectedPurpose, fedRequest.Purpose)

	jsonBody := fmt.Sprintf(`{"token":"%s"}`, expectedSecurityToken)
	buff := bytes.NewBufferString(jsonBody)
	return &http.Response{Body: ioutil.NopCloser(buff)}, nil

}

func TestInstancePrincipalKeyProviderCustomClient(t *testing.T) {

	modifier := func(d common.HTTPRequestDispatcher) (common.HTTPRequestDispatcher, error) {
		return requestVerifier{t, servicePrincipalTokenPurpose}, nil
	}

	provider, e := instancePrincipalConfigurationWithCertsAndPurpose(common.RegionPHX, []byte(leafCertPem), []byte(""),
		[]byte(leafCertPrivateKeyPem), [][]byte{[]byte(intermediateCertPem)}, servicePrincipalTokenPurpose, modifier)
	assert.NoError(t, e)
	_, e = provider.KeyID()
	assert.NoError(t, e)
}

func TestInstancePrincipalKeyProvider_KeyIDError(t *testing.T) {
	mockFederationClient := new(mockFederationClient)
	expectedErrorMessage := "TestSecurityTokenError"
	mockFederationClient.On("SecurityToken").Return("", fmt.Errorf(expectedErrorMessage)).Once()

	keyProvider := &instancePrincipalKeyProvider{FederationClient: mockFederationClient}

	actualKeyID, actualError := keyProvider.KeyID()

	assert.Equal(t, "", actualKeyID)
	assert.EqualError(t, actualError, fmt.Sprintf("failed to get security token: %s", expectedErrorMessage))
	mockFederationClient.AssertExpectations(t)
}

type mockFederationClient struct {
	mock.Mock
}

func (m *mockFederationClient) PrivateKey() (*rsa.PrivateKey, error) {
	args := m.Called()
	return args.Get(0).(*rsa.PrivateKey), args.Error(1)
}

func (m *mockFederationClient) SecurityToken() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *mockFederationClient) GetClaim(key string) (interface{}, error) {
	args := m.Called(key)
	return args.Get(0), args.Error(1)
}
