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

// checks if the HttpServletCrossOriginPolicyListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpServletCrossOriginPolicyListResponse{}

// HttpServletCrossOriginPolicyListResponse struct for HttpServletCrossOriginPolicyListResponse
type HttpServletCrossOriginPolicyListResponse struct {
	Schemas      []string                               `json:"schemas,omitempty"`
	TotalResults *float64                               `json:"totalResults,omitempty"`
	Resources    []HttpServletCrossOriginPolicyResponse `json:"Resources,omitempty"`
}

// NewHttpServletCrossOriginPolicyListResponse instantiates a new HttpServletCrossOriginPolicyListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpServletCrossOriginPolicyListResponse() *HttpServletCrossOriginPolicyListResponse {
	this := HttpServletCrossOriginPolicyListResponse{}
	return &this
}

// NewHttpServletCrossOriginPolicyListResponseWithDefaults instantiates a new HttpServletCrossOriginPolicyListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpServletCrossOriginPolicyListResponseWithDefaults() *HttpServletCrossOriginPolicyListResponse {
	this := HttpServletCrossOriginPolicyListResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *HttpServletCrossOriginPolicyListResponse) GetSchemas() []string {
	if o == nil || IsNil(o.Schemas) {
		var ret []string
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpServletCrossOriginPolicyListResponse) GetSchemasOk() ([]string, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *HttpServletCrossOriginPolicyListResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *HttpServletCrossOriginPolicyListResponse) SetSchemas(v []string) {
	o.Schemas = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *HttpServletCrossOriginPolicyListResponse) GetTotalResults() float64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpServletCrossOriginPolicyListResponse) GetTotalResultsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *HttpServletCrossOriginPolicyListResponse) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float64 and assigns it to the TotalResults field.
func (o *HttpServletCrossOriginPolicyListResponse) SetTotalResults(v float64) {
	o.TotalResults = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *HttpServletCrossOriginPolicyListResponse) GetResources() []HttpServletCrossOriginPolicyResponse {
	if o == nil || IsNil(o.Resources) {
		var ret []HttpServletCrossOriginPolicyResponse
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpServletCrossOriginPolicyListResponse) GetResourcesOk() ([]HttpServletCrossOriginPolicyResponse, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *HttpServletCrossOriginPolicyListResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []HttpServletCrossOriginPolicyResponse and assigns it to the Resources field.
func (o *HttpServletCrossOriginPolicyListResponse) SetResources(v []HttpServletCrossOriginPolicyResponse) {
	o.Resources = v
}

func (o HttpServletCrossOriginPolicyListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpServletCrossOriginPolicyListResponse) ToMap() (map[string]interface{}, error) {
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

type NullableHttpServletCrossOriginPolicyListResponse struct {
	value *HttpServletCrossOriginPolicyListResponse
	isSet bool
}

func (v NullableHttpServletCrossOriginPolicyListResponse) Get() *HttpServletCrossOriginPolicyListResponse {
	return v.value
}

func (v *NullableHttpServletCrossOriginPolicyListResponse) Set(val *HttpServletCrossOriginPolicyListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpServletCrossOriginPolicyListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpServletCrossOriginPolicyListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpServletCrossOriginPolicyListResponse(val *HttpServletCrossOriginPolicyListResponse) *NullableHttpServletCrossOriginPolicyListResponse {
	return &NullableHttpServletCrossOriginPolicyListResponse{value: val, isSet: true}
}

func (v NullableHttpServletCrossOriginPolicyListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpServletCrossOriginPolicyListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}