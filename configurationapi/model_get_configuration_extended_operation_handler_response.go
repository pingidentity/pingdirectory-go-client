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

// checks if the GetConfigurationExtendedOperationHandlerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetConfigurationExtendedOperationHandlerResponse{}

// GetConfigurationExtendedOperationHandlerResponse struct for GetConfigurationExtendedOperationHandlerResponse
type GetConfigurationExtendedOperationHandlerResponse struct {
	Schemas []EnumgetConfigurationExtendedOperationHandlerSchemaUrn `json:"schemas"`
	// Name of the Extended Operation Handler
	Id string `json:"id"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled                                       bool                                               `json:"enabled"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewGetConfigurationExtendedOperationHandlerResponse instantiates a new GetConfigurationExtendedOperationHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConfigurationExtendedOperationHandlerResponse(schemas []EnumgetConfigurationExtendedOperationHandlerSchemaUrn, id string, enabled bool) *GetConfigurationExtendedOperationHandlerResponse {
	this := GetConfigurationExtendedOperationHandlerResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewGetConfigurationExtendedOperationHandlerResponseWithDefaults instantiates a new GetConfigurationExtendedOperationHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConfigurationExtendedOperationHandlerResponseWithDefaults() *GetConfigurationExtendedOperationHandlerResponse {
	this := GetConfigurationExtendedOperationHandlerResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *GetConfigurationExtendedOperationHandlerResponse) GetSchemas() []EnumgetConfigurationExtendedOperationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumgetConfigurationExtendedOperationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GetConfigurationExtendedOperationHandlerResponse) GetSchemasOk() ([]EnumgetConfigurationExtendedOperationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GetConfigurationExtendedOperationHandlerResponse) SetSchemas(v []EnumgetConfigurationExtendedOperationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *GetConfigurationExtendedOperationHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetConfigurationExtendedOperationHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetConfigurationExtendedOperationHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetConfigurationExtendedOperationHandlerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfigurationExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetConfigurationExtendedOperationHandlerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetConfigurationExtendedOperationHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *GetConfigurationExtendedOperationHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GetConfigurationExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GetConfigurationExtendedOperationHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetConfigurationExtendedOperationHandlerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfigurationExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetConfigurationExtendedOperationHandlerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GetConfigurationExtendedOperationHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GetConfigurationExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfigurationExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GetConfigurationExtendedOperationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GetConfigurationExtendedOperationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o GetConfigurationExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetConfigurationExtendedOperationHandlerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableGetConfigurationExtendedOperationHandlerResponse struct {
	value *GetConfigurationExtendedOperationHandlerResponse
	isSet bool
}

func (v NullableGetConfigurationExtendedOperationHandlerResponse) Get() *GetConfigurationExtendedOperationHandlerResponse {
	return v.value
}

func (v *NullableGetConfigurationExtendedOperationHandlerResponse) Set(val *GetConfigurationExtendedOperationHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConfigurationExtendedOperationHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConfigurationExtendedOperationHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConfigurationExtendedOperationHandlerResponse(val *GetConfigurationExtendedOperationHandlerResponse) *NullableGetConfigurationExtendedOperationHandlerResponse {
	return &NullableGetConfigurationExtendedOperationHandlerResponse{value: val, isSet: true}
}

func (v NullableGetConfigurationExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConfigurationExtendedOperationHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
