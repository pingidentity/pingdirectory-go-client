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

// checks if the GroovyScriptedHttpServletExtensionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroovyScriptedHttpServletExtensionResponse{}

// GroovyScriptedHttpServletExtensionResponse struct for GroovyScriptedHttpServletExtensionResponse
type GroovyScriptedHttpServletExtensionResponse struct {
	Schemas []EnumgroovyScriptedHttpServletExtensionSchemaUrn `json:"schemas"`
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
	CorrelationIDResponseHeader                   *string                                            `json:"correlationIDResponseHeader,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the HTTP Servlet Extension
	Id string `json:"id"`
}

// NewGroovyScriptedHttpServletExtensionResponse instantiates a new GroovyScriptedHttpServletExtensionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroovyScriptedHttpServletExtensionResponse(schemas []EnumgroovyScriptedHttpServletExtensionSchemaUrn, scriptClass string, id string) *GroovyScriptedHttpServletExtensionResponse {
	this := GroovyScriptedHttpServletExtensionResponse{}
	this.Schemas = schemas
	this.ScriptClass = scriptClass
	this.Id = id
	return &this
}

// NewGroovyScriptedHttpServletExtensionResponseWithDefaults instantiates a new GroovyScriptedHttpServletExtensionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroovyScriptedHttpServletExtensionResponseWithDefaults() *GroovyScriptedHttpServletExtensionResponse {
	this := GroovyScriptedHttpServletExtensionResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *GroovyScriptedHttpServletExtensionResponse) GetSchemas() []EnumgroovyScriptedHttpServletExtensionSchemaUrn {
	if o == nil {
		var ret []EnumgroovyScriptedHttpServletExtensionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) GetSchemasOk() ([]EnumgroovyScriptedHttpServletExtensionSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GroovyScriptedHttpServletExtensionResponse) SetSchemas(v []EnumgroovyScriptedHttpServletExtensionSchemaUrn) {
	o.Schemas = v
}

// GetScriptClass returns the ScriptClass field value
func (o *GroovyScriptedHttpServletExtensionResponse) GetScriptClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptClass
}

// GetScriptClassOk returns a tuple with the ScriptClass field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) GetScriptClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptClass, true
}

// SetScriptClass sets field value
func (o *GroovyScriptedHttpServletExtensionResponse) SetScriptClass(v string) {
	o.ScriptClass = v
}

// GetScriptArgument returns the ScriptArgument field value if set, zero value otherwise.
func (o *GroovyScriptedHttpServletExtensionResponse) GetScriptArgument() []string {
	if o == nil || IsNil(o.ScriptArgument) {
		var ret []string
		return ret
	}
	return o.ScriptArgument
}

// GetScriptArgumentOk returns a tuple with the ScriptArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) GetScriptArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ScriptArgument) {
		return nil, false
	}
	return o.ScriptArgument, true
}

// HasScriptArgument returns a boolean if a field has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) HasScriptArgument() bool {
	if o != nil && !IsNil(o.ScriptArgument) {
		return true
	}

	return false
}

// SetScriptArgument gets a reference to the given []string and assigns it to the ScriptArgument field.
func (o *GroovyScriptedHttpServletExtensionResponse) SetScriptArgument(v []string) {
	o.ScriptArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroovyScriptedHttpServletExtensionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroovyScriptedHttpServletExtensionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCrossOriginPolicy returns the CrossOriginPolicy field value if set, zero value otherwise.
func (o *GroovyScriptedHttpServletExtensionResponse) GetCrossOriginPolicy() string {
	if o == nil || IsNil(o.CrossOriginPolicy) {
		var ret string
		return ret
	}
	return *o.CrossOriginPolicy
}

// GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.CrossOriginPolicy) {
		return nil, false
	}
	return o.CrossOriginPolicy, true
}

// HasCrossOriginPolicy returns a boolean if a field has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) HasCrossOriginPolicy() bool {
	if o != nil && !IsNil(o.CrossOriginPolicy) {
		return true
	}

	return false
}

// SetCrossOriginPolicy gets a reference to the given string and assigns it to the CrossOriginPolicy field.
func (o *GroovyScriptedHttpServletExtensionResponse) SetCrossOriginPolicy(v string) {
	o.CrossOriginPolicy = &v
}

// GetResponseHeader returns the ResponseHeader field value if set, zero value otherwise.
func (o *GroovyScriptedHttpServletExtensionResponse) GetResponseHeader() []string {
	if o == nil || IsNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || IsNil(o.ResponseHeader) {
		return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) HasResponseHeader() bool {
	if o != nil && !IsNil(o.ResponseHeader) {
		return true
	}

	return false
}

// SetResponseHeader gets a reference to the given []string and assigns it to the ResponseHeader field.
func (o *GroovyScriptedHttpServletExtensionResponse) SetResponseHeader(v []string) {
	o.ResponseHeader = v
}

// GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field value if set, zero value otherwise.
func (o *GroovyScriptedHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string {
	if o == nil || IsNil(o.CorrelationIDResponseHeader) {
		var ret string
		return ret
	}
	return *o.CorrelationIDResponseHeader
}

// GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationIDResponseHeader) {
		return nil, false
	}
	return o.CorrelationIDResponseHeader, true
}

// HasCorrelationIDResponseHeader returns a boolean if a field has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool {
	if o != nil && !IsNil(o.CorrelationIDResponseHeader) {
		return true
	}

	return false
}

// SetCorrelationIDResponseHeader gets a reference to the given string and assigns it to the CorrelationIDResponseHeader field.
func (o *GroovyScriptedHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string) {
	o.CorrelationIDResponseHeader = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GroovyScriptedHttpServletExtensionResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GroovyScriptedHttpServletExtensionResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GroovyScriptedHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GroovyScriptedHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *GroovyScriptedHttpServletExtensionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedHttpServletExtensionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroovyScriptedHttpServletExtensionResponse) SetId(v string) {
	o.Id = v
}

func (o GroovyScriptedHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroovyScriptedHttpServletExtensionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["scriptClass"] = o.ScriptClass
	if !IsNil(o.ScriptArgument) {
		toSerialize["scriptArgument"] = o.ScriptArgument
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CrossOriginPolicy) {
		toSerialize["crossOriginPolicy"] = o.CrossOriginPolicy
	}
	if !IsNil(o.ResponseHeader) {
		toSerialize["responseHeader"] = o.ResponseHeader
	}
	if !IsNil(o.CorrelationIDResponseHeader) {
		toSerialize["correlationIDResponseHeader"] = o.CorrelationIDResponseHeader
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

type NullableGroovyScriptedHttpServletExtensionResponse struct {
	value *GroovyScriptedHttpServletExtensionResponse
	isSet bool
}

func (v NullableGroovyScriptedHttpServletExtensionResponse) Get() *GroovyScriptedHttpServletExtensionResponse {
	return v.value
}

func (v *NullableGroovyScriptedHttpServletExtensionResponse) Set(val *GroovyScriptedHttpServletExtensionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroovyScriptedHttpServletExtensionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroovyScriptedHttpServletExtensionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroovyScriptedHttpServletExtensionResponse(val *GroovyScriptedHttpServletExtensionResponse) *NullableGroovyScriptedHttpServletExtensionResponse {
	return &NullableGroovyScriptedHttpServletExtensionResponse{value: val, isSet: true}
}

func (v NullableGroovyScriptedHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroovyScriptedHttpServletExtensionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
