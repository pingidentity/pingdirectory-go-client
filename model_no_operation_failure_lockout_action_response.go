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

// NoOperationFailureLockoutActionResponse struct for NoOperationFailureLockoutActionResponse
type NoOperationFailureLockoutActionResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Failure Lockout Action
	Id string `json:"id"`
	Schemas []EnumnoOperationFailureLockoutActionSchemaUrn `json:"schemas"`
	// Indicates whether to generate an account status notification for cases in which this failure lockout action is invoked for a bind attempt with too many outstanding authentication failures.
	GenerateAccountStatusNotification *bool `json:"generateAccountStatusNotification,omitempty"`
	// A description for this Failure Lockout Action
	Description *string `json:"description,omitempty"`
}

// NewNoOperationFailureLockoutActionResponse instantiates a new NoOperationFailureLockoutActionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoOperationFailureLockoutActionResponse(id string, schemas []EnumnoOperationFailureLockoutActionSchemaUrn) *NoOperationFailureLockoutActionResponse {
	this := NoOperationFailureLockoutActionResponse{}
	this.Id = id
	this.Schemas = schemas
	return &this
}

// NewNoOperationFailureLockoutActionResponseWithDefaults instantiates a new NoOperationFailureLockoutActionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoOperationFailureLockoutActionResponseWithDefaults() *NoOperationFailureLockoutActionResponse {
	this := NoOperationFailureLockoutActionResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *NoOperationFailureLockoutActionResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoOperationFailureLockoutActionResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *NoOperationFailureLockoutActionResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *NoOperationFailureLockoutActionResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *NoOperationFailureLockoutActionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoOperationFailureLockoutActionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *NoOperationFailureLockoutActionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *NoOperationFailureLockoutActionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *NoOperationFailureLockoutActionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NoOperationFailureLockoutActionResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NoOperationFailureLockoutActionResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *NoOperationFailureLockoutActionResponse) GetSchemas() []EnumnoOperationFailureLockoutActionSchemaUrn {
	if o == nil {
		var ret []EnumnoOperationFailureLockoutActionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *NoOperationFailureLockoutActionResponse) GetSchemasOk() ([]EnumnoOperationFailureLockoutActionSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *NoOperationFailureLockoutActionResponse) SetSchemas(v []EnumnoOperationFailureLockoutActionSchemaUrn) {
	o.Schemas = v
}

// GetGenerateAccountStatusNotification returns the GenerateAccountStatusNotification field value if set, zero value otherwise.
func (o *NoOperationFailureLockoutActionResponse) GetGenerateAccountStatusNotification() bool {
	if o == nil || isNil(o.GenerateAccountStatusNotification) {
		var ret bool
		return ret
	}
	return *o.GenerateAccountStatusNotification
}

// GetGenerateAccountStatusNotificationOk returns a tuple with the GenerateAccountStatusNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoOperationFailureLockoutActionResponse) GetGenerateAccountStatusNotificationOk() (*bool, bool) {
	if o == nil || isNil(o.GenerateAccountStatusNotification) {
    return nil, false
	}
	return o.GenerateAccountStatusNotification, true
}

// HasGenerateAccountStatusNotification returns a boolean if a field has been set.
func (o *NoOperationFailureLockoutActionResponse) HasGenerateAccountStatusNotification() bool {
	if o != nil && !isNil(o.GenerateAccountStatusNotification) {
		return true
	}

	return false
}

// SetGenerateAccountStatusNotification gets a reference to the given bool and assigns it to the GenerateAccountStatusNotification field.
func (o *NoOperationFailureLockoutActionResponse) SetGenerateAccountStatusNotification(v bool) {
	o.GenerateAccountStatusNotification = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NoOperationFailureLockoutActionResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoOperationFailureLockoutActionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NoOperationFailureLockoutActionResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NoOperationFailureLockoutActionResponse) SetDescription(v string) {
	o.Description = &v
}

func (o NoOperationFailureLockoutActionResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.GenerateAccountStatusNotification) {
		toSerialize["generateAccountStatusNotification"] = o.GenerateAccountStatusNotification
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableNoOperationFailureLockoutActionResponse struct {
	value *NoOperationFailureLockoutActionResponse
	isSet bool
}

func (v NullableNoOperationFailureLockoutActionResponse) Get() *NoOperationFailureLockoutActionResponse {
	return v.value
}

func (v *NullableNoOperationFailureLockoutActionResponse) Set(val *NoOperationFailureLockoutActionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNoOperationFailureLockoutActionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNoOperationFailureLockoutActionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoOperationFailureLockoutActionResponse(val *NoOperationFailureLockoutActionResponse) *NullableNoOperationFailureLockoutActionResponse {
	return &NullableNoOperationFailureLockoutActionResponse{value: val, isSet: true}
}

func (v NullableNoOperationFailureLockoutActionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoOperationFailureLockoutActionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


