/*
PingData Location Config - OpenAPI 3.0

This is the PingData Configuration API for the Location config object

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LocationResponseAllOf struct for LocationResponseAllOf
type LocationResponseAllOf struct {
	// Name of the Location
	Id string `json:"id"`
}

// NewLocationResponseAllOf instantiates a new LocationResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationResponseAllOf(id string) *LocationResponseAllOf {
	this := LocationResponseAllOf{}
	this.Id = id
	return &this
}

// NewLocationResponseAllOfWithDefaults instantiates a new LocationResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationResponseAllOfWithDefaults() *LocationResponseAllOf {
	this := LocationResponseAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *LocationResponseAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LocationResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LocationResponseAllOf) SetId(v string) {
	o.Id = v
}

func (o LocationResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableLocationResponseAllOf struct {
	value *LocationResponseAllOf
	isSet bool
}

func (v NullableLocationResponseAllOf) Get() *LocationResponseAllOf {
	return v.value
}

func (v *NullableLocationResponseAllOf) Set(val *LocationResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationResponseAllOf(val *LocationResponseAllOf) *NullableLocationResponseAllOf {
	return &NullableLocationResponseAllOf{value: val, isSet: true}
}

func (v NullableLocationResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


