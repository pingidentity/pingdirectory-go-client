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

// UserPassVaultAuthenticationMethodResponse struct for UserPassVaultAuthenticationMethodResponse
type UserPassVaultAuthenticationMethodResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Vault Authentication Method
	Id      string                                           `json:"id"`
	Schemas []EnumuserPassVaultAuthenticationMethodSchemaUrn `json:"schemas"`
	// The username for the user to authenticate.
	Username string `json:"username"`
	// The password for the user to authenticate.
	Password string `json:"password"`
	// The name used when enabling the desired UserPass authentication mechanism in the Vault server.
	LoginMechanismName *string `json:"loginMechanismName,omitempty"`
	// A description for this Vault Authentication Method
	Description *string `json:"description,omitempty"`
}

// NewUserPassVaultAuthenticationMethodResponse instantiates a new UserPassVaultAuthenticationMethodResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPassVaultAuthenticationMethodResponse(id string, schemas []EnumuserPassVaultAuthenticationMethodSchemaUrn, username string, password string) *UserPassVaultAuthenticationMethodResponse {
	this := UserPassVaultAuthenticationMethodResponse{}
	this.Id = id
	this.Schemas = schemas
	this.Username = username
	this.Password = password
	return &this
}

// NewUserPassVaultAuthenticationMethodResponseWithDefaults instantiates a new UserPassVaultAuthenticationMethodResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPassVaultAuthenticationMethodResponseWithDefaults() *UserPassVaultAuthenticationMethodResponse {
	this := UserPassVaultAuthenticationMethodResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UserPassVaultAuthenticationMethodResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPassVaultAuthenticationMethodResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UserPassVaultAuthenticationMethodResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *UserPassVaultAuthenticationMethodResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *UserPassVaultAuthenticationMethodResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPassVaultAuthenticationMethodResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *UserPassVaultAuthenticationMethodResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *UserPassVaultAuthenticationMethodResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *UserPassVaultAuthenticationMethodResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserPassVaultAuthenticationMethodResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserPassVaultAuthenticationMethodResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *UserPassVaultAuthenticationMethodResponse) GetSchemas() []EnumuserPassVaultAuthenticationMethodSchemaUrn {
	if o == nil {
		var ret []EnumuserPassVaultAuthenticationMethodSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *UserPassVaultAuthenticationMethodResponse) GetSchemasOk() ([]EnumuserPassVaultAuthenticationMethodSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *UserPassVaultAuthenticationMethodResponse) SetSchemas(v []EnumuserPassVaultAuthenticationMethodSchemaUrn) {
	o.Schemas = v
}

// GetUsername returns the Username field value
func (o *UserPassVaultAuthenticationMethodResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserPassVaultAuthenticationMethodResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserPassVaultAuthenticationMethodResponse) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *UserPassVaultAuthenticationMethodResponse) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserPassVaultAuthenticationMethodResponse) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserPassVaultAuthenticationMethodResponse) SetPassword(v string) {
	o.Password = v
}

// GetLoginMechanismName returns the LoginMechanismName field value if set, zero value otherwise.
func (o *UserPassVaultAuthenticationMethodResponse) GetLoginMechanismName() string {
	if o == nil || isNil(o.LoginMechanismName) {
		var ret string
		return ret
	}
	return *o.LoginMechanismName
}

// GetLoginMechanismNameOk returns a tuple with the LoginMechanismName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPassVaultAuthenticationMethodResponse) GetLoginMechanismNameOk() (*string, bool) {
	if o == nil || isNil(o.LoginMechanismName) {
		return nil, false
	}
	return o.LoginMechanismName, true
}

// HasLoginMechanismName returns a boolean if a field has been set.
func (o *UserPassVaultAuthenticationMethodResponse) HasLoginMechanismName() bool {
	if o != nil && !isNil(o.LoginMechanismName) {
		return true
	}

	return false
}

// SetLoginMechanismName gets a reference to the given string and assigns it to the LoginMechanismName field.
func (o *UserPassVaultAuthenticationMethodResponse) SetLoginMechanismName(v string) {
	o.LoginMechanismName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserPassVaultAuthenticationMethodResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPassVaultAuthenticationMethodResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserPassVaultAuthenticationMethodResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserPassVaultAuthenticationMethodResponse) SetDescription(v string) {
	o.Description = &v
}

func (o UserPassVaultAuthenticationMethodResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.LoginMechanismName) {
		toSerialize["loginMechanismName"] = o.LoginMechanismName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableUserPassVaultAuthenticationMethodResponse struct {
	value *UserPassVaultAuthenticationMethodResponse
	isSet bool
}

func (v NullableUserPassVaultAuthenticationMethodResponse) Get() *UserPassVaultAuthenticationMethodResponse {
	return v.value
}

func (v *NullableUserPassVaultAuthenticationMethodResponse) Set(val *UserPassVaultAuthenticationMethodResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPassVaultAuthenticationMethodResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPassVaultAuthenticationMethodResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPassVaultAuthenticationMethodResponse(val *UserPassVaultAuthenticationMethodResponse) *NullableUserPassVaultAuthenticationMethodResponse {
	return &NullableUserPassVaultAuthenticationMethodResponse{value: val, isSet: true}
}

func (v NullableUserPassVaultAuthenticationMethodResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPassVaultAuthenticationMethodResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}