/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AddLastAccessTimeUncachedEntryCriteriaRequest struct for AddLastAccessTimeUncachedEntryCriteriaRequest
type AddLastAccessTimeUncachedEntryCriteriaRequest struct {
	// Name of the new Uncached Entry Criteria
	CriteriaName string `json:"criteriaName"`
	Schemas []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn `json:"schemas"`
	// Specifies the maximum length of time that has passed since an entry was last accessed that it should still be included in the id2entry database. Entries that have not been accessed in more than this length of time may be written into the uncached-id2entry database.
	AccessTimeThreshold string `json:"accessTimeThreshold"`
	// A description for this Uncached Entry Criteria
	Description *string `json:"description,omitempty"`
	// Indicates whether this Uncached Entry Criteria is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewAddLastAccessTimeUncachedEntryCriteriaRequest instantiates a new AddLastAccessTimeUncachedEntryCriteriaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLastAccessTimeUncachedEntryCriteriaRequest(criteriaName string, schemas []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn, accessTimeThreshold string, enabled bool) *AddLastAccessTimeUncachedEntryCriteriaRequest {
	this := AddLastAccessTimeUncachedEntryCriteriaRequest{}
	this.CriteriaName = criteriaName
	this.Schemas = schemas
	this.AccessTimeThreshold = accessTimeThreshold
	this.Enabled = enabled
	return &this
}

// NewAddLastAccessTimeUncachedEntryCriteriaRequestWithDefaults instantiates a new AddLastAccessTimeUncachedEntryCriteriaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLastAccessTimeUncachedEntryCriteriaRequestWithDefaults() *AddLastAccessTimeUncachedEntryCriteriaRequest {
	this := AddLastAccessTimeUncachedEntryCriteriaRequest{}
	return &this
}

// GetCriteriaName returns the CriteriaName field value
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetCriteriaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CriteriaName
}

// GetCriteriaNameOk returns a tuple with the CriteriaName field value
// and a boolean to check if the value has been set.
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetCriteriaNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CriteriaName, true
}

// SetCriteriaName sets field value
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) SetCriteriaName(v string) {
	o.CriteriaName = v
}

// GetSchemas returns the Schemas field value
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetSchemas() []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetSchemasOk() ([]EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) SetSchemas(v []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetAccessTimeThreshold returns the AccessTimeThreshold field value
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetAccessTimeThreshold() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessTimeThreshold
}

// GetAccessTimeThresholdOk returns a tuple with the AccessTimeThreshold field value
// and a boolean to check if the value has been set.
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetAccessTimeThresholdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessTimeThreshold, true
}

// SetAccessTimeThreshold sets field value
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) SetAccessTimeThreshold(v string) {
	o.AccessTimeThreshold = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddLastAccessTimeUncachedEntryCriteriaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["criteriaName"] = o.CriteriaName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["accessTimeThreshold"] = o.AccessTimeThreshold
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddLastAccessTimeUncachedEntryCriteriaRequest struct {
	value *AddLastAccessTimeUncachedEntryCriteriaRequest
	isSet bool
}

func (v NullableAddLastAccessTimeUncachedEntryCriteriaRequest) Get() *AddLastAccessTimeUncachedEntryCriteriaRequest {
	return v.value
}

func (v *NullableAddLastAccessTimeUncachedEntryCriteriaRequest) Set(val *AddLastAccessTimeUncachedEntryCriteriaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLastAccessTimeUncachedEntryCriteriaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLastAccessTimeUncachedEntryCriteriaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLastAccessTimeUncachedEntryCriteriaRequest(val *AddLastAccessTimeUncachedEntryCriteriaRequest) *NullableAddLastAccessTimeUncachedEntryCriteriaRequest {
	return &NullableAddLastAccessTimeUncachedEntryCriteriaRequest{value: val, isSet: true}
}

func (v NullableAddLastAccessTimeUncachedEntryCriteriaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLastAccessTimeUncachedEntryCriteriaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


