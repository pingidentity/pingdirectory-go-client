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

// checks if the CancelExtendedOperationHandlerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelExtendedOperationHandlerResponse{}

// CancelExtendedOperationHandlerResponse struct for CancelExtendedOperationHandlerResponse
type CancelExtendedOperationHandlerResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumcancelExtendedOperationHandlerSchemaUrn      `json:"schemas"`
	// Name of the Extended Operation Handler
	Id string `json:"id"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled bool `json:"enabled"`
}

// NewCancelExtendedOperationHandlerResponse instantiates a new CancelExtendedOperationHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelExtendedOperationHandlerResponse(schemas []EnumcancelExtendedOperationHandlerSchemaUrn, id string, enabled bool) *CancelExtendedOperationHandlerResponse {
	this := CancelExtendedOperationHandlerResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewCancelExtendedOperationHandlerResponseWithDefaults instantiates a new CancelExtendedOperationHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelExtendedOperationHandlerResponseWithDefaults() *CancelExtendedOperationHandlerResponse {
	this := CancelExtendedOperationHandlerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CancelExtendedOperationHandlerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CancelExtendedOperationHandlerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *CancelExtendedOperationHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *CancelExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *CancelExtendedOperationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *CancelExtendedOperationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *CancelExtendedOperationHandlerResponse) GetSchemas() []EnumcancelExtendedOperationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumcancelExtendedOperationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *CancelExtendedOperationHandlerResponse) GetSchemasOk() ([]EnumcancelExtendedOperationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *CancelExtendedOperationHandlerResponse) SetSchemas(v []EnumcancelExtendedOperationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *CancelExtendedOperationHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CancelExtendedOperationHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CancelExtendedOperationHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CancelExtendedOperationHandlerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CancelExtendedOperationHandlerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CancelExtendedOperationHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *CancelExtendedOperationHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CancelExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CancelExtendedOperationHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o CancelExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelExtendedOperationHandlerResponse) ToMap() (map[string]interface{}, error) {
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

type NullableCancelExtendedOperationHandlerResponse struct {
	value *CancelExtendedOperationHandlerResponse
	isSet bool
}

func (v NullableCancelExtendedOperationHandlerResponse) Get() *CancelExtendedOperationHandlerResponse {
	return v.value
}

func (v *NullableCancelExtendedOperationHandlerResponse) Set(val *CancelExtendedOperationHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelExtendedOperationHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelExtendedOperationHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelExtendedOperationHandlerResponse(val *CancelExtendedOperationHandlerResponse) *NullableCancelExtendedOperationHandlerResponse {
	return &NullableCancelExtendedOperationHandlerResponse{value: val, isSet: true}
}

func (v NullableCancelExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelExtendedOperationHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
