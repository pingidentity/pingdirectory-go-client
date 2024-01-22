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

// checks if the AddThirdPartyTrustManagerProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddThirdPartyTrustManagerProviderRequest{}

// AddThirdPartyTrustManagerProviderRequest struct for AddThirdPartyTrustManagerProviderRequest
type AddThirdPartyTrustManagerProviderRequest struct {
	Schemas []EnumthirdPartyTrustManagerProviderSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Trust Manager Provider.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Trust Manager Provider. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// Indicate whether the Trust Manager Provider is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether certificates issued by an authority included in the JVM's set of default issuers should be automatically trusted, even if they would not otherwise be trusted by this provider.
	IncludeJVMDefaultIssuers *bool `json:"includeJVMDefaultIssuers,omitempty"`
	// Name of the new Trust Manager Provider
	ProviderName string `json:"providerName"`
}

// NewAddThirdPartyTrustManagerProviderRequest instantiates a new AddThirdPartyTrustManagerProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyTrustManagerProviderRequest(schemas []EnumthirdPartyTrustManagerProviderSchemaUrn, extensionClass string, enabled bool, providerName string) *AddThirdPartyTrustManagerProviderRequest {
	this := AddThirdPartyTrustManagerProviderRequest{}
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	this.ProviderName = providerName
	return &this
}

// NewAddThirdPartyTrustManagerProviderRequestWithDefaults instantiates a new AddThirdPartyTrustManagerProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyTrustManagerProviderRequestWithDefaults() *AddThirdPartyTrustManagerProviderRequest {
	this := AddThirdPartyTrustManagerProviderRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyTrustManagerProviderRequest) GetSchemas() []EnumthirdPartyTrustManagerProviderSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyTrustManagerProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyTrustManagerProviderRequest) GetSchemasOk() ([]EnumthirdPartyTrustManagerProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyTrustManagerProviderRequest) SetSchemas(v []EnumthirdPartyTrustManagerProviderSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyTrustManagerProviderRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyTrustManagerProviderRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyTrustManagerProviderRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyTrustManagerProviderRequest) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyTrustManagerProviderRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyTrustManagerProviderRequest) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyTrustManagerProviderRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyTrustManagerProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyTrustManagerProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyTrustManagerProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIncludeJVMDefaultIssuers returns the IncludeJVMDefaultIssuers field value if set, zero value otherwise.
func (o *AddThirdPartyTrustManagerProviderRequest) GetIncludeJVMDefaultIssuers() bool {
	if o == nil || IsNil(o.IncludeJVMDefaultIssuers) {
		var ret bool
		return ret
	}
	return *o.IncludeJVMDefaultIssuers
}

// GetIncludeJVMDefaultIssuersOk returns a tuple with the IncludeJVMDefaultIssuers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyTrustManagerProviderRequest) GetIncludeJVMDefaultIssuersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeJVMDefaultIssuers) {
		return nil, false
	}
	return o.IncludeJVMDefaultIssuers, true
}

// HasIncludeJVMDefaultIssuers returns a boolean if a field has been set.
func (o *AddThirdPartyTrustManagerProviderRequest) HasIncludeJVMDefaultIssuers() bool {
	if o != nil && !IsNil(o.IncludeJVMDefaultIssuers) {
		return true
	}

	return false
}

// SetIncludeJVMDefaultIssuers gets a reference to the given bool and assigns it to the IncludeJVMDefaultIssuers field.
func (o *AddThirdPartyTrustManagerProviderRequest) SetIncludeJVMDefaultIssuers(v bool) {
	o.IncludeJVMDefaultIssuers = &v
}

// GetProviderName returns the ProviderName field value
func (o *AddThirdPartyTrustManagerProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyTrustManagerProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddThirdPartyTrustManagerProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

func (o AddThirdPartyTrustManagerProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddThirdPartyTrustManagerProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.IncludeJVMDefaultIssuers) {
		toSerialize["includeJVMDefaultIssuers"] = o.IncludeJVMDefaultIssuers
	}
	toSerialize["providerName"] = o.ProviderName
	return toSerialize, nil
}

type NullableAddThirdPartyTrustManagerProviderRequest struct {
	value *AddThirdPartyTrustManagerProviderRequest
	isSet bool
}

func (v NullableAddThirdPartyTrustManagerProviderRequest) Get() *AddThirdPartyTrustManagerProviderRequest {
	return v.value
}

func (v *NullableAddThirdPartyTrustManagerProviderRequest) Set(val *AddThirdPartyTrustManagerProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyTrustManagerProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyTrustManagerProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyTrustManagerProviderRequest(val *AddThirdPartyTrustManagerProviderRequest) *NullableAddThirdPartyTrustManagerProviderRequest {
	return &NullableAddThirdPartyTrustManagerProviderRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyTrustManagerProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyTrustManagerProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
