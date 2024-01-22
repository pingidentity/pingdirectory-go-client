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

// checks if the StandardHttpServletExtensionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardHttpServletExtensionResponse{}

// StandardHttpServletExtensionResponse struct for StandardHttpServletExtensionResponse
type StandardHttpServletExtensionResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumstandardHttpServletExtensionSchemaUrn        `json:"schemas"`
	// Name of the HTTP Servlet Extension
	Id string `json:"id"`
	// A description for this HTTP Servlet Extension
	Description *string `json:"description,omitempty"`
	// The cross-origin request policy to use for the HTTP Servlet Extension.
	CrossOriginPolicy *string `json:"crossOriginPolicy,omitempty"`
	// Specifies HTTP header fields and values added to response headers for all requests.
	ResponseHeader []string `json:"responseHeader,omitempty"`
	// Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \"Correlation-Id\", \"X-Amzn-Trace-Id\", and \"X-Request-Id\".
	CorrelationIDResponseHeader *string `json:"correlationIDResponseHeader,omitempty"`
}

// NewStandardHttpServletExtensionResponse instantiates a new StandardHttpServletExtensionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardHttpServletExtensionResponse(schemas []EnumstandardHttpServletExtensionSchemaUrn, id string) *StandardHttpServletExtensionResponse {
	this := StandardHttpServletExtensionResponse{}
	this.Schemas = schemas
	this.Id = id
	return &this
}

// NewStandardHttpServletExtensionResponseWithDefaults instantiates a new StandardHttpServletExtensionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardHttpServletExtensionResponseWithDefaults() *StandardHttpServletExtensionResponse {
	this := StandardHttpServletExtensionResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *StandardHttpServletExtensionResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *StandardHttpServletExtensionResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *StandardHttpServletExtensionResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *StandardHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *StandardHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *StandardHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *StandardHttpServletExtensionResponse) GetSchemas() []EnumstandardHttpServletExtensionSchemaUrn {
	if o == nil {
		var ret []EnumstandardHttpServletExtensionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *StandardHttpServletExtensionResponse) GetSchemasOk() ([]EnumstandardHttpServletExtensionSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *StandardHttpServletExtensionResponse) SetSchemas(v []EnumstandardHttpServletExtensionSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *StandardHttpServletExtensionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StandardHttpServletExtensionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StandardHttpServletExtensionResponse) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StandardHttpServletExtensionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardHttpServletExtensionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StandardHttpServletExtensionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StandardHttpServletExtensionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCrossOriginPolicy returns the CrossOriginPolicy field value if set, zero value otherwise.
func (o *StandardHttpServletExtensionResponse) GetCrossOriginPolicy() string {
	if o == nil || IsNil(o.CrossOriginPolicy) {
		var ret string
		return ret
	}
	return *o.CrossOriginPolicy
}

// GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.CrossOriginPolicy) {
		return nil, false
	}
	return o.CrossOriginPolicy, true
}

// HasCrossOriginPolicy returns a boolean if a field has been set.
func (o *StandardHttpServletExtensionResponse) HasCrossOriginPolicy() bool {
	if o != nil && !IsNil(o.CrossOriginPolicy) {
		return true
	}

	return false
}

// SetCrossOriginPolicy gets a reference to the given string and assigns it to the CrossOriginPolicy field.
func (o *StandardHttpServletExtensionResponse) SetCrossOriginPolicy(v string) {
	o.CrossOriginPolicy = &v
}

// GetResponseHeader returns the ResponseHeader field value if set, zero value otherwise.
func (o *StandardHttpServletExtensionResponse) GetResponseHeader() []string {
	if o == nil || IsNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardHttpServletExtensionResponse) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || IsNil(o.ResponseHeader) {
		return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *StandardHttpServletExtensionResponse) HasResponseHeader() bool {
	if o != nil && !IsNil(o.ResponseHeader) {
		return true
	}

	return false
}

// SetResponseHeader gets a reference to the given []string and assigns it to the ResponseHeader field.
func (o *StandardHttpServletExtensionResponse) SetResponseHeader(v []string) {
	o.ResponseHeader = v
}

// GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field value if set, zero value otherwise.
func (o *StandardHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string {
	if o == nil || IsNil(o.CorrelationIDResponseHeader) {
		var ret string
		return ret
	}
	return *o.CorrelationIDResponseHeader
}

// GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationIDResponseHeader) {
		return nil, false
	}
	return o.CorrelationIDResponseHeader, true
}

// HasCorrelationIDResponseHeader returns a boolean if a field has been set.
func (o *StandardHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool {
	if o != nil && !IsNil(o.CorrelationIDResponseHeader) {
		return true
	}

	return false
}

// SetCorrelationIDResponseHeader gets a reference to the given string and assigns it to the CorrelationIDResponseHeader field.
func (o *StandardHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string) {
	o.CorrelationIDResponseHeader = &v
}

func (o StandardHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandardHttpServletExtensionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
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
	return toSerialize, nil
}

type NullableStandardHttpServletExtensionResponse struct {
	value *StandardHttpServletExtensionResponse
	isSet bool
}

func (v NullableStandardHttpServletExtensionResponse) Get() *StandardHttpServletExtensionResponse {
	return v.value
}

func (v *NullableStandardHttpServletExtensionResponse) Set(val *StandardHttpServletExtensionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardHttpServletExtensionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardHttpServletExtensionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardHttpServletExtensionResponse(val *StandardHttpServletExtensionResponse) *NullableStandardHttpServletExtensionResponse {
	return &NullableStandardHttpServletExtensionResponse{value: val, isSet: true}
}

func (v NullableStandardHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardHttpServletExtensionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
