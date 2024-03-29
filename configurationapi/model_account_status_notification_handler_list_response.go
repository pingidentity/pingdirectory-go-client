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

// checks if the AccountStatusNotificationHandlerListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountStatusNotificationHandlerListResponse{}

// AccountStatusNotificationHandlerListResponse struct for AccountStatusNotificationHandlerListResponse
type AccountStatusNotificationHandlerListResponse struct {
	Schemas      []string                                         `json:"schemas,omitempty"`
	TotalResults *float64                                         `json:"totalResults,omitempty"`
	Resources    []AddAccountStatusNotificationHandler200Response `json:"Resources,omitempty"`
}

// NewAccountStatusNotificationHandlerListResponse instantiates a new AccountStatusNotificationHandlerListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountStatusNotificationHandlerListResponse() *AccountStatusNotificationHandlerListResponse {
	this := AccountStatusNotificationHandlerListResponse{}
	return &this
}

// NewAccountStatusNotificationHandlerListResponseWithDefaults instantiates a new AccountStatusNotificationHandlerListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountStatusNotificationHandlerListResponseWithDefaults() *AccountStatusNotificationHandlerListResponse {
	this := AccountStatusNotificationHandlerListResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AccountStatusNotificationHandlerListResponse) GetSchemas() []string {
	if o == nil || IsNil(o.Schemas) {
		var ret []string
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusNotificationHandlerListResponse) GetSchemasOk() ([]string, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AccountStatusNotificationHandlerListResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *AccountStatusNotificationHandlerListResponse) SetSchemas(v []string) {
	o.Schemas = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *AccountStatusNotificationHandlerListResponse) GetTotalResults() float64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusNotificationHandlerListResponse) GetTotalResultsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *AccountStatusNotificationHandlerListResponse) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float64 and assigns it to the TotalResults field.
func (o *AccountStatusNotificationHandlerListResponse) SetTotalResults(v float64) {
	o.TotalResults = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *AccountStatusNotificationHandlerListResponse) GetResources() []AddAccountStatusNotificationHandler200Response {
	if o == nil || IsNil(o.Resources) {
		var ret []AddAccountStatusNotificationHandler200Response
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusNotificationHandlerListResponse) GetResourcesOk() ([]AddAccountStatusNotificationHandler200Response, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *AccountStatusNotificationHandlerListResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []AddAccountStatusNotificationHandler200Response and assigns it to the Resources field.
func (o *AccountStatusNotificationHandlerListResponse) SetResources(v []AddAccountStatusNotificationHandler200Response) {
	o.Resources = v
}

func (o AccountStatusNotificationHandlerListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountStatusNotificationHandlerListResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAccountStatusNotificationHandlerListResponse struct {
	value *AccountStatusNotificationHandlerListResponse
	isSet bool
}

func (v NullableAccountStatusNotificationHandlerListResponse) Get() *AccountStatusNotificationHandlerListResponse {
	return v.value
}

func (v *NullableAccountStatusNotificationHandlerListResponse) Set(val *AccountStatusNotificationHandlerListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountStatusNotificationHandlerListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountStatusNotificationHandlerListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountStatusNotificationHandlerListResponse(val *AccountStatusNotificationHandlerListResponse) *NullableAccountStatusNotificationHandlerListResponse {
	return &NullableAccountStatusNotificationHandlerListResponse{value: val, isSet: true}
}

func (v NullableAccountStatusNotificationHandlerListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountStatusNotificationHandlerListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
