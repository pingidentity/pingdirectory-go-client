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

// checks if the VelocityContextProviderListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VelocityContextProviderListResponse{}

// VelocityContextProviderListResponse struct for VelocityContextProviderListResponse
type VelocityContextProviderListResponse struct {
	Schemas      []string                                `json:"schemas,omitempty"`
	TotalResults *float64                                `json:"totalResults,omitempty"`
	Resources    []GetVelocityContextProvider200Response `json:"Resources,omitempty"`
}

// NewVelocityContextProviderListResponse instantiates a new VelocityContextProviderListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVelocityContextProviderListResponse() *VelocityContextProviderListResponse {
	this := VelocityContextProviderListResponse{}
	return &this
}

// NewVelocityContextProviderListResponseWithDefaults instantiates a new VelocityContextProviderListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVelocityContextProviderListResponseWithDefaults() *VelocityContextProviderListResponse {
	this := VelocityContextProviderListResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *VelocityContextProviderListResponse) GetSchemas() []string {
	if o == nil || IsNil(o.Schemas) {
		var ret []string
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityContextProviderListResponse) GetSchemasOk() ([]string, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *VelocityContextProviderListResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *VelocityContextProviderListResponse) SetSchemas(v []string) {
	o.Schemas = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *VelocityContextProviderListResponse) GetTotalResults() float64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityContextProviderListResponse) GetTotalResultsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *VelocityContextProviderListResponse) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float64 and assigns it to the TotalResults field.
func (o *VelocityContextProviderListResponse) SetTotalResults(v float64) {
	o.TotalResults = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *VelocityContextProviderListResponse) GetResources() []GetVelocityContextProvider200Response {
	if o == nil || IsNil(o.Resources) {
		var ret []GetVelocityContextProvider200Response
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityContextProviderListResponse) GetResourcesOk() ([]GetVelocityContextProvider200Response, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *VelocityContextProviderListResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []GetVelocityContextProvider200Response and assigns it to the Resources field.
func (o *VelocityContextProviderListResponse) SetResources(v []GetVelocityContextProvider200Response) {
	o.Resources = v
}

func (o VelocityContextProviderListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VelocityContextProviderListResponse) ToMap() (map[string]interface{}, error) {
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

type NullableVelocityContextProviderListResponse struct {
	value *VelocityContextProviderListResponse
	isSet bool
}

func (v NullableVelocityContextProviderListResponse) Get() *VelocityContextProviderListResponse {
	return v.value
}

func (v *NullableVelocityContextProviderListResponse) Set(val *VelocityContextProviderListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVelocityContextProviderListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVelocityContextProviderListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVelocityContextProviderListResponse(val *VelocityContextProviderListResponse) *NullableVelocityContextProviderListResponse {
	return &NullableVelocityContextProviderListResponse{value: val, isSet: true}
}

func (v NullableVelocityContextProviderListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVelocityContextProviderListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
