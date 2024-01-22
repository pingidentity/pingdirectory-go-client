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

// checks if the CustomLoggedStatsListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomLoggedStatsListResponse{}

// CustomLoggedStatsListResponse struct for CustomLoggedStatsListResponse
type CustomLoggedStatsListResponse struct {
	Schemas      []string                    `json:"schemas,omitempty"`
	TotalResults *float64                    `json:"totalResults,omitempty"`
	Resources    []CustomLoggedStatsResponse `json:"Resources,omitempty"`
}

// NewCustomLoggedStatsListResponse instantiates a new CustomLoggedStatsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomLoggedStatsListResponse() *CustomLoggedStatsListResponse {
	this := CustomLoggedStatsListResponse{}
	return &this
}

// NewCustomLoggedStatsListResponseWithDefaults instantiates a new CustomLoggedStatsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomLoggedStatsListResponseWithDefaults() *CustomLoggedStatsListResponse {
	this := CustomLoggedStatsListResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *CustomLoggedStatsListResponse) GetSchemas() []string {
	if o == nil || IsNil(o.Schemas) {
		var ret []string
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomLoggedStatsListResponse) GetSchemasOk() ([]string, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *CustomLoggedStatsListResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *CustomLoggedStatsListResponse) SetSchemas(v []string) {
	o.Schemas = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *CustomLoggedStatsListResponse) GetTotalResults() float64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomLoggedStatsListResponse) GetTotalResultsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *CustomLoggedStatsListResponse) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float64 and assigns it to the TotalResults field.
func (o *CustomLoggedStatsListResponse) SetTotalResults(v float64) {
	o.TotalResults = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *CustomLoggedStatsListResponse) GetResources() []CustomLoggedStatsResponse {
	if o == nil || IsNil(o.Resources) {
		var ret []CustomLoggedStatsResponse
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomLoggedStatsListResponse) GetResourcesOk() ([]CustomLoggedStatsResponse, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *CustomLoggedStatsListResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []CustomLoggedStatsResponse and assigns it to the Resources field.
func (o *CustomLoggedStatsListResponse) SetResources(v []CustomLoggedStatsResponse) {
	o.Resources = v
}

func (o CustomLoggedStatsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomLoggedStatsListResponse) ToMap() (map[string]interface{}, error) {
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

type NullableCustomLoggedStatsListResponse struct {
	value *CustomLoggedStatsListResponse
	isSet bool
}

func (v NullableCustomLoggedStatsListResponse) Get() *CustomLoggedStatsListResponse {
	return v.value
}

func (v *NullableCustomLoggedStatsListResponse) Set(val *CustomLoggedStatsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomLoggedStatsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomLoggedStatsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomLoggedStatsListResponse(val *CustomLoggedStatsListResponse) *NullableCustomLoggedStatsListResponse {
	return &NullableCustomLoggedStatsListResponse{value: val, isSet: true}
}

func (v NullableCustomLoggedStatsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomLoggedStatsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}