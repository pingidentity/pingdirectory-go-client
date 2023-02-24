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

// checks if the AddThirdPartyLogFileRotationListenerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddThirdPartyLogFileRotationListenerRequest{}

// AddThirdPartyLogFileRotationListenerRequest struct for AddThirdPartyLogFileRotationListenerRequest
type AddThirdPartyLogFileRotationListenerRequest struct {
	// Name of the new Log File Rotation Listener
	ListenerName string                                           `json:"listenerName"`
	Schemas      []EnumthirdPartyLogFileRotationListenerSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Log File Rotation Listener.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Log File Rotation Listener. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Log File Rotation Listener
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log File Rotation Listener is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddThirdPartyLogFileRotationListenerRequest instantiates a new AddThirdPartyLogFileRotationListenerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyLogFileRotationListenerRequest(listenerName string, schemas []EnumthirdPartyLogFileRotationListenerSchemaUrn, extensionClass string, enabled bool) *AddThirdPartyLogFileRotationListenerRequest {
	this := AddThirdPartyLogFileRotationListenerRequest{}
	this.ListenerName = listenerName
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewAddThirdPartyLogFileRotationListenerRequestWithDefaults instantiates a new AddThirdPartyLogFileRotationListenerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyLogFileRotationListenerRequestWithDefaults() *AddThirdPartyLogFileRotationListenerRequest {
	this := AddThirdPartyLogFileRotationListenerRequest{}
	return &this
}

// GetListenerName returns the ListenerName field value
func (o *AddThirdPartyLogFileRotationListenerRequest) GetListenerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListenerName
}

// GetListenerNameOk returns a tuple with the ListenerName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyLogFileRotationListenerRequest) GetListenerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListenerName, true
}

// SetListenerName sets field value
func (o *AddThirdPartyLogFileRotationListenerRequest) SetListenerName(v string) {
	o.ListenerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyLogFileRotationListenerRequest) GetSchemas() []EnumthirdPartyLogFileRotationListenerSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyLogFileRotationListenerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyLogFileRotationListenerRequest) GetSchemasOk() ([]EnumthirdPartyLogFileRotationListenerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyLogFileRotationListenerRequest) SetSchemas(v []EnumthirdPartyLogFileRotationListenerSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyLogFileRotationListenerRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyLogFileRotationListenerRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyLogFileRotationListenerRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyLogFileRotationListenerRequest) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyLogFileRotationListenerRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyLogFileRotationListenerRequest) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyLogFileRotationListenerRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyLogFileRotationListenerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyLogFileRotationListenerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyLogFileRotationListenerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyLogFileRotationListenerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyLogFileRotationListenerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyLogFileRotationListenerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyLogFileRotationListenerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddThirdPartyLogFileRotationListenerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddThirdPartyLogFileRotationListenerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["listenerName"] = o.ListenerName
	toSerialize["schemas"] = o.Schemas
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAddThirdPartyLogFileRotationListenerRequest struct {
	value *AddThirdPartyLogFileRotationListenerRequest
	isSet bool
}

func (v NullableAddThirdPartyLogFileRotationListenerRequest) Get() *AddThirdPartyLogFileRotationListenerRequest {
	return v.value
}

func (v *NullableAddThirdPartyLogFileRotationListenerRequest) Set(val *AddThirdPartyLogFileRotationListenerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyLogFileRotationListenerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyLogFileRotationListenerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyLogFileRotationListenerRequest(val *AddThirdPartyLogFileRotationListenerRequest) *NullableAddThirdPartyLogFileRotationListenerRequest {
	return &NullableAddThirdPartyLogFileRotationListenerRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyLogFileRotationListenerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyLogFileRotationListenerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
