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

// checks if the ScimSubattributeListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimSubattributeListResponse{}

// ScimSubattributeListResponse struct for ScimSubattributeListResponse
type ScimSubattributeListResponse struct {
	Schemas      []string                   `json:"schemas,omitempty"`
	TotalResults *float64                   `json:"totalResults,omitempty"`
	Resources    []ScimSubattributeResponse `json:"Resources,omitempty"`
}

// NewScimSubattributeListResponse instantiates a new ScimSubattributeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimSubattributeListResponse() *ScimSubattributeListResponse {
	this := ScimSubattributeListResponse{}
	return &this
}

// NewScimSubattributeListResponseWithDefaults instantiates a new ScimSubattributeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimSubattributeListResponseWithDefaults() *ScimSubattributeListResponse {
	this := ScimSubattributeListResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ScimSubattributeListResponse) GetSchemas() []string {
	if o == nil || IsNil(o.Schemas) {
		var ret []string
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSubattributeListResponse) GetSchemasOk() ([]string, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ScimSubattributeListResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *ScimSubattributeListResponse) SetSchemas(v []string) {
	o.Schemas = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ScimSubattributeListResponse) GetTotalResults() float64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSubattributeListResponse) GetTotalResultsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ScimSubattributeListResponse) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float64 and assigns it to the TotalResults field.
func (o *ScimSubattributeListResponse) SetTotalResults(v float64) {
	o.TotalResults = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ScimSubattributeListResponse) GetResources() []ScimSubattributeResponse {
	if o == nil || IsNil(o.Resources) {
		var ret []ScimSubattributeResponse
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSubattributeListResponse) GetResourcesOk() ([]ScimSubattributeResponse, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ScimSubattributeListResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ScimSubattributeResponse and assigns it to the Resources field.
func (o *ScimSubattributeListResponse) SetResources(v []ScimSubattributeResponse) {
	o.Resources = v
}

func (o ScimSubattributeListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimSubattributeListResponse) ToMap() (map[string]interface{}, error) {
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

type NullableScimSubattributeListResponse struct {
	value *ScimSubattributeListResponse
	isSet bool
}

func (v NullableScimSubattributeListResponse) Get() *ScimSubattributeListResponse {
	return v.value
}

func (v *NullableScimSubattributeListResponse) Set(val *ScimSubattributeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScimSubattributeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScimSubattributeListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimSubattributeListResponse(val *ScimSubattributeListResponse) *NullableScimSubattributeListResponse {
	return &NullableScimSubattributeListResponse{value: val, isSet: true}
}

func (v NullableScimSubattributeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimSubattributeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}