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

// checks if the AddThirdPartySaslMechanismHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddThirdPartySaslMechanismHandlerRequest{}

// AddThirdPartySaslMechanismHandlerRequest struct for AddThirdPartySaslMechanismHandlerRequest
type AddThirdPartySaslMechanismHandlerRequest struct {
	Schemas []EnumthirdPartySaslMechanismHandlerSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party SASL Mechanism Handler.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party SASL Mechanism Handler. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// The identity mapper that may be used to map usernames to user entries. If the custom SASL mechanism involves a username or some other form of authentication and/or authorization identity, then this may be used to map that ID to an entry for that user.
	IdentityMapper *string `json:"identityMapper,omitempty"`
	// A description for this SASL Mechanism Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the SASL mechanism handler is enabled for use.
	Enabled bool `json:"enabled"`
	// Name of the new SASL Mechanism Handler
	HandlerName string `json:"handlerName"`
}

// NewAddThirdPartySaslMechanismHandlerRequest instantiates a new AddThirdPartySaslMechanismHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartySaslMechanismHandlerRequest(schemas []EnumthirdPartySaslMechanismHandlerSchemaUrn, extensionClass string, enabled bool, handlerName string) *AddThirdPartySaslMechanismHandlerRequest {
	this := AddThirdPartySaslMechanismHandlerRequest{}
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	this.HandlerName = handlerName
	return &this
}

// NewAddThirdPartySaslMechanismHandlerRequestWithDefaults instantiates a new AddThirdPartySaslMechanismHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartySaslMechanismHandlerRequestWithDefaults() *AddThirdPartySaslMechanismHandlerRequest {
	this := AddThirdPartySaslMechanismHandlerRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartySaslMechanismHandlerRequest) GetSchemas() []EnumthirdPartySaslMechanismHandlerSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartySaslMechanismHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartySaslMechanismHandlerRequest) GetSchemasOk() ([]EnumthirdPartySaslMechanismHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartySaslMechanismHandlerRequest) SetSchemas(v []EnumthirdPartySaslMechanismHandlerSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartySaslMechanismHandlerRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartySaslMechanismHandlerRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartySaslMechanismHandlerRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartySaslMechanismHandlerRequest) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartySaslMechanismHandlerRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartySaslMechanismHandlerRequest) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartySaslMechanismHandlerRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetIdentityMapper returns the IdentityMapper field value if set, zero value otherwise.
func (o *AddThirdPartySaslMechanismHandlerRequest) GetIdentityMapper() string {
	if o == nil || IsNil(o.IdentityMapper) {
		var ret string
		return ret
	}
	return *o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartySaslMechanismHandlerRequest) GetIdentityMapperOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityMapper) {
		return nil, false
	}
	return o.IdentityMapper, true
}

// HasIdentityMapper returns a boolean if a field has been set.
func (o *AddThirdPartySaslMechanismHandlerRequest) HasIdentityMapper() bool {
	if o != nil && !IsNil(o.IdentityMapper) {
		return true
	}

	return false
}

// SetIdentityMapper gets a reference to the given string and assigns it to the IdentityMapper field.
func (o *AddThirdPartySaslMechanismHandlerRequest) SetIdentityMapper(v string) {
	o.IdentityMapper = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartySaslMechanismHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartySaslMechanismHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartySaslMechanismHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartySaslMechanismHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartySaslMechanismHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartySaslMechanismHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartySaslMechanismHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetHandlerName returns the HandlerName field value
func (o *AddThirdPartySaslMechanismHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartySaslMechanismHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddThirdPartySaslMechanismHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

func (o AddThirdPartySaslMechanismHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddThirdPartySaslMechanismHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !IsNil(o.IdentityMapper) {
		toSerialize["identityMapper"] = o.IdentityMapper
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["handlerName"] = o.HandlerName
	return toSerialize, nil
}

type NullableAddThirdPartySaslMechanismHandlerRequest struct {
	value *AddThirdPartySaslMechanismHandlerRequest
	isSet bool
}

func (v NullableAddThirdPartySaslMechanismHandlerRequest) Get() *AddThirdPartySaslMechanismHandlerRequest {
	return v.value
}

func (v *NullableAddThirdPartySaslMechanismHandlerRequest) Set(val *AddThirdPartySaslMechanismHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartySaslMechanismHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartySaslMechanismHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartySaslMechanismHandlerRequest(val *AddThirdPartySaslMechanismHandlerRequest) *NullableAddThirdPartySaslMechanismHandlerRequest {
	return &NullableAddThirdPartySaslMechanismHandlerRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartySaslMechanismHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartySaslMechanismHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
