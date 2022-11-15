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

// AddFifoEntryCacheRequestAllOf struct for AddFifoEntryCacheRequestAllOf
type AddFifoEntryCacheRequestAllOf struct {
	// Name of the new Entry Cache
	CacheName string `json:"cacheName"`
}

// NewAddFifoEntryCacheRequestAllOf instantiates a new AddFifoEntryCacheRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFifoEntryCacheRequestAllOf(cacheName string) *AddFifoEntryCacheRequestAllOf {
	this := AddFifoEntryCacheRequestAllOf{}
	this.CacheName = cacheName
	return &this
}

// NewAddFifoEntryCacheRequestAllOfWithDefaults instantiates a new AddFifoEntryCacheRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFifoEntryCacheRequestAllOfWithDefaults() *AddFifoEntryCacheRequestAllOf {
	this := AddFifoEntryCacheRequestAllOf{}
	return &this
}

// GetCacheName returns the CacheName field value
func (o *AddFifoEntryCacheRequestAllOf) GetCacheName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CacheName
}

// GetCacheNameOk returns a tuple with the CacheName field value
// and a boolean to check if the value has been set.
func (o *AddFifoEntryCacheRequestAllOf) GetCacheNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CacheName, true
}

// SetCacheName sets field value
func (o *AddFifoEntryCacheRequestAllOf) SetCacheName(v string) {
	o.CacheName = v
}

func (o AddFifoEntryCacheRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cacheName"] = o.CacheName
	}
	return json.Marshal(toSerialize)
}

type NullableAddFifoEntryCacheRequestAllOf struct {
	value *AddFifoEntryCacheRequestAllOf
	isSet bool
}

func (v NullableAddFifoEntryCacheRequestAllOf) Get() *AddFifoEntryCacheRequestAllOf {
	return v.value
}

func (v *NullableAddFifoEntryCacheRequestAllOf) Set(val *AddFifoEntryCacheRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFifoEntryCacheRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFifoEntryCacheRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFifoEntryCacheRequestAllOf(val *AddFifoEntryCacheRequestAllOf) *NullableAddFifoEntryCacheRequestAllOf {
	return &NullableAddFifoEntryCacheRequestAllOf{value: val, isSet: true}
}

func (v NullableAddFifoEntryCacheRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFifoEntryCacheRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


