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

// AddGroovyScriptedHttpServletExtensionRequest struct for AddGroovyScriptedHttpServletExtensionRequest
type AddGroovyScriptedHttpServletExtensionRequest struct {
	// Name of the new HTTP Servlet Extension
	ExtensionName string                                            `json:"extensionName"`
	Schemas       []EnumgroovyScriptedHttpServletExtensionSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted HTTP Servlet Extension.
	ScriptClass string `json:"scriptClass"`
	// The set of arguments used to customize the behavior for the Scripted HTTP Servlet Extension. Each configuration property should be given in the form 'name=value'.
	ScriptArgument []string `json:"scriptArgument,omitempty"`
	// A description for this HTTP Servlet Extension
	Description *string `json:"description,omitempty"`
	// The cross-origin request policy to use for the HTTP Servlet Extension.
	CrossOriginPolicy *string `json:"crossOriginPolicy,omitempty"`
	// Specifies HTTP header fields and values added to response headers for all requests.
	ResponseHeader []string `json:"responseHeader,omitempty"`
	// Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \"Correlation-Id\", \"X-Amzn-Trace-Id\", and \"X-Request-Id\".
	CorrelationIDResponseHeader *string `json:"correlationIDResponseHeader,omitempty"`
}

// NewAddGroovyScriptedHttpServletExtensionRequest instantiates a new AddGroovyScriptedHttpServletExtensionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddGroovyScriptedHttpServletExtensionRequest(extensionName string, schemas []EnumgroovyScriptedHttpServletExtensionSchemaUrn, scriptClass string) *AddGroovyScriptedHttpServletExtensionRequest {
	this := AddGroovyScriptedHttpServletExtensionRequest{}
	this.ExtensionName = extensionName
	this.Schemas = schemas
	this.ScriptClass = scriptClass
	return &this
}

// NewAddGroovyScriptedHttpServletExtensionRequestWithDefaults instantiates a new AddGroovyScriptedHttpServletExtensionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddGroovyScriptedHttpServletExtensionRequestWithDefaults() *AddGroovyScriptedHttpServletExtensionRequest {
	this := AddGroovyScriptedHttpServletExtensionRequest{}
	return &this
}

// GetExtensionName returns the ExtensionName field value
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetExtensionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionName
}

// GetExtensionNameOk returns a tuple with the ExtensionName field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetExtensionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionName, true
}

// SetExtensionName sets field value
func (o *AddGroovyScriptedHttpServletExtensionRequest) SetExtensionName(v string) {
	o.ExtensionName = v
}

// GetSchemas returns the Schemas field value
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetSchemas() []EnumgroovyScriptedHttpServletExtensionSchemaUrn {
	if o == nil {
		var ret []EnumgroovyScriptedHttpServletExtensionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetSchemasOk() ([]EnumgroovyScriptedHttpServletExtensionSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddGroovyScriptedHttpServletExtensionRequest) SetSchemas(v []EnumgroovyScriptedHttpServletExtensionSchemaUrn) {
	o.Schemas = v
}

// GetScriptClass returns the ScriptClass field value
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetScriptClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptClass
}

// GetScriptClassOk returns a tuple with the ScriptClass field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetScriptClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptClass, true
}

// SetScriptClass sets field value
func (o *AddGroovyScriptedHttpServletExtensionRequest) SetScriptClass(v string) {
	o.ScriptClass = v
}

// GetScriptArgument returns the ScriptArgument field value if set, zero value otherwise.
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetScriptArgument() []string {
	if o == nil || isNil(o.ScriptArgument) {
		var ret []string
		return ret
	}
	return o.ScriptArgument
}

// GetScriptArgumentOk returns a tuple with the ScriptArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetScriptArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ScriptArgument) {
		return nil, false
	}
	return o.ScriptArgument, true
}

// HasScriptArgument returns a boolean if a field has been set.
func (o *AddGroovyScriptedHttpServletExtensionRequest) HasScriptArgument() bool {
	if o != nil && !isNil(o.ScriptArgument) {
		return true
	}

	return false
}

