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

// checks if the CustomSynchronizationProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomSynchronizationProviderResponse{}

// CustomSynchronizationProviderResponse struct for CustomSynchronizationProviderResponse
type CustomSynchronizationProviderResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumcustomSynchronizationProviderSchemaUrn       `json:"schemas"`
	// Name of the Synchronization Provider
	Id string `json:"id"`
	// A description for this Synchronization Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether the Synchronization Provider is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewCustomSynchronizationProviderResponse instantiates a new CustomSynchronizationProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomSynchronizationProviderResponse(schemas []EnumcustomSynchronizationProviderSchemaUrn, id string, enabled bool) *CustomSynchronizationProviderResponse {
	this := CustomSynchronizationProviderResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewCustomSynchronizationProviderResponseWithDefaults instantiates a new CustomSynchronizationProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomSynchronizationProviderResponseWithDefaults() *CustomSynchronizationProviderResponse {
	this := CustomSynchronizationProviderResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CustomSynchronizationProviderResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomSynchronizationProviderResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CustomSynchronizationProviderResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *CustomSynchronizationProviderResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *CustomSynchronizationProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomSynchronizationProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *CustomSynchronizationProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *CustomSynchronizationProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *CustomSynchronizationProviderResponse) GetSchemas() []EnumcustomSynchronizationProviderSchemaUrn {
	if o == nil {
		var ret []EnumcustomSynchronizationProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *CustomSynchronizationProviderResponse) GetSchemasOk() ([]EnumcustomSynchronizationProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *CustomSynchronizationProviderResponse) SetSchemas(v []EnumcustomSynchronizationProviderSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *CustomSynchronizationProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomSynchronizationProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomSynchronizationProviderResponse) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomSynchronizationProviderResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomSynchronizationProviderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomSynchronizationProviderResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomSynchronizationProviderResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *CustomSynchronizationProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CustomSynchronizationProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CustomSynchronizationProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o CustomSynchronizationProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomSynchronizationProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableCustomSynchronizationProviderResponse struct {
	value *CustomSynchronizationProviderResponse
	isSet bool
}

func (v NullableCustomSynchronizationProviderResponse) Get() *CustomSynchronizationProviderResponse {
	return v.value
}

func (v *NullableCustomSynchronizationProviderResponse) Set(val *CustomSynchronizationProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomSynchronizationProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomSynchronizationProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomSynchronizationProviderResponse(val *CustomSynchronizationProviderResponse) *NullableCustomSynchronizationProviderResponse {
	return &NullableCustomSynchronizationProviderResponse{value: val, isSet: true}
}

func (v NullableCustomSynchronizationProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomSynchronizationProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}