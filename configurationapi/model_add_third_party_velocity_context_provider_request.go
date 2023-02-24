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

// checks if the AddThirdPartyVelocityContextProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddThirdPartyVelocityContextProviderRequest{}

// AddThirdPartyVelocityContextProviderRequest struct for AddThirdPartyVelocityContextProviderRequest
type AddThirdPartyVelocityContextProviderRequest struct {
	// Name of the new Velocity Context Provider
	ProviderName string                                           `json:"providerName"`
	Schemas      []EnumthirdPartyVelocityContextProviderSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Velocity Context Provider.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Velocity Context Provider. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// Indicates whether this Velocity Context Provider is enabled. If set to 'false' this Velocity Context Provider will not contribute context content for any requests.
	Enabled     *bool                                       `json:"enabled,omitempty"`
	ObjectScope *EnumvelocityContextProviderObjectScopeProp `json:"objectScope,omitempty"`
	// The name of a view for which this Velocity Context Provider will contribute content.
	IncludedView []string `json:"includedView,omitempty"`
	// The name of a view for which this Velocity Context Provider will not contribute content.
	ExcludedView []string `json:"excludedView,omitempty"`
	// Specifies the set of HTTP methods handled by this Velocity Context Provider, which will perform actions necessary to fulfill the request before updating the context for the response. The values of this property are not case-sensitive.
	HttpMethod []string `json:"httpMethod,omitempty"`
	// Specifies HTTP header fields and values added to response headers for template page requests to which this Velocity Context Provider contributes content.
	ResponseHeader []string `json:"responseHeader,omitempty"`
}

// NewAddThirdPartyVelocityContextProviderRequest instantiates a new AddThirdPartyVelocityContextProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyVelocityContextProviderRequest(providerName string, schemas []EnumthirdPartyVelocityContextProviderSchemaUrn, extensionClass string) *AddThirdPartyVelocityContextProviderRequest {
	this := AddThirdPartyVelocityContextProviderRequest{}
	this.ProviderName = providerName
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	return &this
}

// NewAddThirdPartyVelocityContextProviderRequestWithDefaults instantiates a new AddThirdPartyVelocityContextProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyVelocityContextProviderRequestWithDefaults() *AddThirdPartyVelocityContextProviderRequest {
	this := AddThirdPartyVelocityContextProviderRequest{}
	return &this
}

// GetProviderName returns the ProviderName field value
func (o *AddThirdPartyVelocityContextProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddThirdPartyVelocityContextProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyVelocityContextProviderRequest) GetSchemas() []EnumthirdPartyVelocityContextProviderSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyVelocityContextProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) GetSchemasOk() ([]EnumthirdPartyVelocityContextProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyVelocityContextProviderRequest) SetSchemas(v []EnumthirdPartyVelocityContextProviderSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyVelocityContextProviderRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyVelocityContextProviderRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyVelocityContextProviderRequest) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyVelocityContextProviderRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AddThirdPartyVelocityContextProviderRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AddThirdPartyVelocityContextProviderRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetObjectScope returns the ObjectScope field value if set, zero value otherwise.
func (o *AddThirdPartyVelocityContextProviderRequest) GetObjectScope() EnumvelocityContextProviderObjectScopeProp {
	if o == nil || IsNil(o.ObjectScope) {
		var ret EnumvelocityContextProviderObjectScopeProp
		return ret
	}
	return *o.ObjectScope
}

// GetObjectScopeOk returns a tuple with the ObjectScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) GetObjectScopeOk() (*EnumvelocityContextProviderObjectScopeProp, bool) {
	if o == nil || IsNil(o.ObjectScope) {
		return nil, false
	}
	return o.ObjectScope, true
}

// HasObjectScope returns a boolean if a field has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) HasObjectScope() bool {
	if o != nil && !IsNil(o.ObjectScope) {
		return true
	}

	return false
}

// SetObjectScope gets a reference to the given EnumvelocityContextProviderObjectScopeProp and assigns it to the ObjectScope field.
func (o *AddThirdPartyVelocityContextProviderRequest) SetObjectScope(v EnumvelocityContextProviderObjectScopeProp) {
	o.ObjectScope = &v
}

// GetIncludedView returns the IncludedView field value if set, zero value otherwise.
func (o *AddThirdPartyVelocityContextProviderRequest) GetIncludedView() []string {
	if o == nil || IsNil(o.IncludedView) {
		var ret []string
		return ret
	}
	return o.IncludedView
}

