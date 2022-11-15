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

// ConfigHttpServletExtensionResponse struct for ConfigHttpServletExtensionResponse
type ConfigHttpServletExtensionResponse struct {
	Schemas []EnumconfigHttpServletExtensionSchemaUrn `json:"schemas"`
	// Specifies the name of the identity mapper that is to be used for associating user entries with basic authentication user names.
	IdentityMapper *string `json:"identityMapper,omitempty"`
	// A description for this HTTP Servlet Extension
	Description *string `json:"description,omitempty"`
	// The cross-origin request policy to use for the HTTP Servlet Extension.
	CrossOriginPolicy *string `json:"crossOriginPolicy,omitempty"`
	ResponseHeader []string `json:"responseHeader,omitempty"`
	// Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \"Correlation-Id\", \"X-Amzn-Trace-Id\", and \"X-Request-Id\".
	CorrelationIDResponseHeader *string `json:"correlationIDResponseHeader,omitempty"`
}

// NewConfigHttpServletExtensionResponse instantiates a new ConfigHttpServletExtensionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigHttpServletExtensionResponse(schemas []EnumconfigHttpServletExtensionSchemaUrn) *ConfigHttpServletExtensionResponse {
	this := ConfigHttpServletExtensionResponse{}
	this.Schemas = schemas
	return &this
}

// NewConfigHttpServletExtensionResponseWithDefaults instantiates a new ConfigHttpServletExtensionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigHttpServletExtensionResponseWithDefaults() *ConfigHttpServletExtensionResponse {
	this := ConfigHttpServletExtensionResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ConfigHttpServletExtensionResponse) GetSchemas() []EnumconfigHttpServletExtensionSchemaUrn {
	if o == nil {
		var ret []EnumconfigHttpServletExtensionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ConfigHttpServletExtensionResponse) GetSchemasOk() ([]EnumconfigHttpServletExtensionSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ConfigHttpServletExtensionResponse) SetSchemas(v []EnumconfigHttpServletExtensionSchemaUrn) {
	o.Schemas = v
}

// GetIdentityMapper returns the IdentityMapper field value if set, zero value otherwise.
func (o *ConfigHttpServletExtensionResponse) GetIdentityMapper() string {
	if o == nil || isNil(o.IdentityMapper) {
		var ret string
		return ret
	}
	return *o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHttpServletExtensionResponse) GetIdentityMapperOk() (*string, bool) {
	if o == nil || isNil(o.IdentityMapper) {
    return nil, false
	}
	return o.IdentityMapper, true
}

// HasIdentityMapper returns a boolean if a field has been set.
func (o *ConfigHttpServletExtensionResponse) HasIdentityMapper() bool {
	if o != nil && !isNil(o.IdentityMapper) {
		return true
	}

	return false
}

// SetIdentityMapper gets a reference to the given string and assigns it to the IdentityMapper field.
func (o *ConfigHttpServletExtensionResponse) SetIdentityMapper(v string) {
	o.IdentityMapper = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigHttpServletExtensionResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHttpServletExtensionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigHttpServletExtensionResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConfigHttpServletExtensionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCrossOriginPolicy returns the CrossOriginPolicy field value if set, zero value otherwise.
func (o *ConfigHttpServletExtensionResponse) GetCrossOriginPolicy() string {
	if o == nil || isNil(o.CrossOriginPolicy) {
		var ret string
		return ret
	}
	return *o.CrossOriginPolicy
}

// GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool) {
	if o == nil || isNil(o.CrossOriginPolicy) {
    return nil, false
	}
	return o.CrossOriginPolicy, true
}

// HasCrossOriginPolicy returns a boolean if a field has been set.
func (o *ConfigHttpServletExtensionResponse) HasCrossOriginPolicy() bool {
	if o != nil && !isNil(o.CrossOriginPolicy) {
		return true
	}

	return false
}

// SetCrossOriginPolicy gets a reference to the given string and assigns it to the CrossOriginPolicy field.
func (o *ConfigHttpServletExtensionResponse) SetCrossOriginPolicy(v string) {
	o.CrossOriginPolicy = &v
}

// GetResponseHeader returns the ResponseHeader field value if set, zero value otherwise.
func (o *ConfigHttpServletExtensionResponse) GetResponseHeader() []string {
	if o == nil || isNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHttpServletExtensionResponse) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || isNil(o.ResponseHeader) {
    return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *ConfigHttpServletExtensionResponse) HasResponseHeader() bool {
	if o != nil && !isNil(o.ResponseHeader) {
		return true
	}

	return false
}

// SetResponseHeader gets a reference to the given []string and assigns it to the ResponseHeader field.
func (o *ConfigHttpServletExtensionResponse) SetResponseHeader(v []string) {
	o.ResponseHeader = v
}

// GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field value if set, zero value otherwise.
func (o *ConfigHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string {
	if o == nil || isNil(o.CorrelationIDResponseHeader) {
		var ret string
		return ret
	}
	return *o.CorrelationIDResponseHeader
}

// GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool) {
	if o == nil || isNil(o.CorrelationIDResponseHeader) {
    return nil, false
	}
	return o.CorrelationIDResponseHeader, true
}

// HasCorrelationIDResponseHeader returns a boolean if a field has been set.
func (o *ConfigHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool {
	if o != nil && !isNil(o.CorrelationIDResponseHeader) {
		return true
	}

	return false
}

// SetCorrelationIDResponseHeader gets a reference to the given string and assigns it to the CorrelationIDResponseHeader field.
func (o *ConfigHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string) {
	o.CorrelationIDResponseHeader = &v
}

func (o ConfigHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.IdentityMapper) {
		toSerialize["identityMapper"] = o.IdentityMapper
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

type NullableConfigHttpServletExtensionResponse struct {
	value *ConfigHttpServletExtensionResponse
	isSet bool
}

func (v NullableConfigHttpServletExtensionResponse) Get() *ConfigHttpServletExtensionResponse {
	return v.value
}

func (v *NullableConfigHttpServletExtensionResponse) Set(val *ConfigHttpServletExtensionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigHttpServletExtensionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigHttpServletExtensionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigHttpServletExtensionResponse(val *ConfigHttpServletExtensionResponse) *NullableConfigHttpServletExtensionResponse {
	return &NullableConfigHttpServletExtensionResponse{value: val, isSet: true}
}

func (v NullableConfigHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigHttpServletExtensionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


