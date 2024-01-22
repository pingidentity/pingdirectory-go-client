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

// checks if the AddAggregateIdentityMapperRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddAggregateIdentityMapperRequest{}

// AddAggregateIdentityMapperRequest struct for AddAggregateIdentityMapperRequest
type AddAggregateIdentityMapperRequest struct {
	Schemas []EnumaggregateIdentityMapperSchemaUrn `json:"schemas"`
	// The set of identity mappers that must all match the target entry. Each identity mapper must uniquely match the same target entry. If any of the identity mappers match multiple entries, if any of them match zero entries, or if any of them match different entries, then the mapping will fail.
	AllIncludedIdentityMapper []string `json:"allIncludedIdentityMapper,omitempty"`
	// The set of identity mappers that will be used to identify the target entry. At least one identity mapper must uniquely match an entry. If multiple identity mappers match entries, then they must all uniquely match the same entry. If none of the identity mappers match any entries, if any of them match multiple entries, or if any of them match different entries, then the mapping will fail.
	AnyIncludedIdentityMapper []string `json:"anyIncludedIdentityMapper,omitempty"`
	// A description for this Identity Mapper
	Description *string `json:"description,omitempty"`
	// Indicates whether the Identity Mapper is enabled for use.
	Enabled bool `json:"enabled"`
	// Name of the new Identity Mapper
	MapperName string `json:"mapperName"`
}

// NewAddAggregateIdentityMapperRequest instantiates a new AddAggregateIdentityMapperRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAggregateIdentityMapperRequest(schemas []EnumaggregateIdentityMapperSchemaUrn, enabled bool, mapperName string) *AddAggregateIdentityMapperRequest {
	this := AddAggregateIdentityMapperRequest{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.MapperName = mapperName
	return &this
}

// NewAddAggregateIdentityMapperRequestWithDefaults instantiates a new AddAggregateIdentityMapperRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAggregateIdentityMapperRequestWithDefaults() *AddAggregateIdentityMapperRequest {
	this := AddAggregateIdentityMapperRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddAggregateIdentityMapperRequest) GetSchemas() []EnumaggregateIdentityMapperSchemaUrn {
	if o == nil {
		var ret []EnumaggregateIdentityMapperSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddAggregateIdentityMapperRequest) GetSchemasOk() ([]EnumaggregateIdentityMapperSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddAggregateIdentityMapperRequest) SetSchemas(v []EnumaggregateIdentityMapperSchemaUrn) {
	o.Schemas = v
}

// GetAllIncludedIdentityMapper returns the AllIncludedIdentityMapper field value if set, zero value otherwise.
func (o *AddAggregateIdentityMapperRequest) GetAllIncludedIdentityMapper() []string {
	if o == nil || IsNil(o.AllIncludedIdentityMapper) {
		var ret []string
		return ret
	}
	return o.AllIncludedIdentityMapper
}

// GetAllIncludedIdentityMapperOk returns a tuple with the AllIncludedIdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAggregateIdentityMapperRequest) GetAllIncludedIdentityMapperOk() ([]string, bool) {
	if o == nil || IsNil(o.AllIncludedIdentityMapper) {
		return nil, false
	}
	return o.AllIncludedIdentityMapper, true
}

// HasAllIncludedIdentityMapper returns a boolean if a field has been set.
func (o *AddAggregateIdentityMapperRequest) HasAllIncludedIdentityMapper() bool {
	if o != nil && !IsNil(o.AllIncludedIdentityMapper) {
		return true
	}

	return false
}

// SetAllIncludedIdentityMapper gets a reference to the given []string and assigns it to the AllIncludedIdentityMapper field.
func (o *AddAggregateIdentityMapperRequest) SetAllIncludedIdentityMapper(v []string) {
	o.AllIncludedIdentityMapper = v
}

// GetAnyIncludedIdentityMapper returns the AnyIncludedIdentityMapper field value if set, zero value otherwise.
func (o *AddAggregateIdentityMapperRequest) GetAnyIncludedIdentityMapper() []string {
	if o == nil || IsNil(o.AnyIncludedIdentityMapper) {
		var ret []string
		return ret
	}
	return o.AnyIncludedIdentityMapper
}

// GetAnyIncludedIdentityMapperOk returns a tuple with the AnyIncludedIdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAggregateIdentityMapperRequest) GetAnyIncludedIdentityMapperOk() ([]string, bool) {
	if o == nil || IsNil(o.AnyIncludedIdentityMapper) {
		return nil, false
	}
	return o.AnyIncludedIdentityMapper, true
}

// HasAnyIncludedIdentityMapper returns a boolean if a field has been set.
func (o *AddAggregateIdentityMapperRequest) HasAnyIncludedIdentityMapper() bool {
	if o != nil && !IsNil(o.AnyIncludedIdentityMapper) {
		return true
	}

	return false
}

// SetAnyIncludedIdentityMapper gets a reference to the given []string and assigns it to the AnyIncludedIdentityMapper field.
func (o *AddAggregateIdentityMapperRequest) SetAnyIncludedIdentityMapper(v []string) {
	o.AnyIncludedIdentityMapper = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddAggregateIdentityMapperRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAggregateIdentityMapperRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddAggregateIdentityMapperRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddAggregateIdentityMapperRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddAggregateIdentityMapperRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddAggregateIdentityMapperRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddAggregateIdentityMapperRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMapperName returns the MapperName field value
func (o *AddAggregateIdentityMapperRequest) GetMapperName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MapperName
}

// GetMapperNameOk returns a tuple with the MapperName field value
// and a boolean to check if the value has been set.
func (o *AddAggregateIdentityMapperRequest) GetMapperNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapperName, true
}

// SetMapperName sets field value
func (o *AddAggregateIdentityMapperRequest) SetMapperName(v string) {
	o.MapperName = v
}

func (o AddAggregateIdentityMapperRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddAggregateIdentityMapperRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.AllIncludedIdentityMapper) {
		toSerialize["allIncludedIdentityMapper"] = o.AllIncludedIdentityMapper
	}
	if !IsNil(o.AnyIncludedIdentityMapper) {
		toSerialize["anyIncludedIdentityMapper"] = o.AnyIncludedIdentityMapper
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["mapperName"] = o.MapperName
	return toSerialize, nil
}

type NullableAddAggregateIdentityMapperRequest struct {
	value *AddAggregateIdentityMapperRequest
	isSet bool
}

func (v NullableAddAggregateIdentityMapperRequest) Get() *AddAggregateIdentityMapperRequest {
	return v.value
}

func (v *NullableAddAggregateIdentityMapperRequest) Set(val *AddAggregateIdentityMapperRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAggregateIdentityMapperRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAggregateIdentityMapperRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAggregateIdentityMapperRequest(val *AddAggregateIdentityMapperRequest) *NullableAddAggregateIdentityMapperRequest {
	return &NullableAddAggregateIdentityMapperRequest{value: val, isSet: true}
}

func (v NullableAddAggregateIdentityMapperRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAggregateIdentityMapperRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
