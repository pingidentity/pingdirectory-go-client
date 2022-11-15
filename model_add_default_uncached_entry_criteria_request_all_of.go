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

// AddDefaultUncachedEntryCriteriaRequestAllOf struct for AddDefaultUncachedEntryCriteriaRequestAllOf
type AddDefaultUncachedEntryCriteriaRequestAllOf struct {
	// Name of the new Uncached Entry Criteria
	CriteriaName string `json:"criteriaName"`
}

// NewAddDefaultUncachedEntryCriteriaRequestAllOf instantiates a new AddDefaultUncachedEntryCriteriaRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDefaultUncachedEntryCriteriaRequestAllOf(criteriaName string) *AddDefaultUncachedEntryCriteriaRequestAllOf {
	this := AddDefaultUncachedEntryCriteriaRequestAllOf{}
	this.CriteriaName = criteriaName
	return &this
}

// NewAddDefaultUncachedEntryCriteriaRequestAllOfWithDefaults instantiates a new AddDefaultUncachedEntryCriteriaRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDefaultUncachedEntryCriteriaRequestAllOfWithDefaults() *AddDefaultUncachedEntryCriteriaRequestAllOf {
	this := AddDefaultUncachedEntryCriteriaRequestAllOf{}
	return &this
}

// GetCriteriaName returns the CriteriaName field value
func (o *AddDefaultUncachedEntryCriteriaRequestAllOf) GetCriteriaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CriteriaName
}

// GetCriteriaNameOk returns a tuple with the CriteriaName field value
// and a boolean to check if the value has been set.
func (o *AddDefaultUncachedEntryCriteriaRequestAllOf) GetCriteriaNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CriteriaName, true
}

// SetCriteriaName sets field value
func (o *AddDefaultUncachedEntryCriteriaRequestAllOf) SetCriteriaName(v string) {
	o.CriteriaName = v
}

func (o AddDefaultUncachedEntryCriteriaRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["criteriaName"] = o.CriteriaName
	}
	return json.Marshal(toSerialize)
}

type NullableAddDefaultUncachedEntryCriteriaRequestAllOf struct {
	value *AddDefaultUncachedEntryCriteriaRequestAllOf
	isSet bool
}

func (v NullableAddDefaultUncachedEntryCriteriaRequestAllOf) Get() *AddDefaultUncachedEntryCriteriaRequestAllOf {
	return v.value
}

func (v *NullableAddDefaultUncachedEntryCriteriaRequestAllOf) Set(val *AddDefaultUncachedEntryCriteriaRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDefaultUncachedEntryCriteriaRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDefaultUncachedEntryCriteriaRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDefaultUncachedEntryCriteriaRequestAllOf(val *AddDefaultUncachedEntryCriteriaRequestAllOf) *NullableAddDefaultUncachedEntryCriteriaRequestAllOf {
	return &NullableAddDefaultUncachedEntryCriteriaRequestAllOf{value: val, isSet: true}
}

func (v NullableAddDefaultUncachedEntryCriteriaRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDefaultUncachedEntryCriteriaRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


