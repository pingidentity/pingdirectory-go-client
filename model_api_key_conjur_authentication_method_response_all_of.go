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

// ApiKeyConjurAuthenticationMethodResponseAllOf struct for ApiKeyConjurAuthenticationMethodResponseAllOf
type ApiKeyConjurAuthenticationMethodResponseAllOf struct {
	// Name of the Conjur Authentication Method
	Id string `json:"id"`
}

// NewApiKeyConjurAuthenticationMethodResponseAllOf instantiates a new ApiKeyConjurAuthenticationMethodResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyConjurAuthenticationMethodResponseAllOf(id string) *ApiKeyConjurAuthenticationMethodResponseAllOf {
	this := ApiKeyConjurAuthenticationMethodResponseAllOf{}
	this.Id = id
	return &this
}

// NewApiKeyConjurAuthenticationMethodResponseAllOfWithDefaults instantiates a new ApiKeyConjurAuthenticationMethodResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyConjurAuthenticationMethodResponseAllOfWithDefaults() *ApiKeyConjurAuthenticationMethodResponseAllOf {
	this := ApiKeyConjurAuthenticationMethodResponseAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *ApiKeyConjurAuthenticationMethodResponseAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiKeyConjurAuthenticationMethodResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiKeyConjurAuthenticationMethodResponseAllOf) SetId(v string) {
	o.Id = v
}

func (o ApiKeyConjurAuthenticationMethodResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableApiKeyConjurAuthenticationMethodResponseAllOf struct {
	value *ApiKeyConjurAuthenticationMethodResponseAllOf
	isSet bool
}

func (v NullableApiKeyConjurAuthenticationMethodResponseAllOf) Get() *ApiKeyConjurAuthenticationMethodResponseAllOf {
	return v.value
}

func (v *NullableApiKeyConjurAuthenticationMethodResponseAllOf) Set(val *ApiKeyConjurAuthenticationMethodResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyConjurAuthenticationMethodResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyConjurAuthenticationMethodResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyConjurAuthenticationMethodResponseAllOf(val *ApiKeyConjurAuthenticationMethodResponseAllOf) *NullableApiKeyConjurAuthenticationMethodResponseAllOf {
	return &NullableApiKeyConjurAuthenticationMethodResponseAllOf{value: val, isSet: true}
}

func (v NullableApiKeyConjurAuthenticationMethodResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyConjurAuthenticationMethodResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


