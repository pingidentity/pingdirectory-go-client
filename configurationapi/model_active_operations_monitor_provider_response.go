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

// checks if the ActiveOperationsMonitorProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveOperationsMonitorProviderResponse{}

// ActiveOperationsMonitorProviderResponse struct for ActiveOperationsMonitorProviderResponse
type ActiveOperationsMonitorProviderResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumactiveOperationsMonitorProviderSchemaUrn     `json:"schemas"`
	// Name of the Monitor Provider
	Id string `json:"id"`
	// A description for this Monitor Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether the Monitor Provider is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewActiveOperationsMonitorProviderResponse instantiates a new ActiveOperationsMonitorProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveOperationsMonitorProviderResponse(schemas []EnumactiveOperationsMonitorProviderSchemaUrn, id string, enabled bool) *ActiveOperationsMonitorProviderResponse {
	this := ActiveOperationsMonitorProviderResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewActiveOperationsMonitorProviderResponseWithDefaults instantiates a new ActiveOperationsMonitorProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveOperationsMonitorProviderResponseWithDefaults() *ActiveOperationsMonitorProviderResponse {
	this := ActiveOperationsMonitorProviderResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ActiveOperationsMonitorProviderResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveOperationsMonitorProviderResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ActiveOperationsMonitorProviderResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ActiveOperationsMonitorProviderResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ActiveOperationsMonitorProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveOperationsMonitorProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ActiveOperationsMonitorProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ActiveOperationsMonitorProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *ActiveOperationsMonitorProviderResponse) GetSchemas() []EnumactiveOperationsMonitorProviderSchemaUrn {
	if o == nil {
		var ret []EnumactiveOperationsMonitorProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ActiveOperationsMonitorProviderResponse) GetSchemasOk() ([]EnumactiveOperationsMonitorProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ActiveOperationsMonitorProviderResponse) SetSchemas(v []EnumactiveOperationsMonitorProviderSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *ActiveOperationsMonitorProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ActiveOperationsMonitorProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ActiveOperationsMonitorProviderResponse) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ActiveOperationsMonitorProviderResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveOperationsMonitorProviderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ActiveOperationsMonitorProviderResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ActiveOperationsMonitorProviderResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ActiveOperationsMonitorProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ActiveOperationsMonitorProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ActiveOperationsMonitorProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ActiveOperationsMonitorProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveOperationsMonitorProviderResponse) ToMap() (map[string]interface{}, error) {
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

type NullableActiveOperationsMonitorProviderResponse struct {
	value *ActiveOperationsMonitorProviderResponse
	isSet bool
}

func (v NullableActiveOperationsMonitorProviderResponse) Get() *ActiveOperationsMonitorProviderResponse {
	return v.value
}

func (v *NullableActiveOperationsMonitorProviderResponse) Set(val *ActiveOperationsMonitorProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveOperationsMonitorProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveOperationsMonitorProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveOperationsMonitorProviderResponse(val *ActiveOperationsMonitorProviderResponse) *NullableActiveOperationsMonitorProviderResponse {
	return &NullableActiveOperationsMonitorProviderResponse{value: val, isSet: true}
}

func (v NullableActiveOperationsMonitorProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveOperationsMonitorProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
