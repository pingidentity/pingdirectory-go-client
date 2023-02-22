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

// AddSearchShutdownPluginRequest struct for AddSearchShutdownPluginRequest
type AddSearchShutdownPluginRequest struct {
	// Name of the new Plugin
	PluginName string                              `json:"pluginName"`
	Schemas    []EnumsearchShutdownPluginSchemaUrn `json:"schemas"`
	// The base DN to use for the search.
	BaseDN *string             `json:"baseDN,omitempty"`
	Scope  EnumpluginScopeProp `json:"scope"`
	// The filter to use for the search.
	Filter string `json:"filter"`
	// The name of an attribute that should be included in the results. This may include any token which is allowed as a requested attribute in search requests, including the name of an attribute, an asterisk (to indicate all user attributes), a plus sign (to indicate all operational attributes), an object class name preceded with an at symbol (to indicate all attributes associated with that object class), an attribute name preceded by a caret (to indicate that attribute should be excluded), or an object class name preceded by a caret and an at symbol (to indicate that all attributes associated with that object class should be excluded).
	IncludeAttribute []string `json:"includeAttribute,omitempty"`
	// The path of an LDIF file that should be created with the results of the search.
	OutputFile string `json:"outputFile"`
	// An extension that should be appended to the name of an existing output file rather than deleting it. If a file already exists with the full previous file name, then it will be deleted before the current file is renamed to become the previous file.
	PreviousFileExtension *string `json:"previousFileExtension,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddSearchShutdownPluginRequest instantiates a new AddSearchShutdownPluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSearchShutdownPluginRequest(pluginName string, schemas []EnumsearchShutdownPluginSchemaUrn, scope EnumpluginScopeProp, filter string, outputFile string, enabled bool) *AddSearchShutdownPluginRequest {
	this := AddSearchShutdownPluginRequest{}
	this.PluginName = pluginName
	this.Schemas = schemas
	this.Scope = scope
	this.Filter = filter
	this.OutputFile = outputFile
	this.Enabled = enabled
	return &this
}

// NewAddSearchShutdownPluginRequestWithDefaults instantiates a new AddSearchShutdownPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSearchShutdownPluginRequestWithDefaults() *AddSearchShutdownPluginRequest {
	this := AddSearchShutdownPluginRequest{}
	return &this
}

// GetPluginName returns the PluginName field value
func (o *AddSearchShutdownPluginRequest) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *AddSearchShutdownPluginRequest) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *AddSearchShutdownPluginRequest) SetPluginName(v string) {
	o.PluginName = v
}

// GetSchemas returns the Schemas field value
func (o *AddSearchShutdownPluginRequest) GetSchemas() []EnumsearchShutdownPluginSchemaUrn {
	if o == nil {
		var ret []EnumsearchShutdownPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSearchShutdownPluginRequest) GetSchemasOk() ([]EnumsearchShutdownPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSearchShutdownPluginRequest) SetSchemas(v []EnumsearchShutdownPluginSchemaUrn) {
	o.Schemas = v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *AddSearchShutdownPluginRequest) GetBaseDN() string {
	if o == nil || isNil(o.BaseDN) {
		var ret string
		return ret
	}
	return *o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSearchShutdownPluginRequest) GetBaseDNOk() (*string, bool) {
	if o == nil || isNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *AddSearchShutdownPluginRequest) HasBaseDN() bool {
	if o != nil && !isNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given string and assigns it to the BaseDN field.
func (o *AddSearchShutdownPluginRequest) SetBaseDN(v string) {
	o.BaseDN = &v
}

// GetScope returns the Scope field value
func (o *AddSearchShutdownPluginRequest) GetScope() EnumpluginScopeProp {
	if o == nil {
		var ret EnumpluginScopeProp
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AddSearchShutdownPluginRequest) GetScopeOk() (*EnumpluginScopeProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *AddSearchShutdownPluginRequest) SetScope(v EnumpluginScopeProp) {
	o.Scope = v
}

// GetFilter returns the Filter field value
func (o *AddSearchShutdownPluginRequest) GetFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *AddSearchShutdownPluginRequest) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *AddSearchShutdownPluginRequest) SetFilter(v string) {
	o.Filter = v
}

// GetIncludeAttribute returns the IncludeAttribute field value if set, zero value otherwise.
func (o *AddSearchShutdownPluginRequest) GetIncludeAttribute() []string {
	if o == nil || isNil(o.IncludeAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeAttribute
}

// GetIncludeAttributeOk returns a tuple with the IncludeAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSearchShutdownPluginRequest) GetIncludeAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeAttribute) {
		return nil, false
	}
	return o.IncludeAttribute, true
}

// HasIncludeAttribute returns a boolean if a field has been set.
func (o *AddSearchShutdownPluginRequest) HasIncludeAttribute() bool {
	if o != nil && !isNil(o.IncludeAttribute) {
		return true
	}

	return false
}

// SetIncludeAttribute gets a reference to the given []string and assigns it to the IncludeAttribute field.
func (o *AddSearchShutdownPluginRequest) SetIncludeAttribute(v []string) {
	o.IncludeAttribute = v
}

// GetOutputFile returns the OutputFile field value
func (o *AddSearchShutdownPluginRequest) GetOutputFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutputFile
}

// GetOutputFileOk returns a tuple with the OutputFile field value
// and a boolean to check if the value has been set.
func (o *AddSearchShutdownPluginRequest) GetOutputFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputFile, true
}

// SetOutputFile sets field value
func (o *AddSearchShutdownPluginRequest) SetOutputFile(v string) {
	o.OutputFile = v
}

// GetPreviousFileExtension returns the PreviousFileExtension field value if set, zero value otherwise.
func (o *AddSearchShutdownPluginRequest) GetPreviousFileExtension() string {
	if o == nil || isNil(o.PreviousFileExtension) {
		var ret string
		return ret
	}
	return *o.PreviousFileExtension
}

// GetPreviousFileExtensionOk returns a tuple with the PreviousFileExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSearchShutdownPluginRequest) GetPreviousFileExtensionOk() (*string, bool) {
	if o == nil || isNil(o.PreviousFileExtension) {
		return nil, false
	}
	return o.PreviousFileExtension, true
}

// HasPreviousFileExtension returns a boolean if a field has been set.
func (o *AddSearchShutdownPluginRequest) HasPreviousFileExtension() bool {
	if o != nil && !isNil(o.PreviousFileExtension) {
		return true
	}

	return false
}

// SetPreviousFileExtension gets a reference to the given string and assigns it to the PreviousFileExtension field.
func (o *AddSearchShutdownPluginRequest) SetPreviousFileExtension(v string) {
	o.PreviousFileExtension = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSearchShutdownPluginRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSearchShutdownPluginRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSearchShutdownPluginRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSearchShutdownPluginRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddSearchShutdownPluginRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddSearchShutdownPluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddSearchShutdownPluginRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddSearchShutdownPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pluginName"] = o.PluginName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.BaseDN) {
		toSerialize["baseDN"] = o.BaseDN
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["filter"] = o.Filter
	}
	if !isNil(o.IncludeAttribute) {
		toSerialize["includeAttribute"] = o.IncludeAttribute
	}
	if true {
		toSerialize["outputFile"] = o.OutputFile
	}
	if !isNil(o.PreviousFileExtension) {
		toSerialize["previousFileExtension"] = o.PreviousFileExtension
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddSearchShutdownPluginRequest struct {
	value *AddSearchShutdownPluginRequest
	isSet bool
}

func (v NullableAddSearchShutdownPluginRequest) Get() *AddSearchShutdownPluginRequest {
	return v.value
}

func (v *NullableAddSearchShutdownPluginRequest) Set(val *AddSearchShutdownPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSearchShutdownPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSearchShutdownPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSearchShutdownPluginRequest(val *AddSearchShutdownPluginRequest) *NullableAddSearchShutdownPluginRequest {
	return &NullableAddSearchShutdownPluginRequest{value: val, isSet: true}
}

func (v NullableAddSearchShutdownPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSearchShutdownPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}