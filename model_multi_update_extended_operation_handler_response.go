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

// MultiUpdateExtendedOperationHandlerResponse struct for MultiUpdateExtendedOperationHandlerResponse
type MultiUpdateExtendedOperationHandlerResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas []EnummultiUpdateExtendedOperationHandlerSchemaUrn `json:"schemas"`
	// Name of the Extended Operation Handler
	Id string `json:"id"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled bool `json:"enabled"`
}

// NewMultiUpdateExtendedOperationHandlerResponse instantiates a new MultiUpdateExtendedOperationHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiUpdateExtendedOperationHandlerResponse(schemas []EnummultiUpdateExtendedOperationHandlerSchemaUrn, id string, enabled bool) *MultiUpdateExtendedOperationHandlerResponse {
	this := MultiUpdateExtendedOperationHandlerResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewMultiUpdateExtendedOperationHandlerResponseWithDefaults instantiates a new MultiUpdateExtendedOperationHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiUpdateExtendedOperationHandlerResponseWithDefaults() *MultiUpdateExtendedOperationHandlerResponse {
	this := MultiUpdateExtendedOperationHandlerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *MultiUpdateExtendedOperationHandlerResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiUpdateExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *MultiUpdateExtendedOperationHandlerResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *MultiUpdateExtendedOperationHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *MultiUpdateExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiUpdateExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *MultiUpdateExtendedOperationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *MultiUpdateExtendedOperationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *MultiUpdateExtendedOperationHandlerResponse) GetSchemas() []EnummultiUpdateExtendedOperationHandlerSchemaUrn {
	if o == nil {
		var ret []EnummultiUpdateExtendedOperationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *MultiUpdateExtendedOperationHandlerResponse) GetSchemasOk() ([]EnummultiUpdateExtendedOperationHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *MultiUpdateExtendedOperationHandlerResponse) SetSchemas(v []EnummultiUpdateExtendedOperationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *MultiUpdateExtendedOperationHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MultiUpdateExtendedOperationHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MultiUpdateExtendedOperationHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MultiUpdateExtendedOperationHandlerResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiUpdateExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MultiUpdateExtendedOperationHandlerResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MultiUpdateExtendedOperationHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *MultiUpdateExtendedOperationHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *MultiUpdateExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *MultiUpdateExtendedOperationHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o MultiUpdateExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableMultiUpdateExtendedOperationHandlerResponse struct {
	value *MultiUpdateExtendedOperationHandlerResponse
	isSet bool
}

func (v NullableMultiUpdateExtendedOperationHandlerResponse) Get() *MultiUpdateExtendedOperationHandlerResponse {
	return v.value
}

func (v *NullableMultiUpdateExtendedOperationHandlerResponse) Set(val *MultiUpdateExtendedOperationHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiUpdateExtendedOperationHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiUpdateExtendedOperationHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiUpdateExtendedOperationHandlerResponse(val *MultiUpdateExtendedOperationHandlerResponse) *NullableMultiUpdateExtendedOperationHandlerResponse {
	return &NullableMultiUpdateExtendedOperationHandlerResponse{value: val, isSet: true}
}

func (v NullableMultiUpdateExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiUpdateExtendedOperationHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


