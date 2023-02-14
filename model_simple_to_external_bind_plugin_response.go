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

// SimpleToExternalBindPluginResponse struct for SimpleToExternalBindPluginResponse
type SimpleToExternalBindPluginResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Plugin
	Id      string                                    `json:"id"`
	Schemas []EnumsimpleToExternalBindPluginSchemaUrn `json:"schemas"`
	// Specifies a connection criteria object that may be used to indicate the set of clients for which this plugin should be used. If a value is provided, then this plugin will only be used for requests from client connections matching this criteria.
	ConnectionCriteria *string `json:"connectionCriteria,omitempty"`
	// Specifies a request criteria object that may be used to indicate the set of requests for which this plugin should be used. If a value is provided, then this plugin will only be used for bind requests matching this criteria.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewSimpleToExternalBindPluginResponse instantiates a new SimpleToExternalBindPluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleToExternalBindPluginResponse(id string, schemas []EnumsimpleToExternalBindPluginSchemaUrn, enabled bool) *SimpleToExternalBindPluginResponse {
	this := SimpleToExternalBindPluginResponse{}
	this.Id = id
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewSimpleToExternalBindPluginResponseWithDefaults instantiates a new SimpleToExternalBindPluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleToExternalBindPluginResponseWithDefaults() *SimpleToExternalBindPluginResponse {
	this := SimpleToExternalBindPluginResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SimpleToExternalBindPluginResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleToExternalBindPluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SimpleToExternalBindPluginResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SimpleToExternalBindPluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SimpleToExternalBindPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleToExternalBindPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SimpleToExternalBindPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SimpleToExternalBindPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *SimpleToExternalBindPluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimpleToExternalBindPluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimpleToExternalBindPluginResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *SimpleToExternalBindPluginResponse) GetSchemas() []EnumsimpleToExternalBindPluginSchemaUrn {
	if o == nil {
		var ret []EnumsimpleToExternalBindPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SimpleToExternalBindPluginResponse) GetSchemasOk() ([]EnumsimpleToExternalBindPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SimpleToExternalBindPluginResponse) SetSchemas(v []EnumsimpleToExternalBindPluginSchemaUrn) {
	o.Schemas = v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *SimpleToExternalBindPluginResponse) GetConnectionCriteria() string {
	if o == nil || isNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleToExternalBindPluginResponse) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.ConnectionCriteria) {
		return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *SimpleToExternalBindPluginResponse) HasConnectionCriteria() bool {
	if o != nil && !isNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *SimpleToExternalBindPluginResponse) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *SimpleToExternalBindPluginResponse) GetRequestCriteria() string {
	if o == nil || isNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleToExternalBindPluginResponse) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *SimpleToExternalBindPluginResponse) HasRequestCriteria() bool {
	if o != nil && !isNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *SimpleToExternalBindPluginResponse) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SimpleToExternalBindPluginResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleToExternalBindPluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SimpleToExternalBindPluginResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SimpleToExternalBindPluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *SimpleToExternalBindPluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SimpleToExternalBindPluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SimpleToExternalBindPluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o SimpleToExternalBindPluginResponse) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.ConnectionCriteria) {
		toSerialize["connectionCriteria"] = o.ConnectionCriteria
	}
	if !isNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSimpleToExternalBindPluginResponse struct {
	value *SimpleToExternalBindPluginResponse
	isSet bool
}

func (v NullableSimpleToExternalBindPluginResponse) Get() *SimpleToExternalBindPluginResponse {
	return v.value
}

func (v *NullableSimpleToExternalBindPluginResponse) Set(val *SimpleToExternalBindPluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleToExternalBindPluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleToExternalBindPluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleToExternalBindPluginResponse(val *SimpleToExternalBindPluginResponse) *NullableSimpleToExternalBindPluginResponse {
	return &NullableSimpleToExternalBindPluginResponse{value: val, isSet: true}
}

func (v NullableSimpleToExternalBindPluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleToExternalBindPluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
