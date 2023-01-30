/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ObscuredValueResponse struct for ObscuredValueResponse
type ObscuredValueResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Obscured Value
	Id      string                       `json:"id"`
	Schemas []EnumobscuredValueSchemaUrn `json:"schemas,omitempty"`
	// A description for this Obscured Value
	Description *string `json:"description,omitempty"`
	// The value to be stored in an obscured form.
	ObscuredValue string `json:"obscuredValue"`
}

// NewObscuredValueResponse instantiates a new ObscuredValueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObscuredValueResponse(id string, obscuredValue string) *ObscuredValueResponse {
	this := ObscuredValueResponse{}
	this.Id = id
	this.ObscuredValue = obscuredValue
	return &this
}

// NewObscuredValueResponseWithDefaults instantiates a new ObscuredValueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObscuredValueResponseWithDefaults() *ObscuredValueResponse {
	this := ObscuredValueResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ObscuredValueResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObscuredValueResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ObscuredValueResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ObscuredValueResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ObscuredValueResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObscuredValueResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ObscuredValueResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ObscuredValueResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ObscuredValueResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObscuredValueResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ObscuredValueResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ObscuredValueResponse) GetSchemas() []EnumobscuredValueSchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumobscuredValueSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObscuredValueResponse) GetSchemasOk() ([]EnumobscuredValueSchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ObscuredValueResponse) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumobscuredValueSchemaUrn and assigns it to the Schemas field.
func (o *ObscuredValueResponse) SetSchemas(v []EnumobscuredValueSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ObscuredValueResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObscuredValueResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ObscuredValueResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ObscuredValueResponse) SetDescription(v string) {
	o.Description = &v
}

// GetObscuredValue returns the ObscuredValue field value
func (o *ObscuredValueResponse) GetObscuredValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObscuredValue
}

// GetObscuredValueOk returns a tuple with the ObscuredValue field value
// and a boolean to check if the value has been set.
func (o *ObscuredValueResponse) GetObscuredValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObscuredValue, true
}

// SetObscuredValue sets field value
func (o *ObscuredValueResponse) SetObscuredValue(v string) {
	o.ObscuredValue = v
}

func (o ObscuredValueResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["obscuredValue"] = o.ObscuredValue
	}
	return json.Marshal(toSerialize)
}

type NullableObscuredValueResponse struct {
	value *ObscuredValueResponse
	isSet bool
}

func (v NullableObscuredValueResponse) Get() *ObscuredValueResponse {
	return v.value
}

func (v *NullableObscuredValueResponse) Set(val *ObscuredValueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableObscuredValueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableObscuredValueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObscuredValueResponse(val *ObscuredValueResponse) *NullableObscuredValueResponse {
	return &NullableObscuredValueResponse{value: val, isSet: true}
}

func (v NullableObscuredValueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObscuredValueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
