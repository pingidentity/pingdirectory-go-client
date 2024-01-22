/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the KeyManagerProviderListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyManagerProviderListResponse{}

// KeyManagerProviderListResponse struct for KeyManagerProviderListResponse
type KeyManagerProviderListResponse struct {
	Schemas      []string                                       `json:"schemas,omitempty"`
	TotalResults *float64                                       `json:"totalResults,omitempty"`
	Resources    []KeyManagerProviderListResponseResourcesInner `json:"Resources,omitempty"`
}

// NewKeyManagerProviderListResponse instantiates a new KeyManagerProviderListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyManagerProviderListResponse() *KeyManagerProviderListResponse {
	this := KeyManagerProviderListResponse{}
	return &this
}

// NewKeyManagerProviderListResponseWithDefaults instantiates a new KeyManagerProviderListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyManagerProviderListResponseWithDefaults() *KeyManagerProviderListResponse {
	this := KeyManagerProviderListResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *KeyManagerProviderListResponse) GetSchemas() []string {
	if o == nil || IsNil(o.Schemas) {
		var ret []string
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyManagerProviderListResponse) GetSchemasOk() ([]string, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *KeyManagerProviderListResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *KeyManagerProviderListResponse) SetSchemas(v []string) {
	o.Schemas = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *KeyManagerProviderListResponse) GetTotalResults() float64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyManagerProviderListResponse) GetTotalResultsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *KeyManagerProviderListResponse) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float64 and assigns it to the TotalResults field.
func (o *KeyManagerProviderListResponse) SetTotalResults(v float64) {
	o.TotalResults = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *KeyManagerProviderListResponse) GetResources() []KeyManagerProviderListResponseResourcesInner {
	if o == nil || IsNil(o.Resources) {
		var ret []KeyManagerProviderListResponseResourcesInner
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyManagerProviderListResponse) GetResourcesOk() ([]KeyManagerProviderListResponseResourcesInner, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *KeyManagerProviderListResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []KeyManagerProviderListResponseResourcesInner and assigns it to the Resources field.
func (o *KeyManagerProviderListResponse) SetResources(v []KeyManagerProviderListResponseResourcesInner) {
	o.Resources = v
}

func (o KeyManagerProviderListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyManagerProviderListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.Resources) {
		toSerialize["Resources"] = o.Resources
	}
	return toSerialize, nil
}

type NullableKeyManagerProviderListResponse struct {
	value *KeyManagerProviderListResponse
	isSet bool
}

func (v NullableKeyManagerProviderListResponse) Get() *KeyManagerProviderListResponse {
	return v.value
}

func (v *NullableKeyManagerProviderListResponse) Set(val *KeyManagerProviderListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyManagerProviderListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyManagerProviderListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyManagerProviderListResponse(val *KeyManagerProviderListResponse) *NullableKeyManagerProviderListResponse {
	return &NullableKeyManagerProviderListResponse{value: val, isSet: true}
}

func (v NullableKeyManagerProviderListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyManagerProviderListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}