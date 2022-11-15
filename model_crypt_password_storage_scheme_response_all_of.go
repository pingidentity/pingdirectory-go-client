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

// CryptPasswordStorageSchemeResponseAllOf struct for CryptPasswordStorageSchemeResponseAllOf
type CryptPasswordStorageSchemeResponseAllOf struct {
	// Name of the Password Storage Scheme
	Id string `json:"id"`
}

// NewCryptPasswordStorageSchemeResponseAllOf instantiates a new CryptPasswordStorageSchemeResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptPasswordStorageSchemeResponseAllOf(id string) *CryptPasswordStorageSchemeResponseAllOf {
	this := CryptPasswordStorageSchemeResponseAllOf{}
	this.Id = id
	return &this
}

// NewCryptPasswordStorageSchemeResponseAllOfWithDefaults instantiates a new CryptPasswordStorageSchemeResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptPasswordStorageSchemeResponseAllOfWithDefaults() *CryptPasswordStorageSchemeResponseAllOf {
	this := CryptPasswordStorageSchemeResponseAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *CryptPasswordStorageSchemeResponseAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CryptPasswordStorageSchemeResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CryptPasswordStorageSchemeResponseAllOf) SetId(v string) {
	o.Id = v
}

func (o CryptPasswordStorageSchemeResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCryptPasswordStorageSchemeResponseAllOf struct {
	value *CryptPasswordStorageSchemeResponseAllOf
	isSet bool
}

func (v NullableCryptPasswordStorageSchemeResponseAllOf) Get() *CryptPasswordStorageSchemeResponseAllOf {
	return v.value
}

func (v *NullableCryptPasswordStorageSchemeResponseAllOf) Set(val *CryptPasswordStorageSchemeResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptPasswordStorageSchemeResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptPasswordStorageSchemeResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptPasswordStorageSchemeResponseAllOf(val *CryptPasswordStorageSchemeResponseAllOf) *NullableCryptPasswordStorageSchemeResponseAllOf {
	return &NullableCryptPasswordStorageSchemeResponseAllOf{value: val, isSet: true}
}

func (v NullableCryptPasswordStorageSchemeResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptPasswordStorageSchemeResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


