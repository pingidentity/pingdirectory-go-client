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

// checks if the ResultCodeMapResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultCodeMapResponse{}

// ResultCodeMapResponse struct for ResultCodeMapResponse
type ResultCodeMapResponse struct {
	// Name of the Result Code Map
	Id      string                       `json:"id"`
	Schemas []EnumresultCodeMapSchemaUrn `json:"schemas,omitempty"`
	// A description for this Result Code Map
	Description *string `json:"description,omitempty"`
	// Specifies the result code that should be returned if a bind attempt fails because the user's account is locked as a result of too many failed authentication attempts.
	BindAccountLockedResultCode *int32 `json:"bindAccountLockedResultCode,omitempty"`
	// Specifies the result code that should be returned if a bind attempt fails because the target user entry does not exist in the server.
	BindMissingUserResultCode *int32 `json:"bindMissingUserResultCode,omitempty"`
	// Specifies the result code that should be returned if a password-based bind attempt fails because the target user entry does not have a password.
	BindMissingPasswordResultCode *int32 `json:"bindMissingPasswordResultCode,omitempty"`
	// Specifies the result code that should be returned if a generic error occurs within the server.
	ServerErrorResultCode                         *int32                                             `json:"serverErrorResultCode,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewResultCodeMapResponse instantiates a new ResultCodeMapResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultCodeMapResponse(id string) *ResultCodeMapResponse {
	this := ResultCodeMapResponse{}
	this.Id = id
	return &this
}

// NewResultCodeMapResponseWithDefaults instantiates a new ResultCodeMapResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultCodeMapResponseWithDefaults() *ResultCodeMapResponse {
	this := ResultCodeMapResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ResultCodeMapResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResultCodeMapResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResultCodeMapResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ResultCodeMapResponse) GetSchemas() []EnumresultCodeMapSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumresultCodeMapSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCodeMapResponse) GetSchemasOk() ([]EnumresultCodeMapSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ResultCodeMapResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumresultCodeMapSchemaUrn and assigns it to the Schemas field.
func (o *ResultCodeMapResponse) SetSchemas(v []EnumresultCodeMapSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResultCodeMapResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCodeMapResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResultCodeMapResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResultCodeMapResponse) SetDescription(v string) {
	o.Description = &v
}

// GetBindAccountLockedResultCode returns the BindAccountLockedResultCode field value if set, zero value otherwise.
func (o *ResultCodeMapResponse) GetBindAccountLockedResultCode() int32 {
	if o == nil || IsNil(o.BindAccountLockedResultCode) {
		var ret int32
		return ret
	}
	return *o.BindAccountLockedResultCode
}

// GetBindAccountLockedResultCodeOk returns a tuple with the BindAccountLockedResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCodeMapResponse) GetBindAccountLockedResultCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.BindAccountLockedResultCode) {
		return nil, false
	}
	return o.BindAccountLockedResultCode, true
}

// HasBindAccountLockedResultCode returns a boolean if a field has been set.
func (o *ResultCodeMapResponse) HasBindAccountLockedResultCode() bool {
	if o != nil && !IsNil(o.BindAccountLockedResultCode) {
		return true
	}

	return false
}

// SetBindAccountLockedResultCode gets a reference to the given int32 and assigns it to the BindAccountLockedResultCode field.
func (o *ResultCodeMapResponse) SetBindAccountLockedResultCode(v int32) {
	o.BindAccountLockedResultCode = &v
}

// GetBindMissingUserResultCode returns the BindMissingUserResultCode field value if set, zero value otherwise.
func (o *ResultCodeMapResponse) GetBindMissingUserResultCode() int32 {
	if o == nil || IsNil(o.BindMissingUserResultCode) {
		var ret int32
		return ret
	}
	return *o.BindMissingUserResultCode
}

// GetBindMissingUserResultCodeOk returns a tuple with the BindMissingUserResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCodeMapResponse) GetBindMissingUserResultCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.BindMissingUserResultCode) {
		return nil, false
	}
	return o.BindMissingUserResultCode, true
}

// HasBindMissingUserResultCode returns a boolean if a field has been set.
func (o *ResultCodeMapResponse) HasBindMissingUserResultCode() bool {
	if o != nil && !IsNil(o.BindMissingUserResultCode) {
		return true
	}

	return false
}

// SetBindMissingUserResultCode gets a reference to the given int32 and assigns it to the BindMissingUserResultCode field.
func (o *ResultCodeMapResponse) SetBindMissingUserResultCode(v int32) {
	o.BindMissingUserResultCode = &v
}

// GetBindMissingPasswordResultCode returns the BindMissingPasswordResultCode field value if set, zero value otherwise.
func (o *ResultCodeMapResponse) GetBindMissingPasswordResultCode() int32 {
	if o == nil || IsNil(o.BindMissingPasswordResultCode) {
		var ret int32
		return ret
	}
	return *o.BindMissingPasswordResultCode
}

// GetBindMissingPasswordResultCodeOk returns a tuple with the BindMissingPasswordResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCodeMapResponse) GetBindMissingPasswordResultCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.BindMissingPasswordResultCode) {
		return nil, false
	}
	return o.BindMissingPasswordResultCode, true
}

// HasBindMissingPasswordResultCode returns a boolean if a field has been set.
func (o *ResultCodeMapResponse) HasBindMissingPasswordResultCode() bool {
	if o != nil && !IsNil(o.BindMissingPasswordResultCode) {
		return true
	}

	return false
}

// SetBindMissingPasswordResultCode gets a reference to the given int32 and assigns it to the BindMissingPasswordResultCode field.
func (o *ResultCodeMapResponse) SetBindMissingPasswordResultCode(v int32) {
	o.BindMissingPasswordResultCode = &v
}

// GetServerErrorResultCode returns the ServerErrorResultCode field value if set, zero value otherwise.
func (o *ResultCodeMapResponse) GetServerErrorResultCode() int32 {
	if o == nil || IsNil(o.ServerErrorResultCode) {
		var ret int32
		return ret
	}
	return *o.ServerErrorResultCode
}

// GetServerErrorResultCodeOk returns a tuple with the ServerErrorResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCodeMapResponse) GetServerErrorResultCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerErrorResultCode) {
		return nil, false
	}
	return o.ServerErrorResultCode, true
}

// HasServerErrorResultCode returns a boolean if a field has been set.
func (o *ResultCodeMapResponse) HasServerErrorResultCode() bool {
	if o != nil && !IsNil(o.ServerErrorResultCode) {
		return true
	}

	return false
}

// SetServerErrorResultCode gets a reference to the given int32 and assigns it to the ServerErrorResultCode field.
func (o *ResultCodeMapResponse) SetServerErrorResultCode(v int32) {
	o.ServerErrorResultCode = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ResultCodeMapResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCodeMapResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ResultCodeMapResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ResultCodeMapResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ResultCodeMapResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCodeMapResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ResultCodeMapResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ResultCodeMapResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o ResultCodeMapResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultCodeMapResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.BindAccountLockedResultCode) {
		toSerialize["bindAccountLockedResultCode"] = o.BindAccountLockedResultCode
	}
	if !IsNil(o.BindMissingUserResultCode) {
		toSerialize["bindMissingUserResultCode"] = o.BindMissingUserResultCode
	}
	if !IsNil(o.BindMissingPasswordResultCode) {
		toSerialize["bindMissingPasswordResultCode"] = o.BindMissingPasswordResultCode
	}
	if !IsNil(o.ServerErrorResultCode) {
		toSerialize["serverErrorResultCode"] = o.ServerErrorResultCode
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableResultCodeMapResponse struct {
	value *ResultCodeMapResponse
	isSet bool
}

func (v NullableResultCodeMapResponse) Get() *ResultCodeMapResponse {
	return v.value
}

func (v *NullableResultCodeMapResponse) Set(val *ResultCodeMapResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResultCodeMapResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResultCodeMapResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultCodeMapResponse(val *ResultCodeMapResponse) *NullableResultCodeMapResponse {
	return &NullableResultCodeMapResponse{value: val, isSet: true}
}

func (v NullableResultCodeMapResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultCodeMapResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
