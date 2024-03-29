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

// checks if the AddLocationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddLocationRequest{}

// AddLocationRequest struct for AddLocationRequest
type AddLocationRequest struct {
	Schemas []EnumlocationSchemaUrn `json:"schemas,omitempty"`
	// A description for this Location
	Description *string `json:"description,omitempty"`
	// Name of the new Location
	LocationName string `json:"locationName"`
}

// NewAddLocationRequest instantiates a new AddLocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLocationRequest(locationName string) *AddLocationRequest {
	this := AddLocationRequest{}
	this.LocationName = locationName
	return &this
}

// NewAddLocationRequestWithDefaults instantiates a new AddLocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLocationRequestWithDefaults() *AddLocationRequest {
	this := AddLocationRequest{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddLocationRequest) GetSchemas() []EnumlocationSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumlocationSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLocationRequest) GetSchemasOk() ([]EnumlocationSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddLocationRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumlocationSchemaUrn and assigns it to the Schemas field.
func (o *AddLocationRequest) SetSchemas(v []EnumlocationSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddLocationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLocationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddLocationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddLocationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetLocationName returns the LocationName field value
func (o *AddLocationRequest) GetLocationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocationName
}

// GetLocationNameOk returns a tuple with the LocationName field value
// and a boolean to check if the value has been set.
func (o *AddLocationRequest) GetLocationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationName, true
}

// SetLocationName sets field value
func (o *AddLocationRequest) SetLocationName(v string) {
	o.LocationName = v
}

func (o AddLocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddLocationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["locationName"] = o.LocationName
	return toSerialize, nil
}

type NullableAddLocationRequest struct {
	value *AddLocationRequest
	isSet bool
}

func (v NullableAddLocationRequest) Get() *AddLocationRequest {
	return v.value
}

func (v *NullableAddLocationRequest) Set(val *AddLocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLocationRequest(val *AddLocationRequest) *NullableAddLocationRequest {
	return &NullableAddLocationRequest{value: val, isSet: true}
}

func (v NullableAddLocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
