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

// checks if the AddSimpleToExternalBindPluginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSimpleToExternalBindPluginRequest{}

// AddSimpleToExternalBindPluginRequest struct for AddSimpleToExternalBindPluginRequest
type AddSimpleToExternalBindPluginRequest struct {
	Schemas []EnumsimpleToExternalBindPluginSchemaUrn `json:"schemas"`
	// Specifies a connection criteria object that may be used to indicate the set of clients for which this plugin should be used. If a value is provided, then this plugin will only be used for requests from client connections matching this criteria.
	ConnectionCriteria *string `json:"connectionCriteria,omitempty"`
	// Specifies a request criteria object that may be used to indicate the set of requests for which this plugin should be used. If a value is provided, then this plugin will only be used for bind requests matching this criteria.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
	// Name of the new Plugin
	PluginName string `json:"pluginName"`
}

// NewAddSimpleToExternalBindPluginRequest instantiates a new AddSimpleToExternalBindPluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSimpleToExternalBindPluginRequest(schemas []EnumsimpleToExternalBindPluginSchemaUrn, enabled bool, pluginName string) *AddSimpleToExternalBindPluginRequest {
	this := AddSimpleToExternalBindPluginRequest{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.PluginName = pluginName
	return &this
}

// NewAddSimpleToExternalBindPluginRequestWithDefaults instantiates a new AddSimpleToExternalBindPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSimpleToExternalBindPluginRequestWithDefaults() *AddSimpleToExternalBindPluginRequest {
	this := AddSimpleToExternalBindPluginRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddSimpleToExternalBindPluginRequest) GetSchemas() []EnumsimpleToExternalBindPluginSchemaUrn {
	if o == nil {
		var ret []EnumsimpleToExternalBindPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSimpleToExternalBindPluginRequest) GetSchemasOk() ([]EnumsimpleToExternalBindPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSimpleToExternalBindPluginRequest) SetSchemas(v []EnumsimpleToExternalBindPluginSchemaUrn) {
	o.Schemas = v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *AddSimpleToExternalBindPluginRequest) GetConnectionCriteria() string {
	if o == nil || IsNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleToExternalBindPluginRequest) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionCriteria) {
		return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *AddSimpleToExternalBindPluginRequest) HasConnectionCriteria() bool {
	if o != nil && !IsNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *AddSimpleToExternalBindPluginRequest) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AddSimpleToExternalBindPluginRequest) GetRequestCriteria() string {
	if o == nil || IsNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleToExternalBindPluginRequest) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AddSimpleToExternalBindPluginRequest) HasRequestCriteria() bool {
	if o != nil && !IsNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AddSimpleToExternalBindPluginRequest) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSimpleToExternalBindPluginRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleToExternalBindPluginRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSimpleToExternalBindPluginRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSimpleToExternalBindPluginRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddSimpleToExternalBindPluginRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddSimpleToExternalBindPluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddSimpleToExternalBindPluginRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetPluginName returns the PluginName field value
func (o *AddSimpleToExternalBindPluginRequest) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *AddSimpleToExternalBindPluginRequest) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *AddSimpleToExternalBindPluginRequest) SetPluginName(v string) {
	o.PluginName = v
}

func (o AddSimpleToExternalBindPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSimpleToExternalBindPluginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.ConnectionCriteria) {
		toSerialize["connectionCriteria"] = o.ConnectionCriteria
	}
	if !IsNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["pluginName"] = o.PluginName
	return toSerialize, nil
}

type NullableAddSimpleToExternalBindPluginRequest struct {
	value *AddSimpleToExternalBindPluginRequest
	isSet bool
}

func (v NullableAddSimpleToExternalBindPluginRequest) Get() *AddSimpleToExternalBindPluginRequest {
	return v.value
}

func (v *NullableAddSimpleToExternalBindPluginRequest) Set(val *AddSimpleToExternalBindPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSimpleToExternalBindPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSimpleToExternalBindPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSimpleToExternalBindPluginRequest(val *AddSimpleToExternalBindPluginRequest) *NullableAddSimpleToExternalBindPluginRequest {
	return &NullableAddSimpleToExternalBindPluginRequest{value: val, isSet: true}
}

func (v NullableAddSimpleToExternalBindPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSimpleToExternalBindPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
