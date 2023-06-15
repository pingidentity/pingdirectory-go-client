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

// checks if the VelocityToolsVelocityContextProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VelocityToolsVelocityContextProviderResponse{}

// VelocityToolsVelocityContextProviderResponse struct for VelocityToolsVelocityContextProviderResponse
type VelocityToolsVelocityContextProviderResponse struct {
	// Name of the Velocity Context Provider
	Id      string                                              `json:"id"`
	Schemas []EnumvelocityToolsVelocityContextProviderSchemaUrn `json:"schemas"`
	// The fully-qualified name of a Velocity Tool class that will be initialized for each request. May optionally include a path to a properties file used to configure this tool separated from the class name by a semi-colon (;). The path may absolute or relative to the server root.
	RequestTool []string `json:"requestTool,omitempty"`
	// The fully-qualified name of a Velocity Tool class that will be initialized for each session. May optionally include a path to a properties file used to configure this tool separated from the class name by a semi-colon (;). The path may absolute or relative to the server root.
	SessionTool []string `json:"sessionTool,omitempty"`
	// The fully-qualified name of a Velocity Tool class that will be initialized once for the life of the server. May optionally include a path to a properties file used to configure this tool separated from the class name by a semi-colon (;). The path may absolute or relative to the server root.
	ApplicationTool []string `json:"applicationTool,omitempty"`
	// Indicates whether this Velocity Context Provider is enabled. If set to 'false' this Velocity Context Provider will not contribute context content for any requests.
	Enabled     *bool                                       `json:"enabled,omitempty"`
	ObjectScope *EnumvelocityContextProviderObjectScopeProp `json:"objectScope,omitempty"`
	// The name of a view for which this Velocity Context Provider will contribute content.
	IncludedView []string `json:"includedView,omitempty"`
	// The name of a view for which this Velocity Context Provider will not contribute content.
	ExcludedView []string `json:"excludedView,omitempty"`
	// Specifies HTTP header fields and values added to response headers for template page requests to which this Velocity Context Provider contributes content.
	ResponseHeader                                []string                                           `json:"responseHeader,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewVelocityToolsVelocityContextProviderResponse instantiates a new VelocityToolsVelocityContextProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVelocityToolsVelocityContextProviderResponse(id string, schemas []EnumvelocityToolsVelocityContextProviderSchemaUrn) *VelocityToolsVelocityContextProviderResponse {
	this := VelocityToolsVelocityContextProviderResponse{}
	this.Id = id
	this.Schemas = schemas
	return &this
}

// NewVelocityToolsVelocityContextProviderResponseWithDefaults instantiates a new VelocityToolsVelocityContextProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVelocityToolsVelocityContextProviderResponseWithDefaults() *VelocityToolsVelocityContextProviderResponse {
	this := VelocityToolsVelocityContextProviderResponse{}
	return &this
}

// GetId returns the Id field value
func (o *VelocityToolsVelocityContextProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VelocityToolsVelocityContextProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VelocityToolsVelocityContextProviderResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *VelocityToolsVelocityContextProviderResponse) GetSchemas() []EnumvelocityToolsVelocityContextProviderSchemaUrn {
	if o == nil {
		var ret []EnumvelocityToolsVelocityContextProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *VelocityToolsVelocityContextProviderResponse) GetSchemasOk() ([]EnumvelocityToolsVelocityContextProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *VelocityToolsVelocityContextProviderResponse) SetSchemas(v []EnumvelocityToolsVelocityContextProviderSchemaUrn) {
	o.Schemas = v
}

// GetRequestTool returns the RequestTool field value if set, zero value otherwise.
func (o *VelocityToolsVelocityContextProviderResponse) GetRequestTool() []string {
	if o == nil || IsNil(o.RequestTool) {
		var ret []string
		return ret
	}
	return o.RequestTool
}

// GetRequestToolOk returns a tuple with the RequestTool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityToolsVelocityContextProviderResponse) GetRequestToolOk() ([]string, bool) {
	if o == nil || IsNil(o.RequestTool) {
		return nil, false
	}
	return o.RequestTool, true
}

// HasRequestTool returns a boolean if a field has been set.
func (o *VelocityToolsVelocityContextProviderResponse) HasRequestTool() bool {
	if o != nil && !IsNil(o.RequestTool) {
		return true
	}

	return false
}

// SetRequestTool gets a reference to the given []string and assigns it to the RequestTool field.
func (o *VelocityToolsVelocityContextProviderResponse) SetRequestTool(v []string) {
	o.RequestTool = v
}

// GetSessionTool returns the SessionTool field value if set, zero value otherwise.
func (o *VelocityToolsVelocityContextProviderResponse) GetSessionTool() []string {
	if o == nil || IsNil(o.SessionTool) {
		var ret []string
		return ret
	}
	return o.SessionTool
}

// GetSessionToolOk returns a tuple with the SessionTool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityToolsVelocityContextProviderResponse) GetSessionToolOk() ([]string, bool) {
	if o == nil || IsNil(o.SessionTool) {
		return nil, false
	}
	return o.SessionTool, true
}

// HasSessionTool returns a boolean if a field has been set.
func (o *VelocityToolsVelocityContextProviderResponse) HasSessionTool() bool {
	if o != nil && !IsNil(o.SessionTool) {
		return true
	}

	return false
}

// SetSessionTool gets a reference to the given []string and assigns it to the SessionTool field.
func (o *VelocityToolsVelocityContextProviderResponse) SetSessionTool(v []string) {
	o.SessionTool = v
}

// GetApplicationTool returns the ApplicationTool field value if set, zero value otherwise.
func (o *VelocityToolsVelocityContextProviderResponse) GetApplicationTool() []string {
	if o == nil || IsNil(o.ApplicationTool) {
		var ret []string
		return ret
	}
	return o.ApplicationTool
}

// GetApplicationToolOk returns a tuple with the ApplicationTool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityToolsVelocityContextProviderResponse) GetApplicationToolOk() ([]string, bool) {
	if o == nil || IsNil(o.ApplicationTool) {
		return nil, false
	}
	return o.ApplicationTool, true
}

// HasApplicationTool returns a boolean if a field has been set.
func (o *VelocityToolsVelocityContextProviderResponse) HasApplicationTool() bool {
	if o != nil && !IsNil(o.ApplicationTool) {
		return true
	}

	return false
}

// SetApplicationTool gets a reference to the given []string and assigns it to the ApplicationTool field.
func (o *VelocityToolsVelocityContextProviderResponse) SetApplicationTool(v []string) {
	o.ApplicationTool = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VelocityToolsVelocityContextProviderResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityToolsVelocityContextProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VelocityToolsVelocityContextProviderResponse) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VelocityToolsVelocityContextProviderResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetObjectScope returns the ObjectScope field value if set, zero value otherwise.
func (o *VelocityToolsVelocityContextProviderResponse) GetObjectScope() EnumvelocityContextProviderObjectScopeProp {
	if o == nil || IsNil(o.ObjectScope) {
		var ret EnumvelocityContextProviderObjectScopeProp
		return ret
	}
	return *o.ObjectScope
}

// GetObjectScopeOk returns a tuple with the ObjectScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityToolsVelocityContextProviderResponse) GetObjectScopeOk() (*EnumvelocityContextProviderObjectScopeProp, bool) {
	if o == nil || IsNil(o.ObjectScope) {
		return nil, false
	}
	return o.ObjectScope, true
}

// HasObjectScope returns a boolean if a field has been set.
func (o *VelocityToolsVelocityContextProviderResponse) HasObjectScope() bool {
	if o != nil && !IsNil(o.ObjectScope) {
		return true
	}

	return false
}

// SetObjectScope gets a reference to the given EnumvelocityContextProviderObjectScopeProp and assigns it to the ObjectScope field.
func (o *VelocityToolsVelocityContextProviderResponse) SetObjectScope(v EnumvelocityContextProviderObjectScopeProp) {
	o.ObjectScope = &v
}

// GetIncludedView returns the IncludedView field value if set, zero value otherwise.
func (o *VelocityToolsVelocityContextProviderResponse) GetIncludedView() []string {
	if o == nil || IsNil(o.IncludedView) {
		var ret []string
		return ret
	}
	return o.IncludedView
}

// GetIncludedViewOk returns a tuple with the IncludedView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityToolsVelocityContextProviderResponse) GetIncludedViewOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedView) {
		return nil, false
	}
	return o.IncludedView, true
}

// HasIncludedView returns a boolean if a field has been set.
func (o *VelocityToolsVelocityContextProviderResponse) HasIncludedView() bool {
	if o != nil && !IsNil(o.IncludedView) {
		return true
	}

	return false
}

// SetIncludedView gets a reference to the given []string and assigns it to the IncludedView field.
func (o *VelocityToolsVelocityContextProviderResponse) SetIncludedView(v []string) {
	o.IncludedView = v
}

// GetExcludedView returns the ExcludedView field value if set, zero value otherwise.
func (o *VelocityToolsVelocityContextProviderResponse) GetExcludedView() []string {
	if o == nil || IsNil(o.ExcludedView) {
		var ret []string
		return ret
	}
	return o.ExcludedView
}

// GetExcludedViewOk returns a tuple with the ExcludedView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityToolsVelocityContextProviderResponse) GetExcludedViewOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludedView) {
		return nil, false
	}
	return o.ExcludedView, true
}

// HasExcludedView returns a boolean if a field has been set.
func (o *VelocityToolsVelocityContextProviderResponse) HasExcludedView() bool {
	if o != nil && !IsNil(o.ExcludedView) {
		return true
	}

	return false
}

// SetExcludedView gets a reference to the given []string and assigns it to the ExcludedView field.
func (o *VelocityToolsVelocityContextProviderResponse) SetExcludedView(v []string) {
	o.ExcludedView = v
}

// GetResponseHeader returns the ResponseHeader field value if set, zero value otherwise.
func (o *VelocityToolsVelocityContextProviderResponse) GetResponseHeader() []string {
	if o == nil || IsNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityToolsVelocityContextProviderResponse) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || IsNil(o.ResponseHeader) {
		return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *VelocityToolsVelocityContextProviderResponse) HasResponseHeader() bool {
	if o != nil && !IsNil(o.ResponseHeader) {
		return true
	}

	return false
}

// SetResponseHeader gets a reference to the given []string and assigns it to the ResponseHeader field.
func (o *VelocityToolsVelocityContextProviderResponse) SetResponseHeader(v []string) {
	o.ResponseHeader = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *VelocityToolsVelocityContextProviderResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityToolsVelocityContextProviderResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *VelocityToolsVelocityContextProviderResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *VelocityToolsVelocityContextProviderResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *VelocityToolsVelocityContextProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityToolsVelocityContextProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *VelocityToolsVelocityContextProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *VelocityToolsVelocityContextProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o VelocityToolsVelocityContextProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VelocityToolsVelocityContextProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.RequestTool) {
		toSerialize["requestTool"] = o.RequestTool
	}
	if !IsNil(o.SessionTool) {
		toSerialize["sessionTool"] = o.SessionTool
	}
	if !IsNil(o.ApplicationTool) {
		toSerialize["applicationTool"] = o.ApplicationTool
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
	if !IsNil(o.ResponseHeader) {
		toSerialize["responseHeader"] = o.ResponseHeader
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableVelocityToolsVelocityContextProviderResponse struct {
	value *VelocityToolsVelocityContextProviderResponse
	isSet bool
}

func (v NullableVelocityToolsVelocityContextProviderResponse) Get() *VelocityToolsVelocityContextProviderResponse {
	return v.value
}

func (v *NullableVelocityToolsVelocityContextProviderResponse) Set(val *VelocityToolsVelocityContextProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVelocityToolsVelocityContextProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVelocityToolsVelocityContextProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVelocityToolsVelocityContextProviderResponse(val *VelocityToolsVelocityContextProviderResponse) *NullableVelocityToolsVelocityContextProviderResponse {
	return &NullableVelocityToolsVelocityContextProviderResponse{value: val, isSet: true}
}

func (v NullableVelocityToolsVelocityContextProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVelocityToolsVelocityContextProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