// GetIncludedViewOk returns a tuple with the IncludedView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) GetIncludedViewOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedView) {
		return nil, false
	}
	return o.IncludedView, true
}

// HasIncludedView returns a boolean if a field has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) HasIncludedView() bool {
	if o != nil && !IsNil(o.IncludedView) {
		return true
	}

	return false
}

// SetIncludedView gets a reference to the given []string and assigns it to the IncludedView field.
func (o *AddThirdPartyVelocityContextProviderRequest) SetIncludedView(v []string) {
	o.IncludedView = v
}

// GetExcludedView returns the ExcludedView field value if set, zero value otherwise.
func (o *AddThirdPartyVelocityContextProviderRequest) GetExcludedView() []string {
	if o == nil || IsNil(o.ExcludedView) {
		var ret []string
		return ret
	}
	return o.ExcludedView
}

// GetExcludedViewOk returns a tuple with the ExcludedView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) GetExcludedViewOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludedView) {
		return nil, false
	}
	return o.ExcludedView, true
}

// HasExcludedView returns a boolean if a field has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) HasExcludedView() bool {
	if o != nil && !IsNil(o.ExcludedView) {
		return true
	}

	return false
}

// SetExcludedView gets a reference to the given []string and assigns it to the ExcludedView field.
func (o *AddThirdPartyVelocityContextProviderRequest) SetExcludedView(v []string) {
	o.ExcludedView = v
}

// GetHttpMethod returns the HttpMethod field value if set, zero value otherwise.
func (o *AddThirdPartyVelocityContextProviderRequest) GetHttpMethod() []string {
	if o == nil || IsNil(o.HttpMethod) {
		var ret []string
		return ret
	}
	return o.HttpMethod
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) GetHttpMethodOk() ([]string, bool) {
	if o == nil || IsNil(o.HttpMethod) {
		return nil, false
	}
	return o.HttpMethod, true
}

// HasHttpMethod returns a boolean if a field has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) HasHttpMethod() bool {
	if o != nil && !IsNil(o.HttpMethod) {
		return true
	}

	return false
}

// SetHttpMethod gets a reference to the given []string and assigns it to the HttpMethod field.
func (o *AddThirdPartyVelocityContextProviderRequest) SetHttpMethod(v []string) {
	o.HttpMethod = v
}

// GetResponseHeader returns the ResponseHeader field value if set, zero value otherwise.
func (o *AddThirdPartyVelocityContextProviderRequest) GetResponseHeader() []string {
	if o == nil || IsNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || IsNil(o.ResponseHeader) {
		return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *AddThirdPartyVelocityContextProviderRequest) HasResponseHeader() bool {
	if o != nil && !IsNil(o.ResponseHeader) {
		return true
	}

	return false
}

// SetResponseHeader gets a reference to the given []string and assigns it to the ResponseHeader field.
func (o *AddThirdPartyVelocityContextProviderRequest) SetResponseHeader(v []string) {
	o.ResponseHeader = v
}

func (o AddThirdPartyVelocityContextProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddThirdPartyVelocityContextProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["providerName"] = o.ProviderName
	toSerialize["schemas"] = o.Schemas
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ObjectScope) {
		toSerialize["objectScope"] = o.ObjectScope
	}
	if !IsNil(o.IncludedView) {
		toSerialize["includedView"] = o.IncludedView
	}
	if !IsNil(o.ExcludedView) {
		toSerialize["excludedView"] = o.ExcludedView
	}
	if !IsNil(o.HttpMethod) {
		toSerialize["httpMethod"] = o.HttpMethod
	}
	if !IsNil(o.ResponseHeader) {
		toSerialize["responseHeader"] = o.ResponseHeader
	}
	return toSerialize, nil
}

type NullableAddThirdPartyVelocityContextProviderRequest struct {
	value *AddThirdPartyVelocityContextProviderRequest
	isSet bool
}

func (v NullableAddThirdPartyVelocityContextProviderRequest) Get() *AddThirdPartyVelocityContextProviderRequest {
	return v.value
}

func (v *NullableAddThirdPartyVelocityContextProviderRequest) Set(val *AddThirdPartyVelocityContextProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyVelocityContextProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyVelocityContextProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyVelocityContextProviderRequest(val *AddThirdPartyVelocityContextProviderRequest) *NullableAddThirdPartyVelocityContextProviderRequest {
	return &NullableAddThirdPartyVelocityContextProviderRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyVelocityContextProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyVelocityContextProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
