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

// DigestMd5SaslMechanismHandlerResponse struct for DigestMd5SaslMechanismHandlerResponse
type DigestMd5SaslMechanismHandlerResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas []EnumdigestMd5SaslMechanismHandlerSchemaUrn `json:"schemas"`
	// Name of the SASL Mechanism Handler
	Id *string `json:"id,omitempty"`
	// Specifies the realm that is to be used by the server for DIGEST-MD5 authentication.
	Realm *string `json:"realm,omitempty"`
	// Specifies the name of the identity mapper that is to be used with this SASL mechanism handler to match the authentication or authorization ID included in the SASL bind request to the corresponding user in the directory.
	IdentityMapper string `json:"identityMapper"`
	// Specifies the DNS-resolvable fully-qualified domain name for the server that is used when validating the digest-uri parameter during the authentication process.
	ServerFqdn *string `json:"serverFqdn,omitempty"`
	// A description for this SASL Mechanism Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the SASL mechanism handler is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewDigestMd5SaslMechanismHandlerResponse instantiates a new DigestMd5SaslMechanismHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigestMd5SaslMechanismHandlerResponse(schemas []EnumdigestMd5SaslMechanismHandlerSchemaUrn, identityMapper string, enabled bool) *DigestMd5SaslMechanismHandlerResponse {
	this := DigestMd5SaslMechanismHandlerResponse{}
	this.Schemas = schemas
	this.IdentityMapper = identityMapper
	this.Enabled = enabled
	return &this
}

// NewDigestMd5SaslMechanismHandlerResponseWithDefaults instantiates a new DigestMd5SaslMechanismHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigestMd5SaslMechanismHandlerResponseWithDefaults() *DigestMd5SaslMechanismHandlerResponse {
	this := DigestMd5SaslMechanismHandlerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DigestMd5SaslMechanismHandlerResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *DigestMd5SaslMechanismHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *DigestMd5SaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *DigestMd5SaslMechanismHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *DigestMd5SaslMechanismHandlerResponse) GetSchemas() []EnumdigestMd5SaslMechanismHandlerSchemaUrn {
	if o == nil {
		var ret []EnumdigestMd5SaslMechanismHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) GetSchemasOk() ([]EnumdigestMd5SaslMechanismHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *DigestMd5SaslMechanismHandlerResponse) SetSchemas(v []EnumdigestMd5SaslMechanismHandlerSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DigestMd5SaslMechanismHandlerResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DigestMd5SaslMechanismHandlerResponse) SetId(v string) {
	o.Id = &v
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *DigestMd5SaslMechanismHandlerResponse) GetRealm() string {
	if o == nil || isNil(o.Realm) {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) GetRealmOk() (*string, bool) {
	if o == nil || isNil(o.Realm) {
    return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) HasRealm() bool {
	if o != nil && !isNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *DigestMd5SaslMechanismHandlerResponse) SetRealm(v string) {
	o.Realm = &v
}

// GetIdentityMapper returns the IdentityMapper field value
func (o *DigestMd5SaslMechanismHandlerResponse) GetIdentityMapper() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value
// and a boolean to check if the value has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) GetIdentityMapperOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IdentityMapper, true
}

// SetIdentityMapper sets field value
func (o *DigestMd5SaslMechanismHandlerResponse) SetIdentityMapper(v string) {
	o.IdentityMapper = v
}

// GetServerFqdn returns the ServerFqdn field value if set, zero value otherwise.
func (o *DigestMd5SaslMechanismHandlerResponse) GetServerFqdn() string {
	if o == nil || isNil(o.ServerFqdn) {
		var ret string
		return ret
	}
	return *o.ServerFqdn
}

// GetServerFqdnOk returns a tuple with the ServerFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) GetServerFqdnOk() (*string, bool) {
	if o == nil || isNil(o.ServerFqdn) {
    return nil, false
	}
	return o.ServerFqdn, true
}

// HasServerFqdn returns a boolean if a field has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) HasServerFqdn() bool {
	if o != nil && !isNil(o.ServerFqdn) {
		return true
	}

	return false
}

// SetServerFqdn gets a reference to the given string and assigns it to the ServerFqdn field.
func (o *DigestMd5SaslMechanismHandlerResponse) SetServerFqdn(v string) {
	o.ServerFqdn = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DigestMd5SaslMechanismHandlerResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DigestMd5SaslMechanismHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *DigestMd5SaslMechanismHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DigestMd5SaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DigestMd5SaslMechanismHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o DigestMd5SaslMechanismHandlerResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	if true {
		toSerialize["identityMapper"] = o.IdentityMapper
	}
	if !isNil(o.ServerFqdn) {
		toSerialize["serverFqdn"] = o.ServerFqdn
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableDigestMd5SaslMechanismHandlerResponse struct {
	value *DigestMd5SaslMechanismHandlerResponse
	isSet bool
}

func (v NullableDigestMd5SaslMechanismHandlerResponse) Get() *DigestMd5SaslMechanismHandlerResponse {
	return v.value
}

func (v *NullableDigestMd5SaslMechanismHandlerResponse) Set(val *DigestMd5SaslMechanismHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDigestMd5SaslMechanismHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDigestMd5SaslMechanismHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigestMd5SaslMechanismHandlerResponse(val *DigestMd5SaslMechanismHandlerResponse) *NullableDigestMd5SaslMechanismHandlerResponse {
	return &NullableDigestMd5SaslMechanismHandlerResponse{value: val, isSet: true}
}

func (v NullableDigestMd5SaslMechanismHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigestMd5SaslMechanismHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


