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

// checks if the UnboundidInterServerSaslMechanismHandlerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnboundidInterServerSaslMechanismHandlerResponse{}

// UnboundidInterServerSaslMechanismHandlerResponse struct for UnboundidInterServerSaslMechanismHandlerResponse
type UnboundidInterServerSaslMechanismHandlerResponse struct {
	Meta                                          *MetaMeta                                               `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20      `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumunboundidInterServerSaslMechanismHandlerSchemaUrn `json:"schemas"`
	// Name of the SASL Mechanism Handler
	Id string `json:"id"`
	// A description for this SASL Mechanism Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the SASL mechanism handler is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewUnboundidInterServerSaslMechanismHandlerResponse instantiates a new UnboundidInterServerSaslMechanismHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnboundidInterServerSaslMechanismHandlerResponse(schemas []EnumunboundidInterServerSaslMechanismHandlerSchemaUrn, id string, enabled bool) *UnboundidInterServerSaslMechanismHandlerResponse {
	this := UnboundidInterServerSaslMechanismHandlerResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewUnboundidInterServerSaslMechanismHandlerResponseWithDefaults instantiates a new UnboundidInterServerSaslMechanismHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnboundidInterServerSaslMechanismHandlerResponseWithDefaults() *UnboundidInterServerSaslMechanismHandlerResponse {
	this := UnboundidInterServerSaslMechanismHandlerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *UnboundidInterServerSaslMechanismHandlerResponse) GetSchemas() []EnumunboundidInterServerSaslMechanismHandlerSchemaUrn {
	if o == nil {
		var ret []EnumunboundidInterServerSaslMechanismHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) GetSchemasOk() ([]EnumunboundidInterServerSaslMechanismHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *UnboundidInterServerSaslMechanismHandlerResponse) SetSchemas(v []EnumunboundidInterServerSaslMechanismHandlerSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *UnboundidInterServerSaslMechanismHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnboundidInterServerSaslMechanismHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *UnboundidInterServerSaslMechanismHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UnboundidInterServerSaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UnboundidInterServerSaslMechanismHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o UnboundidInterServerSaslMechanismHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnboundidInterServerSaslMechanismHandlerResponse) ToMap() (map[string]interface{}, error) {
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

type NullableUnboundidInterServerSaslMechanismHandlerResponse struct {
	value *UnboundidInterServerSaslMechanismHandlerResponse
	isSet bool
}

func (v NullableUnboundidInterServerSaslMechanismHandlerResponse) Get() *UnboundidInterServerSaslMechanismHandlerResponse {
	return v.value
}

func (v *NullableUnboundidInterServerSaslMechanismHandlerResponse) Set(val *UnboundidInterServerSaslMechanismHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUnboundidInterServerSaslMechanismHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUnboundidInterServerSaslMechanismHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnboundidInterServerSaslMechanismHandlerResponse(val *UnboundidInterServerSaslMechanismHandlerResponse) *NullableUnboundidInterServerSaslMechanismHandlerResponse {
	return &NullableUnboundidInterServerSaslMechanismHandlerResponse{value: val, isSet: true}
}

func (v NullableUnboundidInterServerSaslMechanismHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnboundidInterServerSaslMechanismHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
