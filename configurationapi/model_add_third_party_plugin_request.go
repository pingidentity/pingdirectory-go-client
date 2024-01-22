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

// checks if the AddThirdPartyPluginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddThirdPartyPluginRequest{}

// AddThirdPartyPluginRequest struct for AddThirdPartyPluginRequest
type AddThirdPartyPluginRequest struct {
	Schemas []EnumthirdPartyPluginSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Plugin.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Plugin. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// Specifies a set of request criteria that may be used to indicate that this Third Party Plugin should only be invoked for operations in which the associated request matches this criteria.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled    bool                       `json:"enabled"`
	PluginType []EnumpluginPluginTypeProp `json:"pluginType"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations *bool `json:"invokeForInternalOperations,omitempty"`
	// Name of the new Plugin
	PluginName string `json:"pluginName"`
}

// NewAddThirdPartyPluginRequest instantiates a new AddThirdPartyPluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyPluginRequest(schemas []EnumthirdPartyPluginSchemaUrn, extensionClass string, enabled bool, pluginType []EnumpluginPluginTypeProp, pluginName string) *AddThirdPartyPluginRequest {
	this := AddThirdPartyPluginRequest{}
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	this.PluginType = pluginType
	this.PluginName = pluginName
	return &this
}

// NewAddThirdPartyPluginRequestWithDefaults instantiates a new AddThirdPartyPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyPluginRequestWithDefaults() *AddThirdPartyPluginRequest {
	this := AddThirdPartyPluginRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyPluginRequest) GetSchemas() []EnumthirdPartyPluginSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPluginRequest) GetSchemasOk() ([]EnumthirdPartyPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyPluginRequest) SetSchemas(v []EnumthirdPartyPluginSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyPluginRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPluginRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyPluginRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyPluginRequest) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPluginRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyPluginRequest) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyPluginRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AddThirdPartyPluginRequest) GetRequestCriteria() string {
	if o == nil || IsNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPluginRequest) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AddThirdPartyPluginRequest) HasRequestCriteria() bool {
	if o != nil && !IsNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AddThirdPartyPluginRequest) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyPluginRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPluginRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyPluginRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyPluginRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyPluginRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyPluginRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetPluginType returns the PluginType field value
func (o *AddThirdPartyPluginRequest) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil {
		var ret []EnumpluginPluginTypeProp
		return ret
	}

	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPluginRequest) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginType, true
}

// SetPluginType sets field value
func (o *AddThirdPartyPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *AddThirdPartyPluginRequest) GetInvokeForInternalOperations() bool {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *AddThirdPartyPluginRequest) HasInvokeForInternalOperations() bool {
	if o != nil && !IsNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *AddThirdPartyPluginRequest) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

// GetPluginName returns the PluginName field value
func (o *AddThirdPartyPluginRequest) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPluginRequest) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *AddThirdPartyPluginRequest) SetPluginName(v string) {
	o.PluginName = v
}

func (o AddThirdPartyPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddThirdPartyPluginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !IsNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["pluginType"] = o.PluginType
	if !IsNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	toSerialize["pluginName"] = o.PluginName
	return toSerialize, nil
}

type NullableAddThirdPartyPluginRequest struct {
	value *AddThirdPartyPluginRequest
	isSet bool
}

func (v NullableAddThirdPartyPluginRequest) Get() *AddThirdPartyPluginRequest {
	return v.value
}

func (v *NullableAddThirdPartyPluginRequest) Set(val *AddThirdPartyPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyPluginRequest(val *AddThirdPartyPluginRequest) *NullableAddThirdPartyPluginRequest {
	return &NullableAddThirdPartyPluginRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