// SetScriptArgument gets a reference to the given []string and assigns it to the ScriptArgument field.
func (o *AddGroovyScriptedHttpServletExtensionRequest) SetScriptArgument(v []string) {
	o.ScriptArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddGroovyScriptedHttpServletExtensionRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddGroovyScriptedHttpServletExtensionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCrossOriginPolicy returns the CrossOriginPolicy field value if set, zero value otherwise.
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetCrossOriginPolicy() string {
	if o == nil || isNil(o.CrossOriginPolicy) {
		var ret string
		return ret
	}
	return *o.CrossOriginPolicy
}

// GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetCrossOriginPolicyOk() (*string, bool) {
	if o == nil || isNil(o.CrossOriginPolicy) {
		return nil, false
	}
	return o.CrossOriginPolicy, true
}

// HasCrossOriginPolicy returns a boolean if a field has been set.
func (o *AddGroovyScriptedHttpServletExtensionRequest) HasCrossOriginPolicy() bool {
	if o != nil && !isNil(o.CrossOriginPolicy) {
		return true
	}

	return false
}

// SetCrossOriginPolicy gets a reference to the given string and assigns it to the CrossOriginPolicy field.
func (o *AddGroovyScriptedHttpServletExtensionRequest) SetCrossOriginPolicy(v string) {
	o.CrossOriginPolicy = &v
}

// GetResponseHeader returns the ResponseHeader field value if set, zero value otherwise.
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetResponseHeader() []string {
	if o == nil || isNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || isNil(o.ResponseHeader) {
		return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *AddGroovyScriptedHttpServletExtensionRequest) HasResponseHeader() bool {
	if o != nil && !isNil(o.ResponseHeader) {
		return true
	}

	return false
}

// SetResponseHeader gets a reference to the given []string and assigns it to the ResponseHeader field.
func (o *AddGroovyScriptedHttpServletExtensionRequest) SetResponseHeader(v []string) {
	o.ResponseHeader = v
}

// GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field value if set, zero value otherwise.
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetCorrelationIDResponseHeader() string {
	if o == nil || isNil(o.CorrelationIDResponseHeader) {
		var ret string
		return ret
	}
	return *o.CorrelationIDResponseHeader
}

// GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpServletExtensionRequest) GetCorrelationIDResponseHeaderOk() (*string, bool) {
	if o == nil || isNil(o.CorrelationIDResponseHeader) {
		return nil, false
	}
	return o.CorrelationIDResponseHeader, true
}

// HasCorrelationIDResponseHeader returns a boolean if a field has been set.
func (o *AddGroovyScriptedHttpServletExtensionRequest) HasCorrelationIDResponseHeader() bool {
	if o != nil && !isNil(o.CorrelationIDResponseHeader) {
		return true
	}

	return false
}

// SetCorrelationIDResponseHeader gets a reference to the given string and assigns it to the CorrelationIDResponseHeader field.
func (o *AddGroovyScriptedHttpServletExtensionRequest) SetCorrelationIDResponseHeader(v string) {
	o.CorrelationIDResponseHeader = &v
}

func (o AddGroovyScriptedHttpServletExtensionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["extensionName"] = o.ExtensionName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["scriptClass"] = o.ScriptClass
	}
	if !isNil(o.ScriptArgument) {
		toSerialize["scriptArgument"] = o.ScriptArgument
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.CrossOriginPolicy) {
		toSerialize["crossOriginPolicy"] = o.CrossOriginPolicy
	}
	if !isNil(o.ResponseHeader) {
		toSerialize["responseHeader"] = o.ResponseHeader
	}
	if !isNil(o.CorrelationIDResponseHeader) {
		toSerialize["correlationIDResponseHeader"] = o.CorrelationIDResponseHeader
	}
	return json.Marshal(toSerialize)
}

type NullableAddGroovyScriptedHttpServletExtensionRequest struct {
	value *AddGroovyScriptedHttpServletExtensionRequest
	isSet bool
}

func (v NullableAddGroovyScriptedHttpServletExtensionRequest) Get() *AddGroovyScriptedHttpServletExtensionRequest {
	return v.value
}

func (v *NullableAddGroovyScriptedHttpServletExtensionRequest) Set(val *AddGroovyScriptedHttpServletExtensionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddGroovyScriptedHttpServletExtensionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddGroovyScriptedHttpServletExtensionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddGroovyScriptedHttpServletExtensionRequest(val *AddGroovyScriptedHttpServletExtensionRequest) *NullableAddGroovyScriptedHttpServletExtensionRequest {
	return &NullableAddGroovyScriptedHttpServletExtensionRequest{value: val, isSet: true}
}

func (v NullableAddGroovyScriptedHttpServletExtensionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddGroovyScriptedHttpServletExtensionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}