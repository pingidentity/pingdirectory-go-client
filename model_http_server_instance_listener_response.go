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

// HttpServerInstanceListenerResponse struct for HttpServerInstanceListenerResponse
type HttpServerInstanceListenerResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumhttpServerInstanceListenerSchemaUrn          `json:"schemas"`
	// Name of the Server Instance Listener
	Id string `json:"id"`
	// If the server is listening on a particular address different from the hostname, then this property may be used to specify the address on which to listen for connections from HTTP clients.
	ListenAddress *string `json:"listenAddress,omitempty"`
	// The TCP port number on which the HTTP server is listening.
	ServerHTTPPort     *int32                                            `json:"serverHTTPPort,omitempty"`
	ConnectionSecurity *EnumserverInstanceListenerConnectionSecurityProp `json:"connectionSecurity,omitempty"`
	Purpose            []EnumserverInstanceListenerPurposeProp           `json:"purpose,omitempty"`
}

// NewHttpServerInstanceListenerResponse instantiates a new HttpServerInstanceListenerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpServerInstanceListenerResponse(schemas []EnumhttpServerInstanceListenerSchemaUrn, id string) *HttpServerInstanceListenerResponse {
	this := HttpServerInstanceListenerResponse{}
	this.Schemas = schemas
	this.Id = id
	return &this
}

// NewHttpServerInstanceListenerResponseWithDefaults instantiates a new HttpServerInstanceListenerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpServerInstanceListenerResponseWithDefaults() *HttpServerInstanceListenerResponse {
	this := HttpServerInstanceListenerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *HttpServerInstanceListenerResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpServerInstanceListenerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *HttpServerInstanceListenerResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *HttpServerInstanceListenerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *HttpServerInstanceListenerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpServerInstanceListenerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *HttpServerInstanceListenerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *HttpServerInstanceListenerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *HttpServerInstanceListenerResponse) GetSchemas() []EnumhttpServerInstanceListenerSchemaUrn {
	if o == nil {
		var ret []EnumhttpServerInstanceListenerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *HttpServerInstanceListenerResponse) GetSchemasOk() ([]EnumhttpServerInstanceListenerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *HttpServerInstanceListenerResponse) SetSchemas(v []EnumhttpServerInstanceListenerSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *HttpServerInstanceListenerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HttpServerInstanceListenerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HttpServerInstanceListenerResponse) SetId(v string) {
	o.Id = v
}

// GetListenAddress returns the ListenAddress field value if set, zero value otherwise.
func (o *HttpServerInstanceListenerResponse) GetListenAddress() string {
	if o == nil || isNil(o.ListenAddress) {
		var ret string
		return ret
	}
	return *o.ListenAddress
}

// GetListenAddressOk returns a tuple with the ListenAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpServerInstanceListenerResponse) GetListenAddressOk() (*string, bool) {
	if o == nil || isNil(o.ListenAddress) {
		return nil, false
	}
	return o.ListenAddress, true
}

// HasListenAddress returns a boolean if a field has been set.
func (o *HttpServerInstanceListenerResponse) HasListenAddress() bool {
	if o != nil && !isNil(o.ListenAddress) {
		return true
	}

	return false
}

// SetListenAddress gets a reference to the given string and assigns it to the ListenAddress field.
func (o *HttpServerInstanceListenerResponse) SetListenAddress(v string) {
	o.ListenAddress = &v
}

// GetServerHTTPPort returns the ServerHTTPPort field value if set, zero value otherwise.
func (o *HttpServerInstanceListenerResponse) GetServerHTTPPort() int32 {
	if o == nil || isNil(o.ServerHTTPPort) {
		var ret int32
		return ret
	}
	return *o.ServerHTTPPort
}

// GetServerHTTPPortOk returns a tuple with the ServerHTTPPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpServerInstanceListenerResponse) GetServerHTTPPortOk() (*int32, bool) {
	if o == nil || isNil(o.ServerHTTPPort) {
		return nil, false
	}
	return o.ServerHTTPPort, true
}

// HasServerHTTPPort returns a boolean if a field has been set.
func (o *HttpServerInstanceListenerResponse) HasServerHTTPPort() bool {
	if o != nil && !isNil(o.ServerHTTPPort) {
		return true
	}

	return false
}

// SetServerHTTPPort gets a reference to the given int32 and assigns it to the ServerHTTPPort field.
func (o *HttpServerInstanceListenerResponse) SetServerHTTPPort(v int32) {
	o.ServerHTTPPort = &v
}

// GetConnectionSecurity returns the ConnectionSecurity field value if set, zero value otherwise.
func (o *HttpServerInstanceListenerResponse) GetConnectionSecurity() EnumserverInstanceListenerConnectionSecurityProp {
	if o == nil || isNil(o.ConnectionSecurity) {
		var ret EnumserverInstanceListenerConnectionSecurityProp
		return ret
	}
	return *o.ConnectionSecurity
}

// GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpServerInstanceListenerResponse) GetConnectionSecurityOk() (*EnumserverInstanceListenerConnectionSecurityProp, bool) {
	if o == nil || isNil(o.ConnectionSecurity) {
		return nil, false
	}
	return o.ConnectionSecurity, true
}

// HasConnectionSecurity returns a boolean if a field has been set.
func (o *HttpServerInstanceListenerResponse) HasConnectionSecurity() bool {
	if o != nil && !isNil(o.ConnectionSecurity) {
		return true
	}

	return false
}

// SetConnectionSecurity gets a reference to the given EnumserverInstanceListenerConnectionSecurityProp and assigns it to the ConnectionSecurity field.
func (o *HttpServerInstanceListenerResponse) SetConnectionSecurity(v EnumserverInstanceListenerConnectionSecurityProp) {
	o.ConnectionSecurity = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *HttpServerInstanceListenerResponse) GetPurpose() []EnumserverInstanceListenerPurposeProp {
	if o == nil || isNil(o.Purpose) {
		var ret []EnumserverInstanceListenerPurposeProp
		return ret
	}
	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpServerInstanceListenerResponse) GetPurposeOk() ([]EnumserverInstanceListenerPurposeProp, bool) {
	if o == nil || isNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *HttpServerInstanceListenerResponse) HasPurpose() bool {
	if o != nil && !isNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given []EnumserverInstanceListenerPurposeProp and assigns it to the Purpose field.
func (o *HttpServerInstanceListenerResponse) SetPurpose(v []EnumserverInstanceListenerPurposeProp) {
	o.Purpose = v
}

func (o HttpServerInstanceListenerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ListenAddress) {
		toSerialize["listenAddress"] = o.ListenAddress
	}
	if !isNil(o.ServerHTTPPort) {
		toSerialize["serverHTTPPort"] = o.ServerHTTPPort
	}
	if !isNil(o.ConnectionSecurity) {
		toSerialize["connectionSecurity"] = o.ConnectionSecurity
	}
	if !isNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	return json.Marshal(toSerialize)
}

type NullableHttpServerInstanceListenerResponse struct {
	value *HttpServerInstanceListenerResponse
	isSet bool
}

func (v NullableHttpServerInstanceListenerResponse) Get() *HttpServerInstanceListenerResponse {
	return v.value
}

func (v *NullableHttpServerInstanceListenerResponse) Set(val *HttpServerInstanceListenerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpServerInstanceListenerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpServerInstanceListenerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpServerInstanceListenerResponse(val *HttpServerInstanceListenerResponse) *NullableHttpServerInstanceListenerResponse {
	return &NullableHttpServerInstanceListenerResponse{value: val, isSet: true}
}

func (v NullableHttpServerInstanceListenerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpServerInstanceListenerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
