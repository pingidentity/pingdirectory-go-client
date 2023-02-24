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

// checks if the ThirdPartyTrustManagerProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyTrustManagerProviderResponse{}

// ThirdPartyTrustManagerProviderResponse struct for ThirdPartyTrustManagerProviderResponse
type ThirdPartyTrustManagerProviderResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Trust Manager Provider
	Id      string                                        `json:"id"`
	Schemas []EnumthirdPartyTrustManagerProviderSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Trust Manager Provider.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Trust Manager Provider. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// Indicate whether the Trust Manager Provider is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether certificates issued by an authority included in the JVM's set of default issuers should be automatically trusted, even if they would not otherwise be trusted by this provider.
	IncludeJVMDefaultIssuers *bool `json:"includeJVMDefaultIssuers,omitempty"`
}

// NewThirdPartyTrustManagerProviderResponse instantiates a new ThirdPartyTrustManagerProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyTrustManagerProviderResponse(id string, schemas []EnumthirdPartyTrustManagerProviderSchemaUrn, extensionClass string, enabled bool) *ThirdPartyTrustManagerProviderResponse {
	this := ThirdPartyTrustManagerProviderResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewThirdPartyTrustManagerProviderResponseWithDefaults instantiates a new ThirdPartyTrustManagerProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyTrustManagerProviderResponseWithDefaults() *ThirdPartyTrustManagerProviderResponse {
	this := ThirdPartyTrustManagerProviderResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ThirdPartyTrustManagerProviderResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyTrustManagerProviderResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ThirdPartyTrustManagerProviderResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ThirdPartyTrustManagerProviderResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ThirdPartyTrustManagerProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyTrustManagerProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ThirdPartyTrustManagerProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ThirdPartyTrustManagerProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ThirdPartyTrustManagerProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyTrustManagerProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThirdPartyTrustManagerProviderResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ThirdPartyTrustManagerProviderResponse) GetSchemas() []EnumthirdPartyTrustManagerProviderSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyTrustManagerProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyTrustManagerProviderResponse) GetSchemasOk() ([]EnumthirdPartyTrustManagerProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ThirdPartyTrustManagerProviderResponse) SetSchemas(v []EnumthirdPartyTrustManagerProviderSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartyTrustManagerProviderResponse) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyTrustManagerProviderResponse) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartyTrustManagerProviderResponse) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartyTrustManagerProviderResponse) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyTrustManagerProviderResponse) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartyTrustManagerProviderResponse) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartyTrustManagerProviderResponse) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyTrustManagerProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyTrustManagerProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyTrustManagerProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIncludeJVMDefaultIssuers returns the IncludeJVMDefaultIssuers field value if set, zero value otherwise.
func (o *ThirdPartyTrustManagerProviderResponse) GetIncludeJVMDefaultIssuers() bool {
	if o == nil || IsNil(o.IncludeJVMDefaultIssuers) {
		var ret bool
		return ret
	}
	return *o.IncludeJVMDefaultIssuers
}

// GetIncludeJVMDefaultIssuersOk returns a tuple with the IncludeJVMDefaultIssuers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyTrustManagerProviderResponse) GetIncludeJVMDefaultIssuersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeJVMDefaultIssuers) {
		return nil, false
	}
	return o.IncludeJVMDefaultIssuers, true
}

// HasIncludeJVMDefaultIssuers returns a boolean if a field has been set.
func (o *ThirdPartyTrustManagerProviderResponse) HasIncludeJVMDefaultIssuers() bool {
	if o != nil && !IsNil(o.IncludeJVMDefaultIssuers) {
		return true
	}

	return false
}

// SetIncludeJVMDefaultIssuers gets a reference to the given bool and assigns it to the IncludeJVMDefaultIssuers field.
func (o *ThirdPartyTrustManagerProviderResponse) SetIncludeJVMDefaultIssuers(v bool) {
	o.IncludeJVMDefaultIssuers = &v
}

func (o ThirdPartyTrustManagerProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyTrustManagerProviderResponse) ToMap() (map[string]interface{}, error) {
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
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.IncludeJVMDefaultIssuers) {
		toSerialize["includeJVMDefaultIssuers"] = o.IncludeJVMDefaultIssuers
	}
	return toSerialize, nil
}

type NullableThirdPartyTrustManagerProviderResponse struct {
	value *ThirdPartyTrustManagerProviderResponse
	isSet bool
}

func (v NullableThirdPartyTrustManagerProviderResponse) Get() *ThirdPartyTrustManagerProviderResponse {
	return v.value
}

func (v *NullableThirdPartyTrustManagerProviderResponse) Set(val *ThirdPartyTrustManagerProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyTrustManagerProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyTrustManagerProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyTrustManagerProviderResponse(val *ThirdPartyTrustManagerProviderResponse) *NullableThirdPartyTrustManagerProviderResponse {
	return &NullableThirdPartyTrustManagerProviderResponse{value: val, isSet: true}
}

func (v NullableThirdPartyTrustManagerProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyTrustManagerProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
