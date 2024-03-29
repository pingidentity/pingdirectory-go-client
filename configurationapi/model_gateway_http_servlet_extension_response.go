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

// checks if the GatewayHttpServletExtensionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayHttpServletExtensionResponse{}

// GatewayHttpServletExtensionResponse struct for GatewayHttpServletExtensionResponse
type GatewayHttpServletExtensionResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumgatewayHttpServletExtensionSchemaUrn         `json:"schemas"`
	// Name of the HTTP Servlet Extension
	Id string `json:"id"`
	// Specifies any Gateway API Endpoints that will not be handled by the Gateway HTTP Servlet Extension.
	ExcludedAPIEndpoint []string `json:"excludedAPIEndpoint,omitempty"`
	// The maximum number of bytes allowed per request body.
	RequestLimit *string `json:"requestLimit,omitempty"`
	// The maximum number of bytes allowed per response body.
	ResponseLimit *string `json:"responseLimit,omitempty"`
	// The number of threads used to forward responses to the API backend.
	NumForwardThreads *int64 `json:"numForwardThreads,omitempty"`
	// A description for this HTTP Servlet Extension
	Description *string `json:"description,omitempty"`
	// The cross-origin request policy to use for the HTTP Servlet Extension.
	CrossOriginPolicy *string `json:"crossOriginPolicy,omitempty"`
	// Specifies HTTP header fields and values added to response headers for all requests.
	ResponseHeader []string `json:"responseHeader,omitempty"`
}

// NewGatewayHttpServletExtensionResponse instantiates a new GatewayHttpServletExtensionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayHttpServletExtensionResponse(schemas []EnumgatewayHttpServletExtensionSchemaUrn, id string) *GatewayHttpServletExtensionResponse {
	this := GatewayHttpServletExtensionResponse{}
	this.Schemas = schemas
	this.Id = id
	return &this
}

// NewGatewayHttpServletExtensionResponseWithDefaults instantiates a new GatewayHttpServletExtensionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayHttpServletExtensionResponseWithDefaults() *GatewayHttpServletExtensionResponse {
	this := GatewayHttpServletExtensionResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GatewayHttpServletExtensionResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GatewayHttpServletExtensionResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GatewayHttpServletExtensionResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GatewayHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GatewayHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GatewayHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *GatewayHttpServletExtensionResponse) GetSchemas() []EnumgatewayHttpServletExtensionSchemaUrn {
	if o == nil {
		var ret []EnumgatewayHttpServletExtensionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GatewayHttpServletExtensionResponse) GetSchemasOk() ([]EnumgatewayHttpServletExtensionSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GatewayHttpServletExtensionResponse) SetSchemas(v []EnumgatewayHttpServletExtensionSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *GatewayHttpServletExtensionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GatewayHttpServletExtensionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GatewayHttpServletExtensionResponse) SetId(v string) {
	o.Id = v
}

// GetExcludedAPIEndpoint returns the ExcludedAPIEndpoint field value if set, zero value otherwise.
func (o *GatewayHttpServletExtensionResponse) GetExcludedAPIEndpoint() []string {
	if o == nil || IsNil(o.ExcludedAPIEndpoint) {
		var ret []string
		return ret
	}
	return o.ExcludedAPIEndpoint
}

// GetExcludedAPIEndpointOk returns a tuple with the ExcludedAPIEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayHttpServletExtensionResponse) GetExcludedAPIEndpointOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludedAPIEndpoint) {
		return nil, false
	}
	return o.ExcludedAPIEndpoint, true
}

// HasExcludedAPIEndpoint returns a boolean if a field has been set.
func (o *GatewayHttpServletExtensionResponse) HasExcludedAPIEndpoint() bool {
	if o != nil && !IsNil(o.ExcludedAPIEndpoint) {
		return true
	}

	return false
}

// SetExcludedAPIEndpoint gets a reference to the given []string and assigns it to the ExcludedAPIEndpoint field.
func (o *GatewayHttpServletExtensionResponse) SetExcludedAPIEndpoint(v []string) {
	o.ExcludedAPIEndpoint = v
}

// GetRequestLimit returns the RequestLimit field value if set, zero value otherwise.
func (o *GatewayHttpServletExtensionResponse) GetRequestLimit() string {
	if o == nil || IsNil(o.RequestLimit) {
		var ret string
		return ret
	}
	return *o.RequestLimit
}

// GetRequestLimitOk returns a tuple with the RequestLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayHttpServletExtensionResponse) GetRequestLimitOk() (*string, bool) {
	if o == nil || IsNil(o.RequestLimit) {
		return nil, false
	}
	return o.RequestLimit, true
}

// HasRequestLimit returns a boolean if a field has been set.
func (o *GatewayHttpServletExtensionResponse) HasRequestLimit() bool {
	if o != nil && !IsNil(o.RequestLimit) {
		return true
	}

	return false
}

// SetRequestLimit gets a reference to the given string and assigns it to the RequestLimit field.
func (o *GatewayHttpServletExtensionResponse) SetRequestLimit(v string) {
	o.RequestLimit = &v
}

