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

// checks if the LockAccountFailureLockoutActionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LockAccountFailureLockoutActionResponse{}

// LockAccountFailureLockoutActionResponse struct for LockAccountFailureLockoutActionResponse
type LockAccountFailureLockoutActionResponse struct {
	// Name of the Failure Lockout Action
	Id      string                                         `json:"id"`
	Schemas []EnumlockAccountFailureLockoutActionSchemaUrn `json:"schemas"`
	// A description for this Failure Lockout Action
	Description                                   *string                                            `json:"description,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewLockAccountFailureLockoutActionResponse instantiates a new LockAccountFailureLockoutActionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLockAccountFailureLockoutActionResponse(id string, schemas []EnumlockAccountFailureLockoutActionSchemaUrn) *LockAccountFailureLockoutActionResponse {
	this := LockAccountFailureLockoutActionResponse{}
	this.Id = id
	this.Schemas = schemas
	return &this
}

// NewLockAccountFailureLockoutActionResponseWithDefaults instantiates a new LockAccountFailureLockoutActionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLockAccountFailureLockoutActionResponseWithDefaults() *LockAccountFailureLockoutActionResponse {
	this := LockAccountFailureLockoutActionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *LockAccountFailureLockoutActionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LockAccountFailureLockoutActionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LockAccountFailureLockoutActionResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *LockAccountFailureLockoutActionResponse) GetSchemas() []EnumlockAccountFailureLockoutActionSchemaUrn {
	if o == nil {
		var ret []EnumlockAccountFailureLockoutActionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *LockAccountFailureLockoutActionResponse) GetSchemasOk() ([]EnumlockAccountFailureLockoutActionSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *LockAccountFailureLockoutActionResponse) SetSchemas(v []EnumlockAccountFailureLockoutActionSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LockAccountFailureLockoutActionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockAccountFailureLockoutActionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LockAccountFailureLockoutActionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LockAccountFailureLockoutActionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LockAccountFailureLockoutActionResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockAccountFailureLockoutActionResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LockAccountFailureLockoutActionResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *LockAccountFailureLockoutActionResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *LockAccountFailureLockoutActionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockAccountFailureLockoutActionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *LockAccountFailureLockoutActionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *LockAccountFailureLockoutActionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o LockAccountFailureLockoutActionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LockAccountFailureLockoutActionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableLockAccountFailureLockoutActionResponse struct {
	value *LockAccountFailureLockoutActionResponse
	isSet bool
}

func (v NullableLockAccountFailureLockoutActionResponse) Get() *LockAccountFailureLockoutActionResponse {
	return v.value
}

func (v *NullableLockAccountFailureLockoutActionResponse) Set(val *LockAccountFailureLockoutActionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLockAccountFailureLockoutActionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLockAccountFailureLockoutActionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLockAccountFailureLockoutActionResponse(val *LockAccountFailureLockoutActionResponse) *NullableLockAccountFailureLockoutActionResponse {
	return &NullableLockAccountFailureLockoutActionResponse{value: val, isSet: true}
}

func (v NullableLockAccountFailureLockoutActionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLockAccountFailureLockoutActionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
