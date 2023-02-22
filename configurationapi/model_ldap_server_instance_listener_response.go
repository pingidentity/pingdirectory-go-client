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

// LdapServerInstanceListenerResponse struct for LdapServerInstanceListenerResponse
type LdapServerInstanceListenerResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumldapServerInstanceListenerSchemaUrn          `json:"schemas"`
	// Name of the Server Instance Listener
	Id string `json:"id"`
	// The TCP port number on which the LDAP server is listening.
	ServerLDAPPort     *int32                                            `json:"serverLDAPPort,omitempty"`
	ConnectionSecurity *EnumserverInstanceListenerConnectionSecurityProp `json:"connectionSecurity,omitempty"`
	// The public component of the certificate that the listener is expected to present to clients. When establishing a connection to this server, only the certificate(s) listed here will be trusted.
	ListenerCertificate *string                                 `json:"listenerCertificate,omitempty"`
	Purpose             []EnumserverInstanceListenerPurposeProp `json:"purpose,omitempty"`
}

// NewLdapServerInstanceListenerResponse instantiates a new LdapServerInstanceListenerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapServerInstanceListenerResponse(schemas []EnumldapServerInstanceListenerSchemaUrn, id string) *LdapServerInstanceListenerResponse {
	this := LdapServerInstanceListenerResponse{}
	this.Schemas = schemas
	this.Id = id
	return &this
}

// NewLdapServerInstanceListenerResponseWithDefaults instantiates a new LdapServerInstanceListenerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapServerInstanceListenerResponseWithDefaults() *LdapServerInstanceListenerResponse {
	this := LdapServerInstanceListenerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LdapServerInstanceListenerResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapServerInstanceListenerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LdapServerInstanceListenerResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *LdapServerInstanceListenerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *LdapServerInstanceListenerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapServerInstanceListenerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *LdapServerInstanceListenerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *LdapServerInstanceListenerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *LdapServerInstanceListenerResponse) GetSchemas() []EnumldapServerInstanceListenerSchemaUrn {
	if o == nil {
		var ret []EnumldapServerInstanceListenerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *LdapServerInstanceListenerResponse) GetSchemasOk() ([]EnumldapServerInstanceListenerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *LdapServerInstanceListenerResponse) SetSchemas(v []EnumldapServerInstanceListenerSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *LdapServerInstanceListenerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LdapServerInstanceListenerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LdapServerInstanceListenerResponse) SetId(v string) {
	o.Id = v
}

// GetServerLDAPPort returns the ServerLDAPPort field value if set, zero value otherwise.
func (o *LdapServerInstanceListenerResponse) GetServerLDAPPort() int32 {
	if o == nil || isNil(o.ServerLDAPPort) {
		var ret int32
		return ret
	}
	return *o.ServerLDAPPort
}

// GetServerLDAPPortOk returns a tuple with the ServerLDAPPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapServerInstanceListenerResponse) GetServerLDAPPortOk() (*int32, bool) {
	if o == nil || isNil(o.ServerLDAPPort) {
		return nil, false
	}
	return o.ServerLDAPPort, true
}

// HasServerLDAPPort returns a boolean if a field has been set.
func (o *LdapServerInstanceListenerResponse) HasServerLDAPPort() bool {
	if o != nil && !isNil(o.ServerLDAPPort) {
		return true
	}

	return false
}

// SetServerLDAPPort gets a reference to the given int32 and assigns it to the ServerLDAPPort field.
func (o *LdapServerInstanceListenerResponse) SetServerLDAPPort(v int32) {
	o.ServerLDAPPort = &v
}

// GetConnectionSecurity returns the ConnectionSecurity field value if set, zero value otherwise.
func (o *LdapServerInstanceListenerResponse) GetConnectionSecurity() EnumserverInstanceListenerConnectionSecurityProp {
	if o == nil || isNil(o.ConnectionSecurity) {
		var ret EnumserverInstanceListenerConnectionSecurityProp
		return ret
	}
	return *o.ConnectionSecurity
}

// GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapServerInstanceListenerResponse) GetConnectionSecurityOk() (*EnumserverInstanceListenerConnectionSecurityProp, bool) {
	if o == nil || isNil(o.ConnectionSecurity) {
		return nil, false
	}
	return o.ConnectionSecurity, true
}

// HasConnectionSecurity returns a boolean if a field has been set.
func (o *LdapServerInstanceListenerResponse) HasConnectionSecurity() bool {
	if o != nil && !isNil(o.ConnectionSecurity) {
		return true
	}

	return false
}

// SetConnectionSecurity gets a reference to the given EnumserverInstanceListenerConnectionSecurityProp and assigns it to the ConnectionSecurity field.
func (o *LdapServerInstanceListenerResponse) SetConnectionSecurity(v EnumserverInstanceListenerConnectionSecurityProp) {
	o.ConnectionSecurity = &v
}

// GetListenerCertificate returns the ListenerCertificate field value if set, zero value otherwise.
func (o *LdapServerInstanceListenerResponse) GetListenerCertificate() string {
	if o == nil || isNil(o.ListenerCertificate) {
		var ret string
		return ret
	}
	return *o.ListenerCertificate
}

// GetListenerCertificateOk returns a tuple with the ListenerCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapServerInstanceListenerResponse) GetListenerCertificateOk() (*string, bool) {
	if o == nil || isNil(o.ListenerCertificate) {
		return nil, false
	}
	return o.ListenerCertificate, true
}

// HasListenerCertificate returns a boolean if a field has been set.
func (o *LdapServerInstanceListenerResponse) HasListenerCertificate() bool {
	if o != nil && !isNil(o.ListenerCertificate) {
		return true
	}

	return false
}

// SetListenerCertificate gets a reference to the given string and assigns it to the ListenerCertificate field.
func (o *LdapServerInstanceListenerResponse) SetListenerCertificate(v string) {
	o.ListenerCertificate = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *LdapServerInstanceListenerResponse) GetPurpose() []EnumserverInstanceListenerPurposeProp {
	if o == nil || isNil(o.Purpose) {
		var ret []EnumserverInstanceListenerPurposeProp
		return ret
	}
	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapServerInstanceListenerResponse) GetPurposeOk() ([]EnumserverInstanceListenerPurposeProp, bool) {
	if o == nil || isNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *LdapServerInstanceListenerResponse) HasPurpose() bool {
	if o != nil && !isNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given []EnumserverInstanceListenerPurposeProp and assigns it to the Purpose field.
func (o *LdapServerInstanceListenerResponse) SetPurpose(v []EnumserverInstanceListenerPurposeProp) {
	o.Purpose = v
}

func (o LdapServerInstanceListenerResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.ServerLDAPPort) {
		toSerialize["serverLDAPPort"] = o.ServerLDAPPort
	}
	if !isNil(o.ConnectionSecurity) {
		toSerialize["connectionSecurity"] = o.ConnectionSecurity
	}
	if !isNil(o.ListenerCertificate) {
		toSerialize["listenerCertificate"] = o.ListenerCertificate
	}
	if !isNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	return json.Marshal(toSerialize)
}

type NullableLdapServerInstanceListenerResponse struct {
	value *LdapServerInstanceListenerResponse
	isSet bool
}

func (v NullableLdapServerInstanceListenerResponse) Get() *LdapServerInstanceListenerResponse {
	return v.value
}

func (v *NullableLdapServerInstanceListenerResponse) Set(val *LdapServerInstanceListenerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapServerInstanceListenerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapServerInstanceListenerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapServerInstanceListenerResponse(val *LdapServerInstanceListenerResponse) *NullableLdapServerInstanceListenerResponse {
	return &NullableLdapServerInstanceListenerResponse{value: val, isSet: true}
}

func (v NullableLdapServerInstanceListenerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapServerInstanceListenerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}