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

// checks if the DelayBindResponseFailureLockoutActionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DelayBindResponseFailureLockoutActionResponse{}

// DelayBindResponseFailureLockoutActionResponse struct for DelayBindResponseFailureLockoutActionResponse
type DelayBindResponseFailureLockoutActionResponse struct {
	Schemas []EnumdelayBindResponseFailureLockoutActionSchemaUrn `json:"schemas"`
	// The length of time to delay the bind response for accounts with too many failed authentication attempts.
	Delay string `json:"delay"`
	// Indicates whether to delay the response for authentication attempts even if that delay may block the thread being used to process the attempt.
	AllowBlockingDelay *bool `json:"allowBlockingDelay,omitempty"`
	// Indicates whether to generate an account status notification for cases in which a bind response is delayed because of failure lockout.
	GenerateAccountStatusNotification *bool `json:"generateAccountStatusNotification,omitempty"`
	// A description for this Failure Lockout Action
	Description                                   *string                                            `json:"description,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Failure Lockout Action
	Id string `json:"id"`
}

// NewDelayBindResponseFailureLockoutActionResponse instantiates a new DelayBindResponseFailureLockoutActionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelayBindResponseFailureLockoutActionResponse(schemas []EnumdelayBindResponseFailureLockoutActionSchemaUrn, delay string, id string) *DelayBindResponseFailureLockoutActionResponse {
	this := DelayBindResponseFailureLockoutActionResponse{}
	this.Schemas = schemas
	this.Delay = delay
	this.Id = id
	return &this
}

// NewDelayBindResponseFailureLockoutActionResponseWithDefaults instantiates a new DelayBindResponseFailureLockoutActionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelayBindResponseFailureLockoutActionResponseWithDefaults() *DelayBindResponseFailureLockoutActionResponse {
	this := DelayBindResponseFailureLockoutActionResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *DelayBindResponseFailureLockoutActionResponse) GetSchemas() []EnumdelayBindResponseFailureLockoutActionSchemaUrn {
	if o == nil {
		var ret []EnumdelayBindResponseFailureLockoutActionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *DelayBindResponseFailureLockoutActionResponse) GetSchemasOk() ([]EnumdelayBindResponseFailureLockoutActionSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *DelayBindResponseFailureLockoutActionResponse) SetSchemas(v []EnumdelayBindResponseFailureLockoutActionSchemaUrn) {
	o.Schemas = v
}

// GetDelay returns the Delay field value
func (o *DelayBindResponseFailureLockoutActionResponse) GetDelay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Delay
}

// GetDelayOk returns a tuple with the Delay field value
// and a boolean to check if the value has been set.
func (o *DelayBindResponseFailureLockoutActionResponse) GetDelayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delay, true
}

// SetDelay sets field value
func (o *DelayBindResponseFailureLockoutActionResponse) SetDelay(v string) {
	o.Delay = v
}

// GetAllowBlockingDelay returns the AllowBlockingDelay field value if set, zero value otherwise.
func (o *DelayBindResponseFailureLockoutActionResponse) GetAllowBlockingDelay() bool {
	if o == nil || IsNil(o.AllowBlockingDelay) {
		var ret bool
		return ret
	}
	return *o.AllowBlockingDelay
}

// GetAllowBlockingDelayOk returns a tuple with the AllowBlockingDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayBindResponseFailureLockoutActionResponse) GetAllowBlockingDelayOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowBlockingDelay) {
		return nil, false
	}
	return o.AllowBlockingDelay, true
}

// HasAllowBlockingDelay returns a boolean if a field has been set.
func (o *DelayBindResponseFailureLockoutActionResponse) HasAllowBlockingDelay() bool {
	if o != nil && !IsNil(o.AllowBlockingDelay) {
		return true
	}

	return false
}

// SetAllowBlockingDelay gets a reference to the given bool and assigns it to the AllowBlockingDelay field.
func (o *DelayBindResponseFailureLockoutActionResponse) SetAllowBlockingDelay(v bool) {
	o.AllowBlockingDelay = &v
}

// GetGenerateAccountStatusNotification returns the GenerateAccountStatusNotification field value if set, zero value otherwise.
func (o *DelayBindResponseFailureLockoutActionResponse) GetGenerateAccountStatusNotification() bool {
	if o == nil || IsNil(o.GenerateAccountStatusNotification) {
		var ret bool
		return ret
	}
	return *o.GenerateAccountStatusNotification
}

// GetGenerateAccountStatusNotificationOk returns a tuple with the GenerateAccountStatusNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayBindResponseFailureLockoutActionResponse) GetGenerateAccountStatusNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.GenerateAccountStatusNotification) {
		return nil, false
	}
	return o.GenerateAccountStatusNotification, true
}

// HasGenerateAccountStatusNotification returns a boolean if a field has been set.
func (o *DelayBindResponseFailureLockoutActionResponse) HasGenerateAccountStatusNotification() bool {
	if o != nil && !IsNil(o.GenerateAccountStatusNotification) {
		return true
	}

	return false
}

// SetGenerateAccountStatusNotification gets a reference to the given bool and assigns it to the GenerateAccountStatusNotification field.
func (o *DelayBindResponseFailureLockoutActionResponse) SetGenerateAccountStatusNotification(v bool) {
	o.GenerateAccountStatusNotification = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DelayBindResponseFailureLockoutActionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayBindResponseFailureLockoutActionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DelayBindResponseFailureLockoutActionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DelayBindResponseFailureLockoutActionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DelayBindResponseFailureLockoutActionResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayBindResponseFailureLockoutActionResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DelayBindResponseFailureLockoutActionResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *DelayBindResponseFailureLockoutActionResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *DelayBindResponseFailureLockoutActionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayBindResponseFailureLockoutActionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *DelayBindResponseFailureLockoutActionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *DelayBindResponseFailureLockoutActionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *DelayBindResponseFailureLockoutActionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DelayBindResponseFailureLockoutActionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DelayBindResponseFailureLockoutActionResponse) SetId(v string) {
	o.Id = v
}

func (o DelayBindResponseFailureLockoutActionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DelayBindResponseFailureLockoutActionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["delay"] = o.Delay
	if !IsNil(o.AllowBlockingDelay) {
		toSerialize["allowBlockingDelay"] = o.AllowBlockingDelay
	}
	if !IsNil(o.GenerateAccountStatusNotification) {
		toSerialize["generateAccountStatusNotification"] = o.GenerateAccountStatusNotification
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableDelayBindResponseFailureLockoutActionResponse struct {
	value *DelayBindResponseFailureLockoutActionResponse
	isSet bool
}

func (v NullableDelayBindResponseFailureLockoutActionResponse) Get() *DelayBindResponseFailureLockoutActionResponse {
	return v.value
}

func (v *NullableDelayBindResponseFailureLockoutActionResponse) Set(val *DelayBindResponseFailureLockoutActionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDelayBindResponseFailureLockoutActionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDelayBindResponseFailureLockoutActionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelayBindResponseFailureLockoutActionResponse(val *DelayBindResponseFailureLockoutActionResponse) *NullableDelayBindResponseFailureLockoutActionResponse {
	return &NullableDelayBindResponseFailureLockoutActionResponse{value: val, isSet: true}
}

func (v NullableDelayBindResponseFailureLockoutActionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelayBindResponseFailureLockoutActionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
