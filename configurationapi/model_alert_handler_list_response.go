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

// checks if the AlertHandlerListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertHandlerListResponse{}

// AlertHandlerListResponse struct for AlertHandlerListResponse
type AlertHandlerListResponse struct {
	Schemas      []string                     `json:"schemas,omitempty"`
	TotalResults *float64                     `json:"totalResults,omitempty"`
	Resources    []GetAlertHandler200Response `json:"Resources,omitempty"`
}

// NewAlertHandlerListResponse instantiates a new AlertHandlerListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertHandlerListResponse() *AlertHandlerListResponse {
	this := AlertHandlerListResponse{}
	return &this
}

// NewAlertHandlerListResponseWithDefaults instantiates a new AlertHandlerListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertHandlerListResponseWithDefaults() *AlertHandlerListResponse {
	this := AlertHandlerListResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AlertHandlerListResponse) GetSchemas() []string {
	if o == nil || IsNil(o.Schemas) {
		var ret []string
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertHandlerListResponse) GetSchemasOk() ([]string, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AlertHandlerListResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *AlertHandlerListResponse) SetSchemas(v []string) {
	o.Schemas = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *AlertHandlerListResponse) GetTotalResults() float64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertHandlerListResponse) GetTotalResultsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *AlertHandlerListResponse) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float64 and assigns it to the TotalResults field.
func (o *AlertHandlerListResponse) SetTotalResults(v float64) {
	o.TotalResults = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *AlertHandlerListResponse) GetResources() []GetAlertHandler200Response {
	if o == nil || IsNil(o.Resources) {
		var ret []GetAlertHandler200Response
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertHandlerListResponse) GetResourcesOk() ([]GetAlertHandler200Response, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *AlertHandlerListResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []GetAlertHandler200Response and assigns it to the Resources field.
func (o *AlertHandlerListResponse) SetResources(v []GetAlertHandler200Response) {
	o.Resources = v
}

func (o AlertHandlerListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertHandlerListResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAlertHandlerListResponse struct {
	value *AlertHandlerListResponse
	isSet bool
}

func (v NullableAlertHandlerListResponse) Get() *AlertHandlerListResponse {
	return v.value
}

func (v *NullableAlertHandlerListResponse) Set(val *AlertHandlerListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertHandlerListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertHandlerListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertHandlerListResponse(val *AlertHandlerListResponse) *NullableAlertHandlerListResponse {
	return &NullableAlertHandlerListResponse{value: val, isSet: true}
}

func (v NullableAlertHandlerListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertHandlerListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
