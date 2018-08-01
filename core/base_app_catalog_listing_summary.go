// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// BaseAppCatalogListingSummary A summary of a listing in the partner image catalog.
type BaseAppCatalogListingSummary interface {

	// the region free ocid of the listing resource.
	GetId() *string

	// The display name of the listing.
	GetDisplayName() *string

	// The short summary to display in list requests.
	GetSummary() *string

	// Name of the publisher who published this listing.
	GetPublisherName() *string
}

type baseappcataloglistingsummary struct {
	JsonData      []byte
	Id            *string `mandatory:"true" json:"id"`
	DisplayName   *string `mandatory:"true" json:"displayName"`
	Summary       *string `mandatory:"true" json:"summary"`
	PublisherName *string `mandatory:"false" json:"publisherName"`
	Type          string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *baseappcataloglistingsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbaseappcataloglistingsummary baseappcataloglistingsummary
	s := struct {
		Model Unmarshalerbaseappcataloglistingsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.Summary = s.Model.Summary
	m.PublisherName = s.Model.PublisherName
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *baseappcataloglistingsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.Type {
	case "listing":
		mm := AppCatalogListingSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "subscribedListing":
		mm := SubscribedAppCatalogListingSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

//GetId returns Id
func (m baseappcataloglistingsummary) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m baseappcataloglistingsummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetSummary returns Summary
func (m baseappcataloglistingsummary) GetSummary() *string {
	return m.Summary
}

//GetPublisherName returns PublisherName
func (m baseappcataloglistingsummary) GetPublisherName() *string {
	return m.PublisherName
}

func (m baseappcataloglistingsummary) String() string {
	return common.PointerString(m)
}
