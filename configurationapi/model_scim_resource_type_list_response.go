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

// checks if the ScimResourceTypeListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimResourceTypeListResponse{}

// ScimResourceTypeListResponse struct for ScimResourceTypeListResponse
type ScimResourceTypeListResponse struct {
	Schemas      []string                                     `json:"schemas,omitempty"`
	TotalResults *float64                                     `json:"totalResults,omitempty"`
	Resources    []ScimResourceTypeListResponseResourcesInner `json:"Resources,omitempty"`
}

// NewScimResourceTypeListResponse instantiates a new ScimResourceTypeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimResourceTypeListResponse() *ScimResourceTypeListResponse {
	this := ScimResourceTypeListResponse{}
	return &this
}

// NewScimResourceTypeListResponseWithDefaults instantiates a new ScimResourceTypeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimResourceTypeListResponseWithDefaults() *ScimResourceTypeListResponse {
	this := ScimResourceTypeListResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ScimResourceTypeListResponse) GetSchemas() []string {
	if o == nil || IsNil(o.Schemas) {
		var ret []string
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimResourceTypeListResponse) GetSchemasOk() ([]string, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ScimResourceTypeListResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *ScimResourceTypeListResponse) SetSchemas(v []string) {
	o.Schemas = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ScimResourceTypeListResponse) GetTotalResults() float64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimResourceTypeListResponse) GetTotalResultsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ScimResourceTypeListResponse) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float64 and assigns it to the TotalResults field.
func (o *ScimResourceTypeListResponse) SetTotalResults(v float64) {
	o.TotalResults = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ScimResourceTypeListResponse) GetResources() []ScimResourceTypeListResponseResourcesInner {
	if o == nil || IsNil(o.Resources) {
		var ret []ScimResourceTypeListResponseResourcesInner
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimResourceTypeListResponse) GetResourcesOk() ([]ScimResourceTypeListResponseResourcesInner, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ScimResourceTypeListResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ScimResourceTypeListResponseResourcesInner and assigns it to the Resources field.
func (o *ScimResourceTypeListResponse) SetResources(v []ScimResourceTypeListResponseResourcesInner) {
	o.Resources = v
}

func (o ScimResourceTypeListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimResourceTypeListResponse) ToMap() (map[string]interface{}, error) {
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

type NullableScimResourceTypeListResponse struct {
	value *ScimResourceTypeListResponse
	isSet bool
}

func (v NullableScimResourceTypeListResponse) Get() *ScimResourceTypeListResponse {
	return v.value
}

func (v *NullableScimResourceTypeListResponse) Set(val *ScimResourceTypeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScimResourceTypeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScimResourceTypeListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimResourceTypeListResponse(val *ScimResourceTypeListResponse) *NullableScimResourceTypeListResponse {
	return &NullableScimResourceTypeListResponse{value: val, isSet: true}
}

func (v NullableScimResourceTypeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimResourceTypeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
