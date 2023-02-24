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

// checks if the ThirdPartyLogFileRotationListenerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyLogFileRotationListenerResponse{}

// ThirdPartyLogFileRotationListenerResponse struct for ThirdPartyLogFileRotationListenerResponse
type ThirdPartyLogFileRotationListenerResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Log File Rotation Listener
	Id      string                                           `json:"id"`
	Schemas []EnumthirdPartyLogFileRotationListenerSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Log File Rotation Listener.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Log File Rotation Listener. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Log File Rotation Listener
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log File Rotation Listener is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewThirdPartyLogFileRotationListenerResponse instantiates a new ThirdPartyLogFileRotationListenerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyLogFileRotationListenerResponse(id string, schemas []EnumthirdPartyLogFileRotationListenerSchemaUrn, extensionClass string, enabled bool) *ThirdPartyLogFileRotationListenerResponse {
	this := ThirdPartyLogFileRotationListenerResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewThirdPartyLogFileRotationListenerResponseWithDefaults instantiates a new ThirdPartyLogFileRotationListenerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyLogFileRotationListenerResponseWithDefaults() *ThirdPartyLogFileRotationListenerResponse {
	this := ThirdPartyLogFileRotationListenerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ThirdPartyLogFileRotationListenerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyLogFileRotationListenerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ThirdPartyLogFileRotationListenerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ThirdPartyLogFileRotationListenerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ThirdPartyLogFileRotationListenerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyLogFileRotationListenerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ThirdPartyLogFileRotationListenerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ThirdPartyLogFileRotationListenerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ThirdPartyLogFileRotationListenerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyLogFileRotationListenerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThirdPartyLogFileRotationListenerResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ThirdPartyLogFileRotationListenerResponse) GetSchemas() []EnumthirdPartyLogFileRotationListenerSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyLogFileRotationListenerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyLogFileRotationListenerResponse) GetSchemasOk() ([]EnumthirdPartyLogFileRotationListenerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ThirdPartyLogFileRotationListenerResponse) SetSchemas(v []EnumthirdPartyLogFileRotationListenerSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartyLogFileRotationListenerResponse) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyLogFileRotationListenerResponse) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartyLogFileRotationListenerResponse) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartyLogFileRotationListenerResponse) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyLogFileRotationListenerResponse) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartyLogFileRotationListenerResponse) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartyLogFileRotationListenerResponse) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartyLogFileRotationListenerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyLogFileRotationListenerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyLogFileRotationListenerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartyLogFileRotationListenerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyLogFileRotationListenerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyLogFileRotationListenerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyLogFileRotationListenerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ThirdPartyLogFileRotationListenerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyLogFileRotationListenerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableThirdPartyLogFileRotationListenerResponse struct {
	value *ThirdPartyLogFileRotationListenerResponse
	isSet bool
}

func (v NullableThirdPartyLogFileRotationListenerResponse) Get() *ThirdPartyLogFileRotationListenerResponse {
	return v.value
}

func (v *NullableThirdPartyLogFileRotationListenerResponse) Set(val *ThirdPartyLogFileRotationListenerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyLogFileRotationListenerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyLogFileRotationListenerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyLogFileRotationListenerResponse(val *ThirdPartyLogFileRotationListenerResponse) *NullableThirdPartyLogFileRotationListenerResponse {
	return &NullableThirdPartyLogFileRotationListenerResponse{value: val, isSet: true}
}

func (v NullableThirdPartyLogFileRotationListenerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyLogFileRotationListenerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
