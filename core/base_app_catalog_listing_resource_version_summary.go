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

// BaseAppCatalogListingResourceVersionSummary Listing Resource Version summary
type BaseAppCatalogListingResourceVersionSummary interface {

	// The OCID of the listing this resource version belongs to.
	GetListingId() *string

	// Date and time the listing resource version was published, in RFC3339 format.
	// Example: `2018-03-20T12:32:53.532Z`
	GetPublishedDate() *common.SDKTime

	// OCID of the listing resource.
	GetResourceId() *string

	// Resource Version.
	GetResourceVersion() *string
}

type baseappcataloglistingresourceversionsummary struct {
	JsonData        []byte
	ListingId       *string         `mandatory:"true" json:"listingId"`
	PublishedDate   *common.SDKTime `mandatory:"true" json:"publishedDate"`
	ResourceId      *string         `mandatory:"true" json:"resourceId"`
	ResourceVersion *string         `mandatory:"true" json:"resourceVersion"`
	Type            string          `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *baseappcataloglistingresourceversionsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbaseappcataloglistingresourceversionsummary baseappcataloglistingresourceversionsummary
	s := struct {
		Model Unmarshalerbaseappcataloglistingresourceversionsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ListingId = s.Model.ListingId
	m.PublishedDate = s.Model.PublishedDate
	m.ResourceId = s.Model.ResourceId
	m.ResourceVersion = s.Model.ResourceVersion
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *baseappcataloglistingresourceversionsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.Type {
	case "listingResourceVersionSummary":
		mm := AppCatalogListingResourceVersionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "listingResourceVersion":
		mm := AppCatalogListingResourceVersion{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

//GetListingId returns ListingId
func (m baseappcataloglistingresourceversionsummary) GetListingId() *string {
	return m.ListingId
}

//GetPublishedDate returns PublishedDate
func (m baseappcataloglistingresourceversionsummary) GetPublishedDate() *common.SDKTime {
	return m.PublishedDate
}

//GetResourceId returns ResourceId
func (m baseappcataloglistingresourceversionsummary) GetResourceId() *string {
	return m.ResourceId
}

//GetResourceVersion returns ResourceVersion
func (m baseappcataloglistingresourceversionsummary) GetResourceVersion() *string {
	return m.ResourceVersion
}

func (m baseappcataloglistingresourceversionsummary) String() string {
	return common.PointerString(m)
}
