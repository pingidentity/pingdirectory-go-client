/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// QuickstartHttpServletExtensionShared struct for QuickstartHttpServletExtensionShared
type QuickstartHttpServletExtensionShared struct {
	Schemas []EnumquickstartHttpServletExtensionSchemaUrn `json:"schemas"`
	// Specifies the PingFederate server to be configured.
	Server *string `json:"server,omitempty"`
	// A description for this HTTP Servlet Extension
	Description *string `json:"description,omitempty"`
	// The cross-origin request policy to use for the HTTP Servlet Extension.
	CrossOriginPolicy *string `json:"crossOriginPolicy,omitempty"`
	ResponseHeader []string `json:"responseHeader,omitempty"`
	// Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \"Correlation-Id\", \"X-Amzn-Trace-Id\", and \"X-Request-Id\".
	CorrelationIDResponseHeader *string `json:"correlationIDResponseHeader,omitempty"`
}

// NewQuickstartHttpServletExtensionShared instantiates a new QuickstartHttpServletExtensionShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickstartHttpServletExtensionShared(schemas []EnumquickstartHttpServletExtensionSchemaUrn) *QuickstartHttpServletExtensionShared {
	this := QuickstartHttpServletExtensionShared{}
	this.Schemas = schemas
	return &this
}

// NewQuickstartHttpServletExtensionSharedWithDefaults instantiates a new QuickstartHttpServletExtensionShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickstartHttpServletExtensionSharedWithDefaults() *QuickstartHttpServletExtensionShared {
	this := QuickstartHttpServletExtensionShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *QuickstartHttpServletExtensionShared) GetSchemas() []EnumquickstartHttpServletExtensionSchemaUrn {
	if o == nil {
		var ret []EnumquickstartHttpServletExtensionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *QuickstartHttpServletExtensionShared) GetSchemasOk() ([]EnumquickstartHttpServletExtensionSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *QuickstartHttpServletExtensionShared) SetSchemas(v []EnumquickstartHttpServletExtensionSchemaUrn) {
	o.Schemas = v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *QuickstartHttpServletExtensionShared) GetServer() string {
	if o == nil || isNil(o.Server) {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickstartHttpServletExtensionShared) GetServerOk() (*string, bool) {
	if o == nil || isNil(o.Server) {
    return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *QuickstartHttpServletExtensionShared) HasServer() bool {
	if o != nil && !isNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *QuickstartHttpServletExtensionShared) SetServer(v string) {
	o.Server = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *QuickstartHttpServletExtensionShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickstartHttpServletExtensionShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *QuickstartHttpServletExtensionShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *QuickstartHttpServletExtensionShared) SetDescription(v string) {
	o.Description = &v
}

// GetCrossOriginPolicy returns the CrossOriginPolicy field value if set, zero value otherwise.
func (o *QuickstartHttpServletExtensionShared) GetCrossOriginPolicy() string {
	if o == nil || isNil(o.CrossOriginPolicy) {
		var ret string
		return ret
	}
	return *o.CrossOriginPolicy
}

// GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickstartHttpServletExtensionShared) GetCrossOriginPolicyOk() (*string, bool) {
	if o == nil || isNil(o.CrossOriginPolicy) {
    return nil, false
	}
	return o.CrossOriginPolicy, true
}

// HasCrossOriginPolicy returns a boolean if a field has been set.
func (o *QuickstartHttpServletExtensionShared) HasCrossOriginPolicy() bool {
	if o != nil && !isNil(o.CrossOriginPolicy) {
		return true
	}

	return false
}

// SetCrossOriginPolicy gets a reference to the given string and assigns it to the CrossOriginPolicy field.
func (o *QuickstartHttpServletExtensionShared) SetCrossOriginPolicy(v string) {
	o.CrossOriginPolicy = &v
}

// GetResponseHeader returns the ResponseHeader field value if set, zero value otherwise.
func (o *QuickstartHttpServletExtensionShared) GetResponseHeader() []string {
	if o == nil || isNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickstartHttpServletExtensionShared) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || isNil(o.ResponseHeader) {
    return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *QuickstartHttpServletExtensionShared) HasResponseHeader() bool {
	if o != nil && !isNil(o.ResponseHeader) {
		return true
	}

	return false
}

// SetResponseHeader gets a reference to the given []string and assigns it to the ResponseHeader field.
func (o *QuickstartHttpServletExtensionShared) SetResponseHeader(v []string) {
	o.ResponseHeader = v
}

// GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field value if set, zero value otherwise.
func (o *QuickstartHttpServletExtensionShared) GetCorrelationIDResponseHeader() string {
	if o == nil || isNil(o.CorrelationIDResponseHeader) {
		var ret string
		return ret
	}
	return *o.CorrelationIDResponseHeader
}

// GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickstartHttpServletExtensionShared) GetCorrelationIDResponseHeaderOk() (*string, bool) {
	if o == nil || isNil(o.CorrelationIDResponseHeader) {
    return nil, false
	}
	return o.CorrelationIDResponseHeader, true
}

// HasCorrelationIDResponseHeader returns a boolean if a field has been set.
func (o *QuickstartHttpServletExtensionShared) HasCorrelationIDResponseHeader() bool {
	if o != nil && !isNil(o.CorrelationIDResponseHeader) {
		return true
	}

	return false
}

// SetCorrelationIDResponseHeader gets a reference to the given string and assigns it to the CorrelationIDResponseHeader field.
func (o *QuickstartHttpServletExtensionShared) SetCorrelationIDResponseHeader(v string) {
	o.CorrelationIDResponseHeader = &v
}

func (o QuickstartHttpServletExtensionShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Server) {
		toSerialize["server"] = o.Server
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

type NullableQuickstartHttpServletExtensionShared struct {
	value *QuickstartHttpServletExtensionShared
	isSet bool
}

func (v NullableQuickstartHttpServletExtensionShared) Get() *QuickstartHttpServletExtensionShared {
	return v.value
}

func (v *NullableQuickstartHttpServletExtensionShared) Set(val *QuickstartHttpServletExtensionShared) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickstartHttpServletExtensionShared) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickstartHttpServletExtensionShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickstartHttpServletExtensionShared(val *QuickstartHttpServletExtensionShared) *NullableQuickstartHttpServletExtensionShared {
	return &NullableQuickstartHttpServletExtensionShared{value: val, isSet: true}
}

func (v NullableQuickstartHttpServletExtensionShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickstartHttpServletExtensionShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


