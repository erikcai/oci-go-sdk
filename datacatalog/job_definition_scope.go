// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DataCatalog API
//
// A description of the DataCatalog API
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// JobDefinitionScope Defines the rules or criteria based on which the scope for job definition is circumscribed.
type JobDefinitionScope struct {

	// Name of the folder or schema for this metadata harvest
	FolderName *string `mandatory:"false" json:"folderName"`

	// Name of the entity for this metadata harvest
	EntityName *string `mandatory:"false" json:"entityName"`

	// Filter rules with regular expression to specify folder names for this metadata harvest
	FolderNameFilter *string `mandatory:"false" json:"folderNameFilter"`

	// Filter rules with regular expression to specify entity names for this metadata harvest
	EntityNameFilter *string `mandatory:"false" json:"entityNameFilter"`

	// A map of maps which contains the properties which are specific to the job type. Each job type
	// definition may define it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// job definitions have required properties within the "default" category.
	// Example: `{"properties": { "default": { "host": "host1", "port": "1521", "database": "orcl"}}}`
	ConfigurationProperties map[string]map[string]string `mandatory:"false" json:"configurationProperties"`

	// Specify if sample data to be extracted as part of this harvest
	IsSampleDataExtracted *bool `mandatory:"false" json:"isSampleDataExtracted"`

	// Specify the sample data size in MB, specified as number of rows, for this metadata harvest
	SampleDataSizeInMBs *int `mandatory:"false" json:"sampleDataSizeInMBs"`
}

func (m JobDefinitionScope) String() string {
	return common.PointerString(m)
}