// GetResponseLimit returns the ResponseLimit field value if set, zero value otherwise.
func (o *GatewayHttpServletExtensionResponse) GetResponseLimit() string {
	if o == nil || IsNil(o.ResponseLimit) {
		var ret string
		return ret
	}
	return *o.ResponseLimit
}

// GetResponseLimitOk returns a tuple with the ResponseLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayHttpServletExtensionResponse) GetResponseLimitOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseLimit) {
		return nil, false
	}
	return o.ResponseLimit, true
}

// HasResponseLimit returns a boolean if a field has been set.
func (o *GatewayHttpServletExtensionResponse) HasResponseLimit() bool {
	if o != nil && !IsNil(o.ResponseLimit) {
		return true
	}

	return false
}

// SetResponseLimit gets a reference to the given string and assigns it to the ResponseLimit field.
func (o *GatewayHttpServletExtensionResponse) SetResponseLimit(v string) {
	o.ResponseLimit = &v
}

// GetNumForwardThreads returns the NumForwardThreads field value if set, zero value otherwise.
func (o *GatewayHttpServletExtensionResponse) GetNumForwardThreads() int64 {
	if o == nil || IsNil(o.NumForwardThreads) {
		var ret int64
		return ret
	}
	return *o.NumForwardThreads
}

// GetNumForwardThreadsOk returns a tuple with the NumForwardThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayHttpServletExtensionResponse) GetNumForwardThreadsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumForwardThreads) {
		return nil, false
	}
	return o.NumForwardThreads, true
}

// HasNumForwardThreads returns a boolean if a field has been set.
func (o *GatewayHttpServletExtensionResponse) HasNumForwardThreads() bool {
	if o != nil && !IsNil(o.NumForwardThreads) {
		return true
	}

	return false
}

// SetNumForwardThreads gets a reference to the given int64 and assigns it to the NumForwardThreads field.
func (o *GatewayHttpServletExtensionResponse) SetNumForwardThreads(v int64) {
	o.NumForwardThreads = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GatewayHttpServletExtensionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayHttpServletExtensionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GatewayHttpServletExtensionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GatewayHttpServletExtensionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCrossOriginPolicy returns the CrossOriginPolicy field value if set, zero value otherwise.
func (o *GatewayHttpServletExtensionResponse) GetCrossOriginPolicy() string {
	if o == nil || IsNil(o.CrossOriginPolicy) {
		var ret string
		return ret
	}
	return *o.CrossOriginPolicy
}

// GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.CrossOriginPolicy) {
		return nil, false
	}
	return o.CrossOriginPolicy, true
}

// HasCrossOriginPolicy returns a boolean if a field has been set.
func (o *GatewayHttpServletExtensionResponse) HasCrossOriginPolicy() bool {
	if o != nil && !IsNil(o.CrossOriginPolicy) {
		return true
	}

	return false
}

// SetCrossOriginPolicy gets a reference to the given string and assigns it to the CrossOriginPolicy field.
func (o *GatewayHttpServletExtensionResponse) SetCrossOriginPolicy(v string) {
	o.CrossOriginPolicy = &v
}

// GetResponseHeader returns the ResponseHeader field value if set, zero value otherwise.
func (o *GatewayHttpServletExtensionResponse) GetResponseHeader() []string {
	if o == nil || IsNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayHttpServletExtensionResponse) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || IsNil(o.ResponseHeader) {
		return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *GatewayHttpServletExtensionResponse) HasResponseHeader() bool {
	if o != nil && !IsNil(o.ResponseHeader) {
		return true
	}

	return false
}

// SetResponseHeader gets a reference to the given []string and assigns it to the ResponseHeader field.
func (o *GatewayHttpServletExtensionResponse) SetResponseHeader(v []string) {
	o.ResponseHeader = v
}

func (o GatewayHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayHttpServletExtensionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.ExcludedAPIEndpoint) {
		toSerialize["excludedAPIEndpoint"] = o.ExcludedAPIEndpoint
	}
	if !IsNil(o.RequestLimit) {
		toSerialize["requestLimit"] = o.RequestLimit
	}
	if !IsNil(o.ResponseLimit) {
		toSerialize["responseLimit"] = o.ResponseLimit
	}
	if !IsNil(o.NumForwardThreads) {
		toSerialize["numForwardThreads"] = o.NumForwardThreads
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
	return toSerialize, nil
}

type NullableGatewayHttpServletExtensionResponse struct {
	value *GatewayHttpServletExtensionResponse
	isSet bool
}

func (v NullableGatewayHttpServletExtensionResponse) Get() *GatewayHttpServletExtensionResponse {
	return v.value
}

func (v *NullableGatewayHttpServletExtensionResponse) Set(val *GatewayHttpServletExtensionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayHttpServletExtensionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayHttpServletExtensionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayHttpServletExtensionResponse(val *GatewayHttpServletExtensionResponse) *NullableGatewayHttpServletExtensionResponse {
	return &NullableGatewayHttpServletExtensionResponse{value: val, isSet: true}
}

func (v NullableGatewayHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayHttpServletExtensionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
