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

// ThirdPartyIdentityMapperResponse struct for ThirdPartyIdentityMapperResponse
type ThirdPartyIdentityMapperResponse struct {
	// Name of the Identity Mapper
	Id string `json:"id"`
	Schemas []EnumthirdPartyIdentityMapperSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Identity Mapper.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Identity Mapper. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Identity Mapper
	Description *string `json:"description,omitempty"`
	// Indicates whether the Identity Mapper is enabled for use.
	Enabled bool `json:"enabled"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewThirdPartyIdentityMapperResponse instantiates a new ThirdPartyIdentityMapperResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyIdentityMapperResponse(id string, schemas []EnumthirdPartyIdentityMapperSchemaUrn, extensionClass string, enabled bool) *ThirdPartyIdentityMapperResponse {
	this := ThirdPartyIdentityMapperResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewThirdPartyIdentityMapperResponseWithDefaults instantiates a new ThirdPartyIdentityMapperResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyIdentityMapperResponseWithDefaults() *ThirdPartyIdentityMapperResponse {
	this := ThirdPartyIdentityMapperResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ThirdPartyIdentityMapperResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyIdentityMapperResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThirdPartyIdentityMapperResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ThirdPartyIdentityMapperResponse) GetSchemas() []EnumthirdPartyIdentityMapperSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyIdentityMapperSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyIdentityMapperResponse) GetSchemasOk() ([]EnumthirdPartyIdentityMapperSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ThirdPartyIdentityMapperResponse) SetSchemas(v []EnumthirdPartyIdentityMapperSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartyIdentityMapperResponse) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyIdentityMapperResponse) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartyIdentityMapperResponse) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartyIdentityMapperResponse) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIdentityMapperResponse) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartyIdentityMapperResponse) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartyIdentityMapperResponse) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartyIdentityMapperResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIdentityMapperResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyIdentityMapperResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartyIdentityMapperResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyIdentityMapperResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyIdentityMapperResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyIdentityMapperResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ThirdPartyIdentityMapperResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIdentityMapperResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ThirdPartyIdentityMapperResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ThirdPartyIdentityMapperResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o ThirdPartyIdentityMapperResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["extensionClass"] = o.ExtensionClass
	}
	if !isNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableThirdPartyIdentityMapperResponse struct {
	value *ThirdPartyIdentityMapperResponse
	isSet bool
}

func (v NullableThirdPartyIdentityMapperResponse) Get() *ThirdPartyIdentityMapperResponse {
	return v.value
}

func (v *NullableThirdPartyIdentityMapperResponse) Set(val *ThirdPartyIdentityMapperResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyIdentityMapperResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyIdentityMapperResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyIdentityMapperResponse(val *ThirdPartyIdentityMapperResponse) *NullableThirdPartyIdentityMapperResponse {
	return &NullableThirdPartyIdentityMapperResponse{value: val, isSet: true}
}

func (v NullableThirdPartyIdentityMapperResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyIdentityMapperResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


